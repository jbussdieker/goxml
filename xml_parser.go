package libxml
/*
#include <stdio.h>
#include <string.h>
#include <libxml/parser.h>

xmlDocPtr _xmlReadMemory(char *buffer) {
	xmlDocPtr doc;
	doc = xmlReadMemory(buffer, strlen(buffer), "", "UTF-8", 0);
	return doc;
}

*/
import "C"
import "unsafe"

// Parses the input string and returns an XML Document
func NewXmlFromString(buf string) Document {
	cbuf := C.CString(buf)
	cdoc := C._xmlReadMemory(cbuf)
	C.free(unsafe.Pointer(cbuf))
	return xmlDocPtr{ptr:cdoc}
}
