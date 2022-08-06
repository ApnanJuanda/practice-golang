package main 

import "fmt"

func main(){
	firstName, lastName, _ := getFullName()
	fmt.Println(firstName, lastName)
}

func getFullName() (string, string, string){
	return "Apnan","Juanda","ignorre"
}