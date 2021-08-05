package main

import "fmt"

//menampilkan value fungsi dengan parameter dan memasukan pada fungsi yang ada
func main(){
    var mahasiswa = map[string]string{"Tinggi":"160 cm","Tinggi1":"165 cm"}
	var hasil_ke_1,hasil_ke_2 = tampilkan_hasil("Trisna","Syarif") 
	 
	fmt.Println(hasil_ke_1 ,mahasiswa["Tinggi"])
	fmt.Println(hasil_ke_2 ,mahasiswa["Tinggi1"])
    // fmt.Println("Syarif :",mahasiswa["Syarif"])
}
func tampilkan_hasil(x string, y string)(string,string){
	var a = "Trisna :"
	var b = "Syarif :"
	 return a,b
}