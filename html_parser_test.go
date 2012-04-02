package libxml

import "testing"
import "strings"

func parseHtmlDoc(t *testing.T, buf string) Document {
	doc := HtmlParseDoc(buf)
	if doc == nil {
		t.Fatal("HtmlParseDoc returned nil")
	}
	return doc
}

func TestHtmlParseDoc(t *testing.T) {
	doc := parseHtmlDoc(t, "<html></html>")
	buf := doc.Dump()
	if strings.TrimSpace(buf.String()) != "<!DOCTYPE html PUBLIC \"-//W3C//DTD HTML 4.0 Transitional//EN\" \"http://www.w3.org/TR/REC-html40/loose.dtd\">\n<html></html>" {
		t.Fatal(buf.String())
		t.Fail()
	}
	buf.Free()
	doc.Free()
	checkMemory(t)
}

