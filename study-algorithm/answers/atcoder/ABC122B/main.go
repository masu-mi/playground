package main

import (
	"fmt"
)

func main() {

	acceptable := newByteSet("ACGT")
	var str string
	fmt.Scanf("%s", &str)

	var max int

	for i := 0; i < len(str); {
		subL := 0
		for j := i; j < len(str); j++ {
			if !acceptable.doesContain(str[j]) {
				break
			}
			subL++
		}
		if max < subL {
			max = subL
		}
		// next of broken point
		i = i + subL + 1
	}
	fmt.Printf("%d\n", max)
}

func newByteSet(input string) set {
	s := newSet()
	for i := 0; i < len(input); i++ {
		s.add(input[i])
	}
	return s
}

type set map[byte]none

func newSet() set {
	return make(map[byte]none)
}

func (s set) add(item byte) {
	s[item] = mark
}

func (s set) doesContain(item byte) bool {
	_, ok := s[item]
	return ok
}

var mark none

type none struct{}
