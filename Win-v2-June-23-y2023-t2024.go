package main

import (
    "os"        // Used to fetch the name of your system 
    "io/ioutil" // Used for file access 
    "fmt"       // Used for printing etc. 
    "math"      // Used for math.Pow(base, exponent) and other more-complex math funcs 
    //"runtime"   // Used to get information on available CPUs
    "time"      // Used for seeding random number generation; and calculating run times 
    "strconv"   // Used in Spigot 
    "os/signal" // Used in *** nifty scoreboard *** concurrent computation of pi using Nilakantha's formula, by Diego Brener diegosilva13 on Github 
    "math/big"  // Used in ArchimedesBig(), and in bpp formula with mods by rick.woolley@gmail.com (pi calculator from https://github.com/karan/Projects)
    "sort"      // Used in xRootOfy (case 18) to sort an array of structs 
    "regexp"

    "bufio"    // used to count the lines in this file 
    "log"
    "strings"
)

func main() {       // top-level program logic flow -- explore SEVENTEEN ways to calculate Pi, plus THREE other constants
    for 1 == 1 {
        totalLines, nonEmptyLines, err := reportSLOCstats()
            if err != nil {
                log.Fatal(err)
            }
        DisplayMenus(totalLines, nonEmptyLines)
        fmt.Println(string(colorRed), "Finished. Hit Enter to redisplay the Main Menu", string(colorReset)) // this will be the last line of every case #: 
            var JustToPauseTheLoop int 
            fmt.Scanf("%d", &JustToPauseTheLoop) 
            fmt.Scanf("%d", &JustToPauseTheLoop) 
            //fmt.Scan(&JustToPauseTheLoop) // Scan does not work quite as well as the Scanf combo below (a single Scanf worked well in Unix environment)
    }
} 


func reportSLOCstats() (totalLines, nonEmptyLines int, err error) {
    pathNameOfMeAsExe := os.Args[0]
        re := regexp.MustCompile(`\\exe\\(.*?)\.exe`) 
        match := re.FindStringSubmatch(pathNameOfMeAsExe) // grab the stuff between /exe/ and .exe in the full pathname of the generated executable 
                var substring string 
            if len(match) >= 2 {
                substring = match[1]                   // substring is loaded with our base file name 
            } else {
                fmt.Println("Substring not found")
            }

    filename := substring + ".go"                // append the .go on the end of it 
    file, err := os.Open(filename)
    if err != nil {
        return 0, 0, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        totalLines++
        line := strings.TrimSpace(scanner.Text())
        if line != "" {
            nonEmptyLines++
        }
    }

    if scanner.Err() != nil {
        return 0, 0, scanner.Err()
    }

    return totalLines, nonEmptyLines, nil
}


func DisplayMenus(totalLines, nonEmptyLines int) {
    pages_of_code := totalLines / 49 
    selection = 0 // a switching flag? But also definately the scanned value 
    date := time.Now()

fmt.Println(string(colorYellow), 
  "\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nMAIN MENU",

  string(colorGreen), "   build date :", string(colorCyan), date.Format("(01-02-2006)"), "\n", string(colorReset))

fmt.Println(string(colorRed), "Enter your selection from below, or 96 for notes on using this app\n", string(colorReset))
fmt.Println("2:", string(colorCyan), "Pi:", string(colorReset), "Bailey–Borwein–Plouffe formula for π, discovered in 1995", string(colorYellow),
 "(best method)", string(colorReset))
fmt.Println("14:", string(colorCyan), "Pi:", string(colorReset), "Archimedes' method but using Go's math/big objects, up to 2,712 digits\n") 
fmt.Println("15:", string(colorCyan), "Pi:", string(colorReset), "Chudnovsky method, state of the art! But God only knows how it works\n") 

fmt.Println("5:", string(colorCyan), "Pi:", string(colorReset), "An infinite series by Indian astronomer Nilakantha Somayaji circa 1500 ")
fmt.Println("    ... this equation can be found in an ancient Sanskrit verse")
fmt.Println("    π = 3 + 4/(2*3*4) - 4/(4*5*6) + 4/(6*7*8) - 4/(8*9*10) + 4/(10*11*12) ...")
fmt.Println("      One-Hundred-Million iterations will be executed -- in less than a second")
fmt.Println("      14 digits of π -- VERY FAST -- gives 7 digits in just 100 iterations !!!")
fmt.Println("6:", string(colorCyan), "Pi:", string(colorReset), "Gottfried Wilhelm Leibniz, late 17th century ( and by Isaac Newton )")
fmt.Println("    π = 4 * ( 1 - 1/3 + 1/5 - 1/7 + 1/9 ...")
fmt.Println("      4 Billion iterations will be executed")
fmt.Println("      10 digits of π -- 8 digits in 100M iterations in under a minute ")
fmt.Println("7:", string(colorCyan), "Pi:", string(colorReset), "The Gregory-Leibniz series circa 1676")
fmt.Println("    π = (4/1) - (4/3) + (4/5) - (4/7) + (4/9) - (4/11) + (4/13) - (4/15) ...")
fmt.Println("      Three-Hundred-Million iterations will be executed")
fmt.Println("      9 digits of π -- in seconds\n")
fmt.Println("12:  Display prior execution times from longer-running prior selections \n")

//           13:                  v v v v v v v v v v v 
fmt.Println(string(colorYellow), "13:  FOR SECOND MENU\n", string(colorReset))

fmt.Println("47:  to End/Exit", string(colorCyan), "tSLOC =", totalLines, "eSLOC =", nonEmptyLines, string(colorPurple), 
  "\u00a9 2023, by Richard Hart Woolley", string(colorReset))
fmt.Println(string(colorCyan),"                 That is", int(pages_of_code), "pages of code at 49 lines per page\n", string(colorReset))
fmt.Print("Enter your selection, 1 -> x", string(colorRed), " (IS THIS WINDOW MAXIMIZED?  Well, do it!)\n", string(colorReset)) 

    fmt.Scan(&selection)  


if selection == 13 { //==============================================================================================================================

fmt.Print(string(colorYellow))
fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nSECOND MENU:", 

  string(colorRed), "Enter your selection from below, or 96 for notes on using this app", string(colorReset), "\n") 
fmt.Println("18:  Calculate", string(colorGreen), "the second or third Root of y (x\u221Ay)", string(colorReset), "from first-principles")
fmt.Println("    ... via the ratio of y:1 of perfect products (slower general ver of #1)\n")
fmt.Println("19:", string(colorGreen), "Pi:", string(colorReset), "Open the 'Spigot' algorithm, instantly bakes way too much pie :)")
fmt.Println("    ... Can easily spit out 10,000 digits of π !!!!!\n", string(colorCyan))
fmt.Println("8:", string(colorGreen), " Pi:", string(colorReset), "An infinite series by John Wallis circa 1655")
fmt.Println("      π = 2 * ((2/1)*(2/3)) * ((4/3)*(4/5)) * ((6/5)*(6/7)) ... ")
fmt.Println("        One-Billion iterations will be executed; option for 40 billion iterations")
fmt.Println("        9 digits of π -- a billion loops, in seconds -- option for 10 digits\n")
fmt.Println("9:   Calculate", string(colorGreen), "Euler's Number: \u2107 or \u2147", string(colorReset), "the natural logarithmic base")
fmt.Println("        Explore the limit of interest\n")
fmt.Println("10:  Calculate the", string(colorGreen), "Erdős-Borwein constant", string(colorReset), "from a breif infinite series\n")
fmt.Println("40:", string(colorGreen), "Pi:", string(colorReset), "Nifty 'ScoreBoard' using Nilakantha's formula", string(colorYellow), 
  "(Ctrl-C to exit it)", string(colorCyan))
fmt.Println("41:", string(colorGreen), "Pi:", string(colorReset), "Bailey–Borwein–Plouffe formula", string(colorCyan), "[concurent]", 
  string(colorReset))
fmt.Println("42:", string(colorGreen), "Pi:", string(colorReset), "BBP formula to 46 digits")
fmt.Println("44:", string(colorGreen), "Pi:", string(colorReset), "via Leibniz method in one billion iterations [runs a while]")

fmt.Println(string(colorYellow), "0:  To return to main menu", string(colorRed), "\n")

        // ("47:  to End/Exit vvvvvvvvvvvvvvvv
fmt.Println("47: to End/Exit", string(colorCyan), "tSLOC =", totalLines, "eSLOC =", nonEmptyLines, string(colorPurple), 
  "\u00a9 2023, by Richard Hart Woolley\n", string(colorReset))
fmt.Print("Enter your selection, 1 -> x", string(colorRed), " (IS THIS WINDOW MAXIMIZED?  Well, do it!)\n", string(colorReset)) 

    fmt.Scan(&selection)  

} //===============================================================================================================================================

    if selection == 17 || selection == 20 { 
       selection = 120 // 17 and 20 are handled by case 120: 
    }

    switch selection { // if selection == 2 then handel per case 2: etc. 

        case 2:  // (best method) the Bailey–Borwein–Plouffe formula for π, circa 1995 
            BBPF(selection)

        case 5: // a series by Indian astronomer Nilakantha Somayaji circa 1500 AD 
            Nilakantha(selection)

        case 6: // Gottfried Wilhelm Leibniz
            GottfriedWilhelmLeibniz(selection)

        case 7: // the Gregory-Leibniz series
            GregoryLeibniz(selection)

        case 8: // John Wallis circa 1655 --- runs very long 
            JohnWallis(selection)

        case 9: // Euler's Number
            EulersNumber(selection)

        case 10: // Erdős–Borwein constant
            ErdosBorwein(selection)

        case 12: // display contents of prior results file
            content, err := ioutil.ReadFile("dataLog-From_calculate-pi-and-friends.txt")  // 
                if err != nil {   // if the file does not exist ... 
                    fmt.Println(string(colorCyan), "\nNo prior results -- no log file", string(colorWhite), "'dataLog-From_calculate-pi-and-friends.txt'", string(colorCyan), "exists\n")
                } else {
                    fmt.Println(string(content))  // dump/display entire file to command line
                }
        case 32:
            var fileAccessRune = `
        case 12: // display contents of prior results file
            content, err := ioutil.ReadFile("dataLog-From_calculate-pi-and-friends.txt")  // 
                if err != nil {   // if the file does not exist ... 
                    fmt.Println(string(colorCyan), "\nNo prior results -- no log file", string(colorWhite), "'dataLog-From_calculate-pi-and-friends.txt'", string(colorCyan), "exists\n")
                } else {
                    fmt.Println(string(content))  // dump/display entire file to command line
                }`
            fmt.Println(string(colorCyan), fileAccessRune, string(colorReset))

        case 14: 
            ArchimedesBig(selection)

        case 15: 
            chud(selection)

        case 16: 
            fmt.Println("Your selection is really-far-out!  Go Fish!\n")
        
        case 18:
            xRootOfy(selection) 
        
        case 19: 
            TheSpigot()

        case 40: // a special case ... if some time has elapsed running the func then we log the stats to a file (because this func is always killed with a ctrl-c)
                start := time.Now() // saved start time to be compared with end time t 
            nifty_scoreBoard()
                t := time.Now()
                elapsed := t.Sub(start)
            // only if 
                if int(elapsed.Seconds()) != 0 {
                        fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                            check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                        Hostname, _ := os.Hostname()
                        _ , err0 := fmt.Fprintf(fileHandle, "\n  -- ScoreBoard -- selection #%d on %s \n", selection, Hostname)
                            check(err0)
                        current_time := time.Now()
                        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                        check(err6)
                        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
                            check(err7)
                }

        case 42:
            bbp_formula(selection)

        case 44:
            Leibniz_method_one_billion_iters(selection)

        case 47:
            os.Exit(1)

        case 180:
            xRootOfyCcaller(selection)
        case 0: 
            // Only useful for Windows variant as a non-functional case to allow returning directly to main menu 
        case 13:
            // ... another non-functional case to allow returning directly to main menu in the windows variant 
    default:
        fmt.Println("\n ... You entered a value that is", string(colorRed), "not a valid option", string(colorReset), "go fish\n")
        // handles every case not explicitly listed (which has not been otherwise handled) 
    }
}



// case 180:
func xRootOfyCcaller(selection int) {
    precisionOfRoot = 1700
    workPiece = 2
    radical_index_number := 2
        fmt.Println("\nThis proc endlessly calculates the root of all integers beginging with 2; and builds a txt file of its results.\n")
        fmt.Println("\nEnter 2 for SquareRoot, or 3 for CubeRoot\n")
        fmt.Scan(&radical_index_number)  // pause and request input from the user 
    radical_index := radical_index_number
    for 1 == 1 {
        sortedResults = nil
        xRootOfyC(selection, workPiece, radical_index)
        workPiece++
    }
}

