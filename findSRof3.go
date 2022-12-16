package main
import ("fmt")

	var root = 100
	var PerfectSquare = root * root
  	var sliceOfPerfectSquares = []int{PerfectSquare} // initialize slice with 100
  	var index = 0 

  	var sliceOfP_diffs_smlrT3_to_largerOfProspectiveHit = []float32{99.9999}  // init with a large-ish p_diff_from_larger_hit
  	var sliceOfP_diffs_smlrT3_to_smlrOfProspectiveHit = []float32{88.888}

func main() {

  fmt.Printf("PerfectSquare from first index = %v and should be 10000 \n", sliceOfPerfectSquares[index])

  buildTableOfOnlyPerfectSquares()

  readTableOfPerfectSquares()  // 

  lengthOfsliceOfPerfectSquares := len(sliceOfPerfectSquares)
  capOfsliceOfPerfectSquares := cap(sliceOfPerfectSquares)

  fmt.Printf("\nlength of slice is %v after bulid\n", lengthOfsliceOfPerfectSquares)
  fmt.Printf("cap of slice is %v after bulid\n", capOfsliceOfPerfectSquares)


  fmt.Println(sliceOfP_diffs_smlrT3_to_largerOfProspectiveHit)  // dump the entire slice
}


func buildTableOfOnlyPerfectSquares() {

	iter := 0

	  	for iter < 2000 {

  			iter++
  			root++ // began at 1,000 

  			PerfectSquare = root*root

  			sliceOfPerfectSquares = append(sliceOfPerfectSquares, PerfectSquare)  // update sliceOfPerfectSquares with another PerfectSquare
		}
}

func readTableOfPerfectSquares() {  // read and print each perfect square on its own line

	iter := 0
	index = 0 // a global variable 
                                                                      // save a copy of smlr psqr 
			smallerPerfectSquareOnce := sliceOfPerfectSquares[index]  // do this just once per func call 

	for iter < 1990 {

		iter++
  		index++
			largerPerfectSquare := sliceOfPerfectSquares[index]        // get next perfect square 
				if largerPerfectSquare == smallerPerfectSquareOnce*3 { // or, in the case that it is a miricle ...
					fmt.Println("It is a miricle, everyone was wrong! ") // praise god
					break // does this exit the for loop ??? 
				}
			if largerPerfectSquare >= smallerPerfectSquareOnce*3 {     // if it is a candidate based on being just larger than 3x smlr psqr
					ProspectiveHitOnLargeSide := largerPerfectSquare
					p_diff_from_larger_hit:= determinDiff(ProspectiveHitOnLargeSide, smallerPerfectSquareOnce*3)  // ProspectiveHitOnLargeSide_Par, smallerPerfectSquareOnce_T3

					Prospective_smaller_hit := sliceOfPerfectSquares[index-1]

					p_diff_from_smlr_hit := determinDiff(Prospective_smaller_hit, smallerPerfectSquareOnce*3)

					sliceOfP_diffs_smlrT3_to_smlrOfProspectiveHit = append(sliceOfP_diffs_smlrT3_to_smlrOfProspectiveHit, p_diff_from_smlr_hit) // build a slice of ... // may not need this as a slice ????



					//fmt.Println("p_diff, smlr_ps, Prospective_smaller_hit, p_hitOnlrgrSd, p_diff_from_smlr_hit :: ", p_diff, smallerPerfectSquareOnce, Prospective_smaller_hit, ProspectiveHitOnLargeSide, p_diff_from_smlr_hit)
					fmt.Println("smlr_ps, Prospective_smaller_hit, p_diff_from_smlr_hit, ProspectiveHitOnLargeSide, p_diff_from_larger_hit:: ", smallerPerfectSquareOnce, Prospective_smaller_hit, p_diff_from_smlr_hit, ProspectiveHitOnLargeSide, p_diff_from_larger_hit)

					sliceOfP_diffs_smlrT3_to_largerOfProspectiveHit = append(sliceOfP_diffs_smlrT3_to_largerOfProspectiveHit, p_diff_from_larger_hit)

					break 
			}
	}
}

func determinDiff (ProspectiveHitOnLargeSide_Par int, smallerPerfectSquareOnce_T3 int) float32 {
	smallerPerfectSquareOnce3TF := float32(smallerPerfectSquareOnce_T3)
	ProspectiveHitOnLargeSideF := float32(ProspectiveHitOnLargeSide_Par) 

		the_absolute_diff := ProspectiveHitOnLargeSideF - smallerPerfectSquareOnce3TF  // this begins as significant, and just grows to become rediculously large, as it should
		fmt.Println("the absolute_diff is", the_absolute_diff)

		return (the_absolute_diff/100)  // attempting to return the diff as a percent, and now we have it :) 
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