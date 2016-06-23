// +build debug

package debugaid

const IsDebug = true

// Assert : if expression is false, goroutine panic!!
func Assert(expression bool) {
	if !expression {
		panic(expression)
	}
}
