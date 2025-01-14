#include <QAction>
#include <QActionEvent>
#include <QByteArray>
#include <QCloseEvent>
#include <QCompleter>
#include <QContextMenuEvent>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QHideEvent>
#include <QIcon>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QLineEdit>
#include <QMargins>
#include <QMenu>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPoint>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTabletEvent>
#include <QValidator>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>
#include <qlineedit.h>
#include "gen_qlineedit.h"

#ifndef _Bool
#define _Bool bool
#endif
#include "_cgo_export.h"

class MiqtVirtualQLineEdit : public virtual QLineEdit {
public:

	MiqtVirtualQLineEdit(QWidget* parent): QLineEdit(parent) {};
	MiqtVirtualQLineEdit(): QLineEdit() {};
	MiqtVirtualQLineEdit(const QString& param1): QLineEdit(param1) {};
	MiqtVirtualQLineEdit(const QString& param1, QWidget* parent): QLineEdit(param1, parent) {};

	virtual ~MiqtVirtualQLineEdit() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SizeHint = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize sizeHint() const override {
		if (handle__SizeHint == 0) {
			return QLineEdit::sizeHint();
		}
		

		QSize* callback_return_value = miqt_exec_callback_QLineEdit_SizeHint(const_cast<MiqtVirtualQLineEdit*>(this), handle__SizeHint);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSize* virtualbase_SizeHint() const {

		return new QSize(QLineEdit::sizeHint());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MinimumSizeHint = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize minimumSizeHint() const override {
		if (handle__MinimumSizeHint == 0) {
			return QLineEdit::minimumSizeHint();
		}
		

		QSize* callback_return_value = miqt_exec_callback_QLineEdit_MinimumSizeHint(const_cast<MiqtVirtualQLineEdit*>(this), handle__MinimumSizeHint);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSize* virtualbase_MinimumSizeHint() const {

		return new QSize(QLineEdit::minimumSizeHint());

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MousePressEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mousePressEvent(QMouseEvent* param1) override {
		if (handle__MousePressEvent == 0) {
			QLineEdit::mousePressEvent(param1);
			return;
		}
		
		QMouseEvent* sigval1 = param1;

		miqt_exec_callback_QLineEdit_MousePressEvent(this, handle__MousePressEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MousePressEvent(QMouseEvent* param1) {

		QLineEdit::mousePressEvent(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MouseMoveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mouseMoveEvent(QMouseEvent* param1) override {
		if (handle__MouseMoveEvent == 0) {
			QLineEdit::mouseMoveEvent(param1);
			return;
		}
		
		QMouseEvent* sigval1 = param1;

		miqt_exec_callback_QLineEdit_MouseMoveEvent(this, handle__MouseMoveEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MouseMoveEvent(QMouseEvent* param1) {

		QLineEdit::mouseMoveEvent(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MouseReleaseEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mouseReleaseEvent(QMouseEvent* param1) override {
		if (handle__MouseReleaseEvent == 0) {
			QLineEdit::mouseReleaseEvent(param1);
			return;
		}
		
		QMouseEvent* sigval1 = param1;

		miqt_exec_callback_QLineEdit_MouseReleaseEvent(this, handle__MouseReleaseEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MouseReleaseEvent(QMouseEvent* param1) {

		QLineEdit::mouseReleaseEvent(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MouseDoubleClickEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mouseDoubleClickEvent(QMouseEvent* param1) override {
		if (handle__MouseDoubleClickEvent == 0) {
			QLineEdit::mouseDoubleClickEvent(param1);
			return;
		}
		
		QMouseEvent* sigval1 = param1;

		miqt_exec_callback_QLineEdit_MouseDoubleClickEvent(this, handle__MouseDoubleClickEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MouseDoubleClickEvent(QMouseEvent* param1) {

		QLineEdit::mouseDoubleClickEvent(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__KeyPressEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void keyPressEvent(QKeyEvent* param1) override {
		if (handle__KeyPressEvent == 0) {
			QLineEdit::keyPressEvent(param1);
			return;
		}
		
		QKeyEvent* sigval1 = param1;

		miqt_exec_callback_QLineEdit_KeyPressEvent(this, handle__KeyPressEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_KeyPressEvent(QKeyEvent* param1) {

		QLineEdit::keyPressEvent(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__FocusInEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void focusInEvent(QFocusEvent* param1) override {
		if (handle__FocusInEvent == 0) {
			QLineEdit::focusInEvent(param1);
			return;
		}
		
		QFocusEvent* sigval1 = param1;

		miqt_exec_callback_QLineEdit_FocusInEvent(this, handle__FocusInEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_FocusInEvent(QFocusEvent* param1) {

		QLineEdit::focusInEvent(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__FocusOutEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void focusOutEvent(QFocusEvent* param1) override {
		if (handle__FocusOutEvent == 0) {
			QLineEdit::focusOutEvent(param1);
			return;
		}
		
		QFocusEvent* sigval1 = param1;

		miqt_exec_callback_QLineEdit_FocusOutEvent(this, handle__FocusOutEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_FocusOutEvent(QFocusEvent* param1) {

		QLineEdit::focusOutEvent(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__PaintEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void paintEvent(QPaintEvent* param1) override {
		if (handle__PaintEvent == 0) {
			QLineEdit::paintEvent(param1);
			return;
		}
		
		QPaintEvent* sigval1 = param1;

		miqt_exec_callback_QLineEdit_PaintEvent(this, handle__PaintEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_PaintEvent(QPaintEvent* param1) {

		QLineEdit::paintEvent(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DragEnterEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void dragEnterEvent(QDragEnterEvent* param1) override {
		if (handle__DragEnterEvent == 0) {
			QLineEdit::dragEnterEvent(param1);
			return;
		}
		
		QDragEnterEvent* sigval1 = param1;

		miqt_exec_callback_QLineEdit_DragEnterEvent(this, handle__DragEnterEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DragEnterEvent(QDragEnterEvent* param1) {

		QLineEdit::dragEnterEvent(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DragMoveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void dragMoveEvent(QDragMoveEvent* e) override {
		if (handle__DragMoveEvent == 0) {
			QLineEdit::dragMoveEvent(e);
			return;
		}
		
		QDragMoveEvent* sigval1 = e;

		miqt_exec_callback_QLineEdit_DragMoveEvent(this, handle__DragMoveEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DragMoveEvent(QDragMoveEvent* e) {

		QLineEdit::dragMoveEvent(e);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DragLeaveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void dragLeaveEvent(QDragLeaveEvent* e) override {
		if (handle__DragLeaveEvent == 0) {
			QLineEdit::dragLeaveEvent(e);
			return;
		}
		
		QDragLeaveEvent* sigval1 = e;

		miqt_exec_callback_QLineEdit_DragLeaveEvent(this, handle__DragLeaveEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DragLeaveEvent(QDragLeaveEvent* e) {

		QLineEdit::dragLeaveEvent(e);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DropEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void dropEvent(QDropEvent* param1) override {
		if (handle__DropEvent == 0) {
			QLineEdit::dropEvent(param1);
			return;
		}
		
		QDropEvent* sigval1 = param1;

		miqt_exec_callback_QLineEdit_DropEvent(this, handle__DropEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DropEvent(QDropEvent* param1) {

		QLineEdit::dropEvent(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ChangeEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void changeEvent(QEvent* param1) override {
		if (handle__ChangeEvent == 0) {
			QLineEdit::changeEvent(param1);
			return;
		}
		
		QEvent* sigval1 = param1;

		miqt_exec_callback_QLineEdit_ChangeEvent(this, handle__ChangeEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ChangeEvent(QEvent* param1) {

		QLineEdit::changeEvent(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ContextMenuEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void contextMenuEvent(QContextMenuEvent* param1) override {
		if (handle__ContextMenuEvent == 0) {
			QLineEdit::contextMenuEvent(param1);
			return;
		}
		
		QContextMenuEvent* sigval1 = param1;

		miqt_exec_callback_QLineEdit_ContextMenuEvent(this, handle__ContextMenuEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ContextMenuEvent(QContextMenuEvent* param1) {

		QLineEdit::contextMenuEvent(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__InputMethodEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void inputMethodEvent(QInputMethodEvent* param1) override {
		if (handle__InputMethodEvent == 0) {
			QLineEdit::inputMethodEvent(param1);
			return;
		}
		
		QInputMethodEvent* sigval1 = param1;

		miqt_exec_callback_QLineEdit_InputMethodEvent(this, handle__InputMethodEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_InputMethodEvent(QInputMethodEvent* param1) {

		QLineEdit::inputMethodEvent(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__InputMethodQuery = 0;

	// Subclass to allow providing a Go implementation
	virtual QVariant inputMethodQuery(Qt::InputMethodQuery param1) const override {
		if (handle__InputMethodQuery == 0) {
			return QLineEdit::inputMethodQuery(param1);
		}
		
		Qt::InputMethodQuery param1_ret = param1;
		int sigval1 = static_cast<int>(param1_ret);

		QVariant* callback_return_value = miqt_exec_callback_QLineEdit_InputMethodQuery(const_cast<MiqtVirtualQLineEdit*>(this), handle__InputMethodQuery, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QVariant* virtualbase_InputMethodQuery(int param1) const {

		return new QVariant(QLineEdit::inputMethodQuery(static_cast<Qt::InputMethodQuery>(param1)));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Event = 0;

	// Subclass to allow providing a Go implementation
	virtual bool event(QEvent* param1) override {
		if (handle__Event == 0) {
			return QLineEdit::event(param1);
		}
		
		QEvent* sigval1 = param1;

		bool callback_return_value = miqt_exec_callback_QLineEdit_Event(this, handle__Event, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_Event(QEvent* param1) {

		return QLineEdit::event(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DevType = 0;

	// Subclass to allow providing a Go implementation
	virtual int devType() const override {
		if (handle__DevType == 0) {
			return QLineEdit::devType();
		}
		

		int callback_return_value = miqt_exec_callback_QLineEdit_DevType(const_cast<MiqtVirtualQLineEdit*>(this), handle__DevType);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_DevType() const {

		return QLineEdit::devType();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetVisible = 0;

	// Subclass to allow providing a Go implementation
	virtual void setVisible(bool visible) override {
		if (handle__SetVisible == 0) {
			QLineEdit::setVisible(visible);
			return;
		}
		
		bool sigval1 = visible;

		miqt_exec_callback_QLineEdit_SetVisible(this, handle__SetVisible, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetVisible(bool visible) {

		QLineEdit::setVisible(visible);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__HeightForWidth = 0;

	// Subclass to allow providing a Go implementation
	virtual int heightForWidth(int param1) const override {
		if (handle__HeightForWidth == 0) {
			return QLineEdit::heightForWidth(param1);
		}
		
		int sigval1 = param1;

		int callback_return_value = miqt_exec_callback_QLineEdit_HeightForWidth(const_cast<MiqtVirtualQLineEdit*>(this), handle__HeightForWidth, sigval1);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_HeightForWidth(int param1) const {

		return QLineEdit::heightForWidth(static_cast<int>(param1));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__HasHeightForWidth = 0;

	// Subclass to allow providing a Go implementation
	virtual bool hasHeightForWidth() const override {
		if (handle__HasHeightForWidth == 0) {
			return QLineEdit::hasHeightForWidth();
		}
		

		bool callback_return_value = miqt_exec_callback_QLineEdit_HasHeightForWidth(const_cast<MiqtVirtualQLineEdit*>(this), handle__HasHeightForWidth);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_HasHeightForWidth() const {

		return QLineEdit::hasHeightForWidth();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__PaintEngine = 0;

	// Subclass to allow providing a Go implementation
	virtual QPaintEngine* paintEngine() const override {
		if (handle__PaintEngine == 0) {
			return QLineEdit::paintEngine();
		}
		

		QPaintEngine* callback_return_value = miqt_exec_callback_QLineEdit_PaintEngine(const_cast<MiqtVirtualQLineEdit*>(this), handle__PaintEngine);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QPaintEngine* virtualbase_PaintEngine() const {

		return QLineEdit::paintEngine();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__WheelEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void wheelEvent(QWheelEvent* event) override {
		if (handle__WheelEvent == 0) {
			QLineEdit::wheelEvent(event);
			return;
		}
		
		QWheelEvent* sigval1 = event;

		miqt_exec_callback_QLineEdit_WheelEvent(this, handle__WheelEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_WheelEvent(QWheelEvent* event) {

		QLineEdit::wheelEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__KeyReleaseEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void keyReleaseEvent(QKeyEvent* event) override {
		if (handle__KeyReleaseEvent == 0) {
			QLineEdit::keyReleaseEvent(event);
			return;
		}
		
		QKeyEvent* sigval1 = event;

		miqt_exec_callback_QLineEdit_KeyReleaseEvent(this, handle__KeyReleaseEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_KeyReleaseEvent(QKeyEvent* event) {

		QLineEdit::keyReleaseEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__EnterEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void enterEvent(QEvent* event) override {
		if (handle__EnterEvent == 0) {
			QLineEdit::enterEvent(event);
			return;
		}
		
		QEvent* sigval1 = event;

		miqt_exec_callback_QLineEdit_EnterEvent(this, handle__EnterEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_EnterEvent(QEvent* event) {

		QLineEdit::enterEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__LeaveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void leaveEvent(QEvent* event) override {
		if (handle__LeaveEvent == 0) {
			QLineEdit::leaveEvent(event);
			return;
		}
		
		QEvent* sigval1 = event;

		miqt_exec_callback_QLineEdit_LeaveEvent(this, handle__LeaveEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_LeaveEvent(QEvent* event) {

		QLineEdit::leaveEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MoveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void moveEvent(QMoveEvent* event) override {
		if (handle__MoveEvent == 0) {
			QLineEdit::moveEvent(event);
			return;
		}
		
		QMoveEvent* sigval1 = event;

		miqt_exec_callback_QLineEdit_MoveEvent(this, handle__MoveEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_MoveEvent(QMoveEvent* event) {

		QLineEdit::moveEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ResizeEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void resizeEvent(QResizeEvent* event) override {
		if (handle__ResizeEvent == 0) {
			QLineEdit::resizeEvent(event);
			return;
		}
		
		QResizeEvent* sigval1 = event;

		miqt_exec_callback_QLineEdit_ResizeEvent(this, handle__ResizeEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ResizeEvent(QResizeEvent* event) {

		QLineEdit::resizeEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__CloseEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void closeEvent(QCloseEvent* event) override {
		if (handle__CloseEvent == 0) {
			QLineEdit::closeEvent(event);
			return;
		}
		
		QCloseEvent* sigval1 = event;

		miqt_exec_callback_QLineEdit_CloseEvent(this, handle__CloseEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_CloseEvent(QCloseEvent* event) {

		QLineEdit::closeEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__TabletEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void tabletEvent(QTabletEvent* event) override {
		if (handle__TabletEvent == 0) {
			QLineEdit::tabletEvent(event);
			return;
		}
		
		QTabletEvent* sigval1 = event;

		miqt_exec_callback_QLineEdit_TabletEvent(this, handle__TabletEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_TabletEvent(QTabletEvent* event) {

		QLineEdit::tabletEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ActionEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void actionEvent(QActionEvent* event) override {
		if (handle__ActionEvent == 0) {
			QLineEdit::actionEvent(event);
			return;
		}
		
		QActionEvent* sigval1 = event;

		miqt_exec_callback_QLineEdit_ActionEvent(this, handle__ActionEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ActionEvent(QActionEvent* event) {

		QLineEdit::actionEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ShowEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void showEvent(QShowEvent* event) override {
		if (handle__ShowEvent == 0) {
			QLineEdit::showEvent(event);
			return;
		}
		
		QShowEvent* sigval1 = event;

		miqt_exec_callback_QLineEdit_ShowEvent(this, handle__ShowEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ShowEvent(QShowEvent* event) {

		QLineEdit::showEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__HideEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void hideEvent(QHideEvent* event) override {
		if (handle__HideEvent == 0) {
			QLineEdit::hideEvent(event);
			return;
		}
		
		QHideEvent* sigval1 = event;

		miqt_exec_callback_QLineEdit_HideEvent(this, handle__HideEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_HideEvent(QHideEvent* event) {

		QLineEdit::hideEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__NativeEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual bool nativeEvent(const QByteArray& eventType, void* message, long* result) override {
		if (handle__NativeEvent == 0) {
			return QLineEdit::nativeEvent(eventType, message, result);
		}
		
		const QByteArray eventType_qb = eventType;
		struct miqt_string eventType_ms;
		eventType_ms.len = eventType_qb.length();
		eventType_ms.data = static_cast<char*>(malloc(eventType_ms.len));
		memcpy(eventType_ms.data, eventType_qb.data(), eventType_ms.len);
		struct miqt_string sigval1 = eventType_ms;
		void* sigval2 = message;
		long* sigval3 = result;

		bool callback_return_value = miqt_exec_callback_QLineEdit_NativeEvent(this, handle__NativeEvent, sigval1, sigval2, sigval3);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_NativeEvent(struct miqt_string eventType, void* message, long* result) {
		QByteArray eventType_QByteArray(eventType.data, eventType.len);

		return QLineEdit::nativeEvent(eventType_QByteArray, message, static_cast<long*>(result));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metric = 0;

	// Subclass to allow providing a Go implementation
	virtual int metric(QPaintDevice::PaintDeviceMetric param1) const override {
		if (handle__Metric == 0) {
			return QLineEdit::metric(param1);
		}
		
		QPaintDevice::PaintDeviceMetric param1_ret = param1;
		int sigval1 = static_cast<int>(param1_ret);

		int callback_return_value = miqt_exec_callback_QLineEdit_Metric(const_cast<MiqtVirtualQLineEdit*>(this), handle__Metric, sigval1);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_Metric(int param1) const {

		return QLineEdit::metric(static_cast<QPaintDevice::PaintDeviceMetric>(param1));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__InitPainter = 0;

	// Subclass to allow providing a Go implementation
	virtual void initPainter(QPainter* painter) const override {
		if (handle__InitPainter == 0) {
			QLineEdit::initPainter(painter);
			return;
		}
		
		QPainter* sigval1 = painter;

		miqt_exec_callback_QLineEdit_InitPainter(const_cast<MiqtVirtualQLineEdit*>(this), handle__InitPainter, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_InitPainter(QPainter* painter) const {

		QLineEdit::initPainter(painter);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Redirected = 0;

	// Subclass to allow providing a Go implementation
	virtual QPaintDevice* redirected(QPoint* offset) const override {
		if (handle__Redirected == 0) {
			return QLineEdit::redirected(offset);
		}
		
		QPoint* sigval1 = offset;

		QPaintDevice* callback_return_value = miqt_exec_callback_QLineEdit_Redirected(const_cast<MiqtVirtualQLineEdit*>(this), handle__Redirected, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QPaintDevice* virtualbase_Redirected(QPoint* offset) const {

		return QLineEdit::redirected(offset);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SharedPainter = 0;

	// Subclass to allow providing a Go implementation
	virtual QPainter* sharedPainter() const override {
		if (handle__SharedPainter == 0) {
			return QLineEdit::sharedPainter();
		}
		

		QPainter* callback_return_value = miqt_exec_callback_QLineEdit_SharedPainter(const_cast<MiqtVirtualQLineEdit*>(this), handle__SharedPainter);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QPainter* virtualbase_SharedPainter() const {

		return QLineEdit::sharedPainter();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__FocusNextPrevChild = 0;

	// Subclass to allow providing a Go implementation
	virtual bool focusNextPrevChild(bool next) override {
		if (handle__FocusNextPrevChild == 0) {
			return QLineEdit::focusNextPrevChild(next);
		}
		
		bool sigval1 = next;

		bool callback_return_value = miqt_exec_callback_QLineEdit_FocusNextPrevChild(this, handle__FocusNextPrevChild, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_FocusNextPrevChild(bool next) {

		return QLineEdit::focusNextPrevChild(next);

	}

};

QLineEdit* QLineEdit_new(QWidget* parent) {
	return new MiqtVirtualQLineEdit(parent);
}

QLineEdit* QLineEdit_new2() {
	return new MiqtVirtualQLineEdit();
}

QLineEdit* QLineEdit_new3(struct miqt_string param1) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	return new MiqtVirtualQLineEdit(param1_QString);
}

QLineEdit* QLineEdit_new4(struct miqt_string param1, QWidget* parent) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	return new MiqtVirtualQLineEdit(param1_QString, parent);
}

void QLineEdit_virtbase(QLineEdit* src, QWidget** outptr_QWidget) {
	*outptr_QWidget = static_cast<QWidget*>(src);
}

QMetaObject* QLineEdit_MetaObject(const QLineEdit* self) {
	return (QMetaObject*) self->metaObject();
}

void* QLineEdit_Metacast(QLineEdit* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QLineEdit_Tr(const char* s) {
	QString _ret = QLineEdit::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLineEdit_TrUtf8(const char* s) {
	QString _ret = QLineEdit::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLineEdit_Text(const QLineEdit* self) {
	QString _ret = self->text();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLineEdit_DisplayText(const QLineEdit* self) {
	QString _ret = self->displayText();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLineEdit_PlaceholderText(const QLineEdit* self) {
	QString _ret = self->placeholderText();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QLineEdit_SetPlaceholderText(QLineEdit* self, struct miqt_string placeholderText) {
	QString placeholderText_QString = QString::fromUtf8(placeholderText.data, placeholderText.len);
	self->setPlaceholderText(placeholderText_QString);
}

int QLineEdit_MaxLength(const QLineEdit* self) {
	return self->maxLength();
}

void QLineEdit_SetMaxLength(QLineEdit* self, int maxLength) {
	self->setMaxLength(static_cast<int>(maxLength));
}

void QLineEdit_SetFrame(QLineEdit* self, bool frame) {
	self->setFrame(frame);
}

bool QLineEdit_HasFrame(const QLineEdit* self) {
	return self->hasFrame();
}

void QLineEdit_SetClearButtonEnabled(QLineEdit* self, bool enable) {
	self->setClearButtonEnabled(enable);
}

bool QLineEdit_IsClearButtonEnabled(const QLineEdit* self) {
	return self->isClearButtonEnabled();
}

int QLineEdit_EchoMode(const QLineEdit* self) {
	QLineEdit::EchoMode _ret = self->echoMode();
	return static_cast<int>(_ret);
}

void QLineEdit_SetEchoMode(QLineEdit* self, int echoMode) {
	self->setEchoMode(static_cast<QLineEdit::EchoMode>(echoMode));
}

bool QLineEdit_IsReadOnly(const QLineEdit* self) {
	return self->isReadOnly();
}

void QLineEdit_SetReadOnly(QLineEdit* self, bool readOnly) {
	self->setReadOnly(readOnly);
}

void QLineEdit_SetValidator(QLineEdit* self, QValidator* validator) {
	self->setValidator(validator);
}

QValidator* QLineEdit_Validator(const QLineEdit* self) {
	return (QValidator*) self->validator();
}

void QLineEdit_SetCompleter(QLineEdit* self, QCompleter* completer) {
	self->setCompleter(completer);
}

QCompleter* QLineEdit_Completer(const QLineEdit* self) {
	return self->completer();
}

QSize* QLineEdit_SizeHint(const QLineEdit* self) {
	return new QSize(self->sizeHint());
}

QSize* QLineEdit_MinimumSizeHint(const QLineEdit* self) {
	return new QSize(self->minimumSizeHint());
}

int QLineEdit_CursorPosition(const QLineEdit* self) {
	return self->cursorPosition();
}

void QLineEdit_SetCursorPosition(QLineEdit* self, int cursorPosition) {
	self->setCursorPosition(static_cast<int>(cursorPosition));
}

int QLineEdit_CursorPositionAt(QLineEdit* self, QPoint* pos) {
	return self->cursorPositionAt(*pos);
}

void QLineEdit_SetAlignment(QLineEdit* self, int flag) {
	self->setAlignment(static_cast<Qt::Alignment>(flag));
}

int QLineEdit_Alignment(const QLineEdit* self) {
	Qt::Alignment _ret = self->alignment();
	return static_cast<int>(_ret);
}

void QLineEdit_CursorForward(QLineEdit* self, bool mark) {
	self->cursorForward(mark);
}

void QLineEdit_CursorBackward(QLineEdit* self, bool mark) {
	self->cursorBackward(mark);
}

void QLineEdit_CursorWordForward(QLineEdit* self, bool mark) {
	self->cursorWordForward(mark);
}

void QLineEdit_CursorWordBackward(QLineEdit* self, bool mark) {
	self->cursorWordBackward(mark);
}

void QLineEdit_Backspace(QLineEdit* self) {
	self->backspace();
}

void QLineEdit_Del(QLineEdit* self) {
	self->del();
}

void QLineEdit_Home(QLineEdit* self, bool mark) {
	self->home(mark);
}

void QLineEdit_End(QLineEdit* self, bool mark) {
	self->end(mark);
}

bool QLineEdit_IsModified(const QLineEdit* self) {
	return self->isModified();
}

void QLineEdit_SetModified(QLineEdit* self, bool modified) {
	self->setModified(modified);
}

void QLineEdit_SetSelection(QLineEdit* self, int param1, int param2) {
	self->setSelection(static_cast<int>(param1), static_cast<int>(param2));
}

bool QLineEdit_HasSelectedText(const QLineEdit* self) {
	return self->hasSelectedText();
}

struct miqt_string QLineEdit_SelectedText(const QLineEdit* self) {
	QString _ret = self->selectedText();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QLineEdit_SelectionStart(const QLineEdit* self) {
	return self->selectionStart();
}

int QLineEdit_SelectionEnd(const QLineEdit* self) {
	return self->selectionEnd();
}

int QLineEdit_SelectionLength(const QLineEdit* self) {
	return self->selectionLength();
}

bool QLineEdit_IsUndoAvailable(const QLineEdit* self) {
	return self->isUndoAvailable();
}

bool QLineEdit_IsRedoAvailable(const QLineEdit* self) {
	return self->isRedoAvailable();
}

void QLineEdit_SetDragEnabled(QLineEdit* self, bool b) {
	self->setDragEnabled(b);
}

bool QLineEdit_DragEnabled(const QLineEdit* self) {
	return self->dragEnabled();
}

void QLineEdit_SetCursorMoveStyle(QLineEdit* self, int style) {
	self->setCursorMoveStyle(static_cast<Qt::CursorMoveStyle>(style));
}

int QLineEdit_CursorMoveStyle(const QLineEdit* self) {
	Qt::CursorMoveStyle _ret = self->cursorMoveStyle();
	return static_cast<int>(_ret);
}

struct miqt_string QLineEdit_InputMask(const QLineEdit* self) {
	QString _ret = self->inputMask();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QLineEdit_SetInputMask(QLineEdit* self, struct miqt_string inputMask) {
	QString inputMask_QString = QString::fromUtf8(inputMask.data, inputMask.len);
	self->setInputMask(inputMask_QString);
}

bool QLineEdit_HasAcceptableInput(const QLineEdit* self) {
	return self->hasAcceptableInput();
}

void QLineEdit_SetTextMargins(QLineEdit* self, int left, int top, int right, int bottom) {
	self->setTextMargins(static_cast<int>(left), static_cast<int>(top), static_cast<int>(right), static_cast<int>(bottom));
}

void QLineEdit_SetTextMarginsWithMargins(QLineEdit* self, QMargins* margins) {
	self->setTextMargins(*margins);
}

void QLineEdit_GetTextMargins(const QLineEdit* self, int* left, int* top, int* right, int* bottom) {
	self->getTextMargins(static_cast<int*>(left), static_cast<int*>(top), static_cast<int*>(right), static_cast<int*>(bottom));
}

QMargins* QLineEdit_TextMargins(const QLineEdit* self) {
	return new QMargins(self->textMargins());
}

void QLineEdit_AddAction(QLineEdit* self, QAction* action, int position) {
	self->addAction(action, static_cast<QLineEdit::ActionPosition>(position));
}

QAction* QLineEdit_AddAction2(QLineEdit* self, QIcon* icon, int position) {
	return self->addAction(*icon, static_cast<QLineEdit::ActionPosition>(position));
}

void QLineEdit_SetText(QLineEdit* self, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->setText(text_QString);
}

void QLineEdit_Clear(QLineEdit* self) {
	self->clear();
}

void QLineEdit_SelectAll(QLineEdit* self) {
	self->selectAll();
}

void QLineEdit_Undo(QLineEdit* self) {
	self->undo();
}

void QLineEdit_Redo(QLineEdit* self) {
	self->redo();
}

void QLineEdit_Cut(QLineEdit* self) {
	self->cut();
}

void QLineEdit_Copy(const QLineEdit* self) {
	self->copy();
}

void QLineEdit_Paste(QLineEdit* self) {
	self->paste();
}

void QLineEdit_Deselect(QLineEdit* self) {
	self->deselect();
}

void QLineEdit_Insert(QLineEdit* self, struct miqt_string param1) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	self->insert(param1_QString);
}

QMenu* QLineEdit_CreateStandardContextMenu(QLineEdit* self) {
	return self->createStandardContextMenu();
}

void QLineEdit_TextChanged(QLineEdit* self, struct miqt_string param1) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	self->textChanged(param1_QString);
}

void QLineEdit_connect_TextChanged(QLineEdit* self, intptr_t slot) {
	MiqtVirtualQLineEdit::connect(self, static_cast<void (QLineEdit::*)(const QString&)>(&QLineEdit::textChanged), self, [=](const QString& param1) {
		const QString param1_ret = param1;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray param1_b = param1_ret.toUtf8();
		struct miqt_string param1_ms;
		param1_ms.len = param1_b.length();
		param1_ms.data = static_cast<char*>(malloc(param1_ms.len));
		memcpy(param1_ms.data, param1_b.data(), param1_ms.len);
		struct miqt_string sigval1 = param1_ms;
		miqt_exec_callback_QLineEdit_TextChanged(slot, sigval1);
	});
}

void QLineEdit_TextEdited(QLineEdit* self, struct miqt_string param1) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	self->textEdited(param1_QString);
}

void QLineEdit_connect_TextEdited(QLineEdit* self, intptr_t slot) {
	MiqtVirtualQLineEdit::connect(self, static_cast<void (QLineEdit::*)(const QString&)>(&QLineEdit::textEdited), self, [=](const QString& param1) {
		const QString param1_ret = param1;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray param1_b = param1_ret.toUtf8();
		struct miqt_string param1_ms;
		param1_ms.len = param1_b.length();
		param1_ms.data = static_cast<char*>(malloc(param1_ms.len));
		memcpy(param1_ms.data, param1_b.data(), param1_ms.len);
		struct miqt_string sigval1 = param1_ms;
		miqt_exec_callback_QLineEdit_TextEdited(slot, sigval1);
	});
}

void QLineEdit_CursorPositionChanged(QLineEdit* self, int param1, int param2) {
	self->cursorPositionChanged(static_cast<int>(param1), static_cast<int>(param2));
}

void QLineEdit_connect_CursorPositionChanged(QLineEdit* self, intptr_t slot) {
	MiqtVirtualQLineEdit::connect(self, static_cast<void (QLineEdit::*)(int, int)>(&QLineEdit::cursorPositionChanged), self, [=](int param1, int param2) {
		int sigval1 = param1;
		int sigval2 = param2;
		miqt_exec_callback_QLineEdit_CursorPositionChanged(slot, sigval1, sigval2);
	});
}

void QLineEdit_ReturnPressed(QLineEdit* self) {
	self->returnPressed();
}

void QLineEdit_connect_ReturnPressed(QLineEdit* self, intptr_t slot) {
	MiqtVirtualQLineEdit::connect(self, static_cast<void (QLineEdit::*)()>(&QLineEdit::returnPressed), self, [=]() {
		miqt_exec_callback_QLineEdit_ReturnPressed(slot);
	});
}

void QLineEdit_EditingFinished(QLineEdit* self) {
	self->editingFinished();
}

void QLineEdit_connect_EditingFinished(QLineEdit* self, intptr_t slot) {
	MiqtVirtualQLineEdit::connect(self, static_cast<void (QLineEdit::*)()>(&QLineEdit::editingFinished), self, [=]() {
		miqt_exec_callback_QLineEdit_EditingFinished(slot);
	});
}

void QLineEdit_SelectionChanged(QLineEdit* self) {
	self->selectionChanged();
}

void QLineEdit_connect_SelectionChanged(QLineEdit* self, intptr_t slot) {
	MiqtVirtualQLineEdit::connect(self, static_cast<void (QLineEdit::*)()>(&QLineEdit::selectionChanged), self, [=]() {
		miqt_exec_callback_QLineEdit_SelectionChanged(slot);
	});
}

void QLineEdit_InputRejected(QLineEdit* self) {
	self->inputRejected();
}

void QLineEdit_connect_InputRejected(QLineEdit* self, intptr_t slot) {
	MiqtVirtualQLineEdit::connect(self, static_cast<void (QLineEdit::*)()>(&QLineEdit::inputRejected), self, [=]() {
		miqt_exec_callback_QLineEdit_InputRejected(slot);
	});
}

QVariant* QLineEdit_InputMethodQuery(const QLineEdit* self, int param1) {
	return new QVariant(self->inputMethodQuery(static_cast<Qt::InputMethodQuery>(param1)));
}

QVariant* QLineEdit_InputMethodQuery2(const QLineEdit* self, int property, QVariant* argument) {
	return new QVariant(self->inputMethodQuery(static_cast<Qt::InputMethodQuery>(property), *argument));
}

bool QLineEdit_Event(QLineEdit* self, QEvent* param1) {
	return self->event(param1);
}

struct miqt_string QLineEdit_Tr2(const char* s, const char* c) {
	QString _ret = QLineEdit::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLineEdit_Tr3(const char* s, const char* c, int n) {
	QString _ret = QLineEdit::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLineEdit_TrUtf82(const char* s, const char* c) {
	QString _ret = QLineEdit::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLineEdit_TrUtf83(const char* s, const char* c, int n) {
	QString _ret = QLineEdit::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QLineEdit_CursorForward2(QLineEdit* self, bool mark, int steps) {
	self->cursorForward(mark, static_cast<int>(steps));
}

void QLineEdit_CursorBackward2(QLineEdit* self, bool mark, int steps) {
	self->cursorBackward(mark, static_cast<int>(steps));
}

void QLineEdit_override_virtual_SizeHint(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__SizeHint = slot;
}

QSize* QLineEdit_virtualbase_SizeHint(const void* self) {
	return ( (const MiqtVirtualQLineEdit*)(self) )->virtualbase_SizeHint();
}

void QLineEdit_override_virtual_MinimumSizeHint(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__MinimumSizeHint = slot;
}

QSize* QLineEdit_virtualbase_MinimumSizeHint(const void* self) {
	return ( (const MiqtVirtualQLineEdit*)(self) )->virtualbase_MinimumSizeHint();
}

void QLineEdit_override_virtual_MousePressEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__MousePressEvent = slot;
}

void QLineEdit_virtualbase_MousePressEvent(void* self, QMouseEvent* param1) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_MousePressEvent(param1);
}

void QLineEdit_override_virtual_MouseMoveEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__MouseMoveEvent = slot;
}

void QLineEdit_virtualbase_MouseMoveEvent(void* self, QMouseEvent* param1) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_MouseMoveEvent(param1);
}

void QLineEdit_override_virtual_MouseReleaseEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__MouseReleaseEvent = slot;
}

void QLineEdit_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* param1) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_MouseReleaseEvent(param1);
}

void QLineEdit_override_virtual_MouseDoubleClickEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__MouseDoubleClickEvent = slot;
}

void QLineEdit_virtualbase_MouseDoubleClickEvent(void* self, QMouseEvent* param1) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_MouseDoubleClickEvent(param1);
}

void QLineEdit_override_virtual_KeyPressEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__KeyPressEvent = slot;
}

void QLineEdit_virtualbase_KeyPressEvent(void* self, QKeyEvent* param1) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_KeyPressEvent(param1);
}

void QLineEdit_override_virtual_FocusInEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__FocusInEvent = slot;
}

void QLineEdit_virtualbase_FocusInEvent(void* self, QFocusEvent* param1) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_FocusInEvent(param1);
}

void QLineEdit_override_virtual_FocusOutEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__FocusOutEvent = slot;
}

void QLineEdit_virtualbase_FocusOutEvent(void* self, QFocusEvent* param1) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_FocusOutEvent(param1);
}

void QLineEdit_override_virtual_PaintEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__PaintEvent = slot;
}

void QLineEdit_virtualbase_PaintEvent(void* self, QPaintEvent* param1) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_PaintEvent(param1);
}

void QLineEdit_override_virtual_DragEnterEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__DragEnterEvent = slot;
}

void QLineEdit_virtualbase_DragEnterEvent(void* self, QDragEnterEvent* param1) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_DragEnterEvent(param1);
}

void QLineEdit_override_virtual_DragMoveEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__DragMoveEvent = slot;
}

void QLineEdit_virtualbase_DragMoveEvent(void* self, QDragMoveEvent* e) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_DragMoveEvent(e);
}

void QLineEdit_override_virtual_DragLeaveEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__DragLeaveEvent = slot;
}

void QLineEdit_virtualbase_DragLeaveEvent(void* self, QDragLeaveEvent* e) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_DragLeaveEvent(e);
}

void QLineEdit_override_virtual_DropEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__DropEvent = slot;
}

void QLineEdit_virtualbase_DropEvent(void* self, QDropEvent* param1) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_DropEvent(param1);
}

void QLineEdit_override_virtual_ChangeEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__ChangeEvent = slot;
}

void QLineEdit_virtualbase_ChangeEvent(void* self, QEvent* param1) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_ChangeEvent(param1);
}

void QLineEdit_override_virtual_ContextMenuEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__ContextMenuEvent = slot;
}

void QLineEdit_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* param1) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_ContextMenuEvent(param1);
}

void QLineEdit_override_virtual_InputMethodEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__InputMethodEvent = slot;
}

void QLineEdit_virtualbase_InputMethodEvent(void* self, QInputMethodEvent* param1) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_InputMethodEvent(param1);
}

void QLineEdit_override_virtual_InputMethodQuery(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__InputMethodQuery = slot;
}

QVariant* QLineEdit_virtualbase_InputMethodQuery(const void* self, int param1) {
	return ( (const MiqtVirtualQLineEdit*)(self) )->virtualbase_InputMethodQuery(param1);
}

void QLineEdit_override_virtual_Event(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__Event = slot;
}

bool QLineEdit_virtualbase_Event(void* self, QEvent* param1) {
	return ( (MiqtVirtualQLineEdit*)(self) )->virtualbase_Event(param1);
}

void QLineEdit_override_virtual_DevType(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__DevType = slot;
}

int QLineEdit_virtualbase_DevType(const void* self) {
	return ( (const MiqtVirtualQLineEdit*)(self) )->virtualbase_DevType();
}

void QLineEdit_override_virtual_SetVisible(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__SetVisible = slot;
}

void QLineEdit_virtualbase_SetVisible(void* self, bool visible) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_SetVisible(visible);
}

void QLineEdit_override_virtual_HeightForWidth(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__HeightForWidth = slot;
}

int QLineEdit_virtualbase_HeightForWidth(const void* self, int param1) {
	return ( (const MiqtVirtualQLineEdit*)(self) )->virtualbase_HeightForWidth(param1);
}

void QLineEdit_override_virtual_HasHeightForWidth(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__HasHeightForWidth = slot;
}

bool QLineEdit_virtualbase_HasHeightForWidth(const void* self) {
	return ( (const MiqtVirtualQLineEdit*)(self) )->virtualbase_HasHeightForWidth();
}

void QLineEdit_override_virtual_PaintEngine(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__PaintEngine = slot;
}

QPaintEngine* QLineEdit_virtualbase_PaintEngine(const void* self) {
	return ( (const MiqtVirtualQLineEdit*)(self) )->virtualbase_PaintEngine();
}

void QLineEdit_override_virtual_WheelEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__WheelEvent = slot;
}

void QLineEdit_virtualbase_WheelEvent(void* self, QWheelEvent* event) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_WheelEvent(event);
}

void QLineEdit_override_virtual_KeyReleaseEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__KeyReleaseEvent = slot;
}

void QLineEdit_virtualbase_KeyReleaseEvent(void* self, QKeyEvent* event) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_KeyReleaseEvent(event);
}

void QLineEdit_override_virtual_EnterEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__EnterEvent = slot;
}

void QLineEdit_virtualbase_EnterEvent(void* self, QEvent* event) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_EnterEvent(event);
}

void QLineEdit_override_virtual_LeaveEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__LeaveEvent = slot;
}

void QLineEdit_virtualbase_LeaveEvent(void* self, QEvent* event) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_LeaveEvent(event);
}

void QLineEdit_override_virtual_MoveEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__MoveEvent = slot;
}

void QLineEdit_virtualbase_MoveEvent(void* self, QMoveEvent* event) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_MoveEvent(event);
}

void QLineEdit_override_virtual_ResizeEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__ResizeEvent = slot;
}

void QLineEdit_virtualbase_ResizeEvent(void* self, QResizeEvent* event) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_ResizeEvent(event);
}

void QLineEdit_override_virtual_CloseEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__CloseEvent = slot;
}

void QLineEdit_virtualbase_CloseEvent(void* self, QCloseEvent* event) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_CloseEvent(event);
}

void QLineEdit_override_virtual_TabletEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__TabletEvent = slot;
}

void QLineEdit_virtualbase_TabletEvent(void* self, QTabletEvent* event) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_TabletEvent(event);
}

void QLineEdit_override_virtual_ActionEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__ActionEvent = slot;
}

void QLineEdit_virtualbase_ActionEvent(void* self, QActionEvent* event) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_ActionEvent(event);
}

void QLineEdit_override_virtual_ShowEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__ShowEvent = slot;
}

void QLineEdit_virtualbase_ShowEvent(void* self, QShowEvent* event) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_ShowEvent(event);
}

void QLineEdit_override_virtual_HideEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__HideEvent = slot;
}

void QLineEdit_virtualbase_HideEvent(void* self, QHideEvent* event) {
	( (MiqtVirtualQLineEdit*)(self) )->virtualbase_HideEvent(event);
}

void QLineEdit_override_virtual_NativeEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__NativeEvent = slot;
}

bool QLineEdit_virtualbase_NativeEvent(void* self, struct miqt_string eventType, void* message, long* result) {
	return ( (MiqtVirtualQLineEdit*)(self) )->virtualbase_NativeEvent(eventType, message, result);
}

void QLineEdit_override_virtual_Metric(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__Metric = slot;
}

int QLineEdit_virtualbase_Metric(const void* self, int param1) {
	return ( (const MiqtVirtualQLineEdit*)(self) )->virtualbase_Metric(param1);
}

void QLineEdit_override_virtual_InitPainter(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__InitPainter = slot;
}

void QLineEdit_virtualbase_InitPainter(const void* self, QPainter* painter) {
	( (const MiqtVirtualQLineEdit*)(self) )->virtualbase_InitPainter(painter);
}

void QLineEdit_override_virtual_Redirected(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__Redirected = slot;
}

QPaintDevice* QLineEdit_virtualbase_Redirected(const void* self, QPoint* offset) {
	return ( (const MiqtVirtualQLineEdit*)(self) )->virtualbase_Redirected(offset);
}

void QLineEdit_override_virtual_SharedPainter(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__SharedPainter = slot;
}

QPainter* QLineEdit_virtualbase_SharedPainter(const void* self) {
	return ( (const MiqtVirtualQLineEdit*)(self) )->virtualbase_SharedPainter();
}

void QLineEdit_override_virtual_FocusNextPrevChild(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLineEdit*>( (QLineEdit*)(self) )->handle__FocusNextPrevChild = slot;
}

bool QLineEdit_virtualbase_FocusNextPrevChild(void* self, bool next) {
	return ( (MiqtVirtualQLineEdit*)(self) )->virtualbase_FocusNextPrevChild(next);
}

void QLineEdit_Delete(QLineEdit* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQLineEdit*>( self );
	} else {
		delete self;
	}
}

