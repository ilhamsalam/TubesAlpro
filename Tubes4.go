package main

import (
    "fmt"
)

// Konstanta untuk ukuran maksimum array
const maxSpareParts = 100
const maxPelanggan = 100
const maxTransaksi = 100

// Tipe data untuk spare-part
type SparePart struct {
    Nama  string
    Harga int
}

// Tipe data untuk pelanggan
type Pelanggan struct {
    Nama string
}

// Tipe data untuk transaksi
type Transaksi struct {
    Tanggal          string
    Pelanggan        Pelanggan
    SpareParts       [maxSpareParts]SparePart
    JumlahSpareParts [maxSpareParts]int
}

// Array untuk menyimpan data
var spareParts [maxSpareParts]SparePart
var pelangganList [maxPelanggan]Pelanggan
var transaksiList [maxTransaksi]Transaksi

// Variabel untuk menyimpan jumlah data yang tersimpan
var sparePartCount int
var pelangganCount int
var transaksiCount int

// Fungsi untuk menambahkan data spare-part
func tambahSparePart(nama string, harga int) {
    if sparePartCount < maxSpareParts {
        spareParts[sparePartCount] = SparePart{Nama: nama, Harga: harga}
        sparePartCount++
        fmt.Println("Spare-part ditambahkan:", nama, "dengan harga", harga)
    } else {
        fmt.Println("Data spare-part penuh")
    }
}

// Fungsi untuk menambahkan data pelanggan
func tambahPelanggan(nama string) {
    if pelangganCount < maxPelanggan {
        pelangganList[pelangganCount] = Pelanggan{Nama: nama}
        pelangganCount++
        fmt.Println("Pelanggan ditambahkan:", nama)
    } else {
        fmt.Println("Data pelanggan penuh")
    }
}

// Fungsi untuk menambahkan data transaksi
func tambahTransaksi(tanggal string, namaPelanggan string, namaSpareParts [maxSpareParts]string, jumlahSpareParts [maxSpareParts]int, sparePartsCount int) {
    var pelanggan Pelanggan
    var pelangganDitemukan bool
    for i := 0; i < pelangganCount; i++ {
        if pelangganList[i].Nama == namaPelanggan {
            pelanggan = pelangganList[i]
            pelangganDitemukan = true
        }
    }
    if !pelangganDitemukan {
        fmt.Println("Pelanggan tidak ditemukan:", namaPelanggan)
        return
    }
    if transaksiCount < maxTransaksi {
        transaksi := Transaksi{
            Tanggal:   tanggal,
            Pelanggan: pelanggan,
        }
        for i := 0; i < sparePartsCount; i++ {
            var sparePart SparePart
            var sparePartDitemukan bool
            for j := 0; j < sparePartCount; j++ {
                if spareParts[j].Nama == namaSpareParts[i] {
                    sparePart = spareParts[j]
                    sparePartDitemukan = true
                }
            }
            if sparePartDitemukan {
                transaksi.SpareParts[i] = sparePart
                transaksi.JumlahSpareParts[i] = jumlahSpareParts[i]
            } else {
                fmt.Println("Spare-part tidak ditemukan:", namaSpareParts[i])
                return
            }
        }
        transaksiList[transaksiCount] = transaksi
        transaksiCount++
        fmt.Println("Transaksi ditambahkan")
    } else {
        fmt.Println("Data transaksi penuh")
    }
}

// Fungsi untuk mengedit data spare-part
func editSparePart(namaLama string, namaBaru string, hargaBaru int) {
    var sparePartDitemukan bool
    for i := 0; i < sparePartCount; i++ {
        if spareParts[i].Nama == namaLama {
            spareParts[i].Nama = namaBaru
            spareParts[i].Harga = hargaBaru
            sparePartDitemukan = true
            fmt.Println("Spare-part diedit:", namaLama, "menjadi", namaBaru, "dengan harga baru", hargaBaru)
        }
    }
    if !sparePartDitemukan {
        fmt.Println("Spare-part tidak ditemukan:", namaLama)
    }
}

