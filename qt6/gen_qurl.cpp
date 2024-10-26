#include <QByteArray>
#include <QList>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QUrl>
#include <QUrlQuery>
#include <qurl.h>
#include "gen_qurl.h"
#include "_cgo_export.h"

QUrl* QUrl_new() {
	return new QUrl();
}

QUrl* QUrl_new2(QUrl* copyVal) {
	return new QUrl(*copyVal);
}

QUrl* QUrl_new3(struct miqt_string url) {
	QString url_QString = QString::fromUtf8(url.data, url.len);
	return new QUrl(url_QString);
}

QUrl* QUrl_new4(struct miqt_string url, int mode) {
	QString url_QString = QString::fromUtf8(url.data, url.len);
	return new QUrl(url_QString, static_cast<QUrl::ParsingMode>(mode));
}

void QUrl_OperatorAssign(QUrl* self, QUrl* copyVal) {
	self->operator=(*copyVal);
}

void QUrl_OperatorAssignWithUrl(QUrl* self, struct miqt_string url) {
	QString url_QString = QString::fromUtf8(url.data, url.len);
	self->operator=(url_QString);
}

void QUrl_Swap(QUrl* self, QUrl* other) {
	self->swap(*other);
}

void QUrl_SetUrl(QUrl* self, struct miqt_string url) {
	QString url_QString = QString::fromUtf8(url.data, url.len);
	self->setUrl(url_QString);
}

