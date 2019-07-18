package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	var n, l int
	fmt.Scan(&n, &l)
	ss := make([]string, n)
	for i := range ss {
		fmt.Scan(&ss[i])
	}
	sort.Sort(sort.StringSlice(ss))
	buf := bytes.NewBuffer([]byte{})
	for _, s := range ss {
		buf.WriteString(s)
	}
	fmt.Println(buf.String())
}
