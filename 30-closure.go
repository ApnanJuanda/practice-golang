package main

import "fmt"

func main(){
	counter := 0

	increment := func ()  {
		counter := 2
		fmt.Println("call func increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}