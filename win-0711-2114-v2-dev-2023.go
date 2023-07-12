// The initial inspiration for all of this was : Veritassium https://youtu.be/gMlf1ELvRzc?list=LL
/*
    Go.lang is a relatively-new general-purpose freely-avialable programming language devoped by google. 

    One can obtain the Go Compiler from : https://go.dev/dl/   
*/

// Unix/Linux/Mac variant build instructions : ==========================================================================
                        /*
    compile with: "go build -o desired-name-of-your-executable name-of-this-source-code-file.go"
    ... thereafter you can run it on a Unix system with "/fullpathname/desired-name-of-your-executable"
        ( e.g.  "./Richards_go_project" )
    ... or, having first obtained the Go compiler, ... just run the source with: "go run name-of-this-source-code-file.go"
                        */
//=======================================================================================================================

// Windows variant : ----------------------------------------------------------------------------------------------------
                        /*
    build with: "go build -o desired-name-of-your-executable.exe name-of-this-source-code-file.go"
    ... or, having first obtained the Go compiler, ... just run the source with: "go run name-of-this-source-code-file.go"
                        */
//-----------------------------------------------------------------------------------------------------------------------

// -- AMFmainA-a
package main

import (
    "os"          // Used to fetch the name of your system 
    "sync"        // Used in Bailey–Borwein–Plouffe formula [concurent]
    "io/ioutil"   // Used for file access 
    "fmt"         // Used for printing etc. 
    "math"        // Used for math.Pow(base, exponent) and other more-complex math funcs 
    "runtime"     // Used to get information on available CPUs
    "time"        // Used primarily for calculating run times 
    "strconv"     // Used in Spigot 
    "os/signal"   // Used in *** nifty scoreboard *** concurrent computation of pi using Nilakantha's formula, ...
                  // ... by Diego Brener diegosilva13 on Github (with mods by yours-truly)
    "math/big"    // Used in ArchimedesBig(), and in bpp formula with mods by rick.woolley@gmail.com (pi calculator ...
                  // ...from https://github.com/karan/Projects)
    "sort"        // Used in xRootOfy (case 18 and 180) to sort an array of structs 
    "regexp"      // used in multiple instance to create regular expression patterns 

    "bufio"       // used to count the lines in this file for reportSLOCstats() and also to Scan this file in order to display lines of source code 
    "log"         // used just once as: log.Fatal(err)
    "strings"     // used in reportSLOCstats()
)

func main() {       // top-level program logic flow -- explore several ways to calculate Pi, plus THREE other constants --AMFmain-A-b

    filenameOfThisFile := getTheNameOfThisSourceCodeFile()  // a locally-defined func 

    totalLines, nonEmptyLines, err := reportSLOCstats(filenameOfThisFile)  // another locally-defined func; returns, and creates, three local values of predetermined type 
        if err != nil {                                  
            log.Fatal(err)                                
        } 

            // The following is for menu header data generation (inception: stripped file name) //  <<------------ below ---------<<
            //
             /* Unix variant 
                re2 := regexp.MustCompile(`Unix-(.+)\.go`)  
             Unix variant */ 
            //
             // /* Windows variant 
                re2 := regexp.MustCompile(`win-(.+)\.go`)
             // Windows variant */ 

    match2 := re2.FindStringSubmatch(filenameOfThisFile) // grab the stuff between win- and .go (or Unix- and .go) in the name of this file 
        SansVerOfNameOfThisFile := "string"      // this var, having been initialize to "string", is of type string 
        if len(match2) >= 2 {

            /* Unix variant 
                SansVerOfNameOfThisFile = match2[0]                   // Unix SansVerOfNameOfThisFile is loaded with our base file name 
            Unix variant */ 
            //
            // /* Windows variant 
                SansVerOfNameOfThisFile = match2[1]                   // Windows SansVerOfNameOfThisFile is loaded with our base file name 
            // Windows variant */

        } else {
            fmt.Println("Substring2 not found in main")
        }

    for 1 == 1 {
        DisplayMenus(totalLines, nonEmptyLines, filenameOfThisFile, SansVerOfNameOfThisFile) 
        fmt.Println(string(colorRed), "Finished. Hit Enter to redisplay the Main Menu", string(colorReset)) // this will be the last line of every case #: 
            var JustToPauseTheLoop int 

            /* Unix variant 
                fmt.Scanf("%d", &JustToPauseTheLoop) // must pause the endless loop after every switch case to prevent menu redisplay over results
                    //fmt.Scan(&JustToPauseTheLoop) // Scan() would work as a pause, but it requies data input to continue, so we use a Scanf() instead 
            Unix variant */ 
            //
            // /* Windows variant 
                fmt.Scanf("%d", &JustToPauseTheLoop) // must pause the endless loop after every switch case to prevent menu redisplay over results 
                fmt.Scanf("%d", &JustToPauseTheLoop) 
                    //fmt.Scan(&JustToPauseTheLoop) // Scan does not work quite as well as the Scanf combo above (but a single Scanf worked well in Unix environment)
            // Windows variant */ 

    }
} // end of main --AMFmain-A-c


func getTheNameOfThisSourceCodeFile() (filenameOfThisFile string) {  // filenameOfThisFile var is created local to this func by its having been listed as a returned val
    // get the name of this source code file to be used in the display-of-source-code funcs 
    pathNameOfMeAsExe := os.Args[0]      // the objective is to find this source code file by name from the command line Args
                
            /* Unix variant 
                re := regexp.MustCompile(`Unix(.*?)2023`) 
            Unix variant */ 
            //
            // /* Windows variant 
                re := regexp.MustCompile(`\\exe\\(.*?)\.exe`) 
            // Windows variant */ 

    match := re.FindStringSubmatch(pathNameOfMeAsExe) // grab the stuff between /exe/ and .exe in the full pathname of the generated executable 
    var substring string                           // ... both Windows and Unix  ^^^
        if len(match) >= 2 {

            /* Unix variant 
                substring = match[0]                   // Unix substring is loaded with our base file name 
            Unix variant */ 
            //
            // /* Windows variant 
                substring = match[1]                   // Windows substring is loaded with our base file name 
            // Windows variant */ 

        } else {
            fmt.Println("Substring not found in getTheNameOfThisSourceCodeFile()")
        }
        filenameOfThisFile = substring + ".go"         // append the .go on the end of the semi-stripped file name 
        return // a naked return of the defined/specified value 
}


// --AMFreportSLOCstats01a
func reportSLOCstats(filenameOfThisFile string) (totalLines, nonEmptyLines int, err error) {   // returns two values of type int, and one of type error 
    file, err := os.Open(filenameOfThisFile)           // open this source code file and get a handle to it 
    if err != nil {
        return 0, 0, err
    }
    defer file.Close()
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            totalLines++            // this var was created locally by having declared it as a return value in this funcs def
            line := strings.TrimSpace(scanner.Text())
            if line != "" {
                nonEmptyLines++     // this var was created locally by having declared it as a return value in this funcs def
            }
        }
    if scanner.Err() != nil {
        return 0, 0, scanner.Err()  // this func must return three values, in the case of an error it returns 0 for the line counts 
    }
    return totalLines, nonEmptyLines, nil         
} // --AMFreportSLOCstats01b -- AMFmainB



// This next func is extremely long, contains all menus, and also a very long switch statement -- AMFmenusA
func DisplayMenus(totalLines, nonEmptyLines int, filenameOfThisFile, SansVerOfNameOfThisFile string) {

    pages_of_code := totalLines / 49 
    var selection int 

    date := time.Now() // only used here once as date.Format("(01-02-2006)") at the top of the main menu **** NOT FOR RELEASE **********
    //*********** the above line will need to be commented out for release, and the date below entered manually ************************

fmt.Print(string(colorYellow), 
  "\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n MAIN MENU  ",

  string(colorGreen), " inception: ", string(colorReset), SansVerOfNameOfThisFile, string(colorGreen),
   "    build date: ", string(colorCyan), date.Format("(01-02-2006)"), string(colorReset), "\n\n")

fmt.Println(string(colorRed), "Enter your selection from below, or 96 for notes on using this app\n", string(colorReset))
fmt.Print("2: ", string(colorCyan), " Pi: ", string(colorReset), "Bailey–Borwein–Plouffe formula for π, discovered in 1995", string(colorYellow))
fmt.Print(" (best method)\n", string(colorReset))
fmt.Println("        ---> uses big.Float types, and also uses channels: GOMAXPROCS(numCPU)\n")

fmt.Print("14:", string(colorCyan), " Pi: ", string(colorReset), "modified Archimedes' method using Go's math/big objects: 3,012 digits\n")
fmt.Println("        ---> uses big.Float (Richard's favorite: 'cause it's fully understood)")
fmt.Print("         (employs the Pythagorean, ", string(colorGreen), "ENTER '11'", string(colorReset), " for a review of its derivation)\n\n")

fmt.Print("15:", string(colorCyan), " Pi: ", string(colorReset), "Chudnovsky method, state of the art! But God only knows how it works\n") 
fmt.Println("        ---> uses big.Float types \n")

fmt.Print("5: ", string(colorCyan), " Pi: ", string(colorReset), "An infinite series by Indian astronomer Nilakantha Somayaji circa 1500\n")
fmt.Println("        ---> uses big.Floats types         AKA: Keļallur Comatiri: 1444–1545")
fmt.Println("     π = 3 + 4/(2*3*4) - 4/(4*5*6) + 4/(6*7*8) - 4/(8*9*10) ...")
fmt.Println("         ... the equation can be found in an ancient Sanskrit verse\n")


fmt.Println("6:", string(colorCyan), "Pi:", string(colorReset), "Gottfried Wilhelm Leibniz, late 17th century ( and by Isaac Newton )")
fmt.Println("     π = 4 * ( 1 - 1/3 + 1/5 - 1/7 + 1/9 - 1/11 + 1/13 ...")
fmt.Println("       Four-Billion iterations to be executed ------ using float64 types <---")
fmt.Println("       10 digits of π : in about ten seconds \n")

fmt.Println("7:", string(colorCyan), "Pi:", string(colorReset), "The Gregory-Leibniz series circa 1676")
fmt.Println("     π = (4/1) - (4/3) + (4/5) - (4/7) + (4/9) - (4/11) ...")
fmt.Println("       Three-Hundred-Million iterations are executed ----- float64 types <---")
fmt.Println("       9 digits of π : under 4 seconds\n")

fmt.Println(string(colorYellow), "12: Display prior execution times from longer-running prior selections \n")

fmt.Println(" 13: FOR SECOND MENU\n", string(colorReset))

fmt.Println("47: to End/Exit", string(colorCyan), "tSLOC =", totalLines, "eSLOC =", nonEmptyLines, string(colorPurple), 
  "\u00a9 2023, by Richard Hart Woolley", string(colorReset))
fmt.Println(string(colorCyan),"                That is", int(pages_of_code), "pages of code at 49 lines per page\n", string(colorReset))
fmt.Print("Enter your selection, 1 -> x", string(colorRed), " (IS THIS WINDOW MAXIMIZED?  Well, do it!)\n", string(colorReset)) 

    fmt.Scan(&selection)  // Pause and obtain Selection 


if selection == 13 { //==========================================================================================================================

fmt.Print(string(colorYellow))
fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")

// --AMFmenu01a 
fmt.Println(string(colorGreen), "    build date :", string(colorCyan), date.Format("(01-02-2006)"), string(colorReset), "  inception :", SansVerOfNameOfThisFile)
// --AMFmenu01b 

fmt.Print(string(colorYellow))
fmt.Print("\nSECOND MENU:")
fmt.Print(string(colorRed),
" Enter a selection from below, or 96 for notes on using the app", string(colorCyan), "\n") 

fmt.Println("\n18:  Calculate", string(colorGreen), "the second or third Root of y (x\u221Ay)", string(colorReset), "from first-principles")
fmt.Println("      ... via the ratio of y:1 of perfect products \n", string(colorCyan))
fmt.Println("9:   Calculate", string(colorGreen), "Euler's Number: \u2107 or \u2147", string(colorReset), "the natural logarithmic base")
fmt.Println("        Explore the limit of interest\n", string(colorCyan))
fmt.Println("10:  Calculate the", string(colorGreen), "Erdős-Borwein constant", string(colorReset), "from a breif infinite series\n")
fmt.Println("8:", string(colorGreen), " Pi:", string(colorReset), "An infinite series by John Wallis circa 1655")
fmt.Println("      π = 2 * ((2/1)*(2/3)) * ((4/3)*(4/5)) * ((6/5)*(6/7)) ... ")
fmt.Println("        One-Billion iterations will be executed; option for 40 billion")
fmt.Println("        9 digits of π -- a billion loops, in seconds -- option for 10 digits\n")
fmt.Println("19:", string(colorGreen), "Pi:", string(colorReset), "Open the 'Spigot' algorithm, instantly bakes way too much pie :)")
fmt.Println("      ... Can easily spit out 10,000 digits of π !!!!!\n", string(colorCyan))

fmt.Println("37:", string(colorGreen), "Pi:", string(colorReset), "Gauss–Legendre algorithm ")

fmt.Println("40:", string(colorGreen), "Pi:", string(colorReset), "Nifty 'ScoreBoard' using Nilakantha's formula", string(colorYellow), 
  "(Ctrl-C to exit it)", string(colorCyan))
fmt.Println("41:", string(colorGreen), "Pi:", string(colorReset), "Bailey–Borwein–Plouffe formula", string(colorCyan), "[concurent]", 
  string(colorReset))
fmt.Println("42:", string(colorGreen), "Pi:", string(colorReset), "BBP formula to 46 digits")
fmt.Println("44:", string(colorGreen), "Pi:", string(colorReset), "via Leibniz method in one billion iterations [runs a while]")
fmt.Println("50:", string(colorGreen), "Pi:", string(colorReset), "comparison of float64 and big.Float types \n\n")

fmt.Println(string(colorYellow), "12: Display prior execution times from longer-running prior selections ")

fmt.Println(string(colorYellow), "\n  0:  To return to main menu", string(colorRed), "\n")

        // ("47:  to End/Exit vvvvvvvvvvvvvvvv
fmt.Println("47: to End/Exit", string(colorCyan), "tSLOC =", totalLines, "eSLOC =", nonEmptyLines, string(colorPurple), 
  "\u00a9 2023, by Richard Hart Woolley\n", string(colorReset))
fmt.Print("Enter your selection, 1 -> x", string(colorRed), " (IS THIS WINDOW MAXIMIZED?  Well, do it!)\n", string(colorReset)) 

    fmt.Scan(&selection)  // pause and request input from the user

} //===============================================================================================================================================


if selection == 12 { //==============================================================================================================================

        //case 12: // display contents of prior results file
fmt.Print(string(colorYellow))
fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nStats MENU:", 

  string(colorRed), "Enter your selection from below", string(colorReset), "\n") 


fmt.Println("212:   Read log file containing stats from prior runs of ALL long-running cases \n")

fmt.Println("202:   Review the output from prior runs of BBPF (case 2:) \n") 

fmt.Println("214:   Review the output from prior runs of modified Archimedes' method (case 14:) \n")

fmt.Println("215:   Read stats from prior runs of the Chudnovsky method (case 15:) which did > 800 loops \n")
fmt.Println("315:   Look at the really big pie that came out of the Chudnovsky method (case 15:) \n")

fmt.Println("205:   Review the data from prior runs of Nilakantha Somayaji (case 5:) \n")

fmt.Println("241:   Review stats and data from prior runs of Bailey–Borwein–Plouffe [concurent] (case 41:) \n")

fmt.Println("180:  ", string(colorRed), "Run", string(colorReset), "the endless test of case 18 \n")
fmt.Println("280:   Read log file from the aformentioned testing case \n\n")

fmt.Println("999: Delete all log files \n\n")

fmt.Println(string(colorYellow), "0:  To return to main menu", string(colorRed), "\n")

        // ("47:  to End/Exit vvvvvvvvvvvvvvvv
fmt.Println("47: to End/Exit", string(colorCyan), "tSLOC =", totalLines, "eSLOC =", nonEmptyLines, string(colorPurple), 
  "\u00a9 2023, by Richard Hart Woolley\n", string(colorReset))
fmt.Print("Enter your selection, 1 -> x", string(colorRed), " (IS THIS WINDOW MAXIMIZED?  Well, do it!)\n", string(colorReset)) 

    fmt.Scan(&selection)  // pause and request input from the user

} //===============================================================================================================================================


    switch selection { // if selection == 2 then handel per case 2: etc. 

        case 999:
            deleteAllLogFiles()

        case 3:
            ShowMain(filenameOfThisFile)

        case 2:  // (best method) the Bailey–Borwein–Plouffe formula for π, circa 1995 
            BBPF(selection)
            case 22:
            ShowBBPF(filenameOfThisFile)

        case 4:
            ShowMenus(filenameOfThisFile)

        case 5: // a series by Indian astronomer Nilakantha Somayaji circa 1500 AD 
            Nilakantha_Somayaji_with_big_Float_types(selection)
            case 25:
            ShowNilakantha_Somayaji_with_big_Float_types(filenameOfThisFile)

        case 6: // Gottfried Wilhelm Leibniz
            GottfriedWilhelmLeibniz(selection)
            case 26:
            ShowGottfriedWilhelmLeibniz(filenameOfThisFile)

        case 7: // the Gregory-Leibniz series
            GregoryLeibniz(selection)
            case 27: 
            ShowGregoryLeibniz(filenameOfThisFile)

        case 8: // John Wallis circa 1655 --- runs very long 
            JohnWallis(selection)
            case 28: 
            ShowJohnWallis(filenameOfThisFile)

        case 9: // Euler's Number
            EulersNumber(selection)
            case 29: 
            ShowEulersNumber(filenameOfThisFile)

        case 10: // Erdős–Borwein constant
            ErdosBorwein(selection)
            case 30:
            ShowErdosBorwein(filenameOfThisFile)

        case 11:  // display a review of the derivation of the Pythagorean
            DisplayPythagorean(selection)
            case 31: 
            DisplayPythagoreanCode()


        // cases related to the menu for log file viewing 
        case 202:
                content, err := ioutil.ReadFile("dataLog-From_BBPF_Method_lengthy_prints.txt")  // 
                if err != nil {   // if the file does not exist ... 
                    fmt.Println(string(colorCyan), "\nNo prior results -- no log file", string(colorWhite), "'dataLog-From_BBPF_Method_lengthy_prints.txt'", string(colorCyan),
                     "exists\n")
                } else {
                    fmt.Println(string(content))  // dump/display entire file to command line
                }
        case 214:
                content, err := ioutil.ReadFile("dataLog-From_AM_Method_lengthy_prints.txt")  // 
                if err != nil {   // if the file does not exist ... 
                    fmt.Println(string(colorCyan), "\nNo prior results -- no log file", string(colorWhite), "'dataLog-From_AM_Method_lengthy_prints.txt'", string(colorCyan),
                     "exists\n")
                } else {
                    fmt.Println(string(content))  // dump/display entire file to command line
                }
        case 215: 
                content, err := ioutil.ReadFile("dataLog-From_Chudnovsky_Method_lengthy_prints.txt")  // 
                if err != nil {   // if the file does not exist ... 
                    fmt.Println(string(colorCyan), "\nNo prior results -- no log file", string(colorWhite), "'dataLog-From_Chudnovsky_Method_lengthy_prints.txt'", string(colorCyan),
                     "exists\n")
                } else {
                    fmt.Println(string(content))  // dump/display entire file to command line
                }
        case 315:
                content, err := ioutil.ReadFile("big_pie_is_in_here.txt")  // 
                if err != nil {   // if the file does not exist ... 
                    fmt.Println(string(colorCyan), "\nNo prior results -- no log file", string(colorWhite), "'big_pie_is_in_here.txt'", string(colorCyan),
                     "exists\n")
                } else {
                    fmt.Println(string(content))  // dump/display entire file to command line
                }
        case 205:
                content, err := ioutil.ReadFile("dataLog-From_Nilakantha_Method_lengthy_prints.txt")  // 
                if err != nil {   // if the file does not exist ... 
                    fmt.Println(string(colorCyan), "\nNo prior results -- no log file", string(colorWhite), "'dataLog-From_Nilakantha_Method_lengthy_prints.txt'", string(colorCyan),
                     "exists\n")
                } else {
                    fmt.Println(string(content))  // dump/display entire file to command line
                }
        case 212:
                content, err := ioutil.ReadFile("dataLog-From_calculate-pi-and-friends.txt")  // 
                if err != nil {   // if the file does not exist ... 
                    fmt.Println(string(colorCyan), "\nNo prior results -- no log file", string(colorWhite), "'dataLog-From_calculate-pi-and-friends.txt'", string(colorCyan),
                     "exists\n")
                } else {
                    fmt.Println(string(content))  // dump/display entire file to command line
                }
        case 241:
                content, err := ioutil.ReadFile("dataLog-From_BBPfConcurent.txt")  // 
                if err != nil {   // if the file does not exist ... 
                    fmt.Println(string(colorCyan), "\nNo prior results -- no log file", string(colorWhite), "'dataLog-From_BBPfConcurent.txt'", string(colorCyan),
                     "exists\n")
                } else {
                    fmt.Println(string(content))  // dump/display entire file to command line
                }
        case 280:
                content, err := ioutil.ReadFile("logfile_from_selection_180.txt")  // 
                if err != nil {   // if the file does not exist ... 
                    fmt.Println(string(colorCyan), "\nNo prior results -- no log file", string(colorWhite), "'logfile_from_selection_180.txt'", string(colorCyan),
                     "exists\n")
                } else {
                    fmt.Println(string(content))  // dump/display entire file to command line
                }


        case 180:
            xRootOfY_continuousCaller(selection)
            case 380:
            ShowxRootOfY_continuousCaller(filenameOfThisFile)

        case 14: 
            ArchimedesBig(selection)
            case 34:
            ShowArchimedesBig(filenameOfThisFile)

        case 15: 
            chud(selection)
            case 35:
            Showchud(filenameOfThisFile)
        
        case 18:
            xRootOfy(selection) 
            case 38:
            ShowxRootOfy(filenameOfThisFile)
        
        case 19: 
            TheSpigot(selection)
            case 39:
            ShowTheSpigot(filenameOfThisFile)

        case 37:
            Gauss_Legendre(selection)
            case 57:
            Show_Gauss_Legendre(filenameOfThisFile)

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
            case 60:
            Shownifty_scoreBoard(filenameOfThisFile)

        case 41:
            BBPfConcurent(selection)
            case 61:
            ShowBBPfConcurent(filenameOfThisFile)

        case 42:
            bbp_formula(selection)
            case 62:
            Showbbp_formula(filenameOfThisFile)

        case 44:
            Leibniz_method_one_billion_iters(selection)
            case 64:
            ShowLeibniz_method_one_billion_iters(filenameOfThisFile)

        case 47:
            os.Exit(1)

        case 50:
            compareFloat64withBigFloats()
            case 70:
            ShowcompareFloat64withBigFloats(filenameOfThisFile)

        case 96:
            Using_this_app()

        case 0: 
            // Only as a non-functional case to allow returning directly to main menu 

    default:
        fmt.Println("\n ... You entered a value that is", string(colorRed), "not a valid option", string(colorReset), "go fish\n")
        // handles every case not explicitly listed (which has not been otherwise handled) 
    }
}  // end DisplayMenus() -- AMFmenusB


