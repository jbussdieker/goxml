package dummy

import "testing"
//import "fmt"
import . "goxml"

func TestXmlDocument(t *testing.T) {
	d := NewXmlDocument("<html>Test Xml</html>")
	//fmt.Println(d.String())
	d.Free()
	stopTest(t)
}
