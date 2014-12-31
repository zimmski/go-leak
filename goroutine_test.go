package leak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoRoutineLeaks(t *testing.T) {
	leaks := GoRoutineLeaks(func() {
		c := make(chan bool)

		leaks := GoRoutineLeaks(func() {
			go func() {
				<-c
			}()
		})

		assert.Equal(t, 1, leaks)

		c <- true
	})

	assert.Equal(t, 0, leaks)
}

func TestGoRoutineMark(t *testing.T) {
	c := make(chan bool)

	m := MarkGoRoutines()

	go func() {
		<-c
	}()

	assert.Equal(t, 1, m.Release())

	c <- true

	assert.Equal(t, 0, m.Release())
}
