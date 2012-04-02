package libxml

/*
#include <stdio.h>
#include <libxml/tree.h>

xmlDocPtr  _xmlNewDoc(char *version) { return xmlNewDoc((xmlChar *)version); }
xmlNodePtr _xmlDocGetRootElement(xmlDocPtr doc)   { return xmlDocGetRootElement(doc); }
void       _xmlFree(void *obj) { xmlFree(obj); }
void       _xmlFreeDoc(xmlDocPtr doc) { xmlFreeDoc(doc); }
void       _xmlFreeNode(xmlNodePtr node) { xmlFreeNode(node); }
void       _xmlUnlinkNode(xmlNodePtr node) { xmlUnlinkNode(node); }
void       _xmlBufferFree(xmlBufferPtr buf) { xmlBufferFree(buf); }
char       *_xmlBufferContent(xmlBufferPtr buf) { return (char *)xmlBufferContent(buf); }
char       *_xmlGetNodePath(xmlNodePtr node) { return (char *)xmlGetNodePath(node); }
xmlNodePtr _xmlNextElementSibling(xmlNodePtr node) { return xmlNextElementSibling(node); }
xmlNodePtr _xmlDocToNode(xmlDocPtr doc) { return (xmlNodePtr)doc; }
xmlNodePtr _xmlFirstElementChild(xmlNodePtr node) { return xmlFirstElementChild(node); }
char       *_xmlGetNodeName(xmlNodePtr node) { return (char *)node->name; }
xmlElementType _xmlGetNodeType(xmlNodePtr node) { return node->type; }

xmlBufferPtr _xmlDocDump(xmlDocPtr doc) {
	xmlBufferPtr buf = xmlBufferCreate();
	xmlNodeDump(buf, NULL, (xmlNodePtr)doc, 0, 0);
	return buf;
}

xmlBufferPtr _xmlNodeDump(xmlNodePtr node) {
	xmlBufferPtr buf = xmlBufferCreate();
	xmlNodeDump(buf, NULL, node, 0, 0);
	return buf;
}

xmlNodePtr _xmlAddDocChild(xmlDocPtr doc, char *name, char *content) {
	xmlNodePtr node;
	node = xmlNewDocRawNode(doc, NULL, (xmlChar *)name, (xmlChar *)content);
	xmlAddChild((xmlNodePtr)doc, node);
	return node;
}

xmlNodePtr _xmlAddNodeChild(xmlNodePtr node, char *name, char *content) {
	xmlNodePtr newnode;
	newnode = xmlNewDocRawNode(node->doc, NULL, (xmlChar *)name, (xmlChar *)content);
	xmlAddChild((xmlNodePtr)node, newnode);
	return newnode;
}
*/
import "C"
import "unsafe"

const (
	XML_ELEMENT_NODE   = C.XML_ELEMENT_NODE
	XML_ATTRIBUTE_NODE = iota
	XML_TEXT_NODE
	XML_CDATA_SECTION_NODE
	XML_ENTITY_REF_NODE
	XML_ENTITY_NODE
	XML_PI_NODE
	XML_COMMENT_NODE
	XML_DOCUMENT_NODE
	XML_DOCUMENT_TYPE_NODE
	XML_DOCUMENT_FRAG_NODE
	XML_NOTATION_NODE
	XML_HTML_DOCUMENT_NODE
	XML_DTD_NODE
	XML_ELEMENT_DECL
	XML_ATTRIBUTE_DECL
	XML_ENTITY_DECL
	XML_NAMESPACE_DECL
	XML_XINCLUDE_START
	XML_XINCLUDE_END
	XML_DOCB_DOCUMENT_NODE
)

type xmlDocPtr struct {
	ptr C.xmlDocPtr
}

type xmlNodePtr struct {
	ptr C.xmlNodePtr
}

type xmlBufferPtr struct {
	ptr C.xmlBufferPtr
}

// Creates a new blank XML Document
func NewXmlDoc(version string) Document {
	cversion := C.CString(version)
	cdoc := C._xmlNewDoc(cversion)
	doc := &xmlDocPtr{ptr: cdoc}
	return doc
}

