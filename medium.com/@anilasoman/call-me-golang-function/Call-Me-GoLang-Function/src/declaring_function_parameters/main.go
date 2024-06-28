package main

import "fmt"

func main() {
  // Function call
  RepeatALine("hello", 3)
}

// Function definition
func RepeatALine(aLineOfText string, times int) {
    for count := 0 ; count < times ; count += 1 {
      fmt.Println(aLineOfText)
    }
  return
}
