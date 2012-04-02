package libxml

/*
#include <libxml/xpath.h>

xmlXPathCompExprPtr _xmlXPathCompile(char *str) {
	return xmlXPathCompile((xmlChar *)str);
}

void _xmlXPathFreeCompExpr(xmlXPathCompExprPtr xpath) {
	xmlXPathFreeCompExpr(xpath);
}

*/
import "C"
import "unsafe"

type Xpath interface {
	Free()
}

type xmlXPathCompExprPtr struct {
	ptr C.xmlXPathCompExprPtr
}

func XmlXpathCompile(xpath string) Xpath {
	cxpath := C.CString(xpath)
	cxpathcomp := C._xmlXPathCompile(cxpath)
	C.free(unsafe.Pointer(cxpath))
	return &xmlXPathCompExprPtr{ptr: cxpathcomp}
}

func (xpath *xmlXPathCompExprPtr) Free() {
	C._xmlXPathFreeCompExpr(xpath.ptr)
}
