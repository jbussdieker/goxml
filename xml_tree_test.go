package libxml

import "testing"
import "strings"

func TestXmlNewDoc(t *testing.T) {
	doc := blankXmlDoc(t)
	doc.Free()
	checkMemory(t)
}

func TestXmlDocString(t *testing.T) {
	doc := blankXmlDoc(t)
	buf := doc.Dump()
	if strings.TrimSpace(buf.String()) != "<?xml version=\"1.0\" encoding=\"UTF-8\"?>" {
		t.Fail()
	}
	buf.Free()
	doc.Free()
	checkMemory(t)
}

func TestXmlDocAddChild(t *testing.T) {
	doc := blankXmlDoc(t)
	doc.AddChild("html", "")
	buf := doc.Dump()
	if strings.TrimSpace(buf.String()) != "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<html></html>" {
		t.Fail()
	}
	buf.Free()
	doc.Free()
	checkMemory(t)
}

func TestXmlDocFun(t *testing.T) {
	doc := blankXmlDoc(t)
	n1 := doc.AddChild("a", "test1")
	n2 := doc.AddChild("b", "test2")
	
	buf := doc.Dump()
	buf.Free()

	n1.Free()
	n2.Free()

	buf = doc.Dump()
	buf.Free()

	doc.Free()
	checkMemory(t)
}

