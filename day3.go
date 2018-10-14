package main

import (
  "fmt"
  "math"
)

var input = 361527

func task1() int {
  level, db, max := 0, 0, 1
  for max < input {
    db += 8
    max += db
    level += 1
  }
  part := (db - (max - input)) / (db/4)
  position := db - (max - input) - part * db
  absPosition := math.Abs(float64(position - db/8))
  return level + int(absPosition)
}

func task2() int {
  level, db, max := 0, 0, 1
  for max < input {
    db += 8
    max += db
    level += 1
  }
  return 0
}

func main() {
  //fmt.Println(task1())
  fmt.Println(task2())
}
