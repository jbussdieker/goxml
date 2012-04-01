package goxml

type Node interface {
	Free()
	Copy(recursive bool) Document
}

type Document interface {
	AddChild(child Node)
	Node
	Debug()
}
