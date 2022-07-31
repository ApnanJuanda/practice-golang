package main 

import "fmt"

type Address struct {
	City, Province, Country string
}

func main(){
	//Pengenalan tentang pointer, jika dicopy ke address2 kemudian ada yang dirubah maka address1 tidak berubah (pass by value)
	/*address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)*/

	//jika address2 ada yang dirubah maka address1 ikut berubah juga (pass by reference)
	/*var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)*/

	//ketika address2 reference ke address 1, kemudian address2 dibuat ulang isinya secara utuh maka address1 tidak berubah
	/*address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 
	address2 = &Address{"Malang", "Jawa Timur", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)*/

	//ketika address2 reference ke address1, kemudian address2 dibuat ulang isinya, agar address1 juga ikut berubah isinya maka
	//digunakanlah tanda *
	address1 := Address{"Subang", "Jawa Timur", "Indonesia"}
	address2 := &address1
	address3 := &address1

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

}