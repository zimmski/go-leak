# go-leak [![GoDoc](https://godoc.org/github.com/zimmski/go-leak?status.png)](https://godoc.org/github.com/zimmski/go-leak)

The std packages of Go do currently not included detection for leaks. go-leak is a package which should help you find leaks in your code. If you have any ideas on how to improve this package or have problems of any kind with it, please [submit an issue](https://github.com/zimmski/go-leak/issues/new) through the [issue tracker](https://github.com/zimmski/go-leak/issues).

> **Note:** This package does not work anymore. It worked for some use-cases I had with an older Go version but with >= 1.4 it does not work at all. I have it on my TODO to redo the whole thing and add some more use-cases but there are some other open source projects I have to do first. So if you have some time on your hand and would like to contribute mail me or submit an issue in the project's tracker.

> **Note:** Since Go is scoped it is important to avoid new variables that are in the same scope as the mark-and-release calls.

## goroutines

If you want to know if a function is leaking goroutines:

```go
leaks := leak.GoRoutineLeaks(foo)

if leaks > 0 {
	panic("foo is leaking!")
}
```

If you want to know if a code is leaking goroutines:

```go
m := MarkGoRoutines()

// some code

leaks := m.Release()

if leaks > 0 {
	panic("some code is leaking!")
}
```


## memory

If you want to know if a function is leaking memory:

```go
leaks := leak.MemoryLeaks(foo)

if leaks > 0 {
	panic("foo is leaking!")
}
```

If you want to know if a code is leaking memory:

```go
m := MarkMemory()

// some code

leaks := m.Release()

if leaks > 0 {
	panic("some code is leaking!")
}
```
