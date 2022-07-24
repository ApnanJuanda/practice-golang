package main

import "fmt"

func main(){
	goodBye := getGoodbye
	fmt.Println(goodBye("Apnan"))
}

func getGoodbye(name string) string{
	return "Goodbye " + name
}