#include <QMetaObject>
#include <QPaintDeviceWindow>
#include <QRect>
#include <QRegion>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qpaintdevicewindow.h>
#include "gen_qpaintdevicewindow.h"
#include "_cgo_export.h"

QMetaObject* QPaintDeviceWindow_MetaObject(const QPaintDeviceWindow* self) {
	return (QMetaObject*) self->metaObject();
}

void* QPaintDeviceWindow_Metacast(QPaintDeviceWindow* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QPaintDeviceWindow_Tr(const char* s) {
	QString _ret = QPaintDeviceWindow::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QPaintDeviceWindow_Update(QPaintDeviceWindow* self, QRect* rect) {
	self->update(*rect);
}

void QPaintDeviceWindow_UpdateWithRegion(QPaintDeviceWindow* self, QRegion* region) {
	self->update(*region);
}

void QPaintDeviceWindow_Update2(QPaintDeviceWindow* self) {
	self->update();
}

struct miqt_string QPaintDeviceWindow_Tr2(const char* s, const char* c) {
	QString _ret = QPaintDeviceWindow::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QPaintDeviceWindow_Tr3(const char* s, const char* c, int n) {
	QString _ret = QPaintDeviceWindow::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QPaintDeviceWindow_Delete(QPaintDeviceWindow* self) {
	delete self;
}
