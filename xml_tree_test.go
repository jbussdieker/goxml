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
	str := strings.TrimSpace(doc.String())
	if str != "<?xml version=\"1.0\" encoding=\"UTF-8\"?>" {
		t.Fail()
	}
	doc.Free()
	checkMemory(t)
}

func TestXmlDocAddChild(t *testing.T) {
	doc := blankXmlDoc(t)
	doc.AddChild("html", "")
	str := strings.TrimSpace(doc.String())
	if str != "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<html></html>" {
		t.Fail()
	}
	doc.Free()
	checkMemory(t)
}

