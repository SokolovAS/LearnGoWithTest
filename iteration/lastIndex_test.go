package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestLastIndex(t *testing.T) {
	expected := 1
	got := strings.LastIndex("example", "x")

	if expected != got {
		t.Errorf("expected %d got %d", expected, got)
	}
}

func BenchmarkLastIndex(b *testing.B) {
	strings.LastIndex("example", "x")
}

func ExampleLastIndex() {
	lastIndex := strings.LastIndex("example", "x")
	fmt.Println(lastIndex)
	//Output:1
}
