package leak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MemoryLeaker struct {
	Foo int
}

func TestMemoryLeaks(t *testing.T) {
	leaks := MemoryLeaks(func() {
		var leaking *MemoryLeaker

		leaks := MemoryLeaks(func() {
			leaking = &MemoryLeaker{
				Foo: 123,
			}
		})

		assert.Equal(t, 1, leaks)

		leaking = nil
	})

	assert.Equal(t, 0, leaks)
}

func TestMemoryLeaksGoroutine(t *testing.T) {
	leaks := MemoryLeaks(func() {
		ch := func() chan bool {
			ch := make(chan bool)

			go func() {
				ch <- true

				close(ch)
			}()

			return ch
		}()

		<-ch
	})

	assert.Equal(t, 0, leaks)
}

var leaking *MemoryLeaker

func TestMemoryMark(t *testing.T) {
	m := MarkMemory()

	leaking = leakMemory()

	assert.Equal(t, 1, m.Release())

	leaking = nil

	assert.Equal(t, 0, m.Release())
}

func leakMemory() *MemoryLeaker {
	return &MemoryLeaker{
		Foo: 123,
	}
}
