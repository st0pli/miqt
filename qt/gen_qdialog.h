#ifndef GEN_QDIALOG_H
#define GEN_QDIALOG_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QDialog;
class QMetaObject;
class QSize;
class QWidget;
#else
typedef struct QDialog QDialog;
typedef struct QMetaObject QMetaObject;
typedef struct QSize QSize;
typedef struct QWidget QWidget;
#endif

QMetaObject* QDialog_MetaObject(QDialog* self);
void QDialog_Tr(char* s, char** _out, int* _out_Strlen);
void QDialog_TrUtf8(char* s, char** _out, int* _out_Strlen);
int QDialog_Result(QDialog* self);
void QDialog_SetVisible(QDialog* self, bool visible);
void QDialog_SetExtension(QDialog* self, QWidget* extension);
QWidget* QDialog_Extension(QDialog* self);
QSize* QDialog_SizeHint(QDialog* self);
QSize* QDialog_MinimumSizeHint(QDialog* self);
void QDialog_SetSizeGripEnabled(QDialog* self, bool sizeGripEnabled);
bool QDialog_IsSizeGripEnabled(QDialog* self);
void QDialog_SetModal(QDialog* self, bool modal);
void QDialog_SetResult(QDialog* self, int r);
void QDialog_Finished(QDialog* self, int result);
void QDialog_connect_Finished(QDialog* self, void* slot);
void QDialog_Accepted(QDialog* self);
void QDialog_connect_Accepted(QDialog* self, void* slot);
void QDialog_Rejected(QDialog* self);
void QDialog_connect_Rejected(QDialog* self, void* slot);
void QDialog_Open(QDialog* self);
int QDialog_Exec(QDialog* self);
void QDialog_Done(QDialog* self, int param1);
void QDialog_Accept(QDialog* self);
void QDialog_Reject(QDialog* self);
void QDialog_ShowExtension(QDialog* self, bool param1);
void QDialog_Tr2(char* s, char* c, char** _out, int* _out_Strlen);
void QDialog_Tr3(char* s, char* c, int n, char** _out, int* _out_Strlen);
void QDialog_TrUtf82(char* s, char* c, char** _out, int* _out_Strlen);
void QDialog_TrUtf83(char* s, char* c, int n, char** _out, int* _out_Strlen);
void QDialog_Delete(QDialog* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif