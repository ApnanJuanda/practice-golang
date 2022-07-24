package main

import "fmt"

func main(){
	name := "Anda Siapa"
	
	if name == "Apnan"{
		fmt.Println("Hello Apnan")
	}else if name == "Juanda"{
		fmt.Println("Hello Juanda")
	}else{
		fmt.Println("Hello " + name)
	}

	if length := len(name); length > 4{
		fmt.Println("Namamu kepanjangan")
	}else{
		fmt.Println("Nah ini sudah pas")
	}
}