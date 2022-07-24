package main

import "fmt"

func main(){
	result := sumAll(10,10,10)
	fmt.Println(result)

	slice := []int{1,2,3}
	total := sumAll(slice...)
	fmt.Println(total)
}

func sumAll(numbers ...int) int{
	total := 0
	for _, value := range numbers{
		total += value
	}

	return total
}