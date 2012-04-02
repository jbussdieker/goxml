package libxml

// Represents a buffer allocated in libxml. These must be freed if returned.
type Buffer interface {
	String() string
	Free()
}

