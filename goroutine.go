package leak

import (
	"runtime"
)

func GoRoutineLeaks(f func()) int {
	then := runtime.NumGoroutine()

	f()

	now := runtime.NumGoroutine()

	return now - then
}

type GoRoutineMark struct {
	then int
	now  int
}

func MarkGoRoutines() *GoRoutineMark {
	m := &GoRoutineMark{}

	m.then = runtime.NumGoroutine()

	return m
}

func (m *GoRoutineMark) Release() int {
	m.now = runtime.NumGoroutine()

	return m.now - m.then
}
