package goxml
/*
#include <stdlib.h>
#include <libxml/HTMLtree.h>
#include <libxml/HTMLparser.h>

void *_htmlNewDoc(char *uri, char *external_id) {
	return htmlNewDoc((xmlChar *)uri, (xmlChar *)external_id);
}

void *_htmlReadMemory(char *buffer, int size, char *url, char *encoding, int options) {
	return htmlReadMemory(buffer, size, url, encoding, options);
}

char *_htmlNodeDump(void *doc, void *node) {
	xmlBufferPtr buffer = 0;
	htmlNodeDump(buffer, doc, node);
	return (char *)buffer->content;
}
*/
import "C"
import "unsafe"

const (
	HTML_PARSE_RECOVER   = C.HTML_PARSE_RECOVER   //: Relaxed parsing
	HTML_PARSE_NODEFDTD  = C.HTML_PARSE_NODEFDTD  //: do not default a doctype if not found
	HTML_PARSE_NOERROR   = C.HTML_PARSE_NOERROR   //: suppress error reports
	HTML_PARSE_NOWARNING = C.HTML_PARSE_NOWARNING //: suppress warning reports
	HTML_PARSE_PEDANTIC  = C.HTML_PARSE_PEDANTIC  //: pedantic error reporting
	HTML_PARSE_NOBLANKS  = C.HTML_PARSE_NOBLANKS  //: remove blank nodes
	HTML_PARSE_NONET     = C.HTML_PARSE_NONET     //: Forbid network access
	HTML_PARSE_NOIMPLIED = C.HTML_PARSE_NOIMPLIED //: Do not add implied html/body... elements
	HTML_PARSE_COMPACT   = C.HTML_PARSE_COMPACT   //: compact small text nodes
)

const (
	HTML_NA         = C.HTML_NA //: something we don't check at all
	HTML_INVALID    = C.HTML_INVALID
	HTML_DEPRECATED = C.HTML_DEPRECATED
	HTML_VALID      = C.HTML_VALID
	HTML_REQUIRED   = C.HTML_REQUIRED //: VALID bit set so ( & HTML_VALID ) is TRUE
)

type htmlDoc struct {
	xmlNode
}

type htmlNode struct {
	xmlNode
}

func HtmlNewDoc(uri string, external_id string) Document {
	curi := C.CString(uri)
	cexternal_id := C.CString(external_id)
	cdoc := C._htmlNewDoc(curi, cexternal_id)
	C.free(unsafe.Pointer(curi))
	C.free(unsafe.Pointer(cexternal_id))
	node := xmlNode{Cdoc:cdoc}
	return Document(htmlDoc{xmlNode:node})
}

func (doc htmlDoc) Free() {
	xmlDoc(doc).Free()
}

func (doc htmlDoc) Debug() {
	xmlDoc(doc).Debug()
}

func (doc htmlDoc) Copy(recursive bool) Document {
	return Document(xmlDoc(doc).Copy(recursive))
}

func (doc htmlDoc) Dump() string {
	return "I'm a htmlDoc"
}

func (doc htmlDoc) AddChild(child Node) {
	xmlDoc(doc).AddChild(child)
}

func HtmlReadMemory(buffer string, size int, url string, encoding string, options int) Document {
	cbuffer := C.CString(buffer)
	curl := C.CString(url)
	cencoding := C.CString(encoding)
	cdoc := C._htmlReadMemory(cbuffer, _Ctype_int(size), curl, cencoding, _Ctype_int(options))
	C.free(unsafe.Pointer(cbuffer))
	C.free(unsafe.Pointer(curl))
	C.free(unsafe.Pointer(cencoding))
	XmlCleanupParser()
	node := xmlNode{Cdoc:cdoc}
	return Document(htmlDoc{xmlNode:node})
}

