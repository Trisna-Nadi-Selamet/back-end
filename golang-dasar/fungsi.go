package main

import "fmt"

func main(){
	fmt.Println(tampilkan_pesan())
	fmt.Println(tampilkan_angka())
	fmt.Println(tampilkan_hasil(10,5))
}
func tampilkan_pesan()string{ //menggunakan tipe variabel
	return "Berhasil diterima"  // menampilkan dengan cara mengembalikan nilai
	//fmt.Println("Berhasil diterima") //menampilkan fungsi dasar
}

//menampilkan type data integer
func tampilkan_angka()int{
	return 10
}

//menampilkan argumen parameter
func tampilkan_hasil(x int,y int)int{
	var z = x + y
	return z

}