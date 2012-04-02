package libxml

import "testing"
import "strings"

func parseXmlDoc(t *testing.T, buf string) Document {
	doc := XmlParseDoc(buf)
	if doc == nil {
		t.Fatal("HtmlParseDoc returned nil")
	}
	return doc
}

func TestXmlParseDoc(t *testing.T) {
	doc := parseXmlDoc(t, "<html></html>")
	str := strings.TrimSpace(doc.String())
	if str != "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<html/>" {
		t.Fatal(str)
		t.Fail()
	}
	doc.Free()
	checkMemory(t)
}

