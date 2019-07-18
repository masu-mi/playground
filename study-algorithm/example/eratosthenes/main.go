package eratosthenes

import "math"

type Field map[int]struct{}

func New(last int) Field {
	f := map[int]struct{}{}
	for i := 2; i <= last; i++ {
		f[i] = struct{}{}
	}
	l := int(math.Sqrt(float64(last)))
	for i := 2; i <= l; i++ {
		for j := 2; j <= last/i; j++ {
			delete(f, i*j)
		}
	}
	return f
}

func (f Field) IsPrime(num int) bool {
	_, ok := map[int]struct{}(f)[num]
	return ok
}
