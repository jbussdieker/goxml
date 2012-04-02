package libxml

import "testing"

func blankXmlDoc(t *testing.T) Document {
	doc := NewXmlDoc("1.0")
	if doc == nil {
		t.Fatal("XmlNewDoc returned nil")
	}
	return doc
}

func TestNewXmlDoc(t *testing.T) {
	doc := blankXmlDoc(t)
	expectString(t, doc, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	doc.Free()
	checkMemory(t)
}

func TestXmlDocAddChild(t *testing.T) {
	doc := blankXmlDoc(t)
	doc.AddChild("html", "")
	expectString(t, doc, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<html></html>")
	doc.Free()
	checkMemory(t)
}

func TestXmlNodeAddChild(t *testing.T) {
	doc := blankXmlDoc(t)
	doc.AddChild("html", "").AddChild("body", "")
	expectString(t, doc, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<html><body></body></html>")
	doc.Free()
	checkMemory(t)
}

func TestXmlNodeString(t *testing.T) {
	doc := blankXmlDoc(t)
	html := doc.AddChild("html", "")
	html.AddChild("body", "")
	expectString(t, html, "<html><body></body></html>")
	doc.Free()
	checkMemory(t)
}

