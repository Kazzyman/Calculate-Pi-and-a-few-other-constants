package main

import (
	"fmt"
	"math/big"
)

func main() {
	r := big.NewFloat(1)              // radius is a constant 1
	a := new(big.Float)              // height of bisected triangle
	s1 := big.NewFloat(1)            // starts out as 1
	s1_2 := new(big.Float)          // s1/2
	s2 := new(big.Float)            // new hypotenuse / new side
	b := new(big.Float)              // new short side
	numberOfSides := big.NewFloat(6) // the number of sides of the polygon
	p := new(big.Float)              // perimeter
	p_d := new(big.Float)            // pi

	// Set the precision to a higher value
	//precision := 500
	precision := 55000 // was 55000
	p_d.SetPrec(uint(precision))
	a.SetPrec(uint(precision))
	s1_2.SetPrec(uint(precision))
	s2.SetPrec(uint(precision))
	b.SetPrec(uint(precision))
	p.SetPrec(uint(precision))
	r.SetPrec(uint(precision))
	s1.SetPrec(uint(precision))
	numberOfSides.SetPrec(uint(precision))

	// Calculate initial value of p and set p_d to the same value
	numberOfSides.Mul(numberOfSides, big.NewFloat(2))
	s1_2.Quo(s1, big.NewFloat(2))
	a.Sqrt(new(big.Float).Sub(r, new(big.Float).Mul(s1_2, s1_2)))                     // height of bisected triangle
	b.Sub(r, a)                                                                         // new short side
	s2.Sqrt(new(big.Float).Add(new(big.Float).Mul(b, b), new(big.Float).Mul(s1_2, s1_2))) // new hypotenuse / NewSide
	s1.Set(s2)
	p.Mul(numberOfSides, s1)
	p_d.Set(p)

	var iterInt64 int64 = 0
	for iterInt64 < 5000 {
		iterInt64++
		numberOfSides.Mul(numberOfSides, big.NewFloat(2))
		s1_2.Quo(s1, big.NewFloat(2))
		a.Sqrt(new(big.Float).Sub(r, new(big.Float).Mul(s1_2, s1_2)))                     // height of bisected triangle
		b.Sub(r, a)                                                                         // new short side
		s2.Sqrt(new(big.Float).Add(new(big.Float).Mul(b, b), new(big.Float).Mul(s1_2, s1_2))) // new hypotenuse / NewSide
		s1.Set(s2)
		p.Mul(numberOfSides, s1)
		p_d.Set(p)
		p_d.Quo(p_d, big.NewFloat(2)) // Halve the value of p_d

		if iterInt64 == 24 { // last one
			fmt.Println("\n   #1 234567890123456# :: counting the first 16 actual digits of π")
			//fmt.Println("   ", p_d.String(), " is a string of what we have calculated per Archimedes' at 25 iters")
			//fmt.Println("   ", p_d, " is the naked big.Float of what we have calculated per Archimedes' at 25 iters")
			fmt.Printf("    %.20f is the big.Float of what we have calculated per Archimedes' at 24 iters, 20f\n", p_d)
			fmt.Println("    3.141592653589793238 is the value of π from the web")
			fmt.Println("   #1 234567890123456# :: counting the first 16 actual digits of π")

			fmt.Print(" the above was estimated from a ")
			fmt.Printf("%.0f sided polygon\n", numberOfSides)
			fmt.Printf("%.0f as parsed against ...\n", numberOfSides)
			fmt.Println("100000000 which is one-hundred-million, for comparison to the above line")
			fmt.Println("... Which is to say a 201,326,592 sided polygon, Mr. A. would have wept!\n")

		}

		if iterInt64 == 50 {
			fmt.Println("\n   #1 234567890123456789012345678901# :: counting the actual digits of π")
			//fmt.Println("   ", p_d, " is the naked big.Float of what we have calculated per Archimedes' at 50 iters")

			fmt.Printf("    %.33f is the big.Float of what we have calculated per Archimedes' at 50 iters, 33f\n", p_d)

			fmt.Println("    3.141592653589793238462643383279502 is the value of π from the web")
			fmt.Println("   #1 234567890123456789012345678901# :: counting the actual digits of π")


			fmt.Println(iterInt64, " iterations were completed yielding 31 correct digits of π\n")

			fmt.Printf("the above was estimated from a %.0f sided polygon (formatted as a .0f) \n", numberOfSides)

		}

		if iterInt64 == 150 {
fmt.Println(" #1 23456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123#")

			fmt.Printf("  %.120f per Archimedes'\n", p_d)

fmt.Println("  3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706798214808651328230664709 is from web")
fmt.Println(" #1 23456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123#")
fmt.Println("           10        20         30       40        50         60        70        80        90 ")


			fmt.Println(iterInt64, " iterations were completed yielding 92 correct digits of π\n")

			fmt.Printf("Calculated from a %.0f sided polygon\n", numberOfSides)

		}

		if iterInt64 == 200 {
fmt.Println(" #1 23456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890#")

			fmt.Printf("  %.120f per Archimedes'\n", p_d)

fmt.Println("  3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706798214808651328230664709 is from web")
fmt.Println(" #1 23456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890#")
fmt.Println("           10        20         30       40        50         60        70        80        90      100        110      120")


			fmt.Println(iterInt64, " iterations were completed yielding 121 correct digits of π\n")

			fmt.Printf("Calculated from a %.0f sided polygon\n\n\n", numberOfSides)

		}

		if iterInt64 == 5000 { // was 5500

			fmt.Printf("  %.2713f per Archimedes'\n", p_d)


			fmt.Println(iterInt64, " iterations were completed, \n ... which generated a ", numberOfSides, "sided polygon!!\n")

			//fmt.Printf("Calculated from a %.0f sided polygon\n", numberOfSides)
			fmt.Println(iterInt64, " iterations were completed yielding 2,712 correct digits of π!!!\n")


		}
	}
}
