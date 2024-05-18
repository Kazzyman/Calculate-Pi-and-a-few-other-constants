package main

// Unix-0713-0524-v2-dev-2023.go

// The initial inspiration for all of this was : Veritassium https://youtu.be/gMlf1ELvRzc?list=LL
/*
   Go.lang is a relatively-new general-purpose freely-available programming language developed by google.

   One can obtain the Go Compiler from : https://go.dev/dl/
*/

// Unix/Linux/Mac variant build instructions : ==========================================================================
/*
   compile with: "go build -o desired-name-of-your-executable name-of-this-source-code-file.go"
   ... thereafter you can run it on a Unix system with "/fullpathname/desired-name-of-your-executable"
       ( e.g.  "./Richards_go_project" )
   ... or, having first obtained the Go compiler, ... just run the source with: "go run name-of-this-source-code-file.go"
*/
// =======================================================================================================================

// Windows variant : ----------------------------------------------------------------------------------------------------
/*
   build with: "go build -o desired-name-of-your-executable.exe name-of-this-source-code-file.go"
   ... or, having first obtained the Go compiler, ... just run the source with: "go run name-of-this-source-code-file.go"
*/
// -----------------------------------------------------------------------------------------------------------------------
// change
// -- AMFmainA-a

import (
	"fmt"    // Used for printing etc.
	"regexp" // used in multiple instance to create regular expression patterns
)

func main() { // top-level program logic flow -- explore several ways to calculate Pi, plus THREE other constants --AMFmain-A-b

	// filenameOfThisFile := getTheNameOfThisSourceCodeFile() // a locally-defined func
	filenameOfThisFile := "/Users/quasar/pi-main/main.go"
	totalLines, nonEmptyLines := reportSLOCstats(filenameOfThisFile) // another locally-defined func; returns, and creates, local values of predetermined type

	// The following is for menu header data generation (inception: stripped file name) //  <<------------ below ---------<<
	// ... there should be exactly nine of each comment tag in this dev Windows ver
	// /* Unix variant
	re2 := regexp.MustCompile(`Unix-(.+)\.go`)
	// Unix variant */
	//
	/* Windows variant
	    re2 := regexp.MustCompile(`win-(.+)\.go`)
	Windows variant */

	match2 := re2.FindStringSubmatch(filenameOfThisFile) // grab the stuff between win- and .go (or Unix- and .go) in the name of this file
	SansVerOfNameOfThisFile := "pi-main"                 // this var, having been initialize to "string", is of type string
	if len(match2) >= 2 {

		// /* Unix variant
		SansVerOfNameOfThisFile = match2[1] // Unix SansVerOfNameOfThisFile is loaded with our base file name
		// Unix variant */
		//
		/* Windows variant
		    SansVerOfNameOfThisFile = match2[1]                   // Windows SansVerOfNameOfThisFile is loaded with our base file name
		Windows variant */

	} else {
		fmt.Println("SansVerOfNameOfThisFile via match2 not found in main")
	}

	for {
		DisplayMenus(totalLines, nonEmptyLines, filenameOfThisFile, SansVerOfNameOfThisFile)
		fmt.Println(colorRed, "Finished. Hit Enter to redisplay the Main Menu", colorReset) // this will be the last line of every case #:
		var JustToPauseTheLoop int

		// /* Unix variant
		fmt.Scanf("%d", &JustToPauseTheLoop) // must pause the endless loop after every switch case to prevent menu redisplay over results
		// fmt.Scan(&JustToPauseTheLoop) // Scan() would work as a pause, but it requires data input to continue, so we use a Scanf() instead
		// Unix variant */
		//
		/* Windows variant
		    fmt.Scanf("%d", &JustToPauseTheLoop) // must pause the endless loop after every switch case to prevent menu redisplay over results
		    fmt.Scanf("%d", &JustToPauseTheLoop)
		        //fmt.Scan(&JustToPauseTheLoop) // Scan does not work quite as well as the Scanf combo above (but a single Scanf worked well in Unix environment)
		Windows variant */

	}
} // end of main --AMFmain-A-c

func check(e error) { // create a func named check which takes one parameter "e" of type error
	if e != nil {
		panic(e) // use panic() to display error code
	}
}
