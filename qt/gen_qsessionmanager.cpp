#include <QList>
#include <QMetaObject>
#include <QSessionManager>
#include <QString>
#include <QByteArray>
#include <cstring>
#include "qsessionmanager.h"
#include "gen_qsessionmanager.h"
#include "_cgo_export.h"

QMetaObject* QSessionManager_MetaObject(const QSessionManager* self) {
	return (QMetaObject*) self->metaObject();
}

struct miqt_string* QSessionManager_Tr(const char* s) {
	QString _ret = QSessionManager::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QSessionManager_TrUtf8(const char* s) {
	QString _ret = QSessionManager::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QSessionManager_SessionId(const QSessionManager* self) {
	QString _ret = self->sessionId();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QSessionManager_SessionKey(const QSessionManager* self) {
	QString _ret = self->sessionKey();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

bool QSessionManager_AllowsInteraction(QSessionManager* self) {
	return self->allowsInteraction();
}

bool QSessionManager_AllowsErrorInteraction(QSessionManager* self) {
	return self->allowsErrorInteraction();
}

void QSessionManager_Release(QSessionManager* self) {
	self->release();
}

void QSessionManager_Cancel(QSessionManager* self) {
	self->cancel();
}

void QSessionManager_SetRestartHint(QSessionManager* self, uintptr_t restartHint) {
	self->setRestartHint(static_cast<QSessionManager::RestartHint>(restartHint));
}

uintptr_t QSessionManager_RestartHint(const QSessionManager* self) {
	QSessionManager::RestartHint _ret = self->restartHint();
	return static_cast<uintptr_t>(_ret);
}

void QSessionManager_SetRestartCommand(QSessionManager* self, struct miqt_array* /* of struct miqt_string* */ restartCommand) {
	QList<QString> restartCommand_QList;
	restartCommand_QList.reserve(restartCommand->len);
	miqt_string** restartCommand_arr = static_cast<miqt_string**>(restartCommand->data);
	for(size_t i = 0; i < restartCommand->len; ++i) {
		restartCommand_QList.push_back(QString::fromUtf8(& restartCommand_arr[i]->data, restartCommand_arr[i]->len));
	}
	self->setRestartCommand(restartCommand_QList);
}

struct miqt_array* QSessionManager_RestartCommand(const QSessionManager* self) {
	QStringList _ret = self->restartCommand();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string** _arr = static_cast<struct miqt_string**>(malloc(sizeof(struct miqt_string*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		_arr[i] = miqt_strdup(_lv_b.data(), _lv_b.length());
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

void QSessionManager_SetDiscardCommand(QSessionManager* self, struct miqt_array* /* of struct miqt_string* */ discardCommand) {
	QList<QString> discardCommand_QList;
	discardCommand_QList.reserve(discardCommand->len);
	miqt_string** discardCommand_arr = static_cast<miqt_string**>(discardCommand->data);
	for(size_t i = 0; i < discardCommand->len; ++i) {
		discardCommand_QList.push_back(QString::fromUtf8(& discardCommand_arr[i]->data, discardCommand_arr[i]->len));
	}
	self->setDiscardCommand(discardCommand_QList);
}

struct miqt_array* QSessionManager_DiscardCommand(const QSessionManager* self) {
	QStringList _ret = self->discardCommand();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string** _arr = static_cast<struct miqt_string**>(malloc(sizeof(struct miqt_string*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		_arr[i] = miqt_strdup(_lv_b.data(), _lv_b.length());
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

void QSessionManager_SetManagerProperty(QSessionManager* self, struct miqt_string* name, struct miqt_string* value) {
	QString name_QString = QString::fromUtf8(&name->data, name->len);
	QString value_QString = QString::fromUtf8(&value->data, value->len);
	self->setManagerProperty(name_QString, value_QString);
}

void QSessionManager_SetManagerProperty2(QSessionManager* self, struct miqt_string* name, struct miqt_array* /* of struct miqt_string* */ value) {
	QString name_QString = QString::fromUtf8(&name->data, name->len);
	QList<QString> value_QList;
	value_QList.reserve(value->len);
	miqt_string** value_arr = static_cast<miqt_string**>(value->data);
	for(size_t i = 0; i < value->len; ++i) {
		value_QList.push_back(QString::fromUtf8(& value_arr[i]->data, value_arr[i]->len));
	}
	self->setManagerProperty(name_QString, value_QList);
}

bool QSessionManager_IsPhase2(const QSessionManager* self) {
	return self->isPhase2();
}

void QSessionManager_RequestPhase2(QSessionManager* self) {
	self->requestPhase2();
}

struct miqt_string* QSessionManager_Tr2(const char* s, const char* c) {
	QString _ret = QSessionManager::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QSessionManager_Tr3(const char* s, const char* c, int n) {
	QString _ret = QSessionManager::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QSessionManager_TrUtf82(const char* s, const char* c) {
	QString _ret = QSessionManager::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_string* QSessionManager_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QSessionManager::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

