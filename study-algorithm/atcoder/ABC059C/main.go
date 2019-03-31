package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type sign int

const (
	minus sign = -1 + iota
	undefined
	plus
)

type state struct {
	totalOps, currentSum int
	preSign              sign
}

func main() {
	var n int
	fmt.Scan(&n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	var stateStartsPlus, stateStartsMinus state
	stateStartsPlus.preSign = minus
	stateStartsMinus.preSign = plus
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		stateStartsPlus = nextState(stateStartsPlus, a)
		stateStartsMinus = nextState(stateStartsMinus, a)
	}
	var minOps int
	if stateStartsPlus.totalOps < stateStartsMinus.totalOps {
		minOps = stateStartsPlus.totalOps
	} else {
		minOps = stateStartsMinus.totalOps
	}
	fmt.Printf("%d\n", minOps)
}

func nextState(s state, val int) state {
	s.currentSum += val
	var ops int
	ops, s.currentSum, s.preSign = adjust(s.currentSum, s.preSign)
	s.totalOps += ops
	return s
}

func adjust(sum int, preSign sign) (ops, fixedSum int, nextSign sign) {
	switch preSign {
	case plus:
		nextSign = minus
		if sum < 0 {
			fixedSum = sum
			return
		}
		ops = sum + 1
		fixedSum = -1
	case minus:
		nextSign = plus
		if sum > 0 {
			fixedSum = sum
			return
		}
		ops = -sum + 1
		fixedSum = 1
	}
	return
}
