package dummy

import "testing"
import . "goxml"

func TestXmlDocument(t *testing.T) {
	d := XmlNewDoc("1.0")
	d.Debug()
	d.Free()
	stopTest(t)
}

func TestXmlNewNode(t *testing.T) {
	n := XmlNewNode("a")
	n.Free()
	stopTest(t)
}

func TestXmlCopyDocument(t *testing.T) {
	d := XmlNewDoc("1.0")
	dd := d.Copy(false)
	d.Free()
	dd.Free()

	d = XmlNewDoc("1.0")
	dd = d.Copy(true)
	d.Free()
	dd.Free()
}

func TestXmlReadMemory(t *testing.T) {
	d := XmlReadMemory(" ", 1, "/", "UTF-8", XML_PARSE_NOERROR)
	d.Free()
	stopTest(t)
}

