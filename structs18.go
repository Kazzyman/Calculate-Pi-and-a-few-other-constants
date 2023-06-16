package main
import (
    "os"        // Used to fetch the name of your system to be used in the log file 
    "fmt"       // Used for printing etcetera 
    "math"      // Used for math.Pow(base, exponent), Sqrt, and Cbrt funcs 
    "time"      // Used for seeding random number generation; and calculating run times 
    "sort"      // Used in case 18 to sort an array of structs 
)
type Results struct {  // define a new structure called Results with two fields; result, and pdiff 
    result float64
    pdiff float64
}

// common globals 
        var colorReset = "\033[0m"          
        var colorRed = "\033[31m"
        var colorGreen = "\033[32m"
        var colorPurple = "\033[35m"
        var colorCyan = "\033[36m"

    // the following globals, are used in multiple funcs of case 18: calculate either square or cube root of any integer 
    //------------------------------------------------------------------------------------------------------------------

        var sortedResults = []Results{} // sortedResults is an array of type Results as defined at the top of this file 
        var Table_of_perfect_Products = []int{}  
        var diffOfLarger int 
        var diffOfSmaller int 
        var perfectResult2 float64 // will contain the square root result if the workpiece was itself a perfect square 
        var perfectResult3 float64 // will contain the cube root result if the workpiece was itself a perfect cube 
        var precisionOfRoot int // this being global means we do not need to pass it in to the read func 
        var radical_index int // 2 for square roots, 3 for cube roots, no need to pass this one around either  
        var workPiece int // the square or cube of which we are to find a root 
        var skip_redoing_loop int 
            // the following three arrays were my original method of logging my results and the nearness of fit 
                //var List_of_2_results_case18 = []float64{}
                //var corresponding_diffs = []int{}
                //var diffs_as_percent = []float64{}

func main() {   
    xRootOfy(1)
}

func xRootOfy(num int) { // calculates either square or cube root of any integer 

    var index = 0 // counter used in the for loop in this func :: is also passed to the pricipal func readTheTableOfPP // cannot be a global

    getInputFromUserAndSetPrecision() // sets radical index also 

        start := time.Now() // saved start time to be compared with end time t (why have two start times, this local one, and a second global one?)

    buildTableOfPerfectProducts() 


// The following section consists of the principal for loop with a conditional break ------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------------------------------------------------------

    for index < 380000 { // the table has 800,000 entries

        readTheTableOfPP(index, start)  // pass-in the index to the table and start time  

// the next sub-section detects, traps, and reports the detection of either a perfect square of a perfect cube ------------------
// ... it also is responsible for causing the algorithm to terminate if workpiece was a perfect square or cube 
//-------------------------------------------------------------------------------------------------------------------------------
        if diffOfLarger == 0 || diffOfSmaller == 0 {  // Then, it was a perfect square or cube 
            t := time.Now()          // t is current time
            elapsed := t.Sub(start) // initial start time (start) is used to calculate and log final run time as compared to t 

            fmt.Println("Total run time was:", float64(elapsed.Seconds()), "seconds.\n")

                fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) 
                    check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                    defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                Hostname, _ := os.Hostname()
                _ , err0 := fmt.Fprintf(fileHandle, "\n  -- %d root of %d by a ratio of PerfectProducts -- selection #%d on %s \n", radical_index, 
                    workPiece, num, Hostname)
                    check(err0)
                current_time := time.Now()
                _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                check(err6)
                TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
                    check(err7)
                    if radical_index == 2 {
                        _ , err8 := fmt.Fprintf(fileHandle, "the %d root of %d is %0.2f \n", radical_index, workPiece, perfectResult2)
                            check (err8)
                    }
                    if radical_index == 3 {
                        _ , err8 := fmt.Fprintf(fileHandle, "the %d root of %d is %0.2f \n", radical_index, workPiece, perfectResult3)
                            check (err8)
                    }

            break // break out of the for loop because we are done : the workpiece was either a perfect square of a perfect cube 

        } // end of if
        index = index + 2 // 
    } // end of for loop // the above break statement is NOT the only way to exit this for loop, it terminates after 380,000 iterations per index 

