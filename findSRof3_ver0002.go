package main
import ("fmt"; "math")

	var root = 100
	var PerfectSquare = root * root

  	var sliceOfPerfectSquares_and_their_roots = []int{PerfectSquare, root} // initialize slice with 100
  	var sliceOfP_diffs_smlrT3_to_largerOfProspectiveHit = []float32{99.9999}  // init with a large-ish p_diff_from_larger_hit
  	var sliceOfP_diffs_smlrT3_to_smlrOfProspectiveHit = []float32{88.888}     //   "            "      p_diff_from_smlr_hit

func main() {
  // index is now undefined here ... so :
  // fmt.Printf("Root and PerfectSquare from first indexing = %v and %v which should be 100 and 10000 \n", sliceOfPerfectSquares_and_their_roots[0+1], sliceOfPerfectSquares_and_their_roots[0])

  buildTableOfOnlyPerfectSquares()  // easy, no issues here

	var index = 0 // no-longer a global variable 

//test:= sliceOfPerfectSquares_and_their_roots[index]
//fmt.Println("test should print 10000 ", test)

	for index < 1000000 {
  			readTableOfPerfectSquares(index)  // pass-in the index // may need or want more than just one ???
  		index = index + 2 
  	}

  //lengthOfsliceOfPerfectSquares_and_their_roots := len(sliceOfPerfectSquares_and_their_roots)
  //capOfsliceOfPerfectSquares_and_their_roots := cap(sliceOfPerfectSquares_and_their_roots)
  //fmt.Printf("\nlength of slice is %v after bulid\n", lengthOfsliceOfPerfectSquares_and_their_roots)
  //fmt.Printf("cap of slice is %v after bulid\n", capOfsliceOfPerfectSquares_and_their_roots)


  //fmt.Println(sliceOfP_diffs_smlrT3_to_largerOfProspectiveHit)  // dump the entire slice
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
                                                                     // save a copy of smlr psqr 
			smallerPerfectSquareOnce := sliceOfPerfectSquares_and_their_roots[index2]  // do this just once per func call 

//lineCounter := 0 


	for iter < 1000 {

		iter++
  		index2 = index2 + 2 // index 
			largerPerfectSquare := sliceOfPerfectSquares_and_their_roots[index2]        // get next perfect square 
			//fmt.Println(largerPerfectSquare, lineCounter); lineCounter++
				if largerPerfectSquare == smallerPerfectSquareOnce*3 { // or, in the case that it is a miricle ...
					fmt.Println("It is a miricle, everyone was wrong! ") // praise god
					break // exit the for loop 
				}
			if largerPerfectSquare >= smallerPerfectSquareOnce*3 {     // if it is a candidate based on being just larger than 3x smlr psqr

					ProspectiveHitOnLargeSide := largerPerfectSquare
					p_diff_from_larger_hit:= determinDiff(ProspectiveHitOnLargeSide, smallerPerfectSquareOnce*3)  // ProspectiveHitOnLargeSide_Par, smallerPerfectSquareOnce_T3
					sliceOfP_diffs_smlrT3_to_largerOfProspectiveHit = append(sliceOfP_diffs_smlrT3_to_largerOfProspectiveHit, p_diff_from_larger_hit)

					Prospective_smaller_hit := sliceOfPerfectSquares_and_their_roots[index2-2]  // an int
					p_diff_from_smlr_hit := determinDiff(Prospective_smaller_hit, smallerPerfectSquareOnce*3)
					sliceOfP_diffs_smlrT3_to_smlrOfProspectiveHit = append(sliceOfP_diffs_smlrT3_to_smlrOfProspectiveHit, p_diff_from_smlr_hit) // build a slice of ... // may not need this as a slice ????

					// list everything that we have found so far 
					//fmt.Println("\n\nsmlr_ps, Prospective_smaller_hit, p_diff_from_smlr_hit, ProspectiveHitOnLargeSide, p_diff_from_larger_hit:: ", smallerPerfectSquareOnce, Prospective_smaller_hit, p_diff_from_smlr_hit, ProspectiveHitOnLargeSide, p_diff_from_larger_hit)

																																									// find a best_seed_root . between smaller and larger 
					//fmt.Println("   ... the roots of the best smaller pair are :: ", math.Sqrt(float64(smallerPerfectSquareOnce)), math.Sqrt(float64(Prospective_smaller_hit)))  // 
					sr_per_sm := (math.Sqrt(float64(Prospective_smaller_hit)) / math.Sqrt(float64(smallerPerfectSquareOnce)))
					//fmt.Println(sr_per_sm, "is sr per sm")
					if sr_per_sm > 1.73205 {
						if sr_per_sm < 1.73206 {
							fmt.Println("And, the square root of 3 per the smaller is ... ", sr_per_sm, "\n") 
						}
					}

					//fmt.Println("   ... the roots of the best larger pair are :: ", math.Sqrt(float64(smallerPerfectSquareOnce)), math.Sqrt(float64(ProspectiveHitOnLargeSide)))  // 
					sr_per_lg := (math.Sqrt(float64(ProspectiveHitOnLargeSide)) / math.Sqrt(float64(smallerPerfectSquareOnce)))
					//fmt.Println(sr_per_lg, "is sr per lg")
					if sr_per_lg > 1.73205 {
						if sr_per_lg < 1.73206{
							fmt.Println("And, the square root of 3 per the larger is ... ", sr_per_lg, "\n")  
						}	
					}

					break // exit the for loop
			}
	}
}

func determinDiff (ProspectiveHitOnLargeSide_Par int, smallerPerfectSquareOnce_T3 int) float32 {
	smallerPerfectSquareOnce3TF := float32(smallerPerfectSquareOnce_T3)
	ProspectiveHitOnLargeSideF := float32(ProspectiveHitOnLargeSide_Par) 

		the_absolute_diff := ProspectiveHitOnLargeSideF - smallerPerfectSquareOnce3TF  // this begins as significant, and just grows to become rediculously large, as it should
		// fmt.Println("the absolute_diff is", the_absolute_diff)

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