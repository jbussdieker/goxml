package libxml
/*
#include <stdio.h>
#include <libxml/HTMLtree.h>

htmlDocPtr _htmlNewDoc(char *uri, char *external_id) {
	return htmlNewDoc((xmlChar *)uri, (xmlChar *)external_id);
}

xmlBufferPtr _htmlDocDump(htmlDocPtr doc) {
	xmlBufferPtr buf = xmlBufferCreate();
	htmlNodeDump(buf, NULL, (xmlNodePtr)doc);
	return buf;
}
*/
import "C"
import "unsafe"

type htmlDocPtr struct {
	ptr C.htmlDocPtr
}

// Create a new blank HTML Document
func NewHtmlDoc(uri string, external_id string) Document {
	curi := C.CString(uri)
	cexternal_id := C.CString(external_id)
	cdoc := C._htmlNewDoc(curi, cexternal_id)
	C.free(unsafe.Pointer(curi))
	C.free(unsafe.Pointer(cexternal_id))
	doc := htmlDocPtr{ptr:cdoc}
	return doc
}

func (doc htmlDocPtr) AddChild(name string, content string) Node {
	// Implemented in xmlDoc
	xmldoc := xmlDocPtr{ptr:_Ctype_xmlDocPtr(doc.ptr)}
	return xmldoc.AddChild(name, content)
}

func (doc htmlDocPtr) Dump() Buffer {
	cbuf := C._htmlDocDump(doc.ptr)
	buf := xmlBufferPtr{ptr:cbuf}
	return buf
}

func (doc htmlDocPtr) GetRootElement() Node {
	// Implemented in xmlDoc
	xmldoc := xmlDocPtr{ptr:_Ctype_xmlDocPtr(doc.ptr)}
	return xmldoc.GetRootElement()
}

func (doc htmlDocPtr) String() string {
	buf := doc.Dump()
	str := buf.String()
	buf.Free()
	return str
}

func (doc htmlDocPtr) Free() {
	// Implemented in xmlDoc
	if unsafe.Pointer(doc.ptr) != nil {
		xmldoc := xmlDocPtr{ptr:_Ctype_xmlDocPtr(doc.ptr)}
		xmldoc.Free()
	} else {
		panic("Document already freed")
	}
}

