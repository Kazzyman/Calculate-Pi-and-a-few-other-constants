package main

import (
    "fmt"       // Used for printing etc. 
    "time"      // Used for seeding random number generation; and calculating run times 
    "math/big"  // Used in ArchimedesBig(), and in bpp formula with mods by rick.woolley@gmail.com (pi calculator from https://github.com/karan/Projects)
)

        var colorReset = "\033[0m"          
        var colorRed = "\033[31m"
        var colorGreen = "\033[32m"
        var colorYellow = "\033[33m"
        var colorBlue = "\033[34m"
        var colorPurple = "\033[35m"
        var colorCyan = "\033[36m"
        var colorWhite = "\033[37m"

        var four float64  // is initialized to 4 where needed 

        var iterInt64 int64  // to be used primarily in selections which require modulus calculations 

func main(){  // case 5: 
    fmt.Println("    π = 3 + 4/(2*3*4) - 4/(4*5*6) + 4/(6*7*8) - 4/(8*9*10) + 4/(10*11*12) ...")
    fmt.Println("    One-Hundred-Million iterations will be executed ... working ...\n")
    start := time.Now() 
            one := big.NewFloat(1)
            two := big.NewFloat(2)
            three := big.NewFloat(3)
            four := big.NewFloat(4)

        //var LeftOfMantissa float64 
        //    LeftOfMantissa = 3
        LeftOfMantissa := big.NewFloat(3)
        //var digitone float64
        //    digitone = 2
        digitone := new(big.Float)
        digitone = two 
        //var digittwo float64
        //    digittwo = 3
        digittwo := new(big.Float)
        digittwo = three 
        //var digitthree float64
        //    digitthree = 4

        digitthree := new(big.Float)
        digitthree = four
        //four = 4
        //var firstsum float64
        //    firstsum = LeftOfMantissa + (four / (digitone*digittwo*digitthree)) 
        firstsum := new(big.Float) // just to create it 
        digitoneTdigitTwo := new(big.Float)
        //digitoneTdigitTwo = big.NewFloat(1)
        digitoneTdigitTwo.Mul(digitone, digittwo)
        firstsum.Add(LeftOfMantissa, (new(big.Float).Quo(four, (new(big.Float).Mul(digitoneTdigitTwo, digitthree)))))
        //var nextterm float64
        nextterm := new(big.Float)  // again, just to create it 

    precision := 52 // 55000
one.SetPrec(uint(precision))
two.SetPrec(uint(precision))
three.SetPrec(uint(precision))
four.SetPrec(uint(precision))
LeftOfMantissa.SetPrec(uint(precision))
digitthree.SetPrec(uint(precision))
digitone.SetPrec(uint(precision))
digittwo.SetPrec(uint(precision))
digitoneTdigitTwo.SetPrec(uint(precision))
firstsum.SetPrec(uint(precision))
nextterm.SetPrec(uint(precision))



        iterInt64 = 1
    for iterInt64 < 100000000 {
        iterInt64++
        //digitone = digitthree 
        digitone = digitthree
        //digittwo = digitthree + 1
        digittwo.Add(digitthree, one)
        //digitthree = digitthree + 2 
        digitthree.Add(digitthree, two)
        //nextterm = four/(digitone*digittwo*digitthree)
        digitoneTdigitTwo.Mul(digitone, digittwo)
        nextterm.Quo(four, new(big.Float).Mul(digitoneTdigitTwo, digitthree))
            if iterInt64 % 2 == 0 {  // % is modulus operator 
                //firstsum = firstsum - nextterm 
                firstsum.Sub(firstsum, nextterm) // ?????????????????????? probbably right ??
            } else {
                //firstsum = firstsum + nextterm
                firstsum.Add(firstsum, nextterm)
            }

            if iterInt64 == 1000000 {
                fmt.Println("   #1 2345678901234# ")
                fmt.Printf("    %0.22f is what we have calculated per Nilakantha Somayaji\n", firstsum) 
                fmt.Println("    3.1415926535897,93 is, again, the value of π from the web") 
                fmt.Println("   #1 2345678901234 56# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("1,000,000 iterations were completed in ", elapsed, " still yielding 14 digits of π\n") 
            }
            if iterInt64 == 100000000 {
                fmt.Println("   #1 2345678901234# ")
                fmt.Printf("    %0.22f is what we have calculated per Nilakantha Somayaji\n", firstsum) 
                fmt.Println("    3.1415926535897,93 is, again, the value of π from the web") 
                fmt.Println("   #1 2345678901234 56# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("100,000,000 iterations were completed in ", elapsed, " still yielding 14 digits of π\n") 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")

        //
                fmt.Println("  -- If we ran 50 billion more iterations we still would get only those 14 digits :(\n") // =========

            }  // end of last if in first for loop 
} // end of first for loop 
// optional code to show the futility of more iterations follows
        // 50 Billion more iterations option follows 
        fmt.Println(string(colorGreen), "Enter any positive digit to continue with the 50 billion iterations, 0 to exit", string(colorReset))
        option50 := 0
        fmt.Scan(&option50)

        if option50 > 0 {

            fmt.Println(string(colorCyan), "\n\nYou elected to continue the series by Indian astronomer Nilakantha Somayaji circa 1500 AD", string(colorReset))
            fmt.Println("    π = 3 + 4/(2*3*4) - 4/(4*5*6) + 4/(6*7*8) - 4/(8*9*10) + 4/(10*11*12) ...")
            fmt.Println("    50 billion iterations will be executed \n\n", string(colorCyan), "   ... working ...\n", string(colorReset))
/*
            start := time.Now() 
                var LeftOfMantissa float64 
                    LeftOfMantissa = 3
                var digitone float64
                    digitone = 2
                var digittwo float64
                    digittwo = 3
                var digitthree float64
                    digitthree = 4
                four = 4
                var firstsum float64
                    firstsum = LeftOfMantissa + (four / (digitone*digittwo*digitthree)) 
                var nextterm float64
                iterInt64 = 1
*/
// ========== not sure if or why we need to do the below ?????????????????????????????
//----------------------------------------------- but I had it above for some reason??
        LeftOfMantissa := big.NewFloat(3)
        //var digitone float64
        //    digitone = 2
        digitone := new(big.Float)
        digitone = two 
        //var digittwo float64
        //    digittwo = 3
        digittwo := new(big.Float)
        digittwo = three 
        //var digitthree float64
        //    digitthree = 4

        digitthree := new(big.Float)
        digitthree = four
        //four = 4
        //var firstsum float64
        //    firstsum = LeftOfMantissa + (four / (digitone*digittwo*digitthree)) 
        firstsum := new(big.Float) // just to create it 
        digitoneTdigitTwo := new(big.Float)
        //digitoneTdigitTwo = big.NewFloat(1)
        digitoneTdigitTwo.Mul(digitone, digittwo)
        firstsum.Add(LeftOfMantissa, (new(big.Float).Quo(four, (new(big.Float).Mul(digitoneTdigitTwo, digitthree)))))
        //var nextterm float64
        nextterm := new(big.Float)  // again, just to create it 
//===================================================================================


        iterInt64 = 1
                //    50,000,000,000
        for iterInt64 < 50000000000 {
            iterInt64++
/*
            digitone = digitthree 
            digittwo = digitthree + 1
            digitthree = digitthree + 2 
            nextterm = four/(digitone*digittwo*digitthree)
                if iterInt64 % 2 == 0 {  // % is modulus operator 
                    firstsum = firstsum - nextterm 
                } else {
                    firstsum = firstsum + nextterm
                }
*/
            digitone = digitthree
            //digittwo = digitthree + 1
            digittwo.Add(digitthree, one)
            //digitthree = digitthree + 2 
            digitthree.Add(digitthree, two)
            //nextterm = four/(digitone*digittwo*digitthree)

            digitoneTdigitTwo.Mul(digitone, digittwo)

            nextterm.Quo(four, new(big.Float).Mul(digitoneTdigitTwo, digitthree))
                if iterInt64 % 2 == 0 {  // % is modulus operator 
                    //firstsum = firstsum - nextterm 
                    firstsum.Sub(firstsum, nextterm)
                } else {
                    //firstsum = firstsum + nextterm
                    firstsum.Add(firstsum, nextterm)
                }
                        // 9,000,000,000
            if iterInt64 == 9000000000 {  
                fmt.Println("    1 2345678901234 ")
                fmt.Printf("    %0.22f is what we have calculated per Nilakantha Somayaji\n", firstsum) 
                fmt.Println("    3.1415926535897,93 is, again, the value of π from the web") 
                fmt.Println("    1 2345678901234 56 :: counting the first 16 actual digits of π")
                t := time.Now() ;  elapsed := t.Sub(start)
                fmt.Println(iterInt64, " 9 billion iterations were completed in ", elapsed, " yielding still just 14 digits of π\n") 
                fmt.Println(string(colorCyan), "  ... working ...\n", string(colorReset))
            }
            //             50,000,000,000
            if iterInt64 == 50000000000 {  // no additional digit are obtained even after 50 billion iterations
                fmt.Println("    1 2345678901234567 ")
                fmt.Printf("    %0.22f is what we have calculated per Nilakantha Somayaji\n", firstsum) 
                fmt.Println("    3.141592653589793 is, again, the value of π from the web") 
                fmt.Println("    1 2345678901234567 :: counting the first 16 actual digits of π")
                t := time.Now() ;  elapsed := t.Sub(start)
                fmt.Println(iterInt64, " 50 billion iterations were completed in ", elapsed, " yielding still just 14 digits of π\n") 
            }
        } // second for loop

    } // my options50 if
}