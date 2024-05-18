package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"strconv"
	"time"
)

func menu_switch(filenameOfThisFile string) {
	switch selection { // if selection == 2 then handel per case 2: etc.

	case 16: // Monte Carlo method (added May 18 2024)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the grid size to be used (10000 will give an acceptable result in about 21s) : ")
		input, _ := reader.ReadString('\n')
		gridSize, err := strconv.Atoi(input[:len(input)-1])
		if err != nil {
			fmt.Println("Invalid input, please enter an integer.")
			return
		}
		fmt.Printf("Size of the grid has been set to: %d\n", gridSize)
		if gridSize > 1000 {
			fmt.Println("working ...")
		}
		piApprox := GridPi(gridSize)
		fmt.Printf("Approximated Pi as big float: %s\n", piApprox.Text('f', 30))
		piApproxFloat64, _ := piApprox.Float64() // Convert piApprox to a float64 type
		fmt.Printf("Approximated Pi as float64:   %f\n", piApproxFloat64)
		piFromMathLib := math.Pi                       // Obtain Pi from math library
		piFromMathLibBF := big.NewFloat(piFromMathLib) // Create a big float object version of Pi from math library
		fmt.Printf("Pi from Math Library:         %s\n", piFromMathLibBF.Text('f', 30))
		fmt.Printf("Difference: %f\n", math.Abs(piApproxFloat64-math.Pi))

	case 999:
		deleteAllLogFiles()
	case 20999:
		ShowdeleteAllLogFiles()

	case 3:
		ShowMain()

	case 2: // (best method) the Bailey–Borwein–Plouffe formula for π, circa 1995
		BBPF(selection)
	case 22:
		ShowBBPF()

	case 4:
		ShowMenus()

	case 5: // a series by Indian astronomer Nilakantha Somayaji circa 1500 AD
		Nilakantha_Somayaji_with_big_Float_types(selection)
	case 25:
		ShowNilakantha_Somayaji_with_big_Float_types()

	case 6: // Gottfried Wilhelm Leibniz
		GottfriedWilhelmLeibniz(selection)
	case 26:
		ShowGottfriedWilhelmLeibniz()

	case 7: // the Gregory-Leibniz series
		GregoryLeibniz(selection)
	case 27:
		ShowGregoryLeibniz()

	case 8: // John Wallis circa 1655 --- runs very long
		JohnWallis(selection)
	case 28:
		ShowJohnWallis()

	case 9: // Euler's Number
		EulersNumber(selection)
	case 29:
		ShowEulersNumber()

	case 10: // Erdős–Borwein constant
		ErdosBorwein(selection)
	case 30:
		ShowErdosBorwein()

	case 11: // display a review of the derivation of the Pythagorean
		DisplayPythagorean(selection)
	case 31:
		DisplayPythagoreanCode()

	// cases related to the menu for log file viewing
	case 202:
		content, err := ioutil.ReadFile("dataLog-From_BBPF_Method_lengthy_prints.txt") //
		if err != nil {                                                                // if the file does not exist ...
			fmt.Println(colorCyan, "\nNo prior results -- no log file", colorWhite, "'dataLog-From_BBPF_Method_lengthy_prints.txt'", colorCyan,
				"exists\n")
		} else {
			fmt.Println(string(content)) // dump/display entire file to command line
		}
	case 214:
		content, err := ioutil.ReadFile("dataLog-From_AM_Method_lengthy_prints.txt") //
		if err != nil {                                                              // if the file does not exist ...
			fmt.Println(colorCyan, "\nNo prior results -- no log file", colorWhite, "'dataLog-From_AM_Method_lengthy_prints.txt'", colorCyan,
				"exists\n")
		} else {
			fmt.Println(string(content)) // dump/display entire file to command line
		}
	case 215:
		content, err := ioutil.ReadFile("dataLog-From_Chudnovsky_Method_lengthy_prints.txt") //
		if err != nil {                                                                      // if the file does not exist ...
			fmt.Println(colorCyan, "\nNo prior results -- no log file", colorWhite, "'dataLog-From_Chudnovsky_Method_lengthy_prints.txt'", colorCyan,
				"exists\n")
		} else {
			fmt.Println(string(content)) // dump/display entire file to command line
		}
	case 315:
		content, err := ioutil.ReadFile("big_pie_is_in_here.txt") //
		if err != nil {                                           // if the file does not exist ...
			fmt.Println(colorCyan, "\nNo prior results -- no log file", colorWhite, "'big_pie_is_in_here.txt'", colorCyan,
				"exists\n")
		} else {
			fmt.Println(string(content)) // dump/display entire file to command line
		}
	case 205:
		content, err := ioutil.ReadFile("dataLog-From_Nilakantha_Method_lengthy_prints.txt") //
		if err != nil {                                                                      // if the file does not exist ...
			fmt.Println(colorCyan, "\nNo prior results -- no log file", colorWhite, "'dataLog-From_Nilakantha_Method_lengthy_prints.txt'", colorCyan,
				"exists\n")
		} else {
			fmt.Println(string(content)) // dump/display entire file to command line
		}
	case 212:
		content, err := ioutil.ReadFile("dataLog-From_calculate-pi-and-friends.txt") //
		if err != nil {                                                              // if the file does not exist ...
			fmt.Println(colorCyan, "\nNo prior results -- no log file", colorWhite, "'dataLog-From_calculate-pi-and-friends.txt'", colorCyan,
				"exists\n")
		} else {
			fmt.Println(string(content)) // dump/display entire file to command line
		}
	case 241:
		content, err := ioutil.ReadFile("dataLog-From_BBPfConcurent.txt") //
		if err != nil {                                                   // if the file does not exist ...
			fmt.Println(colorCyan, "\nNo prior results -- no log file", colorWhite, "'dataLog-From_BBPfConcurent.txt'", colorCyan,
				"exists\n")
		} else {
			fmt.Println(string(content)) // dump/display entire file to command line
		}
	case 280:
		content, err := ioutil.ReadFile("logfile_from_selection_180.txt") //
		if err != nil {                                                   // if the file does not exist ...
			fmt.Println(colorCyan, "\nNo prior results -- no log file", colorWhite, "'logfile_from_selection_180.txt'", colorCyan,
				"exists\n")
		} else {
			fmt.Println(string(content)) // dump/display entire file to command line
		}

	case 180:
		xRootOfY_continuousCaller(selection)
	case 380:
		ShowxRootOfY_continuousCaller()

	case 14:
		ArchimedesBig(selection)
	case 34:
		ShowArchimedesBig()

	case 15:
		chud(selection)
	case 35:
		Showchud()

	case 18:
		xRootOfy(selection)
	case 38:
		ShowxRootOfy()

	case 19:
		TheSpigot(selection)
	case 39:
		ShowTheSpigot()

	case 37:
		Gauss_Legendre(selection)
	case 57:
		Show_Gauss_Legendre()

	case 40: // a special case ... if some time has elapsed running the func then we log the stats to a file (because this func is always killed with a ctrl-c)
		start := time.Now() // saved start time to be compared with end time t
		nifty_scoreBoard()
		t := time.Now()
		elapsed := t.Sub(start)
		// only if
		if int(elapsed.Seconds()) != 0 {
			fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
			check(err1)                                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
			defer fileHandle.Close()                                                                                                // It’s idiomatic to defer a Close immediately after opening a file.
			Hostname, _ := os.Hostname()
			_, err0 := fmt.Fprintf(fileHandle, "\n  -- ScoreBoard -- selection #%d on %s \n", selection, Hostname)
			check(err0)
			current_time := time.Now()
			_, err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
			check(err6)
			TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
			_, err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun)
			check(err7)
		}
	case 60:
		Shownifty_scoreBoard()

	case 41:
		BBPfConcurent(selection)
	case 61:
		ShowBBPfConcurent()

	case 42:
		bbp_formula(selection)
	case 62:
		Showbbp_formula()

	case 44:
		Leibniz_method_one_billion_iters(selection)
	case 64:
		ShowLeibniz_method_one_billion_iters()

	case 47:
		os.Exit(1)

	case 50:
		compareFloat64withBigFloats()
	case 70:
		ShowcompareFloat64withBigFloats()

	case 96:
		Using_this_app()

	case 0:
	// Only as a non-functional case to allow returning directly to main menu

	default:
		fmt.Println("\n ... You entered a value that is", colorRed, "not a valid option", colorReset, "go fish\n")
		// handles every case not explicitly listed (which has not been otherwise handled)
	}
}
