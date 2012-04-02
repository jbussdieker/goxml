package libxml

import "testing"

func parseXmlDoc(t *testing.T, buf string) Document {
	doc := NewXmlFromString(buf)
	if doc == nil {
		t.Fatal("HtmlParseDoc returned nil")
	}
	return doc
}

func TestXmlParseDoc(t *testing.T) {
	doc := parseXmlDoc(t, "<html></html>")
	expectString(t, doc.String(), "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<html/>")
	doc.Free()
	checkMemory(t)
}
