package qt6

/*

#include "gen_qoffscreensurface.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QOffscreenSurface struct {
	h          *C.QOffscreenSurface
	isSubclass bool
	*QObject
	*QSurface
}

func (this *QOffscreenSurface) cPointer() *C.QOffscreenSurface {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QOffscreenSurface) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQOffscreenSurface constructs the type using only CGO pointers.
func newQOffscreenSurface(h *C.QOffscreenSurface, h_QObject *C.QObject, h_QSurface *C.QSurface) *QOffscreenSurface {
	if h == nil {
		return nil
	}
	return &QOffscreenSurface{h: h,
		QObject:  newQObject(h_QObject),
		QSurface: newQSurface(h_QSurface)}
}

// UnsafeNewQOffscreenSurface constructs the type using only unsafe pointers.
func UnsafeNewQOffscreenSurface(h unsafe.Pointer, h_QObject unsafe.Pointer, h_QSurface unsafe.Pointer) *QOffscreenSurface {
	if h == nil {
		return nil
	}

	return &QOffscreenSurface{h: (*C.QOffscreenSurface)(h),
		QObject:  UnsafeNewQObject(h_QObject),
		QSurface: UnsafeNewQSurface(h_QSurface)}
}

// NewQOffscreenSurface constructs a new QOffscreenSurface object.
func NewQOffscreenSurface() *QOffscreenSurface {
	var outptr_QOffscreenSurface *C.QOffscreenSurface = nil
	var outptr_QObject *C.QObject = nil
	var outptr_QSurface *C.QSurface = nil

	C.QOffscreenSurface_new(&outptr_QOffscreenSurface, &outptr_QObject, &outptr_QSurface)
	ret := newQOffscreenSurface(outptr_QOffscreenSurface, outptr_QObject, outptr_QSurface)
	ret.isSubclass = true
	return ret
}

// NewQOffscreenSurface2 constructs a new QOffscreenSurface object.
func NewQOffscreenSurface2(screen *QScreen) *QOffscreenSurface {
	var outptr_QOffscreenSurface *C.QOffscreenSurface = nil
	var outptr_QObject *C.QObject = nil
	var outptr_QSurface *C.QSurface = nil

	C.QOffscreenSurface_new2(screen.cPointer(), &outptr_QOffscreenSurface, &outptr_QObject, &outptr_QSurface)
	ret := newQOffscreenSurface(outptr_QOffscreenSurface, outptr_QObject, outptr_QSurface)
	ret.isSubclass = true
	return ret
}

// NewQOffscreenSurface3 constructs a new QOffscreenSurface object.
func NewQOffscreenSurface3(screen *QScreen, parent *QObject) *QOffscreenSurface {
	var outptr_QOffscreenSurface *C.QOffscreenSurface = nil
	var outptr_QObject *C.QObject = nil
	var outptr_QSurface *C.QSurface = nil

	C.QOffscreenSurface_new3(screen.cPointer(), parent.cPointer(), &outptr_QOffscreenSurface, &outptr_QObject, &outptr_QSurface)
	ret := newQOffscreenSurface(outptr_QOffscreenSurface, outptr_QObject, outptr_QSurface)
	ret.isSubclass = true
	return ret
}

func (this *QOffscreenSurface) MetaObject() *QMetaObject {
	return UnsafeNewQMetaObject(unsafe.Pointer(C.QOffscreenSurface_MetaObject(this.h)))
}

func (this *QOffscreenSurface) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QOffscreenSurface_Metacast(this.h, param1_Cstring))
}

func QOffscreenSurface_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QOffscreenSurface_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QOffscreenSurface) SurfaceType() QSurface__SurfaceType {
	return (QSurface__SurfaceType)(C.QOffscreenSurface_SurfaceType(this.h))
}

func (this *QOffscreenSurface) Create() {
	C.QOffscreenSurface_Create(this.h)
}

func (this *QOffscreenSurface) Destroy() {
	C.QOffscreenSurface_Destroy(this.h)
}

func (this *QOffscreenSurface) IsValid() bool {
	return (bool)(C.QOffscreenSurface_IsValid(this.h))
}

func (this *QOffscreenSurface) SetFormat(format *QSurfaceFormat) {
	C.QOffscreenSurface_SetFormat(this.h, format.cPointer())
}

func (this *QOffscreenSurface) Format() *QSurfaceFormat {
	_ret := C.QOffscreenSurface_Format(this.h)
	_goptr := newQSurfaceFormat(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QOffscreenSurface) RequestedFormat() *QSurfaceFormat {
	_ret := C.QOffscreenSurface_RequestedFormat(this.h)
	_goptr := newQSurfaceFormat(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QOffscreenSurface) Size() *QSize {
	_ret := C.QOffscreenSurface_Size(this.h)
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QOffscreenSurface) Screen() *QScreen {
	return UnsafeNewQScreen(unsafe.Pointer(C.QOffscreenSurface_Screen(this.h)), nil)
}

func (this *QOffscreenSurface) SetScreen(screen *QScreen) {
	C.QOffscreenSurface_SetScreen(this.h, screen.cPointer())
}

func (this *QOffscreenSurface) ScreenChanged(screen *QScreen) {
	C.QOffscreenSurface_ScreenChanged(this.h, screen.cPointer())
}
func (this *QOffscreenSurface) OnScreenChanged(slot func(screen *QScreen)) {
	C.QOffscreenSurface_connect_ScreenChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_ScreenChanged
func miqt_exec_callback_QOffscreenSurface_ScreenChanged(cb C.intptr_t, screen *C.QScreen) {
	gofunc, ok := cgo.Handle(cb).Value().(func(screen *QScreen))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQScreen(unsafe.Pointer(screen), nil)

	gofunc(slotval1)
}

func QOffscreenSurface_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QOffscreenSurface_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QOffscreenSurface_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QOffscreenSurface_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QOffscreenSurface) callVirtualBase_SurfaceType() QSurface__SurfaceType {

	return (QSurface__SurfaceType)(C.QOffscreenSurface_virtualbase_SurfaceType(unsafe.Pointer(this.h)))

}
func (this *QOffscreenSurface) OnSurfaceType(slot func(super func() QSurface__SurfaceType) QSurface__SurfaceType) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QOffscreenSurface_override_virtual_SurfaceType(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_SurfaceType
func miqt_exec_callback_QOffscreenSurface_SurfaceType(self *C.QOffscreenSurface, cb C.intptr_t) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() QSurface__SurfaceType) QSurface__SurfaceType)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QOffscreenSurface{h: self}).callVirtualBase_SurfaceType)

	return (C.int)(virtualReturn)

}

func (this *QOffscreenSurface) callVirtualBase_Format() *QSurfaceFormat {

	_ret := C.QOffscreenSurface_virtualbase_Format(unsafe.Pointer(this.h))
	_goptr := newQSurfaceFormat(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QOffscreenSurface) OnFormat(slot func(super func() *QSurfaceFormat) *QSurfaceFormat) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QOffscreenSurface_override_virtual_Format(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_Format
func miqt_exec_callback_QOffscreenSurface_Format(self *C.QOffscreenSurface, cb C.intptr_t) *C.QSurfaceFormat {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSurfaceFormat) *QSurfaceFormat)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QOffscreenSurface{h: self}).callVirtualBase_Format)

	return virtualReturn.cPointer()

}

func (this *QOffscreenSurface) callVirtualBase_Size() *QSize {

	_ret := C.QOffscreenSurface_virtualbase_Size(unsafe.Pointer(this.h))
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QOffscreenSurface) OnSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QOffscreenSurface_override_virtual_Size(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_Size
func miqt_exec_callback_QOffscreenSurface_Size(self *C.QOffscreenSurface, cb C.intptr_t) *C.QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QOffscreenSurface{h: self}).callVirtualBase_Size)

	return virtualReturn.cPointer()

}

func (this *QOffscreenSurface) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(C.QOffscreenSurface_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QOffscreenSurface) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QOffscreenSurface_override_virtual_Event(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_Event
func miqt_exec_callback_QOffscreenSurface_Event(self *C.QOffscreenSurface, cb C.intptr_t, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QOffscreenSurface{h: self}).callVirtualBase_Event, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QOffscreenSurface) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(C.QOffscreenSurface_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QOffscreenSurface) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QOffscreenSurface_override_virtual_EventFilter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_EventFilter
func miqt_exec_callback_QOffscreenSurface_EventFilter(self *C.QOffscreenSurface, cb C.intptr_t, watched *C.QObject, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQObject(unsafe.Pointer(watched))
	slotval2 := UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QOffscreenSurface{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QOffscreenSurface) callVirtualBase_TimerEvent(event *QTimerEvent) {

	C.QOffscreenSurface_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QOffscreenSurface) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QOffscreenSurface_override_virtual_TimerEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_TimerEvent
func miqt_exec_callback_QOffscreenSurface_TimerEvent(self *C.QOffscreenSurface, cb C.intptr_t, event *C.QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQTimerEvent(unsafe.Pointer(event), nil)

	gofunc((&QOffscreenSurface{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QOffscreenSurface) callVirtualBase_ChildEvent(event *QChildEvent) {

	C.QOffscreenSurface_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QOffscreenSurface) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QOffscreenSurface_override_virtual_ChildEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_ChildEvent
func miqt_exec_callback_QOffscreenSurface_ChildEvent(self *C.QOffscreenSurface, cb C.intptr_t, event *C.QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQChildEvent(unsafe.Pointer(event), nil)

	gofunc((&QOffscreenSurface{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QOffscreenSurface) callVirtualBase_CustomEvent(event *QEvent) {

	C.QOffscreenSurface_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QOffscreenSurface) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QOffscreenSurface_override_virtual_CustomEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_CustomEvent
func miqt_exec_callback_QOffscreenSurface_CustomEvent(self *C.QOffscreenSurface, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QOffscreenSurface{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QOffscreenSurface) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	C.QOffscreenSurface_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QOffscreenSurface) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QOffscreenSurface_override_virtual_ConnectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_ConnectNotify
func miqt_exec_callback_QOffscreenSurface_ConnectNotify(self *C.QOffscreenSurface, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QOffscreenSurface{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QOffscreenSurface) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	C.QOffscreenSurface_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QOffscreenSurface) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QOffscreenSurface_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_DisconnectNotify
func miqt_exec_callback_QOffscreenSurface_DisconnectNotify(self *C.QOffscreenSurface, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QOffscreenSurface{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

// Delete this object from C++ memory.
func (this *QOffscreenSurface) Delete() {
	C.QOffscreenSurface_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QOffscreenSurface) GoGC() {
	runtime.SetFinalizer(this, func(this *QOffscreenSurface) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
