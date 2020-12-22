package main

import (
	"github.com/k0kubun/pp"
	"gonum.org/v1/gonum/mat"
)

func main() {
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	A := mat.NewDense(3, 4, x)
	pp.Println(A)
	pp.Println(A.At(0, 1))
	A.Set(0, 1, 10000.0)
	pp.Println(A.At(0, 1))
	pp.Println(A.RowView(1))
	pp.Println(A.Slice(0, 3, 0, 3))
	B := mat.NewDense(3, 3, nil)
	B.Inverse(A.Slice(0, 3, 0, 3))
	B.Scale(10, B)
	pp.Println(B)
	C := B.T()
	pp.Println(C)
	B.Set(1, 1, 100000)
	pp.Println(C)
	// sigmoid := func(i, j int, v float64) float64 {
	// 	return 1 / (1 + math.Exp(-v))
	// }
	// result := &mat.Dense{}
	// result.Apply(sigmoid, A)
	// pp.Println(result)
	// ⎡1  2  3  4⎤
	// ⎢5  6  7  8⎥
	// ⎣9 10 11 12⎦
}
