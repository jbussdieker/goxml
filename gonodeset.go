package goxml

/*
#include <libxml/tree.h>
#include <libxml/xpath.h>
int nsLength(void *node_set) {
	if (node_set)
    	return ((xmlNodeSetPtr)node_set)->nodeNr;
	return 0;
}
*/
import "C"
import "unsafe"

type NodeSet interface {
	Length() int
}

type aNodeSet struct {
	cnodes unsafe.Pointer
}

func (ns aNodeSet) Length() int {
	return int(C.nsLength(ns.cnodes))
}

