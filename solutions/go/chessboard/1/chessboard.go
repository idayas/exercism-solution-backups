package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
  total := 0
  for key, f := range cb {
    if key != file {
      continue
    }

    for _, v := range f {
      if v {
        total++
      }
    }
  }

  return total
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
  total := 0

  for _, f := range cb {
    for i, v := range f {
      i++
      if i != rank {
        continue
      }

      if v {
        total++
      }
    }
  }
  return total
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
  squares := 0
  for _, f := range cb {
    squares += len(f)
  }
  return squares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
  squares := 0
  for _, f := range cb {
    for _, v := range f {
      if v {
        squares++
      }
    }
  }
  return squares
}
