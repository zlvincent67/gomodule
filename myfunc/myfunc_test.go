package myfunc

import (
	"fmt"
	"testing"
)

func TestMyfunc(t *testing.T) {
	sum := Myfunc(1, 2)
	expected := 3

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleMyfunc() {
	sum := Myfunc(3, 4)
	fmt.Println(sum)
	// Output: 7
}
