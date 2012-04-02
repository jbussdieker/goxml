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

var (
	LIBXML_VERSION = C.LIBXML_DOTTED_VERSION
	LIBXML_PARSER_VERSION = C.GoString(C.xmlParserVersion)
)

func init() {
	C.goXmlInit()
}

func AllocSize() int {
	return int(C.goXmlAllocSize())
}
