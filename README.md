go-sudoku
=========

A brute force Sudoku solver written in Go.

Example
=======

```go
package main

import (
  "fmt"
  "github.com/duckworthd/sudoku"
)

func main() {
  // define a board as a 2D array. 0 values indicate "to be filled in" parts of
  // the board.
  board := [][]int{
      []int{5, 9, 8, 3, 0, 7, 1, 0, 4},
      []int{0, 4, 0, 0, 9, 0, 3, 0, 5},
      []int{0, 0, 0, 0, 0, 0, 6, 9, 0},
      []int{0, 0, 2, 4, 0, 3, 0, 1, 0},
      []int{0, 0, 0, 0, 2, 0, 0, 0, 0},
      []int{0, 3, 0, 9, 0, 6, 4, 0, 0},
      []int{0, 1, 7, 0, 0, 0, 0, 0, 0},
      []int{8, 0, 9, 0, 4, 0, 0, 6, 0},
      []int{3, 0, 6, 2, 0, 9, 8, 4, 1},
    }

  // construct a game object, then solve it
  game := sudoku.NewGame(board)
  solution := sudoku.Solve(game)

  // check solution and print it
  fmt.Println("Is solution valid?", solution.IsValid())
  fmt.Println("Is solution complete?", solution.IsComplete())
  fmt.Println(solution.ToString())
}
```
