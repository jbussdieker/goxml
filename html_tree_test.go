package libxml

import "testing"
import "strings"

func TestHtmlNewDoc(t *testing.T) {
	doc := blankHtmlDoc(t)
	doc.Free()
	checkMemory(t)
}

func TestHtmlDocString(t *testing.T) {
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
