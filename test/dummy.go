package dummy

import "testing"
import "fmt"
import . "goxml"

func stopTest(t *testing.T) {
	if AllocSize() != 0 {
		fmt.Println("Ending with", AllocSize(), "remaining allocations")
		t.Fail()
	}
}

