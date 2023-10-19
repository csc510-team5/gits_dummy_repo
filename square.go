package main

import "fmt"

// Function to square a number using ** operator
func square(number float64) float64 {
	return number**2
}

func main() {
    num := 5.0
    result := square(num) // Square the number using the ** operator
    fmt.Printf("The square of %.2f is %.2f\n", num, result)
}
