package main 

import "fmt"

func main(){
	var data interface{} = ups(1) 
	fmt.Println(data)
}

func ups(i int) interface{}{
	if i == 1{
		return true
	}else if i == 2{
		return 2
	}else{
		return "ups"
	}
}

