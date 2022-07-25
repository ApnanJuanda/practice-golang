package main

import "fmt"

func main(){
	runApp(true)
	fmt.Println("Test")
}

func logging(){
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int){
	//defer berfungsi untuk tetap menjalankan suatu function meskipus function sebelumnya mengalami error
	defer logging()
	fmt.Println("Run Application")
	result := 10/value
	fmt.Println("result: ", result)
}

func endApp(){
	fmt.Println("End App")
	message := recover()
	if message != nil{
		fmt.Println("Ada error dengan message: ", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(value bool){
	defer endApp()
	if value{
		//panic berfungsi untuk menghentikan program tetapi masih dapat mengeksekusi defer
		panic("APLIKASI ERROR")
	}
}

