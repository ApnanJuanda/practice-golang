package main

import "fmt"

func main(){
	person := NewMap("Apnan")
	fmt.Println(person)
}

func NewMap(name string) map[string]string{
	if(name == ""){
		return nil
	}else{
		return map[string]string{
			"Name" : "Apnan",
		}
	}
}