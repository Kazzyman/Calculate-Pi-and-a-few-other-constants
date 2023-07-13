package main

// refactorUnix2Win.go

import (
	"os"
    "fmt"       // Used for printing etc.
    "regexp"
    "bufio"     // used to count the lines in this file
)

var nameOfOldSourceCodeFile = "Unix-0713-0446-v2-dev-2023.go"  
var nameOfNewSourceCodeFile = "targetFileWindows.go"

var countOfWindowsStart = 0 
var countOfWindowsEnd = 0 
var countOfUnixStart = 0 
var countOfUnixEnd = 0 

func main() {

    fileHandle, err1 := os.Open(nameOfOldSourceCodeFile)       // open the source code file and get a handle to it 
    if err1 != nil {
        fmt.Println("error opening source per rick")
    }
    defer fileHandle.Close()
        scanner := bufio.NewScanner(fileHandle)
        for scanner.Scan() {
            //lineRead := strings.TrimSpace(scanner.Text())   // read a line from the file into lineRead
            lineRead := scanner.Text() // this fixes my main problem re formatting getting wasted 

					regexWindowsStart 		:= regexp.MustCompile(`\s*\/\*\s*Windows\s*variant`) 			// ... /* Windows variant
					regexWindowsEnd			:= regexp.MustCompile(`\s*Windows\s*variant\s*\*\/`)			// ... Windows variant */
					regexUnixStart			:= regexp.MustCompile(`\s*\/\/\s*\/\*\s*Unix\s*variant`)	 	// ... // /* Unix variant
					regexUnixEnd			:= regexp.MustCompile(`\s*\/\/\s*Unix\s*variant\s*\*\/`)		// ... // Unix variant */  

            // write the read line unchanged to the output file if any of these four strings are not found in the read line from the input file
   			if !regexWindowsStart.MatchString(lineRead) && !regexWindowsEnd.MatchString(lineRead) && !regexUnixStart.MatchString(lineRead) && !regexUnixEnd.MatchString(lineRead) {
	            OutputFileHandle, err2 := os.OpenFile(nameOfNewSourceCodeFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // write to file 
	            check(err2)                               // ... gets a file handle for output 
	            defer OutputFileHandle.Close()          // It’s idiomatic to defer a Close immediately after opening a file.
	        		_, err3 := fmt.Fprint(OutputFileHandle, lineRead)
	            	check(err3)
	            	_, err6 := fmt.Fprint(OutputFileHandle, "\n")
	            	check(err6)
	        		//OutputFileHandle.Close() 
            } 


            // if they ARE found, replace them with the following edited lines 
            //if lineRead == `(\s*\/\/\s*\/\*\s*Windows\s*variant)` {
            if regexWindowsStart.MatchString(lineRead) {
            	countOfWindowsStart++
                // write the edited line to the output file
	            OutputFileHandle, err4 := os.OpenFile(nameOfNewSourceCodeFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // write to file 
	            check(err4)                               // ... gets a file handle for output 
	            defer OutputFileHandle.Close()          // It’s idiomatic to defer a Close immediately after opening a file.
	        		_, err5 := fmt.Fprint(OutputFileHandle, "            /* Windows variant\n")
	            	check(err5)
            } 
            //if lineRead == `(\s*\/\/\s*\s*Windows\s*variant)` {
            if regexWindowsEnd.MatchString(lineRead) {
            	countOfWindowsEnd++
                // write the edited line to the output file
	            OutputFileHandle, err4 := os.OpenFile(nameOfNewSourceCodeFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // write to file 
	            check(err4)                               // ... gets a file handle for output 
	            defer OutputFileHandle.Close()          // It’s idiomatic to defer a Close immediately after opening a file.
	        		_, err5 := fmt.Fprint(OutputFileHandle, "            Windows variant */\n")
	            	check(err5)
            } 

            //if lineRead == `(\s*\/\*\s*Unix\s*variant)` {
            if regexUnixStart.MatchString(lineRead) {
            	countOfUnixStart++
                // write the edited line to the output file
	            OutputFileHandle, err4 := os.OpenFile(nameOfNewSourceCodeFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // write to file 
	            check(err4)                               // ... gets a file handle for output 
	            defer OutputFileHandle.Close()          // It’s idiomatic to defer a Close immediately after opening a file.
	        		_, err5 := fmt.Fprint(OutputFileHandle, "            // /* Unix variant\n")
	            	check(err5)
            } 
            //if lineRead == `(\s*Unix\s*variant\s*\*\/)` {
            if regexUnixEnd.MatchString(lineRead) {
            	countOfUnixEnd++
                // write the edited line to the output file
	            OutputFileHandle, err4 := os.OpenFile(nameOfNewSourceCodeFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // write to file 
	            check(err4)                               // ... gets a file handle for output 
	            defer OutputFileHandle.Close()          // It’s idiomatic to defer a Close immediately after opening a file.
	        		_, err5 := fmt.Fprint(OutputFileHandle, "            // Unix variant */\n")
	            	check(err5)
            } 


        }
    if scanner.Err() != nil {
        fmt.Println("there was an error scanning")
    }
    fmt.Println("Refactoring complete, winS:", countOfWindowsStart, "winE:", countOfWindowsEnd, "UnixS:", countOfUnixStart, "UnixE:", countOfUnixEnd)
}

func check(e error) {   // create a func named check which takes one parameter "e" of type error 
    if e != nil {
        panic(e)        // use panic() to display error code 
    }
}