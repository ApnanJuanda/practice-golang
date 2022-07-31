package main

import "fmt"

type Address struct{
	City, Province, Country string
}

func ChangeAddressToIndonesia(address *Address){
	address.Country = "Indonesia"
}

func main(){
	//country belum terubah
	/*var alamat = Address{"Subang", "Jawa Barat", "",}
	ChangeAddressToIndonesia(alamat) 
	fmt.Println(alamat)*/

	//country sudah berubah 
	/*alamat := Address{"Subang", "Jawa Barat", "",}
	ChangeAddressToIndonesia(&alamat)
	fmt.Println(alamat)*/

	alamat := Address{"Subang", "Jawa Barat", "",}
	alamatPointer := &alamat
	ChangeAddressToIndonesia(alamatPointer)
	fmt.Println(alamat)
}