//-------------------------------------------------------------------------------------------------------------------------------


// the remaining section is only reached after having exited the primary for loop above via a break statement or an exaustive reading of the table --------------
//---------------------------------------------------------------------------------------------------------------------------------------------------------------
        // calculate elapsed time 
        t := time.Now()
        elapsed := t.Sub(start)

// the following section logs the results to a text file ----------------------------------------------------------------------------------------------
                    fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file 
                        check(err1)                                // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
                        defer fileHandle.Close()                  // It’s idiomatic to defer a Close immediately after opening a file.
                    Hostname, _ := os.Hostname()
                    _ , err0 := fmt.Fprintf(fileHandle, "\n  -- %d root of %d by a ratio of perfect Products -- selection #%d on %s \n", radical_index, workPiece, num, Hostname)
                        check(err0)
                    current_time := time.Now()
                    _ , err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
                    check(err6)
                    _ , err5 := fmt.Fprintf(fileHandle, "%d was total Iterations \n", index)
                        check(err5)

        // Sort the slice sortedResults by its pdiff field : 
        //-------------------------------------------------------------------------------------------------------------------------------
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

        // print the best fitting result based solely on the lowest pdiff
        //---------------------------------------------------------------
                if radical_index == 2 {
                fmt.Printf("%0.9f, is the best approximation for the Square Root of %d \n", sortedResults[0].result, workPiece)
                }
                if radical_index == 3 {
                        fmt.Printf("%0.9f, is the best approximation for the Cube Root of %d \n", sortedResults[0].result, workPiece)
                }
                    // print the best fitting result 
                    if radical_index == 2 {
                        _ , err8 := fmt.Fprintf(fileHandle, "%0.9f, is the best approximation for the Square Root of %d \n", sortedResults[0].result, workPiece)
                        check(err8)
                    }
                    if radical_index == 3 {
                        _ , err9 := fmt.Fprintf(fileHandle, "%0.9f, is the best approximation for the Cube Root of %d \n", sortedResults[0].result, workPiece)
                        check(err9)
                    }
        //---------------------------------------------------------------

                    TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
                    _ , err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) 
                        check(err7)


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
// -------------------------------------------------------------------------------------------------------------------------------------------------------

            
}