// Fungsi untuk mengedit data pelanggan
func editPelanggan(namaLama string, namaBaru string) {
    var pelangganDitemukan bool
    for i := 0; i < pelangganCount; i++ {
        if pelangganList[i].Nama == namaLama {
            pelangganList[i].Nama = namaBaru
            pelangganDitemukan = true
            fmt.Println("Pelanggan diedit:", namaLama, "menjadi", namaBaru)
        }
    }
    if !pelangganDitemukan {
        fmt.Println("Pelanggan tidak ditemukan:", namaLama)
    }
}

// Fungsi untuk menghapus data spare-part
func hapusSparePart(nama string) {
    var sparePartDitemukan bool
    for i := 0; i < sparePartCount; i++ {
        if spareParts[i].Nama == nama {
            sparePartDitemukan = true
            // Geser elemen-elemen setelah spare-part yang dihapus
            for j := i; j < sparePartCount-1; j++ {
                spareParts[j] = spareParts[j+1]
            }
            sparePartCount--
            fmt.Println("Spare-part dihapus:", nama)
        }
    }
    if !sparePartDitemukan {
        fmt.Println("Spare-part tidak ditemukan:", nama)
    }
}

// Fungsi untuk menghapus data pelanggan
func hapusPelanggan(nama string) {
    var pelangganDitemukan bool
    for i := 0; i < pelangganCount; i++ {
        if pelangganList[i].Nama == nama {
            pelangganDitemukan = true
            // Geser elemen-elemen setelah pelanggan yang dihapus
            for j := i; j < pelangganCount-1; j++ {
                pelangganList[j] = pelangganList[j+1]
            }
            pelangganCount--
            fmt.Println("Pelanggan dihapus:", nama)
        }
    }
    if !pelangganDitemukan {
        fmt.Println("Pelanggan tidak ditemukan:", nama)
    }
}

// Fungsi mencari dan menampilkan data transaksi berdasarkan tanggal
func cariTransaksiBerdasarkanTanggal(tanggal string) {
    var transaksiDitemukan bool
    for i := 0; i < transaksiCount; i++ {
        if transaksiList[i].Tanggal == tanggal {
            transaksiDitemukan = true
            fmt.Println("Tanggal:", transaksiList[i].Tanggal)
            fmt.Println("Pelanggan:", transaksiList[i].Pelanggan.Nama)
            for j := 0; j < maxSpareParts; j++ {
                if transaksiList[i].JumlahSpareParts[j] > 0 {
                    fmt.Println("Spare-part:", transaksiList[i].SpareParts[j].Nama, "Jumlah:", transaksiList[i].JumlahSpareParts[j])
                }
            }
        }
    }
    if !transaksiDitemukan {
        fmt.Println("Transaksi tidak ditemukan pada tanggal:", tanggal)
    }
}

// Fungsi mencari dan menampilkan data transaksi berdasarkan spare-part
func cariTransaksiBerdasarkanSparePart(namaSparePart string) {
    var transaksiDitemukan bool
    for i := 0; i < transaksiCount; i++ {
        var totalHarga int
        for j := 0; j < maxSpareParts; j++ {
            if transaksiList[i].SpareParts[j].Nama == namaSparePart {
                transaksiDitemukan = true
                totalHarga += transaksiList[i].SpareParts[j].Harga * transaksiList[i].JumlahSpareParts[j]
            }
        }
        if totalHarga > 0 {
            fmt.Println("Pelanggan:", transaksiList[i].Pelanggan.Nama, "Total Harga Service:", totalHarga)
        }
    }
    if !transaksiDitemukan {
        fmt.Println("Transaksi dengan spare-part", namaSparePart, "tidak ditemukan")
    }
}

