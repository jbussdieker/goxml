package libxml

// Generic document interface
type Document interface {
	Node
	GetRootElement() Node
}
