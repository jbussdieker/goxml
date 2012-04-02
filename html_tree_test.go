package libxml

import "testing"

func blankHtmlDoc(t *testing.T) Document {
	doc := NewHtmlDoc("", "")
	if doc == nil {
		t.Fatal("HtmlNewDoc returned nil")
	}
	return doc
}

func TestNewHtmlDoc(t *testing.T) {
	doc := blankHtmlDoc(t)
	expectString(t, doc, "<!DOCTYPE html PUBLIC \"\" \"\">")
	doc.Free()
	checkMemory(t)
}

func TestHtmlDocAddChild(t *testing.T) {
	doc := blankHtmlDoc(t)
	doc.AddChild("html", "")
	expectString(t, doc, "<!DOCTYPE html PUBLIC \"\" \"\">\n<html></html>")
	doc.Free()
	checkMemory(t)
}

func TestHtmlDocGetRootElement(t *testing.T) {
	doc := blankHtmlDoc(t)
	doc.AddChild("html", "")
	root := doc.GetRootElement()
	expectString(t, root, "<html></html>")
	doc.Free()
	checkMemory(t)
}

func TestHtmlNodeAddChild(t *testing.T) {
	doc := blankHtmlDoc(t)
	doc.AddChild("html", "").AddChild("body", "")
	expectString(t, doc, "<!DOCTYPE html PUBLIC \"\" \"\">\n<html><body></body>\n</html>")
	doc.Free()
	checkMemory(t)
}

func TestHtmlNodeString(t *testing.T) {
	doc := blankHtmlDoc(t)
	html := doc.AddChild("html", "")
	html.AddChild("body", "")
	expectString(t, html, "<html><body></body></html>")
	doc.Free()
	checkMemory(t)
}
