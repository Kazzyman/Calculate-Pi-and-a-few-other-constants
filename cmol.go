// compile with: "go build -o constants constantsMacOrLinux.go"
// ... thereafter you can run it with "/fullpathname/constants"
// ... or, having first obtained the Go compiler, ... just run the source with: "go run constants.go"

// One can obtain the Go Compiler from : https://go.dev/dl/

package main
import ("fmt"; "time"; "math"; "os"; "io/ioutil") // include additional packages 

                var LinesPerSecond float64  // global variables defined outside any function
                var LinesPerIter float64
                var iterInt64 int64  // to be used primarily in selections which require modulus calculations 
                var iterFloat64 float64  // to be used in selections which do not require modulus calculations 

                // a Println("\\caf\u00e9 , 5*5 is 5\u00b2 , squareroot symbol is \u221A") // Prints \café
                // a Println("\u00b0 is degree symbol") // would demonstrate the degree symbol 

func check(e error) {   // create a func named check which takes e of type error 
    if e != nil {
        panic(e)        // use panic to display error code 
    }
}

func main() {        // top-level program logic flow 
    for 1 == 1 {    // loop endlessly, or Ctrl-C to Exit
        RicksDisplayMenu()   // displays the menu 
        RicksSwitch()       // contains a Scanf to input a menu selection 
/*
    // additional dummy pause -- which is not needed on a Linux System 
            var Pnum int // intended to be used only as a dummy 
            fmt.Println("Hit Enter/Return to redisplay the menu") 
            fmt.Scanf("%d", &Pnum) // often flys right-past this Scanf 
*/
    // another dummy pause (the above block) is "needed" in some environments because they sometimes ... 
    // ... fly rite-past the first and then redisplay the menu (Windows 11 does this)
            var Mnum int
            fmt.Println("Hit Enter/Return again to redisplay the menu") 
            fmt.Scanf("%d", &Mnum) 

    } // end of for 
} // end of main, end of top-level program logic flow 


func RicksDisplayMenu() {
p := fmt.Println  // define p as synonomous with the format Print line function
p("")
p("Enter 1 - 6 to execute one of the following methods for calculating π (pi) ")
p(" ")
// Veritassium https://youtu.be/gMlf1ELvRzc?list=LL  was my initial inspiration for all of this ...
// ... option 1 below having been discussed on his channel 
p("1:   Contains a square root, and only yeilds 4 digits of π")
p("     π = 12 * ( 1/2 - (1/2  * 1/3 * (1/2)exp3) - ...")
p("                                    (1/8   * 1/5 * (1/2)exp5) - ...")
p("                                    (1/16  * 1/7 * (1/2)exp7) - ...")
p("                                    (1/128 * 1/9 * (1/2)exp9) - ... - (\u221A3)/8) )")
p("       4 digits of π -- 4-MAXIMUM -- but instantly")
p(" ")
p("2:   An infinite series by Indian astronomer Nilakantha Somayaji circa 1500 ")
p("     ... this equation can be found in an ancient Sanskrit verse")
p("     π = 3 + 4/(2*3*4) - 4/(4*5*6) + 4/(6*7*8) - 4/(8*9*10) + 4/(10*11*12) ...")
p("       One-Hundred-Million iterations will be executed -- in less than a second")
p("       14 digits of π -- VERY FAST -- gives 7 digits in just 100 iterations !!!")
p(" ")
p("3:   The Gregory-Leibniz series circa 1676")
p("     π = (4/1) - (4/3) + (4/5) - (4/7) + (4/9) - (4/11) + (4/13) - (4/15) ...")
p("       Three-Hundred-Million iterations will be executed")
p("       9 digits of π -- in seconds")
p(" ")
p("4:   Gottfried Wilhelm Leibniz, late 17th century ( and by Isaac Newton )")
p("     π = 4 * ( 1 - 1/3 + 1/5 - 1/7 + 1/9 ...")
p("       4 Billion iterations will be executed")
p("       10 digits of π -- 8 digits in 100M iterations in under a minute ")
p(" ")
p("5:   Archimedes' method (improved) of bisecting triangles, circa 200 BC")
p("     π = begining with a hexagon, iteratively double its number of sides") 
p("       26 iterations is all it will take to get the ... ")
p("       16 digits of π -- v.fast     (enter '6' for Pythagorean theorem)")
p("")
p("7:   An infinite series by John Wallis circa 1655")
p("     π = 2 * ((2/1)*(2/3)) * ((4/3)*(4/5)) * ((6/5)*(6/7)) ... ")
p("       One-Billion iterations will be executed; option for 40 billion iterations")
p("       9 digits of π -- a billion loops, in seconds -- option for 10 digits\n")
p("8:   Euler's Number: \u2107 or \u2147 the natural logarithmic base")
p("       Explore the limit of interest")
p(" ")
p("100: Enter '100' to display prior results\n")

//fmt.Print("Enter your selection from above, 1 - 8 \u00a9 2022, by Richard Hart Woolley\n")
// the above commented line is definately not needed in a Linux environment
fmt.Println("Ctrl-C to End/Exit  SLOC including comments and white space = 1292 \n")
}

