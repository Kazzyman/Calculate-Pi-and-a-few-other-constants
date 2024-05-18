package main

import "time"

// vars & consts for use in multiple localized sections of code =========================================================
// convenience globals:
var selection int

var usingBigFloats = false // a variable of type bool which is passed by many funcs to printResultStatsLong()

const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorYellow = "\033[33m"
const colorPurple = "\033[35m"
const colorCyan = "\033[36m"
const colorWhite = "\033[37m"

var iterationsForMonte16 int
var TotalIterations int
var four float64 // is initialized to 4 where needed
var π float64    // a var can be any character, as in this Pi symbol/character
var LinesPerSecond float64
var LinesPerIter float64
var iterInt64 int64     // to be used primarily in selections which require modulus calculations
var iterFloat64 float64 // to be used in selections which do not require modulus calculations
var Table_of_perfect_squares = []int{}
var t2 time.Time

// the following globals, are used in multiple funcs of case 18: calculate either square or cube root of any integer
// ------------------------------------------------------------------------------------------------------------------
var Tim_win float64             // Time Window
var sortedResults = []Results{} // sortedResults is an array of type Results as defined at the top of this file
var Table_of_perfect_Products = []int{}
var diffOfLarger int
var diffOfSmaller int
var perfectResult2 float64 // will contain the square root result if the workpiece was itself a perfect square
var perfectResult3 float64 // will contain the cube root result if the workpiece was itself a perfect cube
var precisionOfRoot int    // this being global means we do not need to pass it in to the read func
var workPiece int          // the square or cube of which we are to find a root
var skip_redoing_loop int

type Results struct { // define a new structure called Results with two fields; result, and pdiff
	result float64
	pdiff  float64
}

// ========== end of global var section ================================================================================
