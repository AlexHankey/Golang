package iteration

import "testing"

func TestRepeat(t *testing.T) {
	for i := 0; i < 10; i++ {
		Repeat("a")
	}
}
