#pragma once
#ifndef MIQT_QT6_GEN_QICONENGINEPLUGIN_H
#define MIQT_QT6_GEN_QICONENGINEPLUGIN_H

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
class QIconEngine;
class QIconEnginePlugin;
class QMetaMethod;
class QMetaObject;
class QObject;
class QTimerEvent;
#else
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QIconEngine QIconEngine;
typedef struct QIconEnginePlugin QIconEnginePlugin;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTimerEvent QTimerEvent;
#endif

QIconEnginePlugin* QIconEnginePlugin_new();
QIconEnginePlugin* QIconEnginePlugin_new2(QObject* parent);
void QIconEnginePlugin_virtbase(QIconEnginePlugin* src, QObject** outptr_QObject);
QMetaObject* QIconEnginePlugin_MetaObject(const QIconEnginePlugin* self);
void* QIconEnginePlugin_Metacast(QIconEnginePlugin* self, const char* param1);
struct miqt_string QIconEnginePlugin_Tr(const char* s);
QIconEngine* QIconEnginePlugin_Create(QIconEnginePlugin* self, struct miqt_string filename);
struct miqt_string QIconEnginePlugin_Tr2(const char* s, const char* c);
struct miqt_string QIconEnginePlugin_Tr3(const char* s, const char* c, int n);
void QIconEnginePlugin_override_virtual_Create(void* self, intptr_t slot);
QIconEngine* QIconEnginePlugin_virtualbase_Create(void* self, struct miqt_string filename);
void QIconEnginePlugin_override_virtual_Event(void* self, intptr_t slot);
bool QIconEnginePlugin_virtualbase_Event(void* self, QEvent* event);
void QIconEnginePlugin_override_virtual_EventFilter(void* self, intptr_t slot);
bool QIconEnginePlugin_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
void QIconEnginePlugin_override_virtual_TimerEvent(void* self, intptr_t slot);
void QIconEnginePlugin_virtualbase_TimerEvent(void* self, QTimerEvent* event);
void QIconEnginePlugin_override_virtual_ChildEvent(void* self, intptr_t slot);
void QIconEnginePlugin_virtualbase_ChildEvent(void* self, QChildEvent* event);
void QIconEnginePlugin_override_virtual_CustomEvent(void* self, intptr_t slot);
void QIconEnginePlugin_virtualbase_CustomEvent(void* self, QEvent* event);
void QIconEnginePlugin_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QIconEnginePlugin_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
void QIconEnginePlugin_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QIconEnginePlugin_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
void QIconEnginePlugin_Delete(QIconEnginePlugin* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
