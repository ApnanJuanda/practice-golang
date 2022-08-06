package main 

import "fmt"

type Customer struct{
	Name, Address string 
	age int
}

func (customer Customer) sayHello(name string){
	fmt.Println("Hello",name, "my name is", customer.Name)
}

func main(){
	apnan := Customer{
		Name: "Apnan",
	}	
	apnan.sayHello("test")
}