// Fungsi untuk mengurutkan dan menampilkan data transaksi berdasarkan pelanggan
func urutkanTransaksiBerdasarkanPelanggan() {
    // Hitung frekuensi transaksi untuk setiap pelanggan
    var frekuensi [maxPelanggan]int
    for i := 0; i < transaksiCount; i++ {
        for j := 0; j < pelangganCount; j++ {
            if transaksiList[i].Pelanggan.Nama == pelangganList[j].Nama {
                frekuensi[j]++
            }
        }
    }
    // Lakukan selection sort berdasarkan frekuensi transaksi
    for i := 0; i < pelangganCount-1; i++ {
        for j := i + 1; j < pelangganCount; j++ {
            if frekuensi[i] < frekuensi[j] {
                frekuensi[i], frekuensi[j] = frekuensi[j], frekuensi[i]
                pelangganList[i], pelangganList[j] = pelangganList[j], pelangganList[i]
            }
        }
    }
    // Tampilkan hasil pengurutan
    fmt.Println("Transaksi berdasarkan pelanggan (terurut):")
    for i := 0; i < pelangganCount; i++ {
        if frekuensi[i] > 0 {
            fmt.Println("Pelanggan:", pelangganList[i].Nama, "Jumlah transaksi:", frekuensi[i])
        }
    }
}

// Fungsi untuk mengurutkan dan menampilkan data transaksi berdasarkan spare-part
func urutkanTransaksiBerdasarkanSparePart() {
    // Hitung frekuensi penggunaan setiap spare-part
    var frekuensi [maxSpareParts]int
    for i := 0; i < transaksiCount; i++ {
        for j := 0; j < maxSpareParts; j++ {
            for k := 0; k < sparePartCount; k++ {
                if transaksiList[i].SpareParts[j].Nama == spareParts[k].Nama {
                    frekuensi[k] += transaksiList[i].JumlahSpareParts[j]
                }
            }
        }
    }
    // Lakukan selection sort berdasarkan frekuensi penggunaan spare-part
    for i := 0; i < sparePartCount-1; i++ {
        for j := i + 1; j < sparePartCount; j++ {
            if frekuensi[i] < frekuensi[j] {
                frekuensi[i], frekuensi[j] = frekuensi[j], frekuensi[i]
                spareParts[i], spareParts[j] = spareParts[j], spareParts[i]
            }
        }
    }
    // Tampilkan hasil pengurutan
    fmt.Println("Transaksi berdasarkan spare-part (terurut):")
    for i := 0; i < sparePartCount; i++ {
        if frekuensi[i] > 0 {
            fmt.Println("Spare-part:", spareParts[i].Nama, "Jumlah penggunaan:", frekuensi[i])
        }
    }
}

// Fungsi untuk menampilkan menu utama
func menuUtama() {
    fmt.Println("======================")
    fmt.Println("Aplikasi Service Motor")
    fmt.Println("~~~~~~Kelompok 4~~~~~~")
    fmt.Println("======================")
    fmt.Println("1. Data")
    fmt.Println("2. Histori")
    fmt.Println("3. Keluar")
    var pilihan int
    fmt.Print("Masukkan pilihan: ")
    fmt.Scan(&pilihan)

    switch pilihan {
    case 1:
        menuData()
    case 2:
        menuHistori()
    case 3:
        fmt.Println("Keluar dari aplikasi...")
        return
    default:
        fmt.Println("Pilihan tidak valid")
        menuUtama()
    }
}

// Fungsi untuk menampilkan menu data
func menuData() {
    fmt.Println("======================")
    fmt.Println("Aplikasi Service Motor")
    fmt.Println("~~~~~~~Menu Data~~~~~~")
    fmt.Println("======================")
    fmt.Println("1. Tambah data")
    fmt.Println("2. Edit data")
    fmt.Println("3. Hapus data")
    fmt.Println("4. Kembali")
    var pilihan int
    fmt.Print("Masukkan pilihan: ")
    fmt.Scan(&pilihan)

    switch pilihan {
    case 1:
        menuTambahData()
    case 2:
        menuEditData()
    case 3:
        menuHapusData()
    case 4:
        menuUtama()
    default:
        fmt.Println("Pilihan tidak valid")
        menuData()
    }
}

