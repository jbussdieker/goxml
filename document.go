package libxml

// Generic document interface
type Document interface {
	Node
	Node() Node
	GetRootElement() Node
}
