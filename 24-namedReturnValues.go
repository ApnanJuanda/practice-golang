package main

import "fmt"

func main(){
	firstName, lastName := getCompleteName()
	fmt.Println(firstName)
	fmt.Println(lastName)
}

func getCompleteName() (firstName, lastName string){
	firstName = "Apnan"
	lastName = "Juanda"
	return
}