// Fungsi untuk menampilkan menu tambah data
func menuTambahData() {
    fmt.Println("======================")
    fmt.Println("Aplikasi Service Motor")
    fmt.Println("~~~Menu Tambah data~~~")
    fmt.Println("======================")
    fmt.Println("1. Tambah Data spare-part")
    fmt.Println("2. Tambah Data pelanggan")
    fmt.Println("3. Tambah Data transaksi")
    fmt.Println("4. Kembali")
    var pilihan int
    fmt.Print("Masukkan pilihan: ")
    fmt.Scan(&pilihan)

    switch pilihan {
    case 1:
        var nama string
        var harga int
        fmt.Print("Masukkan nama spare-part: ")
        fmt.Scan(&nama)
        fmt.Print("Masukkan harga spare-part: ")
        fmt.Scan(&harga)
        tambahSparePart(nama, harga)
        menuTambahData()
    case 2:
        var nama string
        fmt.Print("Masukkan nama pelanggan: ")
        fmt.Scan(&nama)
        tambahPelanggan(nama)
        menuTambahData()
    case 3:
        var tanggal, namaPelanggan string
        var namaSpareParts [maxSpareParts]string
        var jumlahSpareParts [maxSpareParts]int
        var sparePartsCount int
        fmt.Print("Masukkan tanggal transaksi: ")
        fmt.Scan(&tanggal)
        fmt.Print("Masukkan nama pelanggan: ")
        fmt.Scan(&namaPelanggan)
        fmt.Print("Masukkan jumlah spare-part yang digunakan: ")
        fmt.Scan(&sparePartsCount)
        for i := 0; i < sparePartsCount; i++ {
            fmt.Printf("Masukkan nama spare-part ke-%d: ", i+1)
            fmt.Scan(&namaSpareParts[i])
            fmt.Printf("Masukkan jumlah spare-part ke-%d: ", i+1)
            fmt.Scan(&jumlahSpareParts[i])
        }
        tambahTransaksi(tanggal, namaPelanggan, namaSpareParts, jumlahSpareParts, sparePartsCount)
        menuTambahData()
    case 4:
        menuData()
    default:
        fmt.Println("Pilihan tidak valid")
        menuTambahData()
    }
}

// Fungsi untuk menampilkan menu edit data
func menuEditData() {
    fmt.Println("======================")
    fmt.Println("Aplikasi Service Motor")
    fmt.Println("~~~~Menu Edit data~~~~")
    fmt.Println("======================")
    fmt.Println("1. Edit Data spare-part")
    fmt.Println("2. Edit Data pelanggan")
    fmt.Println("3. Kembali")
    var pilihan int
    fmt.Print("Masukkan pilihan: ")
    fmt.Scan(&pilihan)

    switch pilihan {
    case 1:
        var namaLama, namaBaru string
        var hargaBaru int
        fmt.Print("Masukkan nama lama spare-part: ")
        fmt.Scan(&namaLama)
        fmt.Print("Masukkan nama baru spare-part: ")
        fmt.Scan(&namaBaru)
        fmt.Print("Masukkan harga baru spare-part: ")
        fmt.Scan(&hargaBaru)
        editSparePart(namaLama, namaBaru, hargaBaru)
        menuEditData()
    case 2:
        var namaLama, namaBaru string
        fmt.Print("Masukkan nama lama pelanggan: ")
        fmt.Scan(&namaLama)
        fmt.Print("Masukkan nama baru pelanggan: ")
        fmt.Scan(&namaBaru)
        editPelanggan(namaLama, namaBaru)
        menuEditData()
    case 3:
        menuData()
    default:
        fmt.Println("Pilihan tidak valid")
        menuEditData()
    }
}

