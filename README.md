# Queues
Generic queues for Go.

![](https://github.com/twharmon/queues/workflows/Test/badge.svg) [![](https://goreportcard.com/badge/github.com/twharmon/queues)](https://goreportcard.com/report/github.com/twharmon/queues) [![codecov](https://codecov.io/gh/twharmon/queues/branch/main/graph/badge.svg?token=K0P59TPRAL)](https://codecov.io/gh/twharmon/queues)

## Documentation
For full documentation see [pkg.go.dev](https://pkg.go.dev/github.com/twharmon/queues).

## Install
```bash
go get github.com/twharmon/queues
````

## Usage
```go
package main

import (
	"fmt"

	"github.com/twharmon/queues"
)

func main() {
	// Create a new set of ints
	q := queues.New[int]()

	// Enqueue some values
	q.Enqueue(1, 2, 3)

	// Peek at the head
	q.Peek() // 1

	// Dequeue a value
	q.Dequeue() // 1

	// Get length of the queue
	q.Len() // 2

	// Get an ordered slice of the values in the queue
	q.Slice() // [2, 3]

	// Clear the queue
	q.Clear()
}
```

## Contribute
Make a pull request.

## Benchmarks

```
goos: darwin
goarch: arm64
pkg: github.com/twharmon/queues
BenchmarkEnqueue/enqueue-10         	   68956	     16422 ns/op	   25256 B/op	      13 allocs/op
BenchmarkEnqueue/enqueue_dequeue-10 	   39046	     30208 ns/op	   17072 B/op	      23 allocs/op
```