func readTheTableOfPP (index int, start time.Time) {  // this gets called 380,000 times. 
    // The first time it is called index is 0

    // read it ...
    smallerPerfectProductOnce := Table_of_perfect_Products[index]  
    // ... and save it locally, do this just-once per func call. // ... index may be 0 up to 380,000
    RootOfsmallerPerfectProductOnce := Table_of_perfect_Products[index+1]
    // ^^^ also read the root wich corresponds

        iter := 0
    for iter < 420000 { // 420,000 loops. Why do we need so many?, Because index may be 0 to 380,000 ?? and we need to read through 830,000 table entries 
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
                fmt.Println(string(colorCyan), "\n The", radical_index, "root of", workPiece, "is", string(colorGreen), float64(rootOfProspectiveHitOnLargeSide) / float64(RootOfsmallerPerfectProductOnce), string(colorReset), "\n")
                perfectResult2 = (math.Sqrt(float64(workPiece)))
                perfectResult3 = math.Cbrt(float64(workPiece)) 
                break // out of the for loop because the workPiece is itself a perfect square
            }
            if diffOfSmaller == 0 {
                fmt.Println(string(colorCyan), "\n The", radical_index, "root of", workPiece, "is", string(colorGreen), float64(rootOfProspectiveHitOnSmallerSide) / float64(RootOfsmallerPerfectProductOnce), string(colorReset), "\n")
                perfectResult2 = (math.Sqrt(float64(workPiece)))
                perfectResult3 = math.Cbrt(float64(workPiece)) 
                break // out of the for loop because the workPiece is itself a perfect square
            }
// ---------------------------------------------------------------------------------------------------------

            if diffOfLarger < precisionOfRoot {  // report the prospects, their differences, and the calculated result for the Sqrt of 3
                fmt.Println("small PP is", string(colorCyan), smallerPerfectProductOnce, string(colorReset), "and, slightly on the higher side of", workPiece, "* that we found a PP of", string(colorCyan),
                 ProspectiveHitOnLargeSide, string(colorReset), "a difference of", diffOfLarger)

                fmt.Println("the ", radical_index, " root of ", workPiece, " is calculated as ", string(colorGreen), 
                float64(rootOfProspectiveHitOnLargeSide) / float64(RootOfsmallerPerfectProductOnce),
                 string(colorReset))

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
                elapsed2 := t2.Sub(start)
// if needed, notify the user that we are still working 
                if float64(elapsed2.Seconds()) > 0.14  {
                    fmt.Println(float64(elapsed2.Seconds()), "Seconds have elapsed ... working ...\n")
                }
            }


            if diffOfSmaller < precisionOfRoot {  // report the prospects, their differences, and the calculated result for the Sqrt or Cbrt 
                fmt.Println("small PP is", string(colorCyan), smallerPerfectProductOnce, string(colorReset), "and, slightly on the lesser side of", workPiece, "* that we found a PP of", string(colorCyan),
                 ProspectiveHitOnSmallerSide, string(colorReset), "a difference of", diffOfSmaller)

                fmt.Println("the ", radical_index, " root of ", workPiece, " is calculated as ", string(colorGreen), 
                float64(rootOfProspectiveHitOnSmallerSide) / float64(RootOfsmallerPerfectProductOnce),
                 string(colorReset))

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
                elapsed2 := t2.Sub(start)
// if needed, notify the user that we are still working 
                if float64(elapsed2.Seconds()) > 0.14  {
                    fmt.Println(float64(elapsed2.Seconds()), "Seconds have elapsed ... working ...\n")
                }
            }
            break // out of the for loop if we found any prospects using the current index value 
        }
    }
}

func getInputFromUserAndSetPrecision() {
    radical_index = 2
    skip_redoing_loop = 0 
    for radical_index > 1 && skip_redoing_loop == 0 { 
        // radical_index will initially be 2, and the skip flag will be initially 0, therefore we will initially enter this for loop 
        fmt.Println("\n\nEnter 2 for SquareRoot or 3 for CubeRoot")
        fmt.Scan(&radical_index) // just hitting enter will leave radical_index set as 2  // same in Unix or Windows variants 

        for radical_index > 0 { // but, if we entered any value for radical_index 
            if radical_index < 2 || radical_index > 3 { // only then do we check to see if it is a 2 or a 3
                fmt.Println(string(colorPurple), "\n No, for the radical, enter either 2 or 3, just a 2 or a 3, \n ... no other number for the radical please \n", string(colorReset))
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
            fmt.Println(string(colorPurple), "\n Zero is dopey, for the radical, enter either 2 or 3, just a 2 or a 3, \n ... no other number for the radical please \n", string(colorReset))
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
}


// Build a table of 810,000 pairs of PPs with their roots, does either squares or cubes:
//--------------------------------------------------------------------------------------

func buildTableOfPerfectProducts() {

    var PerfectProduct int 
    Table_of_perfect_Products = nil // this fixed my bug 
    root := 10 
            iter := 0
        for iter < 810000 { // a table of 810,000 pairs of PPs with their roots ought to do it !!
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


// auxilary funcs already present in Pi_and_friends:
//--------------------------------------------------

func check(e error) {   // create a func named check which takes one parameter "e" of type error 
    if e != nil {
        panic(e)        // use panic() to display error code 
    }
}