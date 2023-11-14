package main

import ("fmt")
import("tugas6/library")

func main() {
	var x = library.Nama_mahasiswa{"ari", 18}
	fmt.Println("nama", x.Nama)
	fmt.Println("umur", x.Umur)

}
