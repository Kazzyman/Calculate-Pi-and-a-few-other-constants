package main
import "fmt"

func main() { 

	var faster float64
	var slower float64
	var difference float64
	var changeInPercent float64

// prompt
fmt.Println("enter faster, slower")

// input
fmt.Scanf("%f %f", &faster, &slower)

	difference = slower - faster

	changeInPercent = difference / slower * 100

fmt.Printf("\nfaster is %0.2f percent faster than slower\n", changeInPercent)
}