func xRootOfyC(selection int, workPiece int, radical_index int) { // calculates either square or cube root of any integer 

    var index = 0 // counter used in the for loop in this func :: is also passed to the pricipal func readTheTableOfPP // cannot be a global

        startFromTop := time.Now()

            if radical_index == 3 { // if doing a cube root special tolerances are set here for certain problem values, i.e., 2, 11, 17, 3, 4, or 14
                if workPiece > 4 {
                    precisionOfRoot = 1655
                    fmt.Println("\n Default precision is 1700 \n")
                }
                if workPiece == 2 || workPiece == 11 || workPiece == 17 {
                    precisionOfRoot = 580
                    fmt.Println("\n resetting precision to 600 \n")
                }
                if workPiece == 3 || workPiece == 4 || workPiece == 14 {
                    precisionOfRoot = 875
                    fmt.Println("\n resetting precision to 900 \n")
                }
            }
            if radical_index == 2 { // if doing a square root we just use a tolerance of 4 for all workpieces. 
                precisionOfRoot = 4
            }

    //radical_index := getInputFromUserAndSetPrecision() // this step is not needed in this continuous automated version 

    buildTableOfPerfectProductsC(radical_index) // 800,000 entries, 400,000 pairs 

// The following section consists of the principal for loop with a conditional break ------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------------------------------------------------------
// case 180:

        startBeforeCall := time.Now()

    //for index < 380000 { // the table has 800,000 entries, 400,000 pairs; so index increments by 2 at the bottom of this loop 
    for index < 400000 { // the table has 825,000 entries, > 410,000 pairs; so index increments by 2 at the bottom of this loop 


        readTheTableOfPPC(index, startBeforeCall, radical_index) // pass-in an index to the table: 380,000 indexs corresponding to the number of pairs of entries 

        handlePerfectSquaresAndCubesC(startFromTop, radical_index, selection) // handle the rare case of a perfect square or cube 

                if diffOfLarger == 0 || diffOfSmaller == 0 {  // Then, it was a perfect square or cube; so, we need to ... 
                    break // ... out of the for loop because we are done : the workpiece was either a perfect square or a perfect cube
                }


        fileHandle_180, err31 := os.OpenFile("logfile_from_selection_180.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
            check(err31)                                // ... gets a file handle to logfile_from_selection_180.txt
            defer fileHandle_180.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.

                    if index == 80000 {
                        fmt.Printf("\n80,000 ... still working ... workPiece is: %d, TimeWindow is: %0.4f, precision is: %d\n", workPiece, Tim_win, precisionOfRoot)
                        _ , err01 := fmt.Fprintf(fileHandle_180, "\n80,000 ... still working ... workPiece is: %d, TimeWindow is: %0.4f, precision is: %d\n", workPiece, Tim_win, precisionOfRoot)
                        check(err01)
                    }    
                    if index == 160000 {
                        fmt.Printf("\n160,000 ... still working ... workPiece is: %d, TimeWindow is: %0.4f, precision is: %d\n", workPiece, Tim_win, precisionOfRoot)
                        _ , err02 := fmt.Fprintf(fileHandle_180, "\n160,000 ... still working ... workPiece is: %d, TimeWindow is: %0.4f, precision is: %d\n", workPiece, Tim_win, precisionOfRoot)
                        check(err02)
                    }
                    if index == 240000 {
                        fmt.Printf("\n240,000 ... still working ... workPiece is: %d, TimeWindow is: %0.4f, precision is: %d\n", workPiece, Tim_win, precisionOfRoot)
                        _ , err03 := fmt.Fprintf(fileHandle_180, "\n240,000 ... still working ... workPiece is: %d, TimeWindow is: %0.4f, precision is: %d\n", workPiece, Tim_win, precisionOfRoot)
                        check(err03)
                    }                
                    if index == 320000 {
                        fmt.Printf("\n320,000 ... still working ... workPiece is: %d, TimeWindow is: %0.4f, precision is: %d\n", workPiece, Tim_win, precisionOfRoot)
                        _ , err04 := fmt.Fprintf(fileHandle_180, "\n320,000 ... still working ... workPiece is: %d, TimeWindow is: %0.4f, precision is: %d\n", workPiece, Tim_win, precisionOfRoot)
                        check(err04)
                    }        
                    fileHandle_180.Close()        

        index = index + 2 // increment the index and read the table again 
    } // end of for loop // the above break statement is NOT the only way to exit this for loop, it also terminates after 380,000 iterations of index 

// All of the remaining sections are conditional for workpiece NOT being a perfect square or cube
//if perfectResult2 == 0 && perfectResult3 == 0 {  // Then, it was NOT a perfect square or cube, so handle that case 
// the remaining sections are only reached after having exited the primary for loop above via a break statement or an exaustive reading of the table ------------
//---------------------------------------------------------------------------------------------------------------------------------------------------------------
        // calculate elapsed time 
        t_s2 := time.Now()
        elapsed_s2 := t_s2.Sub(startFromTop)


// the following sections log the final results to a text file (and also does one conditional Printf) -------------------------------------------------
//-----------------------------------------------------------------------------------------------------------------------------------------------------
        fileHandle_180, err311 := os.OpenFile("logfile_from_selection_180.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
            check(err311)                                // ... gets a file handle to logfile_from_selection_180.txt
            defer fileHandle_180.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.

        Hostname, _ := os.Hostname()
        _ , err30 := fmt.Fprintf(fileHandle_180, "\n  -- %d root of %d by a ratio of perfect Products -- selection #%d on %s \n", radical_index, workPiece, selection, Hostname)
            check(err30)

        current_time := time.Now()
        _ , err36 := fmt.Fprint(fileHandle_180, "was run on: ", current_time.Format(time.ANSIC), "\n")
        check(err36)

        _ , err35 := fmt.Fprintf(fileHandle_180, "%d was total Iterations, precisionOfRoot was %d \n", index, precisionOfRoot)
            check(err35)

        fileHandle_180.Close()


    // Sort the slice sortedResults by its pdiff field : 
    //-----------------------------------------------------------------------------------------------------------
        sort.Slice(sortedResults, func(i, j int) bool { return sortedResults[i].pdiff < sortedResults[j].pdiff})


    // display and print the best-fitting result based solely on the lowest pdiff :
    //-----------------------------------------------------------------------------
    // case 180: 

        // display the best fitting result :
            if radical_index == 2 {
                if len(sortedResults) > 0 {
                    fmt.Printf("\n %0.9f, is the best approximation for the Square Root of %d \n", sortedResults[0].result, workPiece)
                }
            }
            if radical_index == 3 {
                if len(sortedResults) > 0 {
                    fmt.Printf("\n %0.9f, is the best approximation for the  Cube  Root of %d \n", sortedResults[0].result, workPiece)
                }
            }

            // Fprint/log the best fitting result :
            fileHandle_180, err310 := os.OpenFile("logfile_from_selection_180.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
            check(err310)                                // ... gets a file handle to logfile_from_selection_180.txt
            defer fileHandle_180.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.

                if radical_index == 2 {
                    fmt.Println("\n len of sorted was ", len(sortedResults))
                    _ , err11 := fmt.Fprintf(fileHandle_180, "len of sorted was %d", len(sortedResults))
                    check(err11)
                    if len(sortedResults) > 0 {
                    _ , err48 := fmt.Fprintf(fileHandle_180, "\n %0.9f, is the best approximation for the Square Root of %d \n", sortedResults[0].result, workPiece)
                    check(err48)
                    }
                    fmt.Println("\nThat's all folks!\n\nNext ...")
                    _ , err22 := fmt.Fprint(fileHandle_180, "\nThat's all folks!\n\nNext ...")
                    check(err22)
                }
                if radical_index == 3 {
                    fmt.Println("\n len of sorted was ", len(sortedResults))
                    _ , err11 := fmt.Fprintf(fileHandle_180, "len of sorted was %d", len(sortedResults))
                    check(err11)
                    if len(sortedResults) > 0 {
                        _ , err49 := fmt.Fprintf(fileHandle_180, "\n %0.9f, is the best approximation for the  Cube  Root of %d \n", sortedResults[0].result, workPiece)
                        check(err49)
                    }
                    fmt.Println("\nThat's all folks!\n\nNext ...")
                    _ , err22 := fmt.Fprint(fileHandle_180, "\nThat's all folks!\n\nNext ...")
                    check(err22)
                }
    //----------------------------------------------------------------------------- case 180:


        TotalRun := elapsed_s2.String() // cast time durations to a String type for Fprintf "formatted print"
        _ , err57 := fmt.Fprintf(fileHandle_180, "Total run was %s \n ", TotalRun) 
            check(err57)

        fileHandle_180.Close()
}

// in case 180:
func handlePerfectSquaresAndCubesC(startFromTop time.Time, radical_index int, selection int) {
// the next sub-section detects, traps, and reports the detection of either a perfect square of a perfect cube ------------------
// ... it also is responsible for causing the algorithm to terminate via a break if workpiece was a perfect square or cube 
//-------------------------------------------------------------------------------------------------------------------------------
        if diffOfLarger == 0 || diffOfSmaller == 0 {  // Then, it was a perfect square or cube 

            t_s1 := time.Now() 
            elapsed_s1 := t_s1.Sub(startFromTop)

                fileHandle_180, err1 := os.OpenFile("logfile_from_selection_180.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) 
                    check(err1)                                // ... gets a file handle to logfile_from_selection_180.txt
                    defer fileHandle_180.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.

                Hostname, _ := os.Hostname()
                _ , err0 := fmt.Fprintf(fileHandle_180, "\n  -- %d root of %d by a ratio of PerfectProducts -- selection #%d on %s \n", 
                    radical_index, workPiece, selection, Hostname)
                    check(err0)

                current_time := time.Now()
                _ , err6 := fmt.Fprint(fileHandle_180, "was run on: ", current_time.Format(time.ANSIC), "\n")
                check(err6)

                TotalRun := elapsed_s1.String() // cast time durations to a String type for Fprintf "formatted print"
                _ , err7 := fmt.Fprintf(fileHandle_180, "Total run was %s \n ", TotalRun) 
                    check(err7)

                    if radical_index == 2 {
                        _ , err8 := fmt.Fprintf(fileHandle_180, "the %d root of %d is %0.2f \n", radical_index, workPiece, perfectResult2)
                            check (err8)
                    }
                    if radical_index == 3 {
                        _ , err38 := fmt.Fprintf(fileHandle_180, "the %d root of %d is %0.2f \n", radical_index, workPiece, perfectResult3)
                            check (err38)
                    }
                    fileHandle_180.Close()

            // break // break out of the for loop because we are done : the workpiece was either a perfect square or a perfect cube 

        } // end of if :: if it was a perfect square or cube 
//-------------------------------------------------------------------------------------------------------------------------------

}

// in case 180:
func readTheTableOfPPC(index int, startBeforeCall time.Time, radical_index int) {  // this gets called 380,000 times. 

    // The first time it is called index is 0

    // read it ...
    smallerPerfectProductOnce := Table_of_perfect_Products[index]  
    // ... and save it locally, do this just-once per func call. // ... index may be 0 up to 380,000
    RootOfsmallerPerfectProductOnce := Table_of_perfect_Products[index+1]
    // ^^^ also read the root wich corresponds

        iter := 0
    for iter < 410000 { // 410,000 loops. Why do we need so many?, Because index may be 0 to 400,000 ?? and we need to read through 825,000 table entries 
        iter++

        index = index + 2
        largerPerfectProduct := Table_of_perfect_Products[index]  
        // to approximate the root of an imperfect square x we will need a ratio of two perfect squares wich is about equal to x
        // ...we need to find two perfect squares such that one is about x times larger than the other
        // get next perfect square from table for testing to see if it is more than x * bigger than smallerPerfectProductOnce

        if largerPerfectProduct > smallerPerfectProductOnce*workPiece {  
        // if largerPerfectProduct is a candidate based on it being just-a-bit larger than workPiece* the smaller PP deal with that, else loop to the next potential 

            ProspectiveHitOnLargeSide := largerPerfectProduct  // make a copy under a more suitable name :) 
            rootOfProspectiveHitOnLargeSide := Table_of_perfect_Products[index+1] // the current value of index plus one holds the root of largerPerfectSquare hence the root of ProspectiveHitOnLargeSide

            ProspectiveHitOnSmallerSide := Table_of_perfect_Products[index-2]  
        // save that smaller one too //                               ^^ 2 now instead of 1 because we have added roots to the slice
            rootOfProspectiveHitOnSmallerSide := Table_of_perfect_Products[index-1]

            diffOfLarger = ProspectiveHitOnLargeSide - workPiece*smallerPerfectProductOnce
                //diffOfSmaller = -(ProspectiveHitOnSmallerSide - workPiece*smallerPerfectProductOnce) // this was dumb ?? 
            diffOfSmaller = workPiece*smallerPerfectProductOnce - ProspectiveHitOnSmallerSide

// detect perfect squares and set global vars to their roots -----------------------------------------------
            if diffOfLarger == 0 {
                fmt.Println(string(colorCyan), "\n The", radical_index, "root of", workPiece, "is", string(colorGreen), 
                float64(rootOfProspectiveHitOnLargeSide) / float64(RootOfsmallerPerfectProductOnce), string(colorReset), "\n")

                perfectResult2 = (math.Sqrt(float64(workPiece))) // these global values are used later to Fprint to a log file
                perfectResult3 = math.Cbrt(float64(workPiece)) 
                break // out of the for loop because the workPiece is itself a perfect square
            }
            if diffOfSmaller == 0 {
                fmt.Println(string(colorCyan), "\n The", radical_index, "root of", workPiece, "is", string(colorGreen), 
                float64(rootOfProspectiveHitOnSmallerSide) / float64(RootOfsmallerPerfectProductOnce), string(colorReset), "\n")

                perfectResult2 = (math.Sqrt(float64(workPiece))) // these global values are used later to Fprint to a log file
                perfectResult3 = math.Cbrt(float64(workPiece)) 
                break // out of the for loop because the workPiece is itself a perfect square
            }
// ---------------------------------------------------------------------------------------------------------
// case 180:


// larger side section: of case 180: ---------------------------------------------------------------------------------------------------------------------------
//--------------------------------------------------------------------------------------------------------------------------------------------------------------

            if diffOfLarger < precisionOfRoot {  // report the prospects, their differences, and the calculated result for the Sqrt or Cbrt 
                fmt.Println("small PP is", string(colorCyan), smallerPerfectProductOnce, string(colorReset), "and, slightly on the higher side of", workPiece, 
                "* that we found a PP of", string(colorCyan), ProspectiveHitOnLargeSide, string(colorReset), "a difference of", diffOfLarger)

                fmt.Println("the ", radical_index, " root of ", workPiece, " is calculated as ", string(colorGreen), 
                float64(rootOfProspectiveHitOnLargeSide) / float64(RootOfsmallerPerfectProductOnce), string(colorReset))

                fmt.Printf("with pdiff of %0.4f \n", (float64(diffOfLarger)/float64(ProspectiveHitOnLargeSide))*100000)

                fileHandle_180, err1 := os.OpenFile("logfile_from_selection_180.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) 
                    check(err1)                                // ... gets a file handle to logfile_from_selection_180.txt
                    defer fileHandle_180.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.

                _ , err2 := fmt.Fprint(fileHandle_180,"\nsmall PP is ", smallerPerfectProductOnce, " and, slightly on the higher side of ", 
                    workPiece, " * that we found a PP of ", ProspectiveHitOnLargeSide, " a difference of ", diffOfLarger)
                        check(err2)
                _ , err3 := fmt.Fprint(fileHandle_180,"\nthe ", radical_index, " root of ", workPiece, " is calculated as ", 
                    float64(rootOfProspectiveHitOnLargeSide) / float64(RootOfsmallerPerfectProductOnce))
                        check(err3)
                _ , err6 := fmt.Fprintf(fileHandle_180,"\nwith pdiff of %0.4f, precisionOfRoot is %d\n", 
                    (float64(diffOfLarger)/float64(ProspectiveHitOnLargeSide))*100000, precisionOfRoot)
                        check(err6)

                fileHandle_180.Close()


                // save the result to an accumulator array so we can Fprint all such hits at the very end
                //List_of_2_results_case18 = append(List_of_2_results_case18, float64(rootOfProspectiveHitOnLargeSide) / float64(RootOfsmallerPerfectProductOnce) )
                //corresponding_diffs = append(corresponding_diffs, diffOfLarger)
                //diffs_as_percent = append(diffs_as_percent, float64(diffOfLarger)/float64(ProspectiveHitOnLargeSide))

                // in the next five lines we load (append) a record into/to the file (array) of Results 
                Result1 := Results{
                    result: float64(rootOfProspectiveHitOnLargeSide) / float64(RootOfsmallerPerfectProductOnce),
                    pdiff: float64(diffOfLarger)/float64(ProspectiveHitOnLargeSide),
                }
                sortedResults = append(sortedResults, Result1)

                t2 := time.Now()
                elapsed2 := t2.Sub(startBeforeCall)
                // if needed, notify the user that we are still working 
                    Tim_win = 0.22
                    if radical_index == 3 {
                        if workPiece > 13 {
                            Tim_win = 0.002 
                        } else {
                        Tim_win = 0.0045 
                        }                    
                    }   
                if float64(elapsed2.Seconds()) > Tim_win {
                    fmt.Println(float64(elapsed2.Seconds()), "Seconds have elapsed ... working ..., time window is: ", Tim_win, "\n")
                }
            }
//---------------------------------------------------------------------------------------------------------------------------------------------------------------
// in case 180:


// smaller side section: of case 180: ---------------------------------------------------------------------------------------------------------------------------
//---------------------------------------------------------------------------------------------------------------------------------------------------------------
            if diffOfSmaller < precisionOfRoot {  // report the prospects, their differences, and the calculated result for the Sqrt or Cbrt 
                fmt.Println("small PP is", string(colorCyan), smallerPerfectProductOnce, string(colorReset), "and, slightly on the lesser side of", workPiece, 
                "* that we found a PP of", string(colorCyan), ProspectiveHitOnSmallerSide, string(colorReset), "a difference of", diffOfSmaller)

                fmt.Println("the ", radical_index, " root of ", workPiece, " is calculated as ", string(colorGreen), 
                float64(rootOfProspectiveHitOnSmallerSide) / float64(RootOfsmallerPerfectProductOnce), string(colorReset))

                fmt.Printf("with pdiff of %0.4f \n", (float64(diffOfSmaller)/float64(ProspectiveHitOnSmallerSide))*100000)


                fileHandle_180, err1 := os.OpenFile("logfile_from_selection_180.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) 
                    check(err1)                                // ... gets a file handle to logfile_from_selection_180.txt
                    defer fileHandle_180.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.

                _ , err2 := fmt.Fprint(fileHandle_180,"\nsmall PP is ", smallerPerfectProductOnce, " and, slightly on the higher side of ", 
                    workPiece, " * that we found a PP of ", ProspectiveHitOnLargeSide, " a difference of ", diffOfLarger)
                        check(err2)
                _ , err3 := fmt.Fprint(fileHandle_180,"\nthe ", radical_index, " root of ", workPiece, " is calculated as ", 
                    float64(rootOfProspectiveHitOnLargeSide) / float64(RootOfsmallerPerfectProductOnce))
                        check(err3)
                _ , err6 := fmt.Fprintf(fileHandle_180,"\nwith pdiff of %0.4f, precisionOfRoot is %d\n", 
                    (float64(diffOfLarger)/float64(ProspectiveHitOnLargeSide))*100000, precisionOfRoot)
                        check(err6)

                fileHandle_180.Close()

                // save the result to three accumulator arrays so we can Fprint all such hits, diffs, and p-diffs, at the very end of run 
                //List_of_2_results_case18 = append(List_of_2_results_case18, float64(rootOfProspectiveHitOnSmallerSide) / float64(RootOfsmallerPerfectProductOnce) )
                //corresponding_diffs = append(corresponding_diffs, diffOfSmaller)
                //diffs_as_percent = append(diffs_as_percent, float64(diffOfSmaller)/float64(ProspectiveHitOnSmallerSide))

                // in the next five lines we load (append) a record into/to the file (array) of Results 
                Result1 := Results{
                    result: float64(rootOfProspectiveHitOnSmallerSide) / float64(RootOfsmallerPerfectProductOnce),
                    pdiff: float64(diffOfSmaller)/float64(ProspectiveHitOnSmallerSide),
                }
                sortedResults = append(sortedResults, Result1)

                t2 := time.Now()
                elapsed2 := t2.Sub(startBeforeCall)
                // if needed, notify the user that we are still working 
                    Tim_win = 0.22
                    if radical_index == 3 {
                        if workPiece > 13 {
                            Tim_win = 0.002 
                        } else {
                        Tim_win = 0.0045 
                        }                    
                    }   
                if float64(elapsed2.Seconds()) > Tim_win  {
                    fmt.Println(float64(elapsed2.Seconds()), "Seconds have elapsed ... working ..., time window is: ", Tim_win, "\n")
                }
            }
//---------------------------------------------------------------------------------------------------------------------------------------------------------------

            break // each time we find a prospect we break out of the for loop --- if we found any prospects using the current index value we break 

        } // end of if :: if largerPerfectProduct > smallerPerfectProductOnce*workPiece  //  we only handle reads that were big enough to be prospects 
 
    } // this is the end of the afforementioned for loop that we break out of each time we have found a prospect and handled it 

} // the end of the readTheTableOfPP func that gets called 380,000 times


// Build a table of 825,000 pairs of PPs with their roots, does either squares or cubes:
//--------------------------------------------------------------------------------------

func buildTableOfPerfectProductsC(radical_index int) {  // C suffix denotes case 180: 

    var PerfectProduct int 
    Table_of_perfect_Products = nil // this fixed my bug 
    root := 10 
            iter := 0
        for iter < 825000 { // a table of 825,000 pairs of PPs with their roots ought to do it !!
            iter++
            root++ 
            if radical_index == 3 { // build an array of perfect cubes 
                PerfectProduct = root*root*root 
            }
            if radical_index == 2 { // build an array of perfect squares 
                PerfectProduct = root*root
            }
            Table_of_perfect_Products = append(Table_of_perfect_Products, PerfectProduct ) 
            Table_of_perfect_Products = append(Table_of_perfect_Products, root ) // the root of the prior PP
        }
}
// end of case 180 



// case 18: 
func xRootOfy(selection int) { // calculates either square or cube root of any integer 

    var index = 0 // counter used in the for loop in this func :: is also passed to the pricipal func readTheTableOfPP // cannot be a global

        startFromTop := time.Now()

    radical_index := getInputFromUserAndSetPrecision() 

    buildTableOfPerfectProducts(radical_index) // 800,000 entries, 400,000 pairs 

// The following section consists of the principal for loop with a conditional break ------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------------------------------------------------------

        startBeforeCall := time.Now()

    for index < 400000 { // the table has 825,000 entries, > 410,000 pairs; so index increments by 2 at the bottom of this loop (200,000 iters)

        readTheTableOfPP(index, startBeforeCall, radical_index) // pass-in an index to the table: 380,000 indexs corresponding to the number of pairs of entries 

        handlePerfectSquaresAndCubes(startFromTop, radical_index, selection) // handle the rare case of a perfect square or cube 

                if diffOfLarger == 0 || diffOfSmaller == 0 {  // Then, it was a perfect square or cube; so, we need to ... 
                    break // ... out of the for loop because we are done : the workpiece was either a perfect square or a perfect cube
                }

        if index == 80000 {fmt.Println("\n80,000 ... still working ...")}    
        if index == 160000 {fmt.Println("\n160,000 ... still working ...")}                
        if index == 240000 {fmt.Println("\n240,000 ... still working ...")}                
        if index == 320000 {fmt.Println("\n320,000 ... still working, almost there ...\n")}                

        index = index + 2 // increment the index and read the table again 
    } // end of for loop // the above break statement is NOT the only way to exit this for loop, it also terminates after 380,000 iterations of index 

// All of the remaining sections are conditional for workpiece NOT being a perfect square or cube
if perfectResult2 == 0 && perfectResult3 == 0 {  // Then, it was NOT a perfect square or cube, so handle that case 
// the remaining sections are only reached after having exited the primary for loop above via a break statement or an exaustive reading of the table ------------
//---------------------------------------------------------------------------------------------------------------------------------------------------------------
        // calculate elapsed time 
        t_s2 := time.Now()
        elapsed_s2 := t_s2.Sub(startFromTop)


// the following sections log the final results to a text file (and also does one conditional Printf) -------------------------------------------------
//-----------------------------------------------------------------------------------------------------------------------------------------------------
        fileHandle, err31 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
            check(err31)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.

        Hostname, _ := os.Hostname()
        _ , err30 := fmt.Fprintf(fileHandle, "\n  -- %d root of %d by a ratio of perfect Products -- selection #%d on %s \n", radical_index, workPiece, selection, Hostname)
            check(err30)

        current_time := time.Now()
        _ , err36 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
            check(err36)

        // index = index 
        _ , err35 := fmt.Fprintf(fileHandle, "%d was total Iterations \n", index)
            check(err35)


    // Sort the slice sortedResults by its pdiff field : 
    //-----------------------------------------------------------------------------------------------------------
        sort.Slice(sortedResults, func(i, j int) bool { return sortedResults[i].pdiff < sortedResults[j].pdiff})

                                /*
                                // print the sorted slice twice; once for each field 
                                    fmt.Println("Here are the results:")
                                    resultCount := 1
                                    for _, result := range sortedResults {
                                        fmt.Printf("%d, %0.16f \n", resultCount, result.result)
                                        resultCount++
                                    }
                                    fmt.Println("And here are the p-diffs:")
                                    pdiffCount := 1
                                    for _, result := range sortedResults {
                                        fmt.Printf("%d, %0.16f \n", pdiffCount, result.pdiff)
                                        pdiffCount++
                                    }
                                */ 

    // display and print the best-fitting result based solely on the lowest pdiff :
    //-----------------------------------------------------------------------------

        // display the best fitting result :
            if radical_index == 2 {
                fmt.Printf("%0.9f, is the best approximation for the Square Root of %d \n", sortedResults[0].result, workPiece)
            }
            if radical_index == 3 {
                fmt.Printf("%0.9f, is the best approximation for the  Cube  Root of %d \n", sortedResults[0].result, workPiece)
            }

            // Fprint/log the best fitting result :
                if radical_index == 2 {
                    _ , err48 := fmt.Fprintf(fileHandle, "%0.9f, is the best approximation for the Square Root of %d \n", sortedResults[0].result, workPiece)
                    check(err48)
                }
                if radical_index == 3 {
                    _ , err49 := fmt.Fprintf(fileHandle, "%0.9f, is the best approximation for the  Cube  Root of %d \n", sortedResults[0].result, workPiece)
                    check(err49)
                }
    //-----------------------------------------------------------------------------


        TotalRun := elapsed_s2.String() // cast time durations to a String type for Fprintf "formatted print"
        _ , err57 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
            check(err57)

        fileHandle.Close()


                    /*
                    // all this crap with these 3 arrays is cute, but it does not sort as a record with fields, so it is not what I want 
                                array_len := len(List_of_2_results_case18)
                                _ , err8 := fmt.Fprintf(fileHandle, "%d was len of array \n", array_len)
                                    check(err8)
                                if array_len > 0 {
                                    index := 0
                                    for array_len > 0 {
                                        result_from_array := List_of_2_results_case18[index]
                                        array_len--
                                         _ , err9 := fmt.Fprintf(fileHandle, "%0.16f with a diff of %d, percent diff of %0.4f percent\n", 
                                            result_from_array, corresponding_diffs[index], diffs_as_percent[index]*100000)
                                                check(err9)
                                        index++
                                    }
                                }
                                List_of_2_results_case18 = nil 
                                corresponding_diffs = nil 
                    */
}
}


func handlePerfectSquaresAndCubes(startFromTop time.Time, radical_index int, selection int) {
// the next sub-section detects, traps, and reports the detection of either a perfect square of a perfect cube ------------------
// ... it also is responsible for causing the algorithm to terminate via a break if workpiece was a perfect square or cube 
//-------------------------------------------------------------------------------------------------------------------------------
        if diffOfLarger == 0 || diffOfSmaller == 0 {  // Then, it was a perfect square or cube 

            t_s1 := time.Now() 
            elapsed_s1 := t_s1.Sub(startFromTop)

                fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) 
                    check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                    defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.

                Hostname, _ := os.Hostname()
                _ , err0 := fmt.Fprintf(fileHandle, "\n  -- %d root of %d by a ratio of PerfectProducts -- selection #%d on %s \n", 
                    radical_index, workPiece, selection, Hostname)
                    check(err0)

                current_time := time.Now()
                _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                check(err6)

                TotalRun := elapsed_s1.String() // cast time durations to a String type for Fprintf "formatted print"
                _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
                    check(err7)

                    if radical_index == 2 {
                        _ , err8 := fmt.Fprintf(fileHandle, "the %d root of %d is %0.2f \n", radical_index, workPiece, perfectResult2)
                            check (err8)
                    }
                    if radical_index == 3 {
                        _ , err38 := fmt.Fprintf(fileHandle, "the %d root of %d is %0.2f \n", radical_index, workPiece, perfectResult3)
                            check (err38)
                    }

                    fileHandle.Close()

            // break // break out of the for loop because we are done : the workpiece was either a perfect square or a perfect cube 

        } // end of if :: if it was a perfect square or cube 
//-------------------------------------------------------------------------------------------------------------------------------

}

func readTheTableOfPP(index int, startBeforeCall time.Time, radical_index int) {  // this gets called 380,000 times. 

    // The first time it is called index is 0

    // read it ...
    smallerPerfectProductOnce := Table_of_perfect_Products[index]  
    // ... and save it locally, do this just-once per func call. // ... index may be 0 up to 380,000
    RootOfsmallerPerfectProductOnce := Table_of_perfect_Products[index+1]
    // ^^^ also read the root wich corresponds

        iter := 0 
    for iter < 410000 { // 410,000 loops. Why do we need so many?, Because we need to read through 825,000 table entries pairs
        iter++         //  ... iters are therefore half the number of pairs. There are actually 1,600,000 items, but who's counting?
        index = index + 2
        largerPerfectProduct := Table_of_perfect_Products[index]  
        // to approximate the root of an imperfect square x we will need a ratio of two perfect squares wich is about equal to x
        // ...we need to find two perfect squares such that one is about x times larger than the other
        // get next perfect square from table for testing to see if it is more than x * bigger than smallerPerfectProductOnce

        if largerPerfectProduct > smallerPerfectProductOnce*workPiece {  
        // if largerPerfectProduct is a candidate based on it being just-a-bit larger than workPiece* the smaller PP deal with that, else loop to the next potential 

            ProspectiveHitOnLargeSide := largerPerfectProduct  // make a copy under a more suitable name :) 
            rootOfProspectiveHitOnLargeSide := Table_of_perfect_Products[index+1] // the current value of index plus one holds the root of largerPerfectSquare hence the root of ProspectiveHitOnLargeSide

            ProspectiveHitOnSmallerSide := Table_of_perfect_Products[index-2]  
        // save that smaller one too //                               ^^ 2 now instead of 1 because we have added roots to the slice
            rootOfProspectiveHitOnSmallerSide := Table_of_perfect_Products[index-1]

            diffOfLarger = ProspectiveHitOnLargeSide - workPiece*smallerPerfectProductOnce
                //diffOfSmaller = -(ProspectiveHitOnSmallerSide - workPiece*smallerPerfectProductOnce) // this was dumb ?? 
            diffOfSmaller = workPiece*smallerPerfectProductOnce - ProspectiveHitOnSmallerSide

// detect perfect squares and set global vars to their roots -----------------------------------------------
            if diffOfLarger == 0 {
                fmt.Println(string(colorCyan), "\n The", radical_index, "root of", workPiece, "is", string(colorGreen), 
                float64(rootOfProspectiveHitOnLargeSide) / float64(RootOfsmallerPerfectProductOnce), string(colorReset), "\n")

                perfectResult2 = (math.Sqrt(float64(workPiece))) // these global values are used later to Fprint to a log file
                perfectResult3 = math.Cbrt(float64(workPiece)) 
                break // out of the for loop because the workPiece is itself a perfect square
            }
            if diffOfSmaller == 0 {
                fmt.Println(string(colorCyan), "\n The", radical_index, "root of", workPiece, "is", string(colorGreen), 
                float64(rootOfProspectiveHitOnSmallerSide) / float64(RootOfsmallerPerfectProductOnce), string(colorReset), "\n")

                perfectResult2 = (math.Sqrt(float64(workPiece))) // these global values are used later to Fprint to a log file
                perfectResult3 = math.Cbrt(float64(workPiece)) 
                break // out of the for loop because the workPiece is itself a perfect square
            }
// ---------------------------------------------------------------------------------------------------------
// in case 18:

// larger side section: ----------------------------------------------------------------------------------------------------------------------------------------
//--------------------------------------------------------------------------------------------------------------------------------------------------------------

            if diffOfLarger < precisionOfRoot {  // report the prospects, their differences, and the calculated result for the Sqrt or Cbrt 
                fmt.Println("small PP is", string(colorCyan), smallerPerfectProductOnce, string(colorReset), "and, slightly on the higher side of", workPiece, 
                "* that we found a PP of", string(colorCyan), ProspectiveHitOnLargeSide, string(colorReset), "a difference of", diffOfLarger)

                fmt.Println("the ", radical_index, " root of ", workPiece, " is calculated as ", string(colorGreen), 
                float64(rootOfProspectiveHitOnLargeSide) / float64(RootOfsmallerPerfectProductOnce), string(colorReset))

                fmt.Printf("with pdiff of %0.4f \n", (float64(diffOfLarger)/float64(ProspectiveHitOnLargeSide))*100000)

                // save the result to an accumulator array so we can Fprint all such hits at the very end
                //List_of_2_results_case18 = append(List_of_2_results_case18, float64(rootOfProspectiveHitOnLargeSide) / float64(RootOfsmallerPerfectProductOnce) )
                //corresponding_diffs = append(corresponding_diffs, diffOfLarger)
                //diffs_as_percent = append(diffs_as_percent, float64(diffOfLarger)/float64(ProspectiveHitOnLargeSide))

                // in the next five lines we load (append) a record into/to the file (array) of Results 
                Result1 := Results{
                    result: float64(rootOfProspectiveHitOnLargeSide) / float64(RootOfsmallerPerfectProductOnce),
                    pdiff: float64(diffOfLarger)/float64(ProspectiveHitOnLargeSide),
                }
                sortedResults = append(sortedResults, Result1)

                t2 := time.Now()
                elapsed2 := t2.Sub(startBeforeCall)
                // if needed, notify the user that we are still working 
                    Tim_win = 0.178
                    if radical_index == 3 {
                        if workPiece > 13 {
                            Tim_win = 0.0012 
                        } else {
                        Tim_win = 0.003 
                        }                    }
                if float64(elapsed2.Seconds()) > Tim_win {
                    fmt.Println(float64(elapsed2.Seconds()), "Seconds have elapsed ... working ...\n")
                }
            }
//---------------------------------------------------------------------------------------------------------------------------------------------------------------

// in case 18:
// smaller side section: ----------------------------------------------------------------------------------------------------------------------------------------
//---------------------------------------------------------------------------------------------------------------------------------------------------------------
            if diffOfSmaller < precisionOfRoot {  // report the prospects, their differences, and the calculated result for the Sqrt or Cbrt 
                fmt.Println("small PP is", string(colorCyan), smallerPerfectProductOnce, string(colorReset), "and, slightly on the lesser side of", workPiece, 
                "* that we found a PP of", string(colorCyan), ProspectiveHitOnSmallerSide, string(colorReset), "a difference of", diffOfSmaller)

                fmt.Println("the ", radical_index, " root of ", workPiece, " is calculated as ", string(colorGreen), 
                float64(rootOfProspectiveHitOnSmallerSide) / float64(RootOfsmallerPerfectProductOnce), string(colorReset))

                fmt.Printf("with pdiff of %0.4f \n", (float64(diffOfSmaller)/float64(ProspectiveHitOnSmallerSide))*100000)

                // save the result to three accumulator arrays so we can Fprint all such hits, diffs, and p-diffs, at the very end of run 
                //List_of_2_results_case18 = append(List_of_2_results_case18, float64(rootOfProspectiveHitOnSmallerSide) / float64(RootOfsmallerPerfectProductOnce) )
                //corresponding_diffs = append(corresponding_diffs, diffOfSmaller)
                //diffs_as_percent = append(diffs_as_percent, float64(diffOfSmaller)/float64(ProspectiveHitOnSmallerSide))

                // in the next five lines we load (append) a record into/to the file (array) of Results 
                Result1 := Results{
                    result: float64(rootOfProspectiveHitOnSmallerSide) / float64(RootOfsmallerPerfectProductOnce),
                    pdiff: float64(diffOfSmaller)/float64(ProspectiveHitOnSmallerSide),
                }
                sortedResults = append(sortedResults, Result1)

                t2 := time.Now()
                elapsed2 := t2.Sub(startBeforeCall)
                // if needed, notify the user that we are still working 
                    Tim_win = 0.178
                    if radical_index == 3 {
                        if workPiece > 13 {
                            Tim_win = 0.0012 
                        } else {
                        Tim_win = 0.003 
                        } 
                    }   
                if float64(elapsed2.Seconds()) > Tim_win  {
                    fmt.Println(float64(elapsed2.Seconds()), "Seconds have elapsed ... working ...\n")
                }
            }    // end of if
                //------------------------------------------------------------------------------------------------------------------------------------------

            break // each time we find a prospect we break out of the for loop --- if we found any prospects using the current index value we break 

        } // end of if :: if largerPerfectProduct > smallerPerfectProductOnce*workPiece  //  we only handle reads that were big enough to be prospects 
 
    } // this is the end of the afforementioned for loop that we break out of each time we have found a prospect and handled it 

} // the end of the readTheTableOfPP func that gets called 380,000 times


func getInputFromUserAndSetPrecision() int {
    radical_index := 2 // we set it to 2 becuase if the user just hits enter we want radical_index set as 2 by default 
    skip_redoing_loop = 0 
    for radical_index > 1 && skip_redoing_loop == 0 { 
        // radical_index will initially be 2, and the skip flag will be initially 0, therefore we will initially enter this for loop 
        fmt.Println("\n\nEnter 2 for SquareRoot or 3 for CubeRoot")
        fmt.Scan(&radical_index) // just hitting enter will leave radical_index set as 2  // same in Unix or Windows variants 

        for radical_index > 0 { // but, if we entered any value for radical_index 
            if radical_index < 2 || radical_index > 3 { // only then do we check to see if it is a 2 or a 3
                fmt.Println(string(colorPurple), 
                "\n No, for the radical, enter either 2 or 3, just a 2 or a 3, \n ... no other number for the radical please \n", string(colorReset))
                radical_index = 2 // an entered 2 or 3 would have bypassed this, but this value of 2 is needed to re-enter/continue the outer for loop 
                break // out of inner loop only, and with skip flag still set to 0 
            } else { // if a 2 or a 3 was entered go ahead and prompt for a workPiece 
                    for 1 == 1 {
                        fmt.Println("\n\nEnter any integer that you wish to find the", radical_index, "Root of")
                        fmt.Scan(&workPiece) 
    // if a float was entered earlier for radical_index, the fractional part will be assigned to workPiece, without prompting (which is buggy on the part of Go)
                        // the above is the same in Unix or Windows variants 
                        if workPiece == 0 || workPiece == 1 {
                            fmt.Println(string(colorRed), "\n You must enter an integer greater than one", string(colorReset))
                        } else {
                            break // prompt for a valid workPiece and loop back to re scan, or break out of said loop herein
                        }
                    }
                skip_redoing_loop = 1 
                break // thus sending us back to the outer loop with the directive to skip redoing it per "skip_redoing_loop = 1" on previous line
            }
        }
        if radical_index < 1 { // 
            fmt.Println(string(colorPurple), 
            "\n Zero is dopey, for the radical, enter either 2 or 3, just a 2 or a 3, \n ... no other number for the radical please \n", string(colorReset))
            radical_index = 2 // go fish 
        }
    }

    if radical_index == 3 { // if doing a cube root special tolerances are set here for certain problem values, i.e., 2, 11, 17, 3, 4, or 14
            if workPiece > 4 {
                precisionOfRoot = 1700
                fmt.Println("\n Default precision is 1700 \n")
            }
            if workPiece == 2 || workPiece == 11 || workPiece == 17 {
                precisionOfRoot = 600
                fmt.Println("\n resetting precision to 600 \n")
            }
            if workPiece == 3 || workPiece == 4 || workPiece == 14 {
                precisionOfRoot = 900
                fmt.Println("\n resetting precision to 900 \n")
            }
    }
    if radical_index == 2 { // if doing a square root we just use a tolerance of 4 for all workpieces. 
        precisionOfRoot = 4
    }
return radical_index 
}


// Build a table of 825,000 pairs of PPs with their roots, does either squares or cubes:
//--------------------------------------------------------------------------------------

func buildTableOfPerfectProducts(radical_index int) {

    var PerfectProduct int 
    Table_of_perfect_Products = nil // this fixed my bug 
    root := 10 
            iter := 0
        for iter < 825000 { // a table of 825,000 pairs of PPs with their roots ought to do it !!
            iter++
            root++ 
            if radical_index == 3 { // build an array of perfect cubes 
                PerfectProduct = root*root*root 
            }
            if radical_index == 2 { // build an array of perfect squares 
                PerfectProduct = root*root
            }
            Table_of_perfect_Products = append(Table_of_perfect_Products, PerfectProduct ) 
            Table_of_perfect_Products = append(Table_of_perfect_Products, root ) // the root of the prior PP
        }
}





func BBPF(selection int) { // case 2: 
    fmt.Println("\nYou selected #", selection, "the Bailey–Borwein–Plouffe formula for π, circa 1995\n")
    fmt.Println("How many digits of pi should we calculate? Enter one integer '4 to 16' ")

          var numAi float64 // aligned with the value below
    fmt.Scan(&numAi)

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




func Nilakantha(selection int){  // case 5: 
    fmt.Println("\n\nYou selected #", selection, " a series by Indian astronomer Nilakantha Somayaji circa 1500 AD")
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

    // store reults in a log file which can be displayed from within the program by selecting option #12
                fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                    check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                    defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                Hostname, _ := os.Hostname()
                _ , err0 := fmt.Fprintf(fileHandle, "\n  -- Nilakantha Somayaji -- selection #%d on %s \n", selection, Hostname)
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
                fmt.Println(" per option #", selection, "  --  the Nilakantha Somayaji series, circa 1500 AD\n")
                fmt.Println("Select 12 at menu to display prior results")
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
        for iterInt64 < 50000000000 {
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
            if iterInt64 == 1000000000 {
                fmt.Println("    1 2345678901234 ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.1415926535897,93 is, again, the value of π from the web") 
                fmt.Println("    1 2345678901234 56 :: counting the first 16 actual digits of π")
                t := time.Now() ; elapsed := t.Sub(start)
                fmt.Println(iterInt64, " iterations were completed in ", elapsed, " yielding still just 14 digits of π\n") 
                fmt.Println(string(colorCyan), "  ... working ...\n", string(colorReset))
            }
            if iterInt64 == 5000000000 {
                fmt.Println("    1 2345678901234 ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.1415926535897,93 is, again, the value of π from the web") 
                fmt.Println("    1 2345678901234 56 :: counting the first 16 actual digits of π")
                t := time.Now() ;  elapsed := t.Sub(start)
                fmt.Println(iterInt64, " 5 billion iterations were completed in ", elapsed, " yielding still just 14 digits of π\n") 
                fmt.Println(string(colorCyan), "  ... working ...\n", string(colorReset))
            }            
            if iterInt64 == 9000000000 {
                fmt.Println("    1 2345678901234 ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.1415926535897,93 is, again, the value of π from the web") 
                fmt.Println("    1 2345678901234 56 :: counting the first 16 actual digits of π")
                t := time.Now() ;  elapsed := t.Sub(start)
                fmt.Println(iterInt64, " 9 billion iterations were completed in ", elapsed, " yielding still just 14 digits of π\n") 
                fmt.Println(string(colorCyan), "  ... working ...\n", string(colorReset))
            }
            if iterInt64 == 50000000000 {  // no additional digit are obtained even after 50 billion iterations
                fmt.Println("    1 2345678901234567 ")
                fmt.Println("   ", firstsum, " is what we have calculated per Nilakantha Somayaji") 
                fmt.Println("    3.141592653589793 is, again, the value of π from the web") 
                fmt.Println("    1 2345678901234567 :: counting the first 16 actual digits of π")
                t := time.Now() ;  elapsed := t.Sub(start)
                fmt.Println(iterInt64, " 50 billion iterations were completed in ", elapsed, " yielding still just 14 digits of π\n") 
            }
        } // second for loop

                t := time.Now() ;  elapsed := t.Sub(start)

            // store reults in a log file which can be displayed from within the program by selecting option #12
                fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                    check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                    defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                Hostname, _ := os.Hostname()
                _ , err0 := fmt.Fprintf(fileHandle, "\n  -- 50B continued Nilakantha Somayaji -- selection #%d on %s \n", selection, Hostname)
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

    } // my options50 if
}


func GottfriedWilhelmLeibniz(selection int){  // case 6: 
    fmt.Println("\n\nYou selected #", selection, " Gottfried Wilhelm Leibniz formula  :  π = 4 * ( 1 - 1/3 + 1/5 - 1/7 + 1/9 ... ")
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
                    fmt.Println(" per option #", selection, "  --  the Gottfried Wilhelm Leibniz formula\n")

                    LinesPerIter = 14
                    fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                    LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                    fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                    fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")

    // store reults in a log file which can be displayed from within the program by selecting option #12
                    fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                        check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                        defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                    Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  -- Gottfried Wilhelm Leibniz -- selection #%d on %s \n", selection, Hostname)  
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
} // end of first for loop

        fmt.Println(string(colorGreen), "Enter any positive digit to continue with an additional 9 billion iterations, 0 to exit", string(colorReset))
        option9 := 0
        fmt.Scan(&option9)

    if option9 > 0 {

        fmt.Println(string(colorCyan), "\n\nYou elected to continue the Gottfried Wilhelm Leibniz formula  :  π = 4 * ( 1 - 1/3 + 1/5 - 1/7 + 1/9 ... ")
        fmt.Println("    π = 3 + 4/(2*3*4) - 4/(4*5*6) + 4/(6*7*8) - 4/(8*9*10) + 4/(10*11*12) ...", string(colorReset))

        fmt.Println("   Infinitesimal calculus was developed independently in the late 17th century by Isaac Newton")
        fmt.Println("    ... and Gottfried Wilhelm Leibniz")
        fmt.Println("    9 billion iterations will be executed \n\n", string(colorCyan), "   ... working ...\n", string(colorReset))

        start := time.Now()
/*        iterFloat64 = 0
        var denom float64
            denom = 3
        var sum float64
        iterInt64 = 1
*/
//            sum = 1-(1/denom)

    for iterInt64 < 9000000000 {
        iterFloat64++
        iterInt64++
        denom = denom + 2
            if iterInt64 % 2 == 0 {
                sum = sum + 1/denom 
            } else { 
                sum = sum - 1/denom 
            }
            π = 4 * sum 

                    if iterInt64 == 6000000000 {fmt.Println("... 6 Billion completed ... \n")
                    fmt.Println("   #1 234567890#")
                    fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                    fmt.Println("    3.141592653,589793 is from the web") 
                    fmt.Println("   #1 234567890 123456# :: counting the first 16 actual digits of π")
                    t := time.Now()
                    elapsed := t.Sub(start)
                    fmt.Println("  6,000,000,000 iterations in ", elapsed, " still yields 10 digits of π\n")
                    fmt.Println(string(colorCyan), "  ... working ...\n", string(colorReset))
                    }
                    if iterInt64 == 8000000000 {fmt.Println("... 8 Billion completed. still ten ...\n")
                    fmt.Println("   #1 234567890#")
                    fmt.Println("   ",π, "was calculated by the Gottfried Wilhelm Leibniz formula") 
                    fmt.Println("    3.141592653,589793 is from the web") 
                    fmt.Println("   #1 234567890 123456# :: counting the first 16 actual digits of π")
                    t := time.Now()
                    elapsed := t.Sub(start)
                    fmt.Println("  8,000,000,000 iterations in ", elapsed, " still yields 10 digits of π\n")
                    fmt.Println(string(colorCyan), "  ... working ...\n", string(colorReset))
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
                    fmt.Println(" per option #", selection, "  --  the Gottfried Wilhelm Leibniz formula\n")

                    t = time.Now()
                    elapsed = t.Sub(start)

                    LinesPerIter = 14
                    fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                    LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                    fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                    fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")

    // store reults in a log file which can be displayed from within the program by selecting option #12
                    fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                        check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                        defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                    Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  -- Gottfried Wilhelm Leibniz (cont.) -- selection #%d on %s \n", selection, Hostname)  
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
        } // end of second for
    } // end of option9 if
}

func GregoryLeibniz(selection int){  // case 7: 
    fmt.Println("\n\nYou selected #", selection, " the Gregory-Leibniz series ...")
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
// 7 digits of Pi were found per the above code
// the next two ifs give eight digits of Pi
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
// 9 digits of Pi are found below 
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
                fmt.Println(" per option #", selection, "  --  the Gregory-Leibniz series, circa 1676\n")

                LinesPerIter = 11  // an estimate of the number of lines per iteration 
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("       %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")

    // store reults in a log file which can be displayed from within the program by selecting option #12
                fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                    check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                    defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                Hostname, _ := os.Hostname()
                _ , err0 := fmt.Fprintf(fileHandle, "\n  -- Gregory-Leibniz -- selection #%d on %s \n", selection, Hostname)
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





func ArchimedesBig (selection int) { // case 14: 
    fmt.Println("\nYou selected #", selection, "  --  An improved version of Archimedes' method")
        start := time.Now()
    r := big.NewFloat(1)              // radius is a constant 1
    s1 := big.NewFloat(1)             // starts out as 1
    numberOfSides := big.NewFloat(6)  // the number of sides of the polygon

    a := new(big.Float)         // height of bisected triangle
    b := new(big.Float)         // new short side
    p := new(big.Float)         // perimeter
    s2 := new(big.Float)        // new hypotenuse / new side
    p_d := new(big.Float)       // pi
    s1_2 := new(big.Float)      // s1/2

    // Set the precision to a higher value
    precision := 55000 
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
    numberOfSides.Mul(numberOfSides, big.NewFloat(2)) // doubleing the number of sides was not orriginally done here ????????????????

    s1_2.Quo(s1, big.NewFloat(2))  // s1_2 = s1/2
    a.Sqrt(new(big.Float).Sub(r, new(big.Float).Mul(s1_2, s1_2)))  // height of bisected triangle
    //     create.new.obj     1 - create.new.obj = same as squaring below 
    // a = math.Sqrt(r-(math.Pow(s1_2,2)))                         // height of bisected triangle
    b.Sub(r, a)                             // new short side, r is initially 1
    // b = 1-a                              // new short side
    s2.Sqrt(new(big.Float).Add(new(big.Float).Mul(b, b), new(big.Float).Mul(s1_2, s1_2))) // new hypotenuse / NewSide
    // s2 = math.Sqrt(math.Pow(b,2) + (math.Pow(s1_2,2)))                                 // new hypotenuse / NewSide

    s1.Set(s2) // where is the corresponding s1 = s2 ???????????????????????????????????????????????????????????????????????????????

    p.Mul(numberOfSides, s1)
    // p = numberOfSides*s1 
    p_d.Set(p) // p_d should instead be set to p/2 (half p) ????????????????????????????????????????????????????????????????????????
    // p_d = p/2


    var iterInt64 int64 = 0
    for iterInt64 < 5000 {
        iterInt64++
        numberOfSides.Mul(numberOfSides, big.NewFloat(2))                    // double the number of sides 
        s1_2.Quo(s1, big.NewFloat(2))                                        // halve s1 
        a.Sqrt(new(big.Float).Sub(r, new(big.Float).Mul(s1_2, s1_2)))        // height of bisected triangle
        b.Sub(r, a)                                                          // new short side
        s2.Sqrt(new(big.Float).Add(new(big.Float).Mul(b, b), new(big.Float).Mul(s1_2, s1_2))) // new hypotenuse / NewSide
        s1.Set(s2)
        p.Mul(numberOfSides, s1)
        p_d.Set(p)
        p_d.Quo(p_d, big.NewFloat(2)) // Halve the value of p_d

        if iterInt64 == 24 { // last inter of prior version was 25 ?? 
            fmt.Println("\n   #1 234567890123456# :: counting the first 16 actual digits of π")
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
            fmt.Printf("    %.33f is the big.Float of what we have calculated per Archimedes' at 50 iters, 33f\n", p_d)
            fmt.Println("    3.141592653589793238462643383279502 is the value of π from the web")
            fmt.Println("   #1 234567890123456789012345678901# :: counting the actual digits of π")
            fmt.Println(iterInt64, " iterations were completed yielding 31 correct digits of π")
            fmt.Printf("the above was estimated from a %.0f sided polygon (formatted as a .0f) \n\n", numberOfSides)
        }

if iterInt64 == 150 {
fmt.Println(" #1 23456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123#")
fmt.Printf("  %.95f per Archimedes'\n", p_d)
fmt.Println("  3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211 is from web")
fmt.Println(" #1 23456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123#")
fmt.Println("           10        20         30       40        50         60        70        80        90 ")
    fmt.Println(iterInt64, " iterations were completed yielding 92 correct digits of π")
    fmt.Printf("Calculated from a %.0f sided polygon\n\n", numberOfSides)
}

if iterInt64 == 200 {
fmt.Println(" #1 23456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890#")
fmt.Printf("  %.122f per Archimedes'\n", p_d)
fmt.Println("  3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706798214808651328230664709 is from web")
fmt.Println(" #1 23456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890#")
fmt.Println("           10        20         30       40        50         60        70        80        90      100        110      120")
    fmt.Println(iterInt64, " iterations were completed yielding 121 correct digits of π")
    fmt.Printf("Calculated from a %.0f sided polygon\n\n\n", numberOfSides)
    fmt.Println(" ... working ...\n\n")
}

if iterInt64 == 1200 { fmt.Println("... still working, 1200 iterations completed ...\n")}
if iterInt64 == 2200 { fmt.Println("... still working, 2200 iterations completed ...\n")}
if iterInt64 == 3200 { fmt.Println("... still working, 3200 iterations completed ...\n")}
if iterInt64 == 4200 { fmt.Println("... still working, 4200 iterations completed ...\n")}


        if iterInt64 == 5000 { // was 5500
            fmt.Printf("  %.2713f per Archimedes'\n", p_d)  // show the first 2,713 digits of calculated pi
            fmt.Println(iterInt64, " iterations were completed, \n ... which generated a ", numberOfSides, "sided polygon!!\n")
            fmt.Println(iterInt64, " iterations were completed yielding 2,712 correct digits of π!!!\n")
            fmt.Println("Go's math/big objects were set to a precision value of:", precision)
        }
    }
        t := time.Now()
        elapsed := t.Sub(start)
        // store reults in a log file which can be displayed from within the program by selecting option #12
                    fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                        check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                        defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                    Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  -- Archimedes of Syracuse -- selection #%d on %s \n", selection, Hostname)
                        check(err0)
                    current_time := time.Now()
                    _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                    check(err6)

                    TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                    _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
                        check(err7)

                fmt.Println("\n   -- After hitting Return for menu redisplay, enter '11' for the derivation and proof of the Pythagorean\n\n")
                fmt.Println("Select 12 at menu to display prior results")

}




func JohnWallis(selection int){  // case 8: 
    fmt.Println("\n   You selected #", selection, " A Go language exercize which can be used to test the speed of your hardware.")
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
      for iterInt64 < 1000000000 {
        iterInt64++
        iterFloat64++
        numerators = numerators + 2
        firstDenom = firstDenom + 2
        secondDenom = secondDenom + 2
        cumulativeProduct = cumulativeProduct * (numerators/firstDenom) * (numerators/secondDenom)
        π = cumulativeProduct * 2  
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
                fmt.Println("\n   #1 234# ")
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
// a brief Red notification follows :
                fmt.Println(string(colorRed), " ... will be working on doing Billions more iterations ...\n\n", string(colorReset))
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
            fmt.Println(" per option #", selection, "  --  an infinite series by John Wallis circa 1655\n") // ----------------------

                LinesPerIter = 36 // an estimate 
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")

    // store reults in a log file which can be displayed from within the program by selecting option #12
                fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                    check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                    defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                Hostname, _ := os.Hostname()
                _ , err0 := fmt.Fprintf(fileHandle, "\n  -- John Wallis -- selection #%d on %s \n", selection, Hostname)
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
} // end of first for loop

        fmt.Println(string(colorGreen), "Enter any positive digit to continue with an additional 39 billion iterations, 0 to exit", string(colorReset))
        option39 := 0
        fmt.Scan(&option39)

    if option39 > 0 {

        fmt.Println(string(colorCyan), "\n\nYou elected to continue the infinite series by John Wallis", string(colorReset))
        fmt.Println("\n    an additionl 39 billion iterations will be executed \n\n", string(colorCyan), "   ... working ...\n", string(colorReset))

fmt.Println(string(colorRed), " ... still working ... on Billions of iterations, 39 to go ...\n", string(colorReset))

fmt.Println("\n ... 39 Billion additional loops now ensue, just to get one additional digit, the tenth digit of pi")

        start := time.Now()

      for iterInt64 < 40000000000 { 
        iterInt64++
        iterFloat64++
        numerators = numerators + 2
        firstDenom = firstDenom + 2
        secondDenom = secondDenom + 2
        cumulativeProduct = cumulativeProduct * (numerators/firstDenom) * (numerators/secondDenom)
        π = cumulativeProduct * 2  

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
        if iterInt64 == 40000000000 {  // 40B completed 
            fmt.Println("   #1 234567890# ")
            fmt.Println("   ", π, " calculated using an infinite series by John Wallis circa 1655") 
            fmt.Println("    3.141592653,589793 is the value of π from the web") 
            fmt.Println("   #1 234567890 123456# :: counting the first 16 actual digits of π")

            t := time.Now()
            elapsed := t.Sub(start)
            fmt.Println("Forty Billion iterations were completed in ", elapsed, " yielding π to 10 digits\n") 
            fmt.Println(" per option #", selection, "  --  an infinite series by John Wallis circa 1655\n") // ----------------------
                LinesPerIter = 36 // an estimate 
                fmt.Println("at aprox", LinesPerIter, "lines of code per iteration ...")
                LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
                fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond) 
                fmt.Println("      1000000000 is one billion, for comparison in liew of commas \n\n")

    // store reults in a log file which can be displayed from within the program by selecting option #12
                fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                    check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                    defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                Hostname, _ := os.Hostname()
                _ , err0 := fmt.Fprintf(fileHandle, "\n  -- John Wallis (cont.) -- selection #%d on %s \n", selection, Hostname)
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
    } // end of second for loop 
    } // end of 40B continuation if 
}


func EulersNumber(selection int){ // case 9: 
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
                   // change to cyan 
        fmt.Println(string(colorCyan), "An account starts with $1.00 and pays 100 percent interest per year. If the interest is credited once, ")
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
fmt.Println(Ricks_rune_Paragraph, string(colorReset))
}


func ErdosBorwein(selection int){  // case 10: 
    rune := `The Erdős–Borwein constant is the sum of the reciprocals of the Mersenne numbers. 
It is named after Paul Erdős and Peter Borwein. 

Paul Erdős was a Hungarian mathematician. He was one of the most prolific mathematicians 
and producers of mathematical conjectures of the 20th century. Erdős pursued and proposed 
problems in discrete mathematics, graph theory, number theory, mathematical analysis, 
approximation theory, set theory, and probability theory. Much of his work centered 
around discrete mathematics, cracking many previously unsolved problems in the field. 
He championed and contributed to Ramsey theory, which studies the conditions in which 
order necessarily appears. Overall, his work leaned towards solving previously open 
problems, rather than developing or exploring new areas of mathematics.

Erdős published around 1,500 mathematical papers during his lifetime, a figure that
remains unsurpassed. He firmly believed mathematics to be a social activity, 
living an itinerant lifestyle with the sole purpose of writing mathematical papers 
with other mathematicians. He was known both for his social practice of mathematics, 
working with more than 500 collaborators, and for his eccentric lifestyle; Time magazine 
called him "The Oddball's Oddball". He devoted his waking hours to mathematics, even 
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
paper that outlined and popularized a proof for computing one billion digits of π. 
The authors won the 1993 Chauvenet Prize and Merten M. Hasse Prize for this paper.

In 1993, he moved to Simon Fraser University, joining his brother Jonathan in establishing 
the Centre for Experimental and Constructive Mathematics (CECM) where he developed the 
Inverse Symbolic Calculator. `
fmt.Println("") // so as not to pad the following rune with an unwanted space 
    fmt.Println(string(colorCyan), rune, "\n", string(colorReset))
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



// and now for the spigot func itself ... (it's just so tight!)
func TheSpigot(){  // case 19: 
    var numberOfDigitsToCalc int
    fmt.Println("How much pi can you handle?")
    fmt.Println("How many digits of pi do you really want? Enter that number now:")  // prompt the user
        fmt.Scan(&numberOfDigitsToCalc)
        fmt.Println("")

    Spigot(numberOfDigitsToCalc)

    fmt.Print("\n\nThis trick was made possible by a bit of code I mooched off of GitHub ...\n\n")
    fmt.Println("... to view the code with attribution Enter '39' at either menu\n")
    fmt.Println("To view an explaination of how it works enter 99 at either menu\n")
}
//func Spigot(n int) string { // called by the previous func
func Spigot(n int) { // Rick's version does not return a string // called by the previous func
    start := time.Now() 
    pi := "" // allocate a string var "pi" 
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
                pi = delChar(pi, i-k) // delChar func follows
                pi = pi[:i-k] + strconv.Itoa(replaced) + pi[i-k:]
            }
            digitsHeld = 1
        default:
            digitsHeld = 1
        }
        // Rick's code :
        if i == 1 { fmt.Print(".") }   // insert the decimal point between the 3 and the 1 
            fmt.Print(strconv.Itoa(q)) // Rick's new method of displaying pi, one digit at a time
            if i == 5000 || i == 10000 || i == 20000 || i == 40000 || i == 80000 || i == 140000 || i == 200000 {
                    t := time.Now()
                    elapsed := t.Sub(start)
                    // log stats to a log file 
                    fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                        check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                        defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                    Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  -- partial Spigot in process -- selection #%d on %s \n", selection, Hostname)
                        check(err0)
                    current_time := time.Now()
                    _ , err6 := fmt.Fprint(fileHandle, "... running at: ", current_time.Format(time.ANSIC), "\n")
                        check(err6)
                    TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                    _ , err7 := fmt.Fprintf(fileHandle, "Runtime this far is %s \n", TotalRun) 
                        check(err7)
                    _ , err8 := fmt.Fprintf(fileHandle, "... while calculating Pi to %d digits, having completed %d digits\n", n, i)
                        check(err8)
            }
        // end Rick's code
        pi += strconv.Itoa(q)          // in orriginal method, entire string accumulator was being returned 
    }
    // Rick's code : 
    t := time.Now()
    elapsed := t.Sub(start)
    // log stats to a log file 
        fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
            check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
        Hostname, _ := os.Hostname()
        _ , err0 := fmt.Fprintf(fileHandle, "\n  -- Spigot -- selection #%d on %s \n", selection, Hostname)
            check(err0)
        current_time := time.Now()
        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
            check(err6)
        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n", TotalRun) 
            check(err7)
        _ , err8 := fmt.Fprintf(fileHandle, "To calculate Pi to %d digits\n", n)
            check(err8)
    // end Rick's code
    //return pi // prior code 
}
func delChar(s string, index int) string {
    tmp := []rune(s)
    return string(append(tmp[0:index], tmp[index+1:]...))
}





// case 40: 
// A concurrent computation of pi using Nilakantha's formula.
// by Diego Brener diegosilva13 on Github 
// ******* nifty scoreboard ***********************************
func nifty_scoreBoard (){ // case 40: 
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




func bbp_formula (selection int){ // case 42: 
iters_bbp := 1 
    start := time.Now()
    //numCPU := runtime.NumCPU()
    //runtime.GOMAXPROCS(numCPU)

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
        fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
            check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
        Hostname, _ := os.Hostname()
        _ , err0 := fmt.Fprintf(fileHandle, "\n  -- calculate pi using the bbp formula -- selection #%d on %s \n", selection, Hostname)
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


func Leibniz_method_one_billion_iters(selection int){ // case 44: 
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

// Rick's code follows
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
        fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
            check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
        Hostname, _ := os.Hostname()
        _ , err0 := fmt.Fprintf(fileHandle, "\n  -- Approximation of PI using Leibniz method -- selection #%d on %s \n", selection, Hostname)
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


// case 15: 

// Chudnovsky-1.go based on https://arxiv.org/pdf/1809.00533.pdf 
/*
The accuracy of the code "may" be improved by using a higher precision floating-point type, such as big.Float or math.BigDecimal.
*/
        // 1,000,000 digits requires 70516 loops, per the run on May 7 2023 at 10:30
        //  was run on: Sun May  7 08:50:23 2023
        //  Total run was 8h4m39.7847064s
        // AND THE CALCULATION WAS INDEPENDANTLY VERIFIED !!!!!!!!!!!

var digits int 
var precR uint 
var loops int 

func chud(selection int) {  // case 15:
    var pi *big.Float
        start := time.Now() // saved start time to be compared with end time t 
    fmt.Println("\n Richard invites you to enter the number of digits of pi to calculate per the Chudnovsky method: \n")
    fmt.Scanf("%d", &digits) 
    fmt.Scanf("%d", &digits) 
    if digits == 1 { // if the user asks for only one digit of pi a testing proc should be run instead 
        //TestCalcPi(*big.Float(pi1000))
        //TestCalcPi(pi1000) // I tried everything I could think of 
        os.Exit(1)
    } 
  //pi, _            = CalcPi(float64(digits))
    pi, precR, loops = CalcPi(float64(digits), start)
    fmt.Printf("%1.[1]*[2]f \n", digits, pi)
    fmt.Printf("\n loops were %d, and digits requested was %d", loops, digits)
    fmt.Println("\nprec is: ", precR)
        t := time.Now()
        elapsed := t.Sub(start)
        // only if 
            if int(elapsed.Seconds()) != 0 {
                    fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                        check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                        defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.

                Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                    check(err0)
                current_time := time.Now()
                    _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                    check(err6)
                // the whole pi is printed to the datalog file on the line below
                    _ , err8 := fmt.Fprintf(fileHandle, "pi was %1.[1]*[2]f \n", digits, pi) 
                        check(err8)
                // ... after printing the whole pi, some nice stats are appended to the file's log entry 
                    TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                    _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %d \n ", TotalRun, digits) 
                        check(err7)
            }
}

// CalcPi calculate Pi for n number of digits
func CalcPi(digits float64, start time.Time) (*big.Float, uint, int) {
    var i int 
    /**
    *   This is an implementation for https://en.wikipedia.org/wiki/Chudnovsky_algorithm
    *   "It can be improved using binary splitting http://numbers.computation.free.fr/Constants/Algorithms/splitting.html
    *   if we split it into two independent parts and simplify the formula for more details: 
    * https://www.craig-wood.com/nick/articles/pi-chudnovsky/""
    */
    n := int64(2 + int(float64(digits)/14.181647462))
    prec := uint(int(math.Ceil(math.Log2(10)*digits)) + int(math.Ceil(math.Log10(digits))) + 2)

    c := new(big.Float).Mul(
        big.NewFloat(float64(426880)),
        new(big.Float).SetPrec(prec).Sqrt(big.NewFloat(float64(10005))),
    )

    k := big.NewInt(int64(6))
    k12 := big.NewInt(int64(12))
    l := big.NewFloat(float64(13591409))
    lc := big.NewFloat(float64(545140134))
    x := big.NewFloat(float64(1))
    xc := big.NewFloat(float64(-262537412640768000))
    m := big.NewFloat(float64(1))
    sum := big.NewFloat(float64(13591409))

    pi := big.NewFloat(0)

    x.SetPrec(prec)
    m.SetPrec(prec)
    sum.SetPrec(prec)
    pi.SetPrec(prec)

    bigI := big.NewInt(0)
    bigOne := big.NewInt(1)

test := 1
i = 1
    for ; n > 0; n-- {
        i++
        // L calculation
        l.Add(l, lc)

        // X calculation
        x.Mul(x, xc)

        // M calculation
        kpower3 := big.NewInt(0)
        kpower3.Exp(k, big.NewInt(3), nil)
        ktimes16 := new(big.Int).Mul(k, big.NewInt(16))
        mtop := big.NewFloat(0).SetPrec(prec)
        mtop.SetInt(new(big.Int).Sub(kpower3, ktimes16))
        mbot := big.NewFloat(0).SetPrec(prec)
        mbot.SetInt(new(big.Int).Exp(new(big.Int).Add(bigI, bigOne), big.NewInt(3), nil))
        mtmp := big.NewFloat(0).SetPrec(prec)
        mtmp.Quo(mtop, mbot)
        m.Mul(m, mtmp)

        // Sum calculation
        t := big.NewFloat(0).SetPrec(prec)
        t.Mul(m, l)
        t.Quo(t, x)
        sum.Add(sum, t)

        // Pi calculation
        pi.Quo(c, sum)
        k.Add(k, k12)
        bigI.Add(bigI, bigOne)
     // if digits > 99000 we skip all the following step tests etc 
        if digits > 99000 { 
            //fmt.Println("\n this will run a while \n") // cannot actually do this or the message will repeat an ungodly amount of times
        } else {
        if i == 100 {
            fmt.Println("\n we are at 100 loops \n")
            fmt.Printf("%.999f", pi)
            fmt.Println("\n we are at 100 loops \n")
            fmt.Println("enter any int to continue, 0 to end")
            t := time.Now()
            elapsed := t.Sub(start)
            // only if 
                if int(elapsed.Seconds()) != 0 {
                        fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                            check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                        Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                            check(err0)
                    current_time := time.Now()
                        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                        check(err6)
                    // the partial pi is printed to the datalog file on the line below
                        _ , err8 := fmt.Fprintf(fileHandle, "a peak at 77 digits of pi was %1.77f \n", pi) 
                            check(err8)
                    // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                            check(err7)
                }
            fmt.Scanf("%d", &test) 
            fmt.Scanf("%d", &test) 
            if test == 0 {
                t := time.Now()
                elapsed := t.Sub(start)
                // only if 
                    if int(elapsed.Seconds()) != 0 {
                            fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                                check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                                defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                            Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                                check(err0)
                        current_time := time.Now()
                            _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                            check(err6)
                        // the partial pi is printed to the datalog file on the line below
                            _ , err8 := fmt.Fprintf(fileHandle, "a peak at 77 digits of pi was %1.77f \n", pi) 
                                check(err8)
                        // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                            TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                            _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                                check(err7)
                    }
                fmt.Printf("%.999f", pi)
                fmt.Println("\nprec is: ", precR)
                os.Exit(1)
            }
        }
        if i == 200 {
            fmt.Println("\n we are at 200 loops \n")
            fmt.Printf("%.9999f", pi)
            fmt.Println("\n we are at 200 loops \n")
            fmt.Println("enter any int to continue, 0 to end")
            t := time.Now()
            elapsed := t.Sub(start)
            // only if 
                if int(elapsed.Seconds()) != 0 {
                        fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                            check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                        Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                            check(err0)
                    current_time := time.Now()
                        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                        check(err6)
                    // the partial pi is printed to the datalog file on the line below
                        _ , err8 := fmt.Fprintf(fileHandle, "a peak at 87 digits of pi was %1.87f \n", pi) 
                            check(err8)
                    // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                            check(err7)
                }
            fmt.Scanf("%d", &test) 
            fmt.Scanf("%d", &test) 
            if test == 0 {
                t := time.Now()
                elapsed := t.Sub(start)
                // only if 
                    if int(elapsed.Seconds()) != 0 {
                            fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                                check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                                defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                            Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                                check(err0)
                        current_time := time.Now()
                            _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                            check(err6)
                        // the partial pi is printed to the datalog file on the line below
                            _ , err8 := fmt.Fprintf(fileHandle, "a peak at 87 digits of pi was %1.87f \n", pi) 
                                check(err8)
                        // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                            TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                            _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                                check(err7)
                    }
                fmt.Printf("%.9999f", pi)
                fmt.Println("\nprec is: ", precR)
                os.Exit(1)
            }
        }
        if i == 400 {
            fmt.Println("\n we are at 400 loops \n")
            fmt.Printf("%.49999f", pi)
            fmt.Println("\n we are at 400 loops \n")
            fmt.Println("enter any int to continue, 0 to end")
            t := time.Now()
            elapsed := t.Sub(start)
            // only if 
                if int(elapsed.Seconds()) != 0 {
                        fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                            check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                        Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                            check(err0)
                    current_time := time.Now()
                        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                        check(err6)
                    // the partial pi is printed to the datalog file on the line below
                        _ , err8 := fmt.Fprintf(fileHandle, "a peak at 97 digits of pi was %1.97f \n", pi) 
                            check(err8)
                    // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                            check(err7)
                }
            fmt.Scanf("%d", &test) 
            fmt.Scanf("%d", &test) 
            if test == 0 {
                t := time.Now()
                elapsed := t.Sub(start)
                // only if 
                    if int(elapsed.Seconds()) != 0 {
                            fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                                check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                                defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                            Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                                check(err0)
                        current_time := time.Now()
                            _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                            check(err6)
                        // the partial pi is printed to the datalog file on the line below
                            _ , err8 := fmt.Fprintf(fileHandle, "a peak at 97 digits of pi was %1.97f \n", pi) 
                                check(err8)
                        // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                            TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                            _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                                check(err7)
                    }
                fmt.Printf("%.49999f", pi)
                fmt.Println("\nprec is: ", precR)
                os.Exit(1)
            }
        }
        if i == 800 {
            fmt.Println("\n we are at 800 loops \n")
            fmt.Printf("%.69999f", pi)
            fmt.Println("\n we are at 800 loops \n")
            fmt.Println("enter any int to continue, 0 to end")
            t := time.Now()
            elapsed := t.Sub(start)
            // only if 
                if int(elapsed.Seconds()) != 0 {
                        fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                            check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                        Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                            check(err0)
                    current_time := time.Now()
                        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                        check(err6)
                    // the partial pi is printed to the datalog file on the line below
                        _ , err8 := fmt.Fprintf(fileHandle, "a peak at 111 digits of pi was %1.111f \n", pi) 
                            check(err8)
                    // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                            check(err7)
                }
            fmt.Scanf("%d", &test) 
            fmt.Scanf("%d", &test) 
            if test == 0 {
                t := time.Now()
                elapsed := t.Sub(start)
                // only if 
                    if int(elapsed.Seconds()) != 0 {
                            fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                                check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                                defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                            Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                                check(err0)
                        current_time := time.Now()
                            _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                            check(err6)
                        // the partial pi is printed to the datalog file on the line below
                            _ , err8 := fmt.Fprintf(fileHandle, "a peak at 111 digits of pi was %1.111f \n", pi) 
                                check(err8)
                        // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                            TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                            _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                                check(err7)
                    }
                fmt.Printf("%.69999f", pi)
                fmt.Println("\nprec is: ", precR)
                os.Exit(1)
            }
        }
        if i == 1600 {
            fmt.Println("\n we are at 1600 loops \n")
            fmt.Printf("%.89999f", pi)
            fmt.Println("\n we are at 1600 loops \n")
            fmt.Println("enter any int to continue, 0 to end")
            t := time.Now()
            elapsed := t.Sub(start)
            // only if 
                if int(elapsed.Seconds()) != 0 {
                        fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                            check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                        Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                            check(err0)
                    current_time := time.Now()
                        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                        check(err6)
                    // the partial pi is printed to the datalog file on the line below
                        _ , err8 := fmt.Fprintf(fileHandle, "a peak at 122 digits of pi was %1.122f \n", pi) 
                            check(err8)
                    // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                            check(err7)
                }
            fmt.Scanf("%d", &test) 
            fmt.Scanf("%d", &test) 
            if test == 0 {
                t := time.Now()
                elapsed := t.Sub(start)
                // only if 
                    if int(elapsed.Seconds()) != 0 {
                            fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                                check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                                defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                            Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                                check(err0)
                        current_time := time.Now()
                            _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                            check(err6)
                        // the partial pi is printed to the datalog file on the line below
                            _ , err8 := fmt.Fprintf(fileHandle, "a peak at 122 digits of pi was %1.122f \n", pi) 
                                check(err8)
                        // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                            TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                            _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                                check(err7)
                    }
                fmt.Printf("%.89999f", pi)
                fmt.Println("\nprec is: ", precR)
                os.Exit(1)
            }
        }
        if i == 2000 {
            fmt.Println("\n we are at 2000 loops \n")
            fmt.Printf("%.99999f", pi)
            fmt.Println("\n we are at 2000 loops \n")
            fmt.Println("enter any int to continue, 0 to end")
            t := time.Now()
            elapsed := t.Sub(start)
            // only if 
                if int(elapsed.Seconds()) != 0 {
                        fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                            check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                        Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                            check(err0)
                    current_time := time.Now()
                        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                        check(err6)
                    // the partial pi is printed to the datalog file on the line below
                        _ , err8 := fmt.Fprintf(fileHandle, "a peak at 132 digits of pi was %1.132f \n", pi) 
                            check(err8)
                    // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                            check(err7)
                }
            fmt.Scanf("%d", &test) 
            fmt.Scanf("%d", &test) 
            if test == 0 {
                t := time.Now()
                elapsed := t.Sub(start)
                // only if 
                    if int(elapsed.Seconds()) != 0 {
                            fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                                check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                                defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                            Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                                check(err0)
                        current_time := time.Now()
                            _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                            check(err6)
                        // the partial pi is printed to the datalog file on the line below
                            _ , err8 := fmt.Fprintf(fileHandle, "a peak at 132 digits of pi was %1.132f \n", pi) 
                                check(err8)
                        // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                            TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                            _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                                check(err7)
                    }
                fmt.Printf("%.99999f", pi)
                fmt.Println("\nprec is: ", precR)
                os.Exit(1)
            }
        }
        if i == 2400 {
            fmt.Println("\n we are at 2400 loops \n")
            fmt.Printf("%.111999f", pi)
            fmt.Println("\n we are at 2400 loops \n")
            fmt.Println("enter any int to continue, 0 to end")
            t := time.Now()
            elapsed := t.Sub(start)
            // only if 
                if int(elapsed.Seconds()) != 0 {
                        fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                            check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                        Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                            check(err0)
                    current_time := time.Now()
                        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                        check(err6)
                    // the partial pi is printed to the datalog file on the line below
                        _ , err8 := fmt.Fprintf(fileHandle, "a peak at 142 digits of pi was %1.142f \n", pi) 
                            check(err8)
                    // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                            check(err7)
                }
            fmt.Scanf("%d", &test) 
            fmt.Scanf("%d", &test) 
            if test == 0 {
                t := time.Now()
                elapsed := t.Sub(start)
                // only if 
                    if int(elapsed.Seconds()) != 0 {
                            fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                                check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                                defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                            Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                                check(err0)
                        current_time := time.Now()
                            _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                            check(err6)
                        // the partial pi is printed to the datalog file on the line below
                            _ , err8 := fmt.Fprintf(fileHandle, "a peak at 142 digits of pi was %1.142f \n", pi) 
                                check(err8)
                        // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                            TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                            _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                                check(err7)
                    }
                fmt.Printf("%.113999f", pi)
                fmt.Println("\nprec is: ", precR)
                os.Exit(1)
            }
        }
        if i == 2800 {
            fmt.Println("\n we are at 2800 loops \n")
            fmt.Printf("%.119999f", pi)
            fmt.Println("\n we are at 2800 loops \n")
            fmt.Println("enter any int to continue, 0 to end")
            t := time.Now()
            elapsed := t.Sub(start)
            // only if 
                if int(elapsed.Seconds()) != 0 {
                        fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                            check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                        Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                            check(err0)
                    current_time := time.Now()
                        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                        check(err6)
                    // the partial pi is printed to the datalog file on the line below
                        _ , err8 := fmt.Fprintf(fileHandle, "a peak at 152 digits of pi was %1.152f \n", pi) 
                            check(err8)
                    // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                            check(err7)
                }
            fmt.Scanf("%d", &test) 
            fmt.Scanf("%d", &test) 
            if test == 0 {
                t := time.Now()
                elapsed := t.Sub(start)
                // only if 
                    if int(elapsed.Seconds()) != 0 {
                            fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                                check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                                defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                            Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                                check(err0)
                        current_time := time.Now()
                            _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                            check(err6)
                        // the partial pi is printed to the datalog file on the line below
                            _ , err8 := fmt.Fprintf(fileHandle, "a peak at 152 digits of pi was %1.152f \n", pi) 
                                check(err8)
                        // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                            TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                            _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                                check(err7)
                    }
                fmt.Printf("%.129999f", pi)
                fmt.Println("\nprec is: ", precR)
                os.Exit(1)
            }
        }
        if i == 3200 {
            fmt.Println("\n we are at 3200 loops \n")
            fmt.Printf("%.129999f", pi)
            fmt.Println("\n we are at 3200 loops \n")
            fmt.Println("enter any int to continue, 0 to end")
            t := time.Now()
            elapsed := t.Sub(start)
            // only if 
                if int(elapsed.Seconds()) != 0 {
                        fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                            check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                        Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                            check(err0)
                    current_time := time.Now()
                        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                        check(err6)
                    // the partial pi is printed to the datalog file on the line below
                        _ , err8 := fmt.Fprintf(fileHandle, "a peak at 162 digits of pi was %1.162f \n", pi) 
                            check(err8)
                    // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                            check(err7)
                }
            fmt.Scanf("%d", &test) 
            fmt.Scanf("%d", &test) 
            if test == 0 {
                t := time.Now()
                elapsed := t.Sub(start)
                // only if 
                    if int(elapsed.Seconds()) != 0 {
                            fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                                check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                                defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                            Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                                check(err0)
                        current_time := time.Now()
                            _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                            check(err6)
                        // the partial pi is printed to the datalog file on the line below
                            _ , err8 := fmt.Fprintf(fileHandle, "a peak at 162 digits of pi was %1.162f \n", pi) 
                                check(err8)
                        // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                            TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                            _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                                check(err7)
                    }
                fmt.Printf("%.139999f", pi)
                fmt.Println("\nprec is: ", precR)
                os.Exit(1)
            }
        }
        if i == 4000 {
            fmt.Println("\n we are at 4000 loops \n")
            fmt.Printf("%.139999f", pi)
            fmt.Println("\n we are at 4000 loops \n")
            fmt.Println("enter any int to continue, 0 to end")
            t := time.Now()
            elapsed := t.Sub(start)
            // only if 
                if int(elapsed.Seconds()) != 0 {
                        fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                            check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                        Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                            check(err0)
                    current_time := time.Now()
                        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                        check(err6)
                    // the partial pi is printed to the datalog file on the line below
                        _ , err8 := fmt.Fprintf(fileHandle, "a peak at 177 digits of pi was %1.177f \n", pi) 
                            check(err8)
                    // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                            check(err7)
                }
            fmt.Scanf("%d", &test) 
            fmt.Scanf("%d", &test) 
            if test == 0 {
                t := time.Now()
                elapsed := t.Sub(start)
                // only if 
                    if int(elapsed.Seconds()) != 0 {
                            fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                                check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                                defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                            Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                                check(err0)
                        current_time := time.Now()
                            _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                            check(err6)
                        // the partial pi is printed to the datalog file on the line below
                            _ , err8 := fmt.Fprintf(fileHandle, "a peak at 177 digits of pi was %1.177f \n", pi) 
                                check(err8)
                        // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                            TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                            _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                                check(err7)
                    }
                fmt.Printf("%.149999f", pi)
                fmt.Println("\nprec is: ", precR)
                os.Exit(1)
            }
        }
        if i == 6000 {
            fmt.Println("\n we are at 6000 loops \n")
            fmt.Printf("%.149999f", pi)
            fmt.Println("\n we are at 4000 loops \n")
            fmt.Println("enter any int to continue, 0 to end")
            t := time.Now()
            elapsed := t.Sub(start)
            // only if 
                if int(elapsed.Seconds()) != 0 {
                        fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                            check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                        Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                            check(err0)
                    current_time := time.Now()
                        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                        check(err6)
                    // the partial pi is printed to the datalog file on the line below
                        _ , err8 := fmt.Fprintf(fileHandle, "a peak at 187 digits of pi was %1.187f \n", pi) 
                            check(err8)
                    // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                            check(err7)
                }
            fmt.Scanf("%d", &test) 
            fmt.Scanf("%d", &test) 
            if test == 0 {
                t := time.Now()
                elapsed := t.Sub(start)
                // only if 
                    if int(elapsed.Seconds()) != 0 {
                            fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                                check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                                defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                            Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                                check(err0)
                        current_time := time.Now()
                            _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                            check(err6)
                        // the partial pi is printed to the datalog file on the line below
                            _ , err8 := fmt.Fprintf(fileHandle, "a peak at 187 digits of pi was %1.187f \n", pi) 
                                check(err8)
                        // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                            TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                            _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                                check(err7)
                    }
                fmt.Printf("%.159999f", pi)
                fmt.Println("\nprec is: ", precR)
                os.Exit(1)
            }
        }
        if i == 8000 {
            fmt.Println("\n we are at 8000 loops \n")
            fmt.Printf("%.159999f", pi)
            fmt.Println("\n we are at 4000 loops \n")
            fmt.Println("enter any int to continue, 0 to end")
            t := time.Now()
            elapsed := t.Sub(start)
            // only if 
                if int(elapsed.Seconds()) != 0 {
                        fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                            check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                            defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                        Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                            check(err0)
                    current_time := time.Now()
                        _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                        check(err6)
                    // the partial pi is printed to the datalog file on the line below
                        _ , err8 := fmt.Fprintf(fileHandle, "a peak at 197 digits of pi was %1.197f \n", pi) 
                            check(err8)
                    // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                        TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                        _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                            check(err7)
                }
            fmt.Scanf("%d", &test) 
            fmt.Scanf("%d", &test) 
            if test == 0 {
                t := time.Now()
                elapsed := t.Sub(start)
                // only if 
                    if int(elapsed.Seconds()) != 0 {
                            fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                                check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                                defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                            Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  --  pi-chudnovsky  --  selection #%d on %s \n", selection, Hostname)
                                check(err0)
                        current_time := time.Now()
                            _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                            check(err6)
                        // the partial pi is printed to the datalog file on the line below
                            _ , err8 := fmt.Fprintf(fileHandle, "a peak at 197 digits of pi was %1.197f \n", pi) 
                                check(err8)
                        // ... after printing the partial pi, some nice stats are appended to the file's log entry 
                            TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                            _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %1.0f \n ", TotalRun, digits) 
                                check(err7)
                    }
                fmt.Printf("%.169999f", pi)
                fmt.Println("\nprec is: ", precR)
                os.Exit(1)
            }
        }
        // 1,000,000 digits requires 70516 loops, per the run on May 7 2023 at 10:30
        //  was run on: Sun May  7 08:50:23 2023
        //  Total run was 8h4m39.7847064s
        // AND THE CALCULATION WAS INDEPENDANTLY VERIFIED !!!!!!!!!!!
        } // *** ending bracket of the else which is used to skip all of the above steps 
    }
    return pi, prec, i
}

/*
    var pages_of_code float32
    var sloc1 float32
    var sloc2 float32
func sloc() {
    sloc1 = 3706                   // total sloc 
    sloc2 = sloc1 * float32(0.42) // effective sloc
    pages_of_code = sloc1 / float32(49)
}
*/

func check(e error) {   // create a func named check which takes one parameter "e" of type error 
    if e != nil {
        panic(e)        // use panic() to display error code 
    }
}


// gloabl vars for use in multiple localized sections of code 


        var which_menu int 

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
        var selection int 
        var precisionOfSquare int 
        var t2 time.Time 

    // the following globals, are used in multiple funcs of case 18: calculate either square or cube root of any integer 
    //------------------------------------------------------------------------------------------------------------------
        var Tim_win float64 // Time Window
        var sortedResults = []Results{} // sortedResults is an array of type Results as defined at the top of this file 
        var Table_of_perfect_Products = []int{}  
        var diffOfLarger int 
        var diffOfSmaller int 
        var perfectResult2 float64 // will contain the square root result if the workpiece was itself a perfect square 
        var perfectResult3 float64 // will contain the cube root result if the workpiece was itself a perfect cube 
        var precisionOfRoot int // this being global means we do not need to pass it in to the read func 
        var workPiece int // the square or cube of which we are to find a root 
        var skip_redoing_loop int 
            // the following three arrays were my original method of logging my results and the nearness of fit 
                //var List_of_2_results_case18 = []float64{}
                //var corresponding_diffs = []int{}
                //var diffs_as_percent = []float64{}

type Results struct {  // define a new structure called Results with two fields; result, and pdiff 
    result float64
    pdiff float64
}