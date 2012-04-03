package node

import "unsafe"

import . "libxml"
import . "libxml/element"

type Node interface {
	Element
	GetAttribute(name string) string
	NewChild(name string, content string) Node
}

type node struct {
	doc unsafe.Pointer
	ptr unsafe.Pointer
}

func NodeFromDoc(doc unsafe.Pointer) Node {
	return &node{doc: doc, ptr: doc}
}

func (n *node) NewChild(name string, content string) Node {
	newnode := XmlNewDocRawNode(n.doc, nil, name, content)
	XmlAddChild(n.ptr, newnode)
	return &node{doc: n.doc, ptr: newnode}
}

func (n *node) GetAttribute(name string) string {
	return XmlGetProp(n.ptr, name)
}

func (n *node) Child() Element {
	return ElementFromNode(n.doc, n.ptr).Child()
}

func (n *node) Children() chan Element {
	return ElementFromNode(n.doc, n.ptr).Children()
}

func (n *node) Name() string {
	return ElementFromNode(n.doc, n.ptr).Name()
}

func (n *node) String() string {
	return XmlNodeDump(n.ptr)
}
