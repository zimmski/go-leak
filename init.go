package leak

import (
	"sync"
)

var fixRan = false
var fixMutex sync.Mutex

func init() {
	// FIXME there is an allocation for the first started go routine, since I do not know where the allocation is coming from or how to fix that, I just want to ignore it by making an early goroutine. This does NOT show up during test execution. This is easily explained since AFAIK the tests are run in goroutines.

	fixMutex.Lock()
	defer fixMutex.Unlock()

	if fixRan {
		return
	}

	fixRan = true

	ch := func() chan bool {
		ch := make(chan bool)

		go func() {
			ch <- true

			close(ch)
		}()

		return ch
	}()

	<-ch
}
