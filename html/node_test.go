package html

import "testing"

func TestNodeString(t *testing.T) {
	if New("", "").GetRootElement().String() != "" {
		t.Fail()
	}
}

func TestNodeAddChild(t *testing.T) {
	doc := New("", "")
	if doc.GetRootElement().String() != "" {
		t.Fail()
	}
	n := doc.NewChild("test", "ing")
	if n.Name() != "test" {
		t.Fail()
	}
	if doc.GetRootElement().String() != "<test>ing</test>" {
		t.Fail()
	}
}

func TestNodeChildren(t *testing.T) {
	doc := NewFromString("<html><body><a></a><b></b><p></p></body></html>", "", "UTF-8", 0)
	p := doc.Child().Child()
	collect := ""
	for n := range p.Children() {
		collect += n.String()
	}
	if collect != "<a/><b/><p/>" {
		t.Fatal(collect)
	}
}
