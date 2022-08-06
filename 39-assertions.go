package main

import "fmt"

func random() interface{}{
	return true
}

func main(){
	result := random()
	switch value := result.(type) {
	case int:
		fmt.Println("Int", value)
	case string:
		fmt.Println("String", value)
	default:
		fmt.Println("unknown")
	}

}