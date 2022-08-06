package main

import "fmt"

func main(){
	person := map[string]string{
		"name": "Apnan",
		"address": "Bandung",
	}
	person["occupation"] = "Java & Go Developer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["occupation"])

	//Implements map without initialize values
	book := make(map[string]string)
	book["author"] = "Apnan"
	book["title"] = "Java & Go"
	book["iniApa"] = "iniApa"

	fmt.Println(book)
	fmt.Println(len(book))
	delete(book, "iniApa")

	fmt.Println(book)
	fmt.Println(len(book))
 }