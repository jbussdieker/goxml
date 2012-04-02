package libxml

import "testing"

func parseHtmlDoc(t *testing.T, buf string) Document {
	doc := NewHtmlFromString(buf)
	if doc == nil {
		t.Fatal("HtmlParseDoc returned nil")
	}
	return doc
}

func TestHtmlParseDoc(t *testing.T) {
	doc := parseHtmlDoc(t, "<html></html>")
	expectString(t, doc, "<!DOCTYPE html PUBLIC \"-//W3C//DTD HTML 4.0 Transitional//EN\" \"http://www.w3.org/TR/REC-html40/loose.dtd\">\n<html></html>")
	doc.Free()
	checkMemory(t)
}
