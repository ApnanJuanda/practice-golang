package main 

import "fmt"

func main(){
	counter := 1

	for counter <= 10{
		fmt.Println("perulangan ke-", counter)
		counter++
	}

	for number := 1; number <= 10; number++{
		fmt.Println("Perulangan ke-", number)
	}

	//making slice
	person := make([]string, 2, 5)
	person[0] = "Apnan"
	person[1] = "Juanda"
	for i := 0; i < len(person); i++{
		fmt.Println(person[i])
	}

	//making map
	book := make(map[string]string)
	book["author"] = "Apnan"
	book["title"] = "Java & Go Language"

	for key, value := range book{
		fmt.Println(key, "=", value)
	}

	//making second map
	laptop := make(map[string]string)
	laptop["merk"] = "Asus"
	laptop["price"] = "15000000"
	for key, value := range laptop{
		fmt.Println(key, "=", value)
	}
}