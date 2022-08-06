package main

import "fmt"

//using type for func
type Filter func(string) string

func main(){
	sayHelloWithFilter("cek", spamFilter)
}

func sayHelloWithFilter(name string, filter Filter){
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string{
	if name == "cek"{
		return "..."
	}else{
		return name
	}
}