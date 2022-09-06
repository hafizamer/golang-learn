package main

// Import multiple packages
// You could use an alias like f "fmt"
import (
	"fmt"
)

// Create alias to long function names
var pl = fmt.Println

func manySum(y int, x int) int {
	return x + y
}

func total(arr ...int) int {
	sum := 0
	for _, x := range arr {
		sum += x

	}
	return sum

}


func main() {
	pl(manySum(5, 4))
	pl(total(1, 2, 3, 4, 5))

	{
	name := []string{"Uno", "Dos", "Tres"}
	num := []int{1, 2, 3}
	v := -1

	for _, i := range name {
		v++

		fmt.Println(i, num[v])

	}

}
}