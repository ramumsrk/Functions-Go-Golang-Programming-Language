package main

import "fmt"

func main() {
  sample_input_number := 5
  fmt.Println("Before function call", sample_input_number)
  ModifySampleInputNumber(sample_input_number)
  fmt.Println("After function call", sample_input_number)
}

// Function definition
func ModifySampleInputNumber(sample_input_number int) {
  sample_input_number *= 2
}
