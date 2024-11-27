package webengine

/*

#include "gen_qwebengineview.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QWebEngineView struct {
	h          *C.QWebEngineView
	isSubclass bool
	*qt.QWidget
}

func (this *QWebEngineView) cPointer() *C.QWebEngineView {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QWebEngineView) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQWebEngineView constructs the type using only CGO pointers.
func newQWebEngineView(h *C.QWebEngineView, h_QWidget *C.QWidget, h_QObject *C.QObject, h_QPaintDevice *C.QPaintDevice) *QWebEngineView {
	if h == nil {
		return nil
	}
	return &QWebEngineView{h: h,
		QWidget: qt.UnsafeNewQWidget(unsafe.Pointer(h_QWidget), unsafe.Pointer(h_QObject), unsafe.Pointer(h_QPaintDevice))}
}

// UnsafeNewQWebEngineView constructs the type using only unsafe pointers.
func UnsafeNewQWebEngineView(h unsafe.Pointer, h_QWidget unsafe.Pointer, h_QObject unsafe.Pointer, h_QPaintDevice unsafe.Pointer) *QWebEngineView {
	if h == nil {
		return nil
	}

	return &QWebEngineView{h: (*C.QWebEngineView)(h),
		QWidget: qt.UnsafeNewQWidget(h_QWidget, h_QObject, h_QPaintDevice)}
}

// NewQWebEngineView constructs a new QWebEngineView object.
func NewQWebEngineView(parent *qt.QWidget) *QWebEngineView {
	var outptr_QWebEngineView *C.QWebEngineView = nil
	var outptr_QWidget *C.QWidget = nil
	var outptr_QObject *C.QObject = nil
	var outptr_QPaintDevice *C.QPaintDevice = nil

	C.QWebEngineView_new((*C.QWidget)(parent.UnsafePointer()), &outptr_QWebEngineView, &outptr_QWidget, &outptr_QObject, &outptr_QPaintDevice)
	ret := newQWebEngineView(outptr_QWebEngineView, outptr_QWidget, outptr_QObject, outptr_QPaintDevice)
	ret.isSubclass = true
	return ret
}

// NewQWebEngineView2 constructs a new QWebEngineView object.
func NewQWebEngineView2() *QWebEngineView {
	var outptr_QWebEngineView *C.QWebEngineView = nil
	var outptr_QWidget *C.QWidget = nil
	var outptr_QObject *C.QObject = nil
	var outptr_QPaintDevice *C.QPaintDevice = nil

	C.QWebEngineView_new2(&outptr_QWebEngineView, &outptr_QWidget, &outptr_QObject, &outptr_QPaintDevice)
	ret := newQWebEngineView(outptr_QWebEngineView, outptr_QWidget, outptr_QObject, outptr_QPaintDevice)
	ret.isSubclass = true
	return ret
}

func (this *QWebEngineView) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(C.QWebEngineView_MetaObject(this.h)))
}

func (this *QWebEngineView) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QWebEngineView_Metacast(this.h, param1_Cstring))
}

func QWebEngineView_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QWebEngineView_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWebEngineView_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QWebEngineView_TrUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWebEngineView) Page() *QWebEnginePage {
	return UnsafeNewQWebEnginePage(unsafe.Pointer(C.QWebEngineView_Page(this.h)), nil)
}

func (this *QWebEngineView) SetPage(page *QWebEnginePage) {
	C.QWebEngineView_SetPage(this.h, page.cPointer())
}

func (this *QWebEngineView) Load(url *qt.QUrl) {
	C.QWebEngineView_Load(this.h, (*C.QUrl)(url.UnsafePointer()))
}

func (this *QWebEngineView) LoadWithRequest(request *QWebEngineHttpRequest) {
	C.QWebEngineView_LoadWithRequest(this.h, request.cPointer())
}

func (this *QWebEngineView) SetHtml(html string) {
	html_ms := C.struct_miqt_string{}
	html_ms.data = C.CString(html)
	html_ms.len = C.size_t(len(html))
	defer C.free(unsafe.Pointer(html_ms.data))
	C.QWebEngineView_SetHtml(this.h, html_ms)
}

func (this *QWebEngineView) SetContent(data []byte) {
	data_alias := C.struct_miqt_string{}
	data_alias.data = (*C.char)(unsafe.Pointer(&data[0]))
	data_alias.len = C.size_t(len(data))
	C.QWebEngineView_SetContent(this.h, data_alias)
}

func (this *QWebEngineView) History() *QWebEngineHistory {
	return UnsafeNewQWebEngineHistory(unsafe.Pointer(C.QWebEngineView_History(this.h)))
}

func (this *QWebEngineView) Title() string {
	var _ms C.struct_miqt_string = C.QWebEngineView_Title(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWebEngineView) SetUrl(url *qt.QUrl) {
	C.QWebEngineView_SetUrl(this.h, (*C.QUrl)(url.UnsafePointer()))
}

func (this *QWebEngineView) Url() *qt.QUrl {
	_ret := C.QWebEngineView_Url(this.h)
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(_ret))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWebEngineView) IconUrl() *qt.QUrl {
	_ret := C.QWebEngineView_IconUrl(this.h)
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(_ret))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWebEngineView) Icon() *qt.QIcon {
	_ret := C.QWebEngineView_Icon(this.h)
	_goptr := qt.UnsafeNewQIcon(unsafe.Pointer(_ret))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWebEngineView) HasSelection() bool {
	return (bool)(C.QWebEngineView_HasSelection(this.h))
}

func (this *QWebEngineView) SelectedText() string {
	var _ms C.struct_miqt_string = C.QWebEngineView_SelectedText(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWebEngineView) PageAction(action QWebEnginePage__WebAction) *qt.QAction {
	return qt.UnsafeNewQAction(unsafe.Pointer(C.QWebEngineView_PageAction(this.h, (C.int)(action))), nil)
}

func (this *QWebEngineView) TriggerPageAction(action QWebEnginePage__WebAction) {
	C.QWebEngineView_TriggerPageAction(this.h, (C.int)(action))
}

func (this *QWebEngineView) ZoomFactor() float64 {
	return (float64)(C.QWebEngineView_ZoomFactor(this.h))
}

func (this *QWebEngineView) SetZoomFactor(factor float64) {
	C.QWebEngineView_SetZoomFactor(this.h, (C.double)(factor))
}

func (this *QWebEngineView) FindText(subString string) {
	subString_ms := C.struct_miqt_string{}
	subString_ms.data = C.CString(subString)
	subString_ms.len = C.size_t(len(subString))
	defer C.free(unsafe.Pointer(subString_ms.data))
	C.QWebEngineView_FindText(this.h, subString_ms)
}

func (this *QWebEngineView) SizeHint() *qt.QSize {
	_ret := C.QWebEngineView_SizeHint(this.h)
	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(_ret))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWebEngineView) Settings() *QWebEngineSettings {
	return UnsafeNewQWebEngineSettings(unsafe.Pointer(C.QWebEngineView_Settings(this.h)))
}

func (this *QWebEngineView) Stop() {
	C.QWebEngineView_Stop(this.h)
}

func (this *QWebEngineView) Back() {
	C.QWebEngineView_Back(this.h)
}

func (this *QWebEngineView) Forward() {
	C.QWebEngineView_Forward(this.h)
}

func (this *QWebEngineView) Reload() {
	C.QWebEngineView_Reload(this.h)
}

func (this *QWebEngineView) LoadStarted() {
	C.QWebEngineView_LoadStarted(this.h)
}
func (this *QWebEngineView) OnLoadStarted(slot func()) {
	C.QWebEngineView_connect_LoadStarted(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_LoadStarted
func miqt_exec_callback_QWebEngineView_LoadStarted(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QWebEngineView) LoadProgress(progress int) {
	C.QWebEngineView_LoadProgress(this.h, (C.int)(progress))
}
func (this *QWebEngineView) OnLoadProgress(slot func(progress int)) {
	C.QWebEngineView_connect_LoadProgress(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_LoadProgress
func miqt_exec_callback_QWebEngineView_LoadProgress(cb C.intptr_t, progress C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(progress int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(progress)

	gofunc(slotval1)
}

func (this *QWebEngineView) LoadFinished(param1 bool) {
	C.QWebEngineView_LoadFinished(this.h, (C.bool)(param1))
}
func (this *QWebEngineView) OnLoadFinished(slot func(param1 bool)) {
	C.QWebEngineView_connect_LoadFinished(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_LoadFinished
func miqt_exec_callback_QWebEngineView_LoadFinished(cb C.intptr_t, param1 C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(param1)

	gofunc(slotval1)
}

func (this *QWebEngineView) TitleChanged(title string) {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	C.QWebEngineView_TitleChanged(this.h, title_ms)
}
func (this *QWebEngineView) OnTitleChanged(slot func(title string)) {
	C.QWebEngineView_connect_TitleChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_TitleChanged
func miqt_exec_callback_QWebEngineView_TitleChanged(cb C.intptr_t, title C.struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(title string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var title_ms C.struct_miqt_string = title
	title_ret := C.GoStringN(title_ms.data, C.int(int64(title_ms.len)))
	C.free(unsafe.Pointer(title_ms.data))
	slotval1 := title_ret

	gofunc(slotval1)
}

func (this *QWebEngineView) SelectionChanged() {
	C.QWebEngineView_SelectionChanged(this.h)
}
func (this *QWebEngineView) OnSelectionChanged(slot func()) {
	C.QWebEngineView_connect_SelectionChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_SelectionChanged
func miqt_exec_callback_QWebEngineView_SelectionChanged(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QWebEngineView) UrlChanged(param1 *qt.QUrl) {
	C.QWebEngineView_UrlChanged(this.h, (*C.QUrl)(param1.UnsafePointer()))
}
func (this *QWebEngineView) OnUrlChanged(slot func(param1 *qt.QUrl)) {
	C.QWebEngineView_connect_UrlChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_UrlChanged
func miqt_exec_callback_QWebEngineView_UrlChanged(cb C.intptr_t, param1 *C.QUrl) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *qt.QUrl))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQUrl(unsafe.Pointer(param1))

	gofunc(slotval1)
}

func (this *QWebEngineView) IconUrlChanged(param1 *qt.QUrl) {
	C.QWebEngineView_IconUrlChanged(this.h, (*C.QUrl)(param1.UnsafePointer()))
}
func (this *QWebEngineView) OnIconUrlChanged(slot func(param1 *qt.QUrl)) {
	C.QWebEngineView_connect_IconUrlChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_IconUrlChanged
func miqt_exec_callback_QWebEngineView_IconUrlChanged(cb C.intptr_t, param1 *C.QUrl) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *qt.QUrl))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQUrl(unsafe.Pointer(param1))

	gofunc(slotval1)
}

func (this *QWebEngineView) IconChanged(param1 *qt.QIcon) {
	C.QWebEngineView_IconChanged(this.h, (*C.QIcon)(param1.UnsafePointer()))
}
func (this *QWebEngineView) OnIconChanged(slot func(param1 *qt.QIcon)) {
	C.QWebEngineView_connect_IconChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_IconChanged
func miqt_exec_callback_QWebEngineView_IconChanged(cb C.intptr_t, param1 *C.QIcon) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *qt.QIcon))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQIcon(unsafe.Pointer(param1))

	gofunc(slotval1)
}

func (this *QWebEngineView) RenderProcessTerminated(terminationStatus QWebEnginePage__RenderProcessTerminationStatus, exitCode int) {
	C.QWebEngineView_RenderProcessTerminated(this.h, (C.int)(terminationStatus), (C.int)(exitCode))
}
func (this *QWebEngineView) OnRenderProcessTerminated(slot func(terminationStatus QWebEnginePage__RenderProcessTerminationStatus, exitCode int)) {
	C.QWebEngineView_connect_RenderProcessTerminated(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_RenderProcessTerminated
func miqt_exec_callback_QWebEngineView_RenderProcessTerminated(cb C.intptr_t, terminationStatus C.int, exitCode C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(terminationStatus QWebEnginePage__RenderProcessTerminationStatus, exitCode int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QWebEnginePage__RenderProcessTerminationStatus)(terminationStatus)

	slotval2 := (int)(exitCode)

	gofunc(slotval1, slotval2)
}

func QWebEngineView_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QWebEngineView_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWebEngineView_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QWebEngineView_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWebEngineView_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QWebEngineView_TrUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWebEngineView_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QWebEngineView_TrUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWebEngineView) SetHtml2(html string, baseUrl *qt.QUrl) {
	html_ms := C.struct_miqt_string{}
	html_ms.data = C.CString(html)
	html_ms.len = C.size_t(len(html))
	defer C.free(unsafe.Pointer(html_ms.data))
	C.QWebEngineView_SetHtml2(this.h, html_ms, (*C.QUrl)(baseUrl.UnsafePointer()))
}

func (this *QWebEngineView) SetContent2(data []byte, mimeType string) {
	data_alias := C.struct_miqt_string{}
	data_alias.data = (*C.char)(unsafe.Pointer(&data[0]))
	data_alias.len = C.size_t(len(data))
	mimeType_ms := C.struct_miqt_string{}
	mimeType_ms.data = C.CString(mimeType)
	mimeType_ms.len = C.size_t(len(mimeType))
	defer C.free(unsafe.Pointer(mimeType_ms.data))
	C.QWebEngineView_SetContent2(this.h, data_alias, mimeType_ms)
}

func (this *QWebEngineView) SetContent3(data []byte, mimeType string, baseUrl *qt.QUrl) {
	data_alias := C.struct_miqt_string{}
	data_alias.data = (*C.char)(unsafe.Pointer(&data[0]))
	data_alias.len = C.size_t(len(data))
	mimeType_ms := C.struct_miqt_string{}
	mimeType_ms.data = C.CString(mimeType)
	mimeType_ms.len = C.size_t(len(mimeType))
	defer C.free(unsafe.Pointer(mimeType_ms.data))
	C.QWebEngineView_SetContent3(this.h, data_alias, mimeType_ms, (*C.QUrl)(baseUrl.UnsafePointer()))
}

func (this *QWebEngineView) TriggerPageAction2(action QWebEnginePage__WebAction, checked bool) {
	C.QWebEngineView_TriggerPageAction2(this.h, (C.int)(action), (C.bool)(checked))
}

func (this *QWebEngineView) FindText2(subString string, options QWebEnginePage__FindFlag) {
	subString_ms := C.struct_miqt_string{}
	subString_ms.data = C.CString(subString)
	subString_ms.len = C.size_t(len(subString))
	defer C.free(unsafe.Pointer(subString_ms.data))
	C.QWebEngineView_FindText2(this.h, subString_ms, (C.int)(options))
}

func (this *QWebEngineView) callVirtualBase_SizeHint() *qt.QSize {

	_ret := C.QWebEngineView_virtualbase_SizeHint(unsafe.Pointer(this.h))
	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(_ret))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QWebEngineView) OnSizeHint(slot func(super func() *qt.QSize) *qt.QSize) {
	C.QWebEngineView_override_virtual_SizeHint(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_SizeHint
func miqt_exec_callback_QWebEngineView_SizeHint(self *C.QWebEngineView, cb C.intptr_t) *C.QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QSize) *qt.QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_SizeHint)

	return (*C.QSize)(virtualReturn.UnsafePointer())

}

func (this *QWebEngineView) callVirtualBase_CreateWindow(typeVal QWebEnginePage__WebWindowType) *QWebEngineView {

	return UnsafeNewQWebEngineView(unsafe.Pointer(C.QWebEngineView_virtualbase_CreateWindow(unsafe.Pointer(this.h), (C.int)(typeVal))), nil, nil, nil)
}
func (this *QWebEngineView) OnCreateWindow(slot func(super func(typeVal QWebEnginePage__WebWindowType) *QWebEngineView, typeVal QWebEnginePage__WebWindowType) *QWebEngineView) {
	C.QWebEngineView_override_virtual_CreateWindow(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_CreateWindow
func miqt_exec_callback_QWebEngineView_CreateWindow(self *C.QWebEngineView, cb C.intptr_t, typeVal C.int) *C.QWebEngineView {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(typeVal QWebEnginePage__WebWindowType) *QWebEngineView, typeVal QWebEnginePage__WebWindowType) *QWebEngineView)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QWebEnginePage__WebWindowType)(typeVal)

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_CreateWindow, slotval1)

	return virtualReturn.cPointer()

}

func (this *QWebEngineView) callVirtualBase_ContextMenuEvent(param1 *qt.QContextMenuEvent) {

	C.QWebEngineView_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), (*C.QContextMenuEvent)(param1.UnsafePointer()))

}
func (this *QWebEngineView) OnContextMenuEvent(slot func(super func(param1 *qt.QContextMenuEvent), param1 *qt.QContextMenuEvent)) {
	C.QWebEngineView_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_ContextMenuEvent
func miqt_exec_callback_QWebEngineView_ContextMenuEvent(self *C.QWebEngineView, cb C.intptr_t, param1 *C.QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QContextMenuEvent), param1 *qt.QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQContextMenuEvent(unsafe.Pointer(param1), nil, nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_Event(param1 *qt.QEvent) bool {

	return (bool)(C.QWebEngineView_virtualbase_Event(unsafe.Pointer(this.h), (*C.QEvent)(param1.UnsafePointer())))

}
func (this *QWebEngineView) OnEvent(slot func(super func(param1 *qt.QEvent) bool, param1 *qt.QEvent) bool) {
	C.QWebEngineView_override_virtual_Event(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_Event
func miqt_exec_callback_QWebEngineView_Event(self *C.QWebEngineView, cb C.intptr_t, param1 *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QEvent) bool, param1 *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(param1))

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_Event, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QWebEngineView) callVirtualBase_ShowEvent(param1 *qt.QShowEvent) {

	C.QWebEngineView_virtualbase_ShowEvent(unsafe.Pointer(this.h), (*C.QShowEvent)(param1.UnsafePointer()))

}
func (this *QWebEngineView) OnShowEvent(slot func(super func(param1 *qt.QShowEvent), param1 *qt.QShowEvent)) {
	C.QWebEngineView_override_virtual_ShowEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_ShowEvent
func miqt_exec_callback_QWebEngineView_ShowEvent(self *C.QWebEngineView, cb C.intptr_t, param1 *C.QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QShowEvent), param1 *qt.QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQShowEvent(unsafe.Pointer(param1), nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_HideEvent(param1 *qt.QHideEvent) {

	C.QWebEngineView_virtualbase_HideEvent(unsafe.Pointer(this.h), (*C.QHideEvent)(param1.UnsafePointer()))

}
func (this *QWebEngineView) OnHideEvent(slot func(super func(param1 *qt.QHideEvent), param1 *qt.QHideEvent)) {
	C.QWebEngineView_override_virtual_HideEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_HideEvent
func miqt_exec_callback_QWebEngineView_HideEvent(self *C.QWebEngineView, cb C.intptr_t, param1 *C.QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QHideEvent), param1 *qt.QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQHideEvent(unsafe.Pointer(param1), nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_CloseEvent(param1 *qt.QCloseEvent) {

	C.QWebEngineView_virtualbase_CloseEvent(unsafe.Pointer(this.h), (*C.QCloseEvent)(param1.UnsafePointer()))

}
func (this *QWebEngineView) OnCloseEvent(slot func(super func(param1 *qt.QCloseEvent), param1 *qt.QCloseEvent)) {
	C.QWebEngineView_override_virtual_CloseEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_CloseEvent
func miqt_exec_callback_QWebEngineView_CloseEvent(self *C.QWebEngineView, cb C.intptr_t, param1 *C.QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QCloseEvent), param1 *qt.QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQCloseEvent(unsafe.Pointer(param1), nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_DragEnterEvent(e *qt.QDragEnterEvent) {

	C.QWebEngineView_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), (*C.QDragEnterEvent)(e.UnsafePointer()))

}
func (this *QWebEngineView) OnDragEnterEvent(slot func(super func(e *qt.QDragEnterEvent), e *qt.QDragEnterEvent)) {
	C.QWebEngineView_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_DragEnterEvent
func miqt_exec_callback_QWebEngineView_DragEnterEvent(self *C.QWebEngineView, cb C.intptr_t, e *C.QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *qt.QDragEnterEvent), e *qt.QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQDragEnterEvent(unsafe.Pointer(e), nil, nil, nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_DragLeaveEvent(e *qt.QDragLeaveEvent) {

	C.QWebEngineView_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), (*C.QDragLeaveEvent)(e.UnsafePointer()))

}
func (this *QWebEngineView) OnDragLeaveEvent(slot func(super func(e *qt.QDragLeaveEvent), e *qt.QDragLeaveEvent)) {
	C.QWebEngineView_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_DragLeaveEvent
func miqt_exec_callback_QWebEngineView_DragLeaveEvent(self *C.QWebEngineView, cb C.intptr_t, e *C.QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *qt.QDragLeaveEvent), e *qt.QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQDragLeaveEvent(unsafe.Pointer(e), nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_DragMoveEvent(e *qt.QDragMoveEvent) {

	C.QWebEngineView_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), (*C.QDragMoveEvent)(e.UnsafePointer()))

}
func (this *QWebEngineView) OnDragMoveEvent(slot func(super func(e *qt.QDragMoveEvent), e *qt.QDragMoveEvent)) {
	C.QWebEngineView_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_DragMoveEvent
func miqt_exec_callback_QWebEngineView_DragMoveEvent(self *C.QWebEngineView, cb C.intptr_t, e *C.QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *qt.QDragMoveEvent), e *qt.QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQDragMoveEvent(unsafe.Pointer(e), nil, nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_DropEvent(e *qt.QDropEvent) {

	C.QWebEngineView_virtualbase_DropEvent(unsafe.Pointer(this.h), (*C.QDropEvent)(e.UnsafePointer()))

}
func (this *QWebEngineView) OnDropEvent(slot func(super func(e *qt.QDropEvent), e *qt.QDropEvent)) {
	C.QWebEngineView_override_virtual_DropEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_DropEvent
func miqt_exec_callback_QWebEngineView_DropEvent(self *C.QWebEngineView, cb C.intptr_t, e *C.QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *qt.QDropEvent), e *qt.QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQDropEvent(unsafe.Pointer(e), nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_DevType() int {

	return (int)(C.QWebEngineView_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QWebEngineView) OnDevType(slot func(super func() int) int) {
	C.QWebEngineView_override_virtual_DevType(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_DevType
func miqt_exec_callback_QWebEngineView_DevType(self *C.QWebEngineView, cb C.intptr_t) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_DevType)

	return (C.int)(virtualReturn)

}

func (this *QWebEngineView) callVirtualBase_SetVisible(visible bool) {

	C.QWebEngineView_virtualbase_SetVisible(unsafe.Pointer(this.h), (C.bool)(visible))

}
func (this *QWebEngineView) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	C.QWebEngineView_override_virtual_SetVisible(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_SetVisible
func miqt_exec_callback_QWebEngineView_SetVisible(self *C.QWebEngineView, cb C.intptr_t, visible C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QWebEngineView) callVirtualBase_MinimumSizeHint() *qt.QSize {

	_ret := C.QWebEngineView_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h))
	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(_ret))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QWebEngineView) OnMinimumSizeHint(slot func(super func() *qt.QSize) *qt.QSize) {
	C.QWebEngineView_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_MinimumSizeHint
func miqt_exec_callback_QWebEngineView_MinimumSizeHint(self *C.QWebEngineView, cb C.intptr_t) *C.QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QSize) *qt.QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_MinimumSizeHint)

	return (*C.QSize)(virtualReturn.UnsafePointer())

}

func (this *QWebEngineView) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(C.QWebEngineView_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (C.int)(param1)))

}
func (this *QWebEngineView) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	C.QWebEngineView_override_virtual_HeightForWidth(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_HeightForWidth
func miqt_exec_callback_QWebEngineView_HeightForWidth(self *C.QWebEngineView, cb C.intptr_t, param1 C.int) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (C.int)(virtualReturn)

}

func (this *QWebEngineView) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(C.QWebEngineView_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QWebEngineView) OnHasHeightForWidth(slot func(super func() bool) bool) {
	C.QWebEngineView_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_HasHeightForWidth
func miqt_exec_callback_QWebEngineView_HasHeightForWidth(self *C.QWebEngineView, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_HasHeightForWidth)

	return (C.bool)(virtualReturn)

}

func (this *QWebEngineView) callVirtualBase_PaintEngine() *qt.QPaintEngine {

	return qt.UnsafeNewQPaintEngine(unsafe.Pointer(C.QWebEngineView_virtualbase_PaintEngine(unsafe.Pointer(this.h))))
}
func (this *QWebEngineView) OnPaintEngine(slot func(super func() *qt.QPaintEngine) *qt.QPaintEngine) {
	C.QWebEngineView_override_virtual_PaintEngine(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_PaintEngine
func miqt_exec_callback_QWebEngineView_PaintEngine(self *C.QWebEngineView, cb C.intptr_t) *C.QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QPaintEngine) *qt.QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_PaintEngine)

	return (*C.QPaintEngine)(virtualReturn.UnsafePointer())

}

func (this *QWebEngineView) callVirtualBase_MousePressEvent(event *qt.QMouseEvent) {

	C.QWebEngineView_virtualbase_MousePressEvent(unsafe.Pointer(this.h), (*C.QMouseEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnMousePressEvent(slot func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent)) {
	C.QWebEngineView_override_virtual_MousePressEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_MousePressEvent
func miqt_exec_callback_QWebEngineView_MousePressEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMouseEvent(unsafe.Pointer(event), nil, nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_MouseReleaseEvent(event *qt.QMouseEvent) {

	C.QWebEngineView_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), (*C.QMouseEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnMouseReleaseEvent(slot func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent)) {
	C.QWebEngineView_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_MouseReleaseEvent
func miqt_exec_callback_QWebEngineView_MouseReleaseEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMouseEvent(unsafe.Pointer(event), nil, nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_MouseDoubleClickEvent(event *qt.QMouseEvent) {

	C.QWebEngineView_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), (*C.QMouseEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnMouseDoubleClickEvent(slot func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent)) {
	C.QWebEngineView_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_MouseDoubleClickEvent
func miqt_exec_callback_QWebEngineView_MouseDoubleClickEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMouseEvent(unsafe.Pointer(event), nil, nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_MouseMoveEvent(event *qt.QMouseEvent) {

	C.QWebEngineView_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), (*C.QMouseEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnMouseMoveEvent(slot func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent)) {
	C.QWebEngineView_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_MouseMoveEvent
func miqt_exec_callback_QWebEngineView_MouseMoveEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMouseEvent(unsafe.Pointer(event), nil, nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_WheelEvent(event *qt.QWheelEvent) {

	C.QWebEngineView_virtualbase_WheelEvent(unsafe.Pointer(this.h), (*C.QWheelEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnWheelEvent(slot func(super func(event *qt.QWheelEvent), event *qt.QWheelEvent)) {
	C.QWebEngineView_override_virtual_WheelEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_WheelEvent
func miqt_exec_callback_QWebEngineView_WheelEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QWheelEvent), event *qt.QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQWheelEvent(unsafe.Pointer(event), nil, nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_KeyPressEvent(event *qt.QKeyEvent) {

	C.QWebEngineView_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), (*C.QKeyEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnKeyPressEvent(slot func(super func(event *qt.QKeyEvent), event *qt.QKeyEvent)) {
	C.QWebEngineView_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_KeyPressEvent
func miqt_exec_callback_QWebEngineView_KeyPressEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QKeyEvent), event *qt.QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQKeyEvent(unsafe.Pointer(event), nil, nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_KeyReleaseEvent(event *qt.QKeyEvent) {

	C.QWebEngineView_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), (*C.QKeyEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnKeyReleaseEvent(slot func(super func(event *qt.QKeyEvent), event *qt.QKeyEvent)) {
	C.QWebEngineView_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_KeyReleaseEvent
func miqt_exec_callback_QWebEngineView_KeyReleaseEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QKeyEvent), event *qt.QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQKeyEvent(unsafe.Pointer(event), nil, nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_FocusInEvent(event *qt.QFocusEvent) {

	C.QWebEngineView_virtualbase_FocusInEvent(unsafe.Pointer(this.h), (*C.QFocusEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnFocusInEvent(slot func(super func(event *qt.QFocusEvent), event *qt.QFocusEvent)) {
	C.QWebEngineView_override_virtual_FocusInEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_FocusInEvent
func miqt_exec_callback_QWebEngineView_FocusInEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QFocusEvent), event *qt.QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQFocusEvent(unsafe.Pointer(event), nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_FocusOutEvent(event *qt.QFocusEvent) {

	C.QWebEngineView_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), (*C.QFocusEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnFocusOutEvent(slot func(super func(event *qt.QFocusEvent), event *qt.QFocusEvent)) {
	C.QWebEngineView_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_FocusOutEvent
func miqt_exec_callback_QWebEngineView_FocusOutEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QFocusEvent), event *qt.QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQFocusEvent(unsafe.Pointer(event), nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_EnterEvent(event *qt.QEvent) {

	C.QWebEngineView_virtualbase_EnterEvent(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnEnterEvent(slot func(super func(event *qt.QEvent), event *qt.QEvent)) {
	C.QWebEngineView_override_virtual_EnterEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_EnterEvent
func miqt_exec_callback_QWebEngineView_EnterEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent), event *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_LeaveEvent(event *qt.QEvent) {

	C.QWebEngineView_virtualbase_LeaveEvent(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnLeaveEvent(slot func(super func(event *qt.QEvent), event *qt.QEvent)) {
	C.QWebEngineView_override_virtual_LeaveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_LeaveEvent
func miqt_exec_callback_QWebEngineView_LeaveEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent), event *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_PaintEvent(event *qt.QPaintEvent) {

	C.QWebEngineView_virtualbase_PaintEvent(unsafe.Pointer(this.h), (*C.QPaintEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnPaintEvent(slot func(super func(event *qt.QPaintEvent), event *qt.QPaintEvent)) {
	C.QWebEngineView_override_virtual_PaintEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_PaintEvent
func miqt_exec_callback_QWebEngineView_PaintEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QPaintEvent), event *qt.QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQPaintEvent(unsafe.Pointer(event), nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_MoveEvent(event *qt.QMoveEvent) {

	C.QWebEngineView_virtualbase_MoveEvent(unsafe.Pointer(this.h), (*C.QMoveEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnMoveEvent(slot func(super func(event *qt.QMoveEvent), event *qt.QMoveEvent)) {
	C.QWebEngineView_override_virtual_MoveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_MoveEvent
func miqt_exec_callback_QWebEngineView_MoveEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QMoveEvent), event *qt.QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMoveEvent(unsafe.Pointer(event), nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_ResizeEvent(event *qt.QResizeEvent) {

	C.QWebEngineView_virtualbase_ResizeEvent(unsafe.Pointer(this.h), (*C.QResizeEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnResizeEvent(slot func(super func(event *qt.QResizeEvent), event *qt.QResizeEvent)) {
	C.QWebEngineView_override_virtual_ResizeEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_ResizeEvent
func miqt_exec_callback_QWebEngineView_ResizeEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QResizeEvent), event *qt.QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQResizeEvent(unsafe.Pointer(event), nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_TabletEvent(event *qt.QTabletEvent) {

	C.QWebEngineView_virtualbase_TabletEvent(unsafe.Pointer(this.h), (*C.QTabletEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnTabletEvent(slot func(super func(event *qt.QTabletEvent), event *qt.QTabletEvent)) {
	C.QWebEngineView_override_virtual_TabletEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_TabletEvent
func miqt_exec_callback_QWebEngineView_TabletEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QTabletEvent), event *qt.QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQTabletEvent(unsafe.Pointer(event), nil, nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_ActionEvent(event *qt.QActionEvent) {

	C.QWebEngineView_virtualbase_ActionEvent(unsafe.Pointer(this.h), (*C.QActionEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnActionEvent(slot func(super func(event *qt.QActionEvent), event *qt.QActionEvent)) {
	C.QWebEngineView_override_virtual_ActionEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_ActionEvent
func miqt_exec_callback_QWebEngineView_ActionEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QActionEvent), event *qt.QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQActionEvent(unsafe.Pointer(event), nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *int64) bool {
	eventType_alias := C.struct_miqt_string{}
	eventType_alias.data = (*C.char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = C.size_t(len(eventType))

	return (bool)(C.QWebEngineView_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*C.long)(unsafe.Pointer(result))))

}
func (this *QWebEngineView) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *int64) bool, eventType []byte, message unsafe.Pointer, result *int64) bool) {
	C.QWebEngineView_override_virtual_NativeEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_NativeEvent
func miqt_exec_callback_QWebEngineView_NativeEvent(self *C.QWebEngineView, cb C.intptr_t, eventType C.struct_miqt_string, message unsafe.Pointer, result *C.long) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(eventType []byte, message unsafe.Pointer, result *int64) bool, eventType []byte, message unsafe.Pointer, result *int64) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var eventType_bytearray C.struct_miqt_string = eventType
	eventType_ret := C.GoBytes(unsafe.Pointer(eventType_bytearray.data), C.int(int64(eventType_bytearray.len)))
	C.free(unsafe.Pointer(eventType_bytearray.data))
	slotval1 := eventType_ret
	slotval2 := (unsafe.Pointer)(message)

	slotval3 := (*int64)(unsafe.Pointer(result))

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (C.bool)(virtualReturn)

}

func (this *QWebEngineView) callVirtualBase_ChangeEvent(param1 *qt.QEvent) {

	C.QWebEngineView_virtualbase_ChangeEvent(unsafe.Pointer(this.h), (*C.QEvent)(param1.UnsafePointer()))

}
func (this *QWebEngineView) OnChangeEvent(slot func(super func(param1 *qt.QEvent), param1 *qt.QEvent)) {
	C.QWebEngineView_override_virtual_ChangeEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_ChangeEvent
func miqt_exec_callback_QWebEngineView_ChangeEvent(self *C.QWebEngineView, cb C.intptr_t, param1 *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QEvent), param1 *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(param1))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_Metric(param1 qt.QPaintDevice__PaintDeviceMetric) int {

	return (int)(C.QWebEngineView_virtualbase_Metric(unsafe.Pointer(this.h), (C.int)(param1)))

}
func (this *QWebEngineView) OnMetric(slot func(super func(param1 qt.QPaintDevice__PaintDeviceMetric) int, param1 qt.QPaintDevice__PaintDeviceMetric) int) {
	C.QWebEngineView_override_virtual_Metric(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_Metric
func miqt_exec_callback_QWebEngineView_Metric(self *C.QWebEngineView, cb C.intptr_t, param1 C.int) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 qt.QPaintDevice__PaintDeviceMetric) int, param1 qt.QPaintDevice__PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (qt.QPaintDevice__PaintDeviceMetric)(param1)

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_Metric, slotval1)

	return (C.int)(virtualReturn)

}

func (this *QWebEngineView) callVirtualBase_InitPainter(painter *qt.QPainter) {

	C.QWebEngineView_virtualbase_InitPainter(unsafe.Pointer(this.h), (*C.QPainter)(painter.UnsafePointer()))

}
func (this *QWebEngineView) OnInitPainter(slot func(super func(painter *qt.QPainter), painter *qt.QPainter)) {
	C.QWebEngineView_override_virtual_InitPainter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_InitPainter
func miqt_exec_callback_QWebEngineView_InitPainter(self *C.QWebEngineView, cb C.intptr_t, painter *C.QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *qt.QPainter), painter *qt.QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQPainter(unsafe.Pointer(painter))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QWebEngineView) callVirtualBase_Redirected(offset *qt.QPoint) *qt.QPaintDevice {

	return qt.UnsafeNewQPaintDevice(unsafe.Pointer(C.QWebEngineView_virtualbase_Redirected(unsafe.Pointer(this.h), (*C.QPoint)(offset.UnsafePointer()))))
}
func (this *QWebEngineView) OnRedirected(slot func(super func(offset *qt.QPoint) *qt.QPaintDevice, offset *qt.QPoint) *qt.QPaintDevice) {
	C.QWebEngineView_override_virtual_Redirected(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_Redirected
func miqt_exec_callback_QWebEngineView_Redirected(self *C.QWebEngineView, cb C.intptr_t, offset *C.QPoint) *C.QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *qt.QPoint) *qt.QPaintDevice, offset *qt.QPoint) *qt.QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQPoint(unsafe.Pointer(offset))

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_Redirected, slotval1)

	return (*C.QPaintDevice)(virtualReturn.UnsafePointer())

}

func (this *QWebEngineView) callVirtualBase_SharedPainter() *qt.QPainter {

	return qt.UnsafeNewQPainter(unsafe.Pointer(C.QWebEngineView_virtualbase_SharedPainter(unsafe.Pointer(this.h))))
}
func (this *QWebEngineView) OnSharedPainter(slot func(super func() *qt.QPainter) *qt.QPainter) {
	C.QWebEngineView_override_virtual_SharedPainter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_SharedPainter
func miqt_exec_callback_QWebEngineView_SharedPainter(self *C.QWebEngineView, cb C.intptr_t) *C.QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QPainter) *qt.QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_SharedPainter)

	return (*C.QPainter)(virtualReturn.UnsafePointer())

}

func (this *QWebEngineView) callVirtualBase_InputMethodEvent(param1 *qt.QInputMethodEvent) {

	C.QWebEngineView_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), (*C.QInputMethodEvent)(param1.UnsafePointer()))

}
func (this *QWebEngineView) OnInputMethodEvent(slot func(super func(param1 *qt.QInputMethodEvent), param1 *qt.QInputMethodEvent)) {
	C.QWebEngineView_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_InputMethodEvent
func miqt_exec_callback_QWebEngineView_InputMethodEvent(self *C.QWebEngineView, cb C.intptr_t, param1 *C.QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QInputMethodEvent), param1 *qt.QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQInputMethodEvent(unsafe.Pointer(param1), nil)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_InputMethodQuery(param1 qt.InputMethodQuery) *qt.QVariant {

	_ret := C.QWebEngineView_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (C.int)(param1))
	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(_ret))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QWebEngineView) OnInputMethodQuery(slot func(super func(param1 qt.InputMethodQuery) *qt.QVariant, param1 qt.InputMethodQuery) *qt.QVariant) {
	C.QWebEngineView_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_InputMethodQuery
func miqt_exec_callback_QWebEngineView_InputMethodQuery(self *C.QWebEngineView, cb C.intptr_t, param1 C.int) *C.QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 qt.InputMethodQuery) *qt.QVariant, param1 qt.InputMethodQuery) *qt.QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (qt.InputMethodQuery)(param1)

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return (*C.QVariant)(virtualReturn.UnsafePointer())

}

func (this *QWebEngineView) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(C.QWebEngineView_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (C.bool)(next)))

}
func (this *QWebEngineView) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	C.QWebEngineView_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_FocusNextPrevChild
func miqt_exec_callback_QWebEngineView_FocusNextPrevChild(self *C.QWebEngineView, cb C.intptr_t, next C.bool) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (C.bool)(virtualReturn)

}

// Delete this object from C++ memory.
func (this *QWebEngineView) Delete() {
	C.QWebEngineView_Delete(this.h, C.bool(this.isSubclass))
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QWebEngineView) GoGC() {
	runtime.SetFinalizer(this, func(this *QWebEngineView) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}