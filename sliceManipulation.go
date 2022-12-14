package main
import ("fmt")

	var dimention = 10
	var area = dimention * dimention
  	var pairs = []int{dimention, area}
  	var index = 0 



func main() {

  fmt.Printf("pairs from first index = %v %v \n", pairs[index], pairs[index+1])

  buildTable()

  readTable()

  fmt.Printf("pairs from final index = %v %v \n", pairs[index], pairs[index+1])
  fmt.Print(pairs)
}



func buildTable() {

	iter := 0

	  	for iter < 500 {

  			iter++
  			dimention++

  			area = dimention*dimention

  			pairs = append(pairs, dimention, area )  

  			index = index+2
		}
}


func readTable() {

	iter := 0
	index = 0

	for iter < 490 {

		iter++

			fmt.Println(pairs[index], pairs[index+1])

  		index = index+2
	}
}