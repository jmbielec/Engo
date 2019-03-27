package main

import (
  "fmt"
  
  "github.com/jmbielec/Engo/board"
)

func main() {
  fmt.Println("hello world")
  board := board.GenerateStartingBoard()
  fmt.Println(board)
}