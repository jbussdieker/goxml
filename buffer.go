package libxml

type Buffer interface {
	String() string
	Free()
}
