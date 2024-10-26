package qt6

/*

#include "gen_qcommandlineparser.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QCommandLineParser__SingleDashWordOptionMode int

const (
	QCommandLineParser__ParseAsCompactedShortOptions QCommandLineParser__SingleDashWordOptionMode = 0
	QCommandLineParser__ParseAsLongOptions           QCommandLineParser__SingleDashWordOptionMode = 1
)

type QCommandLineParser__OptionsAfterPositionalArgumentsMode int

const (
	QCommandLineParser__ParseAsOptions             QCommandLineParser__OptionsAfterPositionalArgumentsMode = 0
	QCommandLineParser__ParseAsPositionalArguments QCommandLineParser__OptionsAfterPositionalArgumentsMode = 1
)

type QCommandLineParser struct {
	h *C.QCommandLineParser
}

func (this *QCommandLineParser) cPointer() *C.QCommandLineParser {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QCommandLineParser) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

func newQCommandLineParser(h *C.QCommandLineParser) *QCommandLineParser {
	if h == nil {
		return nil
	}
	return &QCommandLineParser{h: h}
}

func UnsafeNewQCommandLineParser(h unsafe.Pointer) *QCommandLineParser {
	return newQCommandLineParser((*C.QCommandLineParser)(h))
}

// NewQCommandLineParser constructs a new QCommandLineParser object.
func NewQCommandLineParser() *QCommandLineParser {
	ret := C.QCommandLineParser_new()
	return newQCommandLineParser(ret)
}

func QCommandLineParser_Tr(sourceText string) string {
	sourceText_Cstring := C.CString(sourceText)
	defer C.free(unsafe.Pointer(sourceText_Cstring))
	var _ms C.struct_miqt_string = C.QCommandLineParser_Tr(sourceText_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCommandLineParser) SetSingleDashWordOptionMode(parsingMode QCommandLineParser__SingleDashWordOptionMode) {
	C.QCommandLineParser_SetSingleDashWordOptionMode(this.h, (C.int)(parsingMode))
}

func (this *QCommandLineParser) SetOptionsAfterPositionalArgumentsMode(mode QCommandLineParser__OptionsAfterPositionalArgumentsMode) {
	C.QCommandLineParser_SetOptionsAfterPositionalArgumentsMode(this.h, (C.int)(mode))
}

func (this *QCommandLineParser) AddOption(commandLineOption *QCommandLineOption) bool {
	return (bool)(C.QCommandLineParser_AddOption(this.h, commandLineOption.cPointer()))
}

func (this *QCommandLineParser) AddOptions(options []QCommandLineOption) bool {
	// For the C ABI, malloc a C array of raw pointers
	options_CArray := (*[0xffff]*C.QCommandLineOption)(C.malloc(C.size_t(8 * len(options))))
	defer C.free(unsafe.Pointer(options_CArray))
	for i := range options {
		options_CArray[i] = options[i].cPointer()
	}
	options_ma := &C.struct_miqt_array{len: C.size_t(len(options)), data: unsafe.Pointer(options_CArray)}
	defer runtime.KeepAlive(unsafe.Pointer(options_ma))
	return (bool)(C.QCommandLineParser_AddOptions(this.h, options_ma))
}

func (this *QCommandLineParser) AddVersionOption() *QCommandLineOption {
	_ret := C.QCommandLineParser_AddVersionOption(this.h)
	_goptr := newQCommandLineOption(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCommandLineParser) AddHelpOption() *QCommandLineOption {
	_ret := C.QCommandLineParser_AddHelpOption(this.h)
	_goptr := newQCommandLineOption(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCommandLineParser) SetApplicationDescription(description string) {
	description_ms := C.struct_miqt_string{}
	description_ms.data = C.CString(description)
	description_ms.len = C.size_t(len(description))
	defer C.free(unsafe.Pointer(description_ms.data))
	C.QCommandLineParser_SetApplicationDescription(this.h, description_ms)
}

func (this *QCommandLineParser) ApplicationDescription() string {
	var _ms C.struct_miqt_string = C.QCommandLineParser_ApplicationDescription(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCommandLineParser) AddPositionalArgument(name string, description string) {
	name_ms := C.struct_miqt_string{}
	name_ms.data = C.CString(name)
	name_ms.len = C.size_t(len(name))
	defer C.free(unsafe.Pointer(name_ms.data))
	description_ms := C.struct_miqt_string{}
	description_ms.data = C.CString(description)
	description_ms.len = C.size_t(len(description))
	defer C.free(unsafe.Pointer(description_ms.data))
	C.QCommandLineParser_AddPositionalArgument(this.h, name_ms, description_ms)
}

func (this *QCommandLineParser) ClearPositionalArguments() {
	C.QCommandLineParser_ClearPositionalArguments(this.h)
}

func (this *QCommandLineParser) Process(arguments []string) {
	// For the C ABI, malloc a C array of structs
	arguments_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(arguments))))
	defer C.free(unsafe.Pointer(arguments_CArray))
	for i := range arguments {
		arguments_i_ms := C.struct_miqt_string{}
		arguments_i_ms.data = C.CString(arguments[i])
		arguments_i_ms.len = C.size_t(len(arguments[i]))
		defer C.free(unsafe.Pointer(arguments_i_ms.data))
		arguments_CArray[i] = arguments_i_ms
	}
	arguments_ma := &C.struct_miqt_array{len: C.size_t(len(arguments)), data: unsafe.Pointer(arguments_CArray)}
	defer runtime.KeepAlive(unsafe.Pointer(arguments_ma))
	C.QCommandLineParser_Process(this.h, arguments_ma)
}

func (this *QCommandLineParser) ProcessWithApp(app *QCoreApplication) {
	C.QCommandLineParser_ProcessWithApp(this.h, app.cPointer())
}

func (this *QCommandLineParser) Parse(arguments []string) bool {
	// For the C ABI, malloc a C array of structs
	arguments_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(arguments))))
	defer C.free(unsafe.Pointer(arguments_CArray))
	for i := range arguments {
		arguments_i_ms := C.struct_miqt_string{}
		arguments_i_ms.data = C.CString(arguments[i])
		arguments_i_ms.len = C.size_t(len(arguments[i]))
		defer C.free(unsafe.Pointer(arguments_i_ms.data))
		arguments_CArray[i] = arguments_i_ms
	}
	arguments_ma := &C.struct_miqt_array{len: C.size_t(len(arguments)), data: unsafe.Pointer(arguments_CArray)}
	defer runtime.KeepAlive(unsafe.Pointer(arguments_ma))
	return (bool)(C.QCommandLineParser_Parse(this.h, arguments_ma))
}

func (this *QCommandLineParser) ErrorText() string {
	var _ms C.struct_miqt_string = C.QCommandLineParser_ErrorText(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCommandLineParser) IsSet(name string) bool {
	name_ms := C.struct_miqt_string{}
	name_ms.data = C.CString(name)
	name_ms.len = C.size_t(len(name))
	defer C.free(unsafe.Pointer(name_ms.data))
	return (bool)(C.QCommandLineParser_IsSet(this.h, name_ms))
}

func (this *QCommandLineParser) Value(name string) string {
	name_ms := C.struct_miqt_string{}
	name_ms.data = C.CString(name)
	name_ms.len = C.size_t(len(name))
	defer C.free(unsafe.Pointer(name_ms.data))
	var _ms C.struct_miqt_string = C.QCommandLineParser_Value(this.h, name_ms)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCommandLineParser) Values(name string) []string {
	name_ms := C.struct_miqt_string{}
	name_ms.data = C.CString(name)
	name_ms.len = C.size_t(len(name))
	defer C.free(unsafe.Pointer(name_ms.data))
	var _ma *C.struct_miqt_array = C.QCommandLineParser_Values(this.h, name_ms)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]C.struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms C.struct_miqt_string = _outCast[i]
		_lv_ret := C.GoStringN(_lv_ms.data, C.int(int64(_lv_ms.len)))
		C.free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	C.free(unsafe.Pointer(_ma))
	return _ret
}

func (this *QCommandLineParser) IsSetWithOption(option *QCommandLineOption) bool {
	return (bool)(C.QCommandLineParser_IsSetWithOption(this.h, option.cPointer()))
}

func (this *QCommandLineParser) ValueWithOption(option *QCommandLineOption) string {
	var _ms C.struct_miqt_string = C.QCommandLineParser_ValueWithOption(this.h, option.cPointer())
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCommandLineParser) ValuesWithOption(option *QCommandLineOption) []string {
	var _ma *C.struct_miqt_array = C.QCommandLineParser_ValuesWithOption(this.h, option.cPointer())
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]C.struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms C.struct_miqt_string = _outCast[i]
		_lv_ret := C.GoStringN(_lv_ms.data, C.int(int64(_lv_ms.len)))
		C.free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	C.free(unsafe.Pointer(_ma))
	return _ret
}

func (this *QCommandLineParser) PositionalArguments() []string {
	var _ma *C.struct_miqt_array = C.QCommandLineParser_PositionalArguments(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]C.struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms C.struct_miqt_string = _outCast[i]
		_lv_ret := C.GoStringN(_lv_ms.data, C.int(int64(_lv_ms.len)))
		C.free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	C.free(unsafe.Pointer(_ma))
	return _ret
}

func (this *QCommandLineParser) OptionNames() []string {
	var _ma *C.struct_miqt_array = C.QCommandLineParser_OptionNames(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]C.struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms C.struct_miqt_string = _outCast[i]
		_lv_ret := C.GoStringN(_lv_ms.data, C.int(int64(_lv_ms.len)))
		C.free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	C.free(unsafe.Pointer(_ma))
	return _ret
}

func (this *QCommandLineParser) UnknownOptionNames() []string {
	var _ma *C.struct_miqt_array = C.QCommandLineParser_UnknownOptionNames(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]C.struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms C.struct_miqt_string = _outCast[i]
		_lv_ret := C.GoStringN(_lv_ms.data, C.int(int64(_lv_ms.len)))
		C.free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	C.free(unsafe.Pointer(_ma))
	return _ret
}

func (this *QCommandLineParser) ShowVersion() {
	C.QCommandLineParser_ShowVersion(this.h)
}

func (this *QCommandLineParser) ShowHelp() {
	C.QCommandLineParser_ShowHelp(this.h)
}

func (this *QCommandLineParser) HelpText() string {
	var _ms C.struct_miqt_string = C.QCommandLineParser_HelpText(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCommandLineParser_Tr2(sourceText string, disambiguation string) string {
	sourceText_Cstring := C.CString(sourceText)
	defer C.free(unsafe.Pointer(sourceText_Cstring))
	disambiguation_Cstring := C.CString(disambiguation)
	defer C.free(unsafe.Pointer(disambiguation_Cstring))
	var _ms C.struct_miqt_string = C.QCommandLineParser_Tr2(sourceText_Cstring, disambiguation_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCommandLineParser_Tr3(sourceText string, disambiguation string, n int) string {
	sourceText_Cstring := C.CString(sourceText)
	defer C.free(unsafe.Pointer(sourceText_Cstring))
	disambiguation_Cstring := C.CString(disambiguation)
	defer C.free(unsafe.Pointer(disambiguation_Cstring))
	var _ms C.struct_miqt_string = C.QCommandLineParser_Tr3(sourceText_Cstring, disambiguation_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCommandLineParser) AddPositionalArgument3(name string, description string, syntax string) {
	name_ms := C.struct_miqt_string{}
	name_ms.data = C.CString(name)
	name_ms.len = C.size_t(len(name))
	defer C.free(unsafe.Pointer(name_ms.data))
	description_ms := C.struct_miqt_string{}
	description_ms.data = C.CString(description)
	description_ms.len = C.size_t(len(description))
	defer C.free(unsafe.Pointer(description_ms.data))
	syntax_ms := C.struct_miqt_string{}
	syntax_ms.data = C.CString(syntax)
	syntax_ms.len = C.size_t(len(syntax))
	defer C.free(unsafe.Pointer(syntax_ms.data))
	C.QCommandLineParser_AddPositionalArgument3(this.h, name_ms, description_ms, syntax_ms)
}

func (this *QCommandLineParser) ShowHelp1(exitCode int) {
	C.QCommandLineParser_ShowHelp1(this.h, (C.int)(exitCode))
}

// Delete this object from C++ memory.
func (this *QCommandLineParser) Delete() {
	C.QCommandLineParser_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QCommandLineParser) GoGC() {
	runtime.SetFinalizer(this, func(this *QCommandLineParser) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}