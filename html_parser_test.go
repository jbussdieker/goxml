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
	str := strings.TrimSpace(doc.String())
	if str != "<!DOCTYPE html PUBLIC \"-//W3C//DTD HTML 4.0 Transitional//EN\" \"http://www.w3.org/TR/REC-html40/loose.dtd\">\n<html></html>" {
		t.Fatal(str)
		t.Fail()
	}
	doc.Free()
	checkMemory(t)
}

