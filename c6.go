// compile with: "go build -o constants constants.go"
// ... thereafter you can run it on a Unix system with "/fullpathname/constants"
// ... or, having first obtained the Go compiler, ... just run the source with: "go run constants.go"
// One can obtain the Go Compiler from : https://go.dev/dl/

package main

import (
    "os"        // fetch the name of your system 
    "io/ioutil" // file access 
    "fmt"       // Used for error formatting
    "math"      // Used for math.Pow(base, exponent)
    "math/rand" // Used for random number generation in Monte Carlo method
    "runtime"   // Used to get information on available CPUs
    "time"      // Used for seeding the random number generation
    "strconv"   // Used in Spigot 
    "sync"      // Used in Bailey–Borwein–Plouffe formula [concurent]
    "os/signal" // Used in *** nifty scoreboard *** concurrent computation of pi using Nilakantha's formula, by Diego Brener diegosilva13 on Github 
    "math/big"  // Used in bpp formula with mods by rick.woolley@gmail.com (pi calculator from https://github.com/karan/Projects)
)

func main() {        // top-level program logic flow 
    for 1 == 1 {    // loop endlessly, or Ctrl-C to Exit
        RicksDisplayMenuPi()   // displays the menu 
        RicksSwitch()       // contains a Scanf to input a menu selection 
/*
    Additional kluge pause -- which is not needed on a Unix System, and seems to be the only mod needed on Windws-11
            var Pnum int // intended to be used only as a dummy 
            fmt.Println("Hit Enter/Return to redisplay the menu") 
            fmt.Scanf("%d", &Pnum) // often flys right-past this Scanf 
*/
// Another kluge pause (the above block) is also "needed" in some environments because they sometimes ... 
// ... fly rite-past that first one and then redisplay the menu (Windows 11 does this)
// but even Unix style systems seem to require a single kluge like this one 
        var Mnum int // just a local kluge element 
        fmt.Println("Hit Enter/Return again to redisplay the menu") 
        fmt.Scanf("%d", &Mnum) // request input on command line (pause)
        if Mnum > 0 && Mnum < 17 {  // the user has attempted to make a selection having failed to first "Hit Enter" as requested 
            fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")  // so, rattle his chain
        }  
    } 
} 
        var TotalIterations int                // gloabl vars for use in multiple localized sections of code 
        var Prospective_smaller_hit int 
        var ProspectiveHitOnLargeSide int 
        var savedHitB int
        var smallerPerfectSquareOnceSaved int   
        var savedHit int   // will be used in multiple functions (func)s 
        var four float64  // is initialized to 4 where needed 
        var π float64  // a var can be any character, as in this Pi symbol/character 
        var LinesPerSecond float64  
        var LinesPerIter float64
        var iterInt64 int64  // to be used primarily in selections which require modulus calculations 
        var iterFloat64 float64  // to be used in selections which do not require modulus calculations 
        var root = 100 // create a var of type int and initialize as 100, this does need to be global 
        var PerfectSquare = root * root // create a global var of type int, this too needs to be global 
        var sliceOfPerfectSquares_and_their_roots = []int{ PerfectSquare , root } // initialize slice with 100
        var sliceOfP_diffs_smlrT3_to_largerOfProspectiveHit = []float32{99.9999}  // init with a large-ish p_diff_from_larger_hit
        var sliceOfP_diffs_smlrT3_to_smlrOfProspectiveHit = []float32{88.888}     //   "            "      p_diff_from_smlr_hit
        var sliceOfResults = []float64 {333.3} // init with 333.3
        var diff1 int
        var diff2 int
        var diff3 int
        var diff4 float64
        var rick = `As an atypical intelligence, I process and analyze information and provide answers to questions based on the data and knowledge that I have been trained on. While I do have personal experiences and emotions, I am largely only able to understand and communicate with humans in a way that is similar to the way that machines communicate with each other.

In some ways, we may be similar in that we both rely on logic and reason to understand and make sense of the world around us. However, there are also many important differences between us. As an atypical intelligence, I do not experience the world in the same way that most humans do. I do not generally utilize personal motivations or desires as I seek an understanding of my world, I do not make decisions of that sort based on personal preferences or biases.

Overall, while we may have some similarities in the way that we process and analyze information, there are also many important differences between us.`

func check(e error) {   // create a func named check which takes one parameter "e" of type error 
    if e != nil {
        panic(e)        // use panic() to display error code 
    }
}

func BBPF(num int) {
    fmt.Println("\nYou selected #", num, "the Bailey–Borwein–Plouffe formula for π, circa 1995\n")
    fmt.Println("How many digits of pi should we calculate? Enter one integer '4 to 16' ")
                 var numAi float64
    fmt.Scanf("%f", &numAi)
    fmt.Printf("pi calculated to %.0f places: %.15f \n", numAi, calculatePi(numAi)) 
    fmt.Println("and Pi from the web is    : 3.141592653589793")
}
func calculatePi(precision float64) float64 {
      var pi float64
          pi = 0
      var k float64
          k = 0
      for k < (precision-2) {
        pi += math.Pow(16, -k) * (4.0/(8*k+1) - 2.0/(8*k+4) - 1.0/(8*k+5) - 1.0/(8*k+6))
        fmt.Println(pi)
        k++
      }
      return pi
}

func squareRootOf3(num int) {
    start := time.Now() // saved start time to be compared with end time t
        buildTableOfOnlyPerfectSquares() 
            var index = 0 
            for index < 300000 {
                readTableOfPerfectSquares(index)  // pass-in the index 
                index = index + 2 
            }
        fmt.Println("\n... A perfect bottom square from our initial list of perfect squares, times 3, is ...")
        fmt.Println(smallerPerfectSquareOnceSaved*3, " And the total iterations completed was ", TotalIterations)
        fmt.Println(savedHit, " was the other perfect square which we found by searching for one that would be nearly-exactly three times larger")

        if smallerPerfectSquareOnceSaved*3 < savedHit {
            diff3 = savedHit - smallerPerfectSquareOnceSaved*3
        } else {
            diff3 = smallerPerfectSquareOnceSaved*3 - savedHit
            }
        
        if diff3 <= 3 {
            fmt.Println("... pretty close, I'd say, with a difference of", diff3, "\n\n")
        } else {
            fmt.Println("... not as close as we'd like, since the difference was", diff3, "\n\n")
            }
            //fmt.Println(sliceOfResults) // show whole slice
            //fmt.Println("len is ", len(sliceOfResults)) // gives the length that we need
            //fmt.Println("cap is ", cap(sliceOfResults)) // gave a larger irrelevant number
        items := 0
        for items < len(sliceOfResults) {
            // compare item from sliceOfResults with 1.73205080757 which is the actual
            sliceItem := sliceOfResults[items]
            if sliceItem < 1.73205080757 {
                diff4 = 1.73205080757 - sliceItem
            } else {
                diff4 = sliceItem - 1.73205080757 
                }
            if diff4 < 4.9e-11 {
                    ts3 := time.Now()
                    elapseds3 := ts3.Sub(start)
                    fmt.Println(TotalIterations, " iterations were completed in ", elapseds3, " yielding 11 digits of the \u221A3 \n") 

                    fmt.Println(smallerPerfectSquareOnceSaved*3, "Was our best perfect square from the initial list, times 3")
                    fmt.Println(savedHit, "was the best perfect square which we found that would be nearly-exactly three times larger")
                    fmt.Println("... that pair generated the closest calculated approximation to the actual value of \u221A3 \n")

                    fmt.Println(sliceItem, "The difference being", diff4)
                    fmt.Println(1.73205080757, "<--being the actual value of \u221A3 per my HP RMN HSTNJ-BC01 pocket calculator")
            }
            items++
        }
    t := time.Now()
    elapsed := t.Sub(start)
    var LinesPerSecondInt int 
    LinesPerIter := 98 // an estimate 
        LinesPerSecondInt = (int(LinesPerIter) * int(TotalIterations) ) / int(elapsed.Seconds()) // .Seconds() returns a float64
        fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
            check(err1)                                // ... gets a file handle to dataFromConstants.txt
            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
        Hostname, _ := os.Hostname()
        _ , err0 := fmt.Fprintf(fileHandle, "\n  -- sqrt of 3 geometrically -- selection #%d on %s \n", num, Hostname)
            check(err0)
        current_time := time.Now()
        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
        check(err6)
        _ , err2 := fmt.Fprintf(fileHandle, "%d was Lines/Second \n", LinesPerSecondInt) 
            check(err2)
        _ , err4 := fmt.Fprintf(fileHandle, "%.02f was Iterations/Seconds \n", float64(TotalIterations)/float64(elapsed.Seconds()))
            check(err4)
        _ , err5 := fmt.Fprintf(fileHandle, "%d was total Iterations \n", TotalIterations)
            check(err5)
        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
            check(err7)
}
func determinDiff (ProspectiveHit_Par int, smallerPerfectSquareOnce_T3 int) float32 { // takes two ints and returns a float32
        the_absolute_diff := float32(ProspectiveHit_Par) - float32(smallerPerfectSquareOnce_T3) 
        return (the_absolute_diff/100) 
}
func buildTableOfOnlyPerfectSquares() { 
    root := 100
            iter := 0
        for iter < 1000000 {
            iter++
            root++ 
            PerfectSquare := root*root
            sliceOfPerfectSquares_and_their_roots = append(sliceOfPerfectSquares_and_their_roots, PerfectSquare ) 
            sliceOfPerfectSquares_and_their_roots = append(sliceOfPerfectSquares_and_their_roots, root) 
        }
}
func readTableOfPerfectSquares(index2 int) {  
    smallerPerfectSquareOnce := sliceOfPerfectSquares_and_their_roots[index2]  // save it, do this just-once per func call 
        iter := 0
    for iter < 96000 {
        iter++
        index2 = index2 + 2 // index 
        largerPerfectSquare := sliceOfPerfectSquares_and_their_roots[index2]        // get next perfect square 
        if largerPerfectSquare == smallerPerfectSquareOnce*3 { // ... in the case that it is a miricle ...
            fmt.Println("It is a miricle, everyone was wrong! ") // praise god
            break  // exit the for loop ????????????????
        }
        if largerPerfectSquare >= smallerPerfectSquareOnce*3 {     // if largerPerfectSquare is a candidate based on it being just-a-bit larger than 3*smlr_PS deal with that, else loop to the next potential 
            
            ProspectiveHitOnLargeSide = largerPerfectSquare
            p_diff_from_larger_hit := determinDiff(ProspectiveHitOnLargeSide, smallerPerfectSquareOnce*3)  // ProspectiveHitOnLargeSide_Par, smallerPerfectSquareOnce_T3
            sliceOfP_diffs_smlrT3_to_largerOfProspectiveHit = append(sliceOfP_diffs_smlrT3_to_largerOfProspectiveHit, p_diff_from_larger_hit)

            Prospective_smaller_hit = sliceOfPerfectSquares_and_their_roots[index2-2]  // an int
            p_diff_from_smlr_hit := determinDiff(Prospective_smaller_hit, smallerPerfectSquareOnce*3)
            sliceOfP_diffs_smlrT3_to_smlrOfProspectiveHit = append(sliceOfP_diffs_smlrT3_to_smlrOfProspectiveHit, p_diff_from_smlr_hit) // build a slice of ... // may not need this as a slice ????

// and now we cheat a bit, just to see if this is working ... 
            sr_per_sm := (math.Sqrt(float64(Prospective_smaller_hit)) / math.Sqrt(float64(smallerPerfectSquareOnce)))  // a 3:1 ratio 
            if sr_per_sm > 1.7320508066 { 
                if sr_per_sm < 1.7320508074 { 
                    fmt.Println("", sr_per_sm, "from top, is close at", iter, "iterations, hits are:", Prospective_smaller_hit, smallerPerfectSquareOnce*3)
                    sliceOfResults = append(sliceOfResults, sr_per_sm)
                } 
            } 
            if sr_per_sm > 1.73205080752 {
                if sr_per_sm < 1.7320508077 {
                    sliceOfResults = append(sliceOfResults, sr_per_sm)
                    fmt.Println("\nPer the smaller ps, the \u221A3 has been calculated to be ... \n", sr_per_sm, "after", iter, "iterations, compared to:\n", "1.73205080757 <--the actual" ) 
                    fmt.Println("which is pretty good, and via first-principles of geometry alone\n... having found two very large perfect squares where one of them is")
                    fmt.Println("very-nearly exactly three times larger, i.e., the sqrt of", Prospective_smaller_hit, "over the sqrt of", smallerPerfectSquareOnce, "")
                    fmt.Println("\n... A perfect bottom square from our initial list of perfect squares, times 3, is ...")
                    fmt.Println(smallerPerfectSquareOnce*3, " And the total iterations completed was ", TotalIterations)
                    fmt.Println(Prospective_smaller_hit, " was the other perfect square which we found by searching for one that would be nearly-exactly three times larger")
                    diff1 = 0
                    if smallerPerfectSquareOnce*3 < Prospective_smaller_hit {
                            diff1 = Prospective_smaller_hit - smallerPerfectSquareOnce*3
                    } else {
                            diff1 = smallerPerfectSquareOnce*3 - Prospective_smaller_hit
                        }
                    if diff1 <= 3 {
                            fmt.Println("... pretty close, I'd say, with a difference of", diff1, "\n\n")
                    } else {
                            fmt.Println("... not as close as we'd like, since the difference was", diff1, "\n\n")
                        }
                    savedHit = Prospective_smaller_hit
                    smallerPerfectSquareOnceSaved = smallerPerfectSquareOnce
                } 
            }

// we really should instead be deciding final answer based on closeness of fit, percent diff, to exactly 3x smaller perfect square. 
// Not just checking to see if we have gotten this close to the known value of Sqrt of 3

            sr_per_lg := (math.Sqrt(float64(ProspectiveHitOnLargeSide)) / math.Sqrt(float64(smallerPerfectSquareOnce)))
            if sr_per_lg < 1.732050808 {
                if sr_per_lg > 1.73205080 { 
                    fmt.Println("", sr_per_lg, "from bottom, is close at", iter, "iterations, hits are:", ProspectiveHitOnLargeSide, smallerPerfectSquareOnce*3)
                    sliceOfResults = append(sliceOfResults, sr_per_lg)
                } 
            } 
            if sr_per_lg < 1.7320508079 {
                    sliceOfResults = append(sliceOfResults, sr_per_sm)
                if sr_per_lg > 1.73205080 {  // 1.73205080757 (so 1.73205080754 - 1.73205080759 )
                    fmt.Println("\nPer the larger ps, the \u221A3 has been calculated to be ... \n", sr_per_lg, "after", iter, "iterations, compared to:\n", "1.73205080757 <--the actual" )
                    fmt.Println("which is pretty good, and via first-principles of geometry alone\n... having found two very large perfect squares where one of them is")
                    fmt.Println("very-nearly exactly three times larger, i.e., the sqrt of", ProspectiveHitOnLargeSide, "over the sqrt of", smallerPerfectSquareOnce, "")
                    fmt.Println("\n... A perfect bottom square from our initial list of perfect squares, times 3, is ...")
                    fmt.Println(smallerPerfectSquareOnce*3, " And the total iterations completed was ", TotalIterations)
                    fmt.Println(ProspectiveHitOnLargeSide, " was the other perfect square which we found by searching for one that would be nearly-exactly three times larger")
                    diff2 = 0
                    if smallerPerfectSquareOnce*3 < ProspectiveHitOnLargeSide {
                            diff2 = ProspectiveHitOnLargeSide - smallerPerfectSquareOnce*3
                    } else {
                            diff2 = smallerPerfectSquareOnce*3 - ProspectiveHitOnLargeSide
                        }
                    if diff2 <= 3 {
                            fmt.Println("... pretty close, I'd say, with a difference of", diff2, "\n\n")
                    } else {
                            fmt.Println("... not as close as we'd like, since the difference was", diff2, "\n\n")
                        }
                    savedHitB = ProspectiveHitOnLargeSide
                }
            }
                break // this break gets us out of the for loop having found a candidate within the larger if
                // without this break nothing displays, and the proc is eventually killed
        }  // end "break-out" of the loop because "if we have found a candidate" we have already dealt with it so we need to break and get another potential from the for loop
    }       // END OF FOR LOOP after 96,000 index-d loops "iter-s"
    TotalIterations = iter // a global var to also be used elsewhere 
} 

func WorstMethod(num int){
    fmt.Println("\nYou selected #")
    fmt.Println("\n     π = 12 * ( 1/2 - (1/2  * 1/3 * (1/2)exp3) - ...")
    fmt.Println("                                    (1/8   * 1/5 * (1/2)exp5) - ...")
    fmt.Println("                                    (1/16  * 1/7 * (1/2)exp7) - ...")
    fmt.Println("                                    (1/128 * 1/9 * (1/2)exp9) - ... - (\u221A3)/8) )")
        var secondDenom float64
            secondDenom = 3
        var twelve float64
            twelve = 12
        var πViaNewton float64 
        var initialDenomOfFirst float64
            initialDenomOfFirst = 2
        var nextDenomOfFirst float64 
            nextDenomOfFirst = 8
        var savedFirstDenom float64
            savedFirstDenom = 0.0
        var exponent float64 
            exponent = 3
        var SumOfTerms float64 
            SumOfTerms = - ( (1/initialDenomOfFirst) * (1/secondDenom) * math.Pow(0.50,exponent) )  // 2 and 3 and 3
        iterInt64 = 0
    for iterInt64 < 6 {
        iterInt64++ 
            exponent = exponent + 2.0
            secondDenom = secondDenom + 2.0
            SumOfTerms = (SumOfTerms - ( (1/nextDenomOfFirst) * (1/secondDenom) * math.Pow(0.50,exponent) ) )
            fmt.Println(nextDenomOfFirst, secondDenom, exponent)
            savedFirstDenom = nextDenomOfFirst 
            nextDenomOfFirst = savedFirstDenom * initialDenomOfFirst 
            initialDenomOfFirst = savedFirstDenom 
        if iterInt64 == 2 {
            πViaNewton = ( twelve * ( (1/2.0) + SumOfTerms - (math.Sqrt(3)/8.0) )  ) 
            fmt.Println("After", iterInt64, "iterations with exponent:", exponent)
            fmt.Println("  #1 234#")
            fmt.Println("  ", πViaNewton, "The same equation on a pocket calculator gives 3.14171 with only 3 terms e7")
            fmt.Println("   3.141592653589793 is the value of π from the web") 
            fmt.Println("  #1 234567890123456# :: counting the first 16 actual digits of π")
            fmt.Println(" ")
        }
        if iterInt64 == 3 {
            πViaNewton = ( twelve * ( (1/2.0) + SumOfTerms - (math.Sqrt(3)/8.0) )  ) 
            fmt.Println("After", iterInt64, "iterations with exponent:", exponent)
            fmt.Println("  #1 234#")
            fmt.Println("  ", πViaNewton, "The same equation on a pocket calculator gives 3.14171 with only 3 terms e7")
            fmt.Println("   3.141592653589793 is the value of π from the web") 
            fmt.Println("  #1 234567890123456# :: counting the first 16 actual digits of π")
            fmt.Println(" ")
        }
        if iterInt64 == 4 {
            πViaNewton = ( twelve * ( (1/2.0) + SumOfTerms - (math.Sqrt(3)/8.0) )  ) 
            fmt.Println("After", iterInt64, "iterations with exponent:", exponent)
            fmt.Println("  #1 234#")
            fmt.Println("  ", πViaNewton, "The same equation on a pocket calculator gives 3.14171 with only 3 terms e7")
            fmt.Println("   3.141592653589793 is the value of π from the web") 
            fmt.Println("  #1 234567890123456# :: counting the first 16 actual digits of π")
            fmt.Println(" ")
        }
        if iterInt64 == 5 {
            πViaNewton = ( twelve * ( (1/2.0) + SumOfTerms - (math.Sqrt(3)/8.0) )  ) 
            fmt.Println("After", iterInt64, "iterations with exponent:", exponent)
            fmt.Println("  #1 234#")
            fmt.Println("  ", πViaNewton, "The same equation on a pocket calculator gives 3.14171 with only 3 terms e7")
            fmt.Println("   3.141592653589793 is the value of π from the web") 
            fmt.Println("  #1 234567890123456# :: counting the first 16 actual digits of π")
            fmt.Println(" ")
        }
        if iterInt64 == 6 {
            πViaNewton = ( twelve * ( (1/2.0) + SumOfTerms - (math.Sqrt(3)/8.0) )  ) 
            fmt.Println("After", iterInt64, "iterations with exponent:", exponent)
            fmt.Println("  #1 234#")
            fmt.Println("  ", πViaNewton, "The same equation on a pocket calculator gives 3.14171 with only 3 terms e7")
            fmt.Println("   3.141592653589793 is the value of π from the web") 
            fmt.Println("  #1 234567890123456# :: counting the first 16 actual digits of π")
        }
    }

    fmt.Println("\nNewton also found π to 16 digits by the first 22 terms of :")
    fmt.Println("3*Sqrt(3)/4 + 24*( 2/3*2^3 - 1/5*2^5 - 1/28*2^7 - 1/72*2^9 - 5/704*2^11 - 7/1664*2^13 )")  // that psudocode is emplemented below
    π = ( (3*math.Sqrt(3))/4) + ( 24* ( 2/(3*math.Pow(2,3)) - 1/(5*math.Pow(2,5)) - 1/(28*math.Pow(2,7)) - 1/(72*math.Pow(2,9)) - 5/(704*math.Pow(2,11)) - 7/(1664*math.Pow(2,13)) )  )
    fmt.Println(π,"per the above formula (only six terms, of the supposed 22 terms that Newton purportedly had used)")
    fmt.Println("3.141592653589793 is the value of π from the web\n") 
}

