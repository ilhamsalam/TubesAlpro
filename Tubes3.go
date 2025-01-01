package main

import (
	"fmt"
)

const (
	maxSpareParts = 100
	maxCustomers  = 100
	maxTransactions = 100
)

type SparePart struct {
	Name  string
	Price int
}

type Customer struct {
	Name string
}

type Transaction struct {
	Date       string
	CustomerID int
	SparePartID int
	Quantity   int
}

var (
	spareParts   [maxSpareParts]SparePart
	customers    [maxCustomers]Customer
	transactions [maxTransactions]Transaction

	numSpareParts   int
	numCustomers    int
	numTransactions int
)

func main() {
	for {
		fmt.Println("======================")
		fmt.Println("Aplikasi Service Motor")
		fmt.Println("~~~~~~Kelompok 4~~~~~~")
		fmt.Println("======================")
		fmt.Println("1. Data")
		fmt.Println("2. Histori")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			menuData()
		case 2:
			menuHistori()
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuData() {
	for {
		fmt.Println("======================")
		fmt.Println("Aplikasi Service Motor")
		fmt.Println("~~~~~~~Menu Data~~~~~~")
		fmt.Println("======================")
		fmt.Println("1. Tambah data")
		fmt.Println("2. Edit data")
		fmt.Println("3. Hapus data")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih menu: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			menuTambahData()
		case 2:
			menuEditData()
		case 3:
			menuHapusData()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuTambahData() {
	for {
		fmt.Println("======================")
		fmt.Println("Aplikasi Service Motor")
		fmt.Println("~~~Menu Tambah data~~~")
		fmt.Println("======================")
		fmt.Println("1. Tambah Data spare-part")
		fmt.Println("2. Tambah Data pelanggan")
		fmt.Println("3. Tambah Data transaksi")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih menu: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			tambahSparePart()
		case 2:
			tambahPelanggan()
		case 3:
			tambahTransaksi()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahSparePart() {
	if numSpareParts >= maxSpareParts {
		fmt.Println("Data spare-part sudah penuh.")
		return
	}
	var name string
	var price int

	fmt.Print("Masukkan nama spare-part: ")
	fmt.Scanln(&name)
	fmt.Print("Masukkan harga spare-part: ")
	fmt.Scanln(&price)

	spareParts[numSpareParts] = SparePart{Name: name, Price: price}
	numSpareParts++
	fmt.Println("Data spare-part berhasil ditambahkan.")
}

func tambahPelanggan() {
	if numCustomers >= maxCustomers {
		fmt.Println("Data pelanggan sudah penuh.")
		return
	}
	var name string

	fmt.Print("Masukkan nama pelanggan: ")
	fmt.Scanln(&name)

	customers[numCustomers] = Customer{Name: name}
	numCustomers++
	fmt.Println("Data pelanggan berhasil ditambahkan.")
}

func tambahTransaksi() {
	if numTransactions >= maxTransactions {
		fmt.Println("Data transaksi sudah penuh.")
		return
	}

	var date string
	var customerID, sparePartID, quantity int

	fmt.Print("Masukkan tanggal transaksi (YYYY-MM-DD): ")
	fmt.Scanln(&date)

	fmt.Print("Masukkan ID pelanggan: ")
	fmt.Scanln(&customerID)
	if customerID < 0 || customerID >= numCustomers {
		fmt.Println("ID pelanggan tidak valid.")
		return
	}

	fmt.Print("Masukkan ID spare-part: ")
	fmt.Scanln(&sparePartID)
	if sparePartID < 0 || sparePartID >= numSpareParts {
		fmt.Println("ID spare-part tidak valid.")
		return
	}

	fmt.Print("Masukkan jumlah spare-part: ")
	fmt.Scanln(&quantity)

	transactions[numTransactions] = Transaction{Date: date, CustomerID: customerID, SparePartID: sparePartID, Quantity: quantity}
	numTransactions++
	fmt.Println("Data transaksi berhasil ditambahkan.")
}

func menuEditData() {
	for {
		fmt.Println("======================")
		fmt.Println("Aplikasi Service Motor")
		fmt.Println("~~~~Menu Edit data~~~~")
		fmt.Println("======================")
		fmt.Println("1. Edit Data spare-part")
		fmt.Println("2. Edit Data pelanggan")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih menu: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			editSparePart()
		case 2:
			editPelanggan()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func editSparePart() {
	var id int
	fmt.Print("Masukkan ID spare-part yang ingin diubah: ")
	fmt.Scanln(&id)

	if id < 0 || id >= numSpareParts {
		fmt.Println("ID spare-part tidak valid.")
		return
	}

	var name string
	var price int
	fmt.Print("Masukkan nama spare-part baru: ")
	fmt.Scanln(&name)
	fmt.Print("Masukkan harga spare-part baru: ")
	fmt.Scanln(&price)

	spareParts[id] = SparePart{Name: name, Price: price}
	fmt.Println("Data spare-part berhasil diubah.")
}

func editPelanggan() {
	var id int
	fmt.Print("Masukkan ID pelanggan yang ingin diubah: ")
	fmt.Scanln(&id)

	if id < 0 || id >= numCustomers {
		fmt.Println("ID pelanggan tidak valid.")
		return
	}

	var name string
	fmt.Print("Masukkan nama pelanggan baru: ")
	fmt.Scanln(&name)

	customers[id] = Customer{Name: name}
	fmt.Println("Data pelanggan berhasil diubah.")
}

func menuHapusData() {
	for {
		fmt.Println("======================")
		fmt.Println("Aplikasi Service Motor")
		fmt.Println("~~~~Menu Hapus data~~~")
		fmt.Println("======================")
		fmt.Println("1. Hapus Data spare-part")
		fmt.Println("2. Hapus Data pelanggan")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih menu: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			hapusSparePart()
		case 2:
			hapusPelanggan()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func hapusSparePart() {
	var id int
	fmt.Print("Masukkan ID spare-part yang ingin dihapus: ")
	fmt.Scanln(&id)

	if id < 0 || id >= numSpareParts {
		fmt.Println("ID spare-part tidak valid.")
		return
	}

	for i := id; i < numSpareParts-1; i++ {
		spareParts[i] = spareParts[i+1]
	}
	numSpareParts--
	fmt.Println("Data spare-part berhasil dihapus.")
}

func hapusPelanggan() {
	var id int
	fmt.Print("Masukkan ID pelanggan yang ingin dihapus: ")
	fmt.Scanln(&id)

	if id < 0 || id >= numCustomers {
		fmt.Println("ID pelanggan tidak valid.")
		return
	}

	for i := id; i < numCustomers-1; i++ {
		customers[i] = customers[i+1]
	}
	numCustomers--
	fmt.Println("Data pelanggan berhasil dihapus.")
}

func menuHistori() {
	for {
		fmt.Println("======================")
		fmt.Println("Aplikasi Service Motor")
		fmt.Println("~~~~~Menu Histori~~~~~")
		fmt.Println("======================")
		fmt.Println("1. Histori Service motor")
		fmt.Println("2. Favorit Service motor")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih menu: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			historiServiceMotor()
		case 2:
			favoritServiceMotor()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func historiServiceMotor() {
	for {
		fmt.Println("======================")
		fmt.Println("Histori Service Motor")
		fmt.Println("======================")
		fmt.Println("1. Berdasarkan Tanggal")
		fmt.Println("2. Berdasarkan Spare-part")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih menu: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			historiBerdasarkanTanggal()
		case 2:
			historiBerdasarkanSparePart()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func historiBerdasarkanTanggal() {
	var date string
	fmt.Print("Masukkan tanggal (YYYY-MM-DD): ")
	fmt.Scanln(&date)

	found := false
	fmt.Println("Histori transaksi pada tanggal", date, ":")
	for _, transaction := range transactions[:numTransactions] {
		if transaction.Date == date {
			found = true
			customer := customers[transaction.CustomerID]
			sparePart := spareParts[transaction.SparePartID]
			fmt.Printf("Pelanggan: %s, Spare-part: %s, Jumlah: %d\n", customer.Name, sparePart.Name, transaction.Quantity)
		}
	}

	if !found {
		fmt.Println("Tidak ada transaksi pada tanggal tersebut.")
	}
}

func historiBerdasarkanSparePart() {
	var sparePartID int
	fmt.Print("Masukkan ID spare-part: ")
	fmt.Scanln(&sparePartID)

	if sparePartID < 0 || sparePartID >= numSpareParts {
		fmt.Println("ID spare-part tidak valid.")
		return
	}

	found := false
	fmt.Println("Histori transaksi menggunakan spare-part", spareParts[sparePartID].Name, ":")
	for _, transaction := range transactions[:numTransactions] {
		if transaction.SparePartID == sparePartID {
			found = true
			customer := customers[transaction.CustomerID]
			totalPrice := transaction.Quantity * spareParts[transaction.SparePartID].Price
			fmt.Printf("Pelanggan: %s, Jumlah: %d, Total Harga: %d\n", customer.Name, transaction.Quantity, totalPrice)
		}
	}

	if !found {
		fmt.Println("Tidak ada transaksi menggunakan spare-part tersebut.")
	}
}

func favoritServiceMotor() {
	for {
		fmt.Println("======================")
		fmt.Println("Favorit Service Motor")
		fmt.Println("======================")
		fmt.Println("1. Berdasarkan Pelanggan")
		fmt.Println("2. Berdasarkan Spare-part")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih menu: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			favoritBerdasarkanPelanggan()
		case 2:
			favoritBerdasarkanSparePart()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func favoritBerdasarkanPelanggan() {
	count := make([]int, numCustomers)

	for _, transaction := range transactions[:numTransactions] {
		count[transaction.CustomerID]++
	}

	// Sorting pelanggan berdasarkan jumlah transaksi (descending)
	for i := 0; i < numCustomers-1; i++ {
		for j := i + 1; j < numCustomers; j++ {
			if count[i] < count[j] {
				count[i], count[j] = count[j], count[i]
				customers[i], customers[j] = customers[j], customers[i]
			}
		}
	}

	fmt.Println("Pelanggan favorit berdasarkan jumlah transaksi:")
	for i := 0; i < numCustomers; i++ {
		if count[i] > 0 {
			fmt.Printf("Pelanggan: %s, Jumlah transaksi: %d\n", customers[i].Name, count[i])
		}
	}
}

func favoritBerdasarkanSparePart() {
	count := make([]int, numSpareParts)

	for _, transaction := range transactions[:numTransactions] {
		count[transaction.SparePartID] += transaction.Quantity
	}

	// Sorting spare-part berdasarkan jumlah penggunaan (descending)
	for i := 0; i < numSpareParts-1; i++ {
		for j := i + 1; j < numSpareParts; j++ {
			if count[i] < count[j] {
				count[i], count[j] = count[j], count[i]
				spareParts[i], spareParts[j] = spareParts[j], spareParts[i]
			}
		}
	}

	fmt.Println("Spare-part favorit berdasarkan jumlah penggunaan:")
	for i := 0; i < numSpareParts; i++ {
		if count[i] > 0 {
			fmt.Printf("Spare-part: %s, Jumlah penggunaan: %d\n", spareParts[i].Name, count[i])
		}
	}
}
