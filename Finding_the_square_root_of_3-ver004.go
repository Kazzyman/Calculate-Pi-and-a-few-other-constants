package main
import ("fmt"; "math")

	var root = 100
	var PerfectSquare = root * root

  	var sliceOfPerfectSquares_and_their_roots = []int{PerfectSquare, root} // initialize slice with 100
  	var sliceOfP_diffs_smlrT3_to_largerOfProspectiveHit = []float32{99.9999}  // init with a large-ish p_diff_from_larger_hit
  	var sliceOfP_diffs_smlrT3_to_smlrOfProspectiveHit = []float32{88.888}     //   "            "      p_diff_from_smlr_hit

func main() {
  // fmt.Printf("Root and PerfectSquare from first indexing = %v and %v which should be 100 and 10000 \n", sliceOfPerfectSquares_and_their_roots[0+1], sliceOfPerfectSquares_and_their_roots[0])

		buildTableOfOnlyPerfectSquares()  // easy, no issues here

	var index = 0 // no-longer a global variable 
	for index < 1000000 {
  		readTableOfPerfectSquares(index)  // pass-in the index // may need or want more than just one ???
  		index = index + 2 
  	}
}

func buildTableOfOnlyPerfectSquares() {  // easy, no issues here
	iter := 0
	  	for iter < 20000000 {
  			iter++
  			root++ // began as a global variable  --- at 1,000 
  			PerfectSquare = root*root
  			sliceOfPerfectSquares_and_their_roots = append(sliceOfPerfectSquares_and_their_roots, PerfectSquare)  // update sliceOfPerfectSquares_and_their_roots with another PerfectSquare
  			sliceOfPerfectSquares_and_their_roots = append(sliceOfPerfectSquares_and_their_roots, root) 
		}
}

func readTableOfPerfectSquares(index2 int) {  // 
	iter := 0
			smallerPerfectSquareOnce := sliceOfPerfectSquares_and_their_roots[index2]  // save it, do this just-once per func call 

	for iter < 100000 {
		iter++
  		index2 = index2 + 2 // index 
			largerPerfectSquare := sliceOfPerfectSquares_and_their_roots[index2]        // get next perfect square 
				if largerPerfectSquare == smallerPerfectSquareOnce*3 { // ... in the case that it is a miricle ...
					fmt.Println("It is a miricle, everyone was wrong! ") // praise god
					break  // exit the for loop ?
				}
			if largerPerfectSquare >= smallerPerfectSquareOnce*3 {     // if largerPerfectSquare is a candidate based on it being just-a-bit larger than 3*smlr_PS

					ProspectiveHitOnLargeSide := largerPerfectSquare
					p_diff_from_larger_hit:= determinDiff(ProspectiveHitOnLargeSide, smallerPerfectSquareOnce*3)  // ProspectiveHitOnLargeSide_Par, smallerPerfectSquareOnce_T3
					sliceOfP_diffs_smlrT3_to_largerOfProspectiveHit = append(sliceOfP_diffs_smlrT3_to_largerOfProspectiveHit, p_diff_from_larger_hit)

					Prospective_smaller_hit := sliceOfPerfectSquares_and_their_roots[index2-2]  // an int
					p_diff_from_smlr_hit := determinDiff(Prospective_smaller_hit, smallerPerfectSquareOnce*3)
					sliceOfP_diffs_smlrT3_to_smlrOfProspectiveHit = append(sliceOfP_diffs_smlrT3_to_smlrOfProspectiveHit, p_diff_from_smlr_hit) // build a slice of ... // may not need this as a slice ????

					// list everything that we have found so far 
					//fmt.Println("\n\nsmlr_ps, Prospective_smaller_hit, p_diff_from_smlr_hit, ProspectiveHitOnLargeSide, p_diff_from_larger_hit:: ", smallerPerfectSquareOnce, Prospective_smaller_hit, p_diff_from_smlr_hit, ProspectiveHitOnLargeSide, p_diff_from_larger_hit)
// and now we cheat just to see if this is working ... 
					sr_per_sm := (math.Sqrt(float64(Prospective_smaller_hit)) / math.Sqrt(float64(smallerPerfectSquareOnce)))
					if sr_per_sm > 1.73205080752 {
						if sr_per_sm < 1.7320508077 {
							fmt.Println("\nPer the smaller ps, the \u221A3 is ... ", sr_per_sm) 
							fmt.Println("which is ten correct digits via first-principles of geometry alone\n... having found two very large perfect squares where one of them is")
							fmt.Println("very-nearly exactly three times larger, i.e., the sqrt of", Prospective_smaller_hit, "over the sqrt of", smallerPerfectSquareOnce)
							break  // exit the for loop ?
						}
					}
// should instead be deciding final answer based on closeness of fit, percent diff, to exactly 3x smaller perfect square. Not just checking to see if we have gotten this close to the known value of Sqrt of 3
					sr_per_lg := (math.Sqrt(float64(ProspectiveHitOnLargeSide)) / math.Sqrt(float64(smallerPerfectSquareOnce)))
					if sr_per_lg > 1.73205080752 {
						if sr_per_lg < 1.7320508077 {  // 1.73205080757 (so 1.73205080754 - 1.73205080759 )
							fmt.Println("\nPer the larger ps, the \u221A3 is ... ", sr_per_sm) 
							fmt.Println("which is ten correct digits via first-principles of geometry alone\n... having found two very large perfect squares where one of them is")
							fmt.Println("very-nearly exactly three times larger, i.e., the sqrt of", Prospective_smaller_hit, "over the sqrt of", smallerPerfectSquareOnce)
							break  // exit the for loop ?
						}	
					}
					break  // exit the for loop ?
			}
	}
}

func determinDiff (ProspectiveHit_Par int, smallerPerfectSquareOnce_T3 int) float32 {  
		the_absolute_diff := float32(ProspectiveHit_Par) - float32(smallerPerfectSquareOnce_T3) 
		return (the_absolute_diff/100) 
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