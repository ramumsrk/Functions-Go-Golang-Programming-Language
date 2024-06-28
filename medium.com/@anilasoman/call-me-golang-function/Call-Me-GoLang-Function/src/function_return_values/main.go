package main

import "fmt"

func main() {
  var dozen float64 = doubled(6.0)
  fmt.Println("One dozen: ", dozen)
  fmt.Println(doubled(7.0))
}

// Function definition
func doubled(quantity float64) float64 {
  return 2 * quantity
}
