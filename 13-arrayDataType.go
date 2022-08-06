package main

import "fmt"

func main(){
	//firstWay
	var names [3]string
	names[0] = "Apnan"
	names[1] = "Juanda"
	names[2] = "Wanda"

	fmt.Println(names);

	//secondWay
	var values = [3]int8{
		8, 9, 10,
	}

	fmt.Println(values)

	//get length array 
	var lenth [10]int8
	fmt.Println(len(lenth))
}