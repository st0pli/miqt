#ifndef GEN_QPICTURE_H
#define GEN_QPICTURE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QByteArray;
class QIODevice;
class QPaintEngine;
class QPainter;
class QPicture;
class QPictureIO;
class QRect;
#else
typedef struct QByteArray QByteArray;
typedef struct QIODevice QIODevice;
typedef struct QPaintEngine QPaintEngine;
typedef struct QPainter QPainter;
typedef struct QPicture QPicture;
typedef struct QPictureIO QPictureIO;
typedef struct QRect QRect;
#endif

QPicture* QPicture_new();
QPicture* QPicture_new2(QPicture* param1);
QPicture* QPicture_new3(int formatVersion);
bool QPicture_IsNull(QPicture* self);
int QPicture_DevType(QPicture* self);
unsigned int QPicture_Size(QPicture* self);
char* QPicture_Data(QPicture* self);
void QPicture_SetData(QPicture* self, char* data, unsigned int size);
bool QPicture_Play(QPicture* self, QPainter* p);
bool QPicture_Load(QPicture* self, QIODevice* dev);
bool QPicture_LoadWithFileName(QPicture* self, const char* fileName, size_t fileName_Strlen);
bool QPicture_Save(QPicture* self, QIODevice* dev);
bool QPicture_SaveWithFileName(QPicture* self, const char* fileName, size_t fileName_Strlen);
QRect* QPicture_BoundingRect(QPicture* self);
void QPicture_SetBoundingRect(QPicture* self, QRect* r);
void QPicture_OperatorAssign(QPicture* self, QPicture* p);
void QPicture_Swap(QPicture* self, QPicture* other);
void QPicture_Detach(QPicture* self);
bool QPicture_IsDetached(QPicture* self);
char* QPicture_PictureFormat(const char* fileName, size_t fileName_Strlen);
void QPicture_InputFormats(QByteArray*** _out, size_t* _out_len);
void QPicture_OutputFormats(QByteArray*** _out, size_t* _out_len);
void QPicture_InputFormatList(char*** _out, int** _out_Lengths, size_t* _out_len);
void QPicture_OutputFormatList(char*** _out, int** _out_Lengths, size_t* _out_len);
QPaintEngine* QPicture_PaintEngine(QPicture* self);
bool QPicture_Load2(QPicture* self, QIODevice* dev, char* format);
bool QPicture_Load22(QPicture* self, const char* fileName, size_t fileName_Strlen, char* format);
bool QPicture_Save2(QPicture* self, QIODevice* dev, char* format);
bool QPicture_Save22(QPicture* self, const char* fileName, size_t fileName_Strlen, char* format);
void QPicture_Delete(QPicture* self);

QPictureIO* QPictureIO_new();
QPictureIO* QPictureIO_new2(QIODevice* ioDevice, char* format);
QPictureIO* QPictureIO_new3(const char* fileName, size_t fileName_Strlen, char* format);
QPicture* QPictureIO_Picture(QPictureIO* self);
int QPictureIO_Status(QPictureIO* self);
char* QPictureIO_Format(QPictureIO* self);
QIODevice* QPictureIO_IoDevice(QPictureIO* self);
void QPictureIO_FileName(QPictureIO* self, char** _out, int* _out_Strlen);
int QPictureIO_Quality(QPictureIO* self);
void QPictureIO_Description(QPictureIO* self, char** _out, int* _out_Strlen);
char* QPictureIO_Parameters(QPictureIO* self);
float QPictureIO_Gamma(QPictureIO* self);
void QPictureIO_SetPicture(QPictureIO* self, QPicture* picture);
void QPictureIO_SetStatus(QPictureIO* self, int status);
void QPictureIO_SetFormat(QPictureIO* self, char* format);
void QPictureIO_SetIODevice(QPictureIO* self, QIODevice* iODevice);
void QPictureIO_SetFileName(QPictureIO* self, const char* fileName, size_t fileName_Strlen);
void QPictureIO_SetQuality(QPictureIO* self, int quality);
void QPictureIO_SetDescription(QPictureIO* self, const char* description, size_t description_Strlen);
void QPictureIO_SetParameters(QPictureIO* self, char* parameters);
void QPictureIO_SetGamma(QPictureIO* self, float gamma);
bool QPictureIO_Read(QPictureIO* self);
bool QPictureIO_Write(QPictureIO* self);
QByteArray* QPictureIO_PictureFormat(const char* fileName, size_t fileName_Strlen);
QByteArray* QPictureIO_PictureFormatWithQIODevice(QIODevice* param1);
void QPictureIO_InputFormats(QByteArray*** _out, size_t* _out_len);
void QPictureIO_OutputFormats(QByteArray*** _out, size_t* _out_len);
void QPictureIO_Delete(QPictureIO* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif