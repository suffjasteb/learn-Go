package main

import "fmt"

func main() {
	names := [...]string{
		"shafwan", // 0
		"junindra", // 1
		"putra", // 2
		"pratama",  // 3
		"budi", // 4
		"nugraha",  // 5
		"sidoarjo", // 6
		"last", // 7
	}

	// ambil slice data nya kira kira dari index 4 - 6 => berati yang di ambil 4 dan 5
	slice1 := names[4:6]
	fmt.Println(slice1)

	//, cara kerja slicing di Go adalah indeks awal (4) diambil, sedangkan indeks akhir (6) tidak diikutsertakan.
	// Jadi, slice yang diambil adalah elemen pada indeks 4 dan 5 saja.

	// index ke 4 "budi", tidak di ada karena yang terakhir tidak di ikutsertakan
	slice2 := names[:4]
	fmt.Println(slice2)

	// ambil index ke 3 sampai terakhir
	slice3 := names[3:]
	fmt.Println(slice3)


	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	days1 := days[5:] // sabtu minggu
	
	fmt.Println(days1)
	days1[0] = "sabtu baru"
	days1[1] = "minggu baru"

	fmt.Println(days1)


	days2 := append(days1, "Libur baru")
	fmt.Println(days2)

	// Panjang 2 kapasitas 5
	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Shafwan";
	newSlice[1] = "Junindra";
	// newSlice[2] = "Putra"; // error harusnya menggunakan append

	// cetak slice
	fmt.Println(newSlice)
	// cek panjang newslice
	fmt.Println(len(newSlice))
	// cek capasitas newslice
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "putra")
		// cetak slice
		fmt.Println(newSlice2)
		// cek panjang newslice
		fmt.Println(len(newSlice2))
		// cek capasitas newslice
		fmt.Println(cap(newSlice2))

		newSlice2[0] = "Budi"
		fmt.Println(newSlice)
		fmt.Println(newSlice2)

		// ambil semua array di variable days
		fromSlice := days[:]
		// buat slice yang tipe nya string
		toSlice := make([]string, len(fromSlice), cap(fromSlice));

		copy(toSlice, fromSlice)

		fmt.Println(fromSlice)
		fmt.Println(toSlice)

		 iniArray := [...]int{1,2,3,4,5} // ini Array
		 iniSlice := []int{1,2,3,4,5} // ini slice

		 fmt.Println(iniArray)
		 fmt.Println(iniSlice)

		 // pilih array atau slice ?
		 // Di Golang rata rata dibuat dalam bentuk Slice
}
		