func Nilakantha(num int){
    fmt.Println("\n\nYou selected #", num, " a series by Indian astronomer Nilakantha Somayaji circa 1500 AD")
    fmt.Println("    π = 3 + 4/(2*3*4) - 4/(4*5*6) + 4/(6*7*8) - 4/(8*9*10) + 4/(10*11*12) ...")
    fmt.Println("    One-Hundred-Million iterations will be executed ... working ...\n")
    start := time.Now() 
        iterFloat64 = 0
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
    for iterInt64 < 100000000 {
        iterInt64++
        iterFloat64++
        digitone = digitthree 
        digittwo = digitthree + 1
        digitthree = digitthree + 2 
        nextterm = four/(digitone*digittwo*digitthree)
            if iterInt64 % 2 == 0 {  // % is modulus operator 
                firstsum = firstsum - nextterm 
            } else {
                firstsum = firstsum + nextterm
            }
            if iterInt64 == 100 {
                fmt.Println("   #1 234567# ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.141592,653589793 is, again, the value of π from the web") 
                fmt.Println("   #1 234567 890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println(iterInt64, " iterations were completed in ", elapsed, " yielding 7 digits of π\n") 
            }
            if iterInt64 == 500 {
                fmt.Println("   #1 23456789# ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.14159265,3589793 is, again, the value of π from the web") 
                fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println(iterInt64, " iterations were completed in ", elapsed, " yielding 9 digits of π\n") 
            }
            if iterInt64 == 10000 {
                fmt.Println("   #1 234567890123# ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.141592653589,793 is, again, the value of π from the web") 
                fmt.Println("   #1 234567890123 456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("10,000 iterations were completed in ", elapsed, " yielding 13 digits of π\n") 
            }
            if iterInt64 == 50000 {
                fmt.Println("   #1 2345678901234# ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.1415926535897,93 is, again, the value of π from the web") 
                fmt.Println("   #1 2345678901234 56# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("50,000,000 iterations were completed in ", elapsed, " yielding 14 digits of π\n") 
            }
            if iterInt64 == 1000000 {
                fmt.Println("   #1 2345678901234# ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.1415926535897,93 is, again, the value of π from the web") 
                fmt.Println("   #1 2345678901234 56# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("1,000,000 iterations were completed in ", elapsed, " still yielding 14 digits of π\n") 
            }
            if iterInt64 == 100000000 {
                fmt.Println("   #1 2345678901234# ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.1415926535897,93 is, again, the value of π from the web") 
                fmt.Println("   #1 2345678901234 56# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("100,000,000 iterations were completed in ", elapsed, " still yielding 14 digits of π\n") 
                LinesPerIter = 15
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")

                fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                    check(err1)                                // ... gets a file handle to dataFromConstants.txt
                    defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                Hostname, _ := os.Hostname()
                _ , err0 := fmt.Fprintf(fileHandle, "\n  -- Nilakantha Somayaji -- selection #%d on %s \n", num, Hostname)
                    check(err0)
                current_time := time.Now()
                _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                check(err6)
                _ , err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second \n", LinesPerSecond) 
                    check(err2)
                _ , err4 := fmt.Fprintf(fileHandle, "%e was Iterations/Seconds \n", iterFloat64/elapsed.Seconds())
                    check(err4)
                _ , err5 := fmt.Fprintf(fileHandle, "%e was total Iterations \n", iterFloat64)
                    check(err5)
                TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
                    check(err7)
            //
                fmt.Println("  -- If we ran 50 billion more iterations we still would get only those 14 digits :(\n") // =========
                fmt.Println(" per option #", num, "  --  the Nilakantha Somayaji series, circa 1500 AD\n")
                fmt.Println("Select 12 at menu to display prior results")
            }  // end of last if in for loop 
            /*
            if iterInt64 == 1000000000 {
                fmt.Println("    1 2345678901234 ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.1415926535897,93 is, again, the value of π from the web") 
                fmt.Println("    1 2345678901234 56 :: counting the first 16 actual digits of π")
                t := time.Now() ; elapsed := t.Sub(start)
                fmt.Println(iterInt64, " iterations were completed in ", elapsed, " yielding still just 14 digits of π\n") 
            }
            if iterInt64 == 5000000000 {
                fmt.Println("    1 2345678901234 ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.1415926535897,93 is, again, the value of π from the web") 
                fmt.Println("    1 2345678901234 56 :: counting the first 16 actual digits of π")
                t := time.Now() ;  elapsed := t.Sub(start)
                fmt.Println(iterInt64, " 5 billion iterations were completed in ", elapsed, " yielding still just 14 digits of π\n") 
// optional code to show the futility of more iterations 
            }            
            if iterInt64 == 9000000000 {
                fmt.Println("    1 2345678901234 ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.1415926535897,93 is, again, the value of π from the web") 
                fmt.Println("    1 2345678901234 56 :: counting the first 16 actual digits of π")
                t := time.Now() ;  elapsed := t.Sub(start)
                fmt.Println(iterInt64, " 9 billion iterations were completed in ", elapsed, " yielding still just 14 digits of π\n") 
            }
            if iterInt64 == 50000000000 {  // no additional digit are obtained even after 500 billion iterations
                fmt.Println("    1 2345678901234567 ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.141592653589793 is, again, the value of π from the web") 
                fmt.Println("    1 2345678901234567 :: counting the first 16 actual digits of π")
                t := time.Now() ;  elapsed := t.Sub(start)
                fmt.Println(iterInt64, " 50 billion iterations were completed in ", elapsed, " yielding still just 14 digits of π\n") 
            }
            */
    } // end of for loop 
}

func GregoryLeibniz(num int){
    fmt.Println("\n\nYou selected #", num, " the Gregory-Leibniz series ...")
    fmt.Println("π = (4/1) - (4/3) + (4/5) - (4/7) + (4/9) - (4/11) + (4/13) - (4/15) ...")
    fmt.Println("Three hundred million iterations are being executed ... working ...\n")
    start := time.Now()
    iterFloat64 = 0
    var nextOdd float64
        nextOdd = 1
    four = 4
    var tally float64 
        tally = (four/nextOdd)
        iterInt64 = 0
    for iterInt64 < 300000000 {
        iterInt64++
        iterFloat64++
        nextOdd = nextOdd + 2
        tally = tally-(tally/nextOdd)
        tally = tally+(tally/nextOdd) // pi (tally) is set equl to the sum of a subtraction and an addition, alternatively 

            if iterInt64 == 10000000 { fmt.Println("... 10,000,000 of three hundred million completed. still working, but ...")
                fmt.Println("\n   #1 234567#")
                fmt.Println("   ",tally, "was calculated by the Gregory-Leibniz series") 
                fmt.Println("    3.141592,653589793 is from the web") 
                fmt.Println("   #1 234567 890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Print("  10,000,000 iterations in ", elapsed, " yields 7 digits of π\n\n")
            }
// 7 digits per above
// 8, next two ifs give eight digits
            if iterInt64 == 50000000 { fmt.Println("... 50,000,000 of three hundred million completed. still working, but ...")
                fmt.Println("\n   #1 2345678#")
                fmt.Println("   ",tally, "was calculated by the Gregory-Leibniz series") 
                fmt.Println("    3.1415926,53589793 is from the web") 
                fmt.Println("   #1 2345678 90123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Print("  50,000,000 iterations in ", elapsed, " yields 8 digits of π\n")
                fmt.Println(" ")
            }
            if iterInt64 == 100000000 { fmt.Println("... 100,000,000 of three hundred million completed. still working, but ...")
                fmt.Println("\n   #1 2345678#")
                fmt.Println("   ",tally, "was calculated by the Gregory-Leibniz series") 
                fmt.Println("    3.1415926,53589793 is from the web") 
                fmt.Println("   #1 2345678 90123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Print("  100,000,000 iterations in ", elapsed, " yields 8 digits of π\n\n")
            }
// 9 digits are found below 
            if iterInt64 == 200000000 { fmt.Println("... 200,000,000 of three hundred million completed. still working, but ...")
                fmt.Println("\n   #1 23456789#")
                fmt.Println("   ",tally, "was calculated by the Gregory-Leibniz series") 
                fmt.Println("    3.14159265,3589793 is from the web") 
                fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Print("  200,000,000 iterations in ", elapsed, " yields 9 digits of π\n")
                }
            if iterInt64 == 300000000 {  // last one, still 9 digits
                fmt.Println("\n   #1 23456789#")
                fmt.Print ("    ", tally, " was calculated by the Gregory-Leibniz series")
                fmt.Println("\n    3.141592653589793 is from the web") 
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Print ("  300 million iterations still yields 9 digits, ") // no Println here
                fmt.Print ("in ", elapsed, "\n\n") 
                fmt.Println(" per option #", num, "  --  the Gregory-Leibniz series, circa 1676\n")

                LinesPerIter = 11  // an estimate of the number of lines per iteration 
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("       %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")
//
                fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                    check(err1)                                // ... gets a file handle to dataFromConstants.txt
                    defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                Hostname, _ := os.Hostname()
                _ , err0 := fmt.Fprintf(fileHandle, "\n  -- Gregory-Leibniz -- selection #%d on %s \n", num, Hostname)
                    check(err0)
                current_time := time.Now()
                _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                check(err6)
                _ , err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second  \n", LinesPerSecond) 
                    check(err2)
                _ , err4 := fmt.Fprintf(fileHandle, "%e was Iterations/Seconds  \n", iterFloat64/elapsed.Seconds())
                    check(err4)
                _ , err5 := fmt.Fprintf(fileHandle, "%e was total Iterations  \n", iterFloat64)
                    check(err5)
                TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
                    check(err7)
                fmt.Println("Select 12 at menu to display prior results")
            }
    } 
}

func GottfriedWilhelmLeibniz(num int){
    fmt.Println("\n\nYou selected #", num, " Gottfried Wilhelm Leibniz formula  :  π = 4 * ( 1 - 1/3 + 1/5 - 1/7 + 1/9 ... ")
    fmt.Println("   Infinitesimal calculus was developed independently in the late 17th century by Isaac Newton")
    fmt.Println("    ... and Gottfried Wilhelm Leibniz")
    fmt.Println("   4 Billion iterations will be executed ... ")
    fmt.Println("")
    fmt.Println(" ... working ...\n")
        start := time.Now()
        iterFloat64 = 0
        var denom float64
            denom = 3
        var sum float64
            sum = 1-(1/denom)
        iterInt64 = 1
    for iterInt64 < 4000000000 {
        iterFloat64++
        iterInt64++
        denom = denom + 2
            if iterInt64 % 2 == 0 {
                sum = sum + 1/denom 
            } else { 
                sum = sum - 1/denom 
            }
            π = 4 * sum 
                if iterInt64 == 100000000 {fmt.Println("... 100,000,000 completed iterations ...")
                    fmt.Println("\n   #1 2345678#")
                    fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                    fmt.Println("    3.1415926,53589793 is from the web") 
                    fmt.Println("   #1 2345678 90123456# :: counting the first 16 actual digits of π")
                    t := time.Now()
                    elapsed := t.Sub(start)
                    fmt.Print("  100,000,000 iterations in ", elapsed, " yields 8 digits of π\n\n")
                }
                if iterInt64 == 200000000 {fmt.Println("... 200,000,000 gets another digit ...")
                    fmt.Println("\n   #1 23456789#")
                    fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                    fmt.Println("    3.14159265,3589793 is from the web") 
                    fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                    t := time.Now()
                    elapsed := t.Sub(start)
                    fmt.Print("  200,000,000 iterations in ", elapsed, " yields 9 digits of π\n\n")
                }
                if iterInt64 == 400000000 {fmt.Println("... 400,000,000 iterations completed, still at nine ...")
                    fmt.Println("\n   #1 23456789#")
                    fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                    fmt.Println("    3.14159265,3589793 is from the web") 
                    fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                    t := time.Now()
                    elapsed := t.Sub(start)
                    fmt.Print("  400,000,000 iterations in ", elapsed, " yields 9 digits of π\n\n")
                }
                if iterInt64 == 600000000 {fmt.Println("... 600,000,000 iterations, still at nine ...")
                    fmt.Println("\n   #1 23456789#")
                    fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                    fmt.Println("    3.14159265,3589793 is from the web") 
                    fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                    t := time.Now()
                    elapsed := t.Sub(start)
                    fmt.Print("  600,000,000 iterations in ", elapsed, " yields 9 digits of π\n\n")
                }
                if iterInt64 == 1000000000 {fmt.Println("... 1 Billion iterations completed, still nine ...")
                    fmt.Println("\n   #1 23456789#")
                    fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                    fmt.Println("    3.14159265,3589793 is from the web") 
                    fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                    t := time.Now()
                    elapsed := t.Sub(start)
                    fmt.Print("  1,000,000,000 iterations in ", elapsed, " yields 9 digits of π\n\n")
                }
                if iterInt64 == 2000000000 {fmt.Println("... 2 Billion, and still just nine ...")
                    fmt.Println("\n   #1 23456789#")
                    fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                    fmt.Println("    3.14159265,3589793 is from the web") 
                    fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                    t := time.Now()
                    elapsed := t.Sub(start)
                    fmt.Print("  2,000,000,000 iterations in ", elapsed, " yields 9 digits of π\n\n")
                }
                if iterInt64 == 4000000000 {  // last one
                    fmt.Println("\n... 4 Billion, gets us ten digits  ...")
                    fmt.Println("\n   #1 234567890#")
                    fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                    fmt.Println("    3.141592653,589793 is from the web") 
                    fmt.Println("   #1 234567890 123456# :: counting the first 16 actual digits of π")
                    t := time.Now()
                    elapsed := t.Sub(start)
                    fmt.Print("  4,000,000,000 iterations in ", elapsed, " yields 10 digits of π\n\n")
                    fmt.Println(" per option #", num, "  --  the Gottfried Wilhelm Leibniz formula\n")

                    LinesPerIter = 14
                    fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                    LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                    fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                    fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")
                    fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                        check(err1)                                // ... gets a file handle to dataFromConstants.txt
                        defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.

                    Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  -- Gottfried Wilhelm Leibniz -- selection #%d on %s \n", num, Hostname)  
                        check(err0)
                        current_time := time.Now()
                    _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                        check(err6)
                    _ , err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second  \n", LinesPerSecond) 
                        check(err2)
                    _ , err4 := fmt.Fprintf(fileHandle, "%e was Iterations/Seconds  \n", iterFloat64/elapsed.Seconds())
                        check(err4)
                    _ , err5 := fmt.Fprintf(fileHandle, "%e was total Iterations  \n", iterFloat64)
                        check(err5)
                    TotalRun := elapsed.String() // cast time duration to a String type for Fprintf "formatted print"
                    _ , err7 := fmt.Fprintf(fileHandle, "Total runTime was %s \n", TotalRun)  // add total runtime of this calculation 
                        check(err7)
                    fmt.Println("Select 12 at menu to display prior results")
                }
/* we skip the rest for now 
                if iterInt64 == 6000000000 {fmt.Println("... 6 Billion completed.  ...")
                fmt.Println("   #1 234567890#")
                fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                fmt.Println("    3.141592653,589793 is from the web") 
                fmt.Println("   #1 234567890 123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Print("  6,000,000,000 iterations in ", elapsed, " still yields 10 digits of π\n")
                }
                if iterInt64 == 8000000000 {fmt.Println("... 8 Billion completed. still ten ...")
                fmt.Println("   #1 234567890#")
                fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                fmt.Println("    3.141592653,589793 is from the web") 
                fmt.Println("   #1 234567890 123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Print("  8,000,000,000 iterations in ", elapsed, " still yields 10 digits of π\n")
                }
            if iterInt64 == 9000000000 {
                fmt.Println("   #1 234567890#")
                fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                fmt.Println("    3.141592653,589793 is from the web") 
                fmt.Println("   #1 234567890 123456# :: counting the first 16 actual digits of π")
                // fmt.Print("   ", iter) 
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Print("\n... 9B iterations in ", elapsed, " , but to get 10 digits we only needed 4B iterations\n\n") 
                fmt.Println(" per option #", num, "  --  the Gottfried Wilhelm Leibniz formula")
            }  
*/
    } // end of for
}

func Archimedes(num int){
    fmt.Println("\nYou selected #", num, "  --  An improved version of Archimedes' method")
    fmt.Println("  -- enter \"6' at the main menu for the derivation and proof of the Pythagorean -- ")
                // the above escape does not seem to work as advertised ??
            start := time.Now()
            iterFloat64 = 0
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
        iterInt64 = 0
    for iterInt64 < 27 {
        iterInt64++
        iterFloat64++
        n = n*2
        s1 = s2 
        s1_2 = s1/2 
        a = math.Sqrt(r-(math.Pow(s1_2,2))) 
        b = 1 - a 
        s2 = math.Sqrt(math.Pow(b,2) + (math.Pow(s1_2,2))) 
        p = n * s1 
        p_d = p/2 

        if iterInt64 == 2 {
                fmt.Println("\n   #1 2# :: counting the first 2 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 2# :: counting the first 3 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 2 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", n)
        }    
        if iterInt64 == 3 {
                fmt.Println("\n   #1 23# :: counting the first 3 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 23# :: counting the first 3 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 2 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", n)
        }    
        if iterInt64 == 7 {
                fmt.Println("\n   #1 23456# :: counting the first 6 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 23456# :: counting the first 6 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 5 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", n)
        }    
        if iterInt64 == 12 {
                fmt.Println("\n   #1 23456789# :: counting the first 9 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 23456789# :: counting the first 9 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 8 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", n)
        }    
        if iterInt64 == 17 {
                fmt.Println("\n   #1 23456789012# :: counting the first 12 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 23456789012# :: counting the first 12 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 12 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", n)
        }    
        if iterInt64 == 23 {
                fmt.Println("\n   #1 23456789012345# :: counting the first 15 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 23456789012345# :: counting the first 15 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 15 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", n)
        }
        if iterInt64 == 25 {  // last one 
                fmt.Println("\n   #1 234567890123456# :: counting the first 16 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes' -- 16 -- digits") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 234567890123456789# :: counting the first 19 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 16 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", n)
                fmt.Printf("%.0f as parsed against ...\n", n)
                fmt.Println("100000000 which is one-hundred-million, for comparison to the above line")
                fmt.Println("... Which is to say a 201,326,592 sided polygon, Mr. A. would have wept!\n")
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
            fmt.Println(arch) // prints the above rune
                LinesPerIter = 18  // an estimate
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
                 fmt.Println("   -- After hitting Return for menu redisplay, enter '6' for the derivation and proof of the Pythagorean\n")

                    fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                        check(err1)                                // ... gets a file handle to dataFromConstants.txt
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
                fmt.Println("Select 12 at menu to display prior results")
        }  // end of last if
    }
}

func DisplayPythagorean(num int){
    fmt.Print("\n\n\n -- You entered '", num, "' to review the derivation of the Pythagorean, which was needed in method #5. We will\n")
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
    fmt.Println("Select 12 at menu to display prior results")
}

func JohnWallis(num int){
    fmt.Println("\n   You selected #", num, " A Go language exercize which can be used to test the speed of your hardware.")
    fmt.Println("   We will calculate π to a maximum of ten digits of accuracy using an infinite series by John Wallis circa 1655")
    fmt.Println("   Up to 40 Billion iterations of the following formula will be executed ")
    fmt.Println("   π = 2 * ((2/1)*(2/3)) * ((4/3)*(4/5)) * ((6/5)*(6/7)) ...")
    start := time.Now() 
    iterFloat64 = 0
    var numerators float64
        numerators = 2
    var firstDenom float64
        firstDenom = 1
    var secondDenom float64
        secondDenom = 3
    var cumulativeProduct float64
        cumulativeProduct = (numerators/firstDenom) * (numerators/secondDenom)
          iterInt64 = 0
      for iterInt64 < 40000000000 {
        iterInt64++
        iterFloat64++
        numerators = numerators + 2
        firstDenom = firstDenom + 2
        secondDenom = secondDenom + 2
        cumulativeProduct = cumulativeProduct * (numerators/firstDenom) * (numerators/secondDenom)
        π = cumulativeProduct * 2  
        if iterInt64 == 10 { fmt.Println(" ") }  //  single-line "if" just to display a blank line; for amusement 
/*            if iterInt64 == 100 {
                fmt.Println("   #1 2# ")
                fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                fmt.Println("    3.1,41592653589793 is, again, the value of π from the web") 
                fmt.Println("   #1 2# 34567890123456 :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println(iterInt64, " iterations were completed in ", elapsed, " yielding 2 digits of π\n") 
            }
            if iterInt64 == 500 {
                fmt.Println("   #1 23# ")
                fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                fmt.Println("    3.14,1592653589793 is, again, the value of π from the web") 
                fmt.Println("   #1 23# 4567890123456 :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println(iterInt64, " iterations were completed in ", elapsed, " yielding 3 digits of π\n") 
            }
*/
            if iterInt64 == 2000 {
                fmt.Println("   #1 234# ")
                fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                fmt.Println("    3.14,1592653589793 is, again, the value of π from the web") 
                fmt.Println("   #1 23 4567890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println(iterInt64, " iterations were completed in ", elapsed, " yielding 4 digits of π\n") 
            }
            if iterInt64 == 10000 {
                fmt.Println("   #1 2345# ")
                fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                fmt.Println("    3.1415,92653589793 is, again, the value of π from the web") 
                fmt.Println("   #1 2345 67890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("10,000 iterations were completed in ", elapsed, " yielding 5 digits of π\n") 
            }
            if iterInt64 == 50000 { // 50,000 
                fmt.Println("   #1 2345# ")
                fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                fmt.Println("    3.1415,92653589793 is, again, the value of π from the web") 
                fmt.Println("   #1 2345 67890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("50,000 iterations were completed in ", elapsed, " yielding 5 digits of π\n") 
            }
            if iterInt64 == 500000 {  // 500,000 done
                fmt.Println("   #1 23456# ")
                fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                fmt.Println("    3.14159,2653589793 is, again, the value of π from the web") 
                fmt.Println("   #1 23456 7890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("500,000 iterations were completed in ", elapsed, " yielding 6 digits of π\n") 
            }
            if iterInt64 == 2000000 {  // 2M done 
                fmt.Println("   #1 234567# ")
                fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                fmt.Println("    3.141592,653589793 is, again, the value of π from the web") 
                fmt.Println("   #1 234567 890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("2,000,000 iterations were completed in ", elapsed, " yielding 7 digits of π\n") 
            }
            if iterInt64 == 40000000 {  // 40M done
                fmt.Println("   #1 2345678# ")
                fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                fmt.Println("    3.1415926,53589793 is, again, the value of π from the web") 
                fmt.Println("   #1 2345678 90123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("40,000,000 iterations were completed in ", elapsed, " yielding 8 digits of π\n\n") 
                fmt.Println("  .. working .. on another factor-of-ten iterations\n")
            }
            if iterInt64 == 400000000 {  // 400M done
                fmt.Println("   #1 23456789# ")
                fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                fmt.Println("    3.14159265,3589793 is, again, the value of π from the web") 
                fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("400,000,000 iterations were completed in ", elapsed, " yielding 9 digits of π\n\n") 

                LinesPerIter = 36 // an estimate 
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")
//
                fmt.Println("  ... will be working on doing Billions more iterations ...\n\n")
            }
//
            if iterInt64 == 600000000 {  // 600M done
                fmt.Println("  600M done, still working on another Two-Hundred-Thousand iterations ... working ...\n")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println(elapsed, "\n")
                fmt.Println("Calculating 10th digit (40B iters) which takes a few min \n")
                fmt.Println("- Ctrl-C to End/Exit without saving results\n")
                LinesPerIter = 36 // an estimate 
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")
                fmt.Println(" ... still working ...")
            }

            if iterInt64 == 800000000 {  // 800M done
                fmt.Println("  800M done, still working on yet another Two Hundred Thousand iterations ... working ...\n")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println(elapsed, " \n")
            }
        if iterInt64 == 1000000000 {  // 1B done
            fmt.Println("   #1 23456789# ")
            fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
            fmt.Println("    3.14159265,3589793 is the value of π from the web") 
            fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
            t := time.Now()
            elapsed := t.Sub(start)
            fmt.Println("\nOne Billion iterations were completed in ", elapsed, " still only yielding π to 9 digits\n") 
            fmt.Println(" per option #", num, "  --  an infinite series by John Wallis circa 1655\n") // ----------------------

                LinesPerIter = 36 // an estimate 
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")
//
            fmt.Println("  ... still working ... on Billions of iterations, 39 to go ...\n\n")
        }

//  "Optional" Forty Billion loops just to get one additional digit, the tenth digit of pi 

        if iterInt64 == 2000000000 {  // 2B completed 
            fmt.Println("  2B done, still working ... on another Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 3000000000 {  // 3B completed
            fmt.Println("  3B done, still working ... on another Billion iterations ... working ... Ctrl-C to End/Exit without saving stats")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 4000000000 {  // 4B completed 
            fmt.Println("  4B done, still working ... on another Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 5000000000 {  // 5B completed 
            fmt.Println("  5B done, still working ... on another Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 6000000000 {  // 6B completed 
            fmt.Println("  6B done, still working ... on another Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 7000000000 {  // 7B completed 
            fmt.Println("  7B done, still working ... on another Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 8000000000 {  // 8B completed 
            fmt.Println("  8B done, still working ... on another Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 9000000000 {  // 9B completed 
            fmt.Println("  9B done, still working ... on another five Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 14000000000 {  // 14B completed
            fmt.Println("  14B done, still working ... on another five Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 19000000000 {  // 19B completed 
            fmt.Println("  19B done, still working ... on another five Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 24000000000 { // 24B completed
            fmt.Println("  24B done, still working ... on another five Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 29000000000 {  // 29B completed 
            fmt.Println("  29B done, still working ... on another five Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 34000000000 {  // 34B completed
            fmt.Println("  34B done, still working ... just another six Billion iterations to go! ... ")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
 //  "Optional" Forty Billion loops to get another digit
        if iterInt64 == 40000000000 {  // 40B completed 
            fmt.Println("   #1 234567890# ")
            fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
            fmt.Println("    3.141592653,589793 is the value of π from the web") 
            fmt.Println("   #1 234567890 123456# :: counting the first 16 actual digits of π")
            t := time.Now()
            elapsed := t.Sub(start)
            fmt.Println("Forty Billion iterations were completed in ", elapsed, " yielding π to 10 digits\n") 
            fmt.Println(" per option #", num, "  --  an infinite series by John Wallis circa 1655\n") // ----------------------
                LinesPerIter = 36 // an estimate 
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")
                fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                    check(err1)                                // ... gets a file handle to dataFromConstants.txt
                    defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                Hostname, _ := os.Hostname()
                _ , err0 := fmt.Fprintf(fileHandle, "\n  -- John Wallis -- selection #%d on %s \n", num, Hostname)
                    check(err0)
                current_time := time.Now()
                _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                check(err6)
                _ , err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second  \n", LinesPerSecond) 
                    check(err2)
                _ , err4 := fmt.Fprintf(fileHandle, "%.02f was Iterations/Seconds  \n", iterFloat64/elapsed.Seconds())
                    check(err4)
                _ , err5 := fmt.Fprintf(fileHandle, "%e was total Iterations  \n", iterFloat64)
                    check(err5)
                TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun)  // add total runtime of this calculation 
                    check(err7)
                fmt.Println("Select 12 at menu to display prior results")
        }
    }
}

func EulersNumber(num int){
    var n float64 
    var sum float64
    var Eulers float64
    fmt.Println("\n\n\nEuler's Number \u2107 or \u2147 the natural logarithmic base")
    fmt.Println("\u2147 = (1+1/n)^n")
    fmt.Println("... the limit of an increasing value for n\n\n") 
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
// And now, here comes a fun rune to print a multi-line "string"
// ... define a rune with the ` :: back-quote character located on the ~ tilda key
Ricks_rune_Paragraph := `  
Bernoulli noticed that this sequence approaches a limit (the force of interest) with larger n and, thus, smaller 
compounding intervals. Compounding weekly (n = 52) yields $2.692596..., while compounding daily (n = 365) yields
$2.714567... (approximately two cents more). The limit as n grows large is the number that came to be known as e.
That is, with continuous compounding, the account value will reach $2.718281828 
`
fmt.Println(Ricks_rune_Paragraph)
}


func ErdosBorwein(num int){
       rune := `The Erdős–Borwein constant is the sum of the reciprocals of the Mersenne numbers. 
It is named after Paul Erdős and Peter Borwein. 

Paul Erdős was a Hungarian mathematician. He was one of the most prolific mathematicians 
and producers of mathematical conjectures of the 20th century. Erdős pursued and proposed 
problems in discrete mathematics, graph theory, number theory, mathematical analysis, 
approximation theory, set theory, and probability theory.[4] Much of his work centered 
around discrete mathematics, cracking many previously unsolved problems in the field. 
He championed and contributed to Ramsey theory, which studies the conditions in which 
order necessarily appears. Overall, his work leaned towards solving previously open 
problems, rather than developing or exploring new areas of mathematics.

Erdős published around 1,500 mathematical papers during his lifetime, a figure that
remains unsurpassed.[5] He firmly believed mathematics to be a social activity, 
living an itinerant lifestyle with the sole purpose of writing mathematical papers 
with other mathematicians. He was known both for his social practice of mathematics, 
working with more than 500 collaborators, and for his eccentric lifestyle; Time magazine 
called him "The Oddball's Oddball".[6] He devoted his waking hours to mathematics, even 
into his later years—indeed, his death came only hours after he solved a geometry 
problem at a conference in Warsaw. Erdős's prolific output with co-authors prompted 
the creation of the Erdős number, the number of steps in the shortest path between a 
mathematician and Erdős in terms of co-authorships. 


Peter Benjamin Borwein (born St. Andrews, Scotland, May 10, 1953 – 23 August 2020) 
was a Canadian mathematician and a professor at Simon Fraser University. He is known 
as a co-author of the paper which presented the Bailey–Borwein–Plouffe algorithm 
(discovered by Simon Plouffe) for computing π. 

Borwein was born into a Jewish family. He became interested in number theory and classical 
analysis during his second year of university. He had not previously been interested in 
math, although his father was the head of the University of Western Ontario's mathematics 
department and his mother is associate dean of medicine there. Borwein and his two siblings 
majored in mathematics. 

After completing a Bachelor of Science in Honours Math at the University of Western 
Ontario in 1974, he went on to complete an MSc and Ph.D. at the University of British 
Columbia. He joined the Department of Mathematics at Dalhousie University. While he 
was there, he, his brother Jonathan Borwein and David H. Bailey of NASA wrote the 1989 
paper[1] that outlined and popularized a proof for computing one billion digits of π. 
The authors won the 1993 Chauvenet Prize and Merten M. Hasse Prize for this paper.

In 1993, he moved to Simon Fraser University, joining his brother Jonathan in establishing 
the Centre for Experimental and Constructive Mathematics (CECM) where he developed the 
Inverse Symbolic Calculator. `

fmt.Println("") // so as not to pad the following rune with an unwanted space 
    fmt.Println(rune, "\n")
    fmt.Println("We calculate E as E = the sum of 1/((2^n)-1) as n grows from 1 to 'infinity'")

    var Erdos_Borwein float64
        Erdos_Borwein = 1
    var iter float64 
        iter = 1
    for iter < 100 {
        iter++ // iter will therefore begin at 2
        Erdos_Borwein = Erdos_Borwein + (1/( (math.Pow(2,iter)) - 1) )
        if iter == 10 || iter == 20 || iter == 30 || iter == 40 || iter == 50 || iter == 60 || iter == 70 || iter == 100 || iter == 101 { fmt.Println(Erdos_Borwein)}
                                                                                                    // 100 and 101 prove that we ended on 100 as the final exponent
                                                                                                    // ... we only get 8 results, not nine
    }
    fmt.Println("for 10, 20, 30, 40, 50, 60, 70, and 100 iterations respectively\n")
    fmt.Println("Our calculated Erdos-Borwein constant is ")
    fmt.Println(Erdos_Borwein, "after", iter, "iterations, i.e., with a final exponent of", iter)
    fmt.Println("1.606695152415291763 is what we get from the web\n")
}

func Explain_spigot(){
    fmt.Println(spigot_rune_per_chatGPT) // https://chat.openai.com/chat)
}
var spigot_rune_per_chatGPT =

`According to chatGPT; The spigot algorithm for generating the digits of pi is a mathematical algorithm that can be used to calculate the digits of pi to any desired level of precision. It is called a "spigot" algorithm because it generates the digits of pi one at a time, like water flowing from a spigot.

The spigot algorithm works by starting with a known series expansion for pi and then using a series of mathematical operations to extract the individual digits of pi one at a time. The specific series expansion that is used is the Bailey-Borwein-Plouffe (BBP) formula, which expresses pi as an infinite series of hexadecimal (base 16) digits.

To generate the digits of pi using the spigot algorithm, the BBP formula is evaluated term by term, and the hexadecimal digits of each term are converted to decimal (base 10) and added to an accumulator. The decimal digits of the accumulator are then extracted one at a time, and the remaining fractional part of the accumulator is used to compute the next term in the series. This process is repeated until the desired number of digits has been generated.

The spigot algorithm has the advantage of being relatively simple and easy to implement, and it is capable of generating the digits of pi to very high precision. However, it can be slow for large numbers of digits, and there are other algorithms that are faster and more efficient for generating large numbers of digits. 

The Bailey-Borwein-Plouffe (BBP) formula, which is used in the spigot algorithm for generating the digits of pi, expresses pi as an infinite series of hexadecimal (base 16) digits. This means that each term in the series is represented as a fraction whose numerator and denominator are both powers of 16.

The reason for using base 16 in the BBP formula is that it allows the series to be expressed in a more compact and efficient form. In general, using a larger base for a number representation allows for more digits to be expressed using fewer bits of information. For example, a number represented in base 2 (binary) can have at most 2 digits (0 and 1), while a number represented in base 10 (decimal) can have at most 10 digits (0 through 9). A number represented in base 16 (hexadecimal) can have at most 16 digits (0 through 9 and A through F).

By expressing the series in base 16, the BBP formula is able to represent the digits of pi more efficiently, which makes it easier to compute and manipulate. It also has the advantage of being relatively simple to convert between base 16 and other common bases, such as base 10 (decimal) and base 2 (binary).

The BBP formula is just one of many methods that can be used to generate the digits of pi, and other methods may use different bases or representation schemes. The choice of base or representation scheme is often determined by the specific properties of the algorithm or the requirements of the application. 

It expresses pi as an infinite series of hexadecimal (base 16) digits and can be used to calculate the nth digit of pi directly, without having to compute all of the previous digits.

The general form of the BBP formula is as follows:

pi = SUM[k=0 to infinity] (16^(-k)) * (4/(8k+1) - 2/(8k+4) - 1/(8k+5) - 1/(8k+6))

This formula expresses pi as the sum of an infinite series of fractions, where each fraction is the result of dividing a series of constants (4, 2, 1, and 1) by a series of powers of 8. The exponent (-k) in the 16^(-k) term determines the position of the digit being calculated in the series, with larger values of k corresponding to higher digits in the series.

To calculate a specific digit of pi using the BBP formula, the sum is evaluated up to the desired term, and the resulting hexadecimal fraction is converted to decimal and rounded to the nearest digit. For example, to calculate the first few digits of pi, the sum would be evaluated up to the k=0 term, and the resulting hexadecimal fraction would be converted to decimal and rounded to the nearest digit. To calculate the next digit, the sum would be evaluated up to the k=1 term, and so on.
`


func TheSpigot(){
    var numberOfDigitsToCalc int
    fmt.Println("How much pi can you handle?")
    fmt.Println("How many digits of pi do you really want?")  
            fmt.Scanf("%d", &numberOfDigitsToCalc)
    fmt.Println("The first line below was calculated via a Spigot Algorithm")
    fmt.Println("The second line is the value of pi from the web")
    fmt.Println("\n", Spigot(numberOfDigitsToCalc))

fmt.Println(" 31415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679")
fmt.Println(" 12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901")
fmt.Println("        ten    twenty        30        40        50        60        70        80        90       100 ")
fmt.Print("This last trick made possible by a bit of code I mooched off of GitHub ...\n")
fmt.Println("... to view the code with attribution Enter '20 or 39' at the second menu\n")
fmt.Println("To view an explaination of how it works enter 99 at either menu\n")
}

func showTheSpigotMagic() {
fmt.Println("\n\nThis 'spigot' trick was adapted from a bit of code I mooched off of GitHub:")
fmt.Println("... and I have yet to discover/figure-out how this particular trick works. \n")
fmt.Print(spigotRune, "[the rune of the code would go here]\n")
fmt.Println("\nThis 'spigot' trick was adapted from a bit of code I mooched off of GitHub:")
fmt.Println("... and I have yet to discover/figure-out how this particular trick works. \n")
fmt.Println("To view an explaination of how it works enter 99 at either menu\n")
}

var spigotRune = `
// Generously shared by Ilya Sokolov i-redbyte, https://github.com/i-redbyte
// https://github.com/TheAlgorithms/Go/blob/master/math/pi/spigotpi.go

func Spigot(n int) string {
    pi := ""
    boxes := n * 10 / 3
    remainders := make([]int, boxes)
    for i := 0; i < boxes; i++ {
        remainders[i] = 2
    }
    digitsHeld := 0
    for i := 0; i < n; i++ {
        carriedOver := 0
        sum := 0
        for j := boxes - 1; j >= 0; j-- {
            remainders[j] *= 10
            sum = remainders[j] + carriedOver
            quotient := sum / (j*2 + 1)
            remainders[j] = sum % (j*2 + 1)
            carriedOver = quotient * j
        }
        remainders[0] = sum % 10
        q := sum / 10
        switch q {
        case 9:
            digitsHeld++
        case 10:
            q = 0
            for k := 1; k <= digitsHeld; k++ {
                replaced, _ := strconv.Atoi(pi[i-k : i-k+1])
                if replaced == 9 {
                    replaced = 0
                } else {
                    replaced++
                }
                pi = delChar(pi, i-k)
                pi = pi[:i-k] + strconv.Itoa(replaced) + pi[i-k:]
            }
            digitsHeld = 1
        default:
            digitsHeld = 1
        }
        pi += strconv.Itoa(q)
    }
    return pi
}

func delChar(s string, index int) string {
    tmp := []rune(s)
    return string(append(tmp[0:index], tmp[index+1:]...))
}

// Generously shared by Ilya Sokolov i-redbyte, https://github.com/i-redbyte
// https://github.com/TheAlgorithms/Go/blob/master/math/pi/spigotpi.go

// to which I added my own hooks: 
// ... 
func TheSpigot(){
    var numberOfDigitsToCalc int
    fmt.Println("How much pi can you handle?")
    fmt.Println("How many digits of pi do you really want?")  
            fmt.Scanf("%d", &numberOfDigitsToCalc)
    fmt.Println("The first line below was calculated via a Spigot Algorithm")
    fmt.Println("The second line is the value of pi from the web")
    fmt.Println("\n", Spigot(numberOfDigitsToCalc))

fmt.Println(" 31415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679")
fmt.Println(" 12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901")
fmt.Println("        ten    twenty        30        40        50        60        70        80        90       100 ")
fmt.Print("This last trick made possible by a bit of code I mooched off of GitHub ...\n")
fmt.Println("... to view the code with attribution Enter '20 or 39' at the main menu")
}

func showTheMagic() {
fmt.Print("That last trick made possible by a bit of code I mooched off of GitHub:\n", spigotRune, "\n\n")
}

var spigotRune = [runMark]
func Spigot(n int) string { ... ` 
// end of rune 

func Spigot(n int) string {
    pi := ""
    boxes := n * 10 / 3
    remainders := make([]int, boxes)
    for i := 0; i < boxes; i++ {
        remainders[i] = 2
    }
    digitsHeld := 0
    for i := 0; i < n; i++ {
        carriedOver := 0
        sum := 0
        for j := boxes - 1; j >= 0; j-- {
            remainders[j] *= 10
            sum = remainders[j] + carriedOver
            quotient := sum / (j*2 + 1)
            remainders[j] = sum % (j*2 + 1)
            carriedOver = quotient * j
        }
        remainders[0] = sum % 10
        q := sum / 10
        switch q {
        case 9:
            digitsHeld++
        case 10:
            q = 0
            for k := 1; k <= digitsHeld; k++ {
                replaced, _ := strconv.Atoi(pi[i-k : i-k+1])
                if replaced == 9 {
                    replaced = 0
                } else {
                    replaced++
                }
                pi = delChar(pi, i-k)
                pi = pi[:i-k] + strconv.Itoa(replaced) + pi[i-k:]
            }
            digitsHeld = 1
        default:
            digitsHeld = 1
        }
        pi += strconv.Itoa(q)
    }
    return pi
}

func delChar(s string, index int) string {
    tmp := []rune(s)
    return string(append(tmp[0:index], tmp[index+1:]...))
}
// hey gtp chatbot, see anything wrong with that code up thar? 
/*
to which it responds with its usual bullshit:

"There are a few issues with the above code:

    The boxes variable is being calculated using integer division, which will truncate the result. 
    This can lead to an incorrect number of boxes being used in the calculations.

    The remainders slice is being initialized with all elements equal to 2. This does not match the 
    algorithm described in the code, and will likely result in incorrect digits being generated.

    The digitsHeld variable is not being initialized to 0 at the start of the outer loop. This can 
    cause unexpected behavior or errors.

    The switch statement is not handling all possible cases. If q is not 9 or 10, the code will not e
    xecute the default case and digitsHeld will not be set to 1. This can cause unexpected behavior or errors.

    The delChar function is using index as the index of the character to delete in the string, 
    but it is not being adjusted for the length of the string. This can cause the function to delete 
    the wrong character or index out of bounds.

    The function does not return an accurate value for pi. The spigot algorithm is a method for generating 
    the digits of pi, but it does not converge to the true value of pi.""
*/

func showMagicBehindsquareRootOf3() {
    var squareRootOf3rune = `

var diff1 int
var diff2 int
var diff3 int
var diff4 float64

func buildTableOfOnlyPerfectSquares() { 
    root := 100
            iter := 0
        for iter < 1000000 {
            iter++
            root++ 
            PerfectSquare := root*root
            sliceOfPerfectSquares_and_their_roots = append(sliceOfPerfectSquares_and_their_roots, PerfectSquare ) 
            sliceOfPerfectSquares_and_their_roots = append(sliceOfPerfectSquares_and_their_roots, root) 
        }
}

func readTableOfPerfectSquares(index2 int) {  
    smallerPerfectSquareOnce := sliceOfPerfectSquares_and_their_roots[index2]  // save it, do this just-once per func call 
        iter := 0
    for iter < 96000 {
        iter++
        index2 = index2 + 2 // index 
        largerPerfectSquare := sliceOfPerfectSquares_and_their_roots[index2]        // get next perfect square 
        if largerPerfectSquare == smallerPerfectSquareOnce*3 { // ... in the case that it is a miricle ...
            fmt.Println("It is a miricle, everyone was wrong! ") // praise god
            break  // exit the for loop ????????????????
        }
        if largerPerfectSquare >= smallerPerfectSquareOnce*3 {     // if largerPerfectSquare is a candidate based on it being just-a-bit larger than 3*smlr_PS deal with that, else loop to the next potential 
            
            ProspectiveHitOnLargeSide = largerPerfectSquare
            p_diff_from_larger_hit := determinDiff(ProspectiveHitOnLargeSide, smallerPerfectSquareOnce*3)  // ProspectiveHitOnLargeSide_Par, smallerPerfectSquareOnce_T3
            sliceOfP_diffs_smlrT3_to_largerOfProspectiveHit = append(sliceOfP_diffs_smlrT3_to_largerOfProspectiveHit, p_diff_from_larger_hit)

            Prospective_smaller_hit = sliceOfPerfectSquares_and_their_roots[index2-2]  // an int
            p_diff_from_smlr_hit := determinDiff(Prospective_smaller_hit, smallerPerfectSquareOnce*3)
            sliceOfP_diffs_smlrT3_to_smlrOfProspectiveHit = append(sliceOfP_diffs_smlrT3_to_smlrOfProspectiveHit, p_diff_from_smlr_hit) // build a slice of ... // may not need this as a slice ????

// and now we cheat a bit, just to see if this is working ... 
            sr_per_sm := (math.Sqrt(float64(Prospective_smaller_hit)) / math.Sqrt(float64(smallerPerfectSquareOnce)))  // a 3:1 ratio 
            if sr_per_sm > 1.7320508066 { 
                if sr_per_sm < 1.7320508074 { 
                    fmt.Println("", sr_per_sm, "from top, is close at", iter, "iterations, hits are:", Prospective_smaller_hit, smallerPerfectSquareOnce*3)
                    sliceOfResults = append(sliceOfResults, sr_per_sm)
                } 
            } 
            if sr_per_sm > 1.73205080752 {
                if sr_per_sm < 1.7320508077 {
                    sliceOfResults = append(sliceOfResults, sr_per_sm)
                    fmt.Println("\nPer the smaller ps, the \u221A3 has been calculated to be ... \n", sr_per_sm, "after", iter, "iterations, compared to:\n", "1.73205080757 <--the actual" ) 
                    fmt.Println("which is pretty good, and via first-principles of geometry alone\n... having found two very large perfect squares where one of them is")
                    fmt.Println("very-nearly exactly three times larger, i.e., the sqrt of", Prospective_smaller_hit, "over the sqrt of", smallerPerfectSquareOnce, "")
                    fmt.Println("\n... A perfect bottom square from our initial list of perfect squares, times 3, is ...")
                    fmt.Println(smallerPerfectSquareOnce*3, " And the total iterations completed was ", TotalIterations)
                    fmt.Println(Prospective_smaller_hit, " was the other perfect square which we found by searching for one that would be nearly-exactly three times larger")
                    diff1 = 0
                    if smallerPerfectSquareOnce*3 < Prospective_smaller_hit {
                            diff1 = Prospective_smaller_hit - smallerPerfectSquareOnce*3
                    } else {
                            diff1 = smallerPerfectSquareOnce*3 - Prospective_smaller_hit
                        }
                    if diff1 <= 3 {
                            fmt.Println("... pretty close, I'd say, with a difference of", diff1, "\n\n")
                    } else {
                            fmt.Println("... not as close as we'd like, since the difference was", diff1, "\n\n")
                        }
                    savedHit = Prospective_smaller_hit
                    smallerPerfectSquareOnceSaved = smallerPerfectSquareOnce
                } 
            }

// we really should instead be deciding final answer based on closeness of fit, percent diff, to exactly 3x smaller perfect square. 
// Not just checking to see if we have gotten this close to the known value of Sqrt of 3

            sr_per_lg := (math.Sqrt(float64(ProspectiveHitOnLargeSide)) / math.Sqrt(float64(smallerPerfectSquareOnce)))
            if sr_per_lg < 1.732050808 {
                if sr_per_lg > 1.73205080 { 
                    fmt.Println("", sr_per_lg, "from bottom, is close at", iter, "iterations, hits are:", ProspectiveHitOnLargeSide, smallerPerfectSquareOnce*3)
                    sliceOfResults = append(sliceOfResults, sr_per_lg)
                } 
            } 
            if sr_per_lg < 1.7320508079 {
                    sliceOfResults = append(sliceOfResults, sr_per_sm)
                if sr_per_lg > 1.73205080 {  // 1.73205080757 (so 1.73205080754 - 1.73205080759 )
                    fmt.Println("\nPer the larger ps, the \u221A3 has been calculated to be ... \n", sr_per_lg, "after", iter, "iterations, compared to:\n", "1.73205080757 <--the actual" )
                    fmt.Println("which is pretty good, and via first-principles of geometry alone\n... having found two very large perfect squares where one of them is")
                    fmt.Println("very-nearly exactly three times larger, i.e., the sqrt of", ProspectiveHitOnLargeSide, "over the sqrt of", smallerPerfectSquareOnce, "")
                    fmt.Println("\n... A perfect bottom square from our initial list of perfect squares, times 3, is ...")
                    fmt.Println(smallerPerfectSquareOnce*3, " And the total iterations completed was ", TotalIterations)
                    fmt.Println(ProspectiveHitOnLargeSide, " was the other perfect square which we found by searching for one that would be nearly-exactly three times larger")
                    diff2 = 0
                    if smallerPerfectSquareOnce*3 < ProspectiveHitOnLargeSide {
                            diff2 = ProspectiveHitOnLargeSide - smallerPerfectSquareOnce*3
                    } else {
                            diff2 = smallerPerfectSquareOnce*3 - ProspectiveHitOnLargeSide
                        }
                    if diff2 <= 3 {
                            fmt.Println("... pretty close, I'd say, with a difference of", diff2, "\n\n")
                    } else {
                            fmt.Println("... not as close as we'd like, since the difference was", diff2, "\n\n")
                        }
                    savedHitB = ProspectiveHitOnLargeSide
                }
            }
                break // this break gets us out of the for loop having found a candidate within the larger if
                // without this break nothing displays, and the proc is eventually killed
        }  // end "break-out" of the loop because "if we have found a candidate" we have already dealt with it so we need to break and get another potential from the for loop
    }       // END OF FOR LOOP after 96,000 index-d loops "iter-s"
    TotalIterations = iter // a global var to also be used elsewhere 
} 

func squareRootOf3(num int) {
    start := time.Now() // saved start time to be compared with end time t
        buildTableOfOnlyPerfectSquares() 
            var index = 0 
            for index < 300000 {
                readTableOfPerfectSquares(index)  // pass-in the index 
                index = index + 2 
            }
        fmt.Println("\n... A perfect bottom square from our initial list of perfect squares, times 3, is ...")
        fmt.Println(smallerPerfectSquareOnceSaved*3, " And the total iterations completed was ", TotalIterations)
        fmt.Println(savedHit, " was the other perfect square which we found by searching for one that would be nearly-exactly three times larger")

        if smallerPerfectSquareOnceSaved*3 < savedHit {
            diff3 = savedHit - smallerPerfectSquareOnceSaved*3
        } else {
            diff3 = smallerPerfectSquareOnceSaved*3 - savedHit
            }
        
        if diff3 <= 3 {
            fmt.Println("... pretty close, I'd say, with a difference of", diff3, "\n\n")
        } else {
            fmt.Println("... not as close as we'd like, since the difference was", diff3, "\n\n")
            }
            //fmt.Println(sliceOfResults) // show whole slice
            //fmt.Println("len is ", len(sliceOfResults)) // gives the length that we need
            //fmt.Println("cap is ", cap(sliceOfResults)) // gave a larger irrelevant number
        items := 0
        for items < len(sliceOfResults) {
            // compare item from sliceOfResults with 1.73205080757 which is the actual
            sliceItem := sliceOfResults[items]
            if sliceItem < 1.73205080757 {
                diff4 = 1.73205080757 - sliceItem
            } else {
                diff4 = sliceItem - 1.73205080757 
                }
            if diff4 < 4.9e-11 {
                    ts3 := time.Now()
                    elapseds3 := ts3.Sub(start)
                    fmt.Println(TotalIterations, " iterations were completed in ", elapseds3, " yielding 11 digits of the \u221A3 \n") 

                    fmt.Println(smallerPerfectSquareOnceSaved*3, "Was our best perfect square from the initial list")
                    fmt.Println(savedHit, " was the best perfect square which we found that would be nearly-exactly three times larger")
                    fmt.Println("... that pair generated the closest calculated approximation to the actual value of \u221A3 \n")

                    fmt.Println(sliceItem, "difference is ", diff4)
                    fmt.Println(1.73205080757, "<--actual")
            }
            items++
        }
    t := time.Now()
    elapsed := t.Sub(start)
    var LinesPerSecondInt int 
    LinesPerIter := 98 // an estimate 
        LinesPerSecondInt = (int(LinesPerIter) * int(TotalIterations) ) / int(elapsed.Seconds()) // .Seconds() returns a float64
        fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
            check(err1)                                // ... gets a file handle to dataFromConstants.txt
            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
        Hostname, _ := os.Hostname()
        _ , err0 := fmt.Fprintf(fileHandle, "\n  -- sqrt of 3 geometrically -- selection #%d on %s \n", num, Hostname)
            check(err0)
        current_time := time.Now()
        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
        check(err6)
        _ , err2 := fmt.Fprintf(fileHandle, "%d was Lines/Second \n", LinesPerSecondInt) 
            check(err2)
        _ , err4 := fmt.Fprintf(fileHandle, "%.02f was Iterations/Seconds \n", float64(TotalIterations)/float64(elapsed.Seconds()))
            check(err4)
        _ , err5 := fmt.Fprintf(fileHandle, "%d was total Iterations \n", TotalIterations)
            check(err5)
        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
            check(err7)
}
    `
    fmt.Println(squareRootOf3rune)
}

func showMagicBehindBBPF() {
var BBPFrune = `
func BBPF(num int) {
    fmt.Println("\nYou selected #", num, "the Bailey–Borwein–Plouffe formula for π, circa 1995\n")
    fmt.Println("How many digits of pi should we calculate? Enter one integer '4 to 16' ")
                 var numAi float64
    fmt.Scanf("%f", &numAi)
    fmt.Printf("pi calculated to %.0f places: %.15f \n", numAi, calculatePi(numAi)) 
    fmt.Println("and Pi from the web is   : 3.141592653589793")
}
func calculatePi(precision float64) float64 {
  var pi float64
      pi = 0
  var k float64
      k = 0
  for k < (precision-2) {
    pi += math.Pow(16, -k) * (4.0/(8*k+1) - 2.0/(8*k+4) - 1.0/(8*k+5) - 1.0/(8*k+6))
    fmt.Println(pi)
    k++
  }
  return pi
}`
    fmt.Println(BBPFrune)
}

func showMagicBehindWorstMethod() {
    var WorstMethodRune = `
func WorstMethod(num int){
    fmt.Println("\nYou selected #")
    fmt.Println("\n     π = 12 * ( 1/2 - (1/2  * 1/3 * (1/2)exp3) - ...")
    fmt.Println("                                    (1/8   * 1/5 * (1/2)exp5) - ...")
    fmt.Println("                                    (1/16  * 1/7 * (1/2)exp7) - ...")
    fmt.Println("                                    (1/128 * 1/9 * (1/2)exp9) - ... - (\u221A3)/8) )")
        var secondDenom float64
            secondDenom = 3
        var twelve float64
            twelve = 12
        var πViaNewton float64 
        var initialDenomOfFirst float64
            initialDenomOfFirst = 2
        var nextDenomOfFirst float64 
            nextDenomOfFirst = 8
        var savedFirstDenom float64
            savedFirstDenom = 0.0
        var exponent float64 
            exponent = 3
        var SumOfTerms float64 
            SumOfTerms = - ( (1/initialDenomOfFirst) * (1/secondDenom) * math.Pow(0.50,exponent) )  // 2 and 3 and 3
        iterInt64 = 0
    for iterInt64 < 6 {
        iterInt64++ 
            exponent = exponent + 2.0
            secondDenom = secondDenom + 2.0
            SumOfTerms = (SumOfTerms - ( (1/nextDenomOfFirst) * (1/secondDenom) * math.Pow(0.50,exponent) ) )
            fmt.Println(nextDenomOfFirst, secondDenom, exponent)
            savedFirstDenom = nextDenomOfFirst 
            nextDenomOfFirst = savedFirstDenom * initialDenomOfFirst 
            initialDenomOfFirst = savedFirstDenom 
        if iterInt64 == 2 {
            πViaNewton = ( twelve * ( (1/2.0) + SumOfTerms - (math.Sqrt(3)/8.0) )  ) 
            fmt.Println("After", iterInt64, "iterations with exponent:", exponent)
            fmt.Println("  #1 234#")
            fmt.Println("  ", πViaNewton, "The same equation on a pocket calculator gives 3.14171 with only 3 terms e7")
            fmt.Println("   3.141592653589793 is the value of π from the web") 
            fmt.Println("  #1 234567890123456# :: counting the first 16 actual digits of π")
            fmt.Println(" ")
        }
        if iterInt64 == 3 {
            πViaNewton = ( twelve * ( (1/2.0) + SumOfTerms - (math.Sqrt(3)/8.0) )  ) 
            fmt.Println("After", iterInt64, "iterations with exponent:", exponent)
            fmt.Println("  #1 234#")
            fmt.Println("  ", πViaNewton, "The same equation on a pocket calculator gives 3.14171 with only 3 terms e7")
            fmt.Println("   3.141592653589793 is the value of π from the web") 
            fmt.Println("  #1 234567890123456# :: counting the first 16 actual digits of π")
            fmt.Println(" ")
        }
        if iterInt64 == 4 {
            πViaNewton = ( twelve * ( (1/2.0) + SumOfTerms - (math.Sqrt(3)/8.0) )  ) 
            fmt.Println("After", iterInt64, "iterations with exponent:", exponent)
            fmt.Println("  #1 234#")
            fmt.Println("  ", πViaNewton, "The same equation on a pocket calculator gives 3.14171 with only 3 terms e7")
            fmt.Println("   3.141592653589793 is the value of π from the web") 
            fmt.Println("  #1 234567890123456# :: counting the first 16 actual digits of π")
            fmt.Println(" ")
        }
        if iterInt64 == 5 {
            πViaNewton = ( twelve * ( (1/2.0) + SumOfTerms - (math.Sqrt(3)/8.0) )  ) 
            fmt.Println("After", iterInt64, "iterations with exponent:", exponent)
            fmt.Println("  #1 234#")
            fmt.Println("  ", πViaNewton, "The same equation on a pocket calculator gives 3.14171 with only 3 terms e7")
            fmt.Println("   3.141592653589793 is the value of π from the web") 
            fmt.Println("  #1 234567890123456# :: counting the first 16 actual digits of π")
            fmt.Println(" ")
        }
        if iterInt64 == 6 {
            πViaNewton = ( twelve * ( (1/2.0) + SumOfTerms - (math.Sqrt(3)/8.0) )  ) 
            fmt.Println("After", iterInt64, "iterations with exponent:", exponent)
            fmt.Println("  #1 234#")
            fmt.Println("  ", πViaNewton, "The same equation on a pocket calculator gives 3.14171 with only 3 terms e7")
            fmt.Println("   3.141592653589793 is the value of π from the web") 
            fmt.Println("  #1 234567890123456# :: counting the first 16 actual digits of π")
        }
    }

    fmt.Println("\nNewton also found π to 16 digits by the first 22 terms of :")
    fmt.Println("3*Sqrt(3)/4 + 24*( 2/3*2^3 - 1/5*2^5 - 1/28*2^7 - 1/72*2^9 - 5/704*2^11 - 7/1664*2^13 )")  // that psudocode is emplemented below
    π = ( (3*math.Sqrt(3))/4) + ( 24* ( 2/(3*math.Pow(2,3)) - 1/(5*math.Pow(2,5)) - 1/(28*math.Pow(2,7)) - 1/(72*math.Pow(2,9)) - 5/(704*math.Pow(2,11)) - 7/(1664*math.Pow(2,13)) )  )
    fmt.Println(π,"per the above formula (only six terms, of the supposed 22 terms that Newton purportedly had used)")
    fmt.Println("3.141592653589793 is the value of π from the web\n") 
}`
    fmt.Println(WorstMethodRune)
}

func showMagicBehindArchimedes() {
    var ArchimedesRune = `
    func Archimedes(num int){
    fmt.Println("\nYou selected #", num, "  --  An improved version of Archimedes' method")
    fmt.Println("  -- enter \"6' at the main menu for the derivation and proof of the Pythagorean -- ")
                // the above escape does not seem to work as advertised ??
            start := time.Now()
            iterFloat64 = 0
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
        iterInt64 = 0
    for iterInt64 < 27 {
        iterInt64++
        iterFloat64++
        n = n*2
        s1 = s2 
        s1_2 = s1/2 
        a = math.Sqrt(r-(math.Pow(s1_2,2))) 
        b = 1 - a 
        s2 = math.Sqrt(math.Pow(b,2) + (math.Pow(s1_2,2))) 
        p = n * s1 
        p_d = p/2 

        if iterInt64 == 2 {
                fmt.Println("\n   #1 2# :: counting the first 2 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 2# :: counting the first 3 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 2 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", n)
        }    
        if iterInt64 == 3 {
                fmt.Println("\n   #1 23# :: counting the first 3 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 23# :: counting the first 3 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 2 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", n)
        }    
        if iterInt64 == 7 {
                fmt.Println("\n   #1 23456# :: counting the first 6 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 23456# :: counting the first 6 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 5 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", n)
        }    
        if iterInt64 == 12 {
                fmt.Println("\n   #1 23456789# :: counting the first 9 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 23456789# :: counting the first 9 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 8 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", n)
        }    
        if iterInt64 == 17 {
                fmt.Println("\n   #1 23456789012# :: counting the first 12 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 23456789012# :: counting the first 12 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 12 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", n)
        }    
        if iterInt64 == 23 {
                fmt.Println("\n   #1 23456789012345# :: counting the first 15 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes'") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 23456789012345# :: counting the first 15 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 15 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", n)
        }
        if iterInt64 == 25 {  // last one 
                fmt.Println("\n   #1 234567890123456# :: counting the first 16 actual digits of π")
                fmt.Println("   ", p_d, " is what we have calculated per Archimedes' -- 16 -- digits") 
                fmt.Println("    3.141592653589793238 is the value of π from the web") 
                fmt.Println("   #1 234567890123456789# :: counting the first 19 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println (iterInt64, " iterations were completed in ", elapsed, " yielding 16 digits of π\n") 
                fmt.Printf("the above was estimated from a %.0f sided polygon\n", n)
                fmt.Printf("%.0f as parsed against ...\n", n)
                fmt.Println("100000000 which is one-hundred-million, for comparison to the above line")
                fmt.Println("... Which is to say a 201,326,592 sided polygon, Mr. A. would have wept!\n")
                fmt.Println(" per option #", num, "  --  an improved version of Archimedes' method\n")
            arch := [runeMark] Archimedes of Syracuse (287 – 212 BC) was a Greek mathematician, physicist, engineer, 
            astronomer, and inventor from the ancient city of Syracuse in Sicily. He estimated π 
            by drawing a larger regular hexagon outside a circle then a smaller regular hexagon inside
            the circle, and progressively doubling the number of sides of each regular polygon, 
            calculating the length of a side of each polygon at each step. As the number of sides 
            increases, it becomes a more accurate approximation of a circle. After four such steps, 
            when the polygons had 96 sides each, he was able to determine that the value of π lay between 
            3+1/7 (approx. 3.1429) and 3+10/71 (approx. 3.1408), consistent with its actual value of 
            approximately 3.1416 He also proved that the area of a circle was equal to π multiplied by the 
            square of the radius of the circle.[runeMark]
            fmt.Println(arch) // prints the above rune
                LinesPerIter = 18  // an estimate
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
                 fmt.Println("   -- After hitting Return for menu redisplay, enter '6' for the derivation and proof of the Pythagorean\n")

                    fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                        check(err1)                                // ... gets a file handle to dataFromConstants.txt
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
                fmt.Println("Select 12 at menu to display prior results")
        }  // end of last if
    }
}`
   fmt.Println(ArchimedesRune)
}

func showMagicBehindNilakantha() {
var NilakanthaRune = `
func Nilakantha(num int){
    fmt.Println("\n\nYou selected #", num, " a series by Indian astronomer Nilakantha Somayaji circa 1500 AD")
    fmt.Println("    π = 3 + 4/(2*3*4) - 4/(4*5*6) + 4/(6*7*8) - 4/(8*9*10) + 4/(10*11*12) ...")
    fmt.Println("    One-Hundred-Million iterations will be executed ... working ...\n")
    start := time.Now() 
        iterFloat64 = 0
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
    for iterInt64 < 100000000 {
        iterInt64++
        iterFloat64++
        digitone = digitthree 
        digittwo = digitthree + 1
        digitthree = digitthree + 2 
        nextterm = four/(digitone*digittwo*digitthree)
            if iterInt64 % 2 == 0 {  // % is modulus operator 
                firstsum = firstsum - nextterm 
            } else {
                firstsum = firstsum + nextterm
            }
            if iterInt64 == 100 {
                fmt.Println("   #1 234567# ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.141592,653589793 is, again, the value of π from the web") 
                fmt.Println("   #1 234567 890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println(iterInt64, " iterations were completed in ", elapsed, " yielding 7 digits of π\n") 
            }
            if iterInt64 == 500 {
                fmt.Println("   #1 23456789# ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.14159265,3589793 is, again, the value of π from the web") 
                fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println(iterInt64, " iterations were completed in ", elapsed, " yielding 9 digits of π\n") 
            }
            if iterInt64 == 10000 {
                fmt.Println("   #1 234567890123# ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.141592653589,793 is, again, the value of π from the web") 
                fmt.Println("   #1 234567890123 456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("10,000 iterations were completed in ", elapsed, " yielding 13 digits of π\n") 
            }
            if iterInt64 == 50000 {
                fmt.Println("   #1 2345678901234# ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.1415926535897,93 is, again, the value of π from the web") 
                fmt.Println("   #1 2345678901234 56# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("50,000,000 iterations were completed in ", elapsed, " yielding 14 digits of π\n") 
            }
            if iterInt64 == 1000000 {
                fmt.Println("   #1 2345678901234# ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.1415926535897,93 is, again, the value of π from the web") 
                fmt.Println("   #1 2345678901234 56# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("1,000,000 iterations were completed in ", elapsed, " still yielding 14 digits of π\n") 
            }
            if iterInt64 == 100000000 {
                fmt.Println("   #1 2345678901234# ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.1415926535897,93 is, again, the value of π from the web") 
                fmt.Println("   #1 2345678901234 56# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("100,000,000 iterations were completed in ", elapsed, " still yielding 14 digits of π\n") 
                LinesPerIter = 15
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")

                fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                    check(err1)                                // ... gets a file handle to dataFromConstants.txt
                    defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                Hostname, _ := os.Hostname()
                _ , err0 := fmt.Fprintf(fileHandle, "\n  -- Nilakantha Somayaji -- selection #%d on %s \n", num, Hostname)
                    check(err0)
                current_time := time.Now()
                _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                check(err6)
                _ , err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second \n", LinesPerSecond) 
                    check(err2)
                _ , err4 := fmt.Fprintf(fileHandle, "%e was Iterations/Seconds \n", iterFloat64/elapsed.Seconds())
                    check(err4)
                _ , err5 := fmt.Fprintf(fileHandle, "%e was total Iterations \n", iterFloat64)
                    check(err5)
                TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
                    check(err7)
            //
                fmt.Println("  -- If we ran 50 billion more iterations we still would get only those 14 digits :(\n") // =========
                fmt.Println(" per option #", num, "  --  the Nilakantha Somayaji series, circa 1500 AD\n")
                fmt.Println("Select 12 at menu to display prior results")
            }  // end of last if in for loop 
            /*
            if iterInt64 == 1000000000 {
                fmt.Println("    1 2345678901234 ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.1415926535897,93 is, again, the value of π from the web") 
                fmt.Println("    1 2345678901234 56 :: counting the first 16 actual digits of π")
                t := time.Now() ; elapsed := t.Sub(start)
                fmt.Println(iterInt64, " iterations were completed in ", elapsed, " yielding still just 14 digits of π\n") 
            }
            if iterInt64 == 5000000000 {
                fmt.Println("    1 2345678901234 ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.1415926535897,93 is, again, the value of π from the web") 
                fmt.Println("    1 2345678901234 56 :: counting the first 16 actual digits of π")
                t := time.Now() ;  elapsed := t.Sub(start)
                fmt.Println(iterInt64, " 5 billion iterations were completed in ", elapsed, " yielding still just 14 digits of π\n") 
// optional code to show the futility of more iterations 
            }            
            if iterInt64 == 9000000000 {
                fmt.Println("    1 2345678901234 ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.1415926535897,93 is, again, the value of π from the web") 
                fmt.Println("    1 2345678901234 56 :: counting the first 16 actual digits of π")
                t := time.Now() ;  elapsed := t.Sub(start)
                fmt.Println(iterInt64, " 9 billion iterations were completed in ", elapsed, " yielding still just 14 digits of π\n") 
            }
            if iterInt64 == 50000000000 {  // no additional digit are obtained even after 500 billion iterations
                fmt.Println("    1 2345678901234567 ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.141592653589793 is, again, the value of π from the web") 
                fmt.Println("    1 2345678901234567 :: counting the first 16 actual digits of π")
                t := time.Now() ;  elapsed := t.Sub(start)
                fmt.Println(iterInt64, " 50 billion iterations were completed in ", elapsed, " yielding still just 14 digits of π\n") 
            }
            */
    } // end of for loop 
}
`
   fmt.Println(NilakanthaRune)
}

func showMagicBehindGottfriedWilhelmLeibniz(){
var GottfriedWilhelmLeibnizRune = `
func GottfriedWilhelmLeibniz(num int){
    fmt.Println("\n\nYou selected #", num, " Gottfried Wilhelm Leibniz formula  :  π = 4 * ( 1 - 1/3 + 1/5 - 1/7 + 1/9 ... ")
    fmt.Println("   Infinitesimal calculus was developed independently in the late 17th century by Isaac Newton")
    fmt.Println("    ... and Gottfried Wilhelm Leibniz")
    fmt.Println("   4 Billion iterations will be executed ... ")
    fmt.Println("")
    fmt.Println(" ... working ...\n")
        start := time.Now()
        iterFloat64 = 0
        var denom float64
            denom = 3
        var sum float64
            sum = 1-(1/denom)
        iterInt64 = 1
    for iterInt64 < 4000000000 {
        iterFloat64++
        iterInt64++
        denom = denom + 2
            if iterInt64 % 2 == 0 {
                sum = sum + 1/denom 
            } else { 
                sum = sum - 1/denom 
            }
            π = 4 * sum 
                if iterInt64 == 100000000 {fmt.Println("... 100,000,000 completed iterations ...")
                    fmt.Println("\n   #1 2345678#")
                    fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                    fmt.Println("    3.1415926,53589793 is from the web") 
                    fmt.Println("   #1 2345678 90123456# :: counting the first 16 actual digits of π")
                    t := time.Now()
                    elapsed := t.Sub(start)
                    fmt.Print("  100,000,000 iterations in ", elapsed, " yields 8 digits of π\n\n")
                }
                if iterInt64 == 200000000 {fmt.Println("... 200,000,000 gets another digit ...")
                    fmt.Println("\n   #1 23456789#")
                    fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                    fmt.Println("    3.14159265,3589793 is from the web") 
                    fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                    t := time.Now()
                    elapsed := t.Sub(start)
                    fmt.Print("  200,000,000 iterations in ", elapsed, " yields 9 digits of π\n\n")
                }
                if iterInt64 == 400000000 {fmt.Println("... 400,000,000 iterations completed, still at nine ...")
                    fmt.Println("\n   #1 23456789#")
                    fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                    fmt.Println("    3.14159265,3589793 is from the web") 
                    fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                    t := time.Now()
                    elapsed := t.Sub(start)
                    fmt.Print("  400,000,000 iterations in ", elapsed, " yields 9 digits of π\n\n")
                }
                if iterInt64 == 600000000 {fmt.Println("... 600,000,000 iterations, still at nine ...")
                    fmt.Println("\n   #1 23456789#")
                    fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                    fmt.Println("    3.14159265,3589793 is from the web") 
                    fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                    t := time.Now()
                    elapsed := t.Sub(start)
                    fmt.Print("  600,000,000 iterations in ", elapsed, " yields 9 digits of π\n\n")
                }
                if iterInt64 == 1000000000 {fmt.Println("... 1 Billion iterations completed, still nine ...")
                    fmt.Println("\n   #1 23456789#")
                    fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                    fmt.Println("    3.14159265,3589793 is from the web") 
                    fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                    t := time.Now()
                    elapsed := t.Sub(start)
                    fmt.Print("  1,000,000,000 iterations in ", elapsed, " yields 9 digits of π\n\n")
                }
                if iterInt64 == 2000000000 {fmt.Println("... 2 Billion, and still just nine ...")
                    fmt.Println("\n   #1 23456789#")
                    fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                    fmt.Println("    3.14159265,3589793 is from the web") 
                    fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                    t := time.Now()
                    elapsed := t.Sub(start)
                    fmt.Print("  2,000,000,000 iterations in ", elapsed, " yields 9 digits of π\n\n")
                }
                if iterInt64 == 4000000000 {  // last one
                    fmt.Println("\n... 4 Billion, gets us ten digits  ...")
                    fmt.Println("\n   #1 234567890#")
                    fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                    fmt.Println("    3.141592653,589793 is from the web") 
                    fmt.Println("   #1 234567890 123456# :: counting the first 16 actual digits of π")
                    t := time.Now()
                    elapsed := t.Sub(start)
                    fmt.Print("  4,000,000,000 iterations in ", elapsed, " yields 10 digits of π\n\n")
                    fmt.Println(" per option #", num, "  --  the Gottfried Wilhelm Leibniz formula\n")

                    LinesPerIter = 14
                    fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                    LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                    fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                    fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")
                    fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                        check(err1)                                // ... gets a file handle to dataFromConstants.txt
                        defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.

                    Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  -- Gottfried Wilhelm Leibniz -- selection #%d on %s \n", num, Hostname)  
                        check(err0)
                        current_time := time.Now()
                    _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                        check(err6)
                    _ , err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second  \n", LinesPerSecond) 
                        check(err2)
                    _ , err4 := fmt.Fprintf(fileHandle, "%e was Iterations/Seconds  \n", iterFloat64/elapsed.Seconds())
                        check(err4)
                    _ , err5 := fmt.Fprintf(fileHandle, "%e was total Iterations  \n", iterFloat64)
                        check(err5)
                    TotalRun := elapsed.String() // cast time duration to a String type for Fprintf "formatted print"
                    _ , err7 := fmt.Fprintf(fileHandle, "Total runTime was %s \n", TotalRun)  // add total runtime of this calculation 
                        check(err7)
                    fmt.Println("Select 12 at menu to display prior results")
                }
/* we skip the rest for now 
                if iterInt64 == 6000000000 {fmt.Println("... 6 Billion completed.  ...")
                fmt.Println("   #1 234567890#")
                fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                fmt.Println("    3.141592653,589793 is from the web") 
                fmt.Println("   #1 234567890 123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Print("  6,000,000,000 iterations in ", elapsed, " still yields 10 digits of π\n")
                }
                if iterInt64 == 8000000000 {fmt.Println("... 8 Billion completed. still ten ...")
                fmt.Println("   #1 234567890#")
                fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                fmt.Println("    3.141592653,589793 is from the web") 
                fmt.Println("   #1 234567890 123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Print("  8,000,000,000 iterations in ", elapsed, " still yields 10 digits of π\n")
                }
            if iterInt64 == 9000000000 {
                fmt.Println("   #1 234567890#")
                fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                fmt.Println("    3.141592653,589793 is from the web") 
                fmt.Println("   #1 234567890 123456# :: counting the first 16 actual digits of π")
                // fmt.Print("   ", iter) 
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Print("\n... 9B iterations in ", elapsed, " , but to get 10 digits we only needed 4B iterations\n\n") 
                fmt.Println(" per option #", num, "  --  the Gottfried Wilhelm Leibniz formula")
            }  
*/
    } // end of for
}`
fmt.Println(GottfriedWilhelmLeibnizRune)
}

func showMagicBehindGregoryLeibniz(){
var GregoryLeibnizRune = `func GregoryLeibniz(num int){
    fmt.Println("\n\nYou selected #", num, " the Gregory-Leibniz series ...")
    fmt.Println("π = (4/1) - (4/3) + (4/5) - (4/7) + (4/9) - (4/11) + (4/13) - (4/15) ...")
    fmt.Println("Three hundred million iterations are being executed ... working ...\n")
    start := time.Now()
    iterFloat64 = 0
    var nextOdd float64
        nextOdd = 1
    four = 4
    var tally float64 
        tally = (four/nextOdd)
        iterInt64 = 0
    for iterInt64 < 300000000 {
        iterInt64++
        iterFloat64++
        nextOdd = nextOdd + 2
        tally = tally-(tally/nextOdd)
        tally = tally+(tally/nextOdd) // pi (tally) is set equl to the sum of a subtraction and an addition, alternatively 

            if iterInt64 == 10000000 { fmt.Println("... 10,000,000 of three hundred million completed. still working, but ...")
                fmt.Println("\n   #1 234567#")
                fmt.Println("   ",tally, "was calculated by the Gregory-Leibniz series") 
                fmt.Println("    3.141592,653589793 is from the web") 
                fmt.Println("   #1 234567 890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Print("  10,000,000 iterations in ", elapsed, " yields 7 digits of π\n\n")
            }
// 7 digits per above
// 8, next two ifs give eight digits
            if iterInt64 == 50000000 { fmt.Println("... 50,000,000 of three hundred million completed. still working, but ...")
                fmt.Println("\n   #1 2345678#")
                fmt.Println("   ",tally, "was calculated by the Gregory-Leibniz series") 
                fmt.Println("    3.1415926,53589793 is from the web") 
                fmt.Println("   #1 2345678 90123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Print("  50,000,000 iterations in ", elapsed, " yields 8 digits of π\n")
                fmt.Println(" ")
            }
            if iterInt64 == 100000000 { fmt.Println("... 100,000,000 of three hundred million completed. still working, but ...")
                fmt.Println("\n   #1 2345678#")
                fmt.Println("   ",tally, "was calculated by the Gregory-Leibniz series") 
                fmt.Println("    3.1415926,53589793 is from the web") 
                fmt.Println("   #1 2345678 90123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Print("  100,000,000 iterations in ", elapsed, " yields 8 digits of π\n\n")
            }
// 9 digits are found below 
            if iterInt64 == 200000000 { fmt.Println("... 200,000,000 of three hundred million completed. still working, but ...")
                fmt.Println("\n   #1 23456789#")
                fmt.Println("   ",tally, "was calculated by the Gregory-Leibniz series") 
                fmt.Println("    3.14159265,3589793 is from the web") 
                fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Print("  200,000,000 iterations in ", elapsed, " yields 9 digits of π\n")
                }
            if iterInt64 == 300000000 {  // last one, still 9 digits
                fmt.Println("\n   #1 23456789#")
                fmt.Print ("    ", tally, " was calculated by the Gregory-Leibniz series")
                fmt.Println("\n    3.141592653589793 is from the web") 
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Print ("  300 million iterations still yields 9 digits, ") // no Println here
                fmt.Print ("in ", elapsed, "\n\n") 
                fmt.Println(" per option #", num, "  --  the Gregory-Leibniz series, circa 1676\n")

                LinesPerIter = 11  // an estimate of the number of lines per iteration 
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("       %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")
//
                fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                    check(err1)                                // ... gets a file handle to dataFromConstants.txt
                    defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                Hostname, _ := os.Hostname()
                _ , err0 := fmt.Fprintf(fileHandle, "\n  -- Gregory-Leibniz -- selection #%d on %s \n", num, Hostname)
                    check(err0)
                current_time := time.Now()
                _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                check(err6)
                _ , err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second  \n", LinesPerSecond) 
                    check(err2)
                _ , err4 := fmt.Fprintf(fileHandle, "%e was Iterations/Seconds  \n", iterFloat64/elapsed.Seconds())
                    check(err4)
                _ , err5 := fmt.Fprintf(fileHandle, "%e was total Iterations  \n", iterFloat64)
                    check(err5)
                TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
                    check(err7)
                fmt.Println("Select 12 at menu to display prior results")
            }
    } 
}
`
fmt.Println(GregoryLeibnizRune)
}

func showMagicBehindJohnWallis() {
var JohnWallisRune = `
func JohnWallis(num int){
    fmt.Println("\n   You selected #", num, " A Go language exercize which can be used to test the speed of your hardware.")
    fmt.Println("   We will calculate π to a maximum of ten digits of accuracy using an infinite series by John Wallis circa 1655")
    fmt.Println("   Up to 40 Billion iterations of the following formula will be executed ")
    fmt.Println("   π = 2 * ((2/1)*(2/3)) * ((4/3)*(4/5)) * ((6/5)*(6/7)) ...")
    start := time.Now() 
    iterFloat64 = 0
    var numerators float64
        numerators = 2
    var firstDenom float64
        firstDenom = 1
    var secondDenom float64
        secondDenom = 3
    var cumulativeProduct float64
        cumulativeProduct = (numerators/firstDenom) * (numerators/secondDenom)
          iterInt64 = 0
      for iterInt64 < 40000000000 {
        iterInt64++
        iterFloat64++
        numerators = numerators + 2
        firstDenom = firstDenom + 2
        secondDenom = secondDenom + 2
        cumulativeProduct = cumulativeProduct * (numerators/firstDenom) * (numerators/secondDenom)
        π = cumulativeProduct * 2  
        if iterInt64 == 10 { fmt.Println(" ") }  //  single-line "if" just to display a blank line; for amusement 
/*            if iterInt64 == 100 {
                fmt.Println("   #1 2# ")
                fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                fmt.Println("    3.1,41592653589793 is, again, the value of π from the web") 
                fmt.Println("   #1 2# 34567890123456 :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println(iterInt64, " iterations were completed in ", elapsed, " yielding 2 digits of π\n") 
            }
            if iterInt64 == 500 {
                fmt.Println("   #1 23# ")
                fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                fmt.Println("    3.14,1592653589793 is, again, the value of π from the web") 
                fmt.Println("   #1 23# 4567890123456 :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println(iterInt64, " iterations were completed in ", elapsed, " yielding 3 digits of π\n") 
            }
*/
            if iterInt64 == 2000 {
                fmt.Println("   #1 234# ")
                fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                fmt.Println("    3.14,1592653589793 is, again, the value of π from the web") 
                fmt.Println("   #1 23 4567890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println(iterInt64, " iterations were completed in ", elapsed, " yielding 4 digits of π\n") 
            }
            if iterInt64 == 10000 {
                fmt.Println("   #1 2345# ")
                fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                fmt.Println("    3.1415,92653589793 is, again, the value of π from the web") 
                fmt.Println("   #1 2345 67890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("10,000 iterations were completed in ", elapsed, " yielding 5 digits of π\n") 
            }
            if iterInt64 == 50000 { // 50,000 
                fmt.Println("   #1 2345# ")
                fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                fmt.Println("    3.1415,92653589793 is, again, the value of π from the web") 
                fmt.Println("   #1 2345 67890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("50,000 iterations were completed in ", elapsed, " yielding 5 digits of π\n") 
            }
            if iterInt64 == 500000 {  // 500,000 done
                fmt.Println("   #1 23456# ")
                fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                fmt.Println("    3.14159,2653589793 is, again, the value of π from the web") 
                fmt.Println("   #1 23456 7890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("500,000 iterations were completed in ", elapsed, " yielding 6 digits of π\n") 
            }
            if iterInt64 == 2000000 {  // 2M done 
                fmt.Println("   #1 234567# ")
                fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                fmt.Println("    3.141592,653589793 is, again, the value of π from the web") 
                fmt.Println("   #1 234567 890123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("2,000,000 iterations were completed in ", elapsed, " yielding 7 digits of π\n") 
            }
            if iterInt64 == 40000000 {  // 40M done
                fmt.Println("   #1 2345678# ")
                fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                fmt.Println("    3.1415926,53589793 is, again, the value of π from the web") 
                fmt.Println("   #1 2345678 90123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("40,000,000 iterations were completed in ", elapsed, " yielding 8 digits of π\n\n") 
                fmt.Println("  .. working .. on another factor-of-ten iterations\n")
            }
            if iterInt64 == 400000000 {  // 400M done
                fmt.Println("   #1 23456789# ")
                fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
                fmt.Println("    3.14159265,3589793 is, again, the value of π from the web") 
                fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println("400,000,000 iterations were completed in ", elapsed, " yielding 9 digits of π\n\n") 

                LinesPerIter = 36 // an estimate 
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")
//
                fmt.Println("  ... will be working on doing Billions more iterations ...\n\n")
            }
//
            if iterInt64 == 600000000 {  // 600M done
                fmt.Println("  600M done, still working on another Two-Hundred-Thousand iterations ... working ...\n")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println(elapsed, "\n")
                fmt.Println("Calculating 10th digit (40B iters) which takes a few min \n")
                fmt.Println("- Ctrl-C to End/Exit without saving results\n")
                LinesPerIter = 36 // an estimate 
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")
                fmt.Println(" ... still working ...")
            }

            if iterInt64 == 800000000 {  // 800M done
                fmt.Println("  800M done, still working on yet another Two Hundred Thousand iterations ... working ...\n")
                t := time.Now()
                elapsed := t.Sub(start)
                fmt.Println(elapsed, " \n")
            }
        if iterInt64 == 1000000000 {  // 1B done
            fmt.Println("   #1 23456789# ")
            fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
            fmt.Println("    3.14159265,3589793 is the value of π from the web") 
            fmt.Println("   #1 23456789 0123456# :: counting the first 16 actual digits of π")
            t := time.Now()
            elapsed := t.Sub(start)
            fmt.Println("\nOne Billion iterations were completed in ", elapsed, " still only yielding π to 9 digits\n") 
            fmt.Println(" per option #", num, "  --  an infinite series by John Wallis circa 1655\n") // ----------------------

                LinesPerIter = 36 // an estimate 
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")
//
            fmt.Println("  ... still working ... on Billions of iterations, 39 to go ...\n\n")
        }

//  "Optional" Forty Billion loops just to get one additional digit, the tenth digit of pi 

        if iterInt64 == 2000000000 {  // 2B completed 
            fmt.Println("  2B done, still working ... on another Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 3000000000 {  // 3B completed
            fmt.Println("  3B done, still working ... on another Billion iterations ... working ... Ctrl-C to End/Exit without saving stats")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 4000000000 {  // 4B completed 
            fmt.Println("  4B done, still working ... on another Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 5000000000 {  // 5B completed 
            fmt.Println("  5B done, still working ... on another Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 6000000000 {  // 6B completed 
            fmt.Println("  6B done, still working ... on another Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 7000000000 {  // 7B completed 
            fmt.Println("  7B done, still working ... on another Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 8000000000 {  // 8B completed 
            fmt.Println("  8B done, still working ... on another Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 9000000000 {  // 9B completed 
            fmt.Println("  9B done, still working ... on another five Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 14000000000 {  // 14B completed
            fmt.Println("  14B done, still working ... on another five Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 19000000000 {  // 19B completed 
            fmt.Println("  19B done, still working ... on another five Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 24000000000 { // 24B completed
            fmt.Println("  24B done, still working ... on another five Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 29000000000 {  // 29B completed 
            fmt.Println("  29B done, still working ... on another five Billion iterations ... working ...")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
        if iterInt64 == 34000000000 {  // 34B completed
            fmt.Println("  34B done, still working ... just another six Billion iterations to go! ... ")
                t := time.Now()
                elapsed := t.Sub(start)
            fmt.Println(elapsed)
        }
 //  "Optional" Forty Billion loops to get another digit
        if iterInt64 == 40000000000 {  // 40B completed 
            fmt.Println("   #1 234567890# ")
            fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
            fmt.Println("    3.141592653,589793 is the value of π from the web") 
            fmt.Println("   #1 234567890 123456# :: counting the first 16 actual digits of π")
            t := time.Now()
            elapsed := t.Sub(start)
            fmt.Println("Forty Billion iterations were completed in ", elapsed, " yielding π to 10 digits\n") 
            fmt.Println(" per option #", num, "  --  an infinite series by John Wallis circa 1655\n") // ----------------------
                LinesPerIter = 36 // an estimate 
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")
                fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                    check(err1)                                // ... gets a file handle to dataFromConstants.txt
                    defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                Hostname, _ := os.Hostname()
                _ , err0 := fmt.Fprintf(fileHandle, "\n  -- John Wallis -- selection #%d on %s \n", num, Hostname)
                    check(err0)
                current_time := time.Now()
                _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                check(err6)
                _ , err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second  \n", LinesPerSecond) 
                    check(err2)
                _ , err4 := fmt.Fprintf(fileHandle, "%.02f was Iterations/Seconds  \n", iterFloat64/elapsed.Seconds())
                    check(err4)
                _ , err5 := fmt.Fprintf(fileHandle, "%e was total Iterations  \n", iterFloat64)
                    check(err5)
                TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun)  // add total runtime of this calculation 
                    check(err7)
                fmt.Println("Select 12 at menu to display prior results")
        }
    }
}`
fmt.Println(JohnWallisRune)
}

func showMagicBehindEulersNumber(){
var EulersNumberRune = `
func EulersNumber(num int){
    var n float64 
    var sum float64
    var Eulers float64
    fmt.Println("\n\n\nEuler's Number \u2107 or \u2147 the natural logarithmic base")
    fmt.Println("\u2147 = (1+1/n)^n")
    fmt.Println("... the limit of an increasing value for n\n\n") 
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
// And now, here comes a fun rune to print a multi-line "string"
// ... define a rune with the [runMark] :: back-quote character located on the ~ tilda key
Ricks_rune_Paragraph := [runMark]  
Bernoulli noticed that this sequence approaches a limit (the force of interest) with larger n and, thus, smaller 
compounding intervals. Compounding weekly (n = 52) yields $2.692596..., while compounding daily (n = 365) yields
$2.714567... (approximately two cents more). The limit as n grows large is the number that came to be known as e.
That is, with continuous compounding, the account value will reach $2.718281828 
[runMark]
fmt.Println(Ricks_rune_Paragraph)
}`
fmt.Println(EulersNumberRune)
}

func showMagicBehindErdosBorwein() {
var ErdosBorweinRune = `
func ErdosBorwein(num int){ 
       rune := [runeMark]The Erdős–Borwein constant is the sum of the reciprocals of the Mersenne numbers. 
It is named after Paul Erdős and Peter Borwein. 

Paul Erdős was a Hungarian mathematician. He was one of the most prolific mathematicians 
and producers of mathematical conjectures of the 20th century. Erdős pursued and proposed 
problems in discrete mathematics, graph theory, number theory, mathematical analysis, 
approximation theory, set theory, and probability theory.[4] Much of his work centered 
around discrete mathematics, cracking many previously unsolved problems in the field. 
He championed and contributed to Ramsey theory, which studies the conditions in which 
order necessarily appears. Overall, his work leaned towards solving previously open 
problems, rather than developing or exploring new areas of mathematics.

Erdős published around 1,500 mathematical papers during his lifetime, a figure that
remains unsurpassed.[5] He firmly believed mathematics to be a social activity, 
living an itinerant lifestyle with the sole purpose of writing mathematical papers 
with other mathematicians. He was known both for his social practice of mathematics, 
working with more than 500 collaborators, and for his eccentric lifestyle; Time magazine 
called him "The Oddball's Oddball".[6] He devoted his waking hours to mathematics, even 
into his later years—indeed, his death came only hours after he solved a geometry 
problem at a conference in Warsaw. Erdős's prolific output with co-authors prompted 
the creation of the Erdős number, the number of steps in the shortest path between a 
mathematician and Erdős in terms of co-authorships. 


Peter Benjamin Borwein (born St. Andrews, Scotland, May 10, 1953 – 23 August 2020) 
was a Canadian mathematician and a professor at Simon Fraser University. He is known 
as a co-author of the paper which presented the Bailey–Borwein–Plouffe algorithm 
(discovered by Simon Plouffe) for computing π. 

Borwein was born into a Jewish family. He became interested in number theory and classical 
analysis during his second year of university. He had not previously been interested in 
math, although his father was the head of the University of Western Ontario's mathematics 
department and his mother is associate dean of medicine there. Borwein and his two siblings 
majored in mathematics. 

After completing a Bachelor of Science in Honours Math at the University of Western 
Ontario in 1974, he went on to complete an MSc and Ph.D. at the University of British 
Columbia. He joined the Department of Mathematics at Dalhousie University. While he 
was there, he, his brother Jonathan Borwein and David H. Bailey of NASA wrote the 1989 
paper[1] that outlined and popularized a proof for computing one billion digits of π. 
The authors won the 1993 Chauvenet Prize and Merten M. Hasse Prize for this paper.

In 1993, he moved to Simon Fraser University, joining his brother Jonathan in establishing 
the Centre for Experimental and Constructive Mathematics (CECM) where he developed the 
Inverse Symbolic Calculator. [runeMark]

fmt.Println("") // so as not to pad the following rune with an unwanted space 
    fmt.Println(rune, "\n")
    fmt.Println("We calculate E as E = the sum of 1/((2^n)-1) as n grows from 1 to 'infinity'")

    var Erdos_Borwein float64
        Erdos_Borwein = 1
    var iter float64 
        iter = 1
    for iter < 100 {
        iter++ // iter will therefore begin at 2
        Erdos_Borwein = Erdos_Borwein + (1/( (math.Pow(2,iter)) - 1) )
        if iter == 10 || iter == 20 || iter == 30 || iter == 40 || iter == 50 || iter == 60 || iter == 70 || iter == 100 || iter == 101 { fmt.Println(Erdos_Borwein)}
                                                                                                    // 100 and 101 prove that we ended on 100 as the final exponent
                                                                                                    // ... we only get 8 results, not nine
    }
    fmt.Println("for 10, 20, 30, 40, 50, 60, 70, and 100 iterations respectively\n")
    fmt.Println("Our calculated Erdos-Borwein constant is ")
    fmt.Println(Erdos_Borwein, "after", iter, "iterations, i.e., with a final exponent of", iter)
    fmt.Println("1.606695152415291763 is what we get from the web\n")
}`
fmt.Println(ErdosBorweinRune)
}

func showMagicBehindmainFunc() {
var mainFuncRune = `
// compile with: "go build -o constants constants.go"
// ... thereafter you can run it on a Unix system with "/fullpathname/constants"
// ... or, having first obtained the Go compiler, ... just run the source with: "go run constants.go"
// One can obtain the Go Compiler from : https://go.dev/dl/

package main
import ("fmt"; "time"; "math"; "os"; "io/ioutil"; "strconv") // include additional packages 
func main() {        // top-level program logic flow 
    for 1 == 1 {    // loop endlessly, or Ctrl-C to Exit
        RicksDisplayMenuPi()   // displays the menu 
        RicksSwitch()       // contains a Scanf to input a menu selection 
/*
    Additional kluge pause -- which is not needed on a Unix System, and seems to be the only mod needed on Windws-11
            var Pnum int // intended to be used only as a dummy 
            fmt.Println("Hit Enter/Return to redisplay the menu") 
            fmt.Scanf("%d", &Pnum) // often flys right-past this Scanf 
*/
// Another kluge pause (the above block) is also "needed" in some environments because they sometimes ... 
// ... fly rite-past that first one and then redisplay the menu (Windows 11 does this)
// but even Unix style systems seem to require a single kluge like this one 
        var Mnum int // just a local kluge element 
        fmt.Println("Hit Enter/Return again to redisplay the menu") 
        fmt.Scanf("%d", &Mnum) // request input on command line (pause)
        if Mnum > 0 && Mnum < 17 {  // the user has attempted to make a selection having failed to first "Hit Enter" as requested 
            fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")  // so, rattle his chain
        }  
    } 
} 

        var TotalIterations int                // gloabl vars for use in multiple localized sections of code 
        var Prospective_smaller_hit int 
        var ProspectiveHitOnLargeSide int 
        var savedHitB int
        var smallerPerfectSquareOnceSaved int   
        var savedHit int   // will be used in multiple functions (func)s 
        var four float64  // is initialized to 4 where needed 
        var π float64  // a var can be any character, as in this Pi symbol/character 
        var LinesPerSecond float64  
        var LinesPerIter float64
        var iterInt64 int64  // to be used primarily in selections which require modulus calculations 
        var iterFloat64 float64  // to be used in selections which do not require modulus calculations 
        var root = 100 // create a var of type int and initialize as 100, this does need to be global 
        var PerfectSquare = root * root // create a global var of type int, this too needs to be global 
        var sliceOfPerfectSquares_and_their_roots = []int{ PerfectSquare , root } // initialize slice with 100
        var sliceOfP_diffs_smlrT3_to_largerOfProspectiveHit = []float32{99.9999}  // init with a large-ish p_diff_from_larger_hit
        var sliceOfP_diffs_smlrT3_to_smlrOfProspectiveHit = []float32{88.888}     //   "            "      p_diff_from_smlr_hit


func check(e error) {   // create a func named check which takes one parameter "e" of type error 
    if e != nil {
        panic(e)        // use panic() to display error code 
    }
}
`
fmt.Println(mainFuncRune)
}

// montecarlopi.go
// description: Calculating pi by the Monte Carlo method
// details:
// implementations of Monte Carlo Algorithm for the calculating of Pi - [Monte Carlo method](https://en.wikipedia.org/wiki/Monte_Carlo_method)
// author(s): [red_byte](https://github.com/i-redbyte), [Paul Leydier] (https://github.com/paul-leydier)
// see montecarlopi_test.go

//  as input to "func MonteCarloPiConcurrent(n int) (float64, error)" yeilds: 
// 3.141610315914161
// 3.1415926535897932384626433832795028841971693993 is from the web
// 1 2345 so this method is not even good for 4 digits of pi 
// but be forewarned, this code will run all your cores at full capacity and will toast your CPU ... 
// ... if it has been overclocked in the least and you let this baby knaw at these loops for too long. 


// 9999999999 as input to "func MonteCarloPi(randomPoints int) float64 {" yeilds: 
// 3.1416128591141614
// 3.1415926535897932384626433832795028841971693993 is from the web
// 1 2345 so this method is not even good for 4 digits of pi 
// this method will only max-out one of your cores for a few min, while the other cores act as heat sinks 
var iters_mc int 
func MonteCarloPi(randomPoints int, num2 int) float64 {
    rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
    inside := 0
        start := time.Now() 
    for i := 0; i < randomPoints; i++ {
        x := rnd.Float64()
        y := rnd.Float64()
        if x*x+y*y <= 1 {
            inside += 1
        }
        iters_mc = i
    }
    t := time.Now()
    elapsed := t.Sub(start)
    // log stats to a log file 
        fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
            check(err1)                                // ... gets a file handle to dataFromConstants.txt
            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
        Hostname, _ := os.Hostname()
        _ , err0 := fmt.Fprintf(fileHandle, "\n  -- non-concurrent MonteCarloPi -- selection #%d on %s \n", num2, Hostname)
            check(err0)
        current_time := time.Now()
        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
        check(err6)
        _ , err4 := fmt.Fprintf(fileHandle, "%.02f was Iterations/Seconds \n", float64(iters_mc)/float64(elapsed.Seconds()))
            check(err4)
        _ , err5 := fmt.Fprintf(fileHandle, "%d was total Iterations \n", iters_mc)
            check(err5)
        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
            check(err7)
    pi := float64(inside) / float64(randomPoints) * 4
    return pi
}

// MonteCarloPiConcurrent approximates the value of pi using the Monte Carlo method.
// Unlike the MonteCarloPi function (first version), this implementation uses
// goroutines and channels to parallelize the computation.
// More details on the Monte Carlo method available at https://en.wikipedia.org/wiki/Monte_Carlo_method.
// More details on goroutines parallelization available at https://go.dev/doc/effective_go#parallel.
func MonteCarloPiConcurrent(n int, num2 int) (float64, error) {
//begin Richard H. Woolley's code block
    start := time.Now() 
//end Richard H. Woolley's code block

    numCPU := runtime.GOMAXPROCS(0)
    c := make(chan int, numCPU)
    pointsToDraw, err := splitInt(n, numCPU) // split the task in sub-tasks of approximately equal sizes
    if err != nil {
        return 0, err
    }

    // launch numCPU parallel tasks
    for _, p := range pointsToDraw {
        go drawPoints(p, c)
    }

    // collect the tasks results
    inside := 0
    for i := 0; i < numCPU; i++ {
        inside += <-c
    }
//begin Richard H. Woolley's code block
            t := time.Now()
            elapsed := t.Sub(start)
            fmt.Println(elapsed)

            fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                check(err1)                                // ... gets a file handle to dataFromConstants.txt
                defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
            Hostname, _ := os.Hostname()
            _ , err0 := fmt.Fprintf(fileHandle, "\n  -- MonteCarloPiConcurrent -- selection #%d on %s \n", num2, Hostname)
                check(err0)
            current_time := time.Now()
            _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
            check(err6)
            _ , err8 := fmt.Fprint(fileHandle, "precision was: ", n, "\n")
            check(err8)
            TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
            _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
                check(err7)
//end Richard H. Woolley's code block

    return float64(inside) / float64(n) * 4, nil
}

// drawPoints draws n random two-dimensional points in the interval [0, 1), [0, 1) and sends through c
// the number of points which where within the circle of center 0 and radius 1 (unit circle)
func drawPoints(n int, c chan<- int) {
    rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
    inside := 0
    for i := 0; i < n; i++ {
        x, y := rnd.Float64(), rnd.Float64()
        if x*x+y*y <= 1 {
            inside++
        }
    }
    c <- inside
}

// splitInt takes an integer x and splits it within an integer slice of length n in the most uniform
// way possible.
// For example, splitInt(10, 3) will return []int{4, 3, 3}, nil
func splitInt(x int, n int) ([]int, error) {
    if x < n {
        return nil, fmt.Errorf("x must be < n - given values are x=%d, n=%d", x, n)
    }
    split := make([]int, n)
    if x%n == 0 {
        for i := 0; i < n; i++ {
            split[i] = x / n
        }
    } else {
        limit := x % n
        for i := 0; i < limit; i++ {
            split[i] = x/n + 1
        }
        for i := limit; i < n; i++ {
            split[i] = x / n
        }
    }
    return split, nil
}

func showMagicBehindMonteCarloPiConcurrent() {
var MonteCarloPiConcurrentRune = `
// montecarlopi.go
// description: Calculating pi by the Monte Carlo method
// details:
// implementations of Monte Carlo Algorithm for the calculating of Pi - [Monte Carlo method](https://en.wikipedia.org/wiki/Monte_Carlo_method)
// author(s): [red_byte](https://github.com/i-redbyte), [Paul Leydier] (https://github.com/paul-leydier)
// see montecarlopi_test.go

// 9999999999 as input to "func MonteCarloPiConcurrent(n int) (float64, error)" yeilds: 
// 3.141610315914161
// 3.1415926535897932384626433832795028841971693993 is from the web
// 1 2345 so this method is not even good for 4 digits of pi 
// but be forewarned, this code will run all your cores at full capacity and will toast your CPU ... 
// ... if it has been overclocked in the least and you let this baby knaw at these loops for too long. 


// 9999999999 as input to "func MonteCarloPi(randomPoints int) float64 {" yeilds: 
// 3.1416128591141614
// 3.1415926535897932384626433832795028841971693993 is from the web
// 1 2345 so this method is not even good for 4 digits of pi 
// this method will only max-out one of your cores for a few min, while the other cores act as heat sinks 

func MonteCarloPi(randomPoints int) float64 {
    rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
    inside := 0
    for i := 0; i < randomPoints; i++ {
        x := rnd.Float64()
        y := rnd.Float64()
        if x*x+y*y <= 1 {
            inside += 1
        }
    }
    pi := float64(inside) / float64(randomPoints) * 4
    return pi
}

// MonteCarloPiConcurrent approximates the value of pi using the Monte Carlo method.
// Unlike the MonteCarloPi function (first version), this implementation uses
// goroutines and channels to parallelize the computation.
// More details on the Monte Carlo method available at https://en.wikipedia.org/wiki/Monte_Carlo_method.
// More details on goroutines parallelization available at https://go.dev/doc/effective_go#parallel.
func MonteCarloPiConcurrent(n int, num int) (float64, error) {
//begin Richard H. Woolley's code block
    start := time.Now() 
//end Richard H. Woolley's code block

    numCPU := runtime.GOMAXPROCS(0)
    c := make(chan int, numCPU)
    pointsToDraw, err := splitInt(n, numCPU) // split the task in sub-tasks of approximately equal sizes
    if err != nil {
        return 0, err
    }

    // launch numCPU parallel tasks
    for _, p := range pointsToDraw {
        go drawPoints(p, c)
    }

    // collect the tasks results
    inside := 0
    for i := 0; i < numCPU; i++ {
        inside += <-c
    }
//begin Richard H. Woolley's code block
            t := time.Now()
            elapsed := t.Sub(start)
            fmt.Println(elapsed)

            fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                check(err1)                                // ... gets a file handle to dataFromConstants.txt
                defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
            Hostname, _ := os.Hostname()
            _ , err0 := fmt.Fprintf(fileHandle, "\n  -- MonteCarloPiConcurrent -- selection #%d on %s \n", num, Hostname)
                check(err0)
            current_time := time.Now()
            _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
            check(err6)
            _ , err8 := fmt.Fprint(fileHandle, "precision was: ", n, "\n")
            check(err8)
            TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
            _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
                check(err7)
//end Richard H. Woolley's code block

    return float64(inside) / float64(n) * 4, nil
}

// drawPoints draws n random two-dimensional points in the interval [0, 1), [0, 1) and sends through c
// the number of points which where within the circle of center 0 and radius 1 (unit circle)
func drawPoints(n int, c chan<- int) {
    rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
    inside := 0
    for i := 0; i < n; i++ {
        x, y := rnd.Float64(), rnd.Float64()
        if x*x+y*y <= 1 {
            inside++
        }
    }
    c <- inside
}

// splitInt takes an integer x and splits it within an integer slice of length n in the most uniform
// way possible.
// For example, splitInt(10, 3) will return []int{4, 3, 3}, nil
func splitInt(x int, n int) ([]int, error) {
    if x < n {
        return nil, fmt.Errorf("x must be < n - given values are x=%d, n=%d", x, n)
    }
    split := make([]int, n)
    if x%n == 0 {
        for i := 0; i < n; i++ {
            split[i] = x / n
        }
    } else {
        limit := x % n
        for i := 0; i < limit; i++ {
            split[i] = x/n + 1
        }
        for i := limit; i < n; i++ {
            split[i] = x / n
        }
    }
    return split, nil
}`
fmt.Println(MonteCarloPiConcurrentRune)
}

func showMagicBehindSwitch(){
    var RicksSwitchRune = `
    switch num { 
        case 1:  // Calculate the square root of 3 from first principles of geometry
            squareRootOf3(num)
        case 21: 
            showMagicBehindsquareRootOf3()
        case 2:  // CASE 2 (best method) the Bailey–Borwein–Plouffe formula for π, circa 1995 
            BBPF(num)
        case 22:
            showMagicBehindBBPF() 
        case 3: // (worst method)
            WorstMethod(num)
        case 23: 
            showMagicBehindWorstMethod() 
        case 4:  // An improved version of Archimedes' method
            Archimedes(num)
        case 24:
            showMagicBehindArchimedes()
        case 5: // a series by Indian astronomer Nilakantha Somayaji circa 1500 AD 
            Nilakantha(num)
        case 25:
            showMagicBehindNilakantha()
        case 6: // Gottfried Wilhelm Leibniz
            GottfriedWilhelmLeibniz(num)
        case 26:
            showMagicBehindGottfriedWilhelmLeibniz()
        case 7: // the Gregory-Leibniz series
            GregoryLeibniz(num)
        case 27:
            showMagicBehindGregoryLeibniz()
        case 8: // John Wallis circa 1655 --- runs very long 
            JohnWallis(num)
        case 28: 
            showMagicBehindJohnWallis()
        case 9: // Euler's Number
            EulersNumber(num)
        case 29:
            showMagicBehindEulersNumber()
        case 10: // Erdős–Borwein constant
            ErdosBorwein(num)
        case 30:
            showMagicBehindErdosBorwein()
        case 11:  // display a review of the derivation of the Pythagorean
            DisplayPythagorean(num)
        case 31: 
            DisplayPythagorean(num)
        case 12: // display contents of prior results file
            content, err := ioutil.ReadFile("dataFromConstants.txt")  // 
                if err != nil {   // if the file does not exist ... 
                    fmt.Println("no prior results -- no log file dataFromConstants.txt exists")
                } else {
                    fmt.Println(string(content))  // dump/display entire file to command line
                }
        case 32:
            var fileAccessRune = [runeMark]
    content, err := ioutil.ReadFile("dataFromConstants.txt")  // 
        if err != nil {   // if the file does not exist ... 
            fmt.Println("no prior results -- no log file dataFromConstants.txt exists")
        } else {
            fmt.Println(string(content))  // dump/display entire file to command line
        }[runeMark]
            fmt.Println(fileAccessRune)
        case 33: 
            showMagicBehindmainFunc()
        case 13: 
            fmt.Println(" ... So sorry, but", num, "was not an option")
        case 14: 
            fmt.Println(num, " was not an option! It was not on the menu, go fish!\n")
        case 15: 
            fmt.Println("Your selection of", num, " is right-out!  Go Fish!\n")
        case 16: 
            fmt.Println("Your selection is really-far-out!  Go Fish!\n")
        case 17: 
            fmt.Println("\nOops, how'd we get here? Hit Enter/Return again to possibly redisplay the menu")
        case 18:
            secondMenu()
        case 34:
            /* // a project that I failed at
            var precision int 
            fmt.Println("what shall be thy precision? ")
            fmt.Scanf("%d", &precision) 
            fmt.Println("result is ", calculatePispi(precision))
            */
        case 35:
            showMagicBehindSwitch()
        case 36:
            ConcurrentMCpi(num)
        case 56:
            showMagicBehindMonteCarloPiConcurrent()
        case 37:
            main_juuso()
        case 57:
            showTheMagicBehind_main_juuso()
        case 58:
            os.Exit(1)
        case 98:
            fmt.Println(rick, "\n")
        case 99:
            Explain_spigot()
    default: 
        fmt.Println("this is the switch default code, after a break option ??")


    switch num2 { 
        case 19: 
            TheSpigot()
        case 20:
            showTheSpigotMagic()
        case 36:
            //MonteCarloPiConcurrent(12, num2)
            ConcurrentMCpi_second_menu(num2)
        case 39:
            showTheSpigotMagic()
        case 40:
            nifty_scoreBoard()
        case 60:
            ShowSLOC_behind_scoreBoard_40()
        case 41:
            BBPfConcurent()
        case 61:
            DisplayCodeBehind_BBPfConcurent()
        case 42:
            bbp_formula(num2)
        case 62:
            CodeRuneOf_bbp_formula()
        case 43:
            numerical_integration(num2)
        case 63:
            ShowLOC_behind_numerical_integration()
        case 44:
            Leibniz_method_one_billion_iters(num2)
        case 64:
            Show_Leibniz_method_one_billion_iters()
        case 45:

        case 65:

        case 46:
            // display contents of prior results file
            content, err := ioutil.ReadFile("dataFromConstants.txt")  // 
                if err != nil {   // if the file does not exist ... 
                    fmt.Println("no prior results -- no log file dataFromConstants.txt exists")
                } else {
                    fmt.Println(string(content))  // dump/display entire file to command line
                }
        case 66:
            var rune4above = [runeMark]            // display contents of prior results file
            content, err := ioutil.ReadFile("dataFromConstants.txt")  // 
                if err != nil {   // if the file does not exist ... 
                    fmt.Println("no prior results -- no log file dataFromConstants.txt exists")
                } else {
                    fmt.Println(string(content))  // dump/display entire file to command line
                }[runeMark]
                fmt.Println(rune4above)
        case 47:
            os.Exit(1)
        case 98:
            fmt.Println(rick, "\n")
        case 99:
            Explain_spigot()
} `
    fmt.Println(RicksSwitchRune)
}

func ConcurrentMCpi(num int) {
    fmt.Println("You have discovered 36, the MonteCarloPi method. 56 to see the code")
    fmt.Println("Warning, it is maximally concurent and will fully utilize all processors")
    fmt.Println("as such it has the potential to cook your cores if left running for too long")
                    numMC := 1
    fmt.Println("Enter an integer to specify a precision, and make it BIG, 9999999999 is reasonable")
    fmt.Scanf("%d", &numMC) 
    fmt.Println(MonteCarloPiConcurrent(numMC, num))
    //fmt.Println(MonteCarloPi(numMC, num))

}
func ConcurrentMCpi_second_menu(num2 int) {
    fmt.Println("You have discovered 36, the MonteCarloPi method. 56 to see the code")
    fmt.Println("Warning, it is maximally concurent and will fully utilize all processors")
    fmt.Println("as such it has the potential to cook your cores if left running for too long")
                    numMC := 1
    fmt.Println("Enter an integer to specify a precision, and make it BIG, 9999999999 is reasonable")
    fmt.Scanf("%d", &numMC) 
    fmt.Println(MonteCarloPiConcurrent(numMC, num2))
    //fmt.Println(MonteCarloPi(numMC, num))

}

func showTheMagicBehind_main_juuso(){
    var juuso_Rune = `

// Gauss–Legendre algorithm in Go: Because of how float64 is implemented in Go, this only gives 15 correct digits of Pi
// juuso@mit.edu , and further modified by rick.woolley@gmail.com 
// I am affraid that it would be far too much trouble to figure out all the math that is being implemented here ... 
// ... to do it one should begin by studying the Gauss–Legendre algorithm itself ?? 

var globalValue_in_juuso float64  // Rick's code 

func main_juuso() {  // initialize lists
    pin := []float64{}
    an  := []float64{1}
    pn  := []float64{1}

    tn := []float64{float64(1)/float64(4)}  // Rick's improved code, instead of the following ... 
    //tn := []float64{}                      // should have been able to accomplish the next line here as in initialize, see above
    //tn = append(tn, float64(1)/float64(4)) // append the quotient of 1.0/4.0 to list tn

    bn := []float64{float64(1)/math.Sqrt(2)}  // Rick's improved code, instead of the following ... 
    //bn := []float64{}                         // should have been able to accomplish the next line here as in initialize, as above
    //bn = append(bn, float64(1)/math.Sqrt(2))  // append the quotient of 1.0/sqrt of 2 to list bn

    // run the algorithm 4 times
    for i := 0; i < 4; i++ {  // call the 5 funcs (a,b,t,p, and pi) defined below, each of which returns just one []float64
        an = a(an, bn) 
        bn = b(an, bn)
        tn = t(an, bn, tn, pn) // calls func t and passes to it 4 []float64s 
        pn = p(pn)
        pin = pi(an, bn, tn, pin) // pin array ends up containing the 4 values that were calculated for Pi -- each time pi is called it returns a new version of a []float64 array
    }

    // print results
    // this for loop runs 4 times, therefore range pin yielded _ and 4 ???? There are no ';'s after the for, so, ??
    // ... no, 'range' looks to be a partner to the for itself 
    for _, value := range pin {  // skip initializing a counter? no!, not sure exactly all that 'range' does, but when appllied to pin fetches calculated Pi from pin and apparently also another return that is being tossed via _ ???
        // above 'range' seems to tell the for to "range" accross pin and assign a successive element to 'value', there were 4 elements, so it runs 4 times -- still not sure what the _, is doing here though ???
        fmt.Printf("pin is %.16f, and ... ", pin) // pin is an aray of calculated values for Pi  // Rick's code to discover same 
        fmt.Printf("%.16f Was calculated herewith\n", value)  // 'value' created on prior 'for' line and is set 4 times to a successive element of pin 
        globalValue_in_juuso = value // Rick's code to grab that final 'value' from last iteration 
    }
        fmt.Println(globalValue_in_juuso) // Rick's code
        fmt.Println("3.1415926535897932 <-- compared to the actual value of Pi")
        fmt.Println("1 23456789012345 so, 15 digits were calculated correctly")
        fmt.Println("   ... via the Gauss–Legendre algorithm ...")
} // end of main_juuso 

func a(an, bn []float64) []float64 {           // func a accepts an and bn of type []float64, and returns a []float64
    a := (an[len(an)-1] + bn[len(bn)-1]) / float64(2) // create local 'a' = (element of an indexed by len of an-1) + (element of bn indexed by len of bn-1) ?? 
    an = append(an, a)                               // append a to an
    return an
}
func b(an, bn []float64) []float64 {
    b := math.Sqrt(an[len(an)-2] * bn[len(bn)-1])  // create local 'b' = (sqrt of element an indexed by len of an-2) * (element of bn indexed by len of bn-1) ??
    bn = append(bn, b)
    return bn
}
func t(an, bn, tn, pn []float64) []float64 {  // accepts 4 parameters of type []float64, and returns a []float64
    t := tn[len(tn)-1] - pn[len(pn)-1]*math.Pow((an[len(an)-2]-an[len(an)-1]), 2)
    tn = append(tn, t)
    return tn
}
func p(pn []float64) []float64 {  // accepts one []float64 
    p := 2 * pn[len(pn)-1]        // create local p as 2 * the element of pn indexed by len(pn)-1
    pn = append(pn, p)
    return pn
}
func pi(an, bn, tn, pin []float64) []float64 {  // this func is all about appending just one value to pin ...
    pi := math.Pow((an[len(an)-1] + bn[len(bn)-1]), 2) / (4 * tn[len(tn)-1])  // ... and this is that value ... x^2 / (4 * an element from tn) ... x = two elements added together 
    //                                             ^2
    pin = append(pin, pi)
    return pin
}`
    fmt.Println(juuso_Rune)
}

var globalValue_in_juuso float64  // Rick's code 
var exterior_catcher int 

func main_juuso() {
    // initialize lists
    pin := []float64{}
    an  := []float64{1}
    pn  := []float64{1}

    tn := []float64{float64(1)/float64(4)}  // Rick's improved code, instead of the following ... 
    // ... 
    //tn := []float64{}                      // should have been able to accomplish the next line here as in initialize, see above
    //tn = append(tn, float64(1)/float64(4)) // append the quotient of 1.0/4.0 to list tn

    bn := []float64{float64(1)/math.Sqrt(2)}  // Rick's improved code, instead of the following ... 
    // ...
    //bn := []float64{}                         // should have been able to accomplish the next line here as in initialize, as above
    //bn = append(bn, float64(1)/math.Sqrt(2))  // append the quotient of 1.0/sqrt of 2 to list bn

    // run the algorithm 4 times
    for i := 0; i < 4; i++ {  // call the 5 funcs (a,b,t,p, and pi) defined below, each of which returns just one []float64
        an = a(an, bn) 
        bn = b(an, bn)
        tn = t(an, bn, tn, pn) // calls func t and passes to it 4 []float64s 
        pn = p(pn)
        pin = pi(an, bn, tn, pin) // pin array ends up containing the 4 values that were calculated for Pi 
        // each time pi is called it returns a new version of a []float64 array
    }

    // print results
    // this for loop runs 4 times, therefore range pin yielded _ and 4 ???? There are no ';'s after the for, so, ??
    // ... no, 'range' looks to be a partner to the for itself 
    // ran as ...
    //for exterior_catcher, value := range pin { // with the two test prints below 
    for _, value := range pin {  // skip initializing a counter? no!, not sure exactly all that 'range' does, but when appllied to pin fetches calculated Pi from pin and apparently also another return that is being tossed via _ 
        // above 'range' seems to tell the for to "range" accross pin and assign a successive element to 'value', there were 4 elements, so it runs 4 times -- the _ catches the unneeded return from 'range' which starts at 0 and goes to 3 in this loop 
        fmt.Printf("pin is %.16f, and ... ", pin) // pin is an aray of calculated values for Pi  // Rick's code to discover same 
        fmt.Printf("%.16f Was calculated herewith\n", value)  // 'value' created on prior 'for' line and is set 4 times to a successive element of pin 
        globalValue_in_juuso = value // Rick's code to grab that final 'value' from last iteration 
    //fmt.Printf("\n\nTop underscore is %d \n\n", exterior_catcher) it starts at 0 and goes to 3 

    }
        fmt.Println(globalValue_in_juuso) // Rick's code
        //fmt.Printf("\n\nBottom underscore is %d \n\n", exterior_catcher) this exterior_catcher var is never touched by the for loop 
        fmt.Println("3.1415926535897932 <-- compared to the actual value of Pi")
        fmt.Println("1 23456789012345 so, 15 digits were calculated correctly")
        fmt.Println("   ... via the Gauss–Legendre algorithm ...")
} // end of main_juuso 

func a(an, bn []float64) []float64 {           // func a accepts an and bn of type []float64, and returns a []float64
    a := (an[len(an)-1] + bn[len(bn)-1]) / float64(2) // create local 'a' = (element of an indexed by len of an-1) + (element of bn indexed by len of bn-1) ?? 
    an = append(an, a)                               // append a to an
    return an
}

func b(an, bn []float64) []float64 {
    b := math.Sqrt(an[len(an)-2] * bn[len(bn)-1])  // create local 'b' = (sqrt of element an indexed by len of an-2) * (element of bn indexed by len of bn-1) ??
    bn = append(bn, b)
    return bn
}

func t(an, bn, tn, pn []float64) []float64 {  // accepts 4 parameters of type []float64, and returns a []float64
    t := tn[len(tn)-1] - pn[len(pn)-1]*math.Pow((an[len(an)-2]-an[len(an)-1]), 2)
    tn = append(tn, t)
    return tn
}

func p(pn []float64) []float64 {  // accepts one []float64 
    p := 2 * pn[len(pn)-1]        // create local p as 2 * the element of pn indexed by len(pn)-1
    pn = append(pn, p)
    return pn
}

func pi(an, bn, tn, pin []float64) []float64 {  // this func is all about appending just one value to pin ...
    pi := math.Pow((an[len(an)-1] + bn[len(bn)-1]), 2) / (4 * tn[len(tn)-1])  // ... and this is that value ... x^2 / (4 * an element from tn) ... x = two elements added together 
    //                                             ^2
    pin = append(pin, pi)
    return pin
}

func BBPfConcurent(){
    //pi := PiMultiThread(4, 1024)
    pi := PiMultiThread(4, 512)
    fmt.Println("Calculated π vs actual:\n", pi, " <-- Calculated")
    fmt.Println(" 3.14159265358979323 <-- actual")
    fmt.Println(" 1 234567890123456 so, 15-16 digits were Calculated correctly")
    fmt.Println("   ... via the Bailey–Borwein–Plouffe formula [concurent] ...")
    fmt.Println("Difference between math.Pi and Calculated:", pi-math.Pi)
}
type Float64 struct {
    value float64
    lock  sync.RWMutex
}
func (f *Float64) Inc(diff float64) float64 {
    f.lock.Lock()
    ret := f.value + diff
    f.value = ret
    f.lock.Unlock()
    return ret
}
func (f *Float64) Get() float64 {
    f.lock.RLock()
    ret := f.value
    f.lock.RUnlock()
    return ret
}
func partialSum(kStart int, kOffset int, amount int) (sum float64) {
    for k := float64(kStart); k < float64(kStart+kOffset*amount); k += float64(kOffset) {
        sum += 1 / math.Pow(16, k) * (4/(8*k+1) - 2/(8*k+4) - 1/(8*k+5) - 1/(8*k+6))
    }

    return
}
func PiMultiThread(workers int, iteration int) float64 {
    wg := sync.WaitGroup{}
    ret := Float64{}

    for i := 0; i < workers; i++ {
        wg.Add(1)
        go (func(kStart int) {
            ret.Inc(partialSum(int(kStart), workers, iteration))
               //fmt.Println(ret)
            wg.Done()
            fmt.Println(ret) // Rick's progress report
        })(i)
    }
    wg.Wait()
    return ret.Get()
}

func DisplayCodeBehind_BBPfConcurent(){
    fmt.Println(DisplayCodeBehind_BBPfConcurentRune)
}
var DisplayCodeBehind_BBPfConcurentRune = `func BBPfConcurent(){
    //pi := PiMultiThread(4, 1024)
    pi := PiMultiThread(4, 512)
    fmt.Println("Calculated π vs actual:\n", pi, " <-- Calculated")
    fmt.Println(" 3.14159265358979323 <-- actual")
    fmt.Println(" 1 234567890123456 so, 15-16 digits were Calculated correctly")
    fmt.Println("   ... via the Bailey–Borwein–Plouffe formula [concurent] ...")
    fmt.Println("Difference between math.Pi and Calculated:", pi-math.Pi)
}
type Float64 struct {
    value float64
    lock  sync.RWMutex
}
func (f *Float64) Inc(diff float64) float64 {
    f.lock.Lock()
    ret := f.value + diff
    f.value = ret
    f.lock.Unlock()
    return ret
}
func (f *Float64) Get() float64 {
    f.lock.RLock()
    ret := f.value
    f.lock.RUnlock()
    return ret
}
func partialSum(kStart int, kOffset int, amount int) (sum float64) {
    for k := float64(kStart); k < float64(kStart+kOffset*amount); k += float64(kOffset) {
        sum += 1 / math.Pow(16, k) * (4/(8*k+1) - 2/(8*k+4) - 1/(8*k+5) - 1/(8*k+6))
    }

    return
}
func PiMultiThread(workers int, iteration int) float64 {
    wg := sync.WaitGroup{}
    ret := Float64{}

    for i := 0; i < workers; i++ {
        wg.Add(1)
        go (func(kStart int) {
            ret.Inc(partialSum(int(kStart), workers, iteration))
               //fmt.Println(ret)
            wg.Done()
            fmt.Println(ret) // Rick's progress report
        })(i)
    }
    wg.Wait()
    return ret.Get()
}`

// A concurrent computation of pi using Nilakantha's formula.
// by Diego Brener diegosilva13 on Github 
// ******* nifty scoreboard ***********************************
func nifty_scoreBoard (){
        // We use a ticker to specify the interval to update the values
        // on the scoreboard
        ticker := time.NewTicker(time.Millisecond * 108)

        // If the user wants to prematurely break out of the program by
        // issuing a Ctrl+C, we tell the signal package to notify us over
        // this interrupt channel
        interrupt := make(chan os.Signal, 1)
        signal.Notify(interrupt, os.Interrupt)

        go pi_nf(5000)

        // This anonymous function is responsible for updating the scoreboard
        // as per the interval specified by the ticker
        go func() {
            for range ticker.C {
                printCalculationSummary()
            }

        }()

        for {
            select {

            // Handle the case when the computation has ended, we can
            // end the program (exit out of this infinite loop)
            case <-computationDone:
                ticker.Stop()
                fmt.Println("Program done calculating Pi.")
                os.Exit(0)

            // If the user interrupts the program (Ctrl+C) we end the
            // program (exit out of this infinite loop)
            case <-interrupt:
                ticker.Stop()
                fmt.Println("Program interrupted by the user.")
                return
            }
        }
}
        // We want to create a virtual scoreboard where we can simutaneously
        // show the current value of Pi and how many Nilakantha terms we
        // have calculated.

        // We can create the virtual scoreboard by using some ANSI escape codes
        // Here's a wikipedia article giving you the complete list of ANSI escape
        // codes: https://en.wikipedia.org/wiki/ANSI_escape_code
        const ANSIClearScreenSequence = "\033[H\033[2J"
        const ANSIFirstSlotScreenSequence = "\033[2;0H"
        const ANSISecondSlotScreenSequence = "\033[3;0H"

        // The channel used to update the current value of pi on the scoreboard
        var pichan chan float64 = make(chan float64)

        // The channel that we use to indicate that the program can exit
        var computationDone chan bool = make(chan bool, 1)

        // Number of Nilakantha terms for the scoreboard
        var termsCount int

        // This function serves as our virtual scoreboard to show the current
        // computed value of Pi using Nilakantha's formula
        func printCalculationSummary() {

            fmt.Print(ANSIClearScreenSequence)
            fmt.Println(ANSIFirstSlotScreenSequence, "Computed Value of Pi:\t\t", <-pichan)
            fmt.Println(ANSISecondSlotScreenSequence, "# of Nilakantha Terms:\t\t", termsCount)
        }

        // We are going to use Nilakantha's formula from the Tantrasamgraha (which
        // is more than 500 years old) to compute the value of Pi to several decimal
        // points
        func pi_nf(n int) float64 { 
            ch := make(chan float64)
            // The k variable is considered to be the current step
            for k := 1; k <= n; k++ {
                // Each Nilakantha term is calculated in its own goroutine
                go nilakanthaTerm(ch, float64(k))
            }
            f := 3.0

            // We sum up the calculated Nilakantha terms for n steps
            for k := 1; k <= n; k++ {
                termsCount++
                f += <-ch
                pichan <- f
            }

            // We notify that the computation is done over the channel
            computationDone <- true
            return f
        }

        // This function gives us the nilakanthaTerm for the kth step
        func nilakanthaTerm(ch chan float64, k float64) {
            j := 2 * k
            if int64(k)%2 == 1 {
                ch <- 4.0 / (j * (j + 1) * (j + 2))
            } else {
                ch <- -4.0 / (j * (j + 1) * (j + 2))
            }
        }

func ShowSLOC_behind_scoreBoard_40 () {
    fmt.Print(rune_of_scoreBoard)
}

var rune_of_scoreBoard = `// A concurrent computation of pi using Nilakantha's formula.
// by Diego Brener diegosilva13 on Github 
// ******* nifty scoreboard ***********************************
func nifty_scoreBoard (){
        // We use a ticker to specify the interval to update the values
        // on the scoreboard
        ticker := time.NewTicker(time.Millisecond * 108)

        // If the user wants to prematurely break out of the program by
        // issuing a Ctrl+C, we tell the signal package to notify us over
        // this interrupt channel
        interrupt := make(chan os.Signal, 1)
        signal.Notify(interrupt, os.Interrupt)

        go pi_nf(5000)

        // This anonymous function is responsible for updating the scoreboard
        // as per the interval specified by the ticker
        go func() {
            for range ticker.C {
                printCalculationSummary()
            }

        }()

        for {
            select {

            // Handle the case when the computation has ended, we can
            // end the program (exit out of this infinite loop)
            case <-computationDone:
                ticker.Stop()
                fmt.Println("Program done calculating Pi.")
                os.Exit(0)

            // If the user interrupts the program (Ctrl+C) we end the
            // program (exit out of this infinite loop)
            case <-interrupt:
                ticker.Stop()
                fmt.Println("Program interrupted by the user.")
                return
            }
        }
}
        // We want to create a virtual scoreboard where we can simutaneously
        // show the current value of Pi and how many Nilakantha terms we
        // have calculated.

        // We can create the virtual scoreboard by using some ANSI escape codes
        // Here's a wikipedia article giving you the complete list of ANSI escape
        // codes: https://en.wikipedia.org/wiki/ANSI_escape_code
        const ANSIClearScreenSequence = "\033[H\033[2J"
        const ANSIFirstSlotScreenSequence = "\033[2;0H"
        const ANSISecondSlotScreenSequence = "\033[3;0H"

        // The channel used to update the current value of pi on the scoreboard
        var pichan chan float64 = make(chan float64)

        // The channel that we use to indicate that the program can exit
        var computationDone chan bool = make(chan bool, 1)

        // Number of Nilakantha terms for the scoreboard
        var termsCount int

        // This function serves as our virtual scoreboard to show the current
        // computed value of Pi using Nilakantha's formula
        func printCalculationSummary() {

            fmt.Print(ANSIClearScreenSequence)
            fmt.Println(ANSIFirstSlotScreenSequence, "Computed Value of Pi:\t\t", <-pichan)
            fmt.Println(ANSISecondSlotScreenSequence, "# of Nilakantha Terms:\t\t", termsCount)
        }

        // We are going to use Nilakantha's formula from the Tantrasamgraha (which
        // is more than 500 years old) to compute the value of Pi to several decimal
        // points
        func pi_nf(n int) float64 { 
            ch := make(chan float64)
            // The k variable is considered to be the current step
            for k := 1; k <= n; k++ {
                // Each Nilakantha term is calculated in its own goroutine
                go nilakanthaTerm(ch, float64(k))
            }
            f := 3.0

            // We sum up the calculated Nilakantha terms for n steps
            for k := 1; k <= n; k++ {
                termsCount++
                f += <-ch
                pichan <- f
            }

            // We notify that the computation is done over the channel
            computationDone <- true
            return f
        }

        // This function gives us the nilakanthaTerm for the kth step
        func nilakanthaTerm(ch chan float64, k float64) {
            j := 2 * k
            if int64(k)%2 == 1 {
                ch <- 4.0 / (j * (j + 1) * (j + 2))
            } else {
                ch <- -4.0 / (j * (j + 1) * (j + 2))
            }
        }`

func numerical_integration (num2 int){
        /* Testpad here: http://tour.golang.org/basics/13

        Mods by rick.woolley@gmail.com 

        calculate pi using a numerical integration of 

                   1                                                                   
           y = ---------                                                               
                1 + x²                                                                

        The indefinite integral of (1 + x²)⁻¹ wrt x is tan⁻¹ x (+ constant)

        OIOW:

        \int\frac{1}{1+x^2}\mathrm{d} x = \tan^{-1} x + c                             

        Evaluating that integral between x = 0 and x = 1 gives us π/4

         \int^{1}_0\frac{1}{1+x^2}\mathrm{d} x = \frac{\pi}{4}                         

        So if we add the area of sufficiently many y-high rectangles between 0 and 1
        for y = (1+x²)⁻¹ , and multiply by 4, we should get a decent approximation for π.

        */

    var i int
    var x float64

    var num_steps int
    var err error

    fmt.Println("\nWhen no command line arg is given, it defaults to: 1,000,000,000\n")

    if len(os.Args) > 1 {
        num_steps, err = strconv.Atoi(os.Args[1])
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error parsing argument as step count: %s\n", err)
            os.Exit(1)
        }
    } else {
        num_steps = 1000000000
    }

    fmt.Printf("Calculating PI using:\n  %v slices\n  1 process\n\n", num_steps)

    var sum float64 = 0
    step := 1.0 / float64(num_steps)

    var start time.Time = time.Now()

    for i = 0; i < num_steps; i++ {
        x = (float64(i) + 0.5) * step
        sum += 4.0 / (1.0 + x*x)
    }

    var stop time.Time = time.Now()

    var pi = sum * step

    timeTaken := (float64(stop.Sub(start)) * 0.000000001) // Durations come in ns

    //fmt.Printf("Obtained value for PI: \n%.16g\n Time taken: %v s\n", pi, timeTaken)
    fmt.Println(pi)
    fmt.Println("3.141592653589793 <-- actual")
    fmt.Println("1 234567890123 so, up to 13 digits were Calculated correctly\n")

    //fmt.Println(timeTaken, "seconds runtime")
    fmt.Printf("     %.9g seconds runtime\n\n", timeTaken)

    t := time.Now()
    elapsed := t.Sub(start)
    //fmt.Println("Ten Billion iterations were completed in ", elapsed, " yielding 10 digits of π\n") 

// log stats to a log file 
    var LinesPerSecondInt int 
    LinesPerIter := 2 // an estimate 
        LinesPerSecondInt = (int(LinesPerIter) * int(i) ) / int(elapsed.Seconds()) // .Seconds() returns a float64
        fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
            check(err1)                                // ... gets a file handle to dataFromConstants.txt
            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
        Hostname, _ := os.Hostname()
        _ , err0 := fmt.Fprintf(fileHandle, "\n  -- calculate pi using a numerical integration -- selection #%d on %s \n", num2, Hostname)
            check(err0)
        current_time := time.Now()
        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
        check(err6)
        _ , err2 := fmt.Fprintf(fileHandle, "%d was Lines/Second \n", LinesPerSecondInt) 
            check(err2)
        _ , err4 := fmt.Fprintf(fileHandle, "%.02f was Iterations/Seconds \n", float64(i)/float64(elapsed.Seconds()))
            check(err4)
        _ , err5 := fmt.Fprintf(fileHandle, "%d was total Iterations \n", i)
            check(err5)
        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
            check(err7)
}
func ShowLOC_behind_numerical_integration (){
    fmt.Println(NI_rune)
}
var NI_rune = `func numerical_integration (num2 int){
        /* Testpad here: http://tour.golang.org/basics/13

        Mods by rick.woolley@gmail.com 

        calculate pi using a numerical integration of 

                   1                                                                   
           y = ---------                                                               
                1 + x²                                                                

        The indefinite integral of (1 + x²)⁻¹ wrt x is tan⁻¹ x (+ constant)

        OIOW:

        \int\frac{1}{1+x^2}\mathrm{d} x = \tan^{-1} x + c                             

        Evaluating that integral between x = 0 and x = 1 gives us π/4

         \int^{1}_0\frac{1}{1+x^2}\mathrm{d} x = \frac{\pi}{4}                         

        So if we add the area of sufficiently many y-high rectangles between 0 and 1
        for y = (1+x²)⁻¹ , and multiply by 4, we should get a decent approximation for π.

        */

    var i int
    var x float64

    var num_steps int
    var err error

    fmt.Println("\nWhen no command line arg is given, it defaults to: 1,000,000,000\n")

    if len(os.Args) > 1 {
        num_steps, err = strconv.Atoi(os.Args[1])
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error parsing argument as step count: %s\n", err)
            os.Exit(1)
        }
    } else {
        num_steps = 1000000000
    }

    fmt.Printf("Calculating PI using:\n  %v slices\n  1 process\n\n", num_steps)

    var sum float64 = 0
    step := 1.0 / float64(num_steps)

    var start time.Time = time.Now()

    for i = 0; i < num_steps; i++ {
        x = (float64(i) + 0.5) * step
        sum += 4.0 / (1.0 + x*x)
    }

    var stop time.Time = time.Now()

    var pi = sum * step

    timeTaken := (float64(stop.Sub(start)) * 0.000000001) // Durations come in ns

    //fmt.Printf("Obtained value for PI: \n%.16g\n Time taken: %v s\n", pi, timeTaken)
    fmt.Println(pi)
    fmt.Println("3.141592653589793 <-- actual")
    fmt.Println("1 234567890123 so, up to 13 digits were Calculated correctly\n")

    //fmt.Println(timeTaken, "seconds runtime")
    fmt.Printf("     %.9g seconds runtime\n\n", timeTaken)

    t := time.Now()
    elapsed := t.Sub(start)
    //fmt.Println("Ten Billion iterations were completed in ", elapsed, " yielding 10 digits of π\n") 

// log stats to a log file 
    var LinesPerSecondInt int 
    LinesPerIter := 2 // an estimate 
        LinesPerSecondInt = (int(LinesPerIter) * int(i) ) / int(elapsed.Seconds()) // .Seconds() returns a float64
        fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
            check(err1)                                // ... gets a file handle to dataFromConstants.txt
            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
        Hostname, _ := os.Hostname()
        _ , err0 := fmt.Fprintf(fileHandle, "\n  -- calculate pi using a numerical integration -- selection #%d on %s \n", num2, Hostname)
            check(err0)
        current_time := time.Now()
        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
        check(err6)
        _ , err2 := fmt.Fprintf(fileHandle, "%d was Lines/Second \n", LinesPerSecondInt) 
            check(err2)
        _ , err4 := fmt.Fprintf(fileHandle, "%.02f was Iterations/Seconds \n", float64(i)/float64(elapsed.Seconds()))
            check(err4)
        _ , err5 := fmt.Fprintf(fileHandle, "%d was total Iterations \n", i)
            check(err5)
        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
            check(err7)
}`

func bbp_formula (num2 int){
// fmt.Println(num2)
iters_bbp := 1 
    start := time.Now()
    numCPU := runtime.NumCPU()
    runtime.GOMAXPROCS(numCPU)

/*
    if len(os.Args) != 2 {
        //fmt.Println("pls enter how many digits of Pi you want, as an arg on the command line after the go file name")
        os.Exit(1)
    }

    n, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
*/

    n:= 46
    p := uint((int(math.Log2(10)))*n + 3)

    result := make(chan *big.Float, n)
    worker := workers(p)

    pi := new(big.Float).SetPrec(p).SetInt64(0)

    for i := 0; i < n; i++ {
        go worker(i, result)
        iters_bbp = i
    }

    for i := 0; i < n; i++ {
        pi.Add(pi, <-result)
        iters_bbp = i
    }

    dur := time.Since(start)
    fmt.Printf("took %v to calculate %d digits of pi \n", dur, n)
    fmt.Printf("%[1]*.[2]*[3]f \n", 1, n, pi)

// log run stats to a log file
    t := time.Now()
    elapsed := t.Sub(start)
        fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
            check(err1)                                // ... gets a file handle to dataFromConstants.txt
            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
        Hostname, _ := os.Hostname()
        _ , err0 := fmt.Fprintf(fileHandle, "\n  -- calculate pi using the bbp formula -- selection #%d on %s \n", num2, Hostname)
            check(err0)
        current_time := time.Now()
        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
        check(err6)
        _ , err4 := fmt.Fprintf(fileHandle, "%.02f was Iterations/Seconds \n", float64(iters_bbp)/float64(elapsed.Seconds()))
            check(err4)
        _ , err5 := fmt.Fprintf(fileHandle, "%d was total Iterations \n", iters_bbp)
            check(err5)
        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
            check(err7)
            
}


func workers(p uint) func(id int, result chan *big.Float) {
    B1 := new(big.Float).SetPrec(p).SetInt64(1)
    B2 := new(big.Float).SetPrec(p).SetInt64(2)
    B4 := new(big.Float).SetPrec(p).SetInt64(4)
    B5 := new(big.Float).SetPrec(p).SetInt64(5)
    B6 := new(big.Float).SetPrec(p).SetInt64(6)
    B8 := new(big.Float).SetPrec(p).SetInt64(8)
    B16 := new(big.Float).SetPrec(p).SetInt64(16)

    return func(id int, result chan *big.Float) {
        Bn := new(big.Float).SetPrec(p).SetInt64(int64(id))

        C1 := new(big.Float).SetPrec(p).SetInt64(1)
        for i := 0; i < id; i++ {
            C1.Mul(C1, B16)
        }

        C2 := new(big.Float).SetPrec(p)
        C2.Mul(B8, Bn)

        T1 := new(big.Float).SetPrec(p)
        T1.Add(C2, B1)
        T1.Quo(B4, T1)

        T2 := new(big.Float).SetPrec(p)
        T2.Add(C2, B4)
        T2.Quo(B2, T2)

        T3 := new(big.Float).SetPrec(p)
        T3.Add(C2, B5)
        T3.Quo(B1, T3)

        T4 := new(big.Float).SetPrec(p)
        T4.Add(C2, B6)
        T4.Quo(B1, T4)

        R := new(big.Float).SetPrec(p)
        R.Sub(T1, T2)
        R.Sub(R, T3)
        R.Sub(R, T4)
        R.Quo(R, C1)

        result <- R
    }
}
func CodeRuneOf_bbp_formula(){
    fmt.Println(RuneOf_bbp_formula)
}
var RuneOf_bbp_formula = `
func bbp_formula (num2 int){
// fmt.Println(num2)
iters_bbp := 1 
    start := time.Now()
    numCPU := runtime.NumCPU()
    runtime.GOMAXPROCS(numCPU)

/*
    if len(os.Args) != 2 {
        //fmt.Println("pls enter how many digits of Pi you want, as an arg on the command line after the go file name")
        os.Exit(1)
    }

    n, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
*/

    n:= 46
    p := uint((int(math.Log2(10)))*n + 3)

    result := make(chan *big.Float, n)
    worker := workers(p)

    pi := new(big.Float).SetPrec(p).SetInt64(0)

    for i := 0; i < n; i++ {
        go worker(i, result)
        iters_bbp = i
    }

    for i := 0; i < n; i++ {
        pi.Add(pi, <-result)
        iters_bbp = i
    }

    dur := time.Since(start)
    fmt.Printf("took %v to calculate %d digits of pi \n", dur, n)
    fmt.Printf("%[1]*.[2]*[3]f \n", 1, n, pi)

// log run stats to a log file
    t := time.Now()
    elapsed := t.Sub(start)
        fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
            check(err1)                                // ... gets a file handle to dataFromConstants.txt
            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
        Hostname, _ := os.Hostname()
        _ , err0 := fmt.Fprintf(fileHandle, "\n  -- calculate pi using the bbp formula -- selection #%d on %s \n", num2, Hostname)
            check(err0)
        current_time := time.Now()
        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
        check(err6)
        _ , err4 := fmt.Fprintf(fileHandle, "%.02f was Iterations/Seconds \n", float64(iters_bbp)/float64(elapsed.Seconds()))
            check(err4)
        _ , err5 := fmt.Fprintf(fileHandle, "%d was total Iterations \n", iters_bbp)
            check(err5)
        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
            check(err7)
            
}

func workers(p uint) func(id int, result chan *big.Float) {
    B1 := new(big.Float).SetPrec(p).SetInt64(1)
    B2 := new(big.Float).SetPrec(p).SetInt64(2)
    B4 := new(big.Float).SetPrec(p).SetInt64(4)
    B5 := new(big.Float).SetPrec(p).SetInt64(5)
    B6 := new(big.Float).SetPrec(p).SetInt64(6)
    B8 := new(big.Float).SetPrec(p).SetInt64(8)
    B16 := new(big.Float).SetPrec(p).SetInt64(16)

    return func(id int, result chan *big.Float) {
        Bn := new(big.Float).SetPrec(p).SetInt64(int64(id))

        C1 := new(big.Float).SetPrec(p).SetInt64(1)
        for i := 0; i < id; i++ {
            C1.Mul(C1, B16)
        }

        C2 := new(big.Float).SetPrec(p)
        C2.Mul(B8, Bn)

        T1 := new(big.Float).SetPrec(p)
        T1.Add(C2, B1)
        T1.Quo(B4, T1)

        T2 := new(big.Float).SetPrec(p)
        T2.Add(C2, B4)
        T2.Quo(B2, T2)

        T3 := new(big.Float).SetPrec(p)
        T3.Add(C2, B5)
        T3.Quo(B1, T3)

        T4 := new(big.Float).SetPrec(p)
        T4.Add(C2, B6)
        T4.Quo(B1, T4)

        R := new(big.Float).SetPrec(p)
        R.Sub(T1, T2)
        R.Sub(R, T3)
        R.Sub(R, T4)
        R.Quo(R, C1)

        result <- R
    }
}`


func Leibniz_method_one_billion_iters(num2 int){
        start := time.Now() // mine, option 44 from second menu
        fmt.Println("\n you selected option 44, which is going to run for 'a minute' \n")
    /**
     *   Adrian Statescu <mergesortv@gmail.com> (http://adrianstatescu.com) 
     *
     *   Approximation of PI using Leibniz method.
     *   1-1/3+1/5-1/7
     *
     *   Compile and Run -> go run leibniz_pi.go
     *
     *   License: MIT
     */ 

     // ******* RUNS TEN BIL ITERS SO IT RUNS A LONG TIME ***************************************

    //var t1, t2, r, sign, i float64
        var t1, t2, sign, i float64

     
         t1 = 1.0

         t2 = 1.0 - float64(1.0/3.0)

    //r = t1 - t2; 

         i = 5;

    //eps := 0.00001

         sign = 1

         iters := 0 

    //for float64(4 * r) >= eps {
          for iters < 10000000000 {

              t1 = t2

              t2 += float64( 1.0 / i) * sign

              if t1 > t2 {
                 
    //r = float64(t1 - t2) 

              }  else {

    //r = float64(t2 - t1)
              }

              sign = (-1)*sign

              i += 2

              iters++ 
         }         
    //fmt.Println("%x", 4 * t2)           

    // mine follows
        pi := 4*t2
        fmt.Println("pi is:\n", pi)
        fmt.Println(" 3.141592653589793 <-- actual")
        fmt.Println(" 1 234567890123 so, 10 digits were Calculated correctly after 10,000,000,000 iterations\n")
    t := time.Now()
    elapsed := t.Sub(start)
    fmt.Println("Ten Billion iterations were completed in ", elapsed, " yielding 10 digits of π\n") 

// log stats to a log file 
    var LinesPerSecondInt int 
    LinesPerIter := 6 // an estimate 
        LinesPerSecondInt = (int(LinesPerIter) * int(iters) ) / int(elapsed.Seconds()) // .Seconds() returns a float64
        fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
            check(err1)                                // ... gets a file handle to dataFromConstants.txt
            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
        Hostname, _ := os.Hostname()
        _ , err0 := fmt.Fprintf(fileHandle, "\n  -- Approximation of PI using Leibniz method -- selection #%d on %s \n", num2, Hostname)
            check(err0)
        current_time := time.Now()
        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
        check(err6)
        _ , err2 := fmt.Fprintf(fileHandle, "%d was Lines/Second \n", LinesPerSecondInt) 
            check(err2)
        _ , err4 := fmt.Fprintf(fileHandle, "%.02f was Iterations/Seconds \n", float64(iters)/float64(elapsed.Seconds()))
            check(err4)
        _ , err5 := fmt.Fprintf(fileHandle, "%d was total Iterations \n", iters)
            check(err5)
        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
            check(err7)
}
func Show_Leibniz_method_one_billion_iters(){
    fmt.Println(Leibniz_method_one_billion_itersRune)
}
var Leibniz_method_one_billion_itersRune = `func Leibniz_method_one_billion_iters(num2 int){
        start := time.Now() // mine, option 44 from second menu
        fmt.Println("\n you selected option 44, which is going to run for 'a minute' \n")
    /**
     *   Adrian Statescu <mergesortv@gmail.com> (http://adrianstatescu.com) 
     *
     *   Approximation of PI using Leibniz method.
     *   1-1/3+1/5-1/7
     *
     *   Compile and Run -> go run leibniz_pi.go
     *
     *   License: MIT
     */ 

     // ******* RUNS TEN BIL ITERS SO IT RUNS A LONG TIME ***************************************

    //var t1, t2, r, sign, i float64
        var t1, t2, sign, i float64

     
         t1 = 1.0

         t2 = 1.0 - float64(1.0/3.0)

    //r = t1 - t2; 

         i = 5;

    //eps := 0.00001

         sign = 1

         iters := 0 

    //for float64(4 * r) >= eps {
          for iters < 10000000000 {

              t1 = t2

              t2 += float64( 1.0 / i) * sign

              if t1 > t2 {
                 
    //r = float64(t1 - t2) 

              }  else {

    //r = float64(t2 - t1)
              }

              sign = (-1)*sign

              i += 2

              iters++ 
         }         
    //fmt.Println("%x", 4 * t2)           

    // mine follows
        pi := 4*t2
        fmt.Println("pi is:\n", pi)
        fmt.Println(" 3.141592653589793 <-- actual")
        fmt.Println(" 1 234567890123 so, 10 digits were Calculated correctly after 10,000,000,000 iterations\n")
    t := time.Now()
    elapsed := t.Sub(start)
    fmt.Println("Ten Billion iterations were completed in ", elapsed, " yielding 10 digits of π\n") 

// log stats to a log file 
    var LinesPerSecondInt int 
    LinesPerIter := 6 // an estimate 
        LinesPerSecondInt = (int(LinesPerIter) * int(iters) ) / int(elapsed.Seconds()) // .Seconds() returns a float64
        fileHandle, err1 := os.OpenFile("dataFromConstants.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
            check(err1)                                // ... gets a file handle to dataFromConstants.txt
            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
        Hostname, _ := os.Hostname()
        _ , err0 := fmt.Fprintf(fileHandle, "\n  -- Approximation of PI using Leibniz method -- selection #%d on %s \n", num2, Hostname)
            check(err0)
        current_time := time.Now()
        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
        check(err6)
        _ , err2 := fmt.Fprintf(fileHandle, "%d was Lines/Second \n", LinesPerSecondInt) 
            check(err2)
        _ , err4 := fmt.Fprintf(fileHandle, "%.02f was Iterations/Seconds \n", float64(iters)/float64(elapsed.Seconds()))
            check(err4)
        _ , err5 := fmt.Fprintf(fileHandle, "%d was total Iterations \n", iters)
            check(err5)
        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
            check(err7)
}`


func RicksSwitch() {    // the primary if-then-else routine to execute a selection from the menu
    var num int 
        //fmt.Print("Enter your selection from above, 1 - 12 \u00a9 2022, by Richard Hart Woolley [it's an echo]\n")  // a kluge
        // the above kluge may be needed in some environments (windows) along with the duplicate in RicksDisplayMenuPi()
        fmt.Print("Enter your selection, 1 - 19 (IS THIS WINDOW MAXIMIZED?  Well, do it!)\n")
        fmt.Scanf("%d", &num)  // pause and request input from the user
        if num == 0 {
            fmt.Println("\nYou failed to make a selection, Hit Enter/Return to redisplay the menu, Ctrl-C to End/Exit")
        }
    if num > 100 && num < 10000 { num = 17 }  // to display a funny out-of-range message as case 17:
    switch num { 
        case 1:  // Calculate the square root of 3 from first principles of geometry
            squareRootOf3(num)
        case 21: 
            showMagicBehindsquareRootOf3()
        case 2:  // CASE 2 (best method) the Bailey–Borwein–Plouffe formula for π, circa 1995 
            BBPF(num)
        case 22:
            showMagicBehindBBPF() 
        case 3: // (worst method)
            WorstMethod(num)
        case 23: 
            showMagicBehindWorstMethod() 
        case 4:  // An improved version of Archimedes' method
            Archimedes(num)
        case 24:
            showMagicBehindArchimedes()
        case 5: // a series by Indian astronomer Nilakantha Somayaji circa 1500 AD 
            Nilakantha(num)
        case 25:
            showMagicBehindNilakantha()
        case 6: // Gottfried Wilhelm Leibniz
            GottfriedWilhelmLeibniz(num)
        case 26:
            showMagicBehindGottfriedWilhelmLeibniz()
        case 7: // the Gregory-Leibniz series
            GregoryLeibniz(num)
        case 27:
            showMagicBehindGregoryLeibniz()
        case 8: // John Wallis circa 1655 --- runs very long 
            JohnWallis(num)
        case 28: 
            showMagicBehindJohnWallis()
        case 9: // Euler's Number
            EulersNumber(num)
        case 29:
            showMagicBehindEulersNumber()
        case 10: // Erdős–Borwein constant
            ErdosBorwein(num)
        case 30:
            showMagicBehindErdosBorwein()
        case 11:  // display a review of the derivation of the Pythagorean
            DisplayPythagorean(num)
        case 31: 
            DisplayPythagorean(num)
        case 12: // display contents of prior results file
            content, err := ioutil.ReadFile("dataFromConstants.txt")  // 
                if err != nil {   // if the file does not exist ... 
                    fmt.Println("no prior results -- no log file dataFromConstants.txt exists\n")
                } else {
                    fmt.Println(string(content))  // dump/display entire file to command line
                }
        case 32:
            var fileAccessRune = `
    content, err := ioutil.ReadFile("dataFromConstants.txt")  // 
        if err != nil {   // if the file does not exist ... 
            fmt.Println("no prior results -- no log file dataFromConstants.txt exists")
        } else {
            fmt.Println(string(content))  // dump/display entire file to command line
        }`
            fmt.Println(fileAccessRune)
        case 33: 
            showMagicBehindmainFunc()
        case 13: 
            fmt.Println(" ... So sorry, but", num, "was not an option")
        case 14: 
            fmt.Println(num, " was not an option! It was not on the menu, go fish!\n")
        case 15: 
            fmt.Println("Your selection of", num, " is right-out!  Go Fish!\n")
        case 16: 
            fmt.Println("Your selection is really-far-out!  Go Fish!\n")
        case 17: 
            fmt.Println("\nOops, how'd we get here? Hit Enter/Return again to possibly redisplay the menu")
        case 18:
            secondMenu()
        case 34:
            /* // a project that I failed at
            var precision int 
            fmt.Println("what shall be thy precision? ")
            fmt.Scanf("%d", &precision) 
            fmt.Println("result is ", calculatePispi(precision))
            */
        case 35:
            showMagicBehindSwitch()
        case 36:
            ConcurrentMCpi(num)
        case 56:
            showMagicBehindMonteCarloPiConcurrent()
        case 37:
            main_juuso()
        case 57:
            showTheMagicBehind_main_juuso()
        case 58:
            os.Exit(1)
        case 98:
            fmt.Println(rick, "\n")
        case 99:
            Explain_spigot()
    default: 
        fmt.Println("this is the switch default code, after a break option ??")
    } 
} 

func RicksDisplayMenuPi() {
fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nEnter 33 to see the the magic behind main (selector+20 for all others)") 
fmt.Println("... 35 to see the code for the switches, or poke around for easter eggs\n\n")
// Veritassium https://youtu.be/gMlf1ELvRzc?list=LL  was the initial inspiration for all of this ...
// ... option 1 below having been discussed on his channel 
fmt.Println("1:  Calculate the Square Root of 3 (\u221A3) from first-principles")
fmt.Println("    ... i.e., from a ratio of 3:1 of perfect squares\n")
fmt.Println("2:  Pi: Bailey–Borwein–Plouffe formula for π, discovered in 1995 (best method)\n")
fmt.Println("3:  Pi: (worst method) requires calculating the \u221A3 (see selection 1)")
fmt.Println("     ... only yeilds 4 digits of π \n")
fmt.Println("4:  Pi: Archimedes' method (improved) of bisecting triangles, circa 200 BC")
fmt.Println("    π = begining with a hexagon, iteratively double its number of sides") 
fmt.Println("      26 iterations is all it will take to get the ... ")
fmt.Println("      16 digits of π -- v.fast   (ENTER '11' for Pythagorean theorem)\n")
fmt.Println("5:  Pi: An infinite series by Indian astronomer Nilakantha Somayaji circa 1500 ")
fmt.Println("    ... this equation can be found in an ancient Sanskrit verse")
fmt.Println("    π = 3 + 4/(2*3*4) - 4/(4*5*6) + 4/(6*7*8) - 4/(8*9*10) + 4/(10*11*12) ...")
fmt.Println("      One-Hundred-Million iterations will be executed -- in less than a second")
fmt.Println("      14 digits of π -- VERY FAST -- gives 7 digits in just 100 iterations !!!\n")
fmt.Println("6:  Pi: Gottfried Wilhelm Leibniz, late 17th century ( and by Isaac Newton )")
fmt.Println("    π = 4 * ( 1 - 1/3 + 1/5 - 1/7 + 1/9 ...")
fmt.Println("      4 Billion iterations will be executed")
fmt.Println("      10 digits of π -- 8 digits in 100M iterations in under a minute \n")
fmt.Println("7:  Pi: The Gregory-Leibniz series circa 1676")
fmt.Println("    π = (4/1) - (4/3) + (4/5) - (4/7) + (4/9) - (4/11) + (4/13) - (4/15) ...")
fmt.Println("      Three-Hundred-Million iterations will be executed")
fmt.Println("      9 digits of π -- in seconds\n")
fmt.Println("8:  Pi: An infinite series by John Wallis circa 1655")
fmt.Println("    π = 2 * ((2/1)*(2/3)) * ((4/3)*(4/5)) * ((6/5)*(6/7)) ... ")
fmt.Println("      One-Billion iterations will be executed; option for 40 billion iterations")
fmt.Println("      9 digits of π -- a billion loops, in seconds -- option for 10 digits\n")
fmt.Println("9:   Calculate Euler's Number: \u2107 or \u2147 the natural logarithmic base")
fmt.Println("        Explore the limit of interest\n")
fmt.Println("10:  Calculate the Erdős-Borwein constant from a breif infinite series\n")
fmt.Println("11:  Show a review of the derivation of the Pythagorean\n")
fmt.Println("12:  Display prior execution times from longer-running prior selections \n")
fmt.Println("18:  FOR SECOND MENU\n")
//fmt.Print("Enter your selection from above, 1 - 18 \u00a9 2022, by Richard Hart Woolley\n")
// the above kluge is definately not needed in a Linux environment
fmt.Println("Ctrl-C or 58 to End/Exit  SLOC = 4700ish   \u00a9 2023, by Richard Hart Woolley \n")
}

func secondMenu(){
    fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nSECOND MENU: (add 20 to any selection to show the Go magic)\n") 
    fmt.Println("19:  Pi: Open the 'Spigot' algorithm, instantly bakes way too much pie\n")
    fmt.Println("36:  Pi: Concurrent Monte_Carlo_method\n")
    fmt.Println("40:  Pi: Nifty 'ScoreBoard' using Nilakantha's formula (Ctrl-C to exit it)\n")
    fmt.Println("41:  Pi: Bailey–Borwein–Plouffe formula [concurent]\n")
    fmt.Println("42:  Pi: bbp formula to 46 digits\n")
    fmt.Println("43:  Pi: via Numerical Integration \n")
    fmt.Println("44:  Pi: via Leibniz method in one billion iterations [runs a while]\n")
    fmt.Println("45:  Pi: MonteCarloPi (non-concurrent) [9999999999 as input]\n")
    fmt.Println("99:  Pi: via BBP and spigot explained\n")
    fmt.Println("46:  Display prior execution times from longer-running prior selections\n\n")
    //fmt.Print("Enter your selection from above, 19 - 46 \u00a9 2022, by Richard Hart Woolley\n")
    // the above kluge is definately not needed in a Linux environment
    fmt.Println("47: to End/Exit  SLOC = 4700ish   \u00a9 2023, by Richard Hart Woolley \n")
    var num2 int 
    fmt.Scanf("%d", &num2)  // pause and request input from the user
        if num2 == 0 {
            fmt.Println("\nYou failed to make a selection, Hit Enter/Return to redisplay the menu, Ctrl-C to End/Exit")
        }
    if num2 > 100 && num2 < 10000 { num2 = 17 }  // to display a funny out-of-range message as case 17:
    switch num2 { 
        case 19: 
            TheSpigot()
        case 20:
            showTheSpigotMagic()
        case 36:
            //MonteCarloPiConcurrent(12, num2)
            ConcurrentMCpi_second_menu(num2)
        case 39:
            showTheSpigotMagic()
        case 40:
            nifty_scoreBoard()
        case 60:
            ShowSLOC_behind_scoreBoard_40()
        case 41:
            BBPfConcurent()
        case 61:
            DisplayCodeBehind_BBPfConcurent()
        case 42:
            bbp_formula(num2)
        case 62:
            CodeRuneOf_bbp_formula()
        case 43:
            numerical_integration(num2)
        case 63:
            ShowLOC_behind_numerical_integration()
        case 44:
            Leibniz_method_one_billion_iters(num2)
        case 64:
            Show_Leibniz_method_one_billion_iters()
        case 45:
            fmt.Println(MonteCarloPi(99999999, num2))
        case 65:

        case 46:
            // display contents of prior results file
            content, err := ioutil.ReadFile("dataFromConstants.txt")  // 
                if err != nil {   // if the file does not exist ... 
                    fmt.Println("no prior results -- no log file dataFromConstants.txt exists\n")
                } else {
                    fmt.Println(string(content))  // dump/display entire file to command line
                }
        case 66:
            var rune4above = `            // display contents of prior results file
            content, err := ioutil.ReadFile("dataFromConstants.txt")  // 
                if err != nil {   // if the file does not exist ... 
                    fmt.Println("no prior results -- no log file dataFromConstants.txt exists")
                } else {
                    fmt.Println(string(content))  // dump/display entire file to command line
                }`
                fmt.Println(rune4above)
        case 47:
            os.Exit(1)
        case 98:
            fmt.Println(rick, "\n")
        case 99:
            Explain_spigot()
    }
}
