#include <QByteArray>
#include <QFileInfo>
#include <QIODevice>
#include <QList>
#include <QMimeDatabase>
#include <QMimeType>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QUrl>
#include "qmimedatabase.h"
#include "gen_qmimedatabase.h"
#include "_cgo_export.h"

QMimeDatabase* QMimeDatabase_new() {
	return new QMimeDatabase();
}

QMimeType* QMimeDatabase_MimeTypeForName(const QMimeDatabase* self, struct miqt_string* nameOrAlias) {
	QString nameOrAlias_QString = QString::fromUtf8(&nameOrAlias->data, nameOrAlias->len);
	return new QMimeType(self->mimeTypeForName(nameOrAlias_QString));
}

QMimeType* QMimeDatabase_MimeTypeForFile(const QMimeDatabase* self, struct miqt_string* fileName) {
	QString fileName_QString = QString::fromUtf8(&fileName->data, fileName->len);
	return new QMimeType(self->mimeTypeForFile(fileName_QString));
}

QMimeType* QMimeDatabase_MimeTypeForFileWithFileInfo(const QMimeDatabase* self, QFileInfo* fileInfo) {
	return new QMimeType(self->mimeTypeForFile(*fileInfo));
}

struct miqt_array* QMimeDatabase_MimeTypesForFileName(const QMimeDatabase* self, struct miqt_string* fileName) {
	QString fileName_QString = QString::fromUtf8(&fileName->data, fileName->len);
	QList<QMimeType> _ret = self->mimeTypesForFileName(fileName_QString);
	// Convert QList<> from C++ memory to manually-managed C memory
	QMimeType** _arr = static_cast<QMimeType**>(malloc(sizeof(QMimeType*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QMimeType(_ret[i]);
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

QMimeType* QMimeDatabase_MimeTypeForData(const QMimeDatabase* self, QByteArray* data) {
	return new QMimeType(self->mimeTypeForData(*data));
}

QMimeType* QMimeDatabase_MimeTypeForDataWithDevice(const QMimeDatabase* self, QIODevice* device) {
	return new QMimeType(self->mimeTypeForData(device));
}

QMimeType* QMimeDatabase_MimeTypeForUrl(const QMimeDatabase* self, QUrl* url) {
	return new QMimeType(self->mimeTypeForUrl(*url));
}

QMimeType* QMimeDatabase_MimeTypeForFileNameAndData(const QMimeDatabase* self, struct miqt_string* fileName, QIODevice* device) {
	QString fileName_QString = QString::fromUtf8(&fileName->data, fileName->len);
	return new QMimeType(self->mimeTypeForFileNameAndData(fileName_QString, device));
}

QMimeType* QMimeDatabase_MimeTypeForFileNameAndData2(const QMimeDatabase* self, struct miqt_string* fileName, QByteArray* data) {
	QString fileName_QString = QString::fromUtf8(&fileName->data, fileName->len);
	return new QMimeType(self->mimeTypeForFileNameAndData(fileName_QString, *data));
}

struct miqt_string* QMimeDatabase_SuffixForFileName(const QMimeDatabase* self, struct miqt_string* fileName) {
	QString fileName_QString = QString::fromUtf8(&fileName->data, fileName->len);
	QString _ret = self->suffixForFileName(fileName_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	return miqt_strdup(_b.data(), _b.length());
}

struct miqt_array* QMimeDatabase_AllMimeTypes(const QMimeDatabase* self) {
	QList<QMimeType> _ret = self->allMimeTypes();
	// Convert QList<> from C++ memory to manually-managed C memory
	QMimeType** _arr = static_cast<QMimeType**>(malloc(sizeof(QMimeType*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QMimeType(_ret[i]);
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

QMimeType* QMimeDatabase_MimeTypeForFile2(const QMimeDatabase* self, struct miqt_string* fileName, uintptr_t mode) {
	QString fileName_QString = QString::fromUtf8(&fileName->data, fileName->len);
	return new QMimeType(self->mimeTypeForFile(fileName_QString, static_cast<QMimeDatabase::MatchMode>(mode)));
}

QMimeType* QMimeDatabase_MimeTypeForFile22(const QMimeDatabase* self, QFileInfo* fileInfo, uintptr_t mode) {
	return new QMimeType(self->mimeTypeForFile(*fileInfo, static_cast<QMimeDatabase::MatchMode>(mode)));
}

void QMimeDatabase_Delete(QMimeDatabase* self) {
	delete self;
}

