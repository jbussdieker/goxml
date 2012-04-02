package libxml

type Node interface {
	Free()
	Dump() Buffer
	String() string
	AddChild(name string, content string) Node
}
