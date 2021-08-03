package main

import "fmt"

//task bilangan genap dan ganjil

func main(){
	var angka = 4
	if angka % 2 == 0   {
		fmt.Printf("bilangan genap %d" , angka)
	}else{
		fmt.Printf("bilangan ganjil %d" , angka)
	}
	
}