package libxml

import "testing"
import "fmt"

func init() {
	fmt.Println("Testing libxml ", LIBXML_VERSION, "using parser version", LIBXML_PARSER_VERSION)
}

func checkMemory(t *testing.T) {
	if AllocSize() != 0 {
		t.Fatal(AllocSize(), "remaining unfreed objects")
	}
}

