package libxml

import "testing"
import "strings"
import "fmt"

func recurseTree(depth int, node Node) (found bool) {
	for n := range node.Children() {
		found = true
		fmt.Println(strings.Repeat("  ", depth), n.Name())
		if recurseTree(depth+1, n) {
			fmt.Println(strings.Repeat("  ", depth), "/"+n.Name())
		}
	}
	return
}

func TestTree(t *testing.T) {
	doc := NewXmlFromString("<html><head><title></title></head><body><div><span></span></div><div><span></span></div></body></html>")
	recurseTree(1, doc)
	doc.Free()
}

func TestAttributes(t *testing.T) {
	doc := NewXmlFromString("<html test=\"yup\" test1=\"yup1\"></html>")
	recurseTree(1, doc)
	root := doc.GetRootElement()
	for attr := range root.Attributes() {
		println(attr.Name())
	}
	doc.Free()
}