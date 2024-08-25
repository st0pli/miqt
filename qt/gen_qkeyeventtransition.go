package qt

/*

#cgo CFLAGS: -fPIC
#cgo pkg-config: Qt5Widgets
#include "gen_qkeyeventtransition.h"
#include <stdlib.h>

*/
import "C"

import (
	"unsafe"
)

type QKeyEventTransition struct {
	h *C.QKeyEventTransition
	*QEventTransition
}

func (this *QKeyEventTransition) cPointer() *C.QKeyEventTransition {
	if this == nil {
		return nil
	}
	return this.h
}

func newQKeyEventTransition(h *C.QKeyEventTransition) *QKeyEventTransition {
	return &QKeyEventTransition{h: h, QEventTransition: newQEventTransition_U(unsafe.Pointer(h))}
}

func newQKeyEventTransition_U(h unsafe.Pointer) *QKeyEventTransition {
	return newQKeyEventTransition((*C.QKeyEventTransition)(h))
}

// NewQKeyEventTransition constructs a new QKeyEventTransition object.
func NewQKeyEventTransition() *QKeyEventTransition {
	ret := C.QKeyEventTransition_new()
	return newQKeyEventTransition(ret)
}

// NewQKeyEventTransition2 constructs a new QKeyEventTransition object.
func NewQKeyEventTransition2(sourceState *QState) *QKeyEventTransition {
	ret := C.QKeyEventTransition_new2(sourceState.cPointer())
	return newQKeyEventTransition(ret)
}

func (this *QKeyEventTransition) MetaObject() *QMetaObject {
	ret := C.QKeyEventTransition_MetaObject(this.h)
	return newQMetaObject_U(unsafe.Pointer(ret))
}

func QKeyEventTransition_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QKeyEventTransition_Tr(s_Cstring, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QKeyEventTransition_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QKeyEventTransition_TrUtf8(s_Cstring, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QKeyEventTransition) Key() int {
	ret := C.QKeyEventTransition_Key(this.h)
	return (int)(ret)
}

func (this *QKeyEventTransition) SetKey(key int) {
	C.QKeyEventTransition_SetKey(this.h, (C.int)(key))
}

func QKeyEventTransition_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QKeyEventTransition_Tr2(s_Cstring, c_Cstring, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QKeyEventTransition_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QKeyEventTransition_Tr3(s_Cstring, c_Cstring, (C.int)(n), &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QKeyEventTransition_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QKeyEventTransition_TrUtf82(s_Cstring, c_Cstring, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QKeyEventTransition_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QKeyEventTransition_TrUtf83(s_Cstring, c_Cstring, (C.int)(n), &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QKeyEventTransition) Delete() {
	C.QKeyEventTransition_Delete(this.h)
}