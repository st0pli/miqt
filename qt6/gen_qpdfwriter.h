#pragma once
#ifndef MIQT_QT6_GEN_QPDFWRITER_H
#define MIQT_QT6_GEN_QPDFWRITER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QChildEvent;
class QEvent;
class QIODevice;
class QMarginsF;
class QMetaMethod;
class QMetaObject;
class QObject;
class QPageLayout;
class QPageRanges;
class QPageSize;
class QPagedPaintDevice;
class QPaintDevice;
class QPaintEngine;
class QPdfWriter;
class QTimerEvent;
#else
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QIODevice QIODevice;
typedef struct QMarginsF QMarginsF;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPageLayout QPageLayout;
typedef struct QPageRanges QPageRanges;
typedef struct QPageSize QPageSize;
typedef struct QPagedPaintDevice QPagedPaintDevice;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEngine QPaintEngine;
typedef struct QPdfWriter QPdfWriter;
typedef struct QTimerEvent QTimerEvent;
#endif

QPdfWriter* QPdfWriter_new(struct miqt_string filename);
QPdfWriter* QPdfWriter_new2(QIODevice* device);
void QPdfWriter_virtbase(QPdfWriter* src, QObject** outptr_QObject, QPagedPaintDevice** outptr_QPagedPaintDevice);
QMetaObject* QPdfWriter_MetaObject(const QPdfWriter* self);
void* QPdfWriter_Metacast(QPdfWriter* self, const char* param1);
struct miqt_string QPdfWriter_Tr(const char* s);
void QPdfWriter_SetPdfVersion(QPdfWriter* self, int version);
int QPdfWriter_PdfVersion(const QPdfWriter* self);
struct miqt_string QPdfWriter_Title(const QPdfWriter* self);
void QPdfWriter_SetTitle(QPdfWriter* self, struct miqt_string title);
struct miqt_string QPdfWriter_Creator(const QPdfWriter* self);
void QPdfWriter_SetCreator(QPdfWriter* self, struct miqt_string creator);
bool QPdfWriter_NewPage(QPdfWriter* self);
void QPdfWriter_SetResolution(QPdfWriter* self, int resolution);
int QPdfWriter_Resolution(const QPdfWriter* self);
void QPdfWriter_SetDocumentXmpMetadata(QPdfWriter* self, struct miqt_string xmpMetadata);
struct miqt_string QPdfWriter_DocumentXmpMetadata(const QPdfWriter* self);
void QPdfWriter_AddFileAttachment(QPdfWriter* self, struct miqt_string fileName, struct miqt_string data);
QPaintEngine* QPdfWriter_PaintEngine(const QPdfWriter* self);
int QPdfWriter_Metric(const QPdfWriter* self, int id);
struct miqt_string QPdfWriter_Tr2(const char* s, const char* c);
struct miqt_string QPdfWriter_Tr3(const char* s, const char* c, int n);
void QPdfWriter_AddFileAttachment3(QPdfWriter* self, struct miqt_string fileName, struct miqt_string data, struct miqt_string mimeType);
void QPdfWriter_override_virtual_NewPage(void* self, intptr_t slot);
bool QPdfWriter_virtualbase_NewPage(void* self);
void QPdfWriter_override_virtual_PaintEngine(void* self, intptr_t slot);
QPaintEngine* QPdfWriter_virtualbase_PaintEngine(const void* self);
void QPdfWriter_override_virtual_Metric(void* self, intptr_t slot);
int QPdfWriter_virtualbase_Metric(const void* self, int id);
void QPdfWriter_override_virtual_Event(void* self, intptr_t slot);
bool QPdfWriter_virtualbase_Event(void* self, QEvent* event);
void QPdfWriter_override_virtual_EventFilter(void* self, intptr_t slot);
bool QPdfWriter_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
void QPdfWriter_override_virtual_TimerEvent(void* self, intptr_t slot);
void QPdfWriter_virtualbase_TimerEvent(void* self, QTimerEvent* event);
void QPdfWriter_override_virtual_ChildEvent(void* self, intptr_t slot);
void QPdfWriter_virtualbase_ChildEvent(void* self, QChildEvent* event);
void QPdfWriter_override_virtual_CustomEvent(void* self, intptr_t slot);
void QPdfWriter_virtualbase_CustomEvent(void* self, QEvent* event);
void QPdfWriter_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QPdfWriter_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
void QPdfWriter_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QPdfWriter_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
void QPdfWriter_override_virtual_SetPageLayout(void* self, intptr_t slot);
bool QPdfWriter_virtualbase_SetPageLayout(void* self, QPageLayout* pageLayout);
void QPdfWriter_override_virtual_SetPageSize(void* self, intptr_t slot);
bool QPdfWriter_virtualbase_SetPageSize(void* self, QPageSize* pageSize);
void QPdfWriter_override_virtual_SetPageOrientation(void* self, intptr_t slot);
bool QPdfWriter_virtualbase_SetPageOrientation(void* self, int orientation);
void QPdfWriter_override_virtual_SetPageMargins(void* self, intptr_t slot);
bool QPdfWriter_virtualbase_SetPageMargins(void* self, QMarginsF* margins, int units);
void QPdfWriter_override_virtual_SetPageRanges(void* self, intptr_t slot);
void QPdfWriter_virtualbase_SetPageRanges(void* self, QPageRanges* ranges);
void QPdfWriter_Delete(QPdfWriter* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
