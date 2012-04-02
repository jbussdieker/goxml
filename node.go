package libxml

// Generic Node interface
type Node interface {
	Free()
	Dump() Buffer
	String() string
	Path() string
	AddChild(name string, content string) Node
}
