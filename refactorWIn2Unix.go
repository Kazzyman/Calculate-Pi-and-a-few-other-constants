package main

import (
	"os"
    "fmt"       // Used for printing etc.
    //"regexp"
    "bufio"     // used to count the lines in this file
    //"log"          
    "strings"    
)

var nameOfOldSourceCodeFile = "testFileWin.go"  
var nameOfNewSourceCodeFile = "targetFileUnix.go"

func main() {

    fileHandle, err1 := os.Open(nameOfOldSourceCodeFile)       // open the source code file and get a handle to it 
    if err1 != nil {
        fmt.Println("error opening source per rick")
    }
    defer fileHandle.Close()
        scanner := bufio.NewScanner(fileHandle)
        for scanner.Scan() {
            lineRead := strings.TrimSpace(scanner.Text())   // read a line from the file into lineRead


            // write the read line unchanged to the output file if any of these four strings are not found in the read line from the input file
            if lineRead != `// /* Windows variant` && lineRead != `// Windows variant */` && lineRead != `/* Unix variant` && lineRead != `Unix variant */` {
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
            if lineRead == `// /* Windows variant` {
                // write the edited line to the output file
	            OutputFileHandle, err4 := os.OpenFile(nameOfNewSourceCodeFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // write to file 
	            check(err4)                               // ... gets a file handle for output 
	            defer OutputFileHandle.Close()          // It’s idiomatic to defer a Close immediately after opening a file.
	        		_, err5 := fmt.Fprint(OutputFileHandle, "/* Windows variant\n")
	            	check(err5)
            } 
            if lineRead == `// Windows variant */` {
                // write the edited line to the output file
	            OutputFileHandle, err4 := os.OpenFile(nameOfNewSourceCodeFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // write to file 
	            check(err4)                               // ... gets a file handle for output 
	            defer OutputFileHandle.Close()          // It’s idiomatic to defer a Close immediately after opening a file.
	        		_, err5 := fmt.Fprint(OutputFileHandle, "Windows variant */\n")
	            	check(err5)
            } 

            if lineRead == `/* Unix variant` {
                // write the edited line to the output file
	            OutputFileHandle, err4 := os.OpenFile(nameOfNewSourceCodeFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // write to file 
	            check(err4)                               // ... gets a file handle for output 
	            defer OutputFileHandle.Close()          // It’s idiomatic to defer a Close immediately after opening a file.
	        		_, err5 := fmt.Fprint(OutputFileHandle, "// /* Unix variant\n")
	            	check(err5)
            } 
            if lineRead == `Unix variant */` {
                // write the edited line to the output file
	            OutputFileHandle, err4 := os.OpenFile(nameOfNewSourceCodeFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // write to file 
	            check(err4)                               // ... gets a file handle for output 
	            defer OutputFileHandle.Close()          // It’s idiomatic to defer a Close immediately after opening a file.
	        		_, err5 := fmt.Fprint(OutputFileHandle, "// Unix variant */\n")
	            	check(err5)
            } 


        }
    if scanner.Err() != nil {
        fmt.Println("there was an error scanning")
    }
}

func check(e error) {   // create a func named check which takes one parameter "e" of type error 
    if e != nil {
        panic(e)        // use panic() to display error code 
    }
}