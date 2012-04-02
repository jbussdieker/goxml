package libxml

import "testing"
import "strings"

func TestHtmlParseDoc(t *testing.T) {
	doc := parseHtmlDoc(t, "<html></html>")
	doc.Free()
	checkMemory(t)
}

func TestHtmlParseDocString(t *testing.T) {
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