func RicksSwitch() {    // the main if-then-else routine to select an option from the menu

        iterInt64 = 0     // this will be used to count iterations and for modulus calculations 
        iterFloat64 = 0  // this is needed to calculate and display code lines per second 

    var secondDenom float64  
    secondDenom = 3
    var four float64
    four = 4
    var π float64
    var num int 
    p := fmt.Println

//fmt.Print("Enter your selection from above, 1 - 8 \u00a9 2022, by Richard Hart Woolley [it's an echo]\n")
// the above may be needed in some environments (windows) along with the duplicate in RicksDisplayMenu()
fmt.Print("Enter your selection from above, 1 - 8 \u00a9 2022, by Richard Hart Woolley \n")
fmt.Scanf("%d", &num)  // pause and accept input from the user

if num > 11 && num < 100 { num = 12 }  // to display a funny out-of-range message as case 12
switch num {  // will do only the case selected by inputed variable "num"
case 100: // read and display contents of prior results file ... 
    content, err := ioutil.ReadFile("dataFromConstants.txt")
        if err != nil {   // if the file does not exist ... 
            fmt.Println("no prior results -- no log file dataFromConstants.txt exists")
        } else {
            fmt.Println(string(content))
        }


case 1: // --- if inputed varible num was "1" then only do case 1 ------------------------------------------
p(" ")
p("You selected #", num)
p(" ")
p("     π = 12 * ( 1/2 - (1/2  * 1/3 * (1/2)exp3) - ")
p("                                    (1/8   * 1/5 * (1/2)exp5) - ")
p("                                    (1/16  * 1/7 * (1/2)exp7) - ")
p("                                    (1/128 * 1/9 * (1/2)exp9) - ... - Sqrt(3)/8) )")
p(" ")

var twelve float64
twelve = 12
var πViaNewton float64 
var SumOfTerms float64 

var initialDenomOfFirst float64
initialDenomOfFirst = 2
var nextDenomOfFirst float64 
nextDenomOfFirst = 8

var savedFirstDenom float64
savedFirstDenom = 0.0

var exponent float64 
exponent = 3

SumOfTerms = - ( (1/initialDenomOfFirst) * (1/secondDenom) * math.Pow(0.50,exponent) )  // 2 and 3 and 3

    for iterInt64 < 6 {
        iterInt64 = iterInt64 + 1 
            exponent = exponent + 2.0
            secondDenom = secondDenom + 2.0
            SumOfTerms = (SumOfTerms - ( (1/nextDenomOfFirst) * (1/secondDenom) * math.Pow(0.50,exponent) ) )
            p(nextDenomOfFirst, secondDenom, exponent)
            savedFirstDenom = nextDenomOfFirst 
            nextDenomOfFirst = savedFirstDenom * initialDenomOfFirst 
            initialDenomOfFirst = savedFirstDenom 
        if iterInt64 == 2 {
            πViaNewton = ( twelve * ( (1/2.0) + SumOfTerms - (math.Sqrt(3)/8.0) )  ) 
            p("After", iterInt64, "iterations with exponent:", exponent)
            p("  #1 234#")
            p("  ", πViaNewton, "The same equation on a pocket calculator gives 3.14171 with only 3 terms e7")
            p("   3.141592653589793 is the value of π from the web") 
            p("  #1 234567890123456# :: counting the first 16 actual digits of π")
            p(" ")
        }
        if iterInt64 == 3 {
            πViaNewton = ( twelve * ( (1/2.0) + SumOfTerms - (math.Sqrt(3)/8.0) )  ) 
            p("After", iterInt64, "iterations with exponent:", exponent)
            p("  #1 234#")
            p("  ", πViaNewton, "The same equation on a pocket calculator gives 3.14171 with only 3 terms e7")
            p("   3.141592653589793 is the value of π from the web") 
            p("  #1 234567890123456# :: counting the first 16 actual digits of π")
            p(" ")
        }
        if iterInt64 == 4 {
            πViaNewton = ( twelve * ( (1/2.0) + SumOfTerms - (math.Sqrt(3)/8.0) )  ) 
            p("After", iterInt64, "iterations with exponent:", exponent)
            p("  #1 234#")
            p("  ", πViaNewton, "The same equation on a pocket calculator gives 3.14171 with only 3 terms e7")
            p("   3.141592653589793 is the value of π from the web") 
            p("  #1 234567890123456# :: counting the first 16 actual digits of π")
            p(" ")
        }
        if iterInt64 == 5 {
            πViaNewton = ( twelve * ( (1/2.0) + SumOfTerms - (math.Sqrt(3)/8.0) )  ) 
            p("After", iterInt64, "iterations with exponent:", exponent)
            p("  #1 234#")
            p("  ", πViaNewton, "The same equation on a pocket calculator gives 3.14171 with only 3 terms e7")
            p("   3.141592653589793 is the value of π from the web") 
            p("  #1 234567890123456# :: counting the first 16 actual digits of π")
            p(" ")
        }
        if iterInt64 == 6 {
            πViaNewton = ( twelve * ( (1/2.0) + SumOfTerms - (math.Sqrt(3)/8.0) )  ) 
            p("After", iterInt64, "iterations with exponent:", exponent)
            p("  #1 234#")
            p("  ", πViaNewton, "The same equation on a pocket calculator gives 3.14171 with only 3 terms e7")
            p("   3.141592653589793 is the value of π from the web") 
            p("  #1 234567890123456# :: counting the first 16 actual digits of π")
        }
    }

p(" ")
p("Newton also found π to 16 digits by the first 22 terms of :")
p("3*Sqrt(3)/4 + 24*( 2/3*2^3 - 1/5*2^5 - 1/28*2^7 - 1/72*2^9 - 5/704*2^11 - 7/1664*2^13 )")  // this psudocode is emplemented below

π = ( (3*math.Sqrt(3))/4) + ( 24* ( 2/(3*math.Pow(2,3)) - 1/(5*math.Pow(2,5)) - 1/(28*math.Pow(2,7)) - 1/(72*math.Pow(2,9)) - 5/(704*math.Pow(2,11)) - 7/(1664*math.Pow(2,13)) )  )
p(π,"per the above formula (only six terms, of the supposed 22 terms he used)")
p("3.141592653589793 is the value of π from the web\n") 

//p(12.0*((1/2.0) - (1/2.0)*(1/3.0)*math.Pow(0.50,3) - (1/8.0)*(1/5.0)*math.Pow(0.50,5) - (1/16.0)*(1/7.0)*math.Pow(0.50,7) - math.Sqrt(3)/8.0), "straight (3-terms) calc via go" )
//p(12.0*((1/2.0) - (1/2.0)*(1/3.0)*math.Pow(0.50,3) - (1/8.0)*(1/5.0)*math.Pow(0.50,5) - (1/16.0)*(1/7.0)*math.Pow(0.50,7) - (1/128.0)*(1/9.0)*math.Pow(0.5,9) - math.Sqrt(3)/8.0), "straight (4-terms) calc via go" )


case 2: // ------------------------------------------------------------------------------------------------

p(" ")
p(" ")
p("You selected #", num, " a series by Indian astronomer Nilakantha Somayaji circa 1500 AD")
p("    π = 3 + 4/(2*3*4) - 4/(4*5*6) + 4/(6*7*8) - 4/(8*9*10) + 4/(10*11*12) ...")
p("    One-Hundred-Million iterations will be executed ... working ...")
p(" ")
    start2 := time.Now()

    var LeftOfMantissa float64 
    LeftOfMantissa = 3
    var digitone float64
    digitone = 2
    var digittwo float64
    digittwo = 3
    var digitthree float64
    digitthree = 4
    var firstsum float64
    firstsum = LeftOfMantissa + (four / (digitone*digittwo*digitthree)) 
    var nextterm float64

    for iterInt64 < 100000000 {
        iterFloat64 = iterFloat64 + 1
        iterInt64 = iterInt64 + 1
        digitone = digitthree 
        digittwo = digitthree + 1
        digitthree = digitthree + 2 
        nextterm = four/(digitone*digittwo*digitthree)
            if iterInt64 % 2 == 0 {
                firstsum = firstsum - nextterm 
            } else {
                firstsum = firstsum + nextterm
            }

            if iterInt64 == 100 {
                p("   #1 234567# ")
                p("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                p("    3.141592,653589793 is, again, the value of π from the web") 
                p("   #1 234567 890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start2)
                p (iterInt64, " iterations were completed in ", elapsed, " yielding 7 digits of π\n") 
            }

            if iterInt64 == 500 {
                p("   #1 23456789# ")
                p("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                p("    3.14159265,3589793 is, again, the value of π from the web") 
                p("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start2)
                p (iterInt64, " iterations were completed in ", elapsed, " yielding 9 digits of π\n") 
            }
// case 2
            if iterInt64 == 10000 {
                p("   #1 234567890123# ")
                p("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                p("    3.141592653589,793 is, again, the value of π from the web") 
                p("   #1 234567890123 456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start2)
                p("10,000 iterations were completed in ", elapsed, " yielding 13 digits of π\n") 
            }
            //
            if iterInt64 == 50000 {
                p("   #1 2345678901234# ")
                p("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                p("    3.1415926535897,93 is, again, the value of π from the web") 
                p("   #1 2345678901234 56# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start2)
                p("50,000,000 iterations were completed in ", elapsed, " yielding 14 digits of π\n") 
            }
            if iterInt64 == 1000000 {
                p("   #1 2345678901234# ")
                p("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                p("    3.1415926535897,93 is, again, the value of π from the web") 
                p("   #1 2345678901234 56# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start2)
                p("1,000,000 iterations were completed in ", elapsed, " still yielding 14 digits of π\n") 
            }
// case 2
            if iterInt64 == 100000000 {
                p("   #1 2345678901234# ")
                p("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                p("    3.1415926535897,93 is, again, the value of π from the web") 
                p("   #1 2345678901234 56# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start2)
                p("100,000,000 iterations were completed in ", elapsed, " still yielding 14 digits of π\n") 
// case 2
                LinesPerIter = 15
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")

                fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                    check(err1)                                // ... gets a file handle to dataFromConstants.txt
                    defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
// case 2
                Hostname, _ := os.Hostname()
                _ , err0 := fmt.Fprintf(fileHandle, "\n  -- Nilakantha Somayaji -- case #%d on %s \n", num, Hostname)
                    check(err0)
                _ , err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second from case %d \n", LinesPerSecond, num) 
                    check(err2)
                _ , err3 := fmt.Fprintf(fileHandle, "%.04f was total Elapsed Seconds from case %d \n", elapsed.Seconds(), num) 
                    check(err3)
                _ , err4 := fmt.Fprintf(fileHandle, "%e was Iterations/Seconds from case %d \n", iterFloat64/elapsed.Seconds(), num)
                    check(err4)
                _ , err5 := fmt.Fprintf(fileHandle, "%e was total Iterations from case %d \n", iterFloat64, num)
                    check(err5)
//
    p("  -- If we ran 50 billion more iterations we still would get only those 14 digits :(\n") // =========
//
                p(" per option #", num, "  --  the Nilakantha Somayaji series, circa 1500 AD\n")
                fmt.Println("Select 100 at menu to display prior results")
// case 2
            }
            //
            /*
            if iterInt64 == 1000000000 {
                p("    1 2345678901234 ")
                p("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                p("    3.1415926535897,93 is, again, the value of π from the web") 
                p("    1 2345678901234 56 :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start2)
                p (iterInt64, " iterations were completed in ", elapsed, " yielding still just 14 digits of π\n") 
                p(" per option #3\n")

            }
            if iterInt64 == 5000000000 {
                p("    1 2345678901234 ")
                p("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                p("    3.1415926535897,93 is, again, the value of π from the web") 
                p("    1 2345678901234 56 :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start2)
                p (iterInt64, " 5 billion iterations were completed in ", elapsed, " yielding still just 14 digits of π\n") 
                p(" per option #3\n")
// case 2 
            }            
            if iterInt64 == 9000000000 {
                p("    1 2345678901234 ")
                p("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                p("    3.1415926535897,93 is, again, the value of π from the web") 
                p("    1 2345678901234 56 :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start2)
                p (iterInt64, " 9 billion iterations were completed in ", elapsed, " yielding still just 14 digits of π\n") 
                p(" per option #3\n")

            }
            if iterInt64 == 50000000000 {  // no additional digit are obtained even after 500 billion iterations
                p("    1 2345678901234567 ")
                p("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                p("    3.141592653589793 is, again, the value of π from the web") 
                p("    1 2345678901234567 :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start2)
                p (iterInt64, " 50 billion iterations were completed in ", elapsed, " yielding still just 14 digits of π\n") 
                p(" per option #", num, "\n")
            }
            */
    }

case 3: // ----------------------------------------------------------------------------------------------

p(" ")
p(" ")
p("You selected #", num, " the Gregory-Leibniz series ...")
p("π = (4/1) - (4/3) + (4/5) - (4/7) + (4/9) - (4/11) + (4/13) - (4/15) ...")
p("Three hundred million iterations are being executed ... working ...")
p(" ")
    start3 := time.Now()

    var nextOdd float64
    nextOdd = 1
    var tally float64    // used in selection 3 only 
    tally = (four/nextOdd)

    for iterInt64 < 300000000 {
        iterInt64 = iterInt64 + 1
        iterFloat64 = iterFloat64 + 1
        nextOdd = nextOdd + 2
        tally = tally-(tally/nextOdd)
        tally = tally+(tally/nextOdd) // pi is set equl to the sum of a subtraction and an addition, alternatively 

                if iterInt64 == 10000000 {p("... 10,000,000 of three hundred million completed. still working, but ...")
                p(" ")
                p("   #1 234567#")
                p("   ",tally, "was calculated by the Gregory-Leibniz series") 
                p("    3.141592,653589793 is from the web") 
                p("   #1 234567 890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start3)
                fmt.Print("  10,000,000 iterations in ", elapsed, " yields 7 digits of π\n")
                p(" ")
                }
// 7 digits per above
// 8, next two ifs give eight digits
                if iterInt64 == 50000000 {p("... 50,000,000 of three hundred million completed. still working, but ...")
                p(" ")
                p("   #1 2345678#")
                p("   ",tally, "was calculated by the Gregory-Leibniz series") 
                p("    3.1415926,53589793 is from the web") 
                p("   #1 2345678 90123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start3)
                fmt.Print("  50,000,000 iterations in ", elapsed, " yields 8 digits of π\n")
                p(" ")
                }
                if iterInt64 == 100000000 {p("... 100,000,000 of three hundred million completed. still working, but ...")
                p(" ")
                p("   #1 2345678#")
                p("   ",tally, "was calculated by the Gregory-Leibniz series") 
                p("    3.1415926,53589793 is from the web") 
                p("   #1 2345678 90123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start3)
                fmt.Print("  100,000,000 iterations in ", elapsed, " yields 8 digits of π\n")
                p(" ")
                }
// case 3
// 9
            if iterInt64 == 200000000 {p("... 200,000,000 of three hundred million completed. still working, but ...")
                p(" ")
                p("   #1 23456789#")
                p("   ",tally, "was calculated by the Gregory-Leibniz series") 
                p("    3.14159265,3589793 is from the web") 
                p("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start3)
                fmt.Print("  200,000,000 iterations in ", elapsed, " yields 9 digits of π\n")
                p(" ")
                }
            if iterInt64 == 300000000 {
                fmt.Print ("    ", tally, " was calculated by the Gregory-Leibniz series")
                p("\n    3.141592653589793 is from the web") 
                t := time.Now()
                elapsed := t.Sub(start3)
                fmt.Print ("  300 million iterations still yields 9 digits, ")  
                fmt.Print ("in ", elapsed, "\n\n") 
                p(" per option #", num, "  --  the Gregory-Leibniz series, circa 1676\n")

                LinesPerIter = 11
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("       %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")
// 9
                fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                    check(err1)                                // ... gets a file handle to dataFromConstants.txt
                    defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
// case 3
                Hostname, _ := os.Hostname()
                _ , err0 := fmt.Fprintf(fileHandle, "\n  -- Gregory-Leibniz -- case #%d on %s \n", num, Hostname)
                    check(err0)
                _ , err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second from case %d \n", LinesPerSecond, num) 
                    check(err2)
                _ , err3 := fmt.Fprintf(fileHandle, "%.04f was total Elapsed Seconds from case %d \n", elapsed.Seconds(), num) 
                    check(err3)
                _ , err4 := fmt.Fprintf(fileHandle, "%e was Iterations/Seconds from case %d \n", iterFloat64/elapsed.Seconds(), num)
                    check(err4)
                _ , err5 := fmt.Fprintf(fileHandle, "%e was total Iterations from case %d \n", iterFloat64, num)
                    check(err5)
                fmt.Println("Select 100 at menu to display prior results")

            }
    }

case 4: // --------------------------------------------------------------------------------------

p(" ")
p(" ")
p("You selected #", num, " Gottfried Wilhelm Leibniz formula  :  π = 4 * ( 1 - 1/3 + 1/5 - 1/7 + 1/9 ... ")
p("   Infinitesimal calculus was developed independently in the late 17th century by Isaac Newton")
p("    ... and Gottfried Wilhelm Leibniz")
p("   4 Billion iterations will be executed ... ")
p("")
p(" ... working ...")
p(" ")
    start4 := time.Now()

    var denom float64
    denom = 3
    var sum float64
    sum = 1-(1/denom)

    for iterInt64 < 4000000000 {
        iterFloat64 = iterFloat64 + 1
        iterInt64 = iterInt64 + 1 
        denom = denom + 2
            if iterInt64 % 2 == 0 {
                sum = sum + 1/denom 
            } else { 
                sum = sum - 1/denom 
            }
            π = 4 * sum 
                if iterInt64 == 100000000 {p("... 100,000,000 completed iterations ...")
                p(" ")
                p("   #1 2345678#")
                p("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                p("    3.1415926,53589793 is from the web") 
                p("   #1 2345678 90123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start4)
                fmt.Print("  100,000,000 iterations in ", elapsed, " yields 8 digits of π\n")
                p(" ")
                }
//
// case 4
                if iterInt64 == 200000000 {p("... 200,000,000 gets another digit ...")
                p(" ")
                p("   #1 23456789#")
                p("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                p("    3.14159265,3589793 is from the web") 
                p("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start4)
                fmt.Print("  200,000,000 iterations in ", elapsed, " yields 9 digits of π\n")
                p(" ")
                }
//
                if iterInt64 == 400000000 {p("... 400,000,000 iterations completed, still at nine ...")
                p(" ")
                p("   #1 23456789#")
                p("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                p("    3.14159265,3589793 is from the web") 
                p("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start4)
                fmt.Print("  400,000,000 iterations in ", elapsed, " yields 9 digits of π\n")
                p(" ")
                }
//
                if iterInt64 == 600000000 {p("... 600,000,000 iterations, still at nine ...")
                p(" ")
                p("   #1 23456789#")
                p("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                p("    3.14159265,3589793 is from the web") 
                p("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start4)
                fmt.Print("  600,000,000 iterations in ", elapsed, " yields 9 digits of π\n")
                p(" ")
                }
// case 4
                if iterInt64 == 1000000000 {p("... 1 Billion iterations completed, still nine ...")
                p(" ")
                p("   #1 23456789#")
                p("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                p("    3.14159265,3589793 is from the web") 
                p("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start4)
                fmt.Print("  1,000,000,000 iterations in ", elapsed, " yields 9 digits of π\n")
                p(" ")
                }
//
                if iterInt64 == 2000000000 {p("... 2 Billion, and still just nine ...")
                p(" ")
                p("   #1 23456789#")
                p("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                p("    3.14159265,3589793 is from the web") 
                p("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start4)
                fmt.Print("  2,000,000,000 iterations in ", elapsed, " yields 9 digits of π\n")
                p(" ")
                }
//
// case 4
                if iterInt64 == 4000000000 {
                p("... 4 Billion, gets us ten digits  ...")
                p(" ")
                p("   #1 234567890#")
                p("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                p("    3.141592653,589793 is from the web") 
                p("   #1 234567890 123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start4)
                fmt.Print("  4,000,000,000 iterations in ", elapsed, " yields 10 digits of π\n")
                p(" ")
                p(" per option #", num, "  --  the Gottfried Wilhelm Leibniz formula\n")

                LinesPerIter = 14
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")
// case 4
                fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                    check(err1)                                // ... gets a file handle to dataFromConstants.txt
                    defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.

                Hostname, _ := os.Hostname()
                _ , err0 := fmt.Fprintf(fileHandle, "\n  -- Gottfried Wilhelm Leibniz -- case #%d on %s \n", num, Hostname)  
                    check(err0)
                _ , err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second from case %d \n", LinesPerSecond, num) 
                    check(err2)
                _ , err3 := fmt.Fprintf(fileHandle, "%.04f was total Elapsed Seconds from case %d \n", elapsed.Seconds(), num) 
                    check(err3)
                _ , err4 := fmt.Fprintf(fileHandle, "%e was Iterations/Seconds from case %d \n", iterFloat64/elapsed.Seconds(), num)
                    check(err4)
                _ , err5 := fmt.Fprintf(fileHandle, "%e was total Iterations from case %d \n", iterFloat64, num)
                    check(err5)
                TotalRun := elapsed.String() // cast time duration to a String type for Fprintf "formatted print"
                _ , err7 := fmt.Fprintf(fileHandle, "Total runTime was %s \n", TotalRun)  // add total runtime of this calculation 
                    check(err7)

                fmt.Println("Select 100 at menu to display prior results")
                }
//case 4 
/* we skip the rest

                if iterInt64 == 6000000000 {p("... 6 Billion completed.  ...")
                p(" ")
                p("   #1 234567890#")
                p("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                p("    3.141592653,589793 is from the web") 
                p("   #1 234567890 123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start4)
                fmt.Print("  6,000,000,000 iterations in ", elapsed, " still yields 10 digits of π\n")
                p(" ")
                }
// case 4
                if iterInt64 == 8000000000 {p("... 8 Billion completed. still ten ...")
                p(" ")
                p("   #1 234567890#")
                p("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                p("    3.141592653,589793 is from the web") 
                p("   #1 234567890 123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start4)
                fmt.Print("  8,000,000,000 iterations in ", elapsed, " still yields 10 digits of π\n")
                p(" ")
                }
//
            if iterInt64 == 9000000000 {
                p(" ")
                p("   #1 234567890#")
                p("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                p("    3.141592653,589793 is from the web") 
                p("   #1 234567890 123456# :: counting the first 16 actual digits of π")
                // fmt.Print("   ", iter) 
                t := time.Now()
                elapsed := t.Sub(start4)
                fmt.Print("\n... 9B iterations in ", elapsed, " , but to get 10 digits we only needed 4B iterations\n\n") 
                p(" per option #", num, "  --  the Gottfried Wilhelm Leibniz formula")
            }  
*/
    } // end of for


case 5:  // -----------------------------------------------------------------------------------------

    start6 := time.Now()

            p(" ")
            p("You selected #", num, "  --  An improved version of Archimedes' method")
            p("  -- enter \"6' at the main menu for the derivation and proof of the Pythagorean -- ")
                        // the above escape does not seem to work as advertised ??

            var r float64 // radius is a constant 1 
            var a float64 // height of bisected triangle 
            var s1 float64 // starts-out as 1
            var s1_2 float64 // s1/2
            var s2 float64 // new hypotenuse / new side
            var b float64 // new short side
            var n float64 // number of sides 
            var p float64 // perimeter
            var p_d float64 // pi

            r = 1 // a constant, the radius 

            n = 6 // the number of sides of the polygon 
            s1 = 1 // starts-out as 1

            s1_2 = s1/2
            a = math.Sqrt(r-(math.Pow(s1_2,2))) // height of bisected triangle 
            b = 1-a // new short side
            s2 = math.Sqrt(math.Pow(b,2) + (math.Pow(s1_2,2))) // new hypotenuse / NewSide
            p = n*s1 
            p_d = p/2 

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

// case 5

for iterInt64 < 27 {
    iterInt64 = iterInt64 + 1 
    iterFloat64 = iterFloat64 + 1
    n = n*2
    s1 = s2 
    s1_2 = s1/2 
    a = math.Sqrt(r-(math.Pow(s1_2,2))) 
    b = 1 - a 
    s2 = math.Sqrt(math.Pow(b,2) + (math.Pow(s1_2,2))) 
    p = n * s1 
    p_d = p/2 

    // 3.141592653589793238,46264338327950288419716939937510 -- is just way-too-much pie! 

    if iterInt64 == 2 {
        fmt.Println(" ")
                fmt.Println("   #1 2# :: counting the first 2 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 2# :: counting the first 3 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start6)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 2 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", n)
    }    
// case 5
    if iterInt64 == 3 {
        fmt.Println(" ")
                fmt.Println("   #1 23# :: counting the first 3 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 23# :: counting the first 3 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start6)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 2 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", n)

    }    
    if iterInt64 == 7 {
        fmt.Println(" ")
                fmt.Println("   #1 23456# :: counting the first 6 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 23456# :: counting the first 6 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start6)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 5 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", n)
    }    
    if iterInt64 == 12 {
        fmt.Println(" ")
                fmt.Println("   #1 23456789# :: counting the first 9 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 23456789# :: counting the first 9 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start6)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 8 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", n)
    }    
    if iterInt64 == 17 {
        fmt.Println(" ")
                fmt.Println("   #1 23456789012# :: counting the first 12 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 23456789012# :: counting the first 12 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start6)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 12 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", n)
    }    
    if iterInt64 == 23 {
        fmt.Println(" ")
                fmt.Println("   #1 23456789012345# :: counting the first 15 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 23456789012345# :: counting the first 15 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start6)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 15 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", n)
    }
// case 5
    if iterInt64 == 25 {
        fmt.Println(" ")
                fmt.Println("   #1 234567890123456# :: counting the first 16 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes' -- 16 -- digits") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 234567890123456789# :: counting the first 19 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start6)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 16 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", n)
                
                fmt.Printf("%.0f as parsed against ...\n", n)
                fmt.Println("100000000 which is one-hundred-million, for comparison to the above line")
                fmt.Println("... Which is to say a 201,326,592 sided polygon, Mr. A. would have wept\n")

                fmt.Println(" per option #", num, "  --  an improved version of Archimedes' method\n")

// Don't know what the   codes from WikiPedia were supposed to do? But they do not display or have any obvious consequence.
/*arch := `Archimedes of Syracuse (c. 287 – c. 212 BC) was a Greek mathematician, physicist, engineer, 
astronomer, and inventor from the ancient city of Syracuse in Sicily.`
*/
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
            fmt.Println(arch) // prints the above rune
//case 5
                LinesPerIter = 18
                fmt.Println("\nat aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("Aprox %.0f lines of code were executed per second \n\n", LinesPerSecond) 

         fmt.Println("The above calculation was per the following psudo code ...\n")
         fmt.Println("   let n = n * 2 // n is number of sides. After having begun with 6, then do 12, etc. ")
         fmt.Println("   let s1 = s2  // s1 begins as 1, thereafter it adopts s2, base of unbisected triangle")
         fmt.Println("   s1_2 = s1/2 // set the variable s1_2 to one half of s1")
         fmt.Println("   a = \u221A( 1-(s1/2)\u00b2) ) // length of side 'a' is height of bisected triangle ")
         fmt.Println("   b = 1-a                    // new short side ")
         fmt.Println("   s2 = \u221A( (b)\u00b2 + (s1/2)\u00b2 )  // s2 is new hypotenuse, and New s1 ")
         fmt.Println("   p = n * s1 // p is the length of the perimeter of the constructed polygon ")
         fmt.Println("   p_d = p/2 // the diameter of the polygon is always two, so p/2 = π \n")
         fmt.Println("   -- After hitting Return for menu redisplay, enter '6' for the derivation and proof of the Pythagorean")
         fmt.Println(" ")

                fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                    check(err1)                                // ... gets a file handle to dataFromConstants.txt
                    defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
// case 5
                Hostname, _ := os.Hostname()
                _ , err0 := fmt.Fprintf(fileHandle, "\n  -- Archimedes of Syracuse -- case #%d on %s \n", num, Hostname)
                    check(err0)
                _ , err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second from case %d \n", LinesPerSecond, num) 
                    check(err2)
                _ , err3 := fmt.Fprintf(fileHandle, "%.04f was total Elapsed Seconds from case %d \n", elapsed.Seconds(), num) 
                    check(err3)
                _ , err4 := fmt.Fprintf(fileHandle, "%.02f was Iterations/Seconds from case %d \n", iterFloat64/elapsed.Seconds(), num)
                    check(err4)
                _ , err5 := fmt.Fprintf(fileHandle, "%.0f was total Iterations from case %d \n", iterFloat64, num)
                    check(err5)
                fmt.Println("Select 100 at menu to display prior results")
    }
}

case 6:  // ---- display a review of the derivation of the Pythagorean --------------------------------------------------


fmt.Println("\n\n\n -- You entered '", num, "' to review the derivation of the Pythagorean, which was needed in method #5. We will")
fmt.Println("be geometrically deriving the Pythagorean theorem according to the 12th century Indian, Bhaskara.")
fmt.Println("    We begin with a square of area c\u00b2. We then partially fill that square with four congruent")
fmt.Println("right triangles each with its right angle opposite to one of the sides 'c'. Each of the four congruent")
fmt.Println("triangles now have sides c, a, and b. 'c' being the hypotenuse of each; 'a' being the shortest side and 'b'")
fmt.Println("being the longer of sides 'a' and 'b'. Thus leaving a small square in the center, which we can ignore.")

fmt.Println("     Next, we detach and slide two adjacent right triangles, each to its opposite side of the parent square.")
fmt.Println("We now have two attached squares a x a and b x b. We can then clearly see that c\u00b2 = a\u00b2 + b\u00b2 ")
fmt.Println(" --- The Pythagorean theorem is thereby derived, per Bhaskara --- ")

fmt.Println("\n     Apparently it was not until 1876 that James Garfield (yea, the former U.S. president) discovered")
fmt.Println("the following geometric proof of the Pythagorean theorem. He began with a right triangle similar ")
fmt.Println("to those we encountered from Bhaskara. But any shape of right triangle will suffice for his proof. ")
fmt.Println("Next, he cloned a second triangle placing it such that its most-acutely-angled point touches the other")
fmt.Println("triangle's least-acutely-angled point such that side 'a' of the first is colinear with side 'c' of the")
fmt.Println("second triangle. Thus creating the potential of a new triangle: an isosceles right triangle of dimensions:")
fmt.Println("c, c, x. ")
fmt.Println("     After having converted that potential to an actual we are left with a trapezoid. That is, a rectangle")
fmt.Println("mated to a right triangle; in this case a trapezoid of dimensions a, a+b, b, and x. We know that for any")
fmt.Println("trapezoid the area is equal to the height [in this case a+b] times the average of the lengths of the ")
fmt.Println("two adjacent sides [in this case, half of the sum of a+b]. Which is to say that, in the case of OUR trapezoid, ")
fmt.Println("the Area is  :  A = (a+b) * 1/2(a+b)  -- Or: --  A = 1/2(a+b)\u00b2 ")
fmt.Println("     But the area of this particular trapezoid that we have constructed is, obviously, also a*b")
fmt.Println("(the combined area of the initial two right triangles), plus 1/2*c\u00b2.   Or:  -- ")
fmt.Println("A = ab + 1/2*c\u00b2 \n")
fmt.Println("           a new equality can then be constructed from the two above equalities: \n")
fmt.Println("    1/2(a+b)\u00b2  =  ab + 1/2(c\u00b2) \n")
fmt.Println("                          and reducing it thusly : ")
fmt.Println("        (a+b)\u00b2 = 2ab + c\u00b2 ")
fmt.Println(" a\u00b2 + 2ab + b\u00b2 = 2ab + c\u00b2 ")
fmt.Println("       a\u00b2 + b\u00b2 = c\u00b2 \n")
fmt.Println("... proves the Pythagorean per Garfield; though, obviously, many other proofs do exist.\n\n")
    fmt.Println("Select 100 at menu to display prior results")


case 7: // ------------------------------------------------------------------------------

p(" ")
p("   You selected #", num, " A Go language exercize which can be used to test the speed of your hardware.")
p("   We will calculate π to a maximum of ten digits of accuracy using an infinite series by John Wallis circa 1655")
p("   Up to 40 Billion iterations of the following formula will be executed ")
p("   π = 2 * ((2/1)*(2/3)) * ((4/3)*(4/5)) * ((6/5)*(6/7)) ...")

    start1 := time.Now() // saved start time to be compared with end time t
    //iterInt64 = 1
    //iterFloat64 = 0
    var cumulativeProduct float64
    var numerators float64
    var firstDenom float64

    numerators = 2
    firstDenom = 1
    secondDenom = 3
    cumulativeProduct = (numerators/firstDenom) * (numerators/secondDenom)

      for iterInt64 < 40000000000 {
        iterInt64 = iterInt64 + 1
        iterFloat64 = iterFloat64 + 1
        numerators = numerators + 2
        firstDenom = firstDenom + 2
        secondDenom = secondDenom + 2
        cumulativeProduct = cumulativeProduct * (numerators/firstDenom) * (numerators/secondDenom)
        π = cumulativeProduct * 2  
        if iterInt64 == 10 { p(" ") }  //  single-line "if" just to display a blank line; for amusement 

/*            if iterInt64 == 100 {
                p("   #1 2# ")
                p("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                p("    3.1,41592653589793 is, again, the value of π from the web") 
                p("   #1 2# 34567890123456 :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start1)
                p (iterInt64, " iterations were completed in ", elapsed, " yielding 2 digits of π\n") 
            }
// case 7
            if iterInt64 == 500 {
                p("   #1 23# ")
                p("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                p("    3.14,1592653589793 is, again, the value of π from the web") 
                p("   #1 23# 4567890123456 :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start1)
                p (iterInt64, " iterations were completed in ", elapsed, " yielding 3 digits of π\n") 
            }
*/
            if iterInt64 == 2000 {
                p("   #1 234# ")
                p("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                p("    3.14,1592653589793 is, again, the value of π from the web") 
                p("   #1 23 4567890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start1)
                p (iterInt64, " iterations were completed in ", elapsed, " yielding 4 digits of π\n") 
            }

            if iterInt64 == 10000 {
                p("   #1 2345# ")
                p("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                p("    3.1415,92653589793 is, again, the value of π from the web") 
                p("   #1 2345 67890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start1)
                p("10,000 iterations were completed in ", elapsed, " yielding 5 digits of π\n") 
            }
            if iterInt64 == 50000 { // 50,000 
                p("   #1 2345# ")
                p("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                p("    3.1415,92653589793 is, again, the value of π from the web") 
                p("   #1 2345 67890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start1)
                p("50,000 iterations were completed in ", elapsed, " yielding 5 digits of π\n") 
            }
            if iterInt64 == 500000 {  // 500,000 done
                p("   #1 23456# ")
                p("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                p("    3.14159,2653589793 is, again, the value of π from the web") 
                p("   #1 23456 7890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start1)
                p("500,000 iterations were completed in ", elapsed, " yielding 6 digits of π\n") 
            }
            if iterInt64 == 2000000 {  // 2M done 
                p("   #1 234567# ")
                p("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                p("    3.141592,653589793 is, again, the value of π from the web") 
                p("   #1 234567 890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start1)
                p("2,000,000 iterations were completed in ", elapsed, " yielding 7 digits of π\n") 
            }
            if iterInt64 == 40000000 {  // 40M done
                p("   #1 2345678# ")
                p("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                p("    3.1415926,53589793 is, again, the value of π from the web") 
                p("   #1 2345678 90123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start1)
                p("40,000,000 iterations were completed in ", elapsed, " yielding 8 digits of π\n\n") 
// case 7
                fmt.Println("  .. working .. on another factor-of-ten iterations\n")
            }

            if iterInt64 == 400000000 {  // 400M done
                p("   #1 23456789# ")
                p("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                p("    3.14159265,3589793 is, again, the value of π from the web") 
                p("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start1)
                p("400,000,000 iterations were completed in ", elapsed, " yielding 9 digits of π\n\n") 

                LinesPerIter = 36
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")
//
//
                p("  ... will be working on doing Billions more iterations ...\n\n")
            }
//
//
            if iterInt64 == 600000000 {  // 600M done
                p("  600M done, still working on another Two-Hundred-Thousand iterations ... working ...\n")
                t := time.Now()
                elapsed := t.Sub(start1)
                    p(elapsed)
                    fmt.Println(" ")
                fmt.Println("Calculating 10th digit (40B iters) which takes a few min - Ctrl-C to End/Exit\n")

                LinesPerIter = 36
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")
                fmt.Println(" ... still working ...")
            }

            if iterInt64 == 800000000 {  // 800M done
                 p("  800M done, still working on yet another Two Hundred Thousand iterations ... working ...\n")
                t := time.Now()
                elapsed := t.Sub(start1)
            p(elapsed, " \n")
            }
// case 7
//
        if iterInt64 == 1000000000 {  // 1B done
            p("   #1 23456789# ")
            p("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
            p("    3.14159265,3589793 is the value of π from the web") 
            p("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
            t := time.Now()
            elapsed := t.Sub(start1)
            p("\nOne Billion iterations were completed in ", elapsed, " still only yielding π to 9 digits\n") 
            p(" per option #", num, "  --  an infinite series by John Wallis circa 1655\n") // ----------------------

                LinesPerIter = 36
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")
//
            p("  ... still working ... on Billions of iterations, 39 to go ...\n\n")
        }

//  "Optional" Forty Billion loops just to get one additional digit, the tenth digit of pi 

        if iterInt64 == 2000000000 {  // 2B completed 
            p("  2B done, still working ... on another Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start1)
            p(elapsed)
        }
        if iterInt64 == 3000000000 {  // 3B completed
            p("  3B done, still working ... on another Billion iterations ... working ... Ctrl-C to End/Exit")
                t := time.Now()
                elapsed := t.Sub(start1)
            p(elapsed)
        }
        if iterInt64 == 4000000000 {  // 4B completed 
            p("  4B done, still working ... on another Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start1)
            p(elapsed)
        }
        if iterInt64 == 5000000000 {  // 5B completed 
            p("  5B done, still working ... on another Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start1)
            p(elapsed)
        }
// case 7
        if iterInt64 == 6000000000 {  // 6B completed 
            p("  6B done, still working ... on another Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start1)
            p(elapsed)
        }
        if iterInt64 == 7000000000 {  // 7B completed 
            p("  7B done, still working ... on another Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start1)
            p(elapsed)
        }
        if iterInt64 == 8000000000 {  // 8B completed 
            p("  8B done, still working ... on another Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start1)
            p(elapsed)
        }
        if iterInt64 == 9000000000 {  // 9B completed 
            p("  9B done, still working ... on another five Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start1)
            p(elapsed)
        }
        if iterInt64 == 14000000000 {  // 14B completed
            p("  14B done, still working ... on another five Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start1)
            p(elapsed)
        }
        if iterInt64 == 19000000000 {  // 19B completed 
            p("  19B done, still working ... on another five Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start1)
            p(elapsed)
        }
        if iterInt64 == 24000000000 { // 24B completed
            p("  24B done, still working ... on another five Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start1)
            p(elapsed)
        }
        if iterInt64 == 29000000000 {  // 29B completed 
            p("  29B done, still working ... on another five Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start1)
            p(elapsed)
        }
        if iterInt64 == 34000000000 {  // 34B completed
            p("  34B done, still working ... just another six Billion iterations to go! ... ")
                t := time.Now()
                elapsed := t.Sub(start1)
            p(elapsed)
        }
// case 7
 //  "Optional" Forty Billion loops to get another digit
        //
        if iterInt64 == 40000000000 {  // 40B completed 
            p("   #1 234567890# ")
            p("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
            p("    3.141592653,589793 is the value of π from the web") 
            p("   #1 234567890 123456# :: counting the first 16 actual digits of π")
            t := time.Now()
            elapsed := t.Sub(start1)
            p("Forty Billion iterations were completed in ", elapsed, " yielding π to 10 digits\n") 
            p(" per option #", num, "  --  an infinite series by John Wallis circa 1655\n") // ----------------------
// case 7
                LinesPerIter = 36
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")
// case 7
                fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                    check(err1)                                // ... gets a file handle to dataFromConstants.txt
                    defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
//case 7
                Hostname, _ := os.Hostname()
                _ , err0 := fmt.Fprintf(fileHandle, "\n  -- John Wallis -- case #%d on %s \n", num, Hostname)
                    check(err0)
                _ , err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second from case %d \n", LinesPerSecond, num) 
                    check(err2)
                _ , err3 := fmt.Fprintf(fileHandle, "%.04f was total Elapsed Seconds from case %d \n", elapsed.Seconds(), num) 
                    check(err3)
                _ , err4 := fmt.Fprintf(fileHandle, "%.02f was Iterations/Seconds from case %d \n", iterFloat64/elapsed.Seconds(), num)
                    check(err4)
                _ , err5 := fmt.Fprintf(fileHandle, "%e was total Iterations from case %d \n", iterFloat64, num)
                    check(err5)
                TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun)  // add total runtime of this calculation 
                    check(err7)
                fmt.Println("Select 100 at menu to display prior results")
        }
    }

// -------------------------------------------------------------------------------------------

case 8:

        var n float64 
        var Eulers float64
        var sum float64

        sum = 1 + 1/n 
        Eulers = math.Pow(sum, n)

        fmt.Println("\n\n\nEuler's Number \u2107 or \u2147 the natural logarithmic base")
        fmt.Println("\u2147 = (1+1/n)^n")
        fmt.Println("... the limit of an increasing value for n\n\n") 
// case 8
        n = 9
        sum = 1 + 1/n 
        Eulers = math.Pow(sum, n)
        fmt.Print(Eulers)
        fmt.Printf(" was calculated with an exponent of %0.f \n", n)
        n = 99
        sum = 1 + 1/n 
        Eulers = math.Pow(sum, n)
        fmt.Print(Eulers)
        fmt.Printf("  was calculated with an exponent of %0.f \n", n)
        sum = 1 + 1/n 
        Eulers = math.Pow(sum, n)
        n = 999
        sum = 1 + 1/n 
        Eulers = math.Pow(sum, n)
        fmt.Print(Eulers)
        fmt.Printf(" was calculated with an exponent of %0.f \n", n)
        n = 9999
        sum = 1 + 1/n 
        Eulers = math.Pow(sum, n)
        fmt.Print(Eulers)
        fmt.Printf(" was calculated with an exponent of %0.f \n", n)
        n = 99999999999
        sum = 1 + 1/n 
        Eulers = math.Pow(sum, n)
        fmt.Print(Eulers)
        fmt.Printf(" was calculated with an exponent of %0.f \n", n)
// case 8
        fmt.Println("2.71828182845904523536028747135266249775724 is Euler's Number from the web")
        fmt.Println("2.718281828 is the dollar value of $1 compounded continuously for one year.")
        fmt.Println("2.714567 is from daily compound interest which is near-enough to continuous interest.\n")

        fmt.Println("An account starts with $1.00 and pays 100 percent interest per year. If the interest is credited once, ")
        fmt.Println("at the end of the year, the value of the account at year-end will be $2.00. What happens if the interest")
        fmt.Println("is computed and credited more frequently during the year?\n")

        fmt.Println("If the interest is credited twice in the year, the interest rate for each 6 months will be 50%, so the ")
        fmt.Println("initial $1 is multiplied by 1.5 twice, yielding $2.25 at the end of the year. Compounding quarterly")
        fmt.Println("yields $2.44140625, and compounding monthly yields $2.613035 = $1.00 × (1 + 1/12)^12 Generally, if there")
        fmt.Println("are n compounding intervals, the interest for each interval will be 100%/n and the value at the end of")
        fmt.Println("the year will be $1.00 × (1 + 1/n)^n.")
//
// And, now here comes a fun rune to print a multi-line "string"
// ... define a rune with the ` :: back-quote character located on the ~ tilda key
Ricks_rune_Paragraph := `  
Bernoulli noticed that this sequence approaches a limit (the force of interest) with larger n and, thus, smaller 
compounding intervals. Compounding weekly (n = 52) yields $2.692596..., while compounding daily (n = 365) yields
$2.714567... (approximately two cents more). The limit as n grows large is the number that came to be known as e.
That is, with continuous compounding, the account value will reach $2.718281828 
`
fmt.Println(Ricks_rune_Paragraph)

// ----------------------------------------------------------------------------------------------------

case 9: 
    fmt.Println(" ... So sorry, but", num, "was not an option")

case 10: // -------------------------------------------------------------------------------------------
    fmt.Println(num, " was not an option! It was not on the menu, go fish!\n")

case 11: 
    fmt.Println("Your selection of", num, " is right-out!  Go Fish!\n")

// the num switched-on is set to 12 for all selections greater than 11 & < 100 
case 12: 
    fmt.Println("Your selection is really-far-out!  Go Fish!\n")

default: // ------------------------------------------------------------------------------
        p(" ")
        // fmt.Println("this is the switch default code after the break option ")
        fmt.Println("You failed to make a selection, Hit Enter/Return to redisplay the menu, Ctrl-C to End/Exit")
    break 

p(" ")
fmt.Println("Oops, how'd we get here? Hit Enter/Return again to possibly redisplay the menu")
} // end of switch
} // end of func RicksSwitch(), and end of 29 pages, 8,474 words, 64,610 characters 