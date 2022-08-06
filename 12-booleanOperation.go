package main

import "fmt"

func main(){
	var(
		nilaiAkhir int8 = 90
		presensi int8 = 40
	)

	var(
		lulusNilaiAkhir = nilaiAkhir >= 89
		lulusPresensi = presensi >= 100
	)

	var lulus bool = lulusNilaiAkhir && lulusPresensi;
	if(lulus){
		fmt.Println("Selamat Anda Lulus");
	}else{
		fmt.Println("Belajar yang benerrr!!!")
	}
}