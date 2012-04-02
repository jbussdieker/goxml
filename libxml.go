//
// Package libxml implements the C library interface to libxml.
package libxml

/*
#cgo pkg-config: libxml-2.0
#cgo LDFLAGS: -lrt
#include <libxml/xmlversion.h>
#include <libxml/parser.h>
void goXmlInit();
unsigned long goXmlAllocSize();
*/
import "C"

// This is the version of libxml that the function definitions are being used from
var LIBXML_VERSION = C.LIBXML_DOTTED_VERSION

// This is the version of libxml that is actually being linked against.
var LIBXML_PARSER_VERSION = C.GoString(C.xmlParserVersion)

func init() {
	C.goXmlInit()
}

// Returns the number of allocated objects.
func MemAllocSize() int {
	return int(C.goXmlAllocSize())
}
