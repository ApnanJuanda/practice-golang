package main 

import "fmt"

func main(){
	var(
		a = 10
		b = 5
		c = a + b
	)

	var i = 10
	i += 10
	i++
	fmt.Println(i)
	fmt.Println(c)
}