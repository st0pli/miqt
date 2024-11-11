#include <QSslCertificate>
#include <QSslError>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qsslerror.h>
#include "gen_qsslerror.h"
#include "_cgo_export.h"

QSslError* QSslError_new() {
	return new QSslError();
}

QSslError* QSslError_new2(int error) {
	return new QSslError(static_cast<QSslError::SslError>(error));
}

QSslError* QSslError_new3(int error, QSslCertificate* certificate) {
	return new QSslError(static_cast<QSslError::SslError>(error), *certificate);
}

QSslError* QSslError_new4(QSslError* other) {
	return new QSslError(*other);
}

void QSslError_Swap(QSslError* self, QSslError* other) {
	self->swap(*other);
}

void QSslError_OperatorAssign(QSslError* self, QSslError* other) {
	self->operator=(*other);
}

bool QSslError_OperatorEqual(const QSslError* self, QSslError* other) {
	return self->operator==(*other);
}

bool QSslError_OperatorNotEqual(const QSslError* self, QSslError* other) {
	return self->operator!=(*other);
}

int QSslError_Error(const QSslError* self) {
	QSslError::SslError _ret = self->error();
	return static_cast<int>(_ret);
}

struct miqt_string QSslError_ErrorString(const QSslError* self) {
	QString _ret = self->errorString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QSslCertificate* QSslError_Certificate(const QSslError* self) {
	return new QSslCertificate(self->certificate());
}

void QSslError_Delete(QSslError* self) {
	delete self;
}
