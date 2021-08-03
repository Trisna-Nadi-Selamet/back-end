package main

import "fmt"

//task mengambil vlaue slice menggunakan for
func main(){
	
    var buah = [] string{"Mangga","Apel","Jambu","Angggur"}
	buah = append(buah,"pepaya")
	fmt.Println("jumlah elment ",len(buah))
	fmt.Println("isi elment :",buah)
	for key:=range buah{
		fmt.Println(" elment ke :",key,buah[key])
	}
	
}