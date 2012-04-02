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
	str := strings.TrimSpace(doc.String())
	if str != "<!DOCTYPE html PUBLIC \"\" \"\">" {
		t.Fail()
	}
	doc.Free()
	checkMemory(t)
}

func TestHtmlDocAddChild(t *testing.T) {
	doc := blankHtmlDoc(t)
	doc.AddChild("html", "")
	str := strings.TrimSpace(doc.String())
	if str != "<!DOCTYPE html PUBLIC \"\" \"\">\n<html></html>" {
		t.Fail()
	}
	doc.Free()
	checkMemory(t)
}
