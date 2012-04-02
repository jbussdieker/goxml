package libxml

type Node interface {
	Free()
	Dump() Buffer
	AddChild(name string, content string) Node
}
