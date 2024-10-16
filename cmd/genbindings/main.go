package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

const (
	ClangSubprocessCount = 3
	BaseModule           = "github.com/mappu/miqt"
)

func cacheFilePath(inputHeader string) string {
	return filepath.Join("cachedir", strings.Replace(inputHeader, `/`, `__`, -1)+".json")
}

func findHeadersInDir(srcDir string) []string {
	content, err := os.ReadDir(srcDir)
	if err != nil {
		panic(err)
	}

	var ret []string

	for _, includeFile := range content {
		if includeFile.IsDir() {
			continue
		}
		if !strings.HasSuffix(includeFile.Name(), `.h`) {
			continue
		}
		fullPath := filepath.Join(srcDir, includeFile.Name())
		if !AllowHeader(fullPath) {
			continue
		}
		ret = append(ret, fullPath)
	}

	return ret
}

func cleanGeneratedFilesInDir(dirpath string) {
	log.Printf("Cleaning up output directory %q...", dirpath)

	_ = os.MkdirAll(dirpath, 0755)

	existing, err := os.ReadDir(dirpath)
	if err != nil {
		panic(err)
	}

	cleaned := 0
	for _, e := range existing {
		if e.IsDir() {
			continue
		}
		if !strings.HasPrefix(e.Name(), `gen_`) {
			continue
		}
		// One of ours, clean up
		err := os.Remove(filepath.Join(dirpath, e.Name()))
		if err != nil {
			log.Printf("WARNING: Failed to remove existing file %q", e.Name())
			continue
		}

		cleaned++
	}

	log.Printf("Removed %d file(s).", cleaned)
}

