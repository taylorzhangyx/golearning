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

func BenchmarkMagic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Magic(111, 1111, 1111, 111231, 123, 123, 123, 1231, 24, 455, 41, 231, 45, 123, 123, 1234)
	}
}
