package main

import "fmt"

//tipe data golang positife ,negative , desimal,boolean:
func main(){
	//positif uint8 angka dari 0 - 255
	var number_positif uint8 = 255
	//positif uint16 angka dari 0 - 65535  
    var number_positif2 uint16 = 65535
	//positif uint32 angka dari 0 - 4294967293  
    var number_positif3 uint32 = 4294967293
    //positif uint64 angka dari 0 - 18446744073709557675  
    var number_positif4 uint64 = 4294967293

	//negative angka -128 - 127
    var number_negaif int8 =-127 

	//desimal penggunaan menggunakan titik(.)
	var number_decimal = 2.5

	fmt.Println(number_positif,number_positif2,number_positif3,number_positif4)
    fmt.Println(number_negaif)
	fmt.Println(number_decimal)
}