package hanoi

import (
  "fmt"
  "testing"
)

func TestHanoi(t *testing.T) {
  for _, move := range Hanoi(Stack(15), Peg(1), Peg(3)) {
    fmt.Printf("Move disc %v from peg %v to peg %v.\n", move.disc, move.from, move.to)
  }
}
