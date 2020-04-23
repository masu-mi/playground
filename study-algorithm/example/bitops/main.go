package main

import "fmt"

func main() {
	fmt.Println("int16, uint16, shift")
	for i := int16(-3); i < 2; i++ {
		showIntUint(i)
	}
	fmt.Println("^")
	var i int16 = 1
	fmt.Printf("a           : b%016b\n", i)
	fmt.Printf("^a          : b%016b\n", ^i)
	fmt.Printf("^a          : b%016b\n", ^uint16(i))
	fmt.Printf("^a          : b%016b\n", uint16(^i))
	fmt.Printf("^a+1        : b%016b\n", uint16(^i+1))
	fmt.Printf("^a+1        : b%016b\n", (^i + 1))
	fmt.Printf("-a          : b%016b\n", -i)
	for i := range bitCombination(3) {
		fmt.Printf("%03b\n", i)
	}

	vals := []int{0, 1, 3}

	for i := range bitCombinationOverSubsets(vals...) {
		fmt.Printf("%04b\n", i)
	}

	for i := range bitCombinationWithSize(5, 3) {
		fmt.Printf("%05b\n", i)
	}
}

func bitCombination(num int) chan uint {
	ch := make(chan uint)
	go func() {
		defer close(ch)
		for i := 0; i < 1<<uint(num); i++ {
			ch <- uint(i)
		}
	}()
	return ch
}

func bitCombinationOverSubsets(nums ...int) chan uint {
	ch := make(chan uint)
	s := uint(0)
	for _, v := range nums {
		s |= 1 << uint(v)
	}
	go func() {
		defer close(ch)
		for bit := s; ; bit = (bit - 1) & s {
			ch <- uint(bit)
			if bit == 0 {
				break
			}
		}
	}()
	return ch
}

func bitCombinationWithSize(num, size int) chan uint {
	ch := make(chan uint)
	bit := uint(1<<uint(size) - 1)
	go func() {
		defer close(ch)
		for ; bit < 1<<uint(num); bit = nextBitCombination(uint(bit)) {
			ch <- bit
		}
	}()
	return ch
}

func nextBitCombination(cur uint) uint {
	x := cur & -cur // rightest bit only         '10100' -> '00100'
	y := cur + x    // carry at rightest 1-block '10111' -> '11000'
	return (((cur & ^y) / x) >> 1) | y
}

func showIntUint(a int16) {
	fmt.Printf("a        : %d\n", a)
	fmt.Printf("uint16(a): %d\n", uint16(a))
	fmt.Printf("a           : b%016b\n", a)
	fmt.Printf("uint16(a)   : b%016b\n", uint16(a))
	fmt.Printf("a>>1        : b%016b\n", a>>1)
	fmt.Printf("uint16(a)   : b%016b\n", uint16(a))
	fmt.Printf("uint16(a>>1): b%016b\n", uint16(a>>1))
	fmt.Printf("uint16(a)>>1: b%016b\n", uint16(a)>>1)
	fmt.Printf("uint16(a)>>1: %d\n", int16(uint16(a)>>1))
}
