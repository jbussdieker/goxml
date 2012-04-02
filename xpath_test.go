package libxml

import "testing"

func TestXpath(t *testing.T) {
	xpath := XmlXpathCompile("/html")
	xpath.Free()
	checkMemory(t)
}
