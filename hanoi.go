package hanoi

// Note to self, here's the runtime:
//
//   Let n be the size of the original stack.
//   T(n) = 2*T(n-1) + 1 for n > 1, T(n) = 1 for n <= 1
//
// Ideas for generalization
//
//  - Use an arbitrary number of pegs
//  - Use an arbitrary collection of discs
//  - All arbitrary definitions for which discs can support which other discs
//  - Allow arbitrary restrictions on moves i.e. adjacency, etc...
//  - Terminal interface?
//
// Ideas for solutions
//
//  - Could, depending on which interfaces are implemented by a "game"
//    a solution be decided, in essence, by the compiler?

type Peg int

type Disc int

type Stack int

type Move struct {
  disc Disc
  from Peg
  to Peg
}

func Hanoi(stack Stack, from Peg, to Peg) []Move {
  output := make([]Move, 0)
  if stack.isEmpty() { return output }
  top := stack.topOfStack()
  bottom := stack.bottomDisc()
  other := otherPeg(from, to)
  output = append(output, Hanoi(top, from, other)...)
  output = append(output, Move{bottom, from, to})
  output = append(output, Hanoi(top, other, to)...)
  return output
}

func (stack Stack) isEmpty() bool {
  return stack == 0
}

func (stack Stack) topOfStack() Stack {
  return Stack(stack - 1)
}

func (stack Stack) bottomDisc() Disc {
  return Disc(stack)
}

func otherPeg(p1, p2 Peg) Peg {
  return Peg(6 - p2 - p1)
}
