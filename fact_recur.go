package main
import ("fmt")

func factorial_recursion(x float64) (y float64) {
     if x > 0 {
                    fmt.Println("x is", x, "y is", y, "from top of if")
          y = x * factorial_recursion(x-1) // calling a func within a func (recursion) 
                    fmt.Println("x is", x, "y is", y, "from after recurse")
     } else {
          y = 5
       }
   return
}

func main() {
     fmt.Println(factorial_recursion(4))
}