func main() {
	clang := flag.String("clang", "clang", "Custom path to clang")
	cflags := flag.String("cflags", `-DQT_WIDGETS_LIB -I/usr/include/x86_64-linux-gnu/qt5/QtWidgets -I/usr/include/x86_64-linux-gnu/qt5 -I/usr/include/x86_64-linux-gnu/qt5/QtCore -DQT_GUI_LIB -I/usr/include/x86_64-linux-gnu/qt5/QtGui -DQT_CORE_LIB`, "Cflags to pass to clang (e.g. `pkg-config --cflags Qt5Widgets`)")
	outDir := flag.String("outdir", "../../qt", "Output directory for generated gen_** files")

	flag.Parse()

	var includeFiles []string

	for _, srcDir := range []string{
		"/usr/include/x86_64-linux-gnu/qt5/QtCore",
		"/usr/include/x86_64-linux-gnu/qt5/QtGui",
		"/usr/include/x86_64-linux-gnu/qt5/QtWidgets",
	} {
		includeFiles = append(includeFiles, findHeadersInDir(srcDir)...)
	}

	log.Printf("Found %d header files to process.", len(includeFiles))

	cleanGeneratedFilesInDir(*outDir)

	var processHeaders []*CppParsedHeader
	atr := astTransformRedundant{
		preserve: make(map[string]*CppEnum),
	}

	InsertTypedefs()

	//
	// PASS 0 (Fill clang cache)
	//

	generateClangCaches(includeFiles, *clang, strings.Fields(*cflags))

	// The cache should now be fully populated.

	//
	// PASS 1 (clang2il)
	//

	for _, inputHeader := range includeFiles {

		cacheFile := cacheFilePath(inputHeader)

		astJson, err := ioutil.ReadFile(cacheFile)
		if err != nil {
			panic("Expected cache to be created for " + inputHeader + ", but got error " + err.Error())
		}

		// Json decode
		var astInner []interface{} = nil
		err = json.Unmarshal(astJson, &astInner)
		if err != nil {
			panic(err)
		}

		if astInner == nil {
			panic("Null in cache file for " + inputHeader)
		}

		// Convert it to our intermediate format
		parsed, err := parseHeader(astInner, "")
		if err != nil {
			panic(err)
		}

		parsed.Filename = inputHeader // Stash

		// AST transforms on our IL
		astTransformChildClasses(parsed) // must be first
		astTransformOptional(parsed)
		astTransformOverloads(parsed)
		atr.Process(parsed)

		// Update global state tracker (AFTER astTransformChildClasses)
		// Currently, this is only used for inner classes
		for _, c := range parsed.Classes {
			KnownClassnames[c.ClassName] = struct{}{}
		}
		for _, td := range parsed.Typedefs {
			KnownTypedefs[td.Alias] = td // copy
		}
		for _, en := range parsed.Enums {
			KnownEnums[en.EnumName] = en // copy
		}

		processHeaders = append(processHeaders, parsed)
	}

	//
	// PASS 2
	//

	for _, parsed := range processHeaders {

		log.Printf("Processing %q...", parsed.Filename)

		// More AST transforms on our IL
		astTransformTypedefs(parsed)
		astTransformBlocklist(parsed) // Must happen after typedef transformation

		{
			// Save the IL file for debug inspection
			jb, err := json.MarshalIndent(parsed, "", "\t")
			if err != nil {
				panic(err)
			}

			err = ioutil.WriteFile(cacheFilePath(parsed.Filename)+".ours.json", jb, 0644)
			if err != nil {
				panic(err)
			}
		}

		// Breakout if there is nothing bindable
		if parsed.Empty() {
			log.Printf("Nothing in this header was bindable.")
			continue
		}

		// Emit 3 code files from the intermediate format
		outputName := filepath.Join(*outDir, "gen_"+strings.TrimSuffix(filepath.Base(parsed.Filename), `.h`))

		goSrc, err := emitGo(parsed, filepath.Base(parsed.Filename))
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(outputName+".go", []byte(goSrc), 0644)
		if err != nil {
			panic(err)
		}

		bindingCppSrc, err := emitBindingCpp(parsed, filepath.Base(parsed.Filename))
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(outputName+".cpp", []byte(bindingCppSrc), 0644)
		if err != nil {
			panic(err)
		}

		bindingHSrc, err := emitBindingHeader(parsed, filepath.Base(parsed.Filename))
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(outputName+".h", []byte(bindingHSrc), 0644)
		if err != nil {
			panic(err)
		}

		// Done

	}

	log.Printf("Processing %d file(s) completed", len(includeFiles))
}

func generateClangCaches(includeFiles []string, clangBin string, cflags []string) {

	var clangChan = make(chan string, 0)
	var clangWg sync.WaitGroup
	ctx := context.Background()

	for i := 0; i < ClangSubprocessCount; i++ {
		clangWg.Add(1)
		go func() {
			defer clangWg.Done()
			log.Printf("Clang worker: starting")

			for {
				inputHeader, ok := <-clangChan
				if !ok {
					return // Done
				}

				log.Printf("Clang worker got message for file %q", inputHeader)

				// Parse the file
				// This seems to intermittently fail, so allow retrying
				astInner := mustClangExec(ctx, clangBin, inputHeader, cflags)

				// Write to cache
				jb, err := json.MarshalIndent(astInner, "", "\t")
				if err != nil {
					panic(err)
				}

				err = ioutil.WriteFile(cacheFilePath(inputHeader), jb, 0644)
				if err != nil {
					panic(err)
				}

				astInner = nil
				jb = nil
				runtime.GC()

			}
			log.Printf("Clang worker: exiting")
		}()
	}

	for _, inputHeader := range includeFiles {

		// Check if there is a matching cache hit
		cacheFile := cacheFilePath(inputHeader)

		if _, err := os.Stat(cacheFile); err != nil && os.IsNotExist(err) {

			// Nonexistent cache file, regenerate from clang
			log.Printf("No AST cache for file %q, running clang...", filepath.Base(inputHeader))
			clangChan <- inputHeader
		}
	}

	// Done with all clang workers
	close(clangChan)
	clangWg.Wait()
}
