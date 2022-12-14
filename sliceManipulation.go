package main
import ("fmt")

func main() {
	var dimention uint64
	var area uint64 

	dimention = 100
	area = dimention * dimention
	iter := 0
  	index := 0

  pairs := []uint64{dimention, area}

  for iter < 23 {

  	iter++
  	dimention++

  	area = dimention*dimention

  pairs = append(pairs, dimention, area )  

  fmt.Printf("pairs = %v %v \n", pairs[index], pairs[index+1])
  	index = index+2

}

//fmt.Print(pairs)

//  fmt.Printf("pairs from last = %v %v \n", pairs[6], pairs[7])

  //sliceManipulation.go

}