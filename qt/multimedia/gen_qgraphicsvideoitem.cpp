#include <QAbstractVideoSurface>
#include <QEvent>
#include <QGraphicsItem>
#include <QGraphicsObject>
#include <QGraphicsVideoItem>
#include <QMediaBindableInterface>
#include <QMediaObject>
#include <QMetaObject>
#include <QObject>
#include <QPainter>
#include <QPointF>
#include <QRectF>
#include <QSizeF>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyleOptionGraphicsItem>
#include <QTimerEvent>
#include <QVariant>
#include <QWidget>
#include <qgraphicsvideoitem.h>
#include "gen_qgraphicsvideoitem.h"

#ifndef _Bool
#define _Bool bool
#endif
#include "_cgo_export.h"

class MiqtVirtualQGraphicsVideoItem : public virtual QGraphicsVideoItem {
public:

	MiqtVirtualQGraphicsVideoItem(): QGraphicsVideoItem() {};
	MiqtVirtualQGraphicsVideoItem(QGraphicsItem* parent): QGraphicsVideoItem(parent) {};

	virtual ~MiqtVirtualQGraphicsVideoItem() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MediaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual QMediaObject* mediaObject() const override {
		if (handle__MediaObject == 0) {
			return QGraphicsVideoItem::mediaObject();
		}
		

		QMediaObject* callback_return_value = miqt_exec_callback_QGraphicsVideoItem_MediaObject(const_cast<MiqtVirtualQGraphicsVideoItem*>(this), handle__MediaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMediaObject* virtualbase_MediaObject() const {

		return QGraphicsVideoItem::mediaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__BoundingRect = 0;

	// Subclass to allow providing a Go implementation
	virtual QRectF boundingRect() const override {
		if (handle__BoundingRect == 0) {
			return QGraphicsVideoItem::boundingRect();
		}
		

		QRectF* callback_return_value = miqt_exec_callback_QGraphicsVideoItem_BoundingRect(const_cast<MiqtVirtualQGraphicsVideoItem*>(this), handle__BoundingRect);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QRectF* virtualbase_BoundingRect() const {

		return new QRectF(QGraphicsVideoItem::boundingRect());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Paint = 0;

	// Subclass to allow providing a Go implementation
	virtual void paint(QPainter* painter, const QStyleOptionGraphicsItem* option, QWidget* widget) override {
		if (handle__Paint == 0) {
			QGraphicsVideoItem::paint(painter, option, widget);
			return;
		}
		
		QPainter* sigval1 = painter;
		QStyleOptionGraphicsItem* sigval2 = (QStyleOptionGraphicsItem*) option;
		QWidget* sigval3 = widget;

		miqt_exec_callback_QGraphicsVideoItem_Paint(this, handle__Paint, sigval1, sigval2, sigval3);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_Paint(QPainter* painter, QStyleOptionGraphicsItem* option, QWidget* widget) {

		QGraphicsVideoItem::paint(painter, option, widget);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__TimerEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void timerEvent(QTimerEvent* event) override {
		if (handle__TimerEvent == 0) {
			QGraphicsVideoItem::timerEvent(event);
			return;
		}
		
		QTimerEvent* sigval1 = event;

		miqt_exec_callback_QGraphicsVideoItem_TimerEvent(this, handle__TimerEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_TimerEvent(QTimerEvent* event) {

		QGraphicsVideoItem::timerEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ItemChange = 0;

	// Subclass to allow providing a Go implementation
	virtual QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant& value) override {
		if (handle__ItemChange == 0) {
			return QGraphicsVideoItem::itemChange(change, value);
		}
		
		QGraphicsItem::GraphicsItemChange change_ret = change;
		int sigval1 = static_cast<int>(change_ret);
		const QVariant& value_ret = value;
		// Cast returned reference into pointer
		QVariant* sigval2 = const_cast<QVariant*>(&value_ret);

		QVariant* callback_return_value = miqt_exec_callback_QGraphicsVideoItem_ItemChange(this, handle__ItemChange, sigval1, sigval2);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QVariant* virtualbase_ItemChange(int change, QVariant* value) {

		return new QVariant(QGraphicsVideoItem::itemChange(static_cast<QGraphicsItem::GraphicsItemChange>(change), *value));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetMediaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual bool setMediaObject(QMediaObject* object) override {
		if (handle__SetMediaObject == 0) {
			return QGraphicsVideoItem::setMediaObject(object);
		}
		
		QMediaObject* sigval1 = object;

		bool callback_return_value = miqt_exec_callback_QGraphicsVideoItem_SetMediaObject(this, handle__SetMediaObject, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_SetMediaObject(QMediaObject* object) {

		return QGraphicsVideoItem::setMediaObject(object);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Event = 0;

	// Subclass to allow providing a Go implementation
	virtual bool event(QEvent* ev) override {
		if (handle__Event == 0) {
			return QGraphicsVideoItem::event(ev);
		}
		
		QEvent* sigval1 = ev;

		bool callback_return_value = miqt_exec_callback_QGraphicsVideoItem_Event(this, handle__Event, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_Event(QEvent* ev) {

		return QGraphicsVideoItem::event(ev);

	}

};

QGraphicsVideoItem* QGraphicsVideoItem_new() {
	return new MiqtVirtualQGraphicsVideoItem();
}

QGraphicsVideoItem* QGraphicsVideoItem_new2(QGraphicsItem* parent) {
	return new MiqtVirtualQGraphicsVideoItem(parent);
}

void QGraphicsVideoItem_virtbase(QGraphicsVideoItem* src, QGraphicsObject** outptr_QGraphicsObject, QMediaBindableInterface** outptr_QMediaBindableInterface) {
	*outptr_QGraphicsObject = static_cast<QGraphicsObject*>(src);
	*outptr_QMediaBindableInterface = static_cast<QMediaBindableInterface*>(src);
}

QMetaObject* QGraphicsVideoItem_MetaObject(const QGraphicsVideoItem* self) {
	return (QMetaObject*) self->metaObject();
}

void* QGraphicsVideoItem_Metacast(QGraphicsVideoItem* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QGraphicsVideoItem_Tr(const char* s) {
	QString _ret = QGraphicsVideoItem::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QGraphicsVideoItem_TrUtf8(const char* s) {
	QString _ret = QGraphicsVideoItem::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QMediaObject* QGraphicsVideoItem_MediaObject(const QGraphicsVideoItem* self) {
	return self->mediaObject();
}

QAbstractVideoSurface* QGraphicsVideoItem_VideoSurface(const QGraphicsVideoItem* self) {
	return self->videoSurface();
}

int QGraphicsVideoItem_AspectRatioMode(const QGraphicsVideoItem* self) {
	Qt::AspectRatioMode _ret = self->aspectRatioMode();
	return static_cast<int>(_ret);
}

void QGraphicsVideoItem_SetAspectRatioMode(QGraphicsVideoItem* self, int mode) {
	self->setAspectRatioMode(static_cast<Qt::AspectRatioMode>(mode));
}

QPointF* QGraphicsVideoItem_Offset(const QGraphicsVideoItem* self) {
	return new QPointF(self->offset());
}

void QGraphicsVideoItem_SetOffset(QGraphicsVideoItem* self, QPointF* offset) {
	self->setOffset(*offset);
}

QSizeF* QGraphicsVideoItem_Size(const QGraphicsVideoItem* self) {
	return new QSizeF(self->size());
}

void QGraphicsVideoItem_SetSize(QGraphicsVideoItem* self, QSizeF* size) {
	self->setSize(*size);
}

QSizeF* QGraphicsVideoItem_NativeSize(const QGraphicsVideoItem* self) {
	return new QSizeF(self->nativeSize());
}

QRectF* QGraphicsVideoItem_BoundingRect(const QGraphicsVideoItem* self) {
	return new QRectF(self->boundingRect());
}

void QGraphicsVideoItem_Paint(QGraphicsVideoItem* self, QPainter* painter, QStyleOptionGraphicsItem* option, QWidget* widget) {
	self->paint(painter, option, widget);
}

void QGraphicsVideoItem_NativeSizeChanged(QGraphicsVideoItem* self, QSizeF* size) {
	self->nativeSizeChanged(*size);
}

void QGraphicsVideoItem_connect_NativeSizeChanged(QGraphicsVideoItem* self, intptr_t slot) {
	MiqtVirtualQGraphicsVideoItem::connect(self, static_cast<void (QGraphicsVideoItem::*)(const QSizeF&)>(&QGraphicsVideoItem::nativeSizeChanged), self, [=](const QSizeF& size) {
		const QSizeF& size_ret = size;
		// Cast returned reference into pointer
		QSizeF* sigval1 = const_cast<QSizeF*>(&size_ret);
		miqt_exec_callback_QGraphicsVideoItem_NativeSizeChanged(slot, sigval1);
	});
}

struct miqt_string QGraphicsVideoItem_Tr2(const char* s, const char* c) {
	QString _ret = QGraphicsVideoItem::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QGraphicsVideoItem_Tr3(const char* s, const char* c, int n) {
	QString _ret = QGraphicsVideoItem::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QGraphicsVideoItem_TrUtf82(const char* s, const char* c) {
	QString _ret = QGraphicsVideoItem::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QGraphicsVideoItem_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QGraphicsVideoItem::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QGraphicsVideoItem_override_virtual_MediaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsVideoItem*>( (QGraphicsVideoItem*)(self) )->handle__MediaObject = slot;
}

QMediaObject* QGraphicsVideoItem_virtualbase_MediaObject(const void* self) {
	return ( (const MiqtVirtualQGraphicsVideoItem*)(self) )->virtualbase_MediaObject();
}

void QGraphicsVideoItem_override_virtual_BoundingRect(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsVideoItem*>( (QGraphicsVideoItem*)(self) )->handle__BoundingRect = slot;
}

QRectF* QGraphicsVideoItem_virtualbase_BoundingRect(const void* self) {
	return ( (const MiqtVirtualQGraphicsVideoItem*)(self) )->virtualbase_BoundingRect();
}

void QGraphicsVideoItem_override_virtual_Paint(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsVideoItem*>( (QGraphicsVideoItem*)(self) )->handle__Paint = slot;
}

void QGraphicsVideoItem_virtualbase_Paint(void* self, QPainter* painter, QStyleOptionGraphicsItem* option, QWidget* widget) {
	( (MiqtVirtualQGraphicsVideoItem*)(self) )->virtualbase_Paint(painter, option, widget);
}

void QGraphicsVideoItem_override_virtual_TimerEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsVideoItem*>( (QGraphicsVideoItem*)(self) )->handle__TimerEvent = slot;
}

void QGraphicsVideoItem_virtualbase_TimerEvent(void* self, QTimerEvent* event) {
	( (MiqtVirtualQGraphicsVideoItem*)(self) )->virtualbase_TimerEvent(event);
}

void QGraphicsVideoItem_override_virtual_ItemChange(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsVideoItem*>( (QGraphicsVideoItem*)(self) )->handle__ItemChange = slot;
}

QVariant* QGraphicsVideoItem_virtualbase_ItemChange(void* self, int change, QVariant* value) {
	return ( (MiqtVirtualQGraphicsVideoItem*)(self) )->virtualbase_ItemChange(change, value);
}

void QGraphicsVideoItem_override_virtual_SetMediaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsVideoItem*>( (QGraphicsVideoItem*)(self) )->handle__SetMediaObject = slot;
}

bool QGraphicsVideoItem_virtualbase_SetMediaObject(void* self, QMediaObject* object) {
	return ( (MiqtVirtualQGraphicsVideoItem*)(self) )->virtualbase_SetMediaObject(object);
}

void QGraphicsVideoItem_override_virtual_Event(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsVideoItem*>( (QGraphicsVideoItem*)(self) )->handle__Event = slot;
}

bool QGraphicsVideoItem_virtualbase_Event(void* self, QEvent* ev) {
	return ( (MiqtVirtualQGraphicsVideoItem*)(self) )->virtualbase_Event(ev);
}

void QGraphicsVideoItem_Delete(QGraphicsVideoItem* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQGraphicsVideoItem*>( self );
	} else {
		delete self;
	}
}

