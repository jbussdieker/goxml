package goxml

/*
#cgo pkg-config: libxml-2.0
#cgo LDFLAGS: -lrt
#include <stdlib.h>
#include <libxml/xmlmemory.h>
#include "goxml.h"

void _xmlFree(void *p) {
	xmlFree(p);
}
void _xmlFreeDoc(void *p) {
	xmlFreeDoc(p);
}
*/
import "C"
import "unsafe"

func init() {
	C.goXmlInit()
}

type Node interface {
	String() string
	Set(string)
	Free()
	Children() NodeSet
	Child() Node
}

type Document interface {
	Node
}

type xmlDocument struct {
	cstring *C.char
	cdoc unsafe.Pointer
}

type htmlDocument struct {
	cstring *C.char
	cdoc unsafe.Pointer
}

func AllocSize() int {
	return int(C.goXmlAllocSize())
}

func NewXmlDocument(input string) Document {
	doc := xmlDocument{}
	doc.cstring = C.CString(input) // TIARL
	doc.cdoc = C.goXmlRead(doc.cstring) // TIARL
	return Document(doc)
}

func NewHtmlDocument(value string) Document {
	doc := htmlDocument{}
	doc.cstring = C.CString(value) // TIARL
	doc.cdoc = C.goHtmlRead(doc.cstring) // TIARL
	return Document(doc)
}

func (doc htmlDocument) Free() {
	C._xmlFreeDoc(doc.cdoc) // TIARL-FREE
	C.free(unsafe.Pointer(doc.cstring)) // TIARL-FREE
}

func (doc htmlDocument) String() string {
	str := C.goHtmlDumpDoc(doc.cdoc) // TIARL
	s := C.GoString(str)
	C._xmlFree(unsafe.Pointer(str)) // TIARL-FREE
	return s
}

func (doc htmlDocument) Set(value string) {
	C._xmlFreeDoc(doc.cdoc) // TIARL-FREE
	C.free(unsafe.Pointer(doc.cstring)) // TIARL-FREE
	doc.cstring = C.CString(value) // TIARL
	doc.cdoc = C.goHtmlRead(doc.cstring) // TIARL
}

func (doc htmlDocument) Children() NodeSet {
	children := C.goXmlChildren(doc.cdoc)
	ns := aNodeSet{}
	ns.cnodes = children
	return NodeSet(ns)
}

func (doc htmlDocument) Child() Node {
	node := xmlDocument{}
	node.cdoc = C.goXmlChild(doc.cdoc)
	return node
}

func (doc xmlDocument) Free() {
	C._xmlFreeDoc(doc.cdoc) // TIARL-FREE
	C.free(unsafe.Pointer(doc.cstring)) // TIARL-FREE
}

func (doc xmlDocument) String() string {
	str := C.goXmlDumpDoc(doc.cdoc) // TIARL
	s := C.GoString(str)
	C._xmlFree(unsafe.Pointer(str)) // TIARL-FREE
	return s
}

func (doc xmlDocument) Set(value string) {
}

func (doc xmlDocument) Children() NodeSet {
	return nil
}

func (doc xmlDocument) Child() Node {
	return nil
}
