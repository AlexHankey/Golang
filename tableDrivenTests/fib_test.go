package fib

import "testing"

func TestFib(t *testing.T) {
	for _, tt := range fibTests {
		actual := Fib(tt.expected)
		if actual != tt.expected {
			t.Errorf("Fib(%d): expected %d, actual %d",tt.n, tt.expected, actual)
		}
	}
}

//In this example we range over all the fibTests defined above,
//assigning their value in turn to tt.
//We then call Fib passing in the value of tt.n and compare the result, stored in actual, with the value of tt.expected.
//
//The use of the names actual and expected show my JUnit heritage, others may prefer names like want and got. You should choose something that works for you and gives a clear meaning in your test code.
//
//The use of t.Errorf instead of t.Fatalf is a personal preference. As Fib is a pure function it is safe to continue the loop after a failure. I find this generally reduces test whack-a-mole by returning all the failures at once.