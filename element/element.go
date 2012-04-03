package element

/*
#cgo pkg-config: libxml-2.0
#include <stdlib.h>
#include <libxml/HTMLtree.h>
#include <libxml/tree.h>
*/
import "C"
import "unsafe"
import . "libxml"

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

type Element interface {
	String() string
	Name() string
	GetAttribute(name string) string
	NewChild(name string, content string) Element
	Child() Element
	Children() chan Element
}

type element struct {
	doc unsafe.Pointer
	ptr unsafe.Pointer
}

func ElementFromDoc(doc unsafe.Pointer) Element {
	return &element{doc: doc, ptr: doc}
}

func (e *element) NewChild(name string, content string) Element {
	newnode := XmlNewDocRawNode(e.doc, nil, name, content)
	XmlAddChild(e.ptr, newnode)
	return &element{doc: e.doc, ptr: newnode}
}

func (e *element) Child() Element {
	cur_node := XmlFirstElementChild(e.ptr)
	return &element{doc: e.doc, ptr: cur_node}
}

func (e *element) Children() chan Element {
	cur_element := XmlFirstElementChild(e.ptr)
	c := make(chan Element)
	go func(c chan Element, cur_element unsafe.Pointer) {
		for cur_element != nil {
			c <- &element{doc: e.doc, ptr: cur_element}
			cur_element = XmlNextElementSibling(cur_element)
		}
		close(c)
	}(c, cur_element)
	return c
}

func (e *element) GetAttribute(name string) string {
	return XmlGetProp(e.ptr, name)
}

func (e *element) Name() string {
	node := unsafe.Pointer(C.xmlNodePtr(e.ptr).name)
	name := C.GoString((*C.char)(node))
	return name
}

func (e *element) Type() ElementType {
	node_type := C.xmlNodePtr(e.ptr)._type
	return ElementType(node_type)
}

func (e *element) String() string {
	return XmlNodeDump(e.ptr)
}
