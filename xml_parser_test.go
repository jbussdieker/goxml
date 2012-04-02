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
	buf := doc.Dump()
	if strings.TrimSpace(buf.String()) != "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<html/>" {
		t.Fatal(buf.String())
		t.Fail()
	}
	buf.Free()
	doc.Free()
	checkMemory(t)
}

