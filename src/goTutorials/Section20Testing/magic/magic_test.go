package magic

import (
	"fmt"
	"testing"
)

func TestMagic2(t *testing.T) {
	res := Magic(2)
	if !res {
		t.Error("2 should be the magic numer")
	}
}

func TestMagic3(t *testing.T) {
	res := Magic(3)
	if !res {
		t.Error("3 should be the magic numer")
	}
}
func TestMagic6(t *testing.T) {
	res := Magic(6)
	if res {
		t.Error("6 should not be the magic numer")
	}
}

func ExampleMagic() {
	fmt.Println(Magic(5, 6, 7))
	// Output:
	// true
}

func ExampleMagic_second() {
	fmt.Println(Magic(10))
	// Output:
	// false
}
