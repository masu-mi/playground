package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Calc struct {
	OpeStack   []string
	DigitQueue []int
	IsError    bool
}

func (c *Calc) PushOpe(ope string) {
	a := c.popDigit()
	b := c.popDigit()
	switch ope {
	case "+":
		fmt.Printf("%d+%d\n", b, a)
		c.pushDigit(b + a)
	case "-":
		fmt.Printf("%d-%d\n", b, a)
		c.pushDigit(b - a)
	case "*":
		fmt.Printf("%d*%d\n", b, a)
		c.pushDigit(b * a)
	case "/":
		fmt.Printf("%d/%d\n", b, a)
		c.pushDigit(b / a)
	case "**":
		fmt.Printf("%d**%d\n", b, a)
		c.pushDigit(int(math.Pow(float64(b), float64(a))))
	}
}

func (c *Calc) PushDigit(digit string) {
	n, _ := strconv.Atoi(digit)
	c.pushDigit(n)
}

func (c *Calc) pushDigit(n int) {
	c.DigitQueue = append(c.DigitQueue, n)
}

func (c *Calc) popDigit() int {
	var n int
	n, c.DigitQueue = c.DigitQueue[len(c.DigitQueue)-1], c.DigitQueue[:len(c.DigitQueue)-1]
	return n
}

func (c *Calc) Err(pos int, buffer string) {
	fmt.Println("")
	a := strings.Split(buffer[:pos], "\n")
	row := len(a) - 1
	column := len(a[row]) - 1

	lines := strings.Split(buffer, "\n")
	for i := row - 5; i <= row; i++ {
		if i < 0 {
			i = 0
		}

		fmt.Println(lines[i])
	}

	s := ""
	for i := 0; i <= column; i++ {
		s += " "
	}
	ln := len(strings.Trim(lines[row], " \r\n"))
	for i := column + 1; i < ln; i++ {
		s += "~"
	}
	fmt.Println(s)

	fmt.Println("error")
	c.IsError = true
}

func (c *Calc) Compute() {
	if !c.IsError {
		fmt.Printf("result: %d\n", c.popDigit())
	}
}

func main() {
	parser := &Parser{Buffer: "1+2*3-(4**5)--100"}
	parser.Init()
	err := parser.Parse()
	if err != nil {
		fmt.Println(err)
	} else {
		parser.Execute()
		parser.Compute()
	}
}
