package libxml

import "testing"
import "strings"

func blankXmlDoc(t *testing.T) Document {
	doc := XmlNewDoc("1.0")
	if doc == nil {
		t.Fatal("XmlNewDoc returned nil")
	}
	return doc
}

func TestXmlNewDoc(t *testing.T) {
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

