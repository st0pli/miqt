#pragma once
#ifndef MIQT_QT_GEN_QTEXTEDIT_H
#define MIQT_QT_GEN_QTEXTEDIT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractScrollArea;
class QColor;
class QContextMenuEvent;
class QDragEnterEvent;
class QDragLeaveEvent;
class QDragMoveEvent;
class QDropEvent;
class QEvent;
class QFocusEvent;
class QFont;
class QFrame;
class QInputMethodEvent;
class QKeyEvent;
class QMenu;
class QMetaObject;
class QMimeData;
class QMouseEvent;
class QObject;
class QPagedPaintDevice;
class QPaintDevice;
class QPaintEvent;
class QPoint;
class QRect;
class QRegExp;
class QRegularExpression;
class QResizeEvent;
class QShowEvent;
class QSize;
class QTextCharFormat;
class QTextCursor;
class QTextDocument;
class QTextEdit;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QTextEdit__ExtraSelection)
typedef QTextEdit::ExtraSelection QTextEdit__ExtraSelection;
#else
class QTextEdit__ExtraSelection;
#endif
class QTimerEvent;
class QUrl;
class QVariant;
class QWheelEvent;
class QWidget;
#else
typedef struct QAbstractScrollArea QAbstractScrollArea;
typedef struct QColor QColor;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QDragEnterEvent QDragEnterEvent;
typedef struct QDragLeaveEvent QDragLeaveEvent;
typedef struct QDragMoveEvent QDragMoveEvent;
typedef struct QDropEvent QDropEvent;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QFont QFont;
typedef struct QFrame QFrame;
typedef struct QInputMethodEvent QInputMethodEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMenu QMenu;
typedef struct QMetaObject QMetaObject;
typedef struct QMimeData QMimeData;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPagedPaintDevice QPagedPaintDevice;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QRegExp QRegExp;
typedef struct QRegularExpression QRegularExpression;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QTextCharFormat QTextCharFormat;
typedef struct QTextCursor QTextCursor;
typedef struct QTextDocument QTextDocument;
typedef struct QTextEdit QTextEdit;
typedef struct QTextEdit__ExtraSelection QTextEdit__ExtraSelection;
typedef struct QTimerEvent QTimerEvent;
typedef struct QUrl QUrl;
typedef struct QVariant QVariant;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;
#endif

