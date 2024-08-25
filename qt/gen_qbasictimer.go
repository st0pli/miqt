package qt

/*

#cgo CFLAGS: -fPIC
#cgo pkg-config: Qt5Widgets
#include "gen_qbasictimer.h"
#include <stdlib.h>

*/
import "C"

import (
	"unsafe"
)

type QBasicTimer struct {
	h *C.QBasicTimer
}

func (this *QBasicTimer) cPointer() *C.QBasicTimer {
	if this == nil {
		return nil
	}
	return this.h
}

func newQBasicTimer(h *C.QBasicTimer) *QBasicTimer {
	return &QBasicTimer{h: h}
}

func newQBasicTimer_U(h unsafe.Pointer) *QBasicTimer {
	return newQBasicTimer((*C.QBasicTimer)(h))
}

// NewQBasicTimer constructs a new QBasicTimer object.
func NewQBasicTimer(param1 *QBasicTimer) *QBasicTimer {
	ret := C.QBasicTimer_new(param1.cPointer())
	return newQBasicTimer(ret)
}

// NewQBasicTimer2 constructs a new QBasicTimer object.
func NewQBasicTimer2() *QBasicTimer {
	ret := C.QBasicTimer_new2()
	return newQBasicTimer(ret)
}

func (this *QBasicTimer) OperatorAssign(param1 *QBasicTimer) {
	C.QBasicTimer_OperatorAssign(this.h, param1.cPointer())
}

func (this *QBasicTimer) Swap(other *QBasicTimer) {
	C.QBasicTimer_Swap(this.h, other.cPointer())
}

func (this *QBasicTimer) IsActive() bool {
	ret := C.QBasicTimer_IsActive(this.h)
	return (bool)(ret)
}

func (this *QBasicTimer) TimerId() int {
	ret := C.QBasicTimer_TimerId(this.h)
	return (int)(ret)
}

func (this *QBasicTimer) Start(msec int, obj *QObject) {
	C.QBasicTimer_Start(this.h, (C.int)(msec), obj.cPointer())
}

func (this *QBasicTimer) Stop() {
	C.QBasicTimer_Stop(this.h)
}

func (this *QBasicTimer) Delete() {
	C.QBasicTimer_Delete(this.h)
}