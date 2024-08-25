package qt

/*

#cgo CFLAGS: -fPIC
#cgo pkg-config: Qt5Widgets
#include "gen_qvector3d.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QVector3D struct {
	h *C.QVector3D
}

func (this *QVector3D) cPointer() *C.QVector3D {
	if this == nil {
		return nil
	}
	return this.h
}

func newQVector3D(h *C.QVector3D) *QVector3D {
	return &QVector3D{h: h}
}

func newQVector3D_U(h unsafe.Pointer) *QVector3D {
	return newQVector3D((*C.QVector3D)(h))
}

// NewQVector3D constructs a new QVector3D object.
func NewQVector3D() *QVector3D {
	ret := C.QVector3D_new()
	return newQVector3D(ret)
}

// NewQVector3D2 constructs a new QVector3D object.
func NewQVector3D2(xpos float32, ypos float32, zpos float32) *QVector3D {
	ret := C.QVector3D_new2((C.float)(xpos), (C.float)(ypos), (C.float)(zpos))
	return newQVector3D(ret)
}

// NewQVector3D3 constructs a new QVector3D object.
func NewQVector3D3(point *QPoint) *QVector3D {
	ret := C.QVector3D_new3(point.cPointer())
	return newQVector3D(ret)
}

// NewQVector3D4 constructs a new QVector3D object.
func NewQVector3D4(point *QPointF) *QVector3D {
	ret := C.QVector3D_new4(point.cPointer())
	return newQVector3D(ret)
}

// NewQVector3D5 constructs a new QVector3D object.
func NewQVector3D5(vector *QVector2D) *QVector3D {
	ret := C.QVector3D_new5(vector.cPointer())
	return newQVector3D(ret)
}

// NewQVector3D6 constructs a new QVector3D object.
func NewQVector3D6(vector *QVector2D, zpos float32) *QVector3D {
	ret := C.QVector3D_new6(vector.cPointer(), (C.float)(zpos))
	return newQVector3D(ret)
}

// NewQVector3D7 constructs a new QVector3D object.
func NewQVector3D7(vector *QVector4D) *QVector3D {
	ret := C.QVector3D_new7(vector.cPointer())
	return newQVector3D(ret)
}

// NewQVector3D8 constructs a new QVector3D object.
func NewQVector3D8(param1 *QVector3D) *QVector3D {
	ret := C.QVector3D_new8(param1.cPointer())
	return newQVector3D(ret)
}

func (this *QVector3D) IsNull() bool {
	ret := C.QVector3D_IsNull(this.h)
	return (bool)(ret)
}

func (this *QVector3D) X() float32 {
	ret := C.QVector3D_X(this.h)
	return (float32)(ret)
}

func (this *QVector3D) Y() float32 {
	ret := C.QVector3D_Y(this.h)
	return (float32)(ret)
}

func (this *QVector3D) Z() float32 {
	ret := C.QVector3D_Z(this.h)
	return (float32)(ret)
}

func (this *QVector3D) SetX(x float32) {
	C.QVector3D_SetX(this.h, (C.float)(x))
}

func (this *QVector3D) SetY(y float32) {
	C.QVector3D_SetY(this.h, (C.float)(y))
}

func (this *QVector3D) SetZ(z float32) {
	C.QVector3D_SetZ(this.h, (C.float)(z))
}

func (this *QVector3D) OperatorSubscript(i int) float32 {
	ret := C.QVector3D_OperatorSubscript(this.h, (C.int)(i))
	return (float32)(ret)
}

func (this *QVector3D) Length() float32 {
	ret := C.QVector3D_Length(this.h)
	return (float32)(ret)
}

func (this *QVector3D) LengthSquared() float32 {
	ret := C.QVector3D_LengthSquared(this.h)
	return (float32)(ret)
}

func (this *QVector3D) Normalized() *QVector3D {
	ret := C.QVector3D_Normalized(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQVector3D(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QVector3D) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QVector3D) Normalize() {
	C.QVector3D_Normalize(this.h)
}

func (this *QVector3D) OperatorPlusAssign(vector *QVector3D) *QVector3D {
	ret := C.QVector3D_OperatorPlusAssign(this.h, vector.cPointer())
	return newQVector3D_U(unsafe.Pointer(ret))
}

func (this *QVector3D) OperatorMinusAssign(vector *QVector3D) *QVector3D {
	ret := C.QVector3D_OperatorMinusAssign(this.h, vector.cPointer())
	return newQVector3D_U(unsafe.Pointer(ret))
}

func (this *QVector3D) OperatorMultiplyAssign(factor float32) *QVector3D {
	ret := C.QVector3D_OperatorMultiplyAssign(this.h, (C.float)(factor))
	return newQVector3D_U(unsafe.Pointer(ret))
}

func (this *QVector3D) OperatorMultiplyAssignWithVector(vector *QVector3D) *QVector3D {
	ret := C.QVector3D_OperatorMultiplyAssignWithVector(this.h, vector.cPointer())
	return newQVector3D_U(unsafe.Pointer(ret))
}

func (this *QVector3D) OperatorDivideAssign(divisor float32) *QVector3D {
	ret := C.QVector3D_OperatorDivideAssign(this.h, (C.float)(divisor))
	return newQVector3D_U(unsafe.Pointer(ret))
}

func (this *QVector3D) OperatorDivideAssignWithVector(vector *QVector3D) *QVector3D {
	ret := C.QVector3D_OperatorDivideAssignWithVector(this.h, vector.cPointer())
	return newQVector3D_U(unsafe.Pointer(ret))
}

func QVector3D_DotProduct(v1 *QVector3D, v2 *QVector3D) float32 {
	ret := C.QVector3D_DotProduct(v1.cPointer(), v2.cPointer())
	return (float32)(ret)
}

func QVector3D_CrossProduct(v1 *QVector3D, v2 *QVector3D) *QVector3D {
	ret := C.QVector3D_CrossProduct(v1.cPointer(), v2.cPointer())
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQVector3D(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QVector3D) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func QVector3D_Normal(v1 *QVector3D, v2 *QVector3D) *QVector3D {
	ret := C.QVector3D_Normal(v1.cPointer(), v2.cPointer())
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQVector3D(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QVector3D) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func QVector3D_Normal2(v1 *QVector3D, v2 *QVector3D, v3 *QVector3D) *QVector3D {
	ret := C.QVector3D_Normal2(v1.cPointer(), v2.cPointer(), v3.cPointer())
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQVector3D(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QVector3D) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QVector3D) Project(modelView *QMatrix4x4, projection *QMatrix4x4, viewport *QRect) *QVector3D {
	ret := C.QVector3D_Project(this.h, modelView.cPointer(), projection.cPointer(), viewport.cPointer())
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQVector3D(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QVector3D) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QVector3D) Unproject(modelView *QMatrix4x4, projection *QMatrix4x4, viewport *QRect) *QVector3D {
	ret := C.QVector3D_Unproject(this.h, modelView.cPointer(), projection.cPointer(), viewport.cPointer())
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQVector3D(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QVector3D) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QVector3D) DistanceToPoint(point *QVector3D) float32 {
	ret := C.QVector3D_DistanceToPoint(this.h, point.cPointer())
	return (float32)(ret)
}

func (this *QVector3D) DistanceToPlane(plane *QVector3D, normal *QVector3D) float32 {
	ret := C.QVector3D_DistanceToPlane(this.h, plane.cPointer(), normal.cPointer())
	return (float32)(ret)
}

func (this *QVector3D) DistanceToPlane2(plane1 *QVector3D, plane2 *QVector3D, plane3 *QVector3D) float32 {
	ret := C.QVector3D_DistanceToPlane2(this.h, plane1.cPointer(), plane2.cPointer(), plane3.cPointer())
	return (float32)(ret)
}

func (this *QVector3D) DistanceToLine(point *QVector3D, direction *QVector3D) float32 {
	ret := C.QVector3D_DistanceToLine(this.h, point.cPointer(), direction.cPointer())
	return (float32)(ret)
}

func (this *QVector3D) ToVector2D() *QVector2D {
	ret := C.QVector3D_ToVector2D(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQVector2D(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QVector2D) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QVector3D) ToVector4D() *QVector4D {
	ret := C.QVector3D_ToVector4D(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQVector4D(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QVector4D) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QVector3D) ToPoint() *QPoint {
	ret := C.QVector3D_ToPoint(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPoint(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPoint) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QVector3D) ToPointF() *QPointF {
	ret := C.QVector3D_ToPointF(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPointF(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPointF) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QVector3D) Delete() {
	C.QVector3D_Delete(this.h)
}