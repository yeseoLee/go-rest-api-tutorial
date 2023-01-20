- singleton + concurrency

```go
package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var singleMazeEx Maze

func GetInstanceEx() Maze {
	once.Do(func() {
		initializeEx()
	})
	return singleMaze
}

func initializeEx() {
	fmt.Println("Initialize Maze")
	singleMazeEx = NewMaze()
}
```