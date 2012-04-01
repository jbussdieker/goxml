package goxml
/*
#include <stdlib.h>
#include <libxml/tree.h>
#include <libxml/parser.h>
#include <libxml/HTMLtree.h>
#include <libxml/HTMLparser.h>

void *_xmlNewDoc(char *version) {
	return xmlNewDoc((xmlChar *)version);
}

void _xmlFreeDoc(void *doc) {
	return xmlFreeDoc((xmlDocPtr)doc);
}

void *_xmlCopyDoc(void *doc, int recursive) {
	return xmlCopyDoc((xmlDocPtr)doc, recursive);
}

void *_xmlReadMemory(char *buffer, int size, char *url, char *encoding, int options) {
	return xmlReadMemory(buffer, size, url, encoding, options);
}

void *_xmlNewNode(char *name) {
	return xmlNewNode(NULL, (xmlChar *)name);
}

void *_xmlAddChild(void *node, void *child_node) {
	return xmlAddChild(node, child_node);
}

void _xmlFreeNode(void *node) {
	return xmlFreeNode((xmlNodePtr)node);
}

void _xmlCleanupParser() {
	xmlCleanupParser();
}

void _xmlDebugDoc(void *doc) {
	xmlDocPtr xmldoc = (xmlDocPtr)doc;
	fprintf(stderr, "Doc Type:       %d\n", xmldoc->type);
	fprintf(stderr, "Doc_Properties: %d\n", xmldoc->properties);
}

*/
import "C"
import "unsafe"

const (
	XML_PARSE_RECOVER    = C.XML_PARSE_RECOVER    //: recover on errors
	XML_PARSE_NOENT      = C.XML_PARSE_NOENT      //: substitute entities
	XML_PARSE_DTDLOAD    = C.XML_PARSE_DTDLOAD    //: load the external subset
	XML_PARSE_DTDATTR    = C.XML_PARSE_DTDATTR    //: default DTD attributes
	XML_PARSE_DTDVALID   = C.XML_PARSE_DTDVALID   //: validate with the DTD
	XML_PARSE_NOERROR    = C.XML_PARSE_NOERROR    //: suppress error reports
	XML_PARSE_NOWARNING  = C.XML_PARSE_NOWARNING  //: suppress warning reports
	XML_PARSE_PEDANTIC   = C.XML_PARSE_PEDANTIC   //: pedantic error reporting
	XML_PARSE_NOBLANKS   = C.XML_PARSE_NOBLANKS   //: remove blank nodes
	XML_PARSE_SAX1       = C.XML_PARSE_SAX1       //: use the SAX1 interface internally
	XML_PARSE_XINCLUDE   = C.XML_PARSE_XINCLUDE   //: Implement XInclude substitition
	XML_PARSE_NONET      = C.XML_PARSE_NONET      //: Forbid network access
	XML_PARSE_NODICT     = C.XML_PARSE_NODICT     //: Do not reuse the context dictionnary
	XML_PARSE_NSCLEAN    = C.XML_PARSE_NSCLEAN    //: remove redundant namespaces declarations
	XML_PARSE_NOCDATA    = C.XML_PARSE_NOCDATA    //: merge CDATA as text nodes
	XML_PARSE_NOXINCNODE = C.XML_PARSE_NOXINCNODE //: do not generate XINCLUDE START/END nodes
	XML_PARSE_COMPACT    = C.XML_PARSE_COMPACT    //: compact small text nodes; no modification of the tree allowed afterwards (will possibly crash if you try to modify the tree)
	XML_PARSE_OLD10      = C.XML_PARSE_OLD10      //: parse using XML-1.0 before update 5
	XML_PARSE_NOBASEFIX  = C.XML_PARSE_NOBASEFIX  //: do not fixup XINCLUDE xml:base uris
	XML_PARSE_HUGE       = C.XML_PARSE_HUGE       //: relax any hardcoded limit from the parser
	XML_PARSE_OLDSAX     = C.XML_PARSE_OLDSAX     //: parse using SAX2 interface from before 2.7.0
)

type xmlNode struct {
	Cdoc unsafe.Pointer
	Cnode unsafe.Pointer
}

type xmlDoc struct {
	xmlNode
}

func XmlNewDoc(version string) Document {
	cversion := C.CString(version)
	cdoc := C._xmlNewDoc(cversion)
	C.free(unsafe.Pointer(cversion))
	node := xmlNode{Cdoc:cdoc}
	return Document(xmlDoc{xmlNode:node})
}

func XmlNewNode(name string) Node {
	cname := C.CString(name)
	cnode := C._xmlNewNode(cname)
	C.free(unsafe.Pointer(cname))
	node := xmlNode{Cnode:cnode}
	return Node(node)
}

func (node xmlNode) Copy(extended bool) Document {
	panic("NOPE!")
}

func (node xmlNode) Free() {
	C._xmlFreeNode(node.Cnode)
}

func XmlReadMemory(buffer string, size int, url string, encoding string, options int) Document {
	cbuffer := C.CString(buffer)
	curl := C.CString(url)
	cencoding := C.CString(encoding)
	cdoc := C._xmlReadMemory(cbuffer, _Ctype_int(size), curl, cencoding, _Ctype_int(options))
	C.free(unsafe.Pointer(cbuffer))
	C.free(unsafe.Pointer(curl))
	C.free(unsafe.Pointer(cencoding))
	XmlCleanupParser()
	node := xmlNode{Cdoc:cdoc}
	return Document(xmlDoc{xmlNode:node})
}

func XmlCleanupParser() {
	C._xmlCleanupParser()
}

func (doc xmlDoc) Free() {
	C._xmlFreeDoc(doc.Cdoc)
}

func (doc xmlDoc) Dump() string {
	return "I'm a xmlDoc"
}

func (doc xmlDoc) Debug() {
	C._xmlDebugDoc(doc.Cdoc)
}

func (doc xmlDoc) Copy(recursive bool) Document {
	if recursive {
		cdoc := C._xmlCopyDoc(doc.Cdoc, 1)
		node := xmlNode{Cdoc:cdoc}
		return Document(xmlDoc{xmlNode:node})
	}
	cdoc := C._xmlCopyDoc(doc.Cdoc, 0)
	node := xmlNode{Cdoc:cdoc}
	return Document(xmlDoc{xmlNode:node})
}

func (doc xmlDoc) AddChild(child Node) {
	cn := child.(xmlNode)
	C._xmlAddChild(doc.Cdoc, cn.Cnode)
}

