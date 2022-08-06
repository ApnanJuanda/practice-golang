package main

import "fmt"

func main(){
	sayHello()

	for i := 0; i < 3; i++{
		fmt.Println("Perulangan ke-", i)
		sayHello()
	}
}

func sayHello(){
	fmt.Println("Hello Dunia")
}

