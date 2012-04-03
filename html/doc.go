package html

/*
#cgo pkg-config: libxml-2.0
#include <stdlib.h>
#include <libxml/HTMLtree.h>
*/
import "C"
import "unsafe"
import . "libxml"
import . "libxml/element"
import . "libxml/node"

type Document interface {
	GetRootElement() Element
	Node
}

type doc struct {
	ptr unsafe.Pointer
}

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

func New(uri string, external_id string) Document {
	return &doc{ptr: HtmlNewDoc(uri, external_id)}
}

func NewFromString(buffer string, url string, encoding string, options HtmlParseOptions) Document {
	return &doc{ptr: HtmlReadMemory(buffer, url, encoding, options)}
}

func (doc *doc) String() string {
	return HtmlNodeDump(doc.ptr)
}

func (doc *doc) GetRootElement() Element {
	cnode := XmlDocGetRootElement(doc.ptr)
	return ElementFromDoc(cnode)
}

func (doc *doc) GetAttribute(name string) string {
	return NodeFromDoc(doc.ptr).GetAttribute(name)
}

func (doc *doc) NewChild(name string, content string) Node {
	return NodeFromDoc(doc.ptr).NewChild(name, content)
}

func (doc *doc) Child() Element {
	return ElementFromDoc(doc.ptr).Child()
}

func (doc *doc) Children() chan Element {
	return ElementFromDoc(doc.ptr).Children()
}

func (doc *doc) Name() string {
	return ElementFromDoc(doc.ptr).Name()
}

