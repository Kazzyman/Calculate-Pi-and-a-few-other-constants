package main
import ("fmt")

	var root = 10
	var PerfectSquare = root * root
  	var sliceOfPerfectSquares = []int{PerfectSquare} // initialize slice with 100
  	var index = 0 

  	var sliceOfP_diffs = []float32{99.9999}  // init with a large p_diff 



func main() {

  fmt.Printf("PerfectSquare from first index = %v and should be 100 \n", sliceOfPerfectSquares[index])

  buildTableOfOnlyPerfectSquares()

  readTableOfPerfectSquares()  // 

  lengthOfsliceOfPerfectSquares := len(sliceOfPerfectSquares )
  capOfsliceOfPerfectSquares := cap(sliceOfPerfectSquares)

  fmt.Printf("length of slice is %v after bulid\n", lengthOfsliceOfPerfectSquares)
  fmt.Printf("cap of slice is %v after bulid\n", capOfsliceOfPerfectSquares)


  fmt.Println(sliceOfP_diffs)  // dump the entire slice
}



func buildTableOfOnlyPerfectSquares() {

	iter := 0

	  	for iter < 500 {

  			iter++
  			root++

  			PerfectSquare = root*root

  			sliceOfPerfectSquares = append(sliceOfPerfectSquares, PerfectSquare )  

		}
}


func readTableOfPerfectSquares() {  // read and print each perfect square on its own line

	iter := 0
	index = 0

			smallerPerfectSquareOnce := sliceOfPerfectSquares[index]  // do this just once

	for iter < 490 {

		iter++
  		index++
			largerPerfectSquare := sliceOfPerfectSquares[index]
			if largerPerfectSquare >= smallerPerfectSquareOnce*3 {
				if largerPerfectSquare == smallerPerfectSquareOnce*3 {
					fmt.Println("It is a miricle, everyone was wrong! ")
				}
			}
					ProspectiveHitOnLargeSide := largerPerfectSquare
					p_diff := determinDiff(ProspectiveHitOnLargeSide, smallerPerfectSquareOnce*3)  // a1, a2

					fmt.Println("print this ", p_diff, smallerPerfectSquareOnce, ProspectiveHitOnLargeSide) // this does execute 

					sliceOfP_diffs = append(sliceOfP_diffs, p_diff)
			
	}
}

func determinDiff (a1 int, a2 int) float32 {
	smallerPerfectSquareOnce3TF := float32(a2)
	ProspectiveHitOnLargeSideF := float32(a1) 
	return (smallerPerfectSquareOnce3TF/ProspectiveHitOnLargeSideF)  // (300/growing_squares) ... if this is close to 1.000 then it is a hit
}
/*
1: 
read an initial area (indexed by last initial area plus 1), understood to be the smaller // done as (READ-1)
read the next area (initial area + 1), which will be the larger  // done
if the larger is == or > than 3 * the smaller, and
	if 
		those two values are equal
		print "it is a miracle, everyone else was wrong"  // done up to here 
	else 
		remember that last read larger as having been larger than 3*sml 
		compare the larger by percentage diff (p-diff) with  3 * the smaller
		record that p-diff in a table of prospective smallest difference degrees along with the prospective pair of areas (smaller, and larger) ... so we are to build a slice with triplets: ( smaller, larger, [p-diff between smaller*3 and larger] )

		now go get the prior "larger" square and determing the percent diff from orriginal root*3 and save it too as the next tripplet
		we are building a table of percent deviations looking for an exceedingly small deviation

loop returns to 1: 

func buildPDtable() {
		iter := 0
		index = 0
	for init < 500,000 {
		iter++
		smaller := pairs[index]
		larger := pairs[index+3]

	}
}
*/