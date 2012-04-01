package goxml
/*
#include <stdlib.h>
#include <libxml/xmlerror.h>

void *_xmlGetLastError() {
	return xmlGetLastError();
}
*/
import "C"
import "unsafe"

func XmlGetLastError() unsafe.Pointer {
	return C._xmlGetLastError()
}
