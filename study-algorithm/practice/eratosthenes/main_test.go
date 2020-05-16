package eratosthenes

import "testing"

func Test_IsPrime_True(t *testing.T) {
	f := New(15)
	for _, i := range []int{2, 3, 5, 7, 11, 13} {
		if !f.IsPrime(i) {
			t.Errorf("%d is prime; but IsPrime() returns false", i)
		}
	}
}

func Test_IsPrime_False(t *testing.T) {
	f := New(15)
	for _, i := range []int{1, 4, 6, 8, 9} {
		if f.IsPrime(i) {
			t.Errorf("%d is not prime; but IsPrime() returns true", i)
		}
	}
}
