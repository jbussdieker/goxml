package dummy

import "testing"
import . "goxml"

func TestHtmlDocument(t *testing.T) {
	d := HtmlNewDoc("/", "")
	d.Free()
	stopTest(t)
}

func TestHtmlCopyDocument(t *testing.T) {
	d := HtmlNewDoc("/", "")
	dd := d.Copy(false)
	d.Free()
	dd.Free()

	d = HtmlNewDoc("/", "")
	dd = d.Copy(true)
	d.Free()
	dd.Free()
}

func TestHtmlAddChild(t *testing.T) {
	d := HtmlNewDoc("/", "")
	n := XmlNewNode("test")
	d.AddChild(n)
	d.Free()
	stopTest(t)
}

func TestHtmlReadMemory(t *testing.T) {
	d := HtmlReadMemory(" ", 1, "/", "UTF-8", HTML_PARSE_NOIMPLIED | HTML_PARSE_NOERROR)
	d.Free()
	stopTest(t)
}