func (doc *xmlDocPtr) Free() {
	if doc.ptr != nil {
		C._xmlFreeDoc(doc.ptr)
		doc.ptr = nil
	} else {
		panic("xmlDoc already freed")
	}
}

func (doc *xmlDocPtr) Parse(buffer string) {
}

func (node *xmlNodePtr) AddChild(name string, content string) Node {
	cname := C.CString(name)
	ccontent := C.CString(content)
	cnode := C._xmlAddNodeChild(node.ptr, cname, ccontent)
	C.free(unsafe.Pointer(cname))
	C.free(unsafe.Pointer(ccontent))
	return &xmlNodePtr{ptr: cnode}
}

func (node *xmlNodePtr) Dump() Buffer {
	cbuf := C._xmlNodeDump(node.ptr)
	buf := &xmlBufferPtr{ptr: cbuf}
	return buf
}

func (node *xmlNodePtr) String() string {
	buf := node.Dump()
	str := buf.String()
	buf.Free()
	return str
}

func (node *xmlNodePtr) Name() string {
	cstr := C._xmlGetNodeName(node.ptr)
	str := C.GoString(cstr)
	return str
}

func (node *xmlNodePtr) Type() int {
	return int(C._xmlGetNodeType(node.ptr))
}

func (node *xmlNodePtr) Path() string {
	cpath := C._xmlGetNodePath(node.ptr)
	str := C.GoString(cpath)
	C._xmlFree(unsafe.Pointer(cpath))
	return str
}

func (node *xmlNodePtr) Children() chan Node {
	cnode := C._xmlFirstElementChild(node.ptr)
	c := make(chan Node)
	go func(c chan Node, cnode C.xmlNodePtr) {
		for cnode != nil {
			c <- &xmlNodePtr{ptr: cnode}
			cnode = C._xmlNextElementSibling(cnode)
		}
		close(c)
	}(c, cnode)
	return c
}

func (node *xmlNodePtr) Free() {
	if node.ptr != nil {
		C._xmlUnlinkNode(node.ptr)
		C._xmlFreeNode(node.ptr)
		node.ptr = nil
	} else {
		panic("Node already freed")
	}
}

func (doc *xmlDocPtr) AddChild(name string, content string) Node {
	cname := C.CString(name)
	ccontent := C.CString(content)
	cnode := C._xmlAddDocChild(doc.ptr, cname, ccontent)
	C.free(unsafe.Pointer(cname))
	C.free(unsafe.Pointer(ccontent))
	return &xmlNodePtr{ptr: cnode}
}

func (doc *xmlDocPtr) GetRootElement() Node {
	cnode := C._xmlDocGetRootElement(doc.ptr)
	return &xmlNodePtr{ptr: cnode}
}

func (doc *xmlDocPtr) Dump() Buffer {
	cbuf := C._xmlDocDump(doc.ptr)
	buf := &xmlBufferPtr{ptr: cbuf}
	return buf
}

func (doc *xmlDocPtr) Children() chan Node {
	return doc.Node().Children()
}

func (doc *xmlDocPtr) String() string {
	buf := doc.Dump()
	str := buf.String()
	buf.Free()
	return str
}

func (doc *xmlDocPtr) Node() Node {
	cnode := C._xmlDocToNode(doc.ptr)
	return &xmlNodePtr{ptr: cnode}
}

func (doc *xmlDocPtr) Path() string {
	return doc.Node().Path()
}

func (doc *xmlDocPtr) Name() string {
	return doc.Node().Name()
}

func (doc *xmlDocPtr) Type() int {
	return doc.Node().Type()
}

func (buf *xmlBufferPtr) Free() {
	if buf.ptr != nil {
		C._xmlBufferFree(buf.ptr)
		buf.ptr = nil
	} else {
		panic("Buffer already freed")
	}
}

func (buf *xmlBufferPtr) String() string {
	cstr := C._xmlBufferContent(buf.ptr)
	str := C.GoString(cstr)
	return str
}
