#include <QAction>
#include <QMenu>
#include <QMetaObject>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QToolButton>
#include <QWidget>
#include "qtoolbutton.h"
#include "gen_qtoolbutton.h"
#include "_cgo_export.h"

QToolButton* QToolButton_new() {
	return new QToolButton();
}

QToolButton* QToolButton_new2(QWidget* parent) {
	return new QToolButton(parent);
}

QMetaObject* QToolButton_MetaObject(const QToolButton* self) {
	return (QMetaObject*) self->metaObject();
}

struct miqt_string* QToolButton_Tr(const char* s) {
	QString _ret = QToolButton::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QToolButton_TrUtf8(const char* s) {
	QString _ret = QToolButton::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

QSize* QToolButton_SizeHint(const QToolButton* self) {
	return new QSize(self->sizeHint());
}

QSize* QToolButton_MinimumSizeHint(const QToolButton* self) {
	return new QSize(self->minimumSizeHint());
}

uintptr_t QToolButton_ToolButtonStyle(const QToolButton* self) {
	Qt::ToolButtonStyle _ret = self->toolButtonStyle();
	return static_cast<uintptr_t>(_ret);
}

uintptr_t QToolButton_ArrowType(const QToolButton* self) {
	Qt::ArrowType _ret = self->arrowType();
	return static_cast<uintptr_t>(_ret);
}

void QToolButton_SetArrowType(QToolButton* self, uintptr_t typeVal) {
	self->setArrowType(static_cast<Qt::ArrowType>(typeVal));
}

void QToolButton_SetMenu(QToolButton* self, QMenu* menu) {
	self->setMenu(menu);
}

QMenu* QToolButton_Menu(const QToolButton* self) {
	return self->menu();
}

void QToolButton_SetPopupMode(QToolButton* self, uintptr_t mode) {
	self->setPopupMode(static_cast<QToolButton::ToolButtonPopupMode>(mode));
}

uintptr_t QToolButton_PopupMode(const QToolButton* self) {
	QToolButton::ToolButtonPopupMode _ret = self->popupMode();
	return static_cast<uintptr_t>(_ret);
}

QAction* QToolButton_DefaultAction(const QToolButton* self) {
	return self->defaultAction();
}

void QToolButton_SetAutoRaise(QToolButton* self, bool enable) {
	self->setAutoRaise(enable);
}

bool QToolButton_AutoRaise(const QToolButton* self) {
	return self->autoRaise();
}

void QToolButton_ShowMenu(QToolButton* self) {
	self->showMenu();
}

void QToolButton_SetToolButtonStyle(QToolButton* self, uintptr_t style) {
	self->setToolButtonStyle(static_cast<Qt::ToolButtonStyle>(style));
}

void QToolButton_SetDefaultAction(QToolButton* self, QAction* defaultAction) {
	self->setDefaultAction(defaultAction);
}

void QToolButton_Triggered(QToolButton* self, QAction* param1) {
	self->triggered(param1);
}

void QToolButton_connect_Triggered(QToolButton* self, void* slot) {
	QToolButton::connect(self, static_cast<void (QToolButton::*)(QAction*)>(&QToolButton::triggered), self, [=](QAction* param1) {
		QAction* sigval1 = param1;
		miqt_exec_callback_QToolButton_Triggered(slot, sigval1);
	});
}

struct miqt_string* QToolButton_Tr2(const char* s, const char* c) {
	QString _ret = QToolButton::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QToolButton_Tr3(const char* s, const char* c, int n) {
	QString _ret = QToolButton::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QToolButton_TrUtf82(const char* s, const char* c) {
	QString _ret = QToolButton::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QToolButton_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QToolButton::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

void QToolButton_Delete(QToolButton* self) {
	delete self;
}