// Fungsi untuk menampilkan menu hapus data
func menuHapusData() {
    fmt.Println("======================")
    fmt.Println("Aplikasi Service Motor")
    fmt.Println("~~~~Menu Hapus data~~~")
    fmt.Println("======================")
    fmt.Println("1. Hapus Data spare-part")
    fmt.Println("2. Hapus Data pelanggan")
    fmt.Println("3. Kembali")
    var pilihan int
    fmt.Print("Masukkan pilihan: ")
    fmt.Scan(&pilihan)

    switch pilihan {
    case 1:
        var nama string
        fmt.Print("Masukkan nama spare-part yang akan dihapus: ")
        fmt.Scan(&nama)
        hapusSparePart(nama)
        menuHapusData()
    case 2:
        var nama string
        fmt.Print("Masukkan nama pelanggan yang akan dihapus: ")
        fmt.Scan(&nama)
        hapusPelanggan(nama)
        menuHapusData()
    case 3:
        menuData()
    default:
        fmt.Println("Pilihan tidak valid")
        menuHapusData()
    }
}

// Fungsi untuk menampilkan menu histori
func menuHistori() {
    fmt.Println("======================")
    fmt.Println("Aplikasi Service Motor")
    fmt.Println("~~~~~Menu Histori~~~~~")
    fmt.Println("======================")
    fmt.Println("1. Histori Service motor")
    fmt.Println("2. Favorit Service motor")
    fmt.Println("3. Kembali")
    var pilihan int
    fmt.Print("Masukkan pilihan: ")
    fmt.Scan(&pilihan)

    switch pilihan {
    case 1:
        menuHistoriService()
    case 2:
        menuFavoritService()
    case 3:
        menuUtama()
    default:
        fmt.Println("Pilihan tidak valid")
        menuHistori()
    }
}

// Fungsi untuk menampilkan menu histori service motor
func menuHistoriService() {
    fmt.Println("======================")
    fmt.Println("Aplikasi Service Motor")
    fmt.Println("Histori Service motor")
    fmt.Println("======================")
    fmt.Println("1. Berdasarkan Tanggal")
    fmt.Println("2. Berdasarkan Spare-part")
    fmt.Println("3. Kembali")
    var pilihan int
    fmt.Print("Masukkan pilihan: ")
    fmt.Scan(&pilihan)

    switch pilihan {
    case 1:
        var tanggal string
        fmt.Print("Masukkan tanggal transaksi: ")
        fmt.Scan(&tanggal)
        cariTransaksiBerdasarkanTanggal(tanggal)
        menuHistoriService()
    case 2:
        var namaSparePart string
        fmt.Print("Masukkan nama spare-part: ")
        fmt.Scan(&namaSparePart)
        cariTransaksiBerdasarkanSparePart(namaSparePart)
        menuHistoriService()
    case 3:
        menuHistori()
    default:
        fmt.Println("Pilihan tidak valid")
        menuHistoriService()
    }
}

// Fungsi untuk menampilkan menu favorit service motor
func menuFavoritService() {
    fmt.Println("======================")
    fmt.Println("Aplikasi Service Motor")
    fmt.Println("Favorit Service motor")
    fmt.Println("======================")
    fmt.Println("1. Berdasarkan Pelanggan")
    fmt.Println("2. Berdasarkan Spare-part")
    fmt.Println("3. Kembali")
    var pilihan int
    fmt.Print("Masukkan pilihan: ")
    fmt.Scan(&pilihan)

    switch pilihan {
    case 1:
        urutkanTransaksiBerdasarkanPelanggan()
        menuFavoritService()
    case 2:
        urutkanTransaksiBerdasarkanSparePart()
        menuFavoritService()
    case 3:
        menuHistori()
    default:
        fmt.Println("Pilihan tidak valid")
        menuFavoritService()
    }
}

func main() {
    // Menampilkan menu utama
    menuUtama()
}
