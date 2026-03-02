package check

import (
	"fmt"
	"runtime/debug"
)

func Two[T any](t T, e error) T {
	One(e)
	return t
}

func One(e error) {
	if e != nil {
		print(fmt.Sprintf("Panic-ing at %s!", debug.Stack()))
		print(e)
		panic(e)
	}
}
