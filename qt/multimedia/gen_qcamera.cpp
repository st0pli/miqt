#include <QAbstractVideoSurface>
#include <QByteArray>
#include <QCamera>
#define WORKAROUND_INNER_CLASS_DEFINITION_QCamera__FrameRateRange
#include <QCameraExposure>
#include <QCameraFocus>
#include <QCameraImageProcessing>
#include <QCameraInfo>
#include <QCameraViewfinderSettings>
#include <QGraphicsVideoItem>
#include <QList>
#include <QMediaObject>
#include <QMediaService>
#include <QMetaObject>
#include <QObject>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVideoWidget>
#include <qcamera.h>
#include "gen_qcamera.h"

#ifndef _Bool
#define _Bool bool
#endif
#include "_cgo_export.h"

class MiqtVirtualQCamera : public virtual QCamera {
public:

	MiqtVirtualQCamera(): QCamera() {};
	MiqtVirtualQCamera(const QByteArray& deviceName): QCamera(deviceName) {};
	MiqtVirtualQCamera(const QCameraInfo& cameraInfo): QCamera(cameraInfo) {};
	MiqtVirtualQCamera(QCamera::Position position): QCamera(position) {};
	MiqtVirtualQCamera(QObject* parent): QCamera(parent) {};
	MiqtVirtualQCamera(const QByteArray& deviceName, QObject* parent): QCamera(deviceName, parent) {};
	MiqtVirtualQCamera(const QCameraInfo& cameraInfo, QObject* parent): QCamera(cameraInfo, parent) {};
	MiqtVirtualQCamera(QCamera::Position position, QObject* parent): QCamera(position, parent) {};

	virtual ~MiqtVirtualQCamera() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Availability = 0;

