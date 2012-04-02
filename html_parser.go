package libxml
/*
#include <stdio.h>
#include <string.h>
#include <libxml/HTMLparser.h>

htmlDocPtr _htmlReadMemory(char *buffer) {
	htmlDocPtr doc;
	doc = htmlReadMemory(buffer, strlen(buffer), "", "UTF-8", 0);
	return doc;
}

*/
import "C"
import "unsafe"

// Parses the input string and returns an HTML Document
func NewHtmlFromString(buf string) Document {
	cbuf := C.CString(buf)
	cdoc := C._htmlReadMemory(cbuf)
	C.free(unsafe.Pointer(cbuf))
	return htmlDocPtr{ptr:cdoc}
}
