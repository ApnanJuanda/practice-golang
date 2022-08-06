package main

import (
	"fmt"
)

func main(){
	name := "kauSiapa"

	switch name {
	case "Apnan":
		fmt.Println("Hello Apnan")

	case "Juanda":
		fmt.Println("Hello Juanda")
	
	default:
		fmt.Println("Hello " + name)
	}

	switch length := len(name); length > 4 {
	case true:
		fmt.Println("Namamu Kepanjangan")

	case false:
		fmt.Println("Nah ini baru pas")
	}
}