// arch

package main

import (
    //"sort"      // was used in case 18: until it was not
    //"os"        // fetch the name of your system 
    //"io/ioutil" // file access 
    "fmt"       // Used for printing etc. 
    "math"      // Used for math.Pow(base, exponent)
    //"math/rand" // Used for random number generation in Monte Carlo method
    //"runtime"   // Used to get information on available CPUs
    "time"      // Used for seeding random number generation; and calculating run times 
    //"strconv"   // Used in Spigot 
    //"sync"      // Used in Bailey–Borwein–Plouffe formula [concurent]
    //"os/signal" // Used in *** nifty scoreboard *** concurrent computation of pi using Nilakantha's formula, by Diego Brener diegosilva13 on Github 
    "math/big"  // Used in bpp formula with mods by rick.woolley@gmail.com (pi calculator from https://github.com/karan/Projects)
)

func main() {       // top-level program logic flow -- explore SEVENTEEN ways to calculate Pi, plus THREE other constants
	Archimedes(4)
} 



func Archimedes(num int){ // case 4: 
    fmt.Println("\nYou selected #", num, "  --  An improved version of Archimedes' method")
    fmt.Println("  -- enter \"11' at the main menu for the derivation and proof of the Pythagorean -- ")
                            // the above escape does not seem to work as advertised ??
        start := time.Now()
        iterFloat64 = 0
        //var r float64 // radius is a constant 1 
        //var r *big.Float 
        //r := *big.NewFloat(float64(1))
        r := float64(1)
        //var a float64 // height of bisected triangle 
        //var a *big.Float

        //var s1 float64 // starts-out as 1
        //s1 := big.NewFloat(float64(1))
        s1 := float64(1)
        //var s1_2 float64 // s1/2
        //two := big.NewFloat(float64(2))
        two := float64(2)
        s1_2 := s1/two
        //var s2 float64 // new hypotenuse / new side
        var s2 *big.Float
        //var b float64 // new short side
        //var b *big.Float
        numberOfSides := float64(6) // number of sides 
        //numberOfSides := *big.NewFloat(float64(6))
        //var p float64 // perimeter
        //var p big.NewFloat 
        //var p_d float64 // pi
        var p *big.Float
        var p_d *big.Float

       

        //r = float64(1) // a constant, the radius 


        //numberOfSides = 6 // the number of sides of the polygon 
        //s1 = 1 // starts-out as 1

        //s1_2 = s1/2
        a := math.Sqrt(r-(math.Pow(s1_2,two))) // height of bisected triangle 
        b := 1-a // new short side
        s2 = big.NewFloat(math.Sqrt(math.Pow(b,two) + (math.Pow(s1_2,two)))) // new hypotenuse / NewSide
      //p = (numberOfSides) * (*big.NewFloat(s1)) // * not defined on (numberOfSides) (variable of type big.Float)
        p = big.NewFloat((numberOfSides) * (s1)) // * not defined on (numberOfSides) (variable of type big.Float)
        p_d = p/big.NewFloat(two) 

/* We iterate over the following psudo code:
let n = n * 2  // after having begun with 6 sides, then do 12 etc. 
let s1 = s2   // s1 begins as 1, thereafter it adopts s2 
 
s1_2 = s1/2 // set the variable s1_2 to one half of s1
a = math.Sqrt(r-(math.Pow(s1/2,2))) // length of side 'a' is height of bisected triangle 
b = 1-a                            // new short side 
s2 = math.Sqrt(math.Pow(b,2) + (math.Pow(s1/2,2))) // s2 is new hypotenuse : NewSide
p = n*s1    // p is the length of the perimeter of the constructed polygon
p_d = p/2  // the diameter of the polygon is always two, so p/2 = π 
*/
        iterInt64 = 0
    for iterInt64 < 301 { // was 27
        iterInt64++
        iterFloat64++
        numberOfSides = numberOfSides*2
        s1 = s2 
        s1_2 = s1/2 
        a = math.Sqrt(r-(math.Pow(s1_2,2))) 
        b = 1 - a 
        s2 = math.Sqrt(math.Pow(b,2) + (math.Pow(s1_2,2))) 
        p = numberOfSides * s1 
        p_d = p/2 

        if iterInt64 == 2 {
                fmt.Println("\n   #1 2# :: counting the first 2 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 2# :: counting the first 3 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 2 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", numberOfSides)
        }    
        if iterInt64 == 3 {
                fmt.Println("\n   #1 23# :: counting the first 3 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 23# :: counting the first 3 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 2 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", numberOfSides)
        }    
        if iterInt64 == 7 {
                fmt.Println("\n   #1 23456# :: counting the first 6 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 23456# :: counting the first 6 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 5 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", numberOfSides)
        }    
        if iterInt64 == 12 {
                fmt.Println("\n   #1 23456789# :: counting the first 9 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 23456789# :: counting the first 9 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 8 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", numberOfSides)
        }    
        if iterInt64 == 17 {
                fmt.Println("\n   #1 23456789012# :: counting the first 12 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 23456789012# :: counting the first 12 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 12 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", numberOfSides)
        }    
        if iterInt64 == 23 {
                fmt.Println("\n   #1 23456789012345# :: counting the first 15 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 23456789012345# :: counting the first 15 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 15 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", numberOfSides)
        }
        if iterInt64 == 25 {  // last one 
                fmt.Println("\n   #1 234567890123456# :: counting the first 16 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes' -- 16 -- digits") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 234567890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 16 digits of π\n") 

                //fmt.Printf(string(colorRed), "the above was estimated from a %.0f sided polygon", string(colorReset), "\n", n)
                fmt.Print(string(colorRed), " the above was estimated from a ")
                fmt.Printf("%.0f sided polygon\n", numberOfSides)
                fmt.Print(string(colorReset))
                fmt.Printf("%.0f as parsed against ...\n", numberOfSides)
                fmt.Println("100000000 which is one-hundred-million, for comparison to the above line")
                fmt.Println("... Which is to say a 201,326,592 sided polygon", string(colorRed), "Mr. A. would have wept!\n", string(colorReset))
                fmt.Println(" per option #", num, "  --  an improved version of Archimedes' method\n")

            arch := `Archimedes of Syracuse (287 – 212 BC) was a Greek mathematician, physicist, engineer, 
    astronomer, and inventor from the ancient city of Syracuse in Sicily. He estimated π 
    by drawing a larger regular hexagon outside a circle then a smaller regular hexagon inside
    the circle, and progressively doubling the number of sides of each regular polygon, 
    calculating the length of a side of each polygon at each step. As the number of sides 
    increases, it becomes a more accurate approximation of a circle. After four such steps, 
    when the polygons had 96 sides each, he was able to determine that the value of π lay between 
    3+1/7 (approx. 3.1429) and 3+10/71 (approx. 3.1408), consistent with its actual value of 
    approximately 3.1416 He also proved that the area of a circle was equal to π multiplied by the 
    square of the radius of the circle.`
            fmt.Println(string(colorCyan), arch, string(colorReset)) // prints the above rune

                LinesPerIter = 18  // an estimate
                fmt.Println("\nat aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64

                fmt.Printf("Aprox %.0f lines of code were executed per second \n\n", LinesPerSecond) 
                fmt.Println("The above calculation was per the following psudo code ...\n", string(colorCyan))
                fmt.Println("   let n = n * 2 // n is number of sides. After having begun with 6, then do 12, etc. ")
                fmt.Println("   let s1 = s2  // s1 begins as 1, thereafter it adopts s2, base of unbisected triangle")
                fmt.Println("   s1_2 = s1/2 // set the variable s1_2 to one half of s1")
                fmt.Println("   a = \u221A( 1-(s1/2)\u00b2) ) // length of side 'a' is height of bisected triangle ")
                fmt.Println("   b = 1-a                    // new short side ")
                fmt.Println("   s2 = \u221A( (b)\u00b2 + (s1/2)\u00b2 )  // s2 is new hypotenuse, and New s1 ")
                fmt.Println("   p = n * s1 // p is the length of the perimeter of the constructed polygon ")
                fmt.Println("   p_d = p/2 // the diameter of the polygon is always two, so p/2 = π \n", string(colorReset))
        }  

/*
        if iterInt64 == 300 {
                fmt.Println("\n   #1 234567890123456# :: counting the actual digits of π")
                fmt.Printf("    %.45f is what we have calculated per Archimedes' at the limit of the go language on this machine\n", p_d) 
                fmt.Println("    3.141592653589793238462643383279502884197169399 is the value of π from the web") 
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding still just 16 correct digits of π\n") 
                fmt.Printf("the above was estimated from a %d sided polygon (as a float64 formatted as an int)\n", numberOfSides)
                fmt.Printf(", or a %.0f sided polygon (formatted as a .0f) \n", numberOfSides)


                fmt.Println("\nat aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64

                fmt.Printf("Aprox %.0f lines of code were executed per second \n\n", LinesPerSecond) 

    // store reults in a log file which can be displayed from within the program by selecting option #12
                    fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                        check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                        defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                    Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  -- Archimedes of Syracuse -- selection #%d on %s \n", num, Hostname)
                        check(err0)
                    current_time := time.Now()
                    _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                    check(err6)
                    _ , err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second  \n", LinesPerSecond) 
                        check(err2)
                    _ , err4 := fmt.Fprintf(fileHandle, "%.02f was Iterations/Seconds  \n", iterFloat64/elapsed.Seconds())
                        check(err4)
                    _ , err5 := fmt.Fprintf(fileHandle, "%.0f was total Iterations  \n", iterFloat64)
                        check(err5)
                    TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                    _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
                        check(err7)

                fmt.Println("\n   -- After hitting Return for menu redisplay, enter '11' for the derivation and proof of the Pythagorean\n\n")
                fmt.Println("Select 12 at menu to display prior results")
        }
*/
    }
}

        var colorReset = "\033[0m"          
        var colorRed = "\033[31m"
        var colorGreen = "\033[32m"
        var colorYellow = "\033[33m"
        var colorBlue = "\033[34m"
        var colorPurple = "\033[35m"
        var colorCyan = "\033[36m"
        var colorWhite = "\033[37m"

        var TotalIterations int 
        var four float64  // is initialized to 4 where needed 
        var π float64  // a var can be any character, as in this Pi symbol/character 
        var LinesPerSecond float64  
        var LinesPerIter float64
        var iterInt64 int64  // to be used primarily in selections which require modulus calculations 
        var iterFloat64 float64  // to be used in selections which do not require modulus calculations 
        var Table_of_perfect_squares = []int{}
        var Table_of_perfect_Products = []int{}
        var num int 
        var diffOfLarger int // the following are used in the SquareRoot of x selection #1 case 1: 
        var diffOfSmaller int 
        var workPiece int
        var precisionOfSquare int 
        var t2 time.Time 
        var start2 time.Time
