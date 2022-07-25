package main

import "fmt"

type Customer struct{
	Name, Address string
	Age int
}
 
func main(){
	//first way 
	var apnan Customer
	apnan.Name = "Apnan"
	apnan.Address = "Bandung"
	apnan.Age = 24
	fmt.Println(apnan)

	//second way
	juanda := Customer{
		Name: "Juanda",
		Address:  "Bandung",
		Age: 24,
	}
	fmt.Println(juanda)

	//third way 
	test := Customer{"Test", "Test", 1}
	fmt.Println(test)
}