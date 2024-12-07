package qt6

/*

#include "gen_qgesture.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QGesture__GestureCancelPolicy int

const (
	QGesture__CancelNone         QGesture__GestureCancelPolicy = 0
	QGesture__CancelAllInContext QGesture__GestureCancelPolicy = 1
)

type QPinchGesture__ChangeFlag int

const (
	QPinchGesture__ScaleFactorChanged   QPinchGesture__ChangeFlag = 1
	QPinchGesture__RotationAngleChanged QPinchGesture__ChangeFlag = 2
	QPinchGesture__CenterPointChanged   QPinchGesture__ChangeFlag = 4
)

type QSwipeGesture__SwipeDirection int

const (
	QSwipeGesture__NoDirection QSwipeGesture__SwipeDirection = 0
	QSwipeGesture__Left        QSwipeGesture__SwipeDirection = 1
	QSwipeGesture__Right       QSwipeGesture__SwipeDirection = 2
	QSwipeGesture__Up          QSwipeGesture__SwipeDirection = 3
	QSwipeGesture__Down        QSwipeGesture__SwipeDirection = 4
)

type QGesture struct {
	h          *C.QGesture
	isSubclass bool
	*QObject
}

func (this *QGesture) cPointer() *C.QGesture {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QGesture) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQGesture constructs the type using only CGO pointers.
func newQGesture(h *C.QGesture, h_QObject *C.QObject) *QGesture {
	if h == nil {
		return nil
	}
	return &QGesture{h: h,
		QObject: newQObject(h_QObject)}
}

// UnsafeNewQGesture constructs the type using only unsafe pointers.
func UnsafeNewQGesture(h unsafe.Pointer, h_QObject unsafe.Pointer) *QGesture {
	if h == nil {
		return nil
	}

	return &QGesture{h: (*C.QGesture)(h),
		QObject: UnsafeNewQObject(h_QObject)}
}

// NewQGesture constructs a new QGesture object.
func NewQGesture() *QGesture {
	var outptr_QGesture *C.QGesture = nil
	var outptr_QObject *C.QObject = nil

	C.QGesture_new(&outptr_QGesture, &outptr_QObject)
	ret := newQGesture(outptr_QGesture, outptr_QObject)
	ret.isSubclass = true
	return ret
}

// NewQGesture2 constructs a new QGesture object.
func NewQGesture2(parent *QObject) *QGesture {
	var outptr_QGesture *C.QGesture = nil
	var outptr_QObject *C.QObject = nil

	C.QGesture_new2(parent.cPointer(), &outptr_QGesture, &outptr_QObject)
	ret := newQGesture(outptr_QGesture, outptr_QObject)
	ret.isSubclass = true
	return ret
}

func (this *QGesture) MetaObject() *QMetaObject {
	return UnsafeNewQMetaObject(unsafe.Pointer(C.QGesture_MetaObject(this.h)))
}

func (this *QGesture) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QGesture_Metacast(this.h, param1_Cstring))
}

func QGesture_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QGesture_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGesture) GestureType() GestureType {
	return (GestureType)(C.QGesture_GestureType(this.h))
}

func (this *QGesture) State() GestureState {
	return (GestureState)(C.QGesture_State(this.h))
}

func (this *QGesture) HotSpot() *QPointF {
	_ret := C.QGesture_HotSpot(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGesture) SetHotSpot(value *QPointF) {
	C.QGesture_SetHotSpot(this.h, value.cPointer())
}

func (this *QGesture) HasHotSpot() bool {
	return (bool)(C.QGesture_HasHotSpot(this.h))
}

func (this *QGesture) UnsetHotSpot() {
	C.QGesture_UnsetHotSpot(this.h)
}

func (this *QGesture) SetGestureCancelPolicy(policy QGesture__GestureCancelPolicy) {
	C.QGesture_SetGestureCancelPolicy(this.h, (C.int)(policy))
}

func (this *QGesture) GestureCancelPolicy() QGesture__GestureCancelPolicy {
	return (QGesture__GestureCancelPolicy)(C.QGesture_GestureCancelPolicy(this.h))
}

func QGesture_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QGesture_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGesture_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QGesture_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGesture) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(C.QGesture_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QGesture) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QGesture_override_virtual_Event(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGesture_Event
func miqt_exec_callback_QGesture_Event(self *C.QGesture, cb C.intptr_t, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QGesture{h: self}).callVirtualBase_Event, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QGesture) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(C.QGesture_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QGesture) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QGesture_override_virtual_EventFilter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGesture_EventFilter
func miqt_exec_callback_QGesture_EventFilter(self *C.QGesture, cb C.intptr_t, watched *C.QObject, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQObject(unsafe.Pointer(watched))
	slotval2 := UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QGesture{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QGesture) callVirtualBase_TimerEvent(event *QTimerEvent) {

	C.QGesture_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGesture) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QGesture_override_virtual_TimerEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGesture_TimerEvent
func miqt_exec_callback_QGesture_TimerEvent(self *C.QGesture, cb C.intptr_t, event *C.QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQTimerEvent(unsafe.Pointer(event), nil)

	gofunc((&QGesture{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QGesture) callVirtualBase_ChildEvent(event *QChildEvent) {

	C.QGesture_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGesture) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QGesture_override_virtual_ChildEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGesture_ChildEvent
func miqt_exec_callback_QGesture_ChildEvent(self *C.QGesture, cb C.intptr_t, event *C.QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQChildEvent(unsafe.Pointer(event), nil)

	gofunc((&QGesture{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QGesture) callVirtualBase_CustomEvent(event *QEvent) {

	C.QGesture_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGesture) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QGesture_override_virtual_CustomEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGesture_CustomEvent
func miqt_exec_callback_QGesture_CustomEvent(self *C.QGesture, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QGesture{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QGesture) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	C.QGesture_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QGesture) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QGesture_override_virtual_ConnectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGesture_ConnectNotify
func miqt_exec_callback_QGesture_ConnectNotify(self *C.QGesture, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QGesture{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QGesture) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	C.QGesture_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QGesture) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QGesture_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGesture_DisconnectNotify
func miqt_exec_callback_QGesture_DisconnectNotify(self *C.QGesture, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QGesture{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

// Delete this object from C++ memory.
func (this *QGesture) Delete() {
	C.QGesture_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QGesture) GoGC() {
	runtime.SetFinalizer(this, func(this *QGesture) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QPanGesture struct {
	h          *C.QPanGesture
	isSubclass bool
	*QGesture
}

func (this *QPanGesture) cPointer() *C.QPanGesture {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QPanGesture) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQPanGesture constructs the type using only CGO pointers.
func newQPanGesture(h *C.QPanGesture, h_QGesture *C.QGesture, h_QObject *C.QObject) *QPanGesture {
	if h == nil {
		return nil
	}
	return &QPanGesture{h: h,
		QGesture: newQGesture(h_QGesture, h_QObject)}
}

// UnsafeNewQPanGesture constructs the type using only unsafe pointers.
func UnsafeNewQPanGesture(h unsafe.Pointer, h_QGesture unsafe.Pointer, h_QObject unsafe.Pointer) *QPanGesture {
	if h == nil {
		return nil
	}

	return &QPanGesture{h: (*C.QPanGesture)(h),
		QGesture: UnsafeNewQGesture(h_QGesture, h_QObject)}
}

// NewQPanGesture constructs a new QPanGesture object.
func NewQPanGesture() *QPanGesture {
	var outptr_QPanGesture *C.QPanGesture = nil
	var outptr_QGesture *C.QGesture = nil
	var outptr_QObject *C.QObject = nil

	C.QPanGesture_new(&outptr_QPanGesture, &outptr_QGesture, &outptr_QObject)
	ret := newQPanGesture(outptr_QPanGesture, outptr_QGesture, outptr_QObject)
	ret.isSubclass = true
	return ret
}

// NewQPanGesture2 constructs a new QPanGesture object.
func NewQPanGesture2(parent *QObject) *QPanGesture {
	var outptr_QPanGesture *C.QPanGesture = nil
	var outptr_QGesture *C.QGesture = nil
	var outptr_QObject *C.QObject = nil

	C.QPanGesture_new2(parent.cPointer(), &outptr_QPanGesture, &outptr_QGesture, &outptr_QObject)
	ret := newQPanGesture(outptr_QPanGesture, outptr_QGesture, outptr_QObject)
	ret.isSubclass = true
	return ret
}

func (this *QPanGesture) MetaObject() *QMetaObject {
	return UnsafeNewQMetaObject(unsafe.Pointer(C.QPanGesture_MetaObject(this.h)))
}

func (this *QPanGesture) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QPanGesture_Metacast(this.h, param1_Cstring))
}

func QPanGesture_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QPanGesture_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPanGesture) LastOffset() *QPointF {
	_ret := C.QPanGesture_LastOffset(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPanGesture) Offset() *QPointF {
	_ret := C.QPanGesture_Offset(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPanGesture) Delta() *QPointF {
	_ret := C.QPanGesture_Delta(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPanGesture) Acceleration() float64 {
	return (float64)(C.QPanGesture_Acceleration(this.h))
}

func (this *QPanGesture) SetLastOffset(value *QPointF) {
	C.QPanGesture_SetLastOffset(this.h, value.cPointer())
}

func (this *QPanGesture) SetOffset(value *QPointF) {
	C.QPanGesture_SetOffset(this.h, value.cPointer())
}

func (this *QPanGesture) SetAcceleration(value float64) {
	C.QPanGesture_SetAcceleration(this.h, (C.double)(value))
}

func QPanGesture_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QPanGesture_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPanGesture_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QPanGesture_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Delete this object from C++ memory.
func (this *QPanGesture) Delete() {
	C.QPanGesture_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QPanGesture) GoGC() {
	runtime.SetFinalizer(this, func(this *QPanGesture) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QPinchGesture struct {
	h          *C.QPinchGesture
	isSubclass bool
	*QGesture
}

func (this *QPinchGesture) cPointer() *C.QPinchGesture {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QPinchGesture) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQPinchGesture constructs the type using only CGO pointers.
func newQPinchGesture(h *C.QPinchGesture, h_QGesture *C.QGesture, h_QObject *C.QObject) *QPinchGesture {
	if h == nil {
		return nil
	}
	return &QPinchGesture{h: h,
		QGesture: newQGesture(h_QGesture, h_QObject)}
}

// UnsafeNewQPinchGesture constructs the type using only unsafe pointers.
func UnsafeNewQPinchGesture(h unsafe.Pointer, h_QGesture unsafe.Pointer, h_QObject unsafe.Pointer) *QPinchGesture {
	if h == nil {
		return nil
	}

	return &QPinchGesture{h: (*C.QPinchGesture)(h),
		QGesture: UnsafeNewQGesture(h_QGesture, h_QObject)}
}

// NewQPinchGesture constructs a new QPinchGesture object.
func NewQPinchGesture() *QPinchGesture {
	var outptr_QPinchGesture *C.QPinchGesture = nil
	var outptr_QGesture *C.QGesture = nil
	var outptr_QObject *C.QObject = nil

	C.QPinchGesture_new(&outptr_QPinchGesture, &outptr_QGesture, &outptr_QObject)
	ret := newQPinchGesture(outptr_QPinchGesture, outptr_QGesture, outptr_QObject)
	ret.isSubclass = true
	return ret
}

// NewQPinchGesture2 constructs a new QPinchGesture object.
func NewQPinchGesture2(parent *QObject) *QPinchGesture {
	var outptr_QPinchGesture *C.QPinchGesture = nil
	var outptr_QGesture *C.QGesture = nil
	var outptr_QObject *C.QObject = nil

	C.QPinchGesture_new2(parent.cPointer(), &outptr_QPinchGesture, &outptr_QGesture, &outptr_QObject)
	ret := newQPinchGesture(outptr_QPinchGesture, outptr_QGesture, outptr_QObject)
	ret.isSubclass = true
	return ret
}

func (this *QPinchGesture) MetaObject() *QMetaObject {
	return UnsafeNewQMetaObject(unsafe.Pointer(C.QPinchGesture_MetaObject(this.h)))
}

func (this *QPinchGesture) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QPinchGesture_Metacast(this.h, param1_Cstring))
}

func QPinchGesture_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QPinchGesture_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPinchGesture) TotalChangeFlags() QPinchGesture__ChangeFlag {
	return (QPinchGesture__ChangeFlag)(C.QPinchGesture_TotalChangeFlags(this.h))
}

func (this *QPinchGesture) SetTotalChangeFlags(value QPinchGesture__ChangeFlag) {
	C.QPinchGesture_SetTotalChangeFlags(this.h, (C.int)(value))
}

func (this *QPinchGesture) ChangeFlags() QPinchGesture__ChangeFlag {
	return (QPinchGesture__ChangeFlag)(C.QPinchGesture_ChangeFlags(this.h))
}

func (this *QPinchGesture) SetChangeFlags(value QPinchGesture__ChangeFlag) {
	C.QPinchGesture_SetChangeFlags(this.h, (C.int)(value))
}

func (this *QPinchGesture) StartCenterPoint() *QPointF {
	_ret := C.QPinchGesture_StartCenterPoint(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPinchGesture) LastCenterPoint() *QPointF {
	_ret := C.QPinchGesture_LastCenterPoint(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPinchGesture) CenterPoint() *QPointF {
	_ret := C.QPinchGesture_CenterPoint(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPinchGesture) SetStartCenterPoint(value *QPointF) {
	C.QPinchGesture_SetStartCenterPoint(this.h, value.cPointer())
}

func (this *QPinchGesture) SetLastCenterPoint(value *QPointF) {
	C.QPinchGesture_SetLastCenterPoint(this.h, value.cPointer())
}

func (this *QPinchGesture) SetCenterPoint(value *QPointF) {
	C.QPinchGesture_SetCenterPoint(this.h, value.cPointer())
}

func (this *QPinchGesture) TotalScaleFactor() float64 {
	return (float64)(C.QPinchGesture_TotalScaleFactor(this.h))
}

func (this *QPinchGesture) LastScaleFactor() float64 {
	return (float64)(C.QPinchGesture_LastScaleFactor(this.h))
}

func (this *QPinchGesture) ScaleFactor() float64 {
	return (float64)(C.QPinchGesture_ScaleFactor(this.h))
}

func (this *QPinchGesture) SetTotalScaleFactor(value float64) {
	C.QPinchGesture_SetTotalScaleFactor(this.h, (C.double)(value))
}

func (this *QPinchGesture) SetLastScaleFactor(value float64) {
	C.QPinchGesture_SetLastScaleFactor(this.h, (C.double)(value))
}

func (this *QPinchGesture) SetScaleFactor(value float64) {
	C.QPinchGesture_SetScaleFactor(this.h, (C.double)(value))
}

func (this *QPinchGesture) TotalRotationAngle() float64 {
	return (float64)(C.QPinchGesture_TotalRotationAngle(this.h))
}

func (this *QPinchGesture) LastRotationAngle() float64 {
	return (float64)(C.QPinchGesture_LastRotationAngle(this.h))
}

func (this *QPinchGesture) RotationAngle() float64 {
	return (float64)(C.QPinchGesture_RotationAngle(this.h))
}

func (this *QPinchGesture) SetTotalRotationAngle(value float64) {
	C.QPinchGesture_SetTotalRotationAngle(this.h, (C.double)(value))
}

func (this *QPinchGesture) SetLastRotationAngle(value float64) {
	C.QPinchGesture_SetLastRotationAngle(this.h, (C.double)(value))
}

func (this *QPinchGesture) SetRotationAngle(value float64) {
	C.QPinchGesture_SetRotationAngle(this.h, (C.double)(value))
}

func QPinchGesture_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QPinchGesture_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPinchGesture_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QPinchGesture_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Delete this object from C++ memory.
func (this *QPinchGesture) Delete() {
	C.QPinchGesture_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QPinchGesture) GoGC() {
	runtime.SetFinalizer(this, func(this *QPinchGesture) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QSwipeGesture struct {
	h          *C.QSwipeGesture
	isSubclass bool
	*QGesture
}

func (this *QSwipeGesture) cPointer() *C.QSwipeGesture {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QSwipeGesture) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQSwipeGesture constructs the type using only CGO pointers.
func newQSwipeGesture(h *C.QSwipeGesture, h_QGesture *C.QGesture, h_QObject *C.QObject) *QSwipeGesture {
	if h == nil {
		return nil
	}
	return &QSwipeGesture{h: h,
		QGesture: newQGesture(h_QGesture, h_QObject)}
}

// UnsafeNewQSwipeGesture constructs the type using only unsafe pointers.
func UnsafeNewQSwipeGesture(h unsafe.Pointer, h_QGesture unsafe.Pointer, h_QObject unsafe.Pointer) *QSwipeGesture {
	if h == nil {
		return nil
	}

	return &QSwipeGesture{h: (*C.QSwipeGesture)(h),
		QGesture: UnsafeNewQGesture(h_QGesture, h_QObject)}
}

// NewQSwipeGesture constructs a new QSwipeGesture object.
func NewQSwipeGesture() *QSwipeGesture {
	var outptr_QSwipeGesture *C.QSwipeGesture = nil
	var outptr_QGesture *C.QGesture = nil
	var outptr_QObject *C.QObject = nil

	C.QSwipeGesture_new(&outptr_QSwipeGesture, &outptr_QGesture, &outptr_QObject)
	ret := newQSwipeGesture(outptr_QSwipeGesture, outptr_QGesture, outptr_QObject)
	ret.isSubclass = true
	return ret
}

// NewQSwipeGesture2 constructs a new QSwipeGesture object.
func NewQSwipeGesture2(parent *QObject) *QSwipeGesture {
	var outptr_QSwipeGesture *C.QSwipeGesture = nil
	var outptr_QGesture *C.QGesture = nil
	var outptr_QObject *C.QObject = nil

	C.QSwipeGesture_new2(parent.cPointer(), &outptr_QSwipeGesture, &outptr_QGesture, &outptr_QObject)
	ret := newQSwipeGesture(outptr_QSwipeGesture, outptr_QGesture, outptr_QObject)
	ret.isSubclass = true
	return ret
}

func (this *QSwipeGesture) MetaObject() *QMetaObject {
	return UnsafeNewQMetaObject(unsafe.Pointer(C.QSwipeGesture_MetaObject(this.h)))
}

func (this *QSwipeGesture) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QSwipeGesture_Metacast(this.h, param1_Cstring))
}

func QSwipeGesture_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QSwipeGesture_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSwipeGesture) HorizontalDirection() QSwipeGesture__SwipeDirection {
	return (QSwipeGesture__SwipeDirection)(C.QSwipeGesture_HorizontalDirection(this.h))
}

func (this *QSwipeGesture) VerticalDirection() QSwipeGesture__SwipeDirection {
	return (QSwipeGesture__SwipeDirection)(C.QSwipeGesture_VerticalDirection(this.h))
}

func (this *QSwipeGesture) SwipeAngle() float64 {
	return (float64)(C.QSwipeGesture_SwipeAngle(this.h))
}

func (this *QSwipeGesture) SetSwipeAngle(value float64) {
	C.QSwipeGesture_SetSwipeAngle(this.h, (C.double)(value))
}

func QSwipeGesture_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QSwipeGesture_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSwipeGesture_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QSwipeGesture_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Delete this object from C++ memory.
func (this *QSwipeGesture) Delete() {
	C.QSwipeGesture_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QSwipeGesture) GoGC() {
	runtime.SetFinalizer(this, func(this *QSwipeGesture) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QTapGesture struct {
	h          *C.QTapGesture
	isSubclass bool
	*QGesture
}

func (this *QTapGesture) cPointer() *C.QTapGesture {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QTapGesture) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQTapGesture constructs the type using only CGO pointers.
func newQTapGesture(h *C.QTapGesture, h_QGesture *C.QGesture, h_QObject *C.QObject) *QTapGesture {
	if h == nil {
		return nil
	}
	return &QTapGesture{h: h,
		QGesture: newQGesture(h_QGesture, h_QObject)}
}

// UnsafeNewQTapGesture constructs the type using only unsafe pointers.
func UnsafeNewQTapGesture(h unsafe.Pointer, h_QGesture unsafe.Pointer, h_QObject unsafe.Pointer) *QTapGesture {
	if h == nil {
		return nil
	}

	return &QTapGesture{h: (*C.QTapGesture)(h),
		QGesture: UnsafeNewQGesture(h_QGesture, h_QObject)}
}

// NewQTapGesture constructs a new QTapGesture object.
func NewQTapGesture() *QTapGesture {
	var outptr_QTapGesture *C.QTapGesture = nil
	var outptr_QGesture *C.QGesture = nil
	var outptr_QObject *C.QObject = nil

	C.QTapGesture_new(&outptr_QTapGesture, &outptr_QGesture, &outptr_QObject)
	ret := newQTapGesture(outptr_QTapGesture, outptr_QGesture, outptr_QObject)
	ret.isSubclass = true
	return ret
}

// NewQTapGesture2 constructs a new QTapGesture object.
func NewQTapGesture2(parent *QObject) *QTapGesture {
	var outptr_QTapGesture *C.QTapGesture = nil
	var outptr_QGesture *C.QGesture = nil
	var outptr_QObject *C.QObject = nil

	C.QTapGesture_new2(parent.cPointer(), &outptr_QTapGesture, &outptr_QGesture, &outptr_QObject)
	ret := newQTapGesture(outptr_QTapGesture, outptr_QGesture, outptr_QObject)
	ret.isSubclass = true
	return ret
}

func (this *QTapGesture) MetaObject() *QMetaObject {
	return UnsafeNewQMetaObject(unsafe.Pointer(C.QTapGesture_MetaObject(this.h)))
}

func (this *QTapGesture) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QTapGesture_Metacast(this.h, param1_Cstring))
}

func QTapGesture_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QTapGesture_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTapGesture) Position() *QPointF {
	_ret := C.QTapGesture_Position(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTapGesture) SetPosition(pos *QPointF) {
	C.QTapGesture_SetPosition(this.h, pos.cPointer())
}

func QTapGesture_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QTapGesture_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTapGesture_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QTapGesture_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Delete this object from C++ memory.
func (this *QTapGesture) Delete() {
	C.QTapGesture_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QTapGesture) GoGC() {
	runtime.SetFinalizer(this, func(this *QTapGesture) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QTapAndHoldGesture struct {
	h          *C.QTapAndHoldGesture
	isSubclass bool
	*QGesture
}

func (this *QTapAndHoldGesture) cPointer() *C.QTapAndHoldGesture {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QTapAndHoldGesture) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQTapAndHoldGesture constructs the type using only CGO pointers.
func newQTapAndHoldGesture(h *C.QTapAndHoldGesture, h_QGesture *C.QGesture, h_QObject *C.QObject) *QTapAndHoldGesture {
	if h == nil {
		return nil
	}
	return &QTapAndHoldGesture{h: h,
		QGesture: newQGesture(h_QGesture, h_QObject)}
}

// UnsafeNewQTapAndHoldGesture constructs the type using only unsafe pointers.
func UnsafeNewQTapAndHoldGesture(h unsafe.Pointer, h_QGesture unsafe.Pointer, h_QObject unsafe.Pointer) *QTapAndHoldGesture {
	if h == nil {
		return nil
	}

	return &QTapAndHoldGesture{h: (*C.QTapAndHoldGesture)(h),
		QGesture: UnsafeNewQGesture(h_QGesture, h_QObject)}
}

// NewQTapAndHoldGesture constructs a new QTapAndHoldGesture object.
func NewQTapAndHoldGesture() *QTapAndHoldGesture {
	var outptr_QTapAndHoldGesture *C.QTapAndHoldGesture = nil
	var outptr_QGesture *C.QGesture = nil
	var outptr_QObject *C.QObject = nil

	C.QTapAndHoldGesture_new(&outptr_QTapAndHoldGesture, &outptr_QGesture, &outptr_QObject)
	ret := newQTapAndHoldGesture(outptr_QTapAndHoldGesture, outptr_QGesture, outptr_QObject)
	ret.isSubclass = true
	return ret
}

// NewQTapAndHoldGesture2 constructs a new QTapAndHoldGesture object.
func NewQTapAndHoldGesture2(parent *QObject) *QTapAndHoldGesture {
	var outptr_QTapAndHoldGesture *C.QTapAndHoldGesture = nil
	var outptr_QGesture *C.QGesture = nil
	var outptr_QObject *C.QObject = nil

	C.QTapAndHoldGesture_new2(parent.cPointer(), &outptr_QTapAndHoldGesture, &outptr_QGesture, &outptr_QObject)
	ret := newQTapAndHoldGesture(outptr_QTapAndHoldGesture, outptr_QGesture, outptr_QObject)
	ret.isSubclass = true
	return ret
}

func (this *QTapAndHoldGesture) MetaObject() *QMetaObject {
	return UnsafeNewQMetaObject(unsafe.Pointer(C.QTapAndHoldGesture_MetaObject(this.h)))
}

func (this *QTapAndHoldGesture) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QTapAndHoldGesture_Metacast(this.h, param1_Cstring))
}

func QTapAndHoldGesture_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QTapAndHoldGesture_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTapAndHoldGesture) Position() *QPointF {
	_ret := C.QTapAndHoldGesture_Position(this.h)
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTapAndHoldGesture) SetPosition(pos *QPointF) {
	C.QTapAndHoldGesture_SetPosition(this.h, pos.cPointer())
}

func QTapAndHoldGesture_SetTimeout(msecs int) {
	C.QTapAndHoldGesture_SetTimeout((C.int)(msecs))
}

func QTapAndHoldGesture_Timeout() int {
	return (int)(C.QTapAndHoldGesture_Timeout())
}

func QTapAndHoldGesture_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QTapAndHoldGesture_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTapAndHoldGesture_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QTapAndHoldGesture_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Delete this object from C++ memory.
func (this *QTapAndHoldGesture) Delete() {
	C.QTapAndHoldGesture_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QTapAndHoldGesture) GoGC() {
	runtime.SetFinalizer(this, func(this *QTapAndHoldGesture) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QGestureEvent struct {
	h          *C.QGestureEvent
	isSubclass bool
	*QEvent
}

func (this *QGestureEvent) cPointer() *C.QGestureEvent {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QGestureEvent) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQGestureEvent constructs the type using only CGO pointers.
func newQGestureEvent(h *C.QGestureEvent, h_QEvent *C.QEvent) *QGestureEvent {
	if h == nil {
		return nil
	}
	return &QGestureEvent{h: h,
		QEvent: newQEvent(h_QEvent)}
}

// UnsafeNewQGestureEvent constructs the type using only unsafe pointers.
func UnsafeNewQGestureEvent(h unsafe.Pointer, h_QEvent unsafe.Pointer) *QGestureEvent {
	if h == nil {
		return nil
	}

	return &QGestureEvent{h: (*C.QGestureEvent)(h),
		QEvent: UnsafeNewQEvent(h_QEvent)}
}

// NewQGestureEvent constructs a new QGestureEvent object.
func NewQGestureEvent(gestures []*QGesture) *QGestureEvent {
	gestures_CArray := (*[0xffff]*C.QGesture)(C.malloc(C.size_t(8 * len(gestures))))
	defer C.free(unsafe.Pointer(gestures_CArray))
	for i := range gestures {
		gestures_CArray[i] = gestures[i].cPointer()
	}
	gestures_ma := C.struct_miqt_array{len: C.size_t(len(gestures)), data: unsafe.Pointer(gestures_CArray)}
	var outptr_QGestureEvent *C.QGestureEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QGestureEvent_new(gestures_ma, &outptr_QGestureEvent, &outptr_QEvent)
	ret := newQGestureEvent(outptr_QGestureEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

// NewQGestureEvent2 constructs a new QGestureEvent object.
func NewQGestureEvent2(param1 *QGestureEvent) *QGestureEvent {
	var outptr_QGestureEvent *C.QGestureEvent = nil
	var outptr_QEvent *C.QEvent = nil

	C.QGestureEvent_new2(param1.cPointer(), &outptr_QGestureEvent, &outptr_QEvent)
	ret := newQGestureEvent(outptr_QGestureEvent, outptr_QEvent)
	ret.isSubclass = true
	return ret
}

func (this *QGestureEvent) Gestures() []*QGesture {
	var _ma C.struct_miqt_array = C.QGestureEvent_Gestures(this.h)
	_ret := make([]*QGesture, int(_ma.len))
	_outCast := (*[0xffff]*C.QGesture)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = UnsafeNewQGesture(unsafe.Pointer(_outCast[i]), nil)
	}
	return _ret
}

func (this *QGestureEvent) Gesture(typeVal GestureType) *QGesture {
	return UnsafeNewQGesture(unsafe.Pointer(C.QGestureEvent_Gesture(this.h, (C.int)(typeVal))), nil)
}

func (this *QGestureEvent) ActiveGestures() []*QGesture {
	var _ma C.struct_miqt_array = C.QGestureEvent_ActiveGestures(this.h)
	_ret := make([]*QGesture, int(_ma.len))
	_outCast := (*[0xffff]*C.QGesture)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = UnsafeNewQGesture(unsafe.Pointer(_outCast[i]), nil)
	}
	return _ret
}

func (this *QGestureEvent) CanceledGestures() []*QGesture {
	var _ma C.struct_miqt_array = C.QGestureEvent_CanceledGestures(this.h)
	_ret := make([]*QGesture, int(_ma.len))
	_outCast := (*[0xffff]*C.QGesture)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = UnsafeNewQGesture(unsafe.Pointer(_outCast[i]), nil)
	}
	return _ret
}

func (this *QGestureEvent) SetAccepted(param1 *QGesture, param2 bool) {
	C.QGestureEvent_SetAccepted(this.h, param1.cPointer(), (C.bool)(param2))
}

func (this *QGestureEvent) Accept(param1 *QGesture) {
	C.QGestureEvent_Accept(this.h, param1.cPointer())
}

func (this *QGestureEvent) Ignore(param1 *QGesture) {
	C.QGestureEvent_Ignore(this.h, param1.cPointer())
}

func (this *QGestureEvent) IsAccepted(param1 *QGesture) bool {
	return (bool)(C.QGestureEvent_IsAccepted(this.h, param1.cPointer()))
}

func (this *QGestureEvent) SetAccepted2(param1 GestureType, param2 bool) {
	C.QGestureEvent_SetAccepted2(this.h, (C.int)(param1), (C.bool)(param2))
}

func (this *QGestureEvent) AcceptWithQtGestureType(param1 GestureType) {
	C.QGestureEvent_AcceptWithQtGestureType(this.h, (C.int)(param1))
}

func (this *QGestureEvent) IgnoreWithQtGestureType(param1 GestureType) {
	C.QGestureEvent_IgnoreWithQtGestureType(this.h, (C.int)(param1))
}

func (this *QGestureEvent) IsAcceptedWithQtGestureType(param1 GestureType) bool {
	return (bool)(C.QGestureEvent_IsAcceptedWithQtGestureType(this.h, (C.int)(param1)))
}

func (this *QGestureEvent) SetWidget(widget *QWidget) {
	C.QGestureEvent_SetWidget(this.h, widget.cPointer())
}

func (this *QGestureEvent) Widget() *QWidget {
	return UnsafeNewQWidget(unsafe.Pointer(C.QGestureEvent_Widget(this.h)), nil, nil)
}

func (this *QGestureEvent) MapToGraphicsScene(gesturePoint *QPointF) *QPointF {
	_ret := C.QGestureEvent_MapToGraphicsScene(this.h, gesturePoint.cPointer())
	_goptr := newQPointF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGestureEvent) callVirtualBase_SetAccepted(accepted bool) {

	C.QGestureEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (C.bool)(accepted))

}
func (this *QGestureEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QGestureEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGestureEvent_SetAccepted
func miqt_exec_callback_QGestureEvent_SetAccepted(self *C.QGestureEvent, cb C.intptr_t, accepted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QGestureEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

func (this *QGestureEvent) callVirtualBase_Clone() *QEvent {

	return UnsafeNewQEvent(unsafe.Pointer(C.QGestureEvent_virtualbase_Clone(unsafe.Pointer(this.h))))
}
func (this *QGestureEvent) OnClone(slot func(super func() *QEvent) *QEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	C.QGestureEvent_override_virtual_Clone(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGestureEvent_Clone
func miqt_exec_callback_QGestureEvent_Clone(self *C.QGestureEvent, cb C.intptr_t) *C.QEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QEvent) *QEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGestureEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

// Delete this object from C++ memory.
func (this *QGestureEvent) Delete() {
	C.QGestureEvent_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QGestureEvent) GoGC() {
	runtime.SetFinalizer(this, func(this *QGestureEvent) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