struct miqt_string QUrl_Url(const QUrl* self) {
	QString _ret = self->url();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QUrl_ToString(const QUrl* self) {
	QString _ret = self->toString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QUrl_ToDisplayString(const QUrl* self) {
	QString _ret = self->toDisplayString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QUrl_ToEncoded(const QUrl* self) {
	QByteArray _qb = self->toEncoded();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

QUrl* QUrl_FromEncoded(struct miqt_string url) {
	QByteArray url_QByteArray(url.data, url.len);
	return new QUrl(QUrl::fromEncoded(url_QByteArray));
}

QUrl* QUrl_FromUserInput(struct miqt_string userInput) {
	QString userInput_QString = QString::fromUtf8(userInput.data, userInput.len);
	return new QUrl(QUrl::fromUserInput(userInput_QString));
}

bool QUrl_IsValid(const QUrl* self) {
	return self->isValid();
}

struct miqt_string QUrl_ErrorString(const QUrl* self) {
	QString _ret = self->errorString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QUrl_IsEmpty(const QUrl* self) {
	return self->isEmpty();
}

void QUrl_Clear(QUrl* self) {
	self->clear();
}

void QUrl_SetScheme(QUrl* self, struct miqt_string scheme) {
	QString scheme_QString = QString::fromUtf8(scheme.data, scheme.len);
	self->setScheme(scheme_QString);
}

struct miqt_string QUrl_Scheme(const QUrl* self) {
	QString _ret = self->scheme();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QUrl_SetAuthority(QUrl* self, struct miqt_string authority) {
	QString authority_QString = QString::fromUtf8(authority.data, authority.len);
	self->setAuthority(authority_QString);
}

struct miqt_string QUrl_Authority(const QUrl* self) {
	QString _ret = self->authority();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QUrl_SetUserInfo(QUrl* self, struct miqt_string userInfo) {
	QString userInfo_QString = QString::fromUtf8(userInfo.data, userInfo.len);
	self->setUserInfo(userInfo_QString);
}

struct miqt_string QUrl_UserInfo(const QUrl* self) {
	QString _ret = self->userInfo();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QUrl_SetUserName(QUrl* self, struct miqt_string userName) {
	QString userName_QString = QString::fromUtf8(userName.data, userName.len);
	self->setUserName(userName_QString);
}

struct miqt_string QUrl_UserName(const QUrl* self) {
	QString _ret = self->userName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QUrl_SetPassword(QUrl* self, struct miqt_string password) {
	QString password_QString = QString::fromUtf8(password.data, password.len);
	self->setPassword(password_QString);
}

struct miqt_string QUrl_Password(const QUrl* self) {
	QString _ret = self->password();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QUrl_SetHost(QUrl* self, struct miqt_string host) {
	QString host_QString = QString::fromUtf8(host.data, host.len);
	self->setHost(host_QString);
}

struct miqt_string QUrl_Host(const QUrl* self) {
	QString _ret = self->host();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QUrl_SetPort(QUrl* self, int port) {
	self->setPort(static_cast<int>(port));
}

int QUrl_Port(const QUrl* self) {
	return self->port();
}

void QUrl_SetPath(QUrl* self, struct miqt_string path) {
	QString path_QString = QString::fromUtf8(path.data, path.len);
	self->setPath(path_QString);
}

struct miqt_string QUrl_Path(const QUrl* self) {
	QString _ret = self->path();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QUrl_FileName(const QUrl* self) {
	QString _ret = self->fileName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QUrl_HasQuery(const QUrl* self) {
	return self->hasQuery();
}

void QUrl_SetQuery(QUrl* self, struct miqt_string query) {
	QString query_QString = QString::fromUtf8(query.data, query.len);
	self->setQuery(query_QString);
}

void QUrl_SetQueryWithQuery(QUrl* self, QUrlQuery* query) {
	self->setQuery(*query);
}

struct miqt_string QUrl_Query(const QUrl* self) {
	QString _ret = self->query();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QUrl_HasFragment(const QUrl* self) {
	return self->hasFragment();
}

struct miqt_string QUrl_Fragment(const QUrl* self) {
	QString _ret = self->fragment();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QUrl_SetFragment(QUrl* self, struct miqt_string fragment) {
	QString fragment_QString = QString::fromUtf8(fragment.data, fragment.len);
	self->setFragment(fragment_QString);
}

QUrl* QUrl_Resolved(const QUrl* self, QUrl* relative) {
	return new QUrl(self->resolved(*relative));
}

bool QUrl_IsRelative(const QUrl* self) {
	return self->isRelative();
}

bool QUrl_IsParentOf(const QUrl* self, QUrl* url) {
	return self->isParentOf(*url);
}

bool QUrl_IsLocalFile(const QUrl* self) {
	return self->isLocalFile();
}

QUrl* QUrl_FromLocalFile(struct miqt_string localfile) {
	QString localfile_QString = QString::fromUtf8(localfile.data, localfile.len);
	return new QUrl(QUrl::fromLocalFile(localfile_QString));
}

struct miqt_string QUrl_ToLocalFile(const QUrl* self) {
	QString _ret = self->toLocalFile();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QUrl_Detach(QUrl* self) {
	self->detach();
}

bool QUrl_IsDetached(const QUrl* self) {
	return self->isDetached();
}

bool QUrl_OperatorLesser(const QUrl* self, QUrl* url) {
	return self->operator<(*url);
}

bool QUrl_OperatorEqual(const QUrl* self, QUrl* url) {
	return self->operator==(*url);
}

bool QUrl_OperatorNotEqual(const QUrl* self, QUrl* url) {
	return self->operator!=(*url);
}

struct miqt_string QUrl_FromPercentEncoding(struct miqt_string param1) {
	QByteArray param1_QByteArray(param1.data, param1.len);
	QString _ret = QUrl::fromPercentEncoding(param1_QByteArray);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QUrl_ToPercentEncoding(struct miqt_string param1) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	QByteArray _qb = QUrl::toPercentEncoding(param1_QString);
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

struct miqt_string QUrl_FromAce(struct miqt_string domain) {
	QByteArray domain_QByteArray(domain.data, domain.len);
	QString _ret = QUrl::fromAce(domain_QByteArray);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QUrl_ToAce(struct miqt_string domain) {
	QString domain_QString = QString::fromUtf8(domain.data, domain.len);
	QByteArray _qb = QUrl::toAce(domain_QString);
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

struct miqt_array* QUrl_IdnWhitelist() {
	QStringList _ret = QUrl::idnWhitelist();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array* QUrl_ToStringList(struct miqt_array* /* of QUrl* */ uris) {
	QList<QUrl> uris_QList;
	uris_QList.reserve(uris->len);
	QUrl** uris_arr = static_cast<QUrl**>(uris->data);
	for(size_t i = 0; i < uris->len; ++i) {
		uris_QList.push_back(*(uris_arr[i]));
	}
	QStringList _ret = QUrl::toStringList(uris_QList);
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array* QUrl_FromStringList(struct miqt_array* /* of struct miqt_string */ uris) {
	QStringList uris_QList;
	uris_QList.reserve(uris->len);
	struct miqt_string* uris_arr = static_cast<struct miqt_string*>(uris->data);
	for(size_t i = 0; i < uris->len; ++i) {
		QString uris_arr_i_QString = QString::fromUtf8(uris_arr[i].data, uris_arr[i].len);
		uris_QList.push_back(uris_arr_i_QString);
	}
	QList<QUrl> _ret = QUrl::fromStringList(uris_QList);
	// Convert QList<> from C++ memory to manually-managed C memory
	QUrl** _arr = static_cast<QUrl**>(malloc(sizeof(QUrl*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QUrl(_ret[i]);
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

void QUrl_SetIdnWhitelist(struct miqt_array* /* of struct miqt_string */ idnWhitelist) {
	QStringList idnWhitelist_QList;
	idnWhitelist_QList.reserve(idnWhitelist->len);
	struct miqt_string* idnWhitelist_arr = static_cast<struct miqt_string*>(idnWhitelist->data);
	for(size_t i = 0; i < idnWhitelist->len; ++i) {
		QString idnWhitelist_arr_i_QString = QString::fromUtf8(idnWhitelist_arr[i].data, idnWhitelist_arr[i].len);
		idnWhitelist_QList.push_back(idnWhitelist_arr_i_QString);
	}
	QUrl::setIdnWhitelist(idnWhitelist_QList);
}

void QUrl_SetUrl2(QUrl* self, struct miqt_string url, int mode) {
	QString url_QString = QString::fromUtf8(url.data, url.len);
	self->setUrl(url_QString, static_cast<QUrl::ParsingMode>(mode));
}

QUrl* QUrl_FromEncoded2(struct miqt_string url, int mode) {
	QByteArray url_QByteArray(url.data, url.len);
	return new QUrl(QUrl::fromEncoded(url_QByteArray, static_cast<QUrl::ParsingMode>(mode)));
}

QUrl* QUrl_FromUserInput2(struct miqt_string userInput, struct miqt_string workingDirectory) {
	QString userInput_QString = QString::fromUtf8(userInput.data, userInput.len);
	QString workingDirectory_QString = QString::fromUtf8(workingDirectory.data, workingDirectory.len);
	return new QUrl(QUrl::fromUserInput(userInput_QString, workingDirectory_QString));
}

QUrl* QUrl_FromUserInput3(struct miqt_string userInput, struct miqt_string workingDirectory, int options) {
	QString userInput_QString = QString::fromUtf8(userInput.data, userInput.len);
	QString workingDirectory_QString = QString::fromUtf8(workingDirectory.data, workingDirectory.len);
	return new QUrl(QUrl::fromUserInput(userInput_QString, workingDirectory_QString, static_cast<QUrl::UserInputResolutionOptions>(options)));
}

void QUrl_SetAuthority2(QUrl* self, struct miqt_string authority, int mode) {
	QString authority_QString = QString::fromUtf8(authority.data, authority.len);
	self->setAuthority(authority_QString, static_cast<QUrl::ParsingMode>(mode));
}

struct miqt_string QUrl_Authority1(const QUrl* self, unsigned int options) {
	QString _ret = self->authority(static_cast<QUrl::ComponentFormattingOptions>(options));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QUrl_SetUserInfo2(QUrl* self, struct miqt_string userInfo, int mode) {
	QString userInfo_QString = QString::fromUtf8(userInfo.data, userInfo.len);
	self->setUserInfo(userInfo_QString, static_cast<QUrl::ParsingMode>(mode));
}

struct miqt_string QUrl_UserInfo1(const QUrl* self, unsigned int options) {
	QString _ret = self->userInfo(static_cast<QUrl::ComponentFormattingOptions>(options));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QUrl_SetUserName2(QUrl* self, struct miqt_string userName, int mode) {
	QString userName_QString = QString::fromUtf8(userName.data, userName.len);
	self->setUserName(userName_QString, static_cast<QUrl::ParsingMode>(mode));
}

struct miqt_string QUrl_UserName1(const QUrl* self, unsigned int options) {
	QString _ret = self->userName(static_cast<QUrl::ComponentFormattingOptions>(options));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QUrl_SetPassword2(QUrl* self, struct miqt_string password, int mode) {
	QString password_QString = QString::fromUtf8(password.data, password.len);
	self->setPassword(password_QString, static_cast<QUrl::ParsingMode>(mode));
}

struct miqt_string QUrl_Password1(const QUrl* self, unsigned int param1) {
	QString _ret = self->password(static_cast<QUrl::ComponentFormattingOptions>(param1));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QUrl_SetHost2(QUrl* self, struct miqt_string host, int mode) {
	QString host_QString = QString::fromUtf8(host.data, host.len);
	self->setHost(host_QString, static_cast<QUrl::ParsingMode>(mode));
}

struct miqt_string QUrl_Host1(const QUrl* self, unsigned int param1) {
	QString _ret = self->host(static_cast<QUrl::ComponentFormattingOptions>(param1));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QUrl_Port1(const QUrl* self, int defaultPort) {
	return self->port(static_cast<int>(defaultPort));
}

void QUrl_SetPath2(QUrl* self, struct miqt_string path, int mode) {
	QString path_QString = QString::fromUtf8(path.data, path.len);
	self->setPath(path_QString, static_cast<QUrl::ParsingMode>(mode));
}

struct miqt_string QUrl_Path1(const QUrl* self, unsigned int options) {
	QString _ret = self->path(static_cast<QUrl::ComponentFormattingOptions>(options));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QUrl_FileName1(const QUrl* self, unsigned int options) {
	QString _ret = self->fileName(static_cast<QUrl::ComponentFormattingOptions>(options));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QUrl_SetQuery2(QUrl* self, struct miqt_string query, int mode) {
	QString query_QString = QString::fromUtf8(query.data, query.len);
	self->setQuery(query_QString, static_cast<QUrl::ParsingMode>(mode));
}

struct miqt_string QUrl_Query1(const QUrl* self, unsigned int param1) {
	QString _ret = self->query(static_cast<QUrl::ComponentFormattingOptions>(param1));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QUrl_Fragment1(const QUrl* self, unsigned int options) {
	QString _ret = self->fragment(static_cast<QUrl::ComponentFormattingOptions>(options));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QUrl_SetFragment2(QUrl* self, struct miqt_string fragment, int mode) {
	QString fragment_QString = QString::fromUtf8(fragment.data, fragment.len);
	self->setFragment(fragment_QString, static_cast<QUrl::ParsingMode>(mode));
}

struct miqt_string QUrl_ToPercentEncoding2(struct miqt_string param1, struct miqt_string exclude) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	QByteArray exclude_QByteArray(exclude.data, exclude.len);
	QByteArray _qb = QUrl::toPercentEncoding(param1_QString, exclude_QByteArray);
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

struct miqt_string QUrl_ToPercentEncoding3(struct miqt_string param1, struct miqt_string exclude, struct miqt_string include) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	QByteArray exclude_QByteArray(exclude.data, exclude.len);
	QByteArray include_QByteArray(include.data, include.len);
	QByteArray _qb = QUrl::toPercentEncoding(param1_QString, exclude_QByteArray, include_QByteArray);
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

struct miqt_string QUrl_FromAce2(struct miqt_string domain, unsigned int options) {
	QByteArray domain_QByteArray(domain.data, domain.len);
	QString _ret = QUrl::fromAce(domain_QByteArray, static_cast<QUrl::AceProcessingOptions>(options));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QUrl_ToAce2(struct miqt_string domain, unsigned int options) {
	QString domain_QString = QString::fromUtf8(domain.data, domain.len);
	QByteArray _qb = QUrl::toAce(domain_QString, static_cast<QUrl::AceProcessingOptions>(options));
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

struct miqt_array* QUrl_FromStringList2(struct miqt_array* /* of struct miqt_string */ uris, int mode) {
	QStringList uris_QList;
	uris_QList.reserve(uris->len);
	struct miqt_string* uris_arr = static_cast<struct miqt_string*>(uris->data);
	for(size_t i = 0; i < uris->len; ++i) {
		QString uris_arr_i_QString = QString::fromUtf8(uris_arr[i].data, uris_arr[i].len);
		uris_QList.push_back(uris_arr_i_QString);
	}
	QList<QUrl> _ret = QUrl::fromStringList(uris_QList, static_cast<QUrl::ParsingMode>(mode));
	// Convert QList<> from C++ memory to manually-managed C memory
	QUrl** _arr = static_cast<QUrl**>(malloc(sizeof(QUrl*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QUrl(_ret[i]);
	}
	struct miqt_array* _out = static_cast<struct miqt_array*>(malloc(sizeof(struct miqt_array)));
	_out->len = _ret.length();
	_out->data = static_cast<void*>(_arr);
	return _out;
}

void QUrl_Delete(QUrl* self) {
	delete self;
}
