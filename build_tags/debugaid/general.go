// +build !debug

package debugaid

const IsDebug = false

// Assert do nothing, because build mode isn't debug mode.
func Assert(expression bool) {
}