	// Subclass to allow providing a Go implementation
	virtual QMultimedia::AvailabilityStatus availability() const override {
		if (handle__Availability == 0) {
			return QCamera::availability();
		}
		

		int callback_return_value = miqt_exec_callback_QCamera_Availability(const_cast<MiqtVirtualQCamera*>(this), handle__Availability);

		return static_cast<QMultimedia::AvailabilityStatus>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_Availability() const {

		QMultimedia::AvailabilityStatus _ret = QCamera::availability();
		return static_cast<int>(_ret);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__IsAvailable = 0;

	// Subclass to allow providing a Go implementation
	virtual bool isAvailable() const override {
		if (handle__IsAvailable == 0) {
			return QCamera::isAvailable();
		}
		

		bool callback_return_value = miqt_exec_callback_QCamera_IsAvailable(const_cast<MiqtVirtualQCamera*>(this), handle__IsAvailable);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_IsAvailable() const {

		return QCamera::isAvailable();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Service = 0;

	// Subclass to allow providing a Go implementation
	virtual QMediaService* service() const override {
		if (handle__Service == 0) {
			return QCamera::service();
		}
		

		QMediaService* callback_return_value = miqt_exec_callback_QCamera_Service(const_cast<MiqtVirtualQCamera*>(this), handle__Service);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMediaService* virtualbase_Service() const {

		return QCamera::service();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Bind = 0;

	// Subclass to allow providing a Go implementation
	virtual bool bind(QObject* param1) override {
		if (handle__Bind == 0) {
			return QCamera::bind(param1);
		}
		
		QObject* sigval1 = param1;

		bool callback_return_value = miqt_exec_callback_QCamera_Bind(this, handle__Bind, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_Bind(QObject* param1) {

		return QCamera::bind(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Unbind = 0;

	// Subclass to allow providing a Go implementation
	virtual void unbind(QObject* param1) override {
		if (handle__Unbind == 0) {
			QCamera::unbind(param1);
			return;
		}
		
		QObject* sigval1 = param1;

		miqt_exec_callback_QCamera_Unbind(this, handle__Unbind, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_Unbind(QObject* param1) {

		QCamera::unbind(param1);

	}

};

QCamera* QCamera_new() {
	return new MiqtVirtualQCamera();
}

QCamera* QCamera_new2(struct miqt_string deviceName) {
	QByteArray deviceName_QByteArray(deviceName.data, deviceName.len);
	return new MiqtVirtualQCamera(deviceName_QByteArray);
}

QCamera* QCamera_new3(QCameraInfo* cameraInfo) {
	return new MiqtVirtualQCamera(*cameraInfo);
}

QCamera* QCamera_new4(int position) {
	return new MiqtVirtualQCamera(static_cast<QCamera::Position>(position));
}

QCamera* QCamera_new5(QObject* parent) {
	return new MiqtVirtualQCamera(parent);
}

QCamera* QCamera_new6(struct miqt_string deviceName, QObject* parent) {
	QByteArray deviceName_QByteArray(deviceName.data, deviceName.len);
	return new MiqtVirtualQCamera(deviceName_QByteArray, parent);
}

QCamera* QCamera_new7(QCameraInfo* cameraInfo, QObject* parent) {
	return new MiqtVirtualQCamera(*cameraInfo, parent);
}

QCamera* QCamera_new8(int position, QObject* parent) {
	return new MiqtVirtualQCamera(static_cast<QCamera::Position>(position), parent);
}

void QCamera_virtbase(QCamera* src, QMediaObject** outptr_QMediaObject) {
	*outptr_QMediaObject = static_cast<QMediaObject*>(src);
}

QMetaObject* QCamera_MetaObject(const QCamera* self) {
	return (QMetaObject*) self->metaObject();
}

void* QCamera_Metacast(QCamera* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QCamera_Tr(const char* s) {
	QString _ret = QCamera::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QCamera_TrUtf8(const char* s) {
	QString _ret = QCamera::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_array /* of struct miqt_string */  QCamera_AvailableDevices() {
	QList<QByteArray> _ret = QCamera::availableDevices();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QByteArray _lv_qb = _ret[i];
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_qb.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_qb.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_string QCamera_DeviceDescription(struct miqt_string device) {
	QByteArray device_QByteArray(device.data, device.len);
	QString _ret = QCamera::deviceDescription(device_QByteArray);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QCamera_Availability(const QCamera* self) {
	QMultimedia::AvailabilityStatus _ret = self->availability();
	return static_cast<int>(_ret);
}

int QCamera_State(const QCamera* self) {
	QCamera::State _ret = self->state();
	return static_cast<int>(_ret);
}

int QCamera_Status(const QCamera* self) {
	QCamera::Status _ret = self->status();
	return static_cast<int>(_ret);
}

int QCamera_CaptureMode(const QCamera* self) {
	QCamera::CaptureModes _ret = self->captureMode();
	return static_cast<int>(_ret);
}

bool QCamera_IsCaptureModeSupported(const QCamera* self, int mode) {
	return self->isCaptureModeSupported(static_cast<QCamera::CaptureModes>(mode));
}

QCameraExposure* QCamera_Exposure(const QCamera* self) {
	return self->exposure();
}

QCameraFocus* QCamera_Focus(const QCamera* self) {
	return self->focus();
}

QCameraImageProcessing* QCamera_ImageProcessing(const QCamera* self) {
	return self->imageProcessing();
}

void QCamera_SetViewfinder(QCamera* self, QVideoWidget* viewfinder) {
	self->setViewfinder(viewfinder);
}

void QCamera_SetViewfinderWithViewfinder(QCamera* self, QGraphicsVideoItem* viewfinder) {
	self->setViewfinder(viewfinder);
}

void QCamera_SetViewfinderWithSurface(QCamera* self, QAbstractVideoSurface* surface) {
	self->setViewfinder(surface);
}

QCameraViewfinderSettings* QCamera_ViewfinderSettings(const QCamera* self) {
	return new QCameraViewfinderSettings(self->viewfinderSettings());
}

void QCamera_SetViewfinderSettings(QCamera* self, QCameraViewfinderSettings* settings) {
	self->setViewfinderSettings(*settings);
}

struct miqt_array /* of QCameraViewfinderSettings* */  QCamera_SupportedViewfinderSettings(const QCamera* self) {
	QList<QCameraViewfinderSettings> _ret = self->supportedViewfinderSettings();
	// Convert QList<> from C++ memory to manually-managed C memory
	QCameraViewfinderSettings** _arr = static_cast<QCameraViewfinderSettings**>(malloc(sizeof(QCameraViewfinderSettings*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QCameraViewfinderSettings(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of QSize* */  QCamera_SupportedViewfinderResolutions(const QCamera* self) {
	QList<QSize> _ret = self->supportedViewfinderResolutions();
	// Convert QList<> from C++ memory to manually-managed C memory
	QSize** _arr = static_cast<QSize**>(malloc(sizeof(QSize*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QSize(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of QCamera__FrameRateRange* */  QCamera_SupportedViewfinderFrameRateRanges(const QCamera* self) {
	QList<QCamera::FrameRateRange> _ret = self->supportedViewfinderFrameRateRanges();
	// Convert QList<> from C++ memory to manually-managed C memory
	QCamera__FrameRateRange** _arr = static_cast<QCamera__FrameRateRange**>(malloc(sizeof(QCamera__FrameRateRange*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QCamera::FrameRateRange(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of int */  QCamera_SupportedViewfinderPixelFormats(const QCamera* self) {
	QList<QVideoFrame::PixelFormat> _ret = self->supportedViewfinderPixelFormats();
	// Convert QList<> from C++ memory to manually-managed C memory
	int* _arr = static_cast<int*>(malloc(sizeof(int) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QVideoFrame::PixelFormat _lv_ret = _ret[i];
		_arr[i] = static_cast<int>(_lv_ret);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

int QCamera_Error(const QCamera* self) {
	QCamera::Error _ret = self->error();
	return static_cast<int>(_ret);
}

struct miqt_string QCamera_ErrorString(const QCamera* self) {
	QString _ret = self->errorString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QCamera_SupportedLocks(const QCamera* self) {
	QCamera::LockTypes _ret = self->supportedLocks();
	return static_cast<int>(_ret);
}

int QCamera_RequestedLocks(const QCamera* self) {
	QCamera::LockTypes _ret = self->requestedLocks();
	return static_cast<int>(_ret);
}

int QCamera_LockStatus(const QCamera* self) {
	QCamera::LockStatus _ret = self->lockStatus();
	return static_cast<int>(_ret);
}

int QCamera_LockStatusWithLock(const QCamera* self, int lock) {
	QCamera::LockStatus _ret = self->lockStatus(static_cast<QCamera::LockType>(lock));
	return static_cast<int>(_ret);
}

void QCamera_SetCaptureMode(QCamera* self, int mode) {
	self->setCaptureMode(static_cast<QCamera::CaptureModes>(mode));
}

void QCamera_Load(QCamera* self) {
	self->load();
}

void QCamera_Unload(QCamera* self) {
	self->unload();
}

void QCamera_Start(QCamera* self) {
	self->start();
}

void QCamera_Stop(QCamera* self) {
	self->stop();
}

void QCamera_SearchAndLock(QCamera* self) {
	self->searchAndLock();
}

void QCamera_Unlock(QCamera* self) {
	self->unlock();
}

void QCamera_SearchAndLockWithLocks(QCamera* self, int locks) {
	self->searchAndLock(static_cast<QCamera::LockTypes>(locks));
}

void QCamera_UnlockWithLocks(QCamera* self, int locks) {
	self->unlock(static_cast<QCamera::LockTypes>(locks));
}

void QCamera_StateChanged(QCamera* self, int state) {
	self->stateChanged(static_cast<QCamera::State>(state));
}

void QCamera_connect_StateChanged(QCamera* self, intptr_t slot) {
	MiqtVirtualQCamera::connect(self, static_cast<void (QCamera::*)(QCamera::State)>(&QCamera::stateChanged), self, [=](QCamera::State state) {
		QCamera::State state_ret = state;
		int sigval1 = static_cast<int>(state_ret);
		miqt_exec_callback_QCamera_StateChanged(slot, sigval1);
	});
}

void QCamera_CaptureModeChanged(QCamera* self, int param1) {
	self->captureModeChanged(static_cast<QCamera::CaptureModes>(param1));
}

void QCamera_connect_CaptureModeChanged(QCamera* self, intptr_t slot) {
	MiqtVirtualQCamera::connect(self, static_cast<void (QCamera::*)(QCamera::CaptureModes)>(&QCamera::captureModeChanged), self, [=](QCamera::CaptureModes param1) {
		QCamera::CaptureModes param1_ret = param1;
		int sigval1 = static_cast<int>(param1_ret);
		miqt_exec_callback_QCamera_CaptureModeChanged(slot, sigval1);
	});
}

void QCamera_StatusChanged(QCamera* self, int status) {
	self->statusChanged(static_cast<QCamera::Status>(status));
}

void QCamera_connect_StatusChanged(QCamera* self, intptr_t slot) {
	MiqtVirtualQCamera::connect(self, static_cast<void (QCamera::*)(QCamera::Status)>(&QCamera::statusChanged), self, [=](QCamera::Status status) {
		QCamera::Status status_ret = status;
		int sigval1 = static_cast<int>(status_ret);
		miqt_exec_callback_QCamera_StatusChanged(slot, sigval1);
	});
}

void QCamera_Locked(QCamera* self) {
	self->locked();
}

void QCamera_connect_Locked(QCamera* self, intptr_t slot) {
	MiqtVirtualQCamera::connect(self, static_cast<void (QCamera::*)()>(&QCamera::locked), self, [=]() {
		miqt_exec_callback_QCamera_Locked(slot);
	});
}

void QCamera_LockFailed(QCamera* self) {
	self->lockFailed();
}

void QCamera_connect_LockFailed(QCamera* self, intptr_t slot) {
	MiqtVirtualQCamera::connect(self, static_cast<void (QCamera::*)()>(&QCamera::lockFailed), self, [=]() {
		miqt_exec_callback_QCamera_LockFailed(slot);
	});
}

void QCamera_LockStatusChanged(QCamera* self, int status, int reason) {
	self->lockStatusChanged(static_cast<QCamera::LockStatus>(status), static_cast<QCamera::LockChangeReason>(reason));
}

void QCamera_connect_LockStatusChanged(QCamera* self, intptr_t slot) {
	MiqtVirtualQCamera::connect(self, static_cast<void (QCamera::*)(QCamera::LockStatus, QCamera::LockChangeReason)>(&QCamera::lockStatusChanged), self, [=](QCamera::LockStatus status, QCamera::LockChangeReason reason) {
		QCamera::LockStatus status_ret = status;
		int sigval1 = static_cast<int>(status_ret);
		QCamera::LockChangeReason reason_ret = reason;
		int sigval2 = static_cast<int>(reason_ret);
		miqt_exec_callback_QCamera_LockStatusChanged(slot, sigval1, sigval2);
	});
}

void QCamera_LockStatusChanged2(QCamera* self, int lock, int status, int reason) {
	self->lockStatusChanged(static_cast<QCamera::LockType>(lock), static_cast<QCamera::LockStatus>(status), static_cast<QCamera::LockChangeReason>(reason));
}

void QCamera_connect_LockStatusChanged2(QCamera* self, intptr_t slot) {
	MiqtVirtualQCamera::connect(self, static_cast<void (QCamera::*)(QCamera::LockType, QCamera::LockStatus, QCamera::LockChangeReason)>(&QCamera::lockStatusChanged), self, [=](QCamera::LockType lock, QCamera::LockStatus status, QCamera::LockChangeReason reason) {
		QCamera::LockType lock_ret = lock;
		int sigval1 = static_cast<int>(lock_ret);
		QCamera::LockStatus status_ret = status;
		int sigval2 = static_cast<int>(status_ret);
		QCamera::LockChangeReason reason_ret = reason;
		int sigval3 = static_cast<int>(reason_ret);
		miqt_exec_callback_QCamera_LockStatusChanged2(slot, sigval1, sigval2, sigval3);
	});
}

void QCamera_ErrorWithQCameraError(QCamera* self, int param1) {
	self->error(static_cast<QCamera::Error>(param1));
}

void QCamera_connect_ErrorWithQCameraError(QCamera* self, intptr_t slot) {
	MiqtVirtualQCamera::connect(self, static_cast<void (QCamera::*)(QCamera::Error)>(&QCamera::error), self, [=](QCamera::Error param1) {
		QCamera::Error param1_ret = param1;
		int sigval1 = static_cast<int>(param1_ret);
		miqt_exec_callback_QCamera_ErrorWithQCameraError(slot, sigval1);
	});
}

void QCamera_ErrorOccurred(QCamera* self, int param1) {
	self->errorOccurred(static_cast<QCamera::Error>(param1));
}

void QCamera_connect_ErrorOccurred(QCamera* self, intptr_t slot) {
	MiqtVirtualQCamera::connect(self, static_cast<void (QCamera::*)(QCamera::Error)>(&QCamera::errorOccurred), self, [=](QCamera::Error param1) {
		QCamera::Error param1_ret = param1;
		int sigval1 = static_cast<int>(param1_ret);
		miqt_exec_callback_QCamera_ErrorOccurred(slot, sigval1);
	});
}

struct miqt_string QCamera_Tr2(const char* s, const char* c) {
	QString _ret = QCamera::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QCamera_Tr3(const char* s, const char* c, int n) {
	QString _ret = QCamera::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QCamera_TrUtf82(const char* s, const char* c) {
	QString _ret = QCamera::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QCamera_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QCamera::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_array /* of QCameraViewfinderSettings* */  QCamera_SupportedViewfinderSettings1(const QCamera* self, QCameraViewfinderSettings* settings) {
	QList<QCameraViewfinderSettings> _ret = self->supportedViewfinderSettings(*settings);
	// Convert QList<> from C++ memory to manually-managed C memory
	QCameraViewfinderSettings** _arr = static_cast<QCameraViewfinderSettings**>(malloc(sizeof(QCameraViewfinderSettings*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QCameraViewfinderSettings(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of QSize* */  QCamera_SupportedViewfinderResolutions1(const QCamera* self, QCameraViewfinderSettings* settings) {
	QList<QSize> _ret = self->supportedViewfinderResolutions(*settings);
	// Convert QList<> from C++ memory to manually-managed C memory
	QSize** _arr = static_cast<QSize**>(malloc(sizeof(QSize*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QSize(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of QCamera__FrameRateRange* */  QCamera_SupportedViewfinderFrameRateRanges1(const QCamera* self, QCameraViewfinderSettings* settings) {
	QList<QCamera::FrameRateRange> _ret = self->supportedViewfinderFrameRateRanges(*settings);
	// Convert QList<> from C++ memory to manually-managed C memory
	QCamera__FrameRateRange** _arr = static_cast<QCamera__FrameRateRange**>(malloc(sizeof(QCamera__FrameRateRange*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QCamera::FrameRateRange(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of int */  QCamera_SupportedViewfinderPixelFormats1(const QCamera* self, QCameraViewfinderSettings* settings) {
	QList<QVideoFrame::PixelFormat> _ret = self->supportedViewfinderPixelFormats(*settings);
	// Convert QList<> from C++ memory to manually-managed C memory
	int* _arr = static_cast<int*>(malloc(sizeof(int) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QVideoFrame::PixelFormat _lv_ret = _ret[i];
		_arr[i] = static_cast<int>(_lv_ret);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

void QCamera_override_virtual_Availability(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCamera*>( (QCamera*)(self) )->handle__Availability = slot;
}

int QCamera_virtualbase_Availability(const void* self) {
	return ( (const MiqtVirtualQCamera*)(self) )->virtualbase_Availability();
}

void QCamera_override_virtual_IsAvailable(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCamera*>( (QCamera*)(self) )->handle__IsAvailable = slot;
}

bool QCamera_virtualbase_IsAvailable(const void* self) {
	return ( (const MiqtVirtualQCamera*)(self) )->virtualbase_IsAvailable();
}

void QCamera_override_virtual_Service(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCamera*>( (QCamera*)(self) )->handle__Service = slot;
}

QMediaService* QCamera_virtualbase_Service(const void* self) {
	return ( (const MiqtVirtualQCamera*)(self) )->virtualbase_Service();
}

void QCamera_override_virtual_Bind(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCamera*>( (QCamera*)(self) )->handle__Bind = slot;
}

bool QCamera_virtualbase_Bind(void* self, QObject* param1) {
	return ( (MiqtVirtualQCamera*)(self) )->virtualbase_Bind(param1);
}

void QCamera_override_virtual_Unbind(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCamera*>( (QCamera*)(self) )->handle__Unbind = slot;
}

void QCamera_virtualbase_Unbind(void* self, QObject* param1) {
	( (MiqtVirtualQCamera*)(self) )->virtualbase_Unbind(param1);
}

void QCamera_Delete(QCamera* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQCamera*>( self );
	} else {
		delete self;
	}
}

QCamera__FrameRateRange* QCamera__FrameRateRange_new() {
	return new QCamera::FrameRateRange();
}

QCamera__FrameRateRange* QCamera__FrameRateRange_new2(double minimum, double maximum) {
	return new QCamera::FrameRateRange(static_cast<qreal>(minimum), static_cast<qreal>(maximum));
}

QCamera__FrameRateRange* QCamera__FrameRateRange_new3(QCamera__FrameRateRange* param1) {
	return new QCamera::FrameRateRange(*param1);
}

void QCamera__FrameRateRange_Delete(QCamera__FrameRateRange* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QCamera::FrameRateRange*>( self );
	} else {
		delete self;
	}
}

