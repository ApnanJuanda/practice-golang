package main 

import "fmt"

func main(){
	result := factorialLoop(8)
	fmt.Println(result)

	result2 := factorialRecursive(8)
	fmt.Println(result2)
}

func factorialLoop(number int) int{
	result := 1
	for i := number; i > 0; i--{
		result *= i
	}
	return result
}

func factorialRecursive(number int) int{
	if number == 1{
		return 1
	}else{
		return number * factorialRecursive(number-1)
	}
}