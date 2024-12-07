package qt

/*

#include "gen_qiconengine.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QIconEngine__IconEngineHook int

const (
	QIconEngine__AvailableSizesHook QIconEngine__IconEngineHook = 1
	QIconEngine__IconNameHook       QIconEngine__IconEngineHook = 2
	QIconEngine__IsNullHook         QIconEngine__IconEngineHook = 3
	QIconEngine__ScaledPixmapHook   QIconEngine__IconEngineHook = 4
)

type QIconEngine struct {
	h          *C.QIconEngine
	isSubclass bool
}

func (this *QIconEngine) cPointer() *C.QIconEngine {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QIconEngine) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQIconEngine constructs the type using only CGO pointers.
func newQIconEngine(h *C.QIconEngine) *QIconEngine {
	if h == nil {
		return nil
	}
	return &QIconEngine{h: h}
}

// UnsafeNewQIconEngine constructs the type using only unsafe pointers.
func UnsafeNewQIconEngine(h unsafe.Pointer) *QIconEngine {
	if h == nil {
		return nil
	}

	return &QIconEngine{h: (*C.QIconEngine)(h)}
}

// NewQIconEngine constructs a new QIconEngine object.
func NewQIconEngine() *QIconEngine {
	var outptr_QIconEngine *C.QIconEngine = nil

	C.QIconEngine_new(&outptr_QIconEngine)
	ret := newQIconEngine(outptr_QIconEngine)
	ret.isSubclass = true
	return ret
}

// NewQIconEngine2 constructs a new QIconEngine object.
func NewQIconEngine2(other *QIconEngine) *QIconEngine {
	var outptr_QIconEngine *C.QIconEngine = nil

	C.QIconEngine_new2(other.cPointer(), &outptr_QIconEngine)
	ret := newQIconEngine(outptr_QIconEngine)
	ret.isSubclass = true
	return ret
}

func (this *QIconEngine) Paint(painter *QPainter, rect *QRect, mode QIcon__Mode, state QIcon__State) {
	C.QIconEngine_Paint(this.h, painter.cPointer(), rect.cPointer(), (C.int)(mode), (C.int)(state))
}

func (this *QIconEngine) ActualSize(size *QSize, mode QIcon__Mode, state QIcon__State) *QSize {
	_ret := C.QIconEngine_ActualSize(this.h, size.cPointer(), (C.int)(mode), (C.int)(state))
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIconEngine) Pixmap(size *QSize, mode QIcon__Mode, state QIcon__State) *QPixmap {
	_ret := C.QIconEngine_Pixmap(this.h, size.cPointer(), (C.int)(mode), (C.int)(state))
	_goptr := newQPixmap(_ret, nil)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIconEngine) AddPixmap(pixmap *QPixmap, mode QIcon__Mode, state QIcon__State) {
	C.QIconEngine_AddPixmap(this.h, pixmap.cPointer(), (C.int)(mode), (C.int)(state))
}

func (this *QIconEngine) AddFile(fileName string, size *QSize, mode QIcon__Mode, state QIcon__State) {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	C.QIconEngine_AddFile(this.h, fileName_ms, size.cPointer(), (C.int)(mode), (C.int)(state))
}

func (this *QIconEngine) Key() string {
	var _ms C.struct_miqt_string = C.QIconEngine_Key(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QIconEngine) Clone() *QIconEngine {
	return UnsafeNewQIconEngine(unsafe.Pointer(C.QIconEngine_Clone(this.h)))
}

func (this *QIconEngine) Read(in *QDataStream) bool {
	return (bool)(C.QIconEngine_Read(this.h, in.cPointer()))
}

func (this *QIconEngine) Write(out *QDataStream) bool {
	return (bool)(C.QIconEngine_Write(this.h, out.cPointer()))
}

func (this *QIconEngine) AvailableSizes(mode QIcon__Mode, state QIcon__State) []QSize {
	var _ma C.struct_miqt_array = C.QIconEngine_AvailableSizes(this.h, (C.int)(mode), (C.int)(state))
	_ret := make([]QSize, int(_ma.len))
	_outCast := (*[0xffff]*C.QSize)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_ret := _outCast[i]
		_lv_goptr := newQSize(_lv_ret)
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QIconEngine) IconName() string {
	var _ms C.struct_miqt_string = C.QIconEngine_IconName(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QIconEngine) IsNull() bool {
	return (bool)(C.QIconEngine_IsNull(this.h))
}

func (this *QIconEngine) ScaledPixmap(size *QSize, mode QIcon__Mode, state QIcon__State, scale float64) *QPixmap {
	_ret := C.QIconEngine_ScaledPixmap(this.h, size.cPointer(), (C.int)(mode), (C.int)(state), (C.double)(scale))
	_goptr := newQPixmap(_ret, nil)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIconEngine) VirtualHook(id int, data unsafe.Pointer) {
	C.QIconEngine_VirtualHook(this.h, (C.int)(id), data)
}
func (this *QIconEngine) OnPaint(slot func(painter *QPainter, rect *QRect, mode QIcon__Mode, state QIcon__State)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QIconEngine_override_virtual_Paint(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_Paint
func miqt_exec_callback_QIconEngine_Paint(self *C.QIconEngine, cb C.intptr_t, painter *C.QPainter, rect *C.QRect, mode C.int, state C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(painter *QPainter, rect *QRect, mode QIcon__Mode, state QIcon__State))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQPainter(unsafe.Pointer(painter))
	slotval2 := UnsafeNewQRect(unsafe.Pointer(rect))
	slotval3 := (QIcon__Mode)(mode)

	slotval4 := (QIcon__State)(state)

	gofunc(slotval1, slotval2, slotval3, slotval4)

}

func (this *QIconEngine) callVirtualBase_ActualSize(size *QSize, mode QIcon__Mode, state QIcon__State) *QSize {

	_ret := C.QIconEngine_virtualbase_ActualSize(unsafe.Pointer(this.h), size.cPointer(), (C.int)(mode), (C.int)(state))
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QIconEngine) OnActualSize(slot func(super func(size *QSize, mode QIcon__Mode, state QIcon__State) *QSize, size *QSize, mode QIcon__Mode, state QIcon__State) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QIconEngine_override_virtual_ActualSize(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_ActualSize
func miqt_exec_callback_QIconEngine_ActualSize(self *C.QIconEngine, cb C.intptr_t, size *C.QSize, mode C.int, state C.int) *C.QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(size *QSize, mode QIcon__Mode, state QIcon__State) *QSize, size *QSize, mode QIcon__Mode, state QIcon__State) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQSize(unsafe.Pointer(size))
	slotval2 := (QIcon__Mode)(mode)

	slotval3 := (QIcon__State)(state)

	virtualReturn := gofunc((&QIconEngine{h: self}).callVirtualBase_ActualSize, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QIconEngine) callVirtualBase_Pixmap(size *QSize, mode QIcon__Mode, state QIcon__State) *QPixmap {

	_ret := C.QIconEngine_virtualbase_Pixmap(unsafe.Pointer(this.h), size.cPointer(), (C.int)(mode), (C.int)(state))
	_goptr := newQPixmap(_ret, nil)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QIconEngine) OnPixmap(slot func(super func(size *QSize, mode QIcon__Mode, state QIcon__State) *QPixmap, size *QSize, mode QIcon__Mode, state QIcon__State) *QPixmap) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QIconEngine_override_virtual_Pixmap(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_Pixmap
func miqt_exec_callback_QIconEngine_Pixmap(self *C.QIconEngine, cb C.intptr_t, size *C.QSize, mode C.int, state C.int) *C.QPixmap {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(size *QSize, mode QIcon__Mode, state QIcon__State) *QPixmap, size *QSize, mode QIcon__Mode, state QIcon__State) *QPixmap)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQSize(unsafe.Pointer(size))
	slotval2 := (QIcon__Mode)(mode)

	slotval3 := (QIcon__State)(state)

	virtualReturn := gofunc((&QIconEngine{h: self}).callVirtualBase_Pixmap, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QIconEngine) callVirtualBase_AddPixmap(pixmap *QPixmap, mode QIcon__Mode, state QIcon__State) {

	C.QIconEngine_virtualbase_AddPixmap(unsafe.Pointer(this.h), pixmap.cPointer(), (C.int)(mode), (C.int)(state))

}
func (this *QIconEngine) OnAddPixmap(slot func(super func(pixmap *QPixmap, mode QIcon__Mode, state QIcon__State), pixmap *QPixmap, mode QIcon__Mode, state QIcon__State)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QIconEngine_override_virtual_AddPixmap(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_AddPixmap
func miqt_exec_callback_QIconEngine_AddPixmap(self *C.QIconEngine, cb C.intptr_t, pixmap *C.QPixmap, mode C.int, state C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(pixmap *QPixmap, mode QIcon__Mode, state QIcon__State), pixmap *QPixmap, mode QIcon__Mode, state QIcon__State))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQPixmap(unsafe.Pointer(pixmap), nil)
	slotval2 := (QIcon__Mode)(mode)

	slotval3 := (QIcon__State)(state)

	gofunc((&QIconEngine{h: self}).callVirtualBase_AddPixmap, slotval1, slotval2, slotval3)

}

func (this *QIconEngine) callVirtualBase_AddFile(fileName string, size *QSize, mode QIcon__Mode, state QIcon__State) {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))

	C.QIconEngine_virtualbase_AddFile(unsafe.Pointer(this.h), fileName_ms, size.cPointer(), (C.int)(mode), (C.int)(state))

}
func (this *QIconEngine) OnAddFile(slot func(super func(fileName string, size *QSize, mode QIcon__Mode, state QIcon__State), fileName string, size *QSize, mode QIcon__Mode, state QIcon__State)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QIconEngine_override_virtual_AddFile(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_AddFile
func miqt_exec_callback_QIconEngine_AddFile(self *C.QIconEngine, cb C.intptr_t, fileName C.struct_miqt_string, size *C.QSize, mode C.int, state C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(fileName string, size *QSize, mode QIcon__Mode, state QIcon__State), fileName string, size *QSize, mode QIcon__Mode, state QIcon__State))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var fileName_ms C.struct_miqt_string = fileName
	fileName_ret := C.GoStringN(fileName_ms.data, C.int(int64(fileName_ms.len)))
	C.free(unsafe.Pointer(fileName_ms.data))
	slotval1 := fileName_ret
	slotval2 := UnsafeNewQSize(unsafe.Pointer(size))
	slotval3 := (QIcon__Mode)(mode)

	slotval4 := (QIcon__State)(state)

	gofunc((&QIconEngine{h: self}).callVirtualBase_AddFile, slotval1, slotval2, slotval3, slotval4)

}

func (this *QIconEngine) callVirtualBase_Key() string {

	var _ms C.struct_miqt_string = C.QIconEngine_virtualbase_Key(unsafe.Pointer(this.h))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QIconEngine) OnKey(slot func(super func() string) string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QIconEngine_override_virtual_Key(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_Key
func miqt_exec_callback_QIconEngine_Key(self *C.QIconEngine, cb C.intptr_t) C.struct_miqt_string {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() string) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QIconEngine{h: self}).callVirtualBase_Key)
	virtualReturn_ms := C.struct_miqt_string{}
	virtualReturn_ms.data = C.CString(virtualReturn)
	virtualReturn_ms.len = C.size_t(len(virtualReturn))
	defer C.free(unsafe.Pointer(virtualReturn_ms.data))

	return virtualReturn_ms

}
func (this *QIconEngine) OnClone(slot func() *QIconEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QIconEngine_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_Clone
func miqt_exec_callback_QIconEngine_Clone(self *C.QIconEngine, cb C.intptr_t) *C.QIconEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func() *QIconEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc()

	return virtualReturn.cPointer()

}

func (this *QIconEngine) callVirtualBase_Read(in *QDataStream) bool {

	return (bool)(C.QIconEngine_virtualbase_Read(unsafe.Pointer(this.h), in.cPointer()))

}
func (this *QIconEngine) OnRead(slot func(super func(in *QDataStream) bool, in *QDataStream) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QIconEngine_override_virtual_Read(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_Read
func miqt_exec_callback_QIconEngine_Read(self *C.QIconEngine, cb C.intptr_t, in *C.QDataStream) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(in *QDataStream) bool, in *QDataStream) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQDataStream(unsafe.Pointer(in))

	virtualReturn := gofunc((&QIconEngine{h: self}).callVirtualBase_Read, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QIconEngine) callVirtualBase_Write(out *QDataStream) bool {

	return (bool)(C.QIconEngine_virtualbase_Write(unsafe.Pointer(this.h), out.cPointer()))

}
func (this *QIconEngine) OnWrite(slot func(super func(out *QDataStream) bool, out *QDataStream) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QIconEngine_override_virtual_Write(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_Write
func miqt_exec_callback_QIconEngine_Write(self *C.QIconEngine, cb C.intptr_t, out *C.QDataStream) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(out *QDataStream) bool, out *QDataStream) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQDataStream(unsafe.Pointer(out))

	virtualReturn := gofunc((&QIconEngine{h: self}).callVirtualBase_Write, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QIconEngine) callVirtualBase_AvailableSizes(mode QIcon__Mode, state QIcon__State) []QSize {

	var _ma C.struct_miqt_array = C.QIconEngine_virtualbase_AvailableSizes(unsafe.Pointer(this.h), (C.int)(mode), (C.int)(state))
	_ret := make([]QSize, int(_ma.len))
	_outCast := (*[0xffff]*C.QSize)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_ret := _outCast[i]
		_lv_goptr := newQSize(_lv_ret)
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret

}
func (this *QIconEngine) OnAvailableSizes(slot func(super func(mode QIcon__Mode, state QIcon__State) []QSize, mode QIcon__Mode, state QIcon__State) []QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QIconEngine_override_virtual_AvailableSizes(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_AvailableSizes
func miqt_exec_callback_QIconEngine_AvailableSizes(self *C.QIconEngine, cb C.intptr_t, mode C.int, state C.int) C.struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(mode QIcon__Mode, state QIcon__State) []QSize, mode QIcon__Mode, state QIcon__State) []QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QIcon__Mode)(mode)

	slotval2 := (QIcon__State)(state)

	virtualReturn := gofunc((&QIconEngine{h: self}).callVirtualBase_AvailableSizes, slotval1, slotval2)
	virtualReturn_CArray := (*[0xffff]*C.QSize)(C.malloc(C.size_t(8 * len(virtualReturn))))
	defer C.free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_CArray[i] = virtualReturn[i].cPointer()
	}
	virtualReturn_ma := C.struct_miqt_array{len: C.size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QIconEngine) callVirtualBase_IconName() string {

	var _ms C.struct_miqt_string = C.QIconEngine_virtualbase_IconName(unsafe.Pointer(this.h))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QIconEngine) OnIconName(slot func(super func() string) string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QIconEngine_override_virtual_IconName(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_IconName
func miqt_exec_callback_QIconEngine_IconName(self *C.QIconEngine, cb C.intptr_t) C.struct_miqt_string {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() string) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QIconEngine{h: self}).callVirtualBase_IconName)
	virtualReturn_ms := C.struct_miqt_string{}
	virtualReturn_ms.data = C.CString(virtualReturn)
	virtualReturn_ms.len = C.size_t(len(virtualReturn))
	defer C.free(unsafe.Pointer(virtualReturn_ms.data))

	return virtualReturn_ms

}

func (this *QIconEngine) callVirtualBase_VirtualHook(id int, data unsafe.Pointer) {

	C.QIconEngine_virtualbase_VirtualHook(unsafe.Pointer(this.h), (C.int)(id), data)

}
func (this *QIconEngine) OnVirtualHook(slot func(super func(id int, data unsafe.Pointer), id int, data unsafe.Pointer)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QIconEngine_override_virtual_VirtualHook(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_VirtualHook
func miqt_exec_callback_QIconEngine_VirtualHook(self *C.QIconEngine, cb C.intptr_t, id C.int, data unsafe.Pointer) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(id int, data unsafe.Pointer), id int, data unsafe.Pointer))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(id)

	slotval2 := (unsafe.Pointer)(data)

	gofunc((&QIconEngine{h: self}).callVirtualBase_VirtualHook, slotval1, slotval2)

}

// Delete this object from C++ memory.
func (this *QIconEngine) Delete() {
	C.QIconEngine_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QIconEngine) GoGC() {
	runtime.SetFinalizer(this, func(this *QIconEngine) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QIconEngine__AvailableSizesArgument struct {
	h          *C.QIconEngine__AvailableSizesArgument
	isSubclass bool
}

func (this *QIconEngine__AvailableSizesArgument) cPointer() *C.QIconEngine__AvailableSizesArgument {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QIconEngine__AvailableSizesArgument) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQIconEngine__AvailableSizesArgument constructs the type using only CGO pointers.
func newQIconEngine__AvailableSizesArgument(h *C.QIconEngine__AvailableSizesArgument) *QIconEngine__AvailableSizesArgument {
	if h == nil {
		return nil
	}
	return &QIconEngine__AvailableSizesArgument{h: h}
}

// UnsafeNewQIconEngine__AvailableSizesArgument constructs the type using only unsafe pointers.
func UnsafeNewQIconEngine__AvailableSizesArgument(h unsafe.Pointer) *QIconEngine__AvailableSizesArgument {
	if h == nil {
		return nil
	}

	return &QIconEngine__AvailableSizesArgument{h: (*C.QIconEngine__AvailableSizesArgument)(h)}
}

// NewQIconEngine__AvailableSizesArgument constructs a new QIconEngine::AvailableSizesArgument object.
func NewQIconEngine__AvailableSizesArgument(param1 *QIconEngine__AvailableSizesArgument) *QIconEngine__AvailableSizesArgument {
	var outptr_QIconEngine__AvailableSizesArgument *C.QIconEngine__AvailableSizesArgument = nil

	C.QIconEngine__AvailableSizesArgument_new(param1.cPointer(), &outptr_QIconEngine__AvailableSizesArgument)
	ret := newQIconEngine__AvailableSizesArgument(outptr_QIconEngine__AvailableSizesArgument)
	ret.isSubclass = true
	return ret
}

func (this *QIconEngine__AvailableSizesArgument) OperatorAssign(param1 *QIconEngine__AvailableSizesArgument) {
	C.QIconEngine__AvailableSizesArgument_OperatorAssign(this.h, param1.cPointer())
}

// Delete this object from C++ memory.
func (this *QIconEngine__AvailableSizesArgument) Delete() {
	C.QIconEngine__AvailableSizesArgument_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QIconEngine__AvailableSizesArgument) GoGC() {
	runtime.SetFinalizer(this, func(this *QIconEngine__AvailableSizesArgument) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QIconEngine__ScaledPixmapArgument struct {
	h          *C.QIconEngine__ScaledPixmapArgument
	isSubclass bool
}

func (this *QIconEngine__ScaledPixmapArgument) cPointer() *C.QIconEngine__ScaledPixmapArgument {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QIconEngine__ScaledPixmapArgument) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQIconEngine__ScaledPixmapArgument constructs the type using only CGO pointers.
func newQIconEngine__ScaledPixmapArgument(h *C.QIconEngine__ScaledPixmapArgument) *QIconEngine__ScaledPixmapArgument {
	if h == nil {
		return nil
	}
	return &QIconEngine__ScaledPixmapArgument{h: h}
}

// UnsafeNewQIconEngine__ScaledPixmapArgument constructs the type using only unsafe pointers.
func UnsafeNewQIconEngine__ScaledPixmapArgument(h unsafe.Pointer) *QIconEngine__ScaledPixmapArgument {
	if h == nil {
		return nil
	}

	return &QIconEngine__ScaledPixmapArgument{h: (*C.QIconEngine__ScaledPixmapArgument)(h)}
}

// NewQIconEngine__ScaledPixmapArgument constructs a new QIconEngine::ScaledPixmapArgument object.
func NewQIconEngine__ScaledPixmapArgument(param1 *QIconEngine__ScaledPixmapArgument) *QIconEngine__ScaledPixmapArgument {
	var outptr_QIconEngine__ScaledPixmapArgument *C.QIconEngine__ScaledPixmapArgument = nil

	C.QIconEngine__ScaledPixmapArgument_new(param1.cPointer(), &outptr_QIconEngine__ScaledPixmapArgument)
	ret := newQIconEngine__ScaledPixmapArgument(outptr_QIconEngine__ScaledPixmapArgument)
	ret.isSubclass = true
	return ret
}

func (this *QIconEngine__ScaledPixmapArgument) OperatorAssign(param1 *QIconEngine__ScaledPixmapArgument) {
	C.QIconEngine__ScaledPixmapArgument_OperatorAssign(this.h, param1.cPointer())
}

// Delete this object from C++ memory.
func (this *QIconEngine__ScaledPixmapArgument) Delete() {
	C.QIconEngine__ScaledPixmapArgument_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QIconEngine__ScaledPixmapArgument) GoGC() {
	runtime.SetFinalizer(this, func(this *QIconEngine__ScaledPixmapArgument) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
