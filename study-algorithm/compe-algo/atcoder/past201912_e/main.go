package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	rs := makeRelation(n)
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < q; i++ {
		sc.Scan()
		s := sc.Text()

		ss := strings.Split(s, " ")
		if ss[0] == "1" {
			rs.follow(toInt(ss[1]), toInt(ss[2]))
		} else if ss[0] == "2" {
			rs.followAllReturn(toInt(ss[1]))
		}
		if ss[0] == "3" {
			rs.followFollow(toInt(ss[1]))
		}
	}
	rs.WriteTo(os.Stdout)
}

func makeRelation(n int) *relations {
	rs := &relations{l: make([][]bool, n)}
	for i := 0; i < n; i++ {
		rs.l[i] = make([]bool, n)
	}
	return rs
}

type relations struct {
	l [][]bool
}

func (rs *relations) WriteTo(w io.Writer) (l int64, e error) {
	for _, r := range rs.l {
		for _, v := range r {
			if v {
				fmt.Fprintf(w, "Y")
			} else {
				fmt.Fprintf(w, "N")
			}
		}
		fmt.Fprintln(w)
		l += int64(len(r) + 1)
	}
	return l, nil
}

func (rs *relations) follow(a, b int) {
	if a == b {
		return
	}
	rs.l[a-1][b-1] = true
}

func (rs *relations) followAllReturn(a int) {
	for _, v := range rs.followerList(a) {
		rs.follow(a, v)
	}
}

func (rs *relations) followFollow(a int) {
	for _, x := range rs.followList(a) {
		for _, y := range rs.followList(x) {
			rs.follow(a, y)
		}
	}
}

// xがfollowしている人一覧
func (rs *relations) followList(x int) []int {
	var list []int
	for i, v := range rs.l[x-1] {
		if v {
			list = append(list, i+1)
		}
	}
	return list
}

// xをfollowしている人一覧
func (rs *relations) followerList(x int) []int {
	var list []int
	for i, v := range rs.l {
		if v[x-1] {
			list = append(list, i+1)
		}
	}
	return list
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
