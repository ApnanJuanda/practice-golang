package main

import "fmt"

type BlackList func(string) bool

func main(){
	blackList := func (name string)  bool{
		return name == "admin"
	}
	registerUser("Apnan", blackList)
}

func registerUser(name string, blackList BlackList){
	if blackList(name){
		fmt.Println("you are blocked", name)
	}else{
		fmt.Println("Welcome",name)
	}
}

func blackListAdmin(name string) bool{
	if name == "admin"{
		return true
	}else{
		return false
	}
}