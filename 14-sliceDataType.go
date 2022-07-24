package main

import "fmt"

func main(){
	//Slice itu potongan dari array dan ukuran si slice ini bisa berubah 
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[3:8]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//months[3] = "Ini apa"
	//fmt.Println(slice1)

	// slice1[3] = "Apaini"
	// fmt.Println(months)

	
	//Test Implements Append
	slice2 := months[2:4]
	slice3 := append(slice2, "bukanBulan")
	slice3[1] = "AprilBaru"
	fmt.Println(months)
	fmt.Println(slice2)
	fmt.Println(slice3)

	//Test Make Slice ]Program
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Apnan"
	newSlice[1] = "Juanda"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//Implement Copy Slice 
	copySlice := make([]string, 1, cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

}