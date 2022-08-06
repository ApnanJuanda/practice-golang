package main

import (
	"fmt"
	"practice-golang/helper"
)

func main(){
	helper.SayHello("Apnan")
	fmt.Println(helper.Application)
	//fmt.Println(helper.version) ini error karena attribute version diset tidak public memakai huruf kecil
}