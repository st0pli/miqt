package qt

/*

#cgo CFLAGS: -fPIC
#cgo pkg-config: Qt5Widgets
#include "gen_qabstractscrollarea.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QAbstractScrollArea struct {
	h *C.QAbstractScrollArea
	*QFrame
}

func (this *QAbstractScrollArea) cPointer() *C.QAbstractScrollArea {
	if this == nil {
		return nil
	}
	return this.h
}

func newQAbstractScrollArea(h *C.QAbstractScrollArea) *QAbstractScrollArea {
	return &QAbstractScrollArea{h: h, QFrame: newQFrame_U(unsafe.Pointer(h))}
}

func newQAbstractScrollArea_U(h unsafe.Pointer) *QAbstractScrollArea {
	return newQAbstractScrollArea((*C.QAbstractScrollArea)(h))
}

// NewQAbstractScrollArea constructs a new QAbstractScrollArea object.
func NewQAbstractScrollArea() *QAbstractScrollArea {
	ret := C.QAbstractScrollArea_new()
	return newQAbstractScrollArea(ret)
}

// NewQAbstractScrollArea2 constructs a new QAbstractScrollArea object.
func NewQAbstractScrollArea2(parent *QWidget) *QAbstractScrollArea {
	ret := C.QAbstractScrollArea_new2(parent.cPointer())
	return newQAbstractScrollArea(ret)
}

func (this *QAbstractScrollArea) MetaObject() *QMetaObject {
	ret := C.QAbstractScrollArea_MetaObject(this.h)
	return newQMetaObject_U(unsafe.Pointer(ret))
}

func QAbstractScrollArea_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QAbstractScrollArea_Tr(s_Cstring, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QAbstractScrollArea_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QAbstractScrollArea_TrUtf8(s_Cstring, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QAbstractScrollArea) VerticalScrollBar() *QScrollBar {
	ret := C.QAbstractScrollArea_VerticalScrollBar(this.h)
	return newQScrollBar_U(unsafe.Pointer(ret))
}

func (this *QAbstractScrollArea) SetVerticalScrollBar(scrollbar *QScrollBar) {
	C.QAbstractScrollArea_SetVerticalScrollBar(this.h, scrollbar.cPointer())
}

func (this *QAbstractScrollArea) HorizontalScrollBar() *QScrollBar {
	ret := C.QAbstractScrollArea_HorizontalScrollBar(this.h)
	return newQScrollBar_U(unsafe.Pointer(ret))
}

func (this *QAbstractScrollArea) SetHorizontalScrollBar(scrollbar *QScrollBar) {
	C.QAbstractScrollArea_SetHorizontalScrollBar(this.h, scrollbar.cPointer())
}

func (this *QAbstractScrollArea) CornerWidget() *QWidget {
	ret := C.QAbstractScrollArea_CornerWidget(this.h)
	return newQWidget_U(unsafe.Pointer(ret))
}

func (this *QAbstractScrollArea) SetCornerWidget(widget *QWidget) {
	C.QAbstractScrollArea_SetCornerWidget(this.h, widget.cPointer())
}

func (this *QAbstractScrollArea) Viewport() *QWidget {
	ret := C.QAbstractScrollArea_Viewport(this.h)
	return newQWidget_U(unsafe.Pointer(ret))
}

func (this *QAbstractScrollArea) SetViewport(widget *QWidget) {
	C.QAbstractScrollArea_SetViewport(this.h, widget.cPointer())
}

func (this *QAbstractScrollArea) MaximumViewportSize() *QSize {
	ret := C.QAbstractScrollArea_MaximumViewportSize(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQSize(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QSize) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QAbstractScrollArea) MinimumSizeHint() *QSize {
	ret := C.QAbstractScrollArea_MinimumSizeHint(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQSize(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QSize) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QAbstractScrollArea) SizeHint() *QSize {
	ret := C.QAbstractScrollArea_SizeHint(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQSize(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QSize) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QAbstractScrollArea) SetupViewport(viewport *QWidget) {
	C.QAbstractScrollArea_SetupViewport(this.h, viewport.cPointer())
}

func QAbstractScrollArea_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QAbstractScrollArea_Tr2(s_Cstring, c_Cstring, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QAbstractScrollArea_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QAbstractScrollArea_Tr3(s_Cstring, c_Cstring, (C.int)(n), &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QAbstractScrollArea_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QAbstractScrollArea_TrUtf82(s_Cstring, c_Cstring, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QAbstractScrollArea_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QAbstractScrollArea_TrUtf83(s_Cstring, c_Cstring, (C.int)(n), &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QAbstractScrollArea) Delete() {
	C.QAbstractScrollArea_Delete(this.h)
}