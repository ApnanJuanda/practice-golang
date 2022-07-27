package main

import (
	"errors"
	"fmt"
)

func main(){
	hasil, err := pembagi(10,0)
	if err == nil{
		fmt.Println("Hasil=",hasil)
	}else{
		fmt.Println(err)
	}
}

func pembagi(nilai int, pembagi int) (int, error){
	if pembagi == 0{
		return 0, errors.New("Pembagi tidak boleh 0")
	}else{
		return nilai/pembagi, nil
	}
}