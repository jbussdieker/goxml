package libxml

/*
#cgo pkg-config: libxml-2.0
#include <stdlib.h>
#include <libxml/tree.h>
*/
import "C"
import "unsafe"

func xmlCharToC(c *C.xmlChar) *C.char {
	return (*C.char)(unsafe.Pointer(c))
}

func cToXmlChar(c *C.char) *C.xmlChar {
	return (*C.xmlChar)(unsafe.Pointer(c))
}