func DisplayPythagorean(selection int){  // case 11: 
    fmt.Print("\n\n\n -- You entered '", selection, "' to review the derivation of the Pythagorean, which was needed in method #4. We will\n")
fmt.Println("be geometrically deriving the Pythagorean theorem according to the 12th century Indian, Bhaskara.")
fmt.Println("    We begin with a square of area c\u00b2. We then partially fill that square with four congruent")
fmt.Println("right triangles each with its right angle opposite to one of the sides 'c'. Each of the four congruent")
fmt.Println("triangles now have sides c, a, and b. 'c' being the hypotenuse of each; 'a' being the shortest side and 'b'")
fmt.Println("being the longer of sides 'a' and 'b'. Thus leaving a small square in the center, which we can ignore.")

fmt.Println("     Next, we detach and slide two adjacent right triangles, each to its opposite side of the parent square.")
fmt.Println("We now have two attached squares a x a and b x b. We can then clearly see that", string(colorCyan), "c\u00b2 = a\u00b2 + b\u00b2", string(colorReset))
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
fmt.Println("the Area is  : ", string(colorCyan), "A = (a+b) * 1/2(a+b)  -- Or: --  A = 1/2(a+b)\u00b2 ")
fmt.Println(string(colorReset), "    But the area of this particular trapezoid that we have constructed is, obviously, also a*b")
fmt.Println("(the combined area of the initial two right triangles), plus 1/2*c\u00b2.   Or:  -- ", string(colorCyan))
fmt.Println("A = ab + 1/2*c\u00b2 \n", string(colorReset))
fmt.Println("           a new equality can then be constructed from the two above equalities: \n", string(colorCyan))
fmt.Println("    1/2(a+b)\u00b2  =  ab + 1/2(c\u00b2) \n", string(colorReset))
fmt.Println("                          and reducing it thusly : ", string(colorCyan))
fmt.Println("        (a+b)\u00b2 = 2ab + c\u00b2 ")
fmt.Println(" a\u00b2 + 2ab + b\u00b2 = 2ab + c\u00b2 ")
fmt.Println("       a\u00b2 + b\u00b2 = c\u00b2 \n", string(colorReset))
fmt.Println("... proves the Pythagorean per Garfield; though, obviously, many other proofs do exist.\n\n")
    fmt.Println("Select 12 at menu to display prior results")
}
func DisplayPythagoreanCode() {  // case 31: 
    var DisplayPythagoreanRune = `

This section was conceived of and written entirely by yours-truly. 

func DisplayPythagorean(selection int){  // case 11: 
    fmt.Print("\n\n\n -- You entered '", selection, "' to review the derivation of the Pythagorean, which was needed in method #4. We will\n")
fmt.Println("be geometrically deriving the Pythagorean theorem according to the 12th century Indian, Bhaskara.")
fmt.Println("    We begin with a square of area c\u00b2. We then partially fill that square with four congruent")
fmt.Println("right triangles each with its right angle opposite to one of the sides 'c'. Each of the four congruent")
fmt.Println("triangles now have sides c, a, and b. 'c' being the hypotenuse of each; 'a' being the shortest side and 'b'")
fmt.Println("being the longer of sides 'a' and 'b'. Thus leaving a small square in the center, which we can ignore.")

fmt.Println("     Next, we detach and slide two adjacent right triangles, each to its opposite side of the parent square.")
fmt.Println("We now have two attached squares a x a and b x b. We can then clearly see that", string(colorCyan), "c\u00b2 = a\u00b2 + b\u00b2", string(colorReset))
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
fmt.Println("the Area is  : ", string(colorCyan), "A = (a+b) * 1/2(a+b)  -- Or: --  A = 1/2(a+b)\u00b2 ")
fmt.Println(string(colorReset), "    But the area of this particular trapezoid that we have constructed is, obviously, also a*b")
fmt.Println("(the combined area of the initial two right triangles), plus 1/2*c\u00b2.   Or:  -- ", string(colorCyan))
fmt.Println("A = ab + 1/2*c\u00b2 \n", string(colorReset))
fmt.Println("           a new equality can then be constructed from the two above equalities: \n", string(colorCyan))
fmt.Println("    1/2(a+b)\u00b2  =  ab + 1/2(c\u00b2) \n", string(colorReset))
fmt.Println("                          and reducing it thusly : ", string(colorCyan))
fmt.Println("        (a+b)\u00b2 = 2ab + c\u00b2 ")
fmt.Println(" a\u00b2 + 2ab + b\u00b2 = 2ab + c\u00b2 ")
fmt.Println("       a\u00b2 + b\u00b2 = c\u00b2 \n", string(colorReset))
fmt.Println("... proves the Pythagorean per Garfield; though, obviously, many other proofs do exist.\n\n")
    fmt.Println("Select 12 at menu to display prior results")
} `
fmt.Println(DisplayPythagoreanRune)
}


