package main

import "fmt"

func main(){
	//Mandatory using type data
	var name string;

	name = "Apnan Juanda";
	fmt.Println(name);

	//not mandatory if we directly initiate value of variable
	var occupation = "Java & Go Developer"
	fmt.Println(occupation);

	var age int8 = 24
	fmt.Println(age);

	//using := for variable
	country := "Indonesia"
	fmt.Println(country)

	//Multiple initiate variable
	var(
		firstJob = "Java Developer"
		secondJob = "Golang Developer"
	)

	fmt.Println(firstJob)
	fmt.Println(secondJob)
}