package webkit

/*

#include "gen_qwebfullscreenrequest.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt"
	"runtime"
	"unsafe"
)

type QWebFullScreenRequest struct {
	h          *C.QWebFullScreenRequest
	isSubclass bool
}

func (this *QWebFullScreenRequest) cPointer() *C.QWebFullScreenRequest {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QWebFullScreenRequest) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQWebFullScreenRequest constructs the type using only CGO pointers.
func newQWebFullScreenRequest(h *C.QWebFullScreenRequest) *QWebFullScreenRequest {
	if h == nil {
		return nil
	}
	return &QWebFullScreenRequest{h: h}
}

// UnsafeNewQWebFullScreenRequest constructs the type using only unsafe pointers.
func UnsafeNewQWebFullScreenRequest(h unsafe.Pointer) *QWebFullScreenRequest {
	if h == nil {
		return nil
	}

	return &QWebFullScreenRequest{h: (*C.QWebFullScreenRequest)(h)}
}

// NewQWebFullScreenRequest constructs a new QWebFullScreenRequest object.
func NewQWebFullScreenRequest() *QWebFullScreenRequest {
	var outptr_QWebFullScreenRequest *C.QWebFullScreenRequest = nil

	C.QWebFullScreenRequest_new(&outptr_QWebFullScreenRequest)
	ret := newQWebFullScreenRequest(outptr_QWebFullScreenRequest)
	ret.isSubclass = true
	return ret
}

// NewQWebFullScreenRequest2 constructs a new QWebFullScreenRequest object.
func NewQWebFullScreenRequest2(param1 *QWebFullScreenRequest) *QWebFullScreenRequest {
	var outptr_QWebFullScreenRequest *C.QWebFullScreenRequest = nil

	C.QWebFullScreenRequest_new2(param1.cPointer(), &outptr_QWebFullScreenRequest)
	ret := newQWebFullScreenRequest(outptr_QWebFullScreenRequest)
	ret.isSubclass = true
	return ret
}

func (this *QWebFullScreenRequest) Accept() {
	C.QWebFullScreenRequest_Accept(this.h)
}

func (this *QWebFullScreenRequest) Reject() {
	C.QWebFullScreenRequest_Reject(this.h)
}

func (this *QWebFullScreenRequest) ToggleOn() bool {
	return (bool)(C.QWebFullScreenRequest_ToggleOn(this.h))
}

func (this *QWebFullScreenRequest) Origin() *qt.QUrl {
	_ret := C.QWebFullScreenRequest_Origin(this.h)
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(_ret))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWebFullScreenRequest) Element() *QWebElement {
	return UnsafeNewQWebElement(unsafe.Pointer(C.QWebFullScreenRequest_Element(this.h)))
}

// Delete this object from C++ memory.
func (this *QWebFullScreenRequest) Delete() {
	C.QWebFullScreenRequest_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QWebFullScreenRequest) GoGC() {
	runtime.SetFinalizer(this, func(this *QWebFullScreenRequest) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}