QTextEdit* QTextEdit_new(QWidget* parent);
QTextEdit* QTextEdit_new2();
QTextEdit* QTextEdit_new3(struct miqt_string text);
QTextEdit* QTextEdit_new4(struct miqt_string text, QWidget* parent);
void QTextEdit_virtbase(QTextEdit* src, QAbstractScrollArea** outptr_QAbstractScrollArea);
QMetaObject* QTextEdit_MetaObject(const QTextEdit* self);
void* QTextEdit_Metacast(QTextEdit* self, const char* param1);
struct miqt_string QTextEdit_Tr(const char* s);
struct miqt_string QTextEdit_TrUtf8(const char* s);
void QTextEdit_SetDocument(QTextEdit* self, QTextDocument* document);
QTextDocument* QTextEdit_Document(const QTextEdit* self);
void QTextEdit_SetPlaceholderText(QTextEdit* self, struct miqt_string placeholderText);
struct miqt_string QTextEdit_PlaceholderText(const QTextEdit* self);
void QTextEdit_SetTextCursor(QTextEdit* self, QTextCursor* cursor);
QTextCursor* QTextEdit_TextCursor(const QTextEdit* self);
bool QTextEdit_IsReadOnly(const QTextEdit* self);
void QTextEdit_SetReadOnly(QTextEdit* self, bool ro);
void QTextEdit_SetTextInteractionFlags(QTextEdit* self, int flags);
int QTextEdit_TextInteractionFlags(const QTextEdit* self);
double QTextEdit_FontPointSize(const QTextEdit* self);
struct miqt_string QTextEdit_FontFamily(const QTextEdit* self);
int QTextEdit_FontWeight(const QTextEdit* self);
bool QTextEdit_FontUnderline(const QTextEdit* self);
bool QTextEdit_FontItalic(const QTextEdit* self);
QColor* QTextEdit_TextColor(const QTextEdit* self);
QColor* QTextEdit_TextBackgroundColor(const QTextEdit* self);
QFont* QTextEdit_CurrentFont(const QTextEdit* self);
int QTextEdit_Alignment(const QTextEdit* self);
void QTextEdit_MergeCurrentCharFormat(QTextEdit* self, QTextCharFormat* modifier);
void QTextEdit_SetCurrentCharFormat(QTextEdit* self, QTextCharFormat* format);
QTextCharFormat* QTextEdit_CurrentCharFormat(const QTextEdit* self);
int QTextEdit_AutoFormatting(const QTextEdit* self);
void QTextEdit_SetAutoFormatting(QTextEdit* self, int features);
bool QTextEdit_TabChangesFocus(const QTextEdit* self);
void QTextEdit_SetTabChangesFocus(QTextEdit* self, bool b);
void QTextEdit_SetDocumentTitle(QTextEdit* self, struct miqt_string title);
struct miqt_string QTextEdit_DocumentTitle(const QTextEdit* self);
bool QTextEdit_IsUndoRedoEnabled(const QTextEdit* self);
void QTextEdit_SetUndoRedoEnabled(QTextEdit* self, bool enable);
int QTextEdit_LineWrapMode(const QTextEdit* self);
void QTextEdit_SetLineWrapMode(QTextEdit* self, int mode);
int QTextEdit_LineWrapColumnOrWidth(const QTextEdit* self);
void QTextEdit_SetLineWrapColumnOrWidth(QTextEdit* self, int w);
int QTextEdit_WordWrapMode(const QTextEdit* self);
void QTextEdit_SetWordWrapMode(QTextEdit* self, int policy);
bool QTextEdit_Find(QTextEdit* self, struct miqt_string exp);
bool QTextEdit_FindWithExp(QTextEdit* self, QRegExp* exp);
bool QTextEdit_Find2(QTextEdit* self, QRegularExpression* exp);
struct miqt_string QTextEdit_ToPlainText(const QTextEdit* self);
struct miqt_string QTextEdit_ToHtml(const QTextEdit* self);
struct miqt_string QTextEdit_ToMarkdown(const QTextEdit* self);
void QTextEdit_EnsureCursorVisible(QTextEdit* self);
QVariant* QTextEdit_LoadResource(QTextEdit* self, int typeVal, QUrl* name);
QMenu* QTextEdit_CreateStandardContextMenu(QTextEdit* self);
QMenu* QTextEdit_CreateStandardContextMenuWithPosition(QTextEdit* self, QPoint* position);
QTextCursor* QTextEdit_CursorForPosition(const QTextEdit* self, QPoint* pos);
QRect* QTextEdit_CursorRect(const QTextEdit* self, QTextCursor* cursor);
QRect* QTextEdit_CursorRect2(const QTextEdit* self);
struct miqt_string QTextEdit_AnchorAt(const QTextEdit* self, QPoint* pos);
bool QTextEdit_OverwriteMode(const QTextEdit* self);
void QTextEdit_SetOverwriteMode(QTextEdit* self, bool overwrite);
int QTextEdit_TabStopWidth(const QTextEdit* self);
void QTextEdit_SetTabStopWidth(QTextEdit* self, int width);
double QTextEdit_TabStopDistance(const QTextEdit* self);
void QTextEdit_SetTabStopDistance(QTextEdit* self, double distance);
int QTextEdit_CursorWidth(const QTextEdit* self);
void QTextEdit_SetCursorWidth(QTextEdit* self, int width);
bool QTextEdit_AcceptRichText(const QTextEdit* self);
void QTextEdit_SetAcceptRichText(QTextEdit* self, bool accept);
void QTextEdit_SetExtraSelections(QTextEdit* self, struct miqt_array /* of QTextEdit__ExtraSelection* */  selections);
struct miqt_array /* of QTextEdit__ExtraSelection* */  QTextEdit_ExtraSelections(const QTextEdit* self);
void QTextEdit_MoveCursor(QTextEdit* self, int operation);
bool QTextEdit_CanPaste(const QTextEdit* self);
void QTextEdit_Print(const QTextEdit* self, QPagedPaintDevice* printer);
QVariant* QTextEdit_InputMethodQuery(const QTextEdit* self, int property);
QVariant* QTextEdit_InputMethodQuery2(const QTextEdit* self, int query, QVariant* argument);
void QTextEdit_SetFontPointSize(QTextEdit* self, double s);
void QTextEdit_SetFontFamily(QTextEdit* self, struct miqt_string fontFamily);
void QTextEdit_SetFontWeight(QTextEdit* self, int w);
void QTextEdit_SetFontUnderline(QTextEdit* self, bool b);
void QTextEdit_SetFontItalic(QTextEdit* self, bool b);
void QTextEdit_SetTextColor(QTextEdit* self, QColor* c);
void QTextEdit_SetTextBackgroundColor(QTextEdit* self, QColor* c);
void QTextEdit_SetCurrentFont(QTextEdit* self, QFont* f);
void QTextEdit_SetAlignment(QTextEdit* self, int a);
void QTextEdit_SetPlainText(QTextEdit* self, struct miqt_string text);
void QTextEdit_SetHtml(QTextEdit* self, struct miqt_string text);
void QTextEdit_SetMarkdown(QTextEdit* self, struct miqt_string markdown);
void QTextEdit_SetText(QTextEdit* self, struct miqt_string text);
void QTextEdit_Cut(QTextEdit* self);
void QTextEdit_Copy(QTextEdit* self);
void QTextEdit_Paste(QTextEdit* self);
void QTextEdit_Undo(QTextEdit* self);
void QTextEdit_Redo(QTextEdit* self);
void QTextEdit_Clear(QTextEdit* self);
void QTextEdit_SelectAll(QTextEdit* self);
void QTextEdit_InsertPlainText(QTextEdit* self, struct miqt_string text);
void QTextEdit_InsertHtml(QTextEdit* self, struct miqt_string text);
void QTextEdit_Append(QTextEdit* self, struct miqt_string text);
void QTextEdit_ScrollToAnchor(QTextEdit* self, struct miqt_string name);
void QTextEdit_ZoomIn(QTextEdit* self);
void QTextEdit_ZoomOut(QTextEdit* self);
void QTextEdit_TextChanged(QTextEdit* self);
void QTextEdit_connect_TextChanged(QTextEdit* self, intptr_t slot);
void QTextEdit_UndoAvailable(QTextEdit* self, bool b);
void QTextEdit_connect_UndoAvailable(QTextEdit* self, intptr_t slot);
void QTextEdit_RedoAvailable(QTextEdit* self, bool b);
void QTextEdit_connect_RedoAvailable(QTextEdit* self, intptr_t slot);
void QTextEdit_CurrentCharFormatChanged(QTextEdit* self, QTextCharFormat* format);
void QTextEdit_connect_CurrentCharFormatChanged(QTextEdit* self, intptr_t slot);
void QTextEdit_CopyAvailable(QTextEdit* self, bool b);
void QTextEdit_connect_CopyAvailable(QTextEdit* self, intptr_t slot);
void QTextEdit_SelectionChanged(QTextEdit* self);
void QTextEdit_connect_SelectionChanged(QTextEdit* self, intptr_t slot);
void QTextEdit_CursorPositionChanged(QTextEdit* self);
void QTextEdit_connect_CursorPositionChanged(QTextEdit* self, intptr_t slot);
bool QTextEdit_Event(QTextEdit* self, QEvent* e);
void QTextEdit_TimerEvent(QTextEdit* self, QTimerEvent* e);
void QTextEdit_KeyPressEvent(QTextEdit* self, QKeyEvent* e);
void QTextEdit_KeyReleaseEvent(QTextEdit* self, QKeyEvent* e);
void QTextEdit_ResizeEvent(QTextEdit* self, QResizeEvent* e);
void QTextEdit_PaintEvent(QTextEdit* self, QPaintEvent* e);
void QTextEdit_MousePressEvent(QTextEdit* self, QMouseEvent* e);
void QTextEdit_MouseMoveEvent(QTextEdit* self, QMouseEvent* e);
void QTextEdit_MouseReleaseEvent(QTextEdit* self, QMouseEvent* e);
void QTextEdit_MouseDoubleClickEvent(QTextEdit* self, QMouseEvent* e);
bool QTextEdit_FocusNextPrevChild(QTextEdit* self, bool next);
void QTextEdit_ContextMenuEvent(QTextEdit* self, QContextMenuEvent* e);
void QTextEdit_DragEnterEvent(QTextEdit* self, QDragEnterEvent* e);
void QTextEdit_DragLeaveEvent(QTextEdit* self, QDragLeaveEvent* e);
void QTextEdit_DragMoveEvent(QTextEdit* self, QDragMoveEvent* e);
void QTextEdit_DropEvent(QTextEdit* self, QDropEvent* e);
void QTextEdit_FocusInEvent(QTextEdit* self, QFocusEvent* e);
void QTextEdit_FocusOutEvent(QTextEdit* self, QFocusEvent* e);
void QTextEdit_ShowEvent(QTextEdit* self, QShowEvent* param1);
void QTextEdit_ChangeEvent(QTextEdit* self, QEvent* e);
void QTextEdit_WheelEvent(QTextEdit* self, QWheelEvent* e);
QMimeData* QTextEdit_CreateMimeDataFromSelection(const QTextEdit* self);
bool QTextEdit_CanInsertFromMimeData(const QTextEdit* self, QMimeData* source);
void QTextEdit_InsertFromMimeData(QTextEdit* self, QMimeData* source);
void QTextEdit_InputMethodEvent(QTextEdit* self, QInputMethodEvent* param1);
void QTextEdit_ScrollContentsBy(QTextEdit* self, int dx, int dy);
void QTextEdit_DoSetTextCursor(QTextEdit* self, QTextCursor* cursor);
struct miqt_string QTextEdit_Tr2(const char* s, const char* c);
struct miqt_string QTextEdit_Tr3(const char* s, const char* c, int n);
struct miqt_string QTextEdit_TrUtf82(const char* s, const char* c);
struct miqt_string QTextEdit_TrUtf83(const char* s, const char* c, int n);
bool QTextEdit_Find22(QTextEdit* self, struct miqt_string exp, int options);
bool QTextEdit_Find23(QTextEdit* self, QRegExp* exp, int options);
bool QTextEdit_Find24(QTextEdit* self, QRegularExpression* exp, int options);
struct miqt_string QTextEdit_ToMarkdown1(const QTextEdit* self, int features);
void QTextEdit_MoveCursor2(QTextEdit* self, int operation, int mode);
void QTextEdit_ZoomIn1(QTextEdit* self, int rangeVal);
void QTextEdit_ZoomOut1(QTextEdit* self, int rangeVal);
void QTextEdit_override_virtual_LoadResource(void* self, intptr_t slot);
QVariant* QTextEdit_virtualbase_LoadResource(void* self, int typeVal, QUrl* name);
void QTextEdit_override_virtual_InputMethodQuery(void* self, intptr_t slot);
QVariant* QTextEdit_virtualbase_InputMethodQuery(const void* self, int property);
void QTextEdit_override_virtual_Event(void* self, intptr_t slot);
bool QTextEdit_virtualbase_Event(void* self, QEvent* e);
void QTextEdit_override_virtual_TimerEvent(void* self, intptr_t slot);
void QTextEdit_virtualbase_TimerEvent(void* self, QTimerEvent* e);
void QTextEdit_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QTextEdit_virtualbase_KeyPressEvent(void* self, QKeyEvent* e);
void QTextEdit_override_virtual_KeyReleaseEvent(void* self, intptr_t slot);
void QTextEdit_virtualbase_KeyReleaseEvent(void* self, QKeyEvent* e);
void QTextEdit_override_virtual_ResizeEvent(void* self, intptr_t slot);
void QTextEdit_virtualbase_ResizeEvent(void* self, QResizeEvent* e);
void QTextEdit_override_virtual_PaintEvent(void* self, intptr_t slot);
void QTextEdit_virtualbase_PaintEvent(void* self, QPaintEvent* e);
void QTextEdit_override_virtual_MousePressEvent(void* self, intptr_t slot);
void QTextEdit_virtualbase_MousePressEvent(void* self, QMouseEvent* e);
void QTextEdit_override_virtual_MouseMoveEvent(void* self, intptr_t slot);
void QTextEdit_virtualbase_MouseMoveEvent(void* self, QMouseEvent* e);
void QTextEdit_override_virtual_MouseReleaseEvent(void* self, intptr_t slot);
void QTextEdit_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* e);
void QTextEdit_override_virtual_MouseDoubleClickEvent(void* self, intptr_t slot);
void QTextEdit_virtualbase_MouseDoubleClickEvent(void* self, QMouseEvent* e);
void QTextEdit_override_virtual_FocusNextPrevChild(void* self, intptr_t slot);
bool QTextEdit_virtualbase_FocusNextPrevChild(void* self, bool next);
void QTextEdit_override_virtual_ContextMenuEvent(void* self, intptr_t slot);
void QTextEdit_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* e);
void QTextEdit_override_virtual_DragEnterEvent(void* self, intptr_t slot);
void QTextEdit_virtualbase_DragEnterEvent(void* self, QDragEnterEvent* e);
void QTextEdit_override_virtual_DragLeaveEvent(void* self, intptr_t slot);
void QTextEdit_virtualbase_DragLeaveEvent(void* self, QDragLeaveEvent* e);
void QTextEdit_override_virtual_DragMoveEvent(void* self, intptr_t slot);
void QTextEdit_virtualbase_DragMoveEvent(void* self, QDragMoveEvent* e);
void QTextEdit_override_virtual_DropEvent(void* self, intptr_t slot);
void QTextEdit_virtualbase_DropEvent(void* self, QDropEvent* e);
void QTextEdit_override_virtual_FocusInEvent(void* self, intptr_t slot);
void QTextEdit_virtualbase_FocusInEvent(void* self, QFocusEvent* e);
void QTextEdit_override_virtual_FocusOutEvent(void* self, intptr_t slot);
void QTextEdit_virtualbase_FocusOutEvent(void* self, QFocusEvent* e);
void QTextEdit_override_virtual_ShowEvent(void* self, intptr_t slot);
void QTextEdit_virtualbase_ShowEvent(void* self, QShowEvent* param1);
void QTextEdit_override_virtual_ChangeEvent(void* self, intptr_t slot);
void QTextEdit_virtualbase_ChangeEvent(void* self, QEvent* e);
void QTextEdit_override_virtual_WheelEvent(void* self, intptr_t slot);
void QTextEdit_virtualbase_WheelEvent(void* self, QWheelEvent* e);
void QTextEdit_override_virtual_CreateMimeDataFromSelection(void* self, intptr_t slot);
QMimeData* QTextEdit_virtualbase_CreateMimeDataFromSelection(const void* self);
void QTextEdit_override_virtual_CanInsertFromMimeData(void* self, intptr_t slot);
bool QTextEdit_virtualbase_CanInsertFromMimeData(const void* self, QMimeData* source);
void QTextEdit_override_virtual_InsertFromMimeData(void* self, intptr_t slot);
void QTextEdit_virtualbase_InsertFromMimeData(void* self, QMimeData* source);
void QTextEdit_override_virtual_InputMethodEvent(void* self, intptr_t slot);
void QTextEdit_virtualbase_InputMethodEvent(void* self, QInputMethodEvent* param1);
void QTextEdit_override_virtual_ScrollContentsBy(void* self, intptr_t slot);
void QTextEdit_virtualbase_ScrollContentsBy(void* self, int dx, int dy);
void QTextEdit_override_virtual_DoSetTextCursor(void* self, intptr_t slot);
void QTextEdit_virtualbase_DoSetTextCursor(void* self, QTextCursor* cursor);
void QTextEdit_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QTextEdit_virtualbase_MinimumSizeHint(const void* self);
void QTextEdit_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QTextEdit_virtualbase_SizeHint(const void* self);
void QTextEdit_override_virtual_SetupViewport(void* self, intptr_t slot);
void QTextEdit_virtualbase_SetupViewport(void* self, QWidget* viewport);
void QTextEdit_override_virtual_EventFilter(void* self, intptr_t slot);
bool QTextEdit_virtualbase_EventFilter(void* self, QObject* param1, QEvent* param2);
void QTextEdit_override_virtual_ViewportEvent(void* self, intptr_t slot);
bool QTextEdit_virtualbase_ViewportEvent(void* self, QEvent* param1);
void QTextEdit_override_virtual_ViewportSizeHint(void* self, intptr_t slot);
QSize* QTextEdit_virtualbase_ViewportSizeHint(const void* self);
void QTextEdit_Delete(QTextEdit* self, bool isSubclass);

QTextEdit__ExtraSelection* QTextEdit__ExtraSelection_new(QTextEdit__ExtraSelection* param1);
void QTextEdit__ExtraSelection_OperatorAssign(QTextEdit__ExtraSelection* self, QTextEdit__ExtraSelection* param1);
void QTextEdit__ExtraSelection_Delete(QTextEdit__ExtraSelection* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
