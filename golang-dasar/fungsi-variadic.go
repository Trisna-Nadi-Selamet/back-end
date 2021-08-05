package main

import "fmt"

//fungsi variadic  jumlah parameter yang disisipkan tidak terbatas yang ditmpung sama variabelnya
func main(){
var rata_rata = hitung(10,2,5,8,7,8,2,5,6,12)
var pesan = fmt.Sprintf("Rata-rata : %.2f",rata_rata)
fmt.Println(pesan)
}

func hitung(angka ...int)float64{
	var total int = 0
	for _, angka:= range angka {
		total+=angka
		fmt.Println(angka)
	}
	var avg =float64(total) / float64(len(angka))
	return avg
}