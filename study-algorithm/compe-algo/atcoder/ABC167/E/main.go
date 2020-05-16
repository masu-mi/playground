package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("%d\n", resolve(parseProblem(os.Stdin)))
}

const modulo = 998244353

func parseProblem(r io.Reader) (int, int, int) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n, m, k := scanInt(sc), scanInt(sc), scanInt(sc)
	return n, m, k
}

func resolve(n, m, k int) int {
	moduloCombiInit(modulo)
	result := 0
	for i := 0; i <= k; i++ {
		result = moduloAdd(
			result,
			moduloMul(
				moduloMul(m, moduloCombi(n-1, i, modulo), modulo),
				moduloPow(m-1, n-1-i, modulo),
				modulo,
			),
			modulo,
		)
		// fmt.Printf("%d, %d, %d, %d\n", i, p, cp, result)
		// fmt.Printf("%d, %d, %d, %d, %d\n", i, p, cp, com, result)
	}
	return result
}

// snip-scan-funcs
func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}
func scanString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func moduloAdd(a, b, modulo int) int {
	result := a%modulo + b%modulo
	if result < 0 {
		result += modulo
	}
	return result % modulo
}

func moduloSub(a, b, modulo int) int {
	result := a%modulo - b%modulo
	if result < 0 {
		result += modulo
	}
	return result % modulo
}

func moduloMul(a, b, modulo int) int {
	return a % modulo * b % modulo
}

func moduloDiv(a, b, modulo int) int {
	return a % moduloInv(b, modulo) % modulo
}

func moduloInv(a, modulo int) int {
	b := modulo
	u, v := 1, 0
	for b > 0 {
		t := a / b
		a, b = b, a-t*b
		u, v = v, u-t*v
	}
	u %= modulo
	if u < 0 {
		u += modulo
	}
	return u
}

func moduloPow(a, b, modulo int) int {
	if a == 1 {
		return 1
	}
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = res * a % modulo
		}
		a = a * a % modulo
		b >>= 1
	}
	return res
}

const length = 510000

var (
	_fac  = map[int][]int{}
	_finv = map[int][]int{}
	_inv  = map[int][]int{}
)

func moduloCombiInit(modulo int) {
	fac := make([]int, length)
	finv := make([]int, length)
	inv := make([]int, length)

	defer func() {
		_fac[modulo] = fac
		_finv[modulo] = finv
		_inv[modulo] = inv
	}()

	fac[0], fac[1] = 1, 1
	finv[0], finv[1] = 1, 1
	inv[1] = 1
	for i := 2; i < length; i++ {
		fac[i] = fac[i-1] * i % modulo
		inv[i] = modulo - inv[modulo%i]*(modulo/i)%modulo
		finv[i] = finv[i-1] * inv[i] % modulo
	}
}

func moduloCombi(n, k, modulo int) int {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	return _fac[modulo][n] * (_finv[modulo][k] * _finv[modulo][n-k] % modulo) % modulo
}

func moduloLog(a, b, modulo int) int {
	// log_a(b)
	a %= modulo
	b %= modulo
	m := int(math.Sqrt(float64(modulo)))

	// basy step
	values := map[int]int{}
	val := 1
	for i := 0; i < m+2; i++ {
		if _, ok := values[val]; !ok {
			values[val] = i
		}
		val = moduloMul(val, a, modulo)
	}

	// giant step
	compound := moduloInv(moduloPow(a, m, modulo), modulo)
	val = b
	for i := 0; i < m+2; i++ {
		if l, ok := values[val]; ok {
			return (i*m%modulo + l) % modulo
		}
		val = moduloMul(val, compound, modulo)
	}
	return -1
}
