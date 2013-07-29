package sudoku

import "fmt"


// A sudoku game
type Game struct {
  // board[i][j] = v means the i-th row, j-th column of the game is set to
  // value v
  board [9][9]int

  // fixed[i][j] = true means that the board was initialized with the i,j-th
  // coordinate fixed
  fixed [9][9]bool
}

// create a new game with some elements fixed
func NewGame(board [][]int) *Game {
  game := new(Game)
  for k:=0; k<9*9; k++ {
    i, j := index2coordinates(k)
    game.board[i][j] = board[i][j]
    if board[i][j] != 0 { game.fixed[i][j] = true }
  }
  return game
}

// Load a game from disk
func LoadGame(path string) Game {
  // TODO
  return Game{}
}

// Change one coordinate and return a copy of the game
func (g  Game) Set(i, j, v int) *Game {
  if !g.fixed[i][j] { g.board[i][j] = v }
  return &g
}

// Get a coordinate
func (g *Game) Get(i, j int) int {
  return g.board[i][j]
}

func (g *Game) IsFixed(i, j int) bool {
  return g.fixed[i][j]
}

// Is the name valid?
func (g *Game) IsComplete() bool {
  for i:=0; i < 9; i++ {
    if !(g.isRowComplete(i) &&
         g.isColumnComplete(i) &&
         g.isGridComplete(i)) {
      return false
    }
  }
  return true
}

// Check if a board is still valid
func (g *Game) IsValid() bool {
  for i:=0; i < 9; i++ {
    if !(g.isRowValid(i) &&
         g.isColumnValid(i) &&
         g.isGridValid(i)) {
      return false
    }
  }
  return true
}

// Pretty-print a board
func (g *Game) ToString() string {
  horizontalRow := "+-----+-----+-----+\n"
  s := ""
  for i:=0; i<9; i++ {
    if i % 3 == 0 { s += horizontalRow }
    for j:=0; j<9; j++ {
      if j % 3 == 0 { s += "|" } else { s+= " " }
      if g.Get(i,j) == 0 {
        s += "-"
      } else {
        s += fmt.Sprintf("%d", g.Get(i,j))
      }
    }
    s += "|\n"
  }
  s += horizontalRow
  return s
}

// Solve a game
func Solve(game *Game) *Game {
  return solve(game, 0)
}

/**************************** Private Functions *******************************/

func solve(game *Game, k int) *Game {
  //fmt.Println(game.ToString())
  // game done
  if game.IsComplete() {
    return game
  }

  // game invalid
  if !game.IsValid() {
    return nil
  }

  // find next index that needs to be changed.
  for {
    i, j := index2coordinates(k)
    if game.IsFixed(i, j) { k += 1 } else { break }
  }

  // try setting each value
  i, j := index2coordinates(k)
  for v:=1; v<10; v++ {
    solution := solve(game.Set(i, j, v), k+1)
    if solution == nil {
      // no solution down this path
    } else if solution.IsComplete() {
      // solution found
      return solution
    }
  }
  // no solution down original path
  return nil
}

func (g *Game) isRowComplete(i int) bool {
  return checkComplete(g.rowTaken(i))
}

func (g *Game) isColumnComplete(i int) bool {
  return checkComplete(g.columnTaken(i))
}

func (g *Game) isGridComplete(i int) bool {
  return checkComplete(g.gridTaken(i))
}

func (g *Game) isRowValid(i int) bool {
  return checkValid(g.rowTaken(i))
}

func (g *Game) isColumnValid(i int) bool {
  return checkValid(g.columnTaken(i))
}

func (g *Game) isGridValid(i int) bool {
  return checkValid(g.gridTaken(i))
}

// return number of times each value appears in row i
func (g *Game) rowTaken(i int) [10]int {
  var taken [10]int
  for j:=0; j<9; j++ {
    taken[g.Get(i,j)] += 1
  }
  return taken
}

// return number of times each value appears in column j
func (g *Game) columnTaken(j int) [10]int {
  var taken [10]int
  for i:=0; i<9; i++ {
    taken[g.Get(i,j)] += 1
  }
  return taken
}

func (g *Game) GridTaken(i int) [10]int {
  return g.gridTaken(i)
}

// return number of times each value appears in 3x3 grid i
func (g *Game) gridTaken(i int) [10]int {
  i0, j0 := 3 * (i / 3), 3 * (i % 3)

  var taken [10]int
  for di:=0; di<3; di++ {
    for dj:=0; dj<3; dj++ {
      taken[g.Get(i0+di, j0+dj)] += 1
    }
  }

  return taken
}

// check if all of taken[1]...taken[10] == 1
func checkComplete(taken [10]int) bool {
  for i:=1; i<10; i++ {
    if taken[i] != 1 {
      return false
    }
  }
  return true
}

// check if all of 0 <= taken[1]...taken[10] <= 1
func checkValid(taken [10]int) bool {
  for i:=1; i<10; i++ {
    if taken[i] > 1 {
      return false
    }
  }
  return true
}

// convert linear index to 2D index
func index2coordinates(k int) (int, int) {
  return k / 9, k % 9
}
