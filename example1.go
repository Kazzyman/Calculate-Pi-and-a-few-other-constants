package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Initialize two big ints with the first two numbers in the sequence.
	//a := big.NewInt(2999999999999999999) // largest constants allowed here
	//b := big.NewInt(8888888888888888888)
	a := big.NewFloat(9989999999999999999999999999.054321) // larger constants allowed as Floats
	b := big.NewFloat(99999999999999999999999999999.123678)
	// Initialize limit as 10^99, the smallest integer with 100 digits.
	//var limit big.Int
	//limit.Exp(big.NewInt(10), big.NewInt(99), nil)

	// Loop while a is smaller than 1e100.
	//for a.Cmp(&limit) < 0 {
		// Compute the next Fibonacci number, storing it in a.
		//a.Add(a, b)
		a.Mul(a, b)
		// Swap a and b so that b is the next number in the sequence.
		//a, b = b, a
	//}
	//fmt.Println(a) // 100-digit Fibonacci number, or the sum of a and b
	fmt.Printf("this is formatted: %0.99f", a)

	// Test a for primality.
	// (ProbablyPrimes' argument sets the number of Miller-Rabin
	// rounds to be performed. 20 is a good value.)
	//fmt.Println(a.ProbablyPrime(20)) // prints false, or true 

}