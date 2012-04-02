//
// Package libxml implements the C library interface to libxml.
package libxml

/*
#cgo CFLAGS: -I/tmp/trinitybuild/include/libxml2
#cgo LDFLAGS: -lxml2
#include <libxml/xmlversion.h>
#include <libxml/parser.h>
void goXmlInit();
unsigned long goXmlAllocSize();
*/
import "C"

// This is the version of libxml that the function definitions are being used from
var LIBXML_VERSION = C.LIBXML_DOTTED_VERSION

// This is the version of libxml that is actually being linked against.
var LIBXML_PARSER_VERSION = C.LIBXML_DOTTED_VERSION

func init() {
	C.goXmlInit()
}

// Returns the number of allocated objects.
func MemAllocSize() int {
	return int(C.goXmlAllocSize())
}
