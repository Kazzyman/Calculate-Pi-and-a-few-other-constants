package main
import ("fmt"; "math/big")

func main() {       // top-level program logic flow -- explore SEVENTEEN ways to calculate Pi, plus THREE other constants
	biguns() // 
	FloatToBigInt(93456799999999999) // pass in an 18 digit literal
	//                  12345678901
	//            12345678912345678 an eighteen digit number!
} 
func biguns() {
        resultBF := big.NewFloat(0)
        firstnumBF := big.NewFloat(1234567891234567.12345)
        //firstnumBF := big.NewFloat(123.89)
        secondnumBF := big.NewFloat(999999999999999999999.00933)
        //secondnumBF := big.NewFloat(1.0)
    fmt.Printf("1 -- enter alternate firstnumBF, current is: %0.6f formatted 0.6f\n", firstnumBF)
    fmt.Println("2a- but should be:                           1234567891234567.12345")
    //  prints 1234567891234566955582880920792727552.0000000 for firstnumBF
    // ^^^^^^^ fails to print the .12345 part of the firstnumBF
    //fmt.Scanf("%d", &firstnumBF)
    //fmt.Scanf("%d", &firstnumBF)

        //cval := new(big.Float) // prepare a place for firstnumBF as a big.Float
        //cval.SetFloat64(1234567891234567) // convert passed val (small float64) to big float 64 and store in bigval (a big.Float)
        /* not working
    	if firstnumBF != cval { // 
    		fmt.Printf("2b-The alternate number you entered was: %d DecimallyFormatted, or %0.20f at 0.20f \n", firstnumBF)
		}
		*/
	// calculate a very large product between two big.Floats, yeilding resultBF
	    resultBF.Mul(secondnumBF, firstnumBF) // calculat product of two big.Floats
    fmt.Printf("\n3-this is the product formatted 0.7f: %0.7f \n\n", resultBF)
    // ^^^^^^ prints 123890000000000007864320.0000000 with firstnumBF set to 123.89
    // ^^^^^^ prints 1234567891234566955582880920792727552.0000000 with firstnumBF set to 1234567891234567.12345
    // 1234567891234567.0000000 when firstnumBF is multiplied by 1.0 --- so these floats are not working 

    // create a 19 digit number from a string 
    var smallnum, _ = new(big.Int).SetString("2188824200011112223", 10)
    //                                        1234567890123456789 19 digits is not small
	num := smallnum.Int64()	// transfer smallnum from above to num as Int64 ???
	fmt.Printf("4-this is num from new(big.Int).SetString(2188824200011112223, 10)  and then   := smallnum.Uint64() \n %d", num)
	// ^^^^^^ prints perfectly 

    var largernum, _ = new(big.Float).SetString("2188824123456789123456789.99967")
	lnum, _ := largernum.Float64()		
	fmt.Printf("\n\n5-this is lnum formatted0.8f: %0.8f \nfrom new(big.Float).SetString(2188824123456789123456789.99967)  and  := largernum.Float64()\n", lnum)
	// ^^^^^^ prints                                            2188824123456789019623424.00000000
	//  errors at:                                                              x         xxxxx

}

// float64 -> big.Float -> big.Int 
	func FloatToBigInt(val float64) { // val is an 18 digit float64 number
    bigval := new(big.Float) // prepare a place for val as a big.Float
    bigval.SetFloat64(val) // convert passed val (small float64) to big float 64 and store in bigval (a big.Float)
    // Set precision if required.
    // bigval.SetPrec(64)

    // set coin to an 18 digit big.Float 
    coin := new(big.Float) // coin is actually a Float !!!
    coin.SetInt(big.NewInt(945678123456345819)) // and, as a Float, it contains methods to convert (SetInt) and create a (big.NewInt) from a leteral 
    //                     123456789012345678 an eighteen digit number!

    // calculate a very very large float (a 35 digit number!!)
    bigval.Mul(bigval, coin) // Mul 2 big.Floats and store product in the first of the two vars, as a Float64 (via new(big.Float) and SetFloat64 above)

    // big.Float -> big.Int 
    result := new(big.Int) // prepare a new place (result) for the product (bigval) as a big int 
    // vvv bigval is definately a big.Float !!!!
    bigval.Int(result)    // store converted number in result, result now holds a big.Int version of the big.Float product: bigval
                          // who gets what? 

    // print the big.Int "result"
    //fmt.Printf("\n\n6a-Decimally formatted result(bigInt) prints as: %d and formatted as a float it prints as: %0.9f \n\n", result, result)
    //  ^^^^^ prints perfectly as %d but fails as %0.9f (as should be expected)
    // prints 88380051248235013904008289657552896 
    //        12345678901234567890123456789012345 a 35 digit number!!
    //       8.83800512482 x 10e34 from HP calculator 
    fmt.Printf("\n\n6b-Dmly formatted as d result(bigInt) prints as: %d \n\n", result)
    fmt.Println("7-and as per Println the big.Float prints as:   ", result) // prints perfectly using Println (as should be expected)

    // convert a *big.Int to big float64, then print it formatted as %0.9f
	bigRval := new(big.Float).SetInt(result) // SetInt converts result to a Int from its last form as a Float 
    //fmt.Printf("\n\n-8 the new Int prints formatted-desimal as: %d and prints formatted as float: %0.9f \n\n", bigRval, bigRval)
    // ... so decimally it fails, but works nicely formatted as %0.9f 
    fmt.Printf("\n8- the new Int prints nicely formatted as float: %0.9f \n\n", bigRval)

}
/*
Int    signed integers
Rat    rational numbers
Float  floating-point numbers

var x Int        // &x is an *Int of value 0
var r = &Rat{}   // r is a *Rat of value 0
y := new(Float)  // y is a *Float of value 0

var z1 big.Int
z1.SetUint64(123)                 // z1 := 123
y := new(Float).SetInt(z1)       // y := 123.0

*/
