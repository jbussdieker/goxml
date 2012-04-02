package libxml

// Generic Node interface
type Node interface {
	Free()
	Dump() Buffer
	String() string
	Name() string
	Type() int
	Path() string
	AddChild(name string, content string) Node
	Children() chan Node
}
