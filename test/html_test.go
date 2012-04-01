package dummy

import "testing"
import "fmt"
import . "goxml"

func stopTest(t *testing.T) {
	if AllocSize() != 9 {
		fmt.Println("Ending with", AllocSize(), "remaining allocations")
		t.Fail()
	}
}

func TestHtmlDocument(t *testing.T) {
	d := NewHtmlDocument("<html>Test Html</html>")
	//fmt.Println(d.String())
	d.Free()
	stopTest(t)
}

func TestHtmlDocumentSet(t *testing.T) {
	d := NewHtmlDocument("<html>Test 1</html>")
	//fmt.Println(d.String())
	d.Set("<html>Test 1 Set</html>")
	//fmt.Println(d.String())
	d.Free()
	stopTest(t)
}

func TestHtmlDocumentChildren(t *testing.T) {
	d := NewHtmlDocument("<html><b>Test 1</b><span>Test 2</span></html>")
	d.Children()
	//fmt.Println(children.Length(), "children")
	d.Free()
	stopTest(t)
}

func TestHtmlDocumentChild(t *testing.T) {
	d := NewHtmlDocument("asdf")
	//fmt.Println(d.String())
	d.Child()
	//fmt.Println("Child is: ", child.String(), "BAMF")
	d.Free()
	stopTest(t)
}
