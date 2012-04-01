package goxml

/*
#cgo pkg-config: libxml-2.0
#cgo LDFLAGS: -lrt
#include "goxml.h"
*/
import "C"
//import "unsafe"

/*type xmlNode unsafe.Pointer
type xmlDoc unsafe.Pointer
type xmlDtd unsafe.Pointer
type _xmlAttr unsafe.Pointer
type xmlChar string*/

/*type xmlAttr interface {
	Name() xmlChar
	Children() xmlNode
	Last() xmlNode
	Parent() xmlDtd
	Next() xmlAttr
	Prev() xmlAttr
	Doc() xmlDoc
}*/

func init() {
	C.goXmlInit()
}

func AllocSize() C.long {
	return C.goXmlAllocSize()
}
