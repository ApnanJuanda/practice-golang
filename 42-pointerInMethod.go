package main 

import "fmt"

type Man struct{
	Name string
}

func (man *Man)  Married(){
	man.Name = "Mr. " + man.Name
	fmt.Println("Di Method",man)
}

func main(){
	apnan := Man{"Apnan"}
	apnan.Married()
	fmt.Println(apnan)
}