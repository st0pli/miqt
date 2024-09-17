#include <QList>
#include <QMetaObject>
#include <QObject>
#include <QPointF>
#include <QRectF>
#include <QScroller>
#include <QScrollerProperties>
#include <QString>
#include <QByteArray>
#include <cstring>
#include "qscroller.h"
#include "gen_qscroller.h"
#include "_cgo_export.h"

QMetaObject* QScroller_MetaObject(const QScroller* self) {
	return (QMetaObject*) self->metaObject();
}

struct miqt_string* QScroller_Tr(const char* s) {
	QString _ret = QScroller::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QScroller_TrUtf8(const char* s) {
	QString _ret = QScroller::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

bool QScroller_HasScroller(QObject* target) {
	return QScroller::hasScroller(target);
}

QScroller* QScroller_Scroller(QObject* target) {
	return QScroller::scroller(target);
}

QScroller* QScroller_ScrollerWithTarget(QObject* target) {
	return (QScroller*) QScroller::scroller(target);
}

uintptr_t QScroller_GrabGesture(QObject* target) {
	Qt::GestureType _ret = QScroller::grabGesture(target);
	return static_cast<uintptr_t>(_ret);
}

uintptr_t QScroller_GrabbedGesture(QObject* target) {
	Qt::GestureType _ret = QScroller::grabbedGesture(target);
	return static_cast<uintptr_t>(_ret);
}

void QScroller_UngrabGesture(QObject* target) {
	QScroller::ungrabGesture(target);
}

struct miqt_array* QScroller_ActiveScrollers() {
	QList<QScroller*> _ret = QScroller::activeScrollers();
	// Convert QList<> from C++ memory to manually-managed C memory
	QScroller** _arr = static_cast<QScroller**>(malloc(sizeof(QScroller*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

QObject* QScroller_Target(const QScroller* self) {
	return self->target();
}

uintptr_t QScroller_State(const QScroller* self) {
	QScroller::State _ret = self->state();
	return static_cast<uintptr_t>(_ret);
}

bool QScroller_HandleInput(QScroller* self, uintptr_t input, QPointF* position) {
	return self->handleInput(static_cast<QScroller::Input>(input), *position);
}

void QScroller_Stop(QScroller* self) {
	self->stop();
}

QPointF* QScroller_Velocity(const QScroller* self) {
	return new QPointF(self->velocity());
}

QPointF* QScroller_FinalPosition(const QScroller* self) {
	return new QPointF(self->finalPosition());
}

QPointF* QScroller_PixelPerMeter(const QScroller* self) {
	return new QPointF(self->pixelPerMeter());
}

QScrollerProperties* QScroller_ScrollerProperties(const QScroller* self) {
	return new QScrollerProperties(self->scrollerProperties());
}

void QScroller_SetSnapPositionsX(QScroller* self, struct miqt_array* /* of double */ positions) {
	QList<double> positions_QList;
	positions_QList.reserve(positions->len);
	double* positions_arr = static_cast<double*>(positions->data);
	for(size_t i = 0; i < positions->len; ++i) {
		positions_QList.push_back(positions_arr[i]);
	}
	self->setSnapPositionsX(positions_QList);
}

void QScroller_SetSnapPositionsX2(QScroller* self, double first, double interval) {
	self->setSnapPositionsX(static_cast<qreal>(first), static_cast<qreal>(interval));
}

void QScroller_SetSnapPositionsY(QScroller* self, struct miqt_array* /* of double */ positions) {
	QList<double> positions_QList;
	positions_QList.reserve(positions->len);
	double* positions_arr = static_cast<double*>(positions->data);
	for(size_t i = 0; i < positions->len; ++i) {
		positions_QList.push_back(positions_arr[i]);
	}
	self->setSnapPositionsY(positions_QList);
}

void QScroller_SetSnapPositionsY2(QScroller* self, double first, double interval) {
	self->setSnapPositionsY(static_cast<qreal>(first), static_cast<qreal>(interval));
}

void QScroller_SetScrollerProperties(QScroller* self, QScrollerProperties* prop) {
	self->setScrollerProperties(*prop);
}

void QScroller_ScrollTo(QScroller* self, QPointF* pos) {
	self->scrollTo(*pos);
}

void QScroller_ScrollTo2(QScroller* self, QPointF* pos, int scrollTime) {
	self->scrollTo(*pos, static_cast<int>(scrollTime));
}

void QScroller_EnsureVisible(QScroller* self, QRectF* rect, double xmargin, double ymargin) {
	self->ensureVisible(*rect, static_cast<qreal>(xmargin), static_cast<qreal>(ymargin));
}

void QScroller_EnsureVisible2(QScroller* self, QRectF* rect, double xmargin, double ymargin, int scrollTime) {
	self->ensureVisible(*rect, static_cast<qreal>(xmargin), static_cast<qreal>(ymargin), static_cast<int>(scrollTime));
}

void QScroller_ResendPrepareEvent(QScroller* self) {
	self->resendPrepareEvent();
}

void QScroller_StateChanged(QScroller* self, uintptr_t newstate) {
	self->stateChanged(static_cast<QScroller::State>(newstate));
}

void QScroller_connect_StateChanged(QScroller* self, void* slot) {
	QScroller::connect(self, static_cast<void (QScroller::*)(QScroller::State)>(&QScroller::stateChanged), self, [=](QScroller::State newstate) {
		QScroller::State newstate_ret = newstate;
		uintptr_t sigval1 = static_cast<uintptr_t>(newstate_ret);
		miqt_exec_callback_QScroller_StateChanged(slot, sigval1);
	});
}

void QScroller_ScrollerPropertiesChanged(QScroller* self, QScrollerProperties* param1) {
	self->scrollerPropertiesChanged(*param1);
}

void QScroller_connect_ScrollerPropertiesChanged(QScroller* self, void* slot) {
	QScroller::connect(self, static_cast<void (QScroller::*)(const QScrollerProperties&)>(&QScroller::scrollerPropertiesChanged), self, [=](const QScrollerProperties& param1) {
		const QScrollerProperties& param1_ret = param1;
		// Cast returned reference into pointer
		QScrollerProperties* sigval1 = const_cast<QScrollerProperties*>(&param1_ret);
		miqt_exec_callback_QScroller_ScrollerPropertiesChanged(slot, sigval1);
	});
}

struct miqt_string* QScroller_Tr2(const char* s, const char* c) {
	QString _ret = QScroller::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QScroller_Tr3(const char* s, const char* c, int n) {
	QString _ret = QScroller::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QScroller_TrUtf82(const char* s, const char* c) {
	QString _ret = QScroller::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QScroller_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QScroller::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

uintptr_t QScroller_GrabGesture2(QObject* target, uintptr_t gestureType) {
	Qt::GestureType _ret = QScroller::grabGesture(target, static_cast<QScroller::ScrollerGestureType>(gestureType));
	return static_cast<uintptr_t>(_ret);
}

bool QScroller_HandleInput3(QScroller* self, uintptr_t input, QPointF* position, long long timestamp) {
	return self->handleInput(static_cast<QScroller::Input>(input), *position, static_cast<qint64>(timestamp));
}

