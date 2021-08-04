package main

import "fmt"

func main(){

    var num1,num2 = bilangan_hasil(20,5)
	fmt.Println(tampilkan_pesan())
	fmt.Println(tampilkan_angka())
	fmt.Println(tampilkan_hasil(10,5))
	fmt.Println(num1)
	fmt.Println(num2)
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
//menampilkan pengembalian multi fungsi pada parameter pada golang
func bilangan_hasil(a int, b int)(int,int){ //syarat penulisan untuk pengembalian multi fungsi
	var c = a + b;
	var d = a * b;
	return c,d
}