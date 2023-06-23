// π = 3 + 4/(2*3*4) - 4/(4*5*6) + 4/(6*7*8) - 4/(8*9*10) + 4/(10*11*12) ...

package main

import (
    "fmt" 
    "math/big" 
    "time"
)

var iterBig int
var precision int 
var piAs102chars string 

func main(){
    start := time.Now()
    //with_float64_types() // runs pretty quickly
    with_big_Float_types()  // can take forever 
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
                for iterBig < 50000000 { 
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
                
                stringOfSum := sumBig.Text('f', 100)
                fmt.Printf("\nSum as 102 chars is %s as a string type\n", stringOfSum)
                fmt.Println("                    1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
                fmt.Println("                            10        20        30        40        50        60        70        80        90       99")
                fmt.Printf("per big.Float type: %0.99f \n", sumBig)
                piAs102chars = "3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679"
                fmt.Printf("piAs102chars is     %s \n", piAs102chars)        

                    posInPi := 0
                    var piChar byte  
                    var ppis int 
                    for positionInString, char := range stringOfSum {
                        piChar = piAs102chars[posInPi] 
                            if char == rune(piChar) {
                                //fmt.Printf("we have agreement up to char pos: %d \n", positionInString)
                                ppis = positionInString 
                            } else {
                                break 
                            }
                            //3.14159265358979323846
                            //1 23456789012345678901 21 digits 
                        posInPi++
                        //fmt.Printf("\nPosiion in the string is: %d and the char is: %c\n", positionInString, char)
                    }        
                    fmt.Printf("\nWe have calculated pi correctly to %d digits\n", ppis)
}

/*
3.141592653589793238462643383279 252758886795791428853923554099998679517428007273779498771670848356052 big floats 
3.141592653589793238462643383279 5028841971693993751058209749445923078164062862089986280348253421170679 web
3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679 web 
123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012 102 chars 
         10       20        30        40        50        60         70       80        90       100
1 234567890123456789012345678901 got 31 digits in 1 hour and 26 min using this algorithm with one billion iters at 128 prec 
         10        20        30  
3.1415926535897932 16828765509696097524283686652779579162597656250000000000000000000000000000000000000 got only 17 digits at 500,000,000 and prec=64 but just 3min runtime
3.1415926535897932 384626433832795028841971693993751058209749445923078164062862089986280348253421170679
*/

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