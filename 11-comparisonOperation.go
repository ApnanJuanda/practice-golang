package main

import "fmt"

func main(){
	var(
		name1 string = "Apnan"
		name2 string = "Apnan"
	)

	var result bool = name1 == name2
	fmt.Println(result)

	var(
		number1 int8 = 100
		number2 int8 = 102
	)

	fmt.Println(number1 > number2);
	fmt.Println(number1 < number2);
	fmt.Println(number1 == number2);
	fmt.Println(number1 != number2);
	

}