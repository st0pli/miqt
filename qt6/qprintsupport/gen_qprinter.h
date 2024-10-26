#ifndef GEN_QPRINTER_H
#define GEN_QPRINTER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QPaintEngine;
class QPrintEngine;
class QPrinter;
class QPrinterInfo;
class QRectF;
#else
typedef struct QPaintEngine QPaintEngine;
typedef struct QPrintEngine QPrintEngine;
typedef struct QPrinter QPrinter;
typedef struct QPrinterInfo QPrinterInfo;
typedef struct QRectF QRectF;
#endif

QPrinter* QPrinter_new();
QPrinter* QPrinter_new2(QPrinterInfo* printer);
QPrinter* QPrinter_new3(int mode);
QPrinter* QPrinter_new4(QPrinterInfo* printer, int mode);
int QPrinter_DevType(const QPrinter* self);
void QPrinter_SetOutputFormat(QPrinter* self, int format);
int QPrinter_OutputFormat(const QPrinter* self);
void QPrinter_SetPdfVersion(QPrinter* self, int version);
int QPrinter_PdfVersion(const QPrinter* self);
void QPrinter_SetPrinterName(QPrinter* self, struct miqt_string printerName);
struct miqt_string QPrinter_PrinterName(const QPrinter* self);
bool QPrinter_IsValid(const QPrinter* self);
void QPrinter_SetOutputFileName(QPrinter* self, struct miqt_string outputFileName);
struct miqt_string QPrinter_OutputFileName(const QPrinter* self);
void QPrinter_SetPrintProgram(QPrinter* self, struct miqt_string printProgram);
struct miqt_string QPrinter_PrintProgram(const QPrinter* self);
void QPrinter_SetDocName(QPrinter* self, struct miqt_string docName);
struct miqt_string QPrinter_DocName(const QPrinter* self);
void QPrinter_SetCreator(QPrinter* self, struct miqt_string creator);
struct miqt_string QPrinter_Creator(const QPrinter* self);
void QPrinter_SetPageOrder(QPrinter* self, int pageOrder);
int QPrinter_PageOrder(const QPrinter* self);
void QPrinter_SetResolution(QPrinter* self, int resolution);
int QPrinter_Resolution(const QPrinter* self);
void QPrinter_SetColorMode(QPrinter* self, int colorMode);
int QPrinter_ColorMode(const QPrinter* self);
void QPrinter_SetCollateCopies(QPrinter* self, bool collate);
bool QPrinter_CollateCopies(const QPrinter* self);
void QPrinter_SetFullPage(QPrinter* self, bool fullPage);
bool QPrinter_FullPage(const QPrinter* self);
void QPrinter_SetCopyCount(QPrinter* self, int copyCount);
int QPrinter_CopyCount(const QPrinter* self);
bool QPrinter_SupportsMultipleCopies(const QPrinter* self);
void QPrinter_SetPaperSource(QPrinter* self, int paperSource);
int QPrinter_PaperSource(const QPrinter* self);
void QPrinter_SetDuplex(QPrinter* self, int duplex);
int QPrinter_Duplex(const QPrinter* self);
struct miqt_array* QPrinter_SupportedResolutions(const QPrinter* self);
void QPrinter_SetFontEmbeddingEnabled(QPrinter* self, bool enable);
bool QPrinter_FontEmbeddingEnabled(const QPrinter* self);
QRectF* QPrinter_PaperRect(const QPrinter* self, int param1);
QRectF* QPrinter_PageRect(const QPrinter* self, int param1);
struct miqt_string QPrinter_PrinterSelectionOption(const QPrinter* self);
void QPrinter_SetPrinterSelectionOption(QPrinter* self, struct miqt_string printerSelectionOption);
bool QPrinter_NewPage(QPrinter* self);
bool QPrinter_Abort(QPrinter* self);
int QPrinter_PrinterState(const QPrinter* self);
QPaintEngine* QPrinter_PaintEngine(const QPrinter* self);
QPrintEngine* QPrinter_PrintEngine(const QPrinter* self);
void QPrinter_SetFromTo(QPrinter* self, int fromPage, int toPage);
int QPrinter_FromPage(const QPrinter* self);
int QPrinter_ToPage(const QPrinter* self);
void QPrinter_SetPrintRange(QPrinter* self, int rangeVal);
int QPrinter_PrintRange(const QPrinter* self);
void QPrinter_Delete(QPrinter* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif