package main

import "fmt"
import "math"

func main() {
	// var is int by default as not pre-defined
	var investment_amount = 1000

	// float64 by default
	var expected_ror = 5.5

	// int by default
	var years = 10

	// investment_amount is type-casted to float as the calclulation was (int * float64)
	// no default power function in GO, math is a package to be imported to get power
	// Pow function takes parameters as (float64, float64), and return float64
	var future_value = float64(investment_amount) * math.Pow(1 + expected_ror / 100, float64(years))

	fmt.Println("Future value:", future_value)
	fmt.Println("Int of future value: ", int(future_value))
	fmt.Println("Int8 of future value: ", int8(future_value))
	fmt.Println("Int16 of future value: ", int16(future_value))
	fmt.Println("Int64 of future value: ", int64(future_value))
}