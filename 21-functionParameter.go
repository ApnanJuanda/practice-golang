package main

import "fmt"

func main(){
	sayHelloWithParam("Apnan", "Juanda")
}

func sayHelloWithParam(firstName string, lastName string){
	fmt.Println("Hello",firstName,lastName)
}