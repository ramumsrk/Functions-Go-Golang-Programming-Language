package main

import (
  "fmt"
  "errors"
)

func main() {
  quotient, error_occurred := division(10, 2)
  if error_occurred != nil {
    fmt.Println("Error: ", error_occurred)
  } else {
    fmt.Println("Quotient: ", quotient)
  }
  quotient, error_occurred = division(10, 0)
  if error_occurred != nil {
    fmt.Println("Error: ", error_occurred)
  } else {
    fmt.Println("Quotient: ", quotient)
  }
}

// Function definition
func division(dividend, divisor float64) (float64, error) {
    if divisor == 0 {
      return 0, errors.New("Division by zero")
    }
  return dividend / divisor , nil
}
