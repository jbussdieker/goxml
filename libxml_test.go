package libxml

import "testing"
import "strings"
import "fmt"

func init() {
	fmt.Println("Testing libxml ", LIBXML_VERSION, "using parser version", LIBXML_PARSER_VERSION)
}

func checkMemory(t *testing.T) {
	if MemAllocSize() != 0 {
		t.Fatal(MemAllocSize(), "remaining unfreed objects")
	}
}

func expectString(t *testing.T, str string, expected string) {
	str = strings.TrimSpace(str)
	if str != expected {
		fmt.Println("**GOT*******************************")
		fmt.Println(str)
		fmt.Println("**EXPECTED**************************")
		fmt.Println(expected)
		fmt.Println("************************************")
		t.Fail()
	}
}
