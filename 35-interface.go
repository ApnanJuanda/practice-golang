package main 

import "fmt"

type HasName interface{
	GetName() string
}

type Person struct{
	Name string
}

func (person Person) GetName() string{
	return person.Name
}

func SayHello(hasName HasName){
	fmt.Println("Hello", hasName.GetName())
}

type Animal struct{
	Name string
}

//implement interface
func (animal Animal) GetName() string{
	return animal.Name
}

func main(){
	cat := Animal{
		Name: "Push",
	}

	SayHello(cat)

}