// case 180:  /// this procedure is really only meant to be used for testing or illustration ... 
//                                              ... it calculates the root of all integers; and builds a file of all its results
// -- AMFxroyCCA
func xRootOfY_continuousCaller(selection int) { 
    usingBigFloats = false
    precisionOfRoot = 1700 // default value 
    workPiece = 2
    radical_index_number := 2
        fmt.Println("\nThis proc endlessly calculates the root of all integers beginging with 2; and builds a txt file of its results.\n")
        fmt.Println("\nEnter 2 for SquareRoot, or 3 for CubeRoot\n")
        fmt.Scan(&radical_index_number)  // pause and request input from the user 
    radical_index := radical_index_number
    for 1 == 1 {
        sortedResults = nil // reset the global var *** maybe "fix" this ??
        xRootOfyC(selection, workPiece, radical_index)
        workPiece++
    }
}
// case 180: /// note the C suffix on this and on all func-s used in this func /// should i, can i, just use the non-suffixed func-s ???
func xRootOfyC(selection int, workPiece int, radical_index int) { // calculates either square or cube root of any integer 

    var index = 0 // counter used in the for loop in this func :: is also passed to the pricipal func readTheTableOfPP // cannot be a global

        TimeOfStartFromTop := time.Now()

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

    buildTableOfPerfectProductsC(radical_index) // 825,000 entries

// The following section consists of the principal for loop with a conditional break ------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------------------------------------------------------
// case 180:

        startBeforeCall := time.Now()

    //for index < 380000 { // the table has 800,000 entries, 400,000 pairs; so index increments by 2 at the bottom of this loop 
    for index < 400000 { // the table has 825,000 entries, > 410,000 pairs; so index increments by 2 at the bottom of this loop 


        readTheTableOfPPC(index, startBeforeCall, radical_index) // pass-in an index to the table: 380,000 indexs corresponding to the number of pairs of entries 

        handlePerfectSquaresAndCubesC(TimeOfStartFromTop, radical_index, selection) // handle the rare case of a perfect square or cube 

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
        elapsed_s2 := t_s2.Sub(TimeOfStartFromTop)


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
    //------------------ case 180: ---------------------------------------------------------


        TotalRun := elapsed_s2.String() // cast time durations to a String type for Fprintf "formatted print"
        _ , err57 := fmt.Fprintf(fileHandle_180, "Total run was %s \n ", TotalRun) 
            check(err57)

        fileHandle_180.Close()
}

// in case 180: ----------
func handlePerfectSquaresAndCubesC(TimeOfStartFromTop time.Time, radical_index int, selection int) {
// the next sub-section detects, traps, and reports the detection of either a perfect square of a perfect cube ------------------
// ... it also is responsible for causing the algorithm to terminate via a break if workpiece was a perfect square or cube 
//-------------------------------------------------------------------------------------------------------------------------------
        if diffOfLarger == 0 || diffOfSmaller == 0 {  // Then, it was a perfect square or cube 

            t_s1 := time.Now() 
            elapsed_s1 := t_s1.Sub(TimeOfStartFromTop)

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
        for iter < 825000 { // a table of 825,000 pairs: PPs with their roots. That ought to do it !!
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
// end of case 180: -- AMFxroyCCB



// case 18: -- AMFxRootOfyA
func xRootOfy(selection int) { // calculates either square or cube root of any integer 

    usingBigFloats = false

    var index = 0 // counter used in the for loop in this func :: is also passed to the pricipal func readTheTableOfPP 

        TimeOfStartFromTop := time.Now()

    radical_index := getInputFromUserAndSetPrecision() 

    buildTableOfPerfectProducts(radical_index) // 800,000 entries, 400,000 pairs 

// The following section consists of the principal for loop with a conditional break ------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------------------------------------------------------

        startBeforeCall := time.Now()

    for index < 400000 { // the table has 825,000 entries, > 410,000 pairs; so index increments by 2 at the bottom of this loop (200,000 iters)

        readTheTableOfPP(index, startBeforeCall, radical_index) // pass-in an index to the table: 400,000 indexs corresponding to the number of pairs of entries 

        handlePerfectSquaresAndCubes(TimeOfStartFromTop, radical_index, selection) // handle the rare case of a perfect square or cube 

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
        elapsed_s2 := t_s2.Sub(TimeOfStartFromTop)


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
    //---- we are still in case 18: -------------------------------------------------------------------------


        TotalRun := elapsed_s2.String() // cast time durations to a String type for Fprintf "formatted print"
        _ , err57 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
            check(err57)

        fileHandle.Close()

                    /*
                    // all this crap with these 3 arrays was cute, but it does not sort as a record with fields, so it is not what I want 
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
}  // --- we are in case 18: 

func handlePerfectSquaresAndCubes(TimeOfStartFromTop time.Time, radical_index int, selection int) {
// the next sub-section detects, traps, and reports the detection of either a perfect square of a perfect cube ------------------
// ... it also is responsible for causing the algorithm to terminate via a break if workpiece was a perfect square or cube 
//-------------------------------------------------------------------------------------------------------------------------------
        if diffOfLarger == 0 || diffOfSmaller == 0 {  // Then, it was a perfect square or cube 

            t_s1 := time.Now() 
            elapsed_s1 := t_s1.Sub(TimeOfStartFromTop)  // need to pass this to the func we are planning to build ?? NO, "two" "perfect".

            fileNameToWriteTo := "dataLog-From_calculate-pi-and-friends.txt" // would have been used/needed if we emplement a func for this.

                //fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) 
                fileHandle, err1 := os.OpenFile(fileNameToWriteTo, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
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
//-------------------------------------------------------------  CASE 18: ------------------------------------------------------------------

}

func readTheTableOfPP(index int, startBeforeCall time.Time, radical_index int) {  // this gets called 400,000 times. 

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
// we are in case 18:

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

// we are in case 18:
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
                // ***** ^^^^ ****** the preceeding was replaced with the following five lines *******************************************

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
                //-------------  we are in case 18:   we are in case 18:   we are in case 18:   we are in case 18:   we are in case 18: ----------------

            break // each time we find a prospect we break out of the for loop --- if we found any prospects using the current index value we break 

        } // end of if :: if largerPerfectProduct > smallerPerfectProductOnce*workPiece  //  we only handle reads that were big enough to be prospects 
     }   // this is the end of the afforementioned for loop that we break out of each time we have found a prospect and handled it 
}       // the end of the readTheTableOfPP func that gets called 380,000 times

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
//----- still in case 18: ---------------------------------------------------------------------------------

func buildTableOfPerfectProducts(radical_index int) {

    var PerfectProduct int 
    Table_of_perfect_Products = nil // this fixed my bug 
    root := 10 
            iter := 0
        for iter < 825000 { // a table of 825,000 pairs: PPs with their roots. That ought to do it !!
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
}  // end of xRootOfy() // case 18: set, -- AMFxRootOfyB



func GottfriedWilhelmLeibniz(selection int){  // case 6: -- AMFGottfriedWilhelmLeibnizA
    usingBigFloats = false
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
//      sum = 1-(1/denom)

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
}  // end GottfriedWilhelmLeibniz() // case 6: // -- AMFGottfriedWilhelmLeibnizB


func GregoryLeibniz(selection int){  // case 7: -- AMFGregoryLeibnizA
    usingBigFloats = false
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
}  // end of GregoryLeibniz() // case 7: // -- AMFGregoryLeibnizB


func ArchimedesBig (selection int) { // case 14: // --AMFArchimedesBigA
    usingBigFloats = true 
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
            fmt.Printf(" A peek at the result formatted 1500f is: %.1500f \nper Archimedes'\n", p_d)  // show the first 1,500 digits of calculated pi
            fmt.Println(iterInt64, " iterations were completed, \n ... which generated a ", numberOfSides, "sided polygon!!\n")
            //fmt.Println(iterInt64, " iterations were completed yielding 2,712 correct digits of π!!!\n")
            fmt.Println("Go's math/big objects were set to a precision value of:", precision)

            fmt.Println("\n\n\n")
            printResultStatsLong(p_d, precision, "AM", 1, "", selection) // sumBig *big.Float, precision int, useAlternateFile string

            fmt.Println("\n\nThese digits are also in a text file named: dataLog-From_AM_Method_lengthy_prints.txt")
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
// end of :: ArchimedesBig (selection int) { // case 14:  // --AMFArchimedesBigB



func JohnWallis(selection int){  // case 8: // -- AMFJohnWallisA
    usingBigFloats = false
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
}  // end of JohnWallis() // case 8: // -- AMFJohnWallisB


func EulersNumber(selection int){ // case 9: // -- AMFEulersNumberA
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
}  // end of EulersNumber() // -- AMFEulersNumberB


func ErdosBorwein(selection int){  // case 10: // -- AMFErdosBorweinA
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
fmt.Println("") 

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
}  // // -- AMFErdosBorweinB



// and now for the spigot func itself ... (it's just so tight!)
func TheSpigot(selection int){  // case 19: // -- AMFTheSpigotA
    usingBigFloats = false
    var numberOfDigitsToCalc int
    fmt.Println("How much pi can you handle?")
    fmt.Println("How many digits of pi do you really want? Enter that number now:")  // prompt the user
        fmt.Scan(&numberOfDigitsToCalc)
        fmt.Println("")

    Spigot(numberOfDigitsToCalc, selection)

    fmt.Print("\n\nThis trick was made possible by a bit of code I mooched off of GitHub ...\n\n")
    fmt.Println("... to view the code with attribution Enter '39' at either menu\n")
    fmt.Println("To view an explaination of how it works enter 99 at either menu\n")
}

//func Spigot(n int) string { // called by the previous func
func Spigot(n int, selection int) { // Rick's version does not return a string // called by the previous func
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
}  // end TheSpigot() set // -- AMFTheSpigotB



// case 40: -- AMFnifty_scoreBoardA
// A concurrent computation of pi using Nilakantha's formula.
// by Diego Brener diegosilva13 on Github 
// ******* nifty scoreboard ***********************************
func nifty_scoreBoard() { // case 40: 
        usingBigFloats = false
        // We use a ticker to specify the interval to update the values on the scoreboard
        ticker := time.NewTicker(time.Millisecond * 108)

        // If the user wants to prematurely break out of the program by issuing a Ctrl+C, we tell the 
        // ... signal package to notify us over this interrupt channel
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
        // We want to create a virtual scoreboard where we can simutaneously show the current value of Pi and 
        // ... also how many Nilakantha terms we have calculated.

        // We can create the virtual scoreboard by using some ANSI escape codes
        // Here's a wikipedia article giving you the complete list of ANSI escape
        // codes: https://en.wikipedia.org/wiki/ANSI_escape_code

    // convenience globals: 
        const ANSIClearScreenSequence = "\033[H\033[2J"
        const ANSIFirstSlotScreenSequence = "\033[2;0H"
        const ANSISecondSlotScreenSequence = "\033[3;0H"

        // The channel used to update the current value of pi on the scoreboard
        var pichan chan float64 = make(chan float64)

        // The channel that we use to indicate that the program can exit
        var computationDone chan bool = make(chan bool, 1)

        // Number of Nilakantha terms for the scoreboard
        var termsCount int

// The following function serves as our virtual scoreboard to show the current
// computed value of Pi using Nilakantha's formula

func printCalculationSummary() {

    fmt.Print(ANSIClearScreenSequence)
    fmt.Println(ANSIFirstSlotScreenSequence, "Computed Value of Pi:\t\t", <-pichan)
    fmt.Println(ANSISecondSlotScreenSequence, "# of Nilakantha Terms:\t\t", termsCount)
}

// We are going to use Nilakantha's formula from the Tantrasamgraha (which is more than 500 years old) 
// ... to compute the value of Pi to several decimal points
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
} // end of nifty_scoreBoard() set // -- AMFnifty_scoreBoardB



// case 41: // -- AMFBBPfConcurentA
func BBPfConcurent(selection int){  
    usingBigFloats = false
    workerCount := 4
    prescribedIterations := 8

    fmt.Println("Enter the desired number of workers (recomend 2 to 16, 4 is optimal) ")
    fmt.Scan(&workerCount)

    fmt.Println("Enter your number of iterations (recomend 2 to 16, 8 being optimal) ")
    fmt.Scan(&workerCount)

    start := time.Now()

    pi := PiMultiThread(workerCount, prescribedIterations)

    t := time.Now()
    elapsed := t.Sub(start)
    TotalRun := elapsed.String()
    fmt.Println("elapsed time was: ", TotalRun)

        piAsBF := new(big.Float)
        piAsBF = big.NewFloat(pi) // pi is being cast to big from float64
        printResultStatsLong(piAsBF, prescribedIterations, "BBPfConcurent", workerCount, TotalRun, selection)
}
// global structure 
type Float64 struct {
    value float64
    lock  sync.RWMutex
}

// helper functions: 
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
            //fmt.Println(ret) // Rick's progress report, one line for each worker ? 
                //fmt.Printf("\nWorker number %d came up with: ", workers) // Rick's progress report
                //fmt.Print(ret)
                //fmt.Print("\n") 
        })(i)
            fmt.Printf("\nWorkers numbering %d came up with: ", workers) // Rick's progress report
            fmt.Print(ret)
            fmt.Print("\n") 
    }
    wg.Wait()
    return ret.Get()
}  // end of func BBPfConcurent() set // -- AMFBBPfConcurentB



func bbp_formula (selection int){ // case 42: // -- AMFbbp_formulaA
usingBigFloats = true
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
}  // end of bbp_formula() set // -- AMFbbp_formulaB


// case 44: // -- AMFLeibniz_method_one_billion_itersA
func Leibniz_method_one_billion_iters(selection int){ 
    usingBigFloats = false
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

// Rick's code follows:
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
}  // end of Leibniz_method_one_billion_iters() // -- AMFLeibniz_method_one_billion_itersB



func check(e error) {   // create a func named check which takes one parameter "e" of type error 
    if e != nil {
        panic(e)        // use panic() to display error code 
    }
}



// gloabl vars for use in multiple localized sections of code =========================================================
// convenience globals: 
        var usingBigFloats = false    // a variable of type bool which is passed by many funcs to printResultStatsLong()

        var colorReset = "\033[0m"          
        var colorRed = "\033[31m"
        var colorGreen = "\033[32m"
        var colorYellow = "\033[33m"
        var colorPurple = "\033[35m"
        var colorCyan = "\033[36m"
        var colorWhite = "\033[37m"

        var TotalIterations int 
        var four float64  // is initialized to 4 where needed 
        var π float64    // a var can be any character, as in this Pi symbol/character 
        var LinesPerSecond float64  
        var LinesPerIter float64
        var iterInt64 int64  // to be used primarily in selections which require modulus calculations 
        var iterFloat64 float64  // to be used in selections which do not require modulus calculations 
        var Table_of_perfect_squares = []int{}
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

        type Results struct {  // define a new structure called Results with two fields; result, and pdiff 
            result float64
            pdiff float64
        }
//========== end of global var section ================================================================================


// case 2: // -- AMFBBPFA
func BBPF(selection int) {
    usingBigFloats = true
// Richard's modifications: -------------------------------------
    useAlternateFile := "BBPF" 

    numberOfDigits := 1 // just to create it 
        fmt.Println("\nYou selected #", selection, "the Bailey–Borwein–Plouffe formula for π, circa 1995\n")
        fmt.Println("This version uses channels: GOMAXPROCS(numCPU), and big floats; how many digits of π would you like?")

    fmt.Scan(&numberOfDigits) 

    if numberOfDigits > 25000 {
        fmt.Printf("\nYou requested %d digits of pi. Which is kinda a lot. 25,000 would have taken a minute. \n", numberOfDigits)
        fmt.Println("The amount you asked for could take much longer. Even though I am hammering all of your cores.\n")
    }
// --------------------------------------------------------------

    numCPU := runtime.NumCPU()
    runtime.GOMAXPROCS(numCPU)    
    p := uint((int(math.Log2(10)))*numberOfDigits + 3)

// Richards's mods follow: --------------------------------------------------------
        fmt.Println("log2 etc. (based on numberOfDigits) has set p at: ", p)

    additionalAmount := 1.2 // just have to create this outside these ifs down thar 

    if numberOfDigits <= 12 {
        additionalAmount = float64(p) + (float64(p) * 0.127)
    } else if numberOfDigits <= 500 {
        additionalAmount = float64(p) + (float64(p) * 0.125)
    } else if numberOfDigits <= 5000 {
        additionalAmount = float64(p) + (float64(p) * 0.120)
    } else if numberOfDigits <= 10000 {
        additionalAmount = float64(p) + (float64(p) * 0.119)
    } else if numberOfDigits > 10000 {
        additionalAmount = float64(p) + (float64(p) * 0.115)
    }

    wholePart := uint(math.Floor(additionalAmount))
    p = wholePart
    if p > 85400 {
        fmt.Printf("\n ... we have adjusted that to: %d ... working ...", p)
    } else {
        fmt.Printf("\n ... we have adjusted that to: %d", p)
    }
// Richard's mod end: -------------------------------------------------------------

    result := make(chan *big.Float, numberOfDigits)
    worker := workers2(p)

    pi := new(big.Float)

// Richards modification of SetPrec:
    pi.SetPrec(p) // this gives a lot more results

    for i := 0; i < numberOfDigits; i++ {
        go worker(i, result)
    }
    for i := 0; i < numberOfDigits; i++ {
        pi.Add(pi, <-result)
    }

// Richards's mods follow: --------------------------------------------------------
    fmt.Printf("\n\na peek at pi formatted 250f is: %[1]*.[2]*[3]f \n", 1, 250, pi)
    printResultStatsLong(pi, int(p), useAlternateFile, 1, "", selection) // p is just the precision that was used for the big floats, which will be reported by printResults... 
// Richard's mods end: -------------------------------------------------------------

}
func workers2(p uint) func(id int, result chan *big.Float) {
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
}  // end of BBPF() set // -- AMFBBPFB



// case 15: // -- AMFchudA
// Chudnovsky method, based on https://arxiv.org/pdf/1809.00533.pdf 
/*
The Chudnovsky algorithm is an incredibly fast algorithm for calculating the digits of pi. It was developed by Gregory Chudnovsky and his 
brother David Chudnovsky in the 1980s. It is more efficient than other algorithms and is based on the theory of modular equations. It has 
been used to calculate pi to over 62 trillion digits.
*/
        //  Using this procedure, calculating 1,000,000 digits requires 70516 loops, per the run on:
        //  Sun May  7 08:50:23 2023
        //  Total run was 8h4m39.7847064s
        // AND, THAT CALCULATION WAS INDEPENDANTLY VERIFIED !!!!!!!!!!!

func chud(selection int) {  // case 15:
    usingBigFloats = true
    var digits int 
    var loops int 
        start := time.Now() // saved start time to be compared with end time t 
    fmt.Println("\nRichard invites you to enter the number of digits of pi to calculate per the Chudnovsky method: \n")
    fmt.Println("The sky is the limit with this method, so don't be shy.  \n")

            /* Unix variant 
                fmt.Scanf("%d", &digits) 
            Unix variant */ 
            //
            // /* Windows variant 
                fmt.Scanf("%d", &digits) 
                fmt.Scanf("%d", &digits) 
            // Windows variant */ 

pi := new(big.Float)
loops, pi = CalcPi(float64(digits), start, loops, selection) // returns i which gets a new name "loops"
//          CalcPi is responsible for most of its printing and logging 

if loops < 100 {
    fmt.Printf("\na peek at the prospective value of Pi as a big float, and formatted 0.122f, is : \n%0.122f \n\n", pi)

    fmt.Println("... but : ") // but, also deploy printResultStatsLong() 
    printResultStatsLong(pi, int(digits), "ChudDidLessThanOneHundredLoops", 1, "", selection)
}

// always do the following elapsed time and conditional print work at the end of this selection (case 15:)
// none of the following will be reached if we exit out of CalcPi with an os.Exit(), so we break instead 
    // for now we take this out : 
    //fmt.Printf("%1.[1]*[2]f \n", digits, pi)
    fmt.Printf("\n loops were %d, and digits requested was %d \n", loops, digits)
    //fmt.Printf("\nprecR is: %d, and precision is, as yet, undefined near the top of this case.\n", precR)
    //os.Exit(1) // was testing
        t := time.Now()
        elapsed := t.Sub(start)
        // only if ===== all this print work is conditional on some time having elapsed ==================================================
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
                // the whole pi would be printed to the datalog file on the line below
                    //_ , err8 := fmt.Fprintf(fileHandle, "pi was %1.[1]*[2]f \n", digits, pi) 
                    //    check(err8)
                // ... after printing the whole pi, some nice stats are appended to the file's log entry 
                    TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                    _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s, and digits requested was %d \n ", TotalRun, digits) 
                        check(err7)
            }
} 
// calculate Pi for n number of digits
func CalcPi(digits float64, start time.Time, loops int, selection int) (int, *big.Float) {
    usingBigFloats = true
    var i int 
    /**
    *   This is an implementation for https://en.wikipedia.org/wiki/Chudnovsky_algorithm
    *   "It can be improved using binary splitting http://numbers.computation.free.fr/Constants/Algorithms/splitting.html
    *   if we split it into two independent parts and simplify the formula for more details: 
    * https://www.craig-wood.com/nick/articles/pi-chudnovsky/""
    */

    // n, apparently, is the expected number of loops we may need to produce digits number of digits 
    n := int64(2 + int(float64(digits)/14.181647462))
    //n := int64(2 + int(float64(digits)/12))  // I tried this, and may try something like it again someday?? like /14.0 ? 
    //prec := uint(int(math.Ceil(math.Log2(10)*digits)) + int(math.Ceil(math.Log10(digits))) + 2) // the original 
    //prec := uint(digits) // not good, not large enough, so ...
    digitsPlus := digits + digits*0.10  // because we needed a little more than the orriginal programmer had figured on :) 
    prec := uint(int(math.Ceil(math.Log2(10)*digitsPlus)) + int(math.Ceil(math.Log10(digitsPlus))) + 2)

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

// this is a flag; if it is set to zero we exit 
    queryIfTimeToDie := 1 
    i = 1 // a secondary dedicated loop counter

    if n > 8998 {
        fmt.Println("\n Well, this is going to take a while, because you asked for too much pie (> 8990)\n")
    }

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

        useAlternateFile := "no" // no means to use the standard log file rather than some special one 
        // the compiler is not happy unless it sees this created outside of an if 
        // But, wait. Why is the compiler allowing me to violate the no new var left of the := assignment ??? This IS in a loop !!!!
        if i == 100 {
            //useAlternateFile := "no" // the compiler is not happy unless it sees this created outside of an if
            fmt.Printf("\n we are at %d loops, here comes a 800f of pi as a big float: \n", i)
            fmt.Printf("%0.[1]*[2]f \n", 800, pi)
            //fmt.Printf("%0.[1]*[2]f \n", int(digits), pi) // if digits was known to be the verified string of digits, then this is what we would want 
            finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
        }
        if i == 200 {
            //useAlternateFile = "no" // still no
            fmt.Printf("\n we are at %d loops, here comes a 999f of pi as a big float: \n", i)
            fmt.Printf("%.999f", pi)
            finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
        }
        if i == 400 {
            //useAlternateFile = "no" // still no
            fmt.Printf("\n we are at %d loops, here comes a 1599f of pi as a big float: \n", i)
            fmt.Printf("%.1599f", pi)
            finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
        }
        //
        //
        // note below the: useAlternateFile = "chudDid800orMoreLoops" 
        if i == 800 {
            useAlternateFile = "chudDid800orMoreLoops"
            fmt.Printf("\n we are at %d loops, here comes a 1999f of pi as a big float: \n", i)
            fmt.Printf("%.1999f", pi)
            finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
        }
        if i == 1600 {
            useAlternateFile = "chudDid800orMoreLoops"
            fmt.Printf("\n we are at %d loops, here comes a 2200f of pi as a big float: \n", i)
            fmt.Printf("%.2200f", pi)
            finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
        }
        if i == 2000 {
            useAlternateFile = "chudDid800orMoreLoops"
            fmt.Printf("\n we are at %d loops, here comes a 2300F of pi as a big float: \n", i)
            fmt.Printf("%.2300F", pi)
            finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
        }
        if i == 2400 {
            useAlternateFile = "chudDid800orMoreLoops"
            fmt.Printf("\n we are at %d loops, here comes a 2600F of pi as a big float: \n", i)
            fmt.Printf("%.2600f", pi)
            finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
        }
        if i == 2800 {
            useAlternateFile = "chudDid800orMoreLoops"
            fmt.Printf("\n we are at %d loops, here comes a 2900F of pi as a big float: \n", i)
            fmt.Printf("%.2900f", pi)
            finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
        }
        if i == 3200 {
            useAlternateFile = "chudDid800orMoreLoops"
            fmt.Printf("\n we are at %d loops, here comes a 3100F of pi as a big float: \n", i)
            fmt.Printf("%.3100f", pi)
            finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
        }
        if i == 4000 {
            useAlternateFile = "chudDid800orMoreLoops"
            fmt.Printf("\n we are at %d loops, here comes a 3300F of pi as a big float: \n", i)
            fmt.Printf("%.3300f", pi)
            finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
        }
        if i == 6000 {
            useAlternateFile = "chudDid800orMoreLoops"
            fmt.Printf("\n we are at %d loops, here comes a 3400F of pi as a big float: \n", i)
            fmt.Printf("%.3400f", pi)
            finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
        }
        if i == 8000 {
            useAlternateFile = "chudDid800orMoreLoops"
            fmt.Printf("\n we are at %d loops, here comes a 3500F of pi as a big float: \n", i)
            fmt.Printf("%.3500f", pi)
            finishChudIfs(pi, digits, selection, i, useAlternateFile, queryIfTimeToDie)
        }
        if queryIfTimeToDie == 0 {
                fmt.Println("if queryIfTimeToDie is 0, time to die")
                fmt.Printf("\nprecision was: %d \n", prec)
                break
        }
        // 1,000,000 digits requires 70516 loops, per the run on May 7 2023 at 10:30
        //  was run on: Sun May  7 08:50:23 2023
        //  Total run was 8h4m39.7847064s
        // AND THE CALCULATION WAS INDEPENDANTLY VERIFIED !!!!!!!!!!!
    } // end of for loop way up thar :: it prompts periodically to continue or die

    // we are out of the loop, so we do the following just once:

            fileHandleBig, err1prslc2c := os.OpenFile("big_pie_is_in_here.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                check(err1prslc2c)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                defer fileHandleBig.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.

            _ , err9bigpie := fmt.Fprint(fileHandleBig, pi) // dump this big-assed pie to a special log file
                check(err9bigpie)

            _, errGoesHere := fmt.Fprint(fileHandleBig, "\n\n")
                check(errGoesHere)

            fileHandleBig.Close() 

    return i, pi // asigning i to loops 
} 
// a helper func 
func finishChudIfs(pi *big.Float, digits float64, selection int, i int, useAlternateFile string, queryIfTimeToDie int) {
            printResultStatsLong(pi, int(digits), useAlternateFile, 1, "", selection) 
            fmt.Printf("\n the above was from %d loops \n", i)
            fmt.Println("enter any int to continue, 0 to end")
            /* Unix variant 
                fmt.Scanf("%d", &queryIfTimeToDie)
            Unix variant */ 
            //
            // /* Windows variant 
                fmt.Scanf("%d", &queryIfTimeToDie) 
                fmt.Scanf("%d", &queryIfTimeToDie) // Windows environments seem to just fly past a single Scanf, so I have used two :) 
            // Windows variant */ 
} // end of chud() set // -- AMFchudB



func deleteAllLogFiles() {  // case 999: // -- AMFdeleteAllLogFilesA
    // list the files to be deleted: 
    filename111 := "dataLog-From_calculate-pi-and-friends.txt"
    filename112 := "dataLog-From_BBPfConcurent.txt"
    filename113 := "dataLog-From_Nilakantha_Method_lengthy_prints.txt"
    filename114 := "dataLog-From_AM_Method_lengthy_prints.txt"
    filename115 := "dataLog-From_BBPF_Method_lengthy_prints.txt"
    filename116 := "dataLog-From_Chudnovsky_Method_lengthy_prints.txt"
    filename117 := "big_pie_is_in_here.txt"
    filename118 := "logfile_from_selection_180.txt"
    // Check if those files exist:
    if _, err := os.Stat(filename111); err != nil {
        fmt.Println("File dataLog-From_calculate-pi-and-friends.txt does not exist")
    }
    if _, err := os.Stat(filename112); err != nil {
        fmt.Println("File dataLog-From_BBPfConcurent.txt does not exist")
    }
        if _, err := os.Stat(filename113); err != nil {
        fmt.Println("File dataLog-From_Nilakantha_Method_lengthy_prints.txt does not exist")
    }
        if _, err := os.Stat(filename114); err != nil {
        fmt.Println("File dataLog-From_AM_Method_lengthy_prints.txt does not exist")
    }
        if _, err := os.Stat(filename115); err != nil {
        fmt.Println("File From_BBPF_Method_lengthy_prints.txt does not exist")
    }
        if _, err := os.Stat(filename116); err != nil {
        fmt.Println("File dataLog-From_Chudnovsky_Method_lengthy_prints.txt does not exist")
    }
        if _, err := os.Stat(filename117); err != nil {
        fmt.Println("File big_pie_is_in_here.txt does not exist")
    }
        if _, err := os.Stat(filename118); err != nil {
        fmt.Println("File logfile_from_selection_180.txt does not exist")
    }
    // Delete those files:
    err := os.Remove(filename118)
    if err != nil {
        fmt.Println(err)
    }
    err1 := os.Remove(filename111)
    if err1 != nil {
        fmt.Println(err)
    }
    err2 := os.Remove(filename112)
    if err2 != nil {
        fmt.Println(err)
    }
    err3 := os.Remove(filename113)
    if err3 != nil {
        fmt.Println(err)
    }
    err4 := os.Remove(filename114)
    if err4 != nil {
        fmt.Println(err)
    }
    err5 := os.Remove(filename115)
    if err5 != nil {
        fmt.Println(err)
    }
    err6 := os.Remove(filename116)
    if err6 != nil {
        fmt.Println(err)
    }
    err7 := os.Remove(filename117)
    if err7 != nil {
        fmt.Println(err)
    }
}  // end of deleteAllLogFiles() // -- AMFdeleteAllLogFilesB


// case 5: // -- AMFNilakantha_Somayaji_with_big_Float_typesA
func Nilakantha_Somayaji_with_big_Float_types(selection int) {

    usingBigFloats = true

    fmt.Println("\nYou have selected the Nilakantha Somayaji method using big.Float types, and with some ")
    fmt.Print("patience one can generate 31 correct digits of pi using it.\n\n")

    fmt.Println("Enter the number of iterations (suggest between 100,000 and 100,000,000)")

    var iters int

            /* Unix variant 
                fmt.Scanf("%d", &iters) 
            Unix variant */ 
            // 
            // /* Windows variant 
                fmt.Scanf("%d", &iters) 
                fmt.Scanf("%d", &iters) 
            // Windows variant */ 

    fmt.Println("Enter the precision: (suggest between 128 and 512)")

    var precision int 

            /* Unix variant 
                fmt.Scanf("%d", &precision) 
            Unix variant */ 
            //
            // /* Windows variant 
                fmt.Scanf("%d", &precision) 
                fmt.Scanf("%d", &precision) 
            // Windows variant */ 

    if iters > 36111222 {
        fmt.Println(" ... ")
    }
    if iters > 42000000 {
        fmt.Println("... werkin ...")
    }
    if iters > 55111222 {
        fmt.Println("... working for a while ...")
    }
    if iters > 69111222 {
        fmt.Println("... will be working for quite a while ...")
    }
    if iters > 80111222 {
        fmt.Println("... a very long while ... working ...")
    }

    start := time.Now()
    var iterBig int


    // big.Float constants:

        twoBig := big.NewFloat(2)
        threeBig := big.NewFloat(3)
        fourBig := big.NewFloat(4)

    // big.Float variables:

        digitoneBig := new(big.Float)
        *digitoneBig = *twoBig 

        digittwoBig := new(big.Float)
        *digittwoBig = *threeBig 

        digitthreeBig := new(big.Float)
        *digitthreeBig = *fourBig

        sumBig := new(big.Float) 
        nexttermBig := new(big.Float) 

    // set precision to a user-specified value 
        sumBig.SetPrec(uint(precision))
        twoBig.SetPrec(uint(precision))
        threeBig.SetPrec(uint(precision))
        fourBig.SetPrec(uint(precision))
        digitoneBig.SetPrec(uint(precision))       
        digittwoBig.SetPrec(uint(precision))
        digitthreeBig.SetPrec(uint(precision))
        nexttermBig.SetPrec(uint(precision))

        sumBig.Add(threeBig,  new(big.Float).Quo(fourBig, new(big.Float).Mul(digitoneBig, new(big.Float).Mul(digittwoBig, digitthreeBig))))

                    iterBig = 1
        for iterBig < iters { // 1,000,000,000 yeilds 19 digits in 13 sec

                        // 256p 100,000,000 : 56s 25 digits digits June 26 2023 
                        // 1,280p and 1Bil : 23 min without ending
                        // 128p and 100,000,000 49s gave 25 digits June 26 2023 
                        // Total run with SetPrec at: 128 and iters of 1,000,000,000 was 7m57.3179415s
                        // got 31 digits in 1 hour and 26 min using this algorithm with one billion iters at 128 prec
                        // 1,000,000,002 and 64 bits prec yielded :: 17 digits in 5m41s
            
            iterBig++

            digitoneBig.Add(digitoneBig, twoBig) 
            digittwoBig.Add(digittwoBig, twoBig) 
            digitthreeBig.Add(digitthreeBig, twoBig)

            nexttermBig.Quo(fourBig, new(big.Float).Mul(digitoneBig, new(big.Float).Mul(digittwoBig, digitthreeBig)))

                if iterBig % 2 == 0 {  // % is modulus operator 
                    sumBig.Sub(sumBig, nexttermBig) 
                } else {
                    sumBig.Add(sumBig, nexttermBig)
                }

                if iterBig == 20111222 {
                    fmt.Println(" ... doin some ... ")
                }
                if iterBig == 36111222 {
                    fmt.Println(" ... werkin ... ")
                }
                if iterBig == 42000000 {
                    fmt.Println("... still werkin ...")
                }
                if iterBig == 55111222 {
                    fmt.Println("... been working for a while ...")
                }
                if iterBig == 69111222 {
                    fmt.Println("... been working for quite a while ...")
                }
                if iterBig == 80111222 {
                    fmt.Println("... it's been a very long while ... but still working ...")
                }
                if iterBig == 180111222 {
                    fmt.Println("... it's been a very long while, 180,111,222 down, ... and still working ...")
                }
                if iterBig == 280111222 {
                    fmt.Println("... it's been a very long while, 280,111,222 down, ... and still working ...")
                }
                if iterBig == 480111222 {
                    fmt.Println("... it's been a very long while, 480,111,222 down, ... still working ...")
                }
                if iterBig == 680111222 {
                    fmt.Println("... it's been a very long while, 680,111,222 down, ...  working ...")
                }
                if iterBig == 880111222 {
                    fmt.Println("... it's been a very long while, down, 880,111,222, down ... still, working ...")
                }
                if iterBig == 977111222 {
                    fmt.Println("... it's been a very long while, 977,111,222 already ... why am I still working? ...")
                }
        }
    t := time.Now()
    elapsed := t.Sub(start)     
    TotalRun := elapsed.String()    
        // print to file:
        printResultStatsLong(sumBig, precision, "Nilakantha", iterBig, TotalRun, selection) 
        // print to screen:
    fmt.Println(" via Nilakantha \n")
    fmt.Printf("Total run with SetPrec at: %d and iters of %d was %s \n\n ", precision, iterBig, TotalRun) 
} // end of Nilakantha_Somayaji_with_big_Float_types() // -- AMFNilakantha_Somayaji_with_big_Float_typesB


func Using_this_app() {  // case 96: 
    fmt.Print(string(colorCyan))
    var rune_Using_this_app = `
Any selection from any menu can be made at any menu. 

If you have instantiated this app from the command line with "go run appname.go", or have the source-code .go file installed,  
each selection has a corresponding selector which displays the source code for that particular algorithm. For example, to view the code 
for selection #18 one simply enters 38 at any menu – one reason that you might want to do this is to discover the section’s authorship. 

Additionally, there are a few other operational selections which are not found on any of the menus. For example, can enter 33 to see the 
magic behind main, or enter 120 to see the code for the switches, and you can enter 95 to read about the state of the art when it comes to 
calculating maximum digits of π. 

===============
About this app:
===============

The majority of the source code which comprises this app was conceived of, designed, and implemented by Richard (Rick) Hart Wolley in late 
2022 and early 2023. Sections of code that were mooched off GitHub or other sites have proper attributions which are viewable as per the 
instructions given above.  

Why does this app exist? Well, it was a rather rainy day sometime late in October and I had some time to kill. I had not done any 
software engineering for a few years and there were two languages that I had never really tried before, so that seemed like fun. 
Fortran, and Python were on my bucket list. I had also been hearing a lot of good things about Google’s new language go.lang (simply Go 
within the inner circles at Google). Revisiting Lisp (both the Emacs and the Common variants) had also been on my mind but I’ve yet to
get around to it. I had also always wanted to try constructing an algorithm that would calculate Pi. I was especially curious to see how 
many digits one could easily calculate from first principles using a home computer and a simple algorithm. I coded up an identical 
prototype in the three languages that were new to me and found that Go was so much better in every way that I can now not imagine messing 
around with any other language. Though, admittedly, I have found that Go is a bit “buggy” on Windows11, Go being intended mainly for use 
on Unix variants. Which is Ok because Linux and Mac are my preferred programming environments. 

I then got a bit carried away, and a few thousand lines of code later here we are. 

The sections that I am the most proud of are #1, #18, and #4/#14. Two variants of brute forcing the extraction of 
irrational roots|Radicals|radicalis, and the geometric derivation of Pi respectively.

    `
    fmt.Println(rune_Using_this_app, string(colorReset))
}


// case 50: // -- AMFcompareFloat64withBigFloatsA
// comparison of float64 vs big.Float types using a Nilakantha Somayaji example 
// π = 3 + 4/(2*3*4) - 4/(4*5*6) + 4/(6*7*8) - 4/(8*9*10) + 4/(10*11*12) ...
// Rick just loves github.com, Sublime Text, and Go.lang 

func compareFloat64withBigFloats() {  // case 50:
    usingBigFloats = true
    fmt.Println("\n ... just a moment, working ... \n")
    start := time.Now()
        with_float64_types() // runs pretty quickly through tens of billions of iters 
        precision, iterBig := with_big_Float_types()  // can take forever with large values of prec and iters 
    t := time.Now()
    elapsed := t.Sub(start)
    TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
    fmt.Printf("\n\nTotal run with SetPrec at: %d and iters of %d was %s \n ", precision, iterBig, TotalRun) 
}

func with_float64_types() {
    usingBigFloats = false
    // float64 constants:

            var four64 float64    
                four64 = 4

            var three64 float64
                three64 = 3 

    // float64 variables:

            var digitone64 float64
                digitone64 = 2

            var digittwo64 float64
                digittwo64 = 3

            var digitthree64 float64
                digitthree64 = 4

            var sum64 float64    
                sum64 = three64 + (four64 / (digitone64*digittwo64*digitthree64)) 

            var nextterm64 float64

            iters_fl64 := 1
        for iters_fl64 < 10000000000 { // 10,000,000,000 ten billion 
            iters_fl64++
            digitone64 = digitone64 + 2 
            digittwo64 = digittwo64 + 2 
            digitthree64 = digitthree64 + 2 
            nextterm64 = four64/(digitone64*digittwo64*digitthree64)

                if iters_fl64 % 2 == 0 {  
                    sum64 = sum64 - nextterm64 
                } else {
                    sum64 = sum64 + nextterm64
                }
        }
        printResultStats64comp(sum64, iters_fl64)
}

func printResultStats64comp(sum64 float64, iterBig int) {
    stringOfSum := strconv.FormatFloat(sum64, 'f', 160, 64)
    piAs100chars := "3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706"

    fmt.Printf("\npi as calculated per float64s is: %0.76f \n", sum64)
    fmt.Println("                                                ^ ")
    fmt.Println("                                  1 234567890123456789012345678901234567890123456789012345")
    fmt.Println("                                           10   |    20        30        40        50     ")
    fmt.Printf("pi from the web is:               %s \n", piAs100chars)   


    posInPi := 0 // to be the incremented offset : piChar = piAs86chars[posInPi]
    var piChar byte  // one byte (character) of pi as string, e.g. piChar = piAs86chars[posInPi]
    var copyOfLastPosition int // an external (to the loop) copy of positionInString
    var stringVerOfCorrectDigits = []string{}  
    for positionInString, charAtRangePos := range stringOfSum {
        piChar = piAs100chars[posInPi] 
            if charAtRangePos == rune(piChar) {
                stringVerOfCorrectDigits = append(stringVerOfCorrectDigits, string(charAtRangePos))
                //fmt.Printf("we have agreement up to character pos: %d \n", positionInString)
                copyOfLastPosition = positionInString // save an external copy, of the last position found to have matched pi, as an int 
            } else {
                break // to print result and info below 
            }
        posInPi++
    }        
    fmt.Printf("\nWe have calculated pi correctly to %d digits using %d iters and float64s types \n", copyOfLastPosition, iterBig)
    fmt.Print(" .... that is ten Billion iterations!!! \n\n")
    fmt.Print("The correctly calculated digits are: ")
        for _, oneChar := range stringVerOfCorrectDigits {
            fmt.Print(oneChar)
        }
    fmt.Print("\n")
}

func with_big_Float_types() (precision uint, iterBig int) {
    usingBigFloats = true
    // big.Float constants:

            twoBig := big.NewFloat(2)
            threeBig := big.NewFloat(3)
            fourBig := big.NewFloat(4)

    // big.Float variables:

            digitoneBig := new(big.Float)
            *digitoneBig = *twoBig 

            digittwoBig := new(big.Float)
            *digittwoBig = *threeBig 

            digitthreeBig := new(big.Float)
            *digitthreeBig = *fourBig

            sumBig := new(big.Float) 
            nexttermBig := new(big.Float) 

        precision = 164 
            sumBig.SetPrec(precision)
            twoBig.SetPrec(precision)
            threeBig.SetPrec(precision)
            fourBig.SetPrec(precision)
            digitoneBig.SetPrec(precision)       
            digittwoBig.SetPrec(precision)
            digitthreeBig.SetPrec(precision)
            nexttermBig.SetPrec(precision)

            sumBig.Add(threeBig,  new(big.Float).Quo(fourBig, new(big.Float).Mul(digitoneBig, new(big.Float).Mul(digittwoBig, digitthreeBig))))

                    iterBig = 1
                for iterBig < 960000 { // 1,000,000,000
                    // Total run with SetPrec at:  128 and iters of 50000000  was 23.1879034s    :: 3.14159265358979323846264
                    // Total run with SetPrec at: 1024 and iters of 600000000 was 12m23.9554367s :: 3.14159265358979323846264338
                    /* We have calculated pi correctly to 28 digits using 1000000000 iters and Prec of 128
                        The correctly calculated digits are:                                        3.141592653589793238462643383
                        Total run with SetPrec at: 128 and iters of 1000000000 was 7m57.3179415s
                // got 31 digits in 1 hour and 26 min using this algorithm with one billion (must have been ten) iters at 128 prec
                    */
                    iterBig++

                    digitoneBig.Add(digitoneBig, twoBig) 
                    digittwoBig.Add(digittwoBig, twoBig) 
                    digitthreeBig.Add(digitthreeBig, twoBig)

                    nexttermBig.Quo(fourBig, new(big.Float).Mul(digitoneBig, new(big.Float).Mul(digittwoBig, digitthreeBig)))

                        if iterBig % 2 == 0 {  // % is modulus operator 
                            sumBig.Sub(sumBig, nexttermBig) 
                        } else {
                            sumBig.Add(sumBig, nexttermBig)
                        }
                }
                printResultStatsLongComp(sumBig, iterBig, int(precision))
                return precision, iterBig  // still is an int ^^^^^^^^^
}

func printResultStatsLongComp(sumBig *big.Float, iterBig int, precision int) {
    piAs100chars := "3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706"
    stringOfSum := sumBig.Text('f', 160) // create a string version of a big result, 700 some odd chars in length 

    fmt.Printf("\n\npi as calculated per big.Floats is: %0.76f \n", sumBig)
    fmt.Println("                                                   ^ ")
    fmt.Println("                                    1 234567890123456789012345678901234567890123456789012345")
    fmt.Println("                                             10    |   20        30        40        50     ")
    fmt.Printf("pi from the web is:                 %s \n", piAs100chars)        

    posInPi := 0 // to be the incremented offset : piChar = piAs86chars[posInPi]
    var piChar byte  // one byte (character) of pi as string, e.g. piChar = piAs86chars[posInPi]
    var copyOfLastPosition int // an external (to the loop) copy of positionInString
    var stringVerOfCorrectDigits = []string{}  
    for positionInString, charAtRangePos := range stringOfSum {
        piChar = piAs100chars[posInPi] 
            if charAtRangePos == rune(piChar) {
                stringVerOfCorrectDigits = append(stringVerOfCorrectDigits, string(charAtRangePos))
                //fmt.Printf("we have agreement up to character pos: %d \n", positionInString)
                copyOfLastPosition = positionInString // save an external copy, of the last position found to have matched pi, as an int 
            } else {
                break // to print result and info below 
            }
        posInPi++
    }        
    fmt.Printf("\nWe have calculated pi correctly to %d digits using %d iters and Prec of %d on big.Floats \n", copyOfLastPosition, iterBig, precision)
    fmt.Print(" ... and here it is only 960,000 iterations !!!\n\n")
    fmt.Print("The correctly calculated digits are: ")
        for _, oneChar := range stringVerOfCorrectDigits {
            fmt.Print(oneChar)
        }
    fmt.Print("\n")
}  // end of compareFloat64withBigFloats() set // -- AMFcompareFloat64withBigFloatsB


// case 37: // -- AMFGauss_LegendreA
func Gauss_Legendre(selection int) {  
    usingBigFloats = false
        start := time.Now()
    fmt.Printf("\n ... running case %d: \n\n", selection)
    var Ricks_value float64             // Rick's code 
    //var exterior_catcher int           // Rick's code 
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

    iters := 3 
    // run the algorithm iters times
    for i := 0; i < iters; i++ {  // call the 5 funcs (a,b,t,p, and pi) defined below, each of which returns just one []float64
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
        Ricks_value = value // Rick's code to grab that final 'value' from last iteration 
    //fmt.Printf("\n\nTop underscore is %d \n\n", exterior_catcher) it starts at 0 and goes to 3 

    }
        fmt.Println(Ricks_value) // Rick's code
        //fmt.Printf("\n\nBottom underscore is %d \n\n", exterior_catcher) this exterior_catcher var is never touched by the for loop 
        fmt.Println("3.1415926535897932 <-- compared to the actual value of Pi")
        fmt.Println("1 23456789012345 counting to fifteen \n")
        fmt.Println("   ... via the Gauss–Legendre algorithm ... \n")

        piAsBF := new(big.Float)
        piAsBF = big.NewFloat(Ricks_value) // pi is being cast to big from float64

            t := time.Now()
            elapsed := t.Sub(start)
            TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
        printResultStatsLong(piAsBF, 0, "Gauss–Legendre", iters, TotalRun, selection)
} 
// helper functions follow: 
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
} // End of Gauss_Legendre Set // case 37: // -- AMFGauss_LegendreB


// // -- AMFprintResultStatsLongA
// A HELPER FUNCTION ==== A HELPER FUNCTION ==== A HELPER FUNCTION ==== A HELPER FUNCTION ==== A HELPER FUNCTION ==== A HELPER FUNCTION ==== 
//------------------------------------------------------------------------------------------------------------------------------------------
// This next func is quite long as it handles all printing and logging of all results for all the sub-apps of this app 
func printResultStatsLong(sumBig *big.Float, precision int, useAlternateFile string, arg01 int, TotalRun string, selection int) {  // prslc2d==default, prslc2c==chud
    var piAs59766chars string 
    //  59,766 is the limit of the size of token Go can handle 
    piAs59766chars = "3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679821480865132823066470938446095505822317253594081284811174502841027019385211055596446229489549303819644288109756659334461284756482337867831652712019091456485669234603486104543266482133936072602491412737245870066063155881748815209209628292540917153643678925903600113305305488204665213841469519415116094330572703657595919530921861173819326117931051185480744623799627495673518857527248912279381830119491298336733624406566430860213949463952247371907021798609437027705392171762931767523846748184676694051320005681271452635608277857713427577896091736371787214684409012249534301465495853710507922796892589235420199561121290219608640344181598136297747713099605187072113499999983729780499510597317328160963185950244594553469083026425223082533446850352619311881710100031378387528865875332083814206171776691473035982534904287554687311595628638823537875937519577818577805321712268066130019278766111959092164201989380952572010654858632788659361533818279682303019520353018529689957736225994138912497217752834791315155748572424541506959508295331168617278558890750983817546374649393192550604009277016711390098488240128583616035637076601047101819429555961989467678374494482553797747268471040475346462080466842590694912933136770289891521047521620569660240580381501935112533824300355876402474964732639141992726042699227967823547816360093417216412199245863150302861829745557067498385054945885869269956909272107975093029553211653449872027559602364806654991198818347977535663698074265425278625518184175746728909777727938000816470600161452491921732172147723501414419735685481613611573525521334757418494684385233239073941433345477624168625189835694855620992192221842725502542568876717904946016534668049886272327917860857843838279679766814541009538837863609506800642251252051173929848960841284886269456042419652850222106611863067442786220391949450471237137869609563643719172874677646575739624138908658326459958133904780275900994657640789512694683983525957098258226205224894077267194782684826014769909026401363944374553050682034962524517493996514314298091906592509372216964615157098583874105978859597729754989301617539284681382686838689427741559918559252459539594310499725246808459872736446958486538367362226260991246080512438843904512441365497627807977156914359977001296160894416948685558484063534220722258284886481584560285060168427394522674676788952521385225499546667278239864565961163548862305774564980355936345681743241125150760694794510965960940252288797108931456691368672287489405601015033086179286809208747609178249385890097149096759852613655497818931297848216829989487226588048575640142704775551323796414515237462343645428584447952658678210511413547357395231134271661021359695362314429524849371871101457654035902799344037420073105785390621983874478084784896833214457138687519435064302184531910484810053706146806749192781911979399520614196634287544406437451237181921799983910159195618146751426912397489409071864942319615679452080951465502252316038819301420937621378559566389377870830390697920773467221825625996615014215030680384477345492026054146659252014974428507325186660021324340881907104863317346496514539057962685610055081066587969981635747363840525714591028970641401109712062804390397595156771577004203378699360072305587631763594218731251471205329281918261861258673215791984148488291644706095752706957220917567116722910981690915280173506712748583222871835209353965725121083579151369882091444210067510334671103141267111369908658516398315019701651511685171437657618351556508849099898599823873455283316355076479185358932261854896321329330898570642046752590709154814165498594616371802709819943099244889575712828905923233260972997120844335732654893823911932597463667305836041428138830320382490375898524374417029132765618093773444030707469211201913020330380197621101100449293215160842444859637669838952286847831235526582131449576857262433441893039686426243410773226978028073189154411010446823252716201052652272111660396665573092547110557853763466820653109896526918620564769312570586356620185581007293606598764861179104533488503461136576867532494416680396265797877185560845529654126654085306143444318586769751456614068007002378776591344017127494704205622305389945613140711270004078547332699390814546646458807972708266830634328587856983052358089330657574067954571637752542021149557615814002501262285941302164715509792592309907965473761255176567513575178296664547791745011299614890304639947132962107340437518957359614589019389713111790429782856475032031986915140287080859904801094121472213179476477726224142548545403321571853061422881375850430633217518297986622371721591607716692547487389866549494501146540628433663937900397692656721463853067360965712091807638327166416274888800786925602902284721040317211860820419000422966171196377921337575114959501566049631862947265473642523081770367515906735023507283540567040386743513622224771589150495309844489333096340878076932599397805419341447377441842631298608099888687413260472156951623965864573021631598193195167353812974167729478672422924654366800980676928238280689964004824354037014163149658979409243237896907069779422362508221688957383798623001593776471651228935786015881617557829735233446042815126272037343146531977774160319906655418763979293344195215413418994854447345673831624993419131814809277771038638773431772075456545322077709212019051660962804909263601975988281613323166636528619326686336062735676303544776280350450777235547105859548702790814356240145171806246436267945612753181340783303362542327839449753824372058353114771199260638133467768796959703098339130771098704085913374641442822772634659470474587847787201927715280731767907707157213444730605700733492436931138350493163128404251219256517980694113528013147013047816437885185290928545201165839341965621349143415956258658655705526904965209858033850722426482939728584783163057777560688876446248246857926039535277348030480290058760758251047470916439613626760449256274204208320856611906254543372131535958450687724602901618766795240616342522577195429162991930645537799140373404328752628889639958794757291746426357455254079091451357111369410911939325191076020825202618798531887705842972591677813149699009019211697173727847684726860849003377024242916513005005168323364350389517029893922334517220138128069650117844087451960121228599371623130171144484640903890644954440061986907548516026327505298349187407866808818338510228334508504860825039302133219715518430635455007668282949304137765527939751754613953984683393638304746119966538581538420568533862186725233402830871123282789212507712629463229563989898935821167456270102183564622013496715188190973038119800497340723961036854066431939509790190699639552453005450580685501956730229219139339185680344903982059551002263535361920419947455385938102343955449597783779023742161727111723643435439478221818528624085140066604433258885698670543154706965747458550332323342107301545940516553790686627333799585115625784322988273723198987571415957811196358330059408730681216028764962867446047746491599505497374256269010490377819868359381465741268049256487985561453723478673303904688383436346553794986419270563872931748723320837601123029911367938627089438799362016295154133714248928307220126901475466847653576164773794675200490757155527819653621323926406160136358155907422020203187277605277219005561484255518792530343513984425322341576233610642506390497500865627109535919465897514131034822769306247435363256916078154781811528436679570611086153315044521274739245449454236828860613408414863776700961207151249140430272538607648236341433462351897576645216413767969031495019108575984423919862916421939949072362346468441173940326591840443780513338945257423995082965912285085558215725031071257012668302402929525220118726767562204154205161841634847565169998116141010029960783869092916030288400269104140792886215078424516709087000699282120660418371806535567252532567532861291042487761825829765157959847035622262934860034158722980534989650226291748788202734209222245339856264766914905562842503912757710284027998066365825488926488025456610172967026640765590429099456815065265305371829412703369313785178609040708667114965583434347693385781711386455873678123014587687126603489139095620099393610310291616152881384379099042317473363948045759314931405297634757481193567091101377517210080315590248530906692037671922033229094334676851422144773793937517034436619910403375111735471918550464490263655128162288244625759163330391072253837421821408835086573917715096828874782656995995744906617583441375223970968340800535598491754173818839994469748676265516582765848358845314277568790029095170283529716344562129640435231176006651012412006597558512761785838292041974844236080071930457618932349229279650198751872127267507981255470958904556357921221033346697499235630254947802490114195212382815309114079073860251522742995818072471625916685451333123948049470791191532673430282441860414263639548000448002670496248201792896476697583183271314251702969234889627668440323260927524960357996469256504936818360900323809293459588970695365349406034021665443755890045632882250545255640564482465151875471196218443965825337543885690941130315095261793780029741207665147939425902989695946995565761218656196733786236256125216320862869222103274889218654364802296780705765615144632046927906821207388377814233562823608963208068222468012248261177185896381409183903673672220888321513755600372798394004152970028783076670944474560134556417254370906979396122571429894671543578468788614445812314593571984922528471605049221242470141214780573455105008019086996033027634787081081754501193071412233908663938339529425786905076431006383519834389341596131854347546495569781038293097164651438407007073604112373599843452251610507027056235266012764848308407611830130527932054274628654036036745328651057065874882256981579367897669742205750596834408697350201410206723585020072452256326513410559240190274216248439140359989535394590944070469120914093870012645600162374288021092764579310657922955249887275846101264836999892256959688159205600101655256375678566722796619885782794848855834397518744545512965634434803966420557982936804352202770984294232533022576341807039476994159791594530069752148293366555661567873640053666564165473217043903521329543529169414599041608753201868379370234888689479151071637852902345292440773659495630510074210871426134974595615138498713757047101787957310422969066670214498637464595280824369445789772330048764765241339075920434019634039114732023380715095222010682563427471646024335440051521266932493419673977041595683753555166730273900749729736354964533288869844061196496162773449518273695588220757355176651589855190986665393549481068873206859907540792342402300925900701731960362254756478940647548346647760411463233905651343306844953979070903023460461470961696886885014083470405460742958699138296682468185710318879065287036650832431974404771855678934823089431068287027228097362480939962706074726455399253994428081137369433887294063079261595995462624629707062594845569034711972996409089418059534393251236235508134949004364278527138315912568989295196427287573946914272534366941532361004537304881985517065941217352462589548730167600298865925786628561249665523533829428785425340483083307016537228563559152534784459818313411290019992059813522051173365856407826484942764411376393866924803118364453698589175442647399882284621844900877769776312795722672655562596282542765318300134070922334365779160128093179401718598599933849235495640057099558561134980252499066984233017350358044081168552653117099570899427328709258487894436460050410892266917835258707859512983441729535195378855345737426085902908176515578039059464087350612322611200937310804854852635722825768203416050484662775045003126200800799804925485346941469775164932709504934639382432227188515974054702148289711177792376122578873477188196825462981268685817050740272550263329044976277894423621674119186269439650671515779586756482399391760426017633870454990176143641204692182370764887834196896861181558158736062938603810171215855272668300823834046564758804051380801633638874216371406435495561868964112282140753302655100424104896783528588290243670904887118190909494533144218287661810310073547705498159680772009474696134360928614849417850171807793068108546900094458995279424398139213505586422196483491512639012803832001097738680662877923971801461343244572640097374257007359210031541508936793008169980536520276007277496745840028362405346037263416554259027601834840306811381855105979705664007509426087885735796037324514146786703688098806097164258497595138069309449401515422221943291302173912538355915031003330325111749156969174502714943315155885403922164097229101129035521815762823283182342548326111912800928252561902052630163911477247331485739107775874425387611746578671169414776421441111263583553871361011023267987756410246824032264834641766369806637857681349204530224081972785647198396308781543221166912246415911776732253264335686146186545222681268872684459684424161078540167681420808850280054143613146230821025941737562389942075713627516745731891894562835257044133543758575342698699472547031656613991999682628247270641336222178923903176085428943733935618891651250424404008952719837873864805847268954624388234375178852014395600571048119498842390606136957342315590796703461491434478863604103182350736502778590897578272731305048893989009923913503373250855982655867089242612429473670193907727130706869170926462548423240748550366080136046689511840093668609546325002145852930950000907151058236267293264537382104938724996699339424685516483261134146110680267446637334375340764294026682973865220935701626384648528514903629320199199688285171839536691345222444708045923966028171565515656661113598231122506289058549145097157553900243931535190902107119457300243880176615035270862602537881797519478061013715004489917210022201335013106016391541589578037117792775225978742891917915522417189585361680594741234193398420218745649256443462392531953135103311476394911995072858430658361935369329699289837914941939406085724863968836903265564364216644257607914710869984315733749648835292769328220762947282381537409961545598798259891093717126218283025848112389011968221429457667580718653806506487026133892822994972574530332838963818439447707794022843598834100358385423897354243956475556840952248445541392394100016207693636846776413017819659379971557468541946334893748439129742391433659360410035234377706588867781139498616478747140793263858738624732889645643598774667638479466504074111825658378878454858148962961273998413442726086061872455452360643153710112746809778704464094758280348769758948328241239292960582948619196670918958089833201210318430340128495116203534280144127617285830243559830032042024512072872535581195840149180969253395075778400067465526031446167050827682772223534191102634163157147406123850425845988419907611287258059113935689601431668283176323567325417073420817332230462987992804908514094790368878687894930546955703072619009502076433493359106024545086453628935456862958531315337183868265617862273637169757741830239860065914816164049449650117321313895747062088474802365371031150898427992754426853277974311395143574172219759799359685252285745263796289612691572357986620573408375766873884266405990993505000813375432454635967504844235284874701443545419576258473564216198134073468541117668831186544893776979566517279662326714810338643913751865946730024434500544995399742372328712494834706044063471606325830649829795510109541836235030309453097335834462839476304775645015008507578949548931393944899216125525597701436858943585877526379625597081677643800125436502371412783467926101995585224717220177723700417808419423948725406801556035998390548985723546745642390585850216719031395262944554391316631345308939062046784387785054239390524731362012947691874975191011472315289326772533918146607300089027768963114810902209724520759167297007850580717186381054967973100167870850694207092232908070383263453452038027860990556900134137182368370991949516489600755049341267876436746384902063964019766685592335654639138363185745698147196210841080961884605456039038455343729141446513474940784884423772175154334260306698831768331001133108690421939031080143784334151370924353013677631084913516156422698475074303297167469640666531527035325467112667522460551199581831963763707617991919203579582007595605302346267757943936307463056901080114942714100939136913810725813781357894005599500183542511841721360557275221035268037357265279224173736057511278872181908449006178013889710770822931002797665935838758909395688148560263224393726562472776037890814458837855019702843779362407825052704875816470324581290878395232453237896029841669225489649715606981192186584926770403956481278102179913217416305810554598801300484562997651121241536374515005635070127815926714241342103301566165356024733807843028655257222753049998837015348793008062601809623815161366903341111386538510919367393835229345888322550887064507539473952043968079067086806445096986548801682874343786126453815834280753061845485903798217994599681154419742536344399602902510015888272164745006820704193761584547123183460072629339550548239557137256840232268213012476794522644820910235647752723082081063518899152692889108455571126603965034397896278250016110153235160519655904211844949907789992007329476905868577878720982901352956613978884860509786085957017731298155314951681467176959760994210036183559138777817698458758104466283998806006162298486169353373865787735983361613384133853684211978938900185295691967804554482858483701170967212535338758621582310133103877668272115726949518179589754693992642197915523385766231676275475703546994148929041301863861194391962838870543677743224276809132365449485366768000001065262485473055861598999140170769838548318875014293890899506854530765116803337322265175662207526951791442252808165171667766727930354851542040238174608923283917032754257508676551178593950027933895920576682789677644531840404185540104351348389531201326378369283580827193783126549617459970567450718332065034556644034490453627560011250184335607361222765949278393706478426456763388188075656121689605041611390390639601620221536849410926053876887148379895599991120991646464411918568277004574243434021672276445589330127781586869525069499364610175685060167145354315814801054588605645501332037586454858403240298717093480910556211671546848477803944756979804263180991756422809873998766973237695737015808068229045992123661689025962730430679316531149401764737693873514093361833216142802149763399189835484875625298752423873077559555955465196394401821840998412489826236737714672260616336432964063357281070788758164043814850188411431885988276944901193212968271588841338694346828590066640806314077757725705630729400492940302420498416565479736705485580445865720227637840466823379852827105784319753541795011347273625774080213476826045022851579795797647467022840999561601569108903845824502679265942055503958792298185264800706837650418365620945554346135134152570065974881916341359556719649654032187271602648593049039787489589066127250794828276938953521753621850796297785146188432719223223810158744450528665238022532843891375273845892384422535472653098171578447834215822327020690287232330053862163479885094695472004795231120150432932266282727632177908840087861480221475376578105819702226309717495072127248479478169572961423658595782090830733233560348465318730293026659645013718375428897557971449924654038681799213893469244741985097334626793321072686870768062639919361965044099542167627840914669856925715074315740793805323925239477557441591845821562518192155233709607483329234921034514626437449805596103307994145347784574699992128599999399612281615219314888769388022281083001986016549416542616968586788372609587745676182507275992950893180521872924610867639958916145855058397274209809097817293239301067663868240401113040247007350857828724627134946368531815469690466968693925472519413992914652423857762550047485295476814795467007050347999588867695016124972282040303995463278830695976249361510102436555352230690612949388599015734661023712235478911292547696176005047974928060721268039226911027772261025441492215765045081206771735712027180242968106203776578837166909109418074487814049075517820385653909910477594141321543284406250301802757169650820964273484146957263978842560084531214065935809041271135920041975985136254796160632288736181367373244506079244117639975974619383584574915988097667447093006546342423460634237474666080431701260052055928493695941434081468529815053947178900451835755154125223590590687264878635752541911288877371766374860276606349603536794702692322971868327717393236192007774522126247518698334951510198642698878471719396649769070825217423365662725928440620430214113719922785269984698847702323823840055655517889087661360130477098438611687052310553149162517283732728676007248172987637569816335415074608838663640693470437206688651275688266149730788657015685016918647488541679154596507234287730699853713904300266530783987763850323818215535597323530686043010675760838908627049841888595138091030423595782495143988590113185835840667472370297149785084145853085781339156270760356390763947311455495832266945702494139831634332378975955680856836297253867913275055542524491943589128405045226953812179131914513500993846311774017971512283785460116035955402864405902496466930707769055481028850208085800878115773817191741776017330738554758006056014337743299012728677253043182519757916792969965041460706645712588834697979642931622965520168797300035646304579308840327480771811555330909887025505207680463034608658165394876951960044084820659673794731680864156456505300498816164905788311543454850526600698230931577765003780704661264706021457505793270962047825615247145918965223608396645624105195510522357239739512881816405978591427914816542632892004281609136937773722299983327082082969955737727375667615527113922588055201898876201141680054687365580633471603734291703907986396522961312801782679717289822936070288069087768660593252746378405397691848082041021944719713869256084162451123980620113184541244782050110798760717155683154078865439041210873032402010685341947230476666721749869868547076781205124736792479193150856444775379853799732234456122785843296846647513336573692387201464723679427870042503255589926884349592876124007558756946413705625140011797133166207153715436006876477318675587148783989081074295309410605969443158477539700943988394914432353668539209946879645066533985738887866147629443414010498889931600512076781035886116602029611936396821349607501116498327856353161451684576956871090029997698412632665023477167286573785790857466460772283415403114415294188047825438761770790430001566986776795760909966936075594965152736349811896413043311662774712338817406037317439705406703109676765748695358789670031925866259410510533584384656023391796749267844763708474978333655579007384191473198862713525954625181604342253729962863267496824058060296421146386436864224724887283434170441573482481833301640566959668866769563491416328426414974533349999480002669987588815935073578151958899005395120853510357261373640343675347141048360175464883004078464167452167371904831096767113443494819262681110739948250607394950735031690197318521195526356325843390998224986240670310768318446607291248747540316179699411397387765899868554170318847788675929026070043212666179192235209382278788809886335991160819235355570464634911320859189796132791319756490976000139962344455350143464268604644958624769094347048293294140411146540923988344435159133201077394411184074107684981066347241048239358274019449356651610884631256785297769734684303061462418035852933159734583038455410337010916767763742762102137013548544509263071901147318485749233181672072137279355679528443925481560913728128406333039373562420016045664557414588166052166608738748047243391212955877763906969037078828527753894052460758496231574369171131761347838827194168606625721036851321566478001476752310393578606896111259960281839309548709059073861351914591819510297327875571049729011487171897180046961697770017913919613791417162707018958469214343696762927459109940060084983568425201915593703701011049747339493877885989417433031785348707603221982970579751191440510994235883034546353492349826883624043327267415540301619505680654180939409982020609994140216890900708213307230896621197755306659188141191577836272927461561857103721724710095214236964830864102592887457999322374955191221951903424452307535133806856807354464995127203174487195403976107308060269906258076020292731455252078079914184290638844373499681458273372072663917670201183004648190002413083508846584152148991276106513741539435657211390328574918769094413702090517031487773461652879848235338297260136110984514841823808120540996125274580881099486972216128524897425555516076371675054896173016809613803811914361143992106380050832140987604599309324851025168294467260666138151745712559754953580239983146982203613380828499356705575524712902745397762140493182014658008021566536067765508783804304134310591804606800834591136640834887408005741272586704792258319127415739080914383138456424150940849133918096840251163991936853225557338966953749026620923261318855891580832455571948453875628786128859004106006073746501402627824027346962528217174941582331749239683530136178653673760642166778137739951006589528877427662636841830680190804609849809469763667335662282915132352788806157768278159588669180238940333076441912403412022316368577860357276941541778826435238131905028087018575047046312933353757285386605888904583111450773942935201994321971171642235005644042979892081594307167019857469273848653833436145794634175922573898588001698014757420542995801242958105456510831046297282937584161162532562516572498078492099897990620035936509934721582965174135798491047111660791587436986541222348341887722929446335178653856731962559852026072947674072616767145573649812105677716893484917660771705277187601199908144113058645577910525684304811440261938402322470939249802933550731845890355397133088446174107959162511714864874468611247605428673436709046678468670274091881014249711149657817724279347070216688295610877794405048437528443375108828264771978540006509704033021862556147332117771174413350281608840351781452541964320309576018694649088681545285621346988355444560249556668436602922195124830910605377201980218310103270417838665447181260397190688462370857518080035327047185659499476124248110999288679158969049563947624608424065930948621507690314987020673533848349550836366017848771060809804269247132410009464014373603265645184566792456669551001502298330798496079949882497061723674493612262229617908143114146609412341593593095854079139087208322733549572080757165171876599449856937956238755516175754380917805280294642004472153962807463602113294255916002570735628126387331060058910652457080244749375431841494014821199962764531068006631183823761639663180931444671298615527598201451410275600689297502463040173514891945763607893528555053173314164570504996443890936308438744847839616840518452732884032345202470568516465716477139323775517294795126132398229602394548579754586517458787713318138752959809412174227300352296508089177705068259248822322154938048371454781647213976820963320508305647920482085920475499857320388876391601995240918938945576768749730856955958010659526503036266159750662225084067428898265907510637563569968211510949669744580547288693631020367823250182323708459790111548472087618212477813266330412076216587312970811230758159821248639807212407868878114501655825136178903070860870198975889807456643955157415363193191981070575336633738038272152798849350397480015890519420879711308051233933221903466249917169150948541401871060354603794643379005890957721180804465743962806186717861017156740967662080295766577051291209907944304632892947306159510430902221439371849560634056189342513057268291465783293340524635028929175470872564842600349629611654138230077313327298305001602567240141851520418907011542885799208121984493156999059182011819733500126187728036812481995877070207532406361259313438595542547781961142935163561223496661522614735399674051584998603552953329245752388810136202347624669055816438967863097627365504724348643071218494373485300606387644566272186661701238127715621379746149861328744117714552444708997144522885662942440230184791205478498574521634696448973892062401943518310088283480249249085403077863875165911302873958787098100772718271874529013972836614842142871705531796543076504534324600536361472618180969976933486264077435199928686323835088756683595097265574815431940195576850437248001020413749831872259677387154958399718444907279141965845930083942637020875635398216962055324803212267498911402678528599673405242031091797899905718821949391320753431707980023736590985375520238911643467185582906853711897952626234492483392496342449714656846591248918556629589329909035239233333647435203707701010843880032907598342170185542283861617210417603011645918780539367447472059985023582891833692922337323999480437108419659473162654825748099482509991833006976569367159689364493348864744213500840700660883597235039532340179582557036016936990988671132109798897070517280755855191269930673099250704070245568507786790694766126298082251633136399521170984528092630375922426742575599892892783704744452189363203489415521044597261883800300677617931381399162058062701651024458869247649246891924612125310275731390840470007143561362316992371694848132554200914530410371354532966206392105479824392125172540132314902740585892063217589494345489068463993137570910346332714153162232805522972979538018801628590735729554162788676498274186164218789885741071649069191851162815285486794173638906653885764229158342500673612453849160674137340173572779956341043326883569507814931378007362354180070619180267328551191942676091221035987469241172837493126163395001239599240508454375698507957046222664619000103500490183034153545842833764378111988556318777792537201166718539541835984438305203762819440761594106820716970302285152250573126093046898423433152732131361216582808075212631547730604423774753505952287174402666389148817173086436111389069420279088143119448799417154042103412190847094080254023932942945493878640230512927119097513536000921971105412096683111516328705423028470073120658032626417116165957613272351566662536672718998534199895236884830999302757419916463841427077988708874229277053891227172486322028898425125287217826030500994510824783572905691988555467886079462805371227042466543192145281760741482403827835829719301017888345674167811398954750448339314689630763396657226727043393216745421824557062524797219978668542798977992339579057581890622525473582205236424850783407110144980478726691990186438822932305382318559732869780922253529591017341407334884761005564018242392192695062083183814546983923664613639891012102177095976704908305081854704194664371312299692358895384930136356576186106062228705599423371631021278457446463989738188566746260879482018647487672727222062676465338099801966883680994159075776852639865146253336312450536402610569605513183813174261184420189088853196356986962795036738424313011331753305329802016688817481342988681585577810343231753064784983210629718425184385534427620128234570716988530518326179641178579608888150329602290705614476220915094739035946646916235396809201394578175891088931992112260073928149169481615273842736264298098234063200244024495894456129167049508235812487391799648641133480324757775219708932772262349486015046652681439877051615317026696929704928316285504212898146706195331970269507214378230476875280287354126166391708245925170010714180854800636923259462019002278087409859771921805158532147392653251559035410209284665925299914353791825314545290598415817637058927906909896911164381187809435371521332261443625314490127454772695739393481546916311624928873574718824071503995009446731954316193855485207665738825139639163576723151005556037263394867208207808653734942440115799667507360711159351331959197120948964717553024531364770942094635696982226673775209945168450643623824211853534887989395673187806606107885440005508276570305587448541805778891719207881423351138662929667179643468760077047999537883387870348718021842437342112273940255717690819603092018240188427057046092622564178375265263358324240661253311529423457965569502506810018310900411245379015332966156970522379210325706937051090830789479999004999395322153622748476603613677697978567386584670936679588583788795625946464891376652199588286933801836011932368578558558195556042156250883650203322024513762158204618106705195330653060606501054887167245377942831338871631395596905832083416898476065607118347136218123246227258841990286142087284956879639325464285343075301105285713829643709990356948885285190402956047346131138263878897551788560424998748316382804046848618938189590542039889872650697620201995548412650005394428203930127481638158530396439925470201672759328574366661644110962566337305409219519675148328734808957477775278344221091073111351828046036347198185655572957144747682552857863349342858423118749440003229690697758315903858039353521358860079600342097547392296733310649395601812237812854584317605561733861126734780745850676063048229409653041118306671081893031108871728167519579675347188537229309616143204006381322465841111157758358581135018569047815368938137718472814751998350504781297718599084707621974605887423256995828892535041937958260616211842368768511418316068315867994601652057740529423053601780313357263267054790338401257305912339601880137825421927094767337191987287385248057421248921183470876629667207272325650565129333126059505777727542471241648312832982072361750574673870128209575544305968395555686861188397135522084452852640081252027665557677495969626612604565245684086139238265768583384698499778726706555191854468698469478495734622606294219624557085371272776523098955450193037732166649182578154677292005212667143463209637891852323215018976126034373684067194193037746880999296877582441047878123266253181845960453853543839114496775312864260925211537673258866722604042523491087026958099647595805794663973419064010036361904042033113579336542426303561457009011244800890020801478056603710154122328891465722393145076071670643556827437743965789067972687438473076346451677562103098604092717090951280863090297385044527182892749689212106670081648583395537735919136950153162018908887484210798706899114804669270650940762046502772528650728905328548561433160812693005693785417861096969202538865034577183176686885923681488475276498468821949739729707737187188400414323127636504814531122850990020742409255859252926103021067368154347015252348786351643976235860419194129697690405264832347009911154242601273438022089331096686367898694977994001260164227609260823493041180643829138347354679725399262338791582998486459271734059225620749105308531537182911681637219395188700957788181586850464507699343940987433514431626330317247747486897918209239480833143970840673084079589358108966564775859905563769525232653614424780230826811831037735887089240613031336477371011628214614661679404090518615260360092521947218890918107335871964142144478654899528582343947050079830388538860831035719306002771194558021911942899922722353458707566246926177663178855144350218287026685610665003531050216318206017609217984684936863161293727951873078972637353717150256378733579771808184878458866504335824377004147710414934927438457587107159731559439426412570270965125108115548247939403597681188117282472158250109496096625393395380922195591918188552678062149923172763163218339896938075616855911752998450132067129392404144593862398809381240452191484831646210147389182510109096773869066404158973610476436500068077105656718486281496371118832192445663945814491486165500495676982690308911185687986929470513524816091743243015383684707292898982846022237301452655679898627767968091469798378268764311598832109043715611299766521539635464420869197567370005738764978437686287681792497469438427465256316323005551304174227341646455127812784577772457520386543754282825671412885834544435132562054464241011037955464190581168623059644769587054072141985212106734332410756767575818456990693046047522770167005684543969234041711089888993416350585157887353430815520811772071880379104046983069578685473937656433631979786803671873079693924236321448450354776315670255390065423117920153464977929066241508328858395290542637687668968805033317227800185885069736232403894700471897619347344308437443759925034178807972235859134245813144049847701732361694719765715353197754997162785663119046912609182591249890367654176979903623755286526375733763526969344354400473067198868901968147428767790866979688522501636949856730217523132529265375896415171479559538784278499866456302878831962099830494519874396369070682762657485810439112232618794059941554063270131989895703761105323606298674803779153767511583043208498720920280929752649812569163425000522908872646925284666104665392171482080130502298052637836426959733707053922789153510568883938113249757071331029504430346715989448786847116438328050692507766274500122003526203709466023414648998390252588830148678162196775194583167718762757200505439794412459900771152051546199305098386982542846407255540927403132571632640792934183342147090412542533523248021932277075355546795871638358750181593387174236061551171013123525633485820365146141870049205704372018261733194715700867578539336078622739558185797587258744102542077105475361294047460100094095444959662881486915903899071865980563617137692227290764197755177720104276496949611056220592502420217704269622154958726453989227697660310524980855759471631075870133208861463266412591148633881220284440694169488261529577625325019870359870674380469821942056381255833436421949232275937221289056420943082352544084110864545369404969271494003319782861318186188811118408257865928757426384450059944229568586460481033015388911499486935436030221810943466764000022362550573631294626296096198760564259963946138692330837196265954739234624134597795748524647837980795693198650815977675350553918991151335252298736112779182748542008689539658359421963331502869561192012298889887006079992795411188269023078913107603617634779489432032102773359416908650071932804017163840644987871753756781185321328408216571107549528294974936214608215583205687232185574065161096274874375098092230211609982633033915469494644491004515280925089745074896760324090768983652940657920198315265410658136823791984090645712468948470209357761193139980246813405200394781949866202624008902150166163813538381515037735022966074627952910384068685569070157516624192987244482719429331004854824454580718897633003232525821581280327467962002814762431828622171054352898348208273451680186131719593324711074662228508710666117703465352839577625997744672185715816126411143271794347885990892808486694914139097716736900277758502686646540565950394867841110790116104008572744562938425494167594605487117235946429105850909950214958793112196135908315882620682332156153086833730838173279328196983875087083483880463884784418840031847126974543709373298362402875197920802321878744882872843727378017827008058782410749357514889978911739746129320351081432703251409030487462262942344327571260086642508333187688650756429271605525289544921537651751492196367181049435317858383453865255656640657251363575064353236508936790431702597878177190314867963840828810209461490079715137717099061954969640070867667102330048672631475510537231757114322317411411680622864206388906210192355223546711662137499693269321737043105987225039456574924616978260970253359475020913836673772894438696400028110344026084712899000746807764844088711341352503367877316797709372778682166117865344231732264637847697875144332095340001650692130546476890985050203015044880834261845208730530973189492916425322933612431514306578264070283898409841602950309241897120971601649265613413433422298827909921786042679812457285345801338260995877178113102167340256562744007296834066198480676615805021691833723680399027931606420436812079900316264449146190219458229690992122788553948783538305646864881655562294315673128274390826450611628942803501661336697824051770155219626522725455850738640585299830379180350432876703809252167907571204061237596327685674845079151147313440001832570344920909712435809447900462494313455028900680648704293534037436032625820535790118395649089354345101342969617545249573960621490288728932792520696535386396443225388327522499605986974759882329916263545973324445163755334377492928990581175786355555626937426910947117002165411718219750519831787137106051063795558588905568852887989084750915764639074693619881507814685262133252473837651192990156109189777922008705793396463827490680698769168197492365624226087154176100430608904377976678519661891404144925270480881971498801542057787006521594009289777601330756847966992955433656139847738060394368895887646054983871478968482805384701730871117761159663505039979343869339119789887109156541709133082607647406305711411098839388095481437828474528838368079418884342666222070438722887413947801017721392281911992365405516395893474263953824829609036900288359327745855060801317988407162446563997948275783650195514221551339281978226984278638391679715091262410548725700924070045488485692950448110738087996547481568913935380943474556972128919827177020766613602489581468119133614121258783895577357194986317210844398901423948496659251731388171602663261931065366535041473070804414939169363262373767777095850313255990095762731957308648042467701212327020533742667053142448208168130306397378736642483672539837487690980602182785786216512738563513290148903509883270617258932575363993979055729175160097615459044771692265806315111028038436017374742152476085152099016158582312571590733421736576267142390478279587281505095633092802668458937649649770232973641319060982740633531089792464242134583740901169391964250459128813403498810635400887596820054408364386516617880557608956896727531538081942077332597917278437625661184319891025007491829086475149794003160703845549465385946027452447466812314687943441610993338908992638411847425257044572517459325738989565185716575961481266020310797628254165590506042479114016957900338356574869252800743025623419498286467914476322774005529460903940177536335655471931000175430047504719144899841040015867946179241610016454716551337074073950260442769538553834397550548871099785205401175169747581344926079433689543783221172450687344231989878844128542064742809735625807066983106979935260693392135685881391214807354728463227784908087002467776303605551232386656295178853719673034634701222939581606792509153217489030840886516061119011498443412350124646928028805996134283511884715449771278473361766285062169778717743824362565711779450064477718370221999106695021656757644044997940765037999954845002710665987813603802314126836905783190460792765297277694043613023051787080546511542469395265127101052927070306673024447125973939950514628404767431363739978259184541176413327906460636584152927019030276017339474866960348694976541752429306040727005059039503148522921392575594845078867977925253931765156416197168443524369794447355964260633391055126826061595726217036698506473281266724521989060549880280782881429796336696744124805982192146339565745722102298677599746738126069367069134081559412016115960190237753525556300606247983261249881288192937343476862689219239777833910733106588256813777172328315329082525092733047850724977139448333892552081175608452966590553940965568541706001179857293813998258319293679100391844099286575605993598910002969864460974714718470101531283762631146774209145574041815908800064943237855839308530828305476076799524357391631221886057549673832243195650655460852881201902363644712703748634421727257879503428486312944916318475347531435041392096108796057730987201352484075057637199253650470908582513936863463863368042891767107602111159828875539940120076013947033661793715396306139863655492213741597905119083588290097656647300733879314678913181465109316761575821351424860442292445304113160652700974330088499034675405518640677342603583409608605533747362760935658853109760994238347382222087292464497684560579562516765574088410321731345627735856052358236389532038534024842273371639123973215995440828421666636023296545694703577184873442034227706653837387506169212768015766181095420097708363604361110592409117889540338021426523948929686439808926114635414571535194342850721353453018315875628275733898268898523557799295727645229391567477566676051087887648453493636068278050564622813598885879259940946446041705204470046315137975431737187756039815962647501410906658866162180038266989961965580587208639721176995219466789857011798332440601811575658074284182910615193917630059194314434605154047710570054339000182453117733718955857603607182860506356479979004139761808955363669603162193113250223851791672055180659263518036251214575926238369348222665895576994660491938112486609099798128571823494006615552196112207203092277646200999315244273589488710576623894693889446495093960330454340842102462401048723328750081749179875543879387381439894238011762700837196053094383940063756116458560943129517597713935396074322792489221267045808183313764165818269562105872892447740035947009268662659651422050630078592002488291860839743732353849083964326147000532423540647042089499210250404726781059083644007466380020870126664209457181702946752278540074508552377720890581683918446592829417018288233014971554235235911774818628592967605048203864343108779562892925405638946621948268711042828163893975711757786915430165058602965217459581988878680408110328432739867198621306205559855266036405046282152306154594474489908839081999738747452969810776201487134000122535522246695409315213115337915798026979555710508507473874750758068765376445782524432638046143042889235934852961058269382103498000405248407084403561167817170512813378805705643450616119330424440798260377951198548694559152051960093041271007277849301555038895360338261929343797081874320949914159593396368110627557295278004254863060054523839151068998913578820019411786535682149118528207852130125518518493711503422159542244511900207393539627400208110465530207932867254740543652717595893500716336076321614725815407642053020045340183572338292661915308354095120226329165054426123619197051613839357326693760156914429944943744856809775696303129588719161129294681884936338647392747601226964158848900965717086160598147204467428664208765334799858222090619802173211614230419477754990738738567941189824660913091691772274207233367635032678340586301930193242996397204445179288122854478211953530898910125342975524727635730226281382091807439748671453590778633530160821559911314144205091447293535022230817193663509346865858656314855575862447818620108711889760652969899269328178705576435143382060141077329261063431525337182243385263520217735440715281898137698755157574546939727150488469793619500477720970561793913828989845327426227288647108883270173723258818244658436249580592560338105215606206155713299156084892064340303395262263451454283678698288074251422567451806184149564686111635404971897682154227722479474033571527436819409892050113653400123846714296551867344153741615042563256713430247655125219218035780169240326699541746087592409207004669340396510178134857835694440760470232540755557764728450751826890418293966113310160131119077398632462778219023650660374041606724962490137433217246454097412995570529142438208076098364823465973886691349919784013108015581343979194852830436739012482082444814128095443773898320059864909159505322857914576884962578665885999179867520554558099004556461178755249370124553217170194282884617402736649978475508294228020232901221630102309772151569446427909802190826689868834263071609207914085197695235553488657743425277531197247430873043619511396119080030255878387644206085044730631299277888942729189727169890575925244679660189707482960949190648764693702750773866432391919042254290235318923377293166736086996228032557185308919284403805071030064776847863243191000223929785255372375566213644740096760539439838235764606992465260089090624105904215453927904411529580345334500256244101006359530039598864466169595626351878060688513723462707997327233134693971456285542615467650632465676620279245208581347717608521691340946520307673391841147504140168924121319826881568664561485380287539331160232292555618941042995335640095786495340935115266454024418775949316930560448686420862757201172319526405023099774567647838488973464317215980626787671838005247696884084989185086149003432403476742686245952395890358582135006450998178244636087317754378859677672919526111213859194725451400301180503437875277664402762618941017576872680428176623860680477885242887430259145247073950546525135339459598789619778911041890292943818567205070964606263541732944649576612651953495701860015412623962286413897796733329070567376962156498184506842263690367849555970026079867996261019039331263768556968767029295371162528005543100786408728939225714512481135778627664902425161990277471090335933309304948380597856628844787441469841499067123764789582263294904679812089984857163571087831191848630254501620929805829208334813638405421720056121989353669371336733392464416125223196943471206417375491216357008573694397305979709719726666642267431117762176403068681310351899112271339724036887000996862922546465006385288620393800504778276912835603372548255793912985251506829969107754257647488325341412132800626717094009098223529657957997803018282428490221470748111124018607613415150387569830918652780658896682362523937845272634530420418802508442363190383318384550522367992357752929106925043261446950109861088899914658551881873582528164302520939285258077969737620845637482114433988162710031703151334402309526351929588680690821355853680161000213740851154484912685841268695899174149133820578492800698255195740201818105641297250836070356851055331787840829000041552511865779453963317538532092149720526607831260281961164858098684587525129997404092797683176639914655386108937587952214971731728131517932904431121815871023518740757222100123768721944747209349312324107065080618562372526732540733324875754482967573450019321902199119960797989373383673242576103938985349278777473980508080015544764061053522202325409443567718794565430406735896491017610775948364540823486130254718476485189575836674399791508512858020607820554462991723202028222914886959399729974297471155371858924238493855858595407438104882624648788053304271463011941589896328792678327322456103852197011130466587100500083285177311776489735230926661234588873102883515626446023671996644554727608310118788389151149340939344750073025855814756190881398752357812331342279866503522725367171230756861045004548970360079569827626392344107146584895780241408158405229536937499710665594894459246286619963556350652623405339439142111271810691052290024657423604130093691889255865784668461215679554256605416005071276641766056874274200329577160643448606201239821698271723197826816628249938714995449137302051843669076723577400053932662622760323659751718925901801104290384274185507894887438832703063283279963007200698012244365116394086922220745320244624121155804354542064215121585056896157356414313068883443185280853975927734433655384188340303517822946253702015782157373265523185763554098954033236382319219892171177449469403678296185920803403867575834111518824177439145077366384071880489358256868542011645031357633355509440319236720348651010561049872726472131986543435450409131859513145181276437310438972507004981987052176272494065214619959232142314439776546708351714749367986186552791715824080651063799500184295938799158350171580759883784962257398512129810326379376218322456594236685376799113140108043139732335449090824910499143325843298821033984698141715756010829706583065211347076803680695322971990599904451209087275776225351040902392888779424630483280319132710495478599180196967835321464441189260631526618167443193550817081875477050802654025294109218264858213857526688155584113198560022135158887210365696087515063187533002942118682221893775546027227291290504292259787710667873840000616772154638441292371193521828499824350920891801685572798156421858191197490985730570332667646460728757430565372602768982373259745084479649545648030771598153955827779139373601717422996027353102768719449444917939785144631597314435351850491413941557329382048542123508173912549749819308714396615132942045919380106231421774199184060180347949887691051557905554806953878540066453375981862846419905220452803306263695626490910827627115903856995051246529996062855443838330327638599800792922846659503551211245284087516229060262011857775313747949362055496401073001348853150735487353905602908933526400713274732621960311773433943673385759124508149335736911664541281788171454023054750667136518258284898099512139193995633241336556777098003081910272040997148687418134667006094051021462690280449159646545330107754695413088714165312544813061192407821188690056027781824235022696189344352547633573536485619363254417756613981703930632872166905722259745209192917262199844409646158269456380239502837121686446561785235565164127712826918688615572716201474934052276946595712198314943381622114006936307430444173284786101777743837977037231795255434107223445512555589998646183876764903972461167959018100035098928641204195163551108763204267612979826529425882951141275841262732790798807559751851576841264742209479721843309352972665210015662514552994745127631550917636730259462132930190402837954246323258550301096706922720227074863419005438302650681214142135057154175057508639907673946335146209082888934938376439399256900604067311422093312195936202982972351163259386772241477911629572780752395056251581603133359382311500518626890530658368129988108663263271980611271548858798093487912913707498230575929091862939195014721197586067270092547718025750337730799397134539532646195269996596385654917590458333585799102012713204583903200853878881633637685182083727885131175227769609787962142372162545214591281831798216044111311671406914827170981015457781939202311563871950805024679725792497605772625913328559726371211201905720771409148645074094926718035815157571514050397610963846755569298970383547314100223802583468767350129775413279532060971154506484212185936490997917766874774481882870632315515865032898164228288232746866106592732197907162384642153489852476216789050260998045266483929542357287343977680495774091449538391575565485459058976495198513801007958010783759945775299196700547602252552034453988712538780171960718164078124847847257912407824544361682345239570689514272269750431873633263011103053423335821609333191218806608268341428910415173247216053355849993224548730778822905252324234861531520976938461042582849714963475341837562003014915703279685301868631572488401526639835689563634657435321783493199825542117308467745297085839507616458229630324424328237737450517028560698067889521768198156710781633405266759539424926280756968326107495323390536223090807081455919837355377748742029039018142937311529334644468151212945097596534306284215319445727118614900017650558177095302468875263250119705209476159416768727784472000192789137251841622857783792284439084301181121496366424659033634194540657183544771912446621259392656620306888520055599121235363718226922531781458792593750441448933981608657900876165024635197045828895481793756681046474614105142498870252139936870509372305447734112641354892806841059107716677821238332810262185587751312721179344448201440425745083063944738363793906283008973306241380614589414227694747931665717623182472168350678076487573420491557628217583972975134478990696589532548940335615613167403276472469212505759116251529654568544633498114317670257295661844775487469378464233737238981920662048511894378868224807279352022501796545343757274163910791972952950812942922205347717304184477915673991738418311710362524395716152714669005814700002633010452643547865903290733205468338872078735444762647925297690170912007874183736735087713376977683496344252419949951388315074877537433849458259765560996555954318040920178497184685497370696212088524377013853757681416632722412634423982152941645378000492507262765150789085071265997036708726692764308377229685985169122305037462744310852934305273078865283977335246017463527703205938179125396915621063637625882937571373840754406468964783100704580613446731271591194608435935825987782835266531151065041623295329047772174083559349723758552138048305090009646676088301540612824308740645594431853413755220166305812111033453120745086824339432159043594430312431227471385842030390106070940315235556172767994160020393975099897629335325855575624808996691829864222677502360193257974726742578211119734709402357457222271212526852384295874273501563660093188045493338989741571490544182559738080871565281430102670460284316819230392535297795765862414392701549740879273131051636119137577008929564823323648298263024607975875767745377160102490804624301856524161756655600160859121534556267602192689982855377872583145144082654583484409478463178777374794653580169960779405568701192328608041130904629350871827125934668712766694873899824598527786499569165464029458935064964335809824765965165142090986755203808309203230487342703468288751604071546653834619611223013759451579252696743642531927390036038608236450762698827497618723575476762889950752114804852527950845033958570838130476937881321123674281319487950228066320170022460331989671970649163741175854851878484012054844672588851401562725019821719066960812627785485964818369621410721714214986361918774754509650308957099470934337856981674465828267911940611956037845397855839240761276344105766751024307559814552786167815949657062559755074306521085301597908073343736079432866757890533483669555486803913433720156498834220893399971641479746938696905480089193067138057171505857307148815649920714086758259602876056459782423770242469805328056632787041926768467116266879463486950464507420219373945259262668613552940624781361206202636498199999498405143868285258956342264328707663299304891723400725471764188685351372332667877921738347541480022803392997357936152412755829569276837231234798989446274330454566790062032420516396282588443085438307201495672106460533238537203143242112607424485845094580494081820927639140008540422023556260218564348994145439950410980591817948882628052066441086319001688568155169229486203010738897181007709290590480749092427141018933542818429995988169660993836961644381528877214085268088757488293258735809905670755817017949161906114001908553744882726200936685604475596557476485674008177381703307380305476973609786543859382187220583902344443508867499866506040645874346005331827436296177862518081893144363251205107094690813586440519229512932450078833398788429339342435126343365204385812912834345297308652909783300671261798130316794385535726296998740359570458452230856390098913179475948752126397078375944861139451960286751210561638976008880092746115860800207803341591451797073036835196977766076373785333012024120112046988609209339085365773222392412449051532780950955866459477634482269986074813297302630975028812103517723124465095349653693090018637764094094349837313251321862080214809922685502948454661814715557444709669530177690434272031892770604717784527939160472281534379803539679861424370956683221491465438014593829277393396032754048009552231816667380357183932757077142046723838624617803976292377131209580789363841447929802588065522129262093623930637313496640186619510811583471173312025805866727639992763579078063818813069156366274125431259589936119647626101405563503399523140323113819656236327198961837254845333702062563464223952766943568376761368711962921818754576081617053031590728828700712313666308722754918661395773730546065997437810987649802414011242142773668082751390959313404155826266789510846776118665957660165998178089414985754976284387856100263796543178313634025135814161151902096499133548733131115022700681930135929595971640197196053625033558479980963488718039111612813595968565478868325856437896173159762002419621552896297904819822199462269487137462444729093456470028537694958859591606789282491054412515996300781368367490209374915732896270028656829344431342347351239298259166739503425995868970697267332582735903121288746660451461487850346142827765991608090398652575717263081833494441820193533385071292345774375579344062178711330063106003324053991693682603746176638565758877580201229366353270267100681261825172914608202541892885935244491070138206211553827793565296914576502048643282865557934707209634807372692141186895467322767751335690190153723669036865389161291688887876407525493494249733427181178892759931596719354758988097924525262363659036320070854440784544797348291802082044926670634420437555325050527522833778887040804033531923407685630109347772125639"
    stringOfSum := sumBig.Text('f', 999999) // create a string version of a big result, up to a million digits can be handled, if not verified, here-with

    posInPi := 0 // to be the incremented offset : piChar = piAs59766chars[posInPi]
    var piChar byte  // one byte (character) of pi as string, e.g. piChar = piAs59766chars[posInPi]
    var copyOfLastPosition int // an external (to the loop) copy of positionInString
    var stringVerOfCorrectDigits = []string{}  
    for positionInString, charAtRangePos := range stringOfSum {
        piChar = piAs59766chars[posInPi] 
            if charAtRangePos == rune(piChar) {
                stringVerOfCorrectDigits = append(stringVerOfCorrectDigits, string(charAtRangePos))
                copyOfLastPosition = positionInString // save an external copy, of the last position found to have matched pi, as an int 
            } else {
                break // to print result and info below 
            }
        posInPi++
    }

    if copyOfLastPosition > 55000 { // if length of pi is > 55,000 digits we have something really big
        // print (log) to a special file
        fmt.Printf("\n\n\nWe have been tasked with making a lot of pie and it was sooo big it needed its own file ...\n")
        fmt.Printf("\n\n  After allowing this process to finish (you may have to continue prodding this thing along for a while) ... \n")
        fmt.Println("... Go have a look in /.big_pie_is_in_here.txt to find all the digits of π you had requested. \n\n")

            fileHandleBig, err1prslc2c := os.OpenFile("big_pie_is_in_here.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                check(err1prslc2c)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                defer fileHandleBig.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
            _, err2prslc2c := fmt.Fprintf(fileHandleBig, "\nThese are the %d verified digits we have calculated per case: %d  :: \n", copyOfLastPosition, selection)
                check(err2prslc2c)
            for _, oneChar := range stringVerOfCorrectDigits {
                //fmt.Print(oneChar) // to the console // the whole point of using an alternate file is to not clutter up the console or the default file 
    // *************************************** this is the one and only logging loop ******************************************************************************        
                    _ , err8prslc2c := fmt.Fprint(fileHandleBig, oneChar)  // to a file
                    check(err8prslc2c)
             }
            _ , err9prslc2c := fmt.Fprintf(fileHandleBig, "\n...the preceeding was logged one char at a time by case: %d \n", selection)
                check(err9prslc2c)  
            fileHandleBig.Close() 
    } else {
        if useAlternateFile == "BBPfConcurent" {
            fmt.Printf("\n... we have actually calculated pi correctly to %d digits after %d loops and using %d workers\n", copyOfLastPosition, precision, arg01)
        } else {
            if usingBigFloats {
                fmt.Printf("\n... we have actually calculated pi correctly to %d digits using precision of %d \n", copyOfLastPosition, precision)
            } else {
                fmt.Printf("\n... we have actually calculated pi correctly to %d digits \n", copyOfLastPosition)
            }
        }   
        /*
        if useAlternateFile == "BBPF" {
            fmt.Printf("Those %d digits can be found in ", copyOfLastPosition)
            fmt.Printf("dataLog-From_BBPF_Method_lengthy_prints.txt\n")
        }
        */


    // } else { continues below: (in other words, the following if-else conditions are only checked if length of pi was < 55,000 digits)
        if useAlternateFile == "chudDid800orMoreLoops" {
            fileHandleChud, err1prslc2c := os.OpenFile("dataLog-From_Chudnovsky_Method_lengthy_prints.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                check(err1prslc2c)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                defer fileHandleChud.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
            _, err2prslc2c := fmt.Fprintf(fileHandleChud, "\n\nThese are the %d verified digits we have calculated via Chudnovsky: \n", copyOfLastPosition)
                check(err2prslc2c)
                    for _, oneChar := range stringVerOfCorrectDigits {
                        //fmt.Print(oneChar) // to the console // the whole point of using this alternate file is to not clutter up the console or the default file 
                            _ , err8prslc2c := fmt.Fprint(fileHandleChud, oneChar)  // to a file
                                check(err8prslc2c)
                    }

            fileHandleChud.Close() 
        } else if useAlternateFile == "BBPF" {
            fileHandleBBPF, err1prslc2c := os.OpenFile("dataLog-From_BBPF_Method_lengthy_prints.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                check(err1prslc2c)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                defer fileHandleBBPF.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
            _, err2prslc2c := fmt.Fprintf(fileHandleBBPF, "\n\nThese are the %d verified digits we have calculated via BBPF: \n", copyOfLastPosition)
                check(err2prslc2c)
                    for _, oneChar := range stringVerOfCorrectDigits {
                        //fmt.Print(oneChar) // to the console // the whole point of using an alternate file is to not clutter up the console or the default file 
                            _ , err8prslc2c := fmt.Fprint(fileHandleBBPF, oneChar)  // to a file
                                check(err8prslc2c)
                    }
            fileHandleBBPF.Close() 
        } else if useAlternateFile == "AM" {
            fileHandleAM, err1prslc2c := os.OpenFile("dataLog-From_AM_Method_lengthy_prints.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                check(err1prslc2c)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                defer fileHandleAM.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
            _, err2prslc2c := fmt.Fprintf(fileHandleAM, "\n\nThese are the %d verified digits we have calculated via Archimedes: \n", copyOfLastPosition)
                check(err2prslc2c)
                    for _, oneChar := range stringVerOfCorrectDigits {
                        fmt.Print(oneChar) // to the console  
                            _ , err8prslc2c := fmt.Fprint(fileHandleAM, oneChar)  // to a file
                                check(err8prslc2c)
                    }
            fileHandleAM.Close() 
        } else if useAlternateFile == "Nilakantha" {
            regularExpression := regexp.MustCompile(`^3.1.........................................................................................`)
            firstSectionOfPiFromWeb := regularExpression.FindStringSubmatch(piAs59766chars)
            fmt.Printf("\npi from the web begins thusly: %s ", firstSectionOfPiFromWeb)
            fmt.Printf("\npi as calculated herein begins: %0.90f ", sumBig)
            fmt.Printf("\n.... we have matched %d digits: ", copyOfLastPosition)

            fileHandleNilakan, err1prslc2c := os.OpenFile("dataLog-From_Nilakantha_Method_lengthy_prints.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                check(err1prslc2c)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                defer fileHandleNilakan.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
            _, err2prslc2c := fmt.Fprintf(fileHandleNilakan,
             "\n\nBelow are the %d verified digits we have calculated via Nilakantha using precision of %d and iterations of %d: \n",
              copyOfLastPosition, precision, arg01)
                check(err2prslc2c)
                    for _, oneChar := range stringVerOfCorrectDigits {
                        //fmt.Print(oneChar) // to the console we log pi, one digit at a time
                            _ , err8prslc2c := fmt.Fprint(fileHandleNilakan, oneChar)  // to a file we log pi one digit at a time 
                                check(err8prslc2c)
                     }
            fileHandleNilakan.Close() 
        } else if useAlternateFile == "ChudDidLessThanOneHundredLoops" {
            fileHandleDefault, err1prslc2d := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                check(err1prslc2d)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                defer fileHandleDefault.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
            _, err2prslc2d := fmt.Fprintf(fileHandleDefault, "\n\nThese are the %d verified digits we have calculated via Chudnovsky: \n", copyOfLastPosition)
                check(err2prslc2d)
                fmt.Printf("\n These are the %d verified digits we have calculated via Chudnovsky: \n", copyOfLastPosition)
                    for _, oneChar := range stringVerOfCorrectDigits {
                        fmt.Print(oneChar)
                            _ , err8prslc2d := fmt.Fprint(fileHandleDefault, oneChar)
                                check(err8prslc2d)  
                    }
            fileHandleDefault.Close()
        } else if useAlternateFile == "BBPfConcurent" {
            fileHandleBBPfConcurent, err1BBPfConcurent := os.OpenFile("dataLog-From_BBPfConcurent.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                check(err1BBPfConcurent)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                defer fileHandleBBPfConcurent.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
            _, err2BBPfConcurent := fmt.Fprintf(fileHandleBBPfConcurent,
             "\n\nBelow are the %d verified digits we have calculated via BBPfConcurent using %d workers and %d loops: \n",
             copyOfLastPosition, arg01, precision)
                check(err2BBPfConcurent)
                fmt.Printf("\nThese are the %d verified digits we have calculated via BBPfConcurent using %d workers and %d loops: \n", copyOfLastPosition, arg01, precision)
                    for _, oneChar := range stringVerOfCorrectDigits {
                        fmt.Print(oneChar)
                            _ , err8BBPfConcurent := fmt.Fprint(fileHandleBBPfConcurent, oneChar)
                                check(err8BBPfConcurent)  
                    }
            _, err9BBPc := fmt.Fprintf(fileHandleBBPfConcurent, "  TotalRun was: %s", TotalRun) 
                check(err9BBPc)
                        /*
                        // this bug came and went, very strange 
                            fmt.Println("\nµs")
                            testCase01 := "µs"
                            fmt.Fprintf(fileHandleBBPfConcurent, "\n\nThis does not render correctly µs and how about this: %s \n\n", testCase01)
                            failingString := "µs" // micro seconds 
                            fmt.Fprintf(fileHandleBBPfConcurent, "Micro Seconds should look like this: µs and this: %s", failingString)
                        */
            fileHandleBBPfConcurent.Close()

        // final else handles any instances of useAlternateFile not caught above 
        } else {
            fileHandleDefault, err1prslc2d := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                check(err1prslc2d)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                defer fileHandleDefault.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
            // to file
            _, err2prslc2d := fmt.Fprintf(fileHandleDefault, "\nThese are the %d verified digits we have calculated per case: %d  :: \n", copyOfLastPosition, selection)
                check(err2prslc2d)
            // to screen
            fmt.Printf("\n These are the %d verified digits we have calculated: \n", copyOfLastPosition)
                for _, oneChar := range stringVerOfCorrectDigits {
                    // to screen
                    fmt.Print(oneChar)
                    // to file
                    _ , err8prslc2d := fmt.Fprint(fileHandleDefault, oneChar)
                        check(err8prslc2d)  
                }
            fileHandleDefault.Close()
        }
        fmt.Print("\n")
    } // end of if's else, way up thar "if copyOfLastPosition > 55000 {} else {"   so, this has been the intance where pi is shorter than 55,000


    // next, in any instance, we add comments to the default log file if an alternate file has been used ... 
    // LOG TO THE MASTER LOGGING FILE :: dataLog-From_calculate-pi-and-friends.txt  =============================================================================

    if useAlternateFile == "chudDid800orMoreLoops" {
        fileHandleDefault, err1prslc2d := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
            check(err1prslc2d)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
            defer fileHandleDefault.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
        _, err2prslc2da := fmt.Fprint(fileHandleDefault, "\n\nThe remaining longer entries can be viewed in dataLog-From_Chudnovsky_Method_lengthy_prints.txt")
            check(err2prslc2da)
        fileHandleDefault.Close() 


    } else if useAlternateFile == "BBPF" {
        fileHandleDefault, err1prslc2d := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
            check(err1prslc2d)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
            defer fileHandleDefault.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
        _, err2prslc2da := fmt.Fprint(fileHandleDefault, "\nResults from running BBPF (case 2:) can be viewed in dataLog-From_BBPF_Method_lengthy_prints.txt\n")
            check(err2prslc2da)
        fileHandleDefault.Close() 


    } else if useAlternateFile == "BBPfConcurent" {
        fileHandleDefault, err1prslc2d := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
            check(err1prslc2d)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
            defer fileHandleDefault.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
        _, err2prslc2da := fmt.Fprint(fileHandleDefault, "\nResults from running BBPfConcurent (case 41:) can be viewed in dataLog-From_BBPfConcurent.txt\n")
            check(err2prslc2da)
        fileHandleDefault.Close() 

    } else if useAlternateFile == "Gauss–Legendre" {
        fileHandleDefault, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
            check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
            defer fileHandleDefault.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
        Hostname, _ := os.Hostname()
        _ , err0 := fmt.Fprintf(fileHandleDefault, "\n  -- Gauss–Legendre -- selection #%d on %s \n", selection, Hostname)
            check(err0)
        current_time := time.Now()
        _ , err6 := fmt.Fprint(fileHandleDefault, "was run on: ", current_time.Format(time.ANSIC), "\n")
        check(err6)
        _ , err5 := fmt.Fprintf(fileHandleDefault, "%d was total Iterations\n", arg01)
            check(err5)
        _ , err7 := fmt.Fprintf(fileHandleDefault, "Total run was %s \n ", TotalRun) 
            check(err7)
        fileHandleDefault.Close() 
    
    } else if useAlternateFile == "Nilakantha" {
        fileHandleDefault, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
            check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
            defer fileHandleDefault.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
        Hostname, _ := os.Hostname()
        _ , err0 := fmt.Fprintf(fileHandleDefault, "\n  -- Nilakantha Somayaji -- selection #%d on %s \n", selection, Hostname)
            check(err0)
        current_time := time.Now()
        _ , err6 := fmt.Fprint(fileHandleDefault, "was run on: ", current_time.Format(time.ANSIC), "\n")
        check(err6)
        _ , err5 := fmt.Fprintf(fileHandleDefault, "%d was total Iterations; %d was precision setting for the big.Float types \n", arg01, precision)
            check(err5)
        //TotalRun := elapsed.String() // e.g., it was cast to a String type in another func and passed to this func as TotalRun
        _ , err7 := fmt.Fprintf(fileHandleDefault, "Total run was %s \n ", TotalRun) 
            check(err7)
        _, err2prslc2da := fmt.Fprint(fileHandleDefault, "\nResults from running Nilakantha (case 5:) can be viewed in dataLog-From_Nilakantha_Method_lengthy_prints.txt\n")
            check(err2prslc2da)
        fileHandleDefault.Close() 
    }
} // end of func printResultStatsLong(sumBig *big.Float, precision int, useAlternateFile string, arg01 int, TotalRun string, selection int) 
// -- AMFprintResultStatsLongB 


// --------------- Display code to user functions follow -----------------------------------------------------------------------------------------//
//------------------------------------------------------------------------------------------------------------------------------------------------//

func Show_Gauss_Legendre(filenameOfThisFile string) { // case 57: 
    pattern1 := regexp.MustCompile(`(?s)AMFGauss_LegendreA.*?`)  // Create a regular expression that matches the first flag anywhere in the line
    pattern2 := regexp.MustCompile(`(?s)AMFGauss_LegendreB.*?`)  // Create a regular expression that matches the second flag anywhere in the line
    doTheRest(pattern1, pattern2, filenameOfThisFile)
}
func ShowArchimedesBig(filenameOfThisFile string) {  // case 34: 

    pattern1 := regexp.MustCompile(`(?s)AMFArchimedesBigA.*?`)  // Create a regular expression that matches the first flag anywhere in the line
    pattern2 := regexp.MustCompile(`(?s)AMFArchimedesBigB.*?`)  // Create a regular expression that matches the second flag anywhere in the line
    doTheRest(pattern1, pattern2, filenameOfThisFile)
}
func ShowMain(filenameOfThisFile string) {  // case 3: 
    pattern1 := regexp.MustCompile(`(?s)AMFmainA.*?`)  // Create a regular expression that matches the first flag anywhere in the line
    pattern2 := regexp.MustCompile(`(?s)AMFmainB.*?`)  // Create a regular expression that matches the second flag anywhere in the line
    doTheRest(pattern1, pattern2, filenameOfThisFile)
}
func ShowMenus(filenameOfThisFile string) {  // case 4: 
    pattern1 := regexp.MustCompile(`(?s)AMFmenusA.*?`)  // Create a regular expression that matches the first flag anywhere in the line
    pattern2 := regexp.MustCompile(`(?s)AMFmenusB.*?`)  // Create a regular expression that matches the second flag anywhere in the line
    doTheRest(pattern1, pattern2, filenameOfThisFile)
}
func ShowxRootOfY_continuousCaller(filenameOfThisFile string) {  // case 380: 
    pattern1 := regexp.MustCompile(`(?s)AMFxroyCCA.*?`)  // Create a regular expression that matches the first flag anywhere in the line
    pattern2 := regexp.MustCompile(`(?s)AMFxroyCCB.*?`)  // Create a regular expression that matches the second flag anywhere in the line
    doTheRest(pattern1, pattern2, filenameOfThisFile)
}
func ShowxRootOfy(filenameOfThisFile string) {  // case 38: 
    pattern1 := regexp.MustCompile(`(?s)AMFxRootOfyA.*?`)  // Create a regular expression that matches the first flag anywhere in the line
    pattern2 := regexp.MustCompile(`(?s)AMFxRootOfyB.*?`)  // Create a regular expression that matches the second flag anywhere in the line
    doTheRest(pattern1, pattern2, filenameOfThisFile)
}
func ShowGottfriedWilhelmLeibniz(filenameOfThisFile string) {  // case 26: 
    pattern1 := regexp.MustCompile(`(?s)AMFGottfriedWilhelmLeibnizA.*?`)  // Create a regular expression that matches the first flag anywhere in the line
    pattern2 := regexp.MustCompile(`(?s)AMFGottfriedWilhelmLeibnizB.*?`)  // Create a regular expression that matches the second flag anywhere in the line
    doTheRest(pattern1, pattern2, filenameOfThisFile)
}
func ShowGregoryLeibniz(filenameOfThisFile string) {  // case 27: 
    pattern1 := regexp.MustCompile(`(?s)AMFGregoryLeibnizA.*?`)  // Create a regular expression that matches the first flag anywhere in the line
    pattern2 := regexp.MustCompile(`(?s)AMFGregoryLeibnizB.*?`)  // Create a regular expression that matches the second flag anywhere in the line
    doTheRest(pattern1, pattern2, filenameOfThisFile)
}
func ShowJohnWallis(filenameOfThisFile string) {  // case 28: 
    pattern1 := regexp.MustCompile(`(?s)AMFJohnWallisA.*?`)  // Create a regular expression that matches the first flag anywhere in the line
    pattern2 := regexp.MustCompile(`(?s)AMFJohnWallisB.*?`)  // Create a regular expression that matches the second flag anywhere in the line
    doTheRest(pattern1, pattern2, filenameOfThisFile)
}
func ShowEulersNumber(filenameOfThisFile string) {  // case 29: 
    pattern1 := regexp.MustCompile(`(?s)AMFEulersNumberA.*?`)  // Create a regular expression that matches the first flag anywhere in the line
    pattern2 := regexp.MustCompile(`(?s)AMFEulersNumberB.*?`)  // Create a regular expression that matches the second flag anywhere in the line
    doTheRest(pattern1, pattern2, filenameOfThisFile)
}
func ShowErdosBorwein(filenameOfThisFile string) {  // case 30: 
    pattern1 := regexp.MustCompile(`(?s)AMFErdosBorweinA.*?`)  // Create a regular expression that matches the first flag anywhere in the line
    pattern2 := regexp.MustCompile(`(?s)AMFErdosBorweinB.*?`)  // Create a regular expression that matches the second flag anywhere in the line
    doTheRest(pattern1, pattern2, filenameOfThisFile)
}
func ShowcompareFloat64withBigFloats(filenameOfThisFile string) {  // case 70: 
    pattern1 := regexp.MustCompile(`(?s)AMFcompareFloat64withBigFloatsA.*?`)  // Create a regular expression that matches the first flag anywhere in the line
    pattern2 := regexp.MustCompile(`(?s)AMFcompareFloat64withBigFloatsB.*?`)  // Create a regular expression that matches the second flag anywhere in the line
    doTheRest(pattern1, pattern2, filenameOfThisFile)
}
func ShowNilakantha_Somayaji_with_big_Float_types(filenameOfThisFile string) {  // case 25: 
    pattern1 := regexp.MustCompile(`(?s)AMFNilakantha_Somayaji_with_big_Float_typesA.*?`)  // Create a regular expression that matches the first flag anywhere in the line
    pattern2 := regexp.MustCompile(`(?s)AMFNilakantha_Somayaji_with_big_Float_typesB.*?`)  // Create a regular expression that matches the second flag anywhere in the line
    doTheRest(pattern1, pattern2, filenameOfThisFile)
}
func Showchud(filenameOfThisFile string) {  // case 35: 
    pattern1 := regexp.MustCompile(`(?s)AMFchudA.*?`)  // Create a regular expression that matches the first flag anywhere in the line
    pattern2 := regexp.MustCompile(`(?s)AMFchudB.*?`)  // Create a regular expression that matches the second flag anywhere in the line
    doTheRest(pattern1, pattern2, filenameOfThisFile)
}
func ShowBBPF(filenameOfThisFile string) {  // case 22: 
    pattern1 := regexp.MustCompile(`(?s)AMFBBPFA.*?`)  // Create a regular expression that matches the first flag anywhere in the line
    pattern2 := regexp.MustCompile(`(?s)AMFBBPFB.*?`)  // Create a regular expression that matches the second flag anywhere in the line
    doTheRest(pattern1, pattern2, filenameOfThisFile)
}
func ShowLeibniz_method_one_billion_iters(filenameOfThisFile string) {  // case 64: 
    pattern1 := regexp.MustCompile(`(?s)AMFLeibniz_method_one_billion_itersA.*?`)  // Create a regular expression that matches the first flag anywhere in the line
    pattern2 := regexp.MustCompile(`(?s)AMFLeibniz_method_one_billion_itersB.*?`)  // Create a regular expression that matches the second flag anywhere in the line
    doTheRest(pattern1, pattern2, filenameOfThisFile)
}
func ShowBBPfConcurent(filenameOfThisFile string) {  // case 61: 
    pattern1 := regexp.MustCompile(`(?s)AMFBBPfConcurentA.*?`)  // Create a regular expression that matches the first flag anywhere in the line
    pattern2 := regexp.MustCompile(`(?s)AMFBBPfConcurentB.*?`)  // Create a regular expression that matches the second flag anywhere in the line
    doTheRest(pattern1, pattern2, filenameOfThisFile)
}
func Showbbp_formula(filenameOfThisFile string) {  // case 62: 
    pattern1 := regexp.MustCompile(`(?s)AMFbbp_formulaA.*?`)  // Create a regular expression that matches the first flag anywhere in the line
    pattern2 := regexp.MustCompile(`(?s)AMFbbp_formulaB.*?`)  // Create a regular expression that matches the second flag anywhere in the line
    doTheRest(pattern1, pattern2, filenameOfThisFile)
}
func Shownifty_scoreBoard(filenameOfThisFile string) {  // case 60: 
    pattern1 := regexp.MustCompile(`(?s)AMFnifty_scoreBoardA.*?`)  // Create a regular expression that matches the first flag anywhere in the line
    pattern2 := regexp.MustCompile(`(?s)AMFnifty_scoreBoardB.*?`)  // Create a regular expression that matches the second flag anywhere in the line
    doTheRest(pattern1, pattern2, filenameOfThisFile)
}
func ShowTheSpigot(filenameOfThisFile string) {  // case 39: 
    pattern1 := regexp.MustCompile(`(?s)AMFTheSpigotA.*?`)  // Create a regular expression that matches the first flag anywhere in the line
    pattern2 := regexp.MustCompile(`(?s)AMFTheSpigotB.*?`)  // Create a regular expression that matches the second flag anywhere in the line
    doTheRest(pattern1, pattern2, filenameOfThisFile)
}
//func doTheRest(pattern1 *regexp.Regexp, pattern2 *regexp.Regexp, file *os.File) {  // decided that it would not be prudent to pass the file handle
func doTheRest(pattern1 *regexp.Regexp, pattern2 *regexp.Regexp, filenameOfThisFile string) {
    file, err := os.Open(filenameOfThisFile)  // create a handle to our file
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
        alreadyFoundFirstTag := false   // Create a boolean variable to track if we are within the desired range of lines we want to print
        scanner := bufio.NewScanner(file)    // Create a scanner to read the file line by line
        countTheLines := 1
    for scanner.Scan() {        // Iterate over the lines in our file
        countTheLines++
        line := scanner.Text()  // read a line and store the line in var line
        // deal with the case of having read the line containing the first flag 
        if pattern1.MatchString(line) { // if line contains pattern1
            alreadyFoundFirstTag = !alreadyFoundFirstTag // flips it to true 
            fmt.Println(line) // print line containing first flag 
            continue // we read another line
        }
        // deal with the case of having encountered the line containing the second flag
        if pattern2.MatchString(line) {   // if line contains pattern2
            fmt.Println(line)
            break  // we break to stop reading or printing additional lines
        }
        // deal with the case of having already found and printed the line containing the first flag 
        // print the line following the line containing the first flag, and all interviening lines
        if alreadyFoundFirstTag { 
            fmt.Println(line)
            continue // read another line
        }
    }
    fmt.Println("The total number of lines in this set is: ", countTheLines)
}