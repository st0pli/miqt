#ifndef GEN_QACCESSIBLEWIDGET_H
#define GEN_QACCESSIBLEWIDGET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAccessibleInterface;
class QAccessibleWidget;
class QColor;
class QRect;
class QWindow;
#else
typedef struct QAccessibleInterface QAccessibleInterface;
typedef struct QAccessibleWidget QAccessibleWidget;
typedef struct QColor QColor;
typedef struct QRect QRect;
typedef struct QWindow QWindow;
#endif

bool QAccessibleWidget_IsValid(QAccessibleWidget* self);
QWindow* QAccessibleWidget_Window(QAccessibleWidget* self);
int QAccessibleWidget_ChildCount(QAccessibleWidget* self);
int QAccessibleWidget_IndexOfChild(QAccessibleWidget* self, QAccessibleInterface* child);
QAccessibleInterface* QAccessibleWidget_FocusChild(QAccessibleWidget* self);
QRect* QAccessibleWidget_Rect(QAccessibleWidget* self);
QAccessibleInterface* QAccessibleWidget_Parent(QAccessibleWidget* self);
QAccessibleInterface* QAccessibleWidget_Child(QAccessibleWidget* self, int index);
QColor* QAccessibleWidget_ForegroundColor(QAccessibleWidget* self);
QColor* QAccessibleWidget_BackgroundColor(QAccessibleWidget* self);
void QAccessibleWidget_ActionNames(QAccessibleWidget* self, char*** _out, int** _out_Lengths, size_t* _out_len);
void QAccessibleWidget_DoAction(QAccessibleWidget* self, const char* actionName, size_t actionName_Strlen);
void QAccessibleWidget_KeyBindingsForAction(QAccessibleWidget* self, const char* actionName, size_t actionName_Strlen, char*** _out, int** _out_Lengths, size_t* _out_len);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif