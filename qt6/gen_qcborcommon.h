#ifndef GEN_QCBORCOMMON_H
#define GEN_QCBORCOMMON_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QCborError;
#else
typedef struct QCborError QCborError;
#endif

struct miqt_string QCborError_ToString(const QCborError* self);
void QCborError_Delete(QCborError* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif