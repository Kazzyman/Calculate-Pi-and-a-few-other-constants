// π = 3 + 4/(2*3*4) - 4/(4*5*6) + 4/(6*7*8) - 4/(8*9*10) + 4/(10*11*12) ...

package main

import (
    "fmt" 
    "math/big" 
    "time"
)

var iterBig int
var precision int 

func main(){
    start := time.Now()
    //with_float64_types() // runs pretty quickly
    with_big_Float_types()  // can take forever 
    // got 31 digits in 1 hour and 26 min using this algorithm with one billion iters at 128 prec
    t := time.Now()
    elapsed := t.Sub(start)
    TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
    fmt.Printf("Total run with SetPrec at: %d and iters of %d was %s \n ", precision, iterBig, TotalRun) 
}

func with_big_Float_types() {

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

        precision = 128 
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
                for iterBig < 5000000 { 
                    // Total run with SetPrec at:  128 and iters of 50000000  was 23.1879034s    :: 3.14159265358979323846264
                    // Total run with SetPrec at: 1024 and iters of 600000000 was 12m23.9554367s :: 3.14159265358979323846264338
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
                printResultStats(sumBig, iterBig, precision) 

}


func printResultStats(sumBig *big.Float, iterBig int, precision int) {
    var piAs86chars string 
    piAs86chars =                   "3.141592653589793238462643383279502884197169399375105820974944592307816406286208998628"
    stringOfSum := sumBig.Text('f', 86) // create a string version of a big result, 80 some odd chars in length 

    //fmt.Printf("\nSum as 86  chars is %s as string type\n", stringOfSum)

    fmt.Printf("pi as calculated is: %0.84f \n", sumBig)
    fmt.Println("                     1 234567890123456789012345678901234567890123456789012345678901234567890123456789012345")
    fmt.Println("                              10        20        30        40        50        60        70        80   85")
    fmt.Printf("pi from the web is:  %s \n", piAs86chars)        

    posInPi := 0
    var piChar byte  
    var ppis int 
    var stringVerOfCorrectDigits = []string{}  
    for positionInString, char := range stringOfSum {
        piChar = piAs86chars[posInPi] 
            if char == rune(piChar) {
                stringVerOfCorrectDigits = append(stringVerOfCorrectDigits, string(char))
                //fmt.Printf("we have agreement up to char pos: %d \n", positionInString)
                ppis = positionInString 
            } else {
                break // to print result and info below 
            }
        posInPi++
    }        
    fmt.Printf("\nWe have calculated pi correctly to %d digits using %d iters and Prec of %d \n", ppis, iterBig, precision)
    fmt.Print("The correctly calculated digits are: ")
        for _, oneChar := range stringVerOfCorrectDigits {
            fmt.Print(oneChar)
        }
    fmt.Print("\n")
}

func with_float64_types() {

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
//fmt.Printf("digitthree64 is: %0.8f and should be 4.00\n", digitthree64)

            var sum64 float64    
                sum64 = three64 + (four64 / (digitone64*digittwo64*digitthree64)) 
//fmt.Printf("firstSum 64: %0.8f \n", sum64) // agrees 
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
fmt.Printf("\nakantha Somayaji per float64 types is: %0.15f \n", sum64)
}
