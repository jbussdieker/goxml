package libxml

import "testing"
import "strings"

func blankHtmlDoc(t *testing.T) Document {
	doc := HtmlNewDoc("", "")
	if doc == nil {
		t.Fatal("HtmlNewDoc returned nil")
	}
	return doc
}

func TestHtmlNewDoc(t *testing.T) {
	doc := blankHtmlDoc(t)
	buf := doc.Dump()
	if strings.TrimSpace(buf.String()) != "<!DOCTYPE html PUBLIC \"\" \"\">" {
		t.Fail()
	}
	buf.Free()
	doc.Free()
	checkMemory(t)
}

func TestHtmlDocAddChild(t *testing.T) {
	doc := blankHtmlDoc(t)
	doc.AddChild("html", "")
	buf := doc.Dump()
	if strings.TrimSpace(buf.String()) != "<!DOCTYPE html PUBLIC \"\" \"\">\n<html></html>" {
		t.Fail()
	}
	buf.Free()
	doc.Free()
	checkMemory(t)
}
