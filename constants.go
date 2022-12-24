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

func determinDiff (ProspectiveHit_Par int, smallerPerfectSquareOnce_T3 int) float32 { // takes two ints and returns a float32
        the_absolute_diff := float32(ProspectiveHit_Par) - float32(smallerPerfectSquareOnce_T3) 
        return (the_absolute_diff/100) 
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

func BBPF(num int) {
    fmt.Println("\nYou selected #", num, "the Bailey–Borwein–Plouffe formula for π, circa 1995\n")
    fmt.Println("How many digits of pi should we calculate? Enter one integer '4 to 16' ")
                 var numAi float64
    fmt.Scanf("%f", &numAi)
    fmt.Printf("pi calculated to %.0f places: %.15f \n", numAi, calculatePi(numAi)) 
    fmt.Println("and Pi from the web is   : 3.141592653589793")
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

// and now we cheat just to see if this is working ... 
                sr_per_sm := (math.Sqrt(float64(Prospective_smaller_hit)) / math.Sqrt(float64(smallerPerfectSquareOnce)))  // a 3:1 ratio 
                if sr_per_sm > 1.7320508066 { 
                    if sr_per_sm < 1.7320508074 { 
                        fmt.Println(sr_per_sm, "from top, is close at", iter, "iterations, hits are:", Prospective_smaller_hit, smallerPerfectSquareOnce)
                    } 
                } 
                    if sr_per_sm > 1.73205080752 {
                        if sr_per_sm < 1.7320508077 {
                            fmt.Println("\nPer the smaller ps, the \u221A3 has been calculated to be ... \n", sr_per_sm, " compared to:\n", "1.73205080757, the actual, after", iter, "iterations") 
                            fmt.Println("which is pretty good, and via first-principles of geometry alone\n... having found two very large perfect squares where one of them is")
                            fmt.Println("very-nearly exactly three times larger, i.e., the sqrt of", Prospective_smaller_hit, "over the sqrt of", smallerPerfectSquareOnce, "\n")
                            savedHit = Prospective_smaller_hit
                            smallerPerfectSquareOnceSaved = smallerPerfectSquareOnce
                        } 
                    }

// should instead be deciding final answer based on closeness of fit, percent diff, to exactly 3x smaller perfect square. Not just checking to see if we have gotten this close to the known value of Sqrt of 3

                sr_per_lg := (math.Sqrt(float64(ProspectiveHitOnLargeSide)) / math.Sqrt(float64(smallerPerfectSquareOnce)))
                if sr_per_lg < 1.732050808 {
                    if sr_per_lg > 1.73205080 { 
                        fmt.Println(sr_per_lg, "from bottom, is close at", iter, "iterations, hits are:", ProspectiveHitOnLargeSide, smallerPerfectSquareOnce)
                        } 
                    } 
                    if sr_per_lg < 1.7320508079 {
                        if sr_per_lg > 1.73205080 {  // 1.73205080757 (so 1.73205080754 - 1.73205080759 )
                            fmt.Println("\nPer the larger ps, the \u221A3 has been calculated to be ... \n", sr_per_lg, " compared to:\n", "1.73205080757, the actual, after", iter, "iterations") 
                            fmt.Println("which is pretty good, and via first-principles of geometry alone\n... having found two very large perfect squares where one of them is")
                            fmt.Println("very-nearly exactly three times larger, i.e., the sqrt of", ProspectiveHitOnLargeSide, "over the sqrt of", smallerPerfectSquareOnce, "\n")
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
        fmt.Println("a perfect bottom square from our initial list of perfect squares, times 3, is ...")
        fmt.Println(smallerPerfectSquareOnceSaved*3, " And the total iterations completed was ", TotalIterations)
        fmt.Println(savedHit, " was the another perfect square which we found by searching for a one that would be nearly-exactly three times larger\n")
        fmt.Println("... pretty close, I'd say\n")

            t := time.Now()
            elapsed := t.Sub(start)
    fmt.Println(TotalIterations, " iterations were completed in ", elapsed, " yielding 10 digits of sqrt of 3 \n") 
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
fmt.Println("... to view the code with attribution Enter '20' at the main menu")
}

func showTheMagic() {
fmt.Print("That last trick made possible by a bit of code I mooched off of GitHub:\n", spigotRune, "\n\n")
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
fmt.Println("... to view the code with attribution Enter '20' at the main menu")
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

func showMagicBehindsquareRootOf3() {
    var squareRootOf3rune = `
func squareRootOf3(num int) {
    start := time.Now() // saved start time to be compared with end time t
        buildTableOfOnlyPerfectSquares() 
            var index = 0 
            for index < 300000 {
                readTableOfPerfectSquares(index)  // pass-in the index 
                index = index + 2 
            }
        fmt.Println("a perfect bottom square from our initial list of perfect squares, times 3, is ...")
        fmt.Println(smallerPerfectSquareOnceSaved*3, " And the total iterations completed was ", TotalIterations)
        fmt.Println(savedHit, " was the another perfect square which we found by searching for a one that would be nearly-exactly three times larger\n")
        fmt.Println("... pretty close, I'd say\n")

            t := time.Now()
            elapsed := t.Sub(start)
    fmt.Println(TotalIterations, " iterations were completed in ", elapsed, " yielding 10 digits of sqrt of 3 \n") 
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

// and now we cheat just to see if this is working ... 
                sr_per_sm := (math.Sqrt(float64(Prospective_smaller_hit)) / math.Sqrt(float64(smallerPerfectSquareOnce)))  // a 3:1 ratio 
                if sr_per_sm > 1.7320508066 { 
                    if sr_per_sm < 1.7320508074 { 
                        fmt.Println(sr_per_sm, "from top, is close at", iter, "iterations, hits are:", Prospective_smaller_hit, smallerPerfectSquareOnce)
                    } 
                } 
                    if sr_per_sm > 1.73205080752 {
                        if sr_per_sm < 1.7320508077 {
                            fmt.Println("\nPer the smaller ps, the \u221A3 has been calculated to be ... \n", sr_per_sm, " compared to:\n", "1.73205080757, the actual, after", iter, "iterations") 
                            fmt.Println("which is pretty good, and via first-principles of geometry alone\n... having found two very large perfect squares where one of them is")
                            fmt.Println("very-nearly exactly three times larger, i.e., the sqrt of", Prospective_smaller_hit, "over the sqrt of", smallerPerfectSquareOnce, "\n")
                            savedHit = Prospective_smaller_hit
                            smallerPerfectSquareOnceSaved = smallerPerfectSquareOnce
                        } 
                    }

// should instead be deciding final answer based on closeness of fit, percent diff, to exactly 3x smaller perfect square. Not just checking to see if we have gotten this close to the known value of Sqrt of 3

                sr_per_lg := (math.Sqrt(float64(ProspectiveHitOnLargeSide)) / math.Sqrt(float64(smallerPerfectSquareOnce)))
                if sr_per_lg < 1.732050808 {
                    if sr_per_lg > 1.73205080 { 
                        fmt.Println(sr_per_lg, "from bottom, is close at", iter, "iterations, hits are:", ProspectiveHitOnLargeSide, smallerPerfectSquareOnce)
                        } 
                    } 
                    if sr_per_lg < 1.7320508079 {
                        if sr_per_lg > 1.73205080 {  // 1.73205080757 (so 1.73205080754 - 1.73205080759 )
                            fmt.Println("\nPer the larger ps, the \u221A3 has been calculated to be ... \n", sr_per_lg, " compared to:\n", "1.73205080757, the actual, after", iter, "iterations") 
                            fmt.Println("which is pretty good, and via first-principles of geometry alone\n... having found two very large perfect squares where one of them is")
                            fmt.Println("very-nearly exactly three times larger, i.e., the sqrt of", ProspectiveHitOnLargeSide, "over the sqrt of", smallerPerfectSquareOnce, "\n")
                            savedHitB = ProspectiveHitOnLargeSide
                        }
                    }
                    break // this break gets us out of the for loop having found a candidate within the larger if
                    // without this break nothing displays, and the proc is eventually killed
            }  // end "break-out" of the loop because "if we have found a candidate" we have already dealt with it so we need to break and get another potential from the for loop
    }       // END OF FOR LOOP after 96,000 index-d loops "iter-s"
    TotalIterations = iter // a global var to also be used elsewhere 
} 
    `
    fmt.Println(squareRootOf3rune)
}

func showMagicBehindBBPF() {
var BBPFrune = `
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
func BBPF(num int) {
    fmt.Println("\nYou selected #", num, "the Bailey–Borwein–Plouffe formula for π, circa 1995\n")
    fmt.Println("How many digits of pi should we calculate? Enter one integer '4 to 16' ")
                 var numAi float64
    fmt.Scanf("%f", &numAi)
    fmt.Printf("pi calculated to %.0f places: %.15f \n", numAi, calculatePi(numAi)) 
    fmt.Println("and Pi from the web is   : 3.141592653589793")
}
`
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

func RicksSwitch() {    // the primary if-then-else routine to execute a selection from the menu
    var num int 
        //fmt.Print("Enter your selection from above, 1 - 12 \u00a9 2022, by Richard Hart Woolley [it's an echo]\n")  // a kluge
        // the above kluge may be needed in some environments (windows) along with the duplicate in RicksDisplayMenuPi()
        fmt.Print("Enter your selection from above, 1 - 12 \n")
        fmt.Scanf("%d", &num)  // pause and request input from the user
        if num == 0 {
            fmt.Println("\nYou failed to make a selection, Hit Enter/Return to redisplay the menu, Ctrl-C to End/Exit")
        }
    if num > 40 && num < 10000 { num = 17 }  // to display a funny out-of-range message as case 17:
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
        case 12: // display contents of prior results file
            content, err := ioutil.ReadFile("dataFromConstants.txt")  // 
                if err != nil {   // if the file does not exist ... 
                    fmt.Println("no prior results -- no log file dataFromConstants.txt exists")
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

        case 19: 
            TheSpigot()
        case 20:
            showTheMagic()
        case 39:
            showTheMagic()

        case 13: 
            fmt.Println(" ... So sorry, but", num, "was not an option")
        case 14: 
            fmt.Println(num, " was not an option! It was not on the menu, go fish!\n")
        case 34:
            /*
            var precision int 
            fmt.Println("what shall be thy precision ")
            fmt.Scanf("%d", &precision) 
            fmt.Println("result is ", calculatePispi(precision))
            */
        case 15: 
            fmt.Println("Your selection of", num, " is right-out!  Go Fish!\n")
        case 16: 
            fmt.Println("Your selection is really-far-out!  Go Fish!\n")
        case 17: 
            fmt.Println("\nOops, how'd we get here? Hit Enter/Return again to possibly redisplay the menu")
        case 18:
            //fmt.Println("\n result via the spigot ... ", calculatePispi()) 
    default: 
        fmt.Println("this is the switch default code, after a break option ??")
    } 
} 

/*
// never got anything to work with this 
// 999999999 gives 0.06427714842
// 99        gives 0.06427714842
// 9         gives 246.5114231795
// 20        gives 0.06427714842
func calculatePispi(precision int) float64 {
    // Initialize variables to store the current and previous digits of pi
    var currentDigit int
        currentDigit = 0 // this cannot matter
    var previousDigit int
        previousDigit = 3 // affects currentDigit

    // Initialize a variable to store the result
    var result float64
        result = 0.0
    iters := 1
        for iters < 2 {
            i := 1
            // Iterate through the digits of pi, starting with the most significant digit
            for i <= precision {
                i++
                // Calculate the next digit of pi using the spigot algorithm
                    //fmt.Println("previous digit top of inner is ", previousDigit, "\n") 
                currentDigit = (previousDigit * 10.0) / (2.0 * (i - 1.0))

                // Update the result with the new digit
                result = result + (float64(currentDigit) / (math.Pow(10, float64(i))))

                // Update the previous digit
                previousDigit = currentDigit
                //i++
            //fmt.Print("\ncurrentl digit inner is ", currentDigit, "\n")
            //fmt.Println("previous digit inner is ", previousDigit, "\n") 
            }
            iters++
            //fmt.Print("\ncurrentl digit outer is ", currentDigit, "\n")
            //fmt.Println("previous digit outer is ", previousDigit, "\n") 
            //fmt.Println("iters is ", iters)
        }
    // Return the result
    return result
}
*/

func RicksDisplayMenuPi() {
fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nEnter 33 to see the the magic behind main (selector+20 for all others)\n") 
// Veritassium https://youtu.be/gMlf1ELvRzc?list=LL  was the initial inspiration for all of this ...
// ... option 1 below having been discussed on his channel 
fmt.Println("1:  Calculate the Square Root of 3 (\u221A3) from first-principles")
fmt.Println("    ... i.e., from a ratio of 3:1 of perfect squares\n")
fmt.Println("2:  Pi: Bailey–Borwein–Plouffe formula for π, discovered in 1995 (best method)\n")
fmt.Println("3:  Pi: (worst method) requires calculating the \u221A3 (see selection 1)")
fmt.Println("     ... as it only yeilds 4 digits of π \n")
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
fmt.Println("19:  Pi: Open the 'Spigot' algorithm, instantly calculates way too much pie\n")
//fmt.Print("Enter your selection from above, 1 - 12 \u00a9 2022, by Richard Hart Woolley\n")
// the above kluge is definately not needed in a Linux environment
fmt.Println("Ctrl-C to End/Exit  SLOC = 2850   \u00a9 2022, by Richard Hart Woolley \n")
}