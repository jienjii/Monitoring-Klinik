package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

const MAX = 100

type pasien struct {
	id       int
	nama     string
	umur     int
	penyakit string
}

type obat struct {
	kodeObat string
	namaObat string
	stok     int
}

type dokter struct {
	idDokter   int
	namaDokter string
	spesialis  string
}

var Pasien [MAX]pasien
var Obat [MAX]obat
var Dokter [MAX]dokter

func TambahPasien(p *[MAX]pasien, npasien *int, tambah pasien) {
	if *npasien < MAX {
		p[*npasien] = tambah
		*npasien++
	}
	fmt.Println("Data Pasien Berhasil Ditambahkan!")
}

func TampilkanDataPasien(p *[MAX]pasien, npasien int) {
	for i := 0; i < npasien; i++ {
		fmt.Printf ("Data Pasien ke-%d :\n", i+1)
		fmt.Printf("ID: %d\nNama: %s\nUmur: %d\nPenyakit: %s\n", p[i].id, p[i].nama, p[i].umur, p[i].penyakit)
	}
}

func SelectionSortPasien(p *[MAX]pasien, npasien int) {
	for i := 0; i < npasien; i++ {
		min := i
		for j := i + 1; j < npasien; j++ {
			if p[j].id < p[min].id {
				min = j
			}
		}
		p[i], p[min] = p[min], p[i]
	}
}
func SelectionsortPDescending(p*[MAX]pasien, npasien int){
	for i := 0; i < npasien; i++ {
		min := i
		for j := i + 1; j < npasien; j++ {
			if p[j].id > p[min].id {
				min = j
			}
		}
		p[i], p[min] = p[min], p[i]
	}
	fmt.Println("Data Pasien Berhasil Diurutkan Secara Descending!")
}
func InsertionSortPDescending(p*[MAX]pasien, npasien int){
	for i := 1; i < npasien; i++ {
		x := p[i]
		j := i - 1
		for j >= 0 && strings.ToLower(p[j].nama) < strings.ToLower(x.nama) {
			p[j+1] = p[j]
			j--
		}
		p[j+1] = x
	}
		fmt.Println("Data Pasien Berhasil Diurutkan Secara Descending!")
}

func CariPasienID(p *[MAX]pasien, npasien, id int) int {
	SelectionSortPasien(p, npasien)
	kiri := 0
	kanan := npasien - 1
	for kiri <= kanan {
		mid := (kiri + kanan) / 2
		if p[mid].id == id {
			return mid
		} else if id < p[mid].id {
			kanan = mid - 1
		} else {
			kiri = mid + 1
		}
	}
	return -1
}

func CariPasienNama(p [MAX]pasien, npasien int, nama string) int {
	for i := 0; i < npasien; i++ {
		if p[i].nama == nama {
			return i
		}
	}
	return -1
}

func hapusPasien(p *[MAX]pasien, npasien *int, id int) {
	for i := 0; i < *npasien; i++ {
		if p[i].id == id {
			for j := i; j < *npasien-1; j++ {
				p[j] = p[j+1]
			}
			*npasien--
			fmt.Println("Pasien berhasil dihapus")
			return
		}
	}
	fmt.Println("Pasien tidak ditemukan")
}

func InsertionSortPasien(p *[MAX]pasien, npasien int) {
	for i := 1; i < npasien; i++ {
		x := p[i]
		j := i - 1
		for j >= 0 && strings.ToLower(p[j].nama) > strings.ToLower(x.nama) {
			p[j+1] = p[j]
			j--
		}
		p[j+1] = x
	}

	fmt.Println("Data Pasien Berhasil Diurutkan Secara Ascending!")
}

func TambahObat(o *[MAX]obat, nobat *int, tambah obat) {
	if *nobat < MAX {
		o[*nobat] = tambah
		*nobat++
	}
	fmt.Println("Data Obat Berhasil Ditambahkan!")
}

func TampilkanRincianObat(o [MAX]obat, nobat int) {
	for i := 0; i < nobat; i++ {
		fmt.Printf("Data Obat ke-%d :\n", i+1)
		fmt.Printf("Kode Obat : %s\nNama obat : %s\nStok: %d\n", o[i].kodeObat, o[i].namaObat, o[i].stok)
	}
}

func CariObat(o [MAX]obat, nobat int, kode string) int {
	for i := 0; i < nobat; i++ {
		if o[i].kodeObat == kode {
			return i
		}
	}
	return -1
}

func HapusObat(o *[MAX]obat, nobat *int, kode string) {
	for i := 0; i < *nobat; i++ {
		if o[i].kodeObat == kode {
			for j := i; j < *nobat-1; j++ {
				o[j] = o[j+1]
			}
			*nobat--
			fmt.Println("Obat berhasil dihapus")
			return
		}
	}
	fmt.Println("Obat tidak ditemukan")
}

func UrutObatKode(o *[MAX]obat, nobat int) {
	for i := 1; i < nobat; i++ {
		x := o[i]
		j := i - 1
		for j >= 0 && strings.ToLower(o[j].kodeObat) > strings.ToLower(x.kodeObat) {
			o[j+1] = o[j]
			j--
		}
		o[j+1] = x
	}
	fmt.Println("Berhasil Diurutkan Secara Ascending!")
}

func UrutObatNama(o *[MAX]obat, nobat int) {
	for i := 1; i < nobat; i++ {
		x := o[i]
		j := i - 1
		for j >= 0 && strings.ToLower(o[j].namaObat) > strings.ToLower(x.namaObat) {
			o[j+1] = o[j]
			j--
		}
		o[j+1] = x
	}
	fmt.Println("Berhasil Diurutkan Secara Ascending!")
}

func UrutObatStok(o *[MAX]obat, nobat int) {
	for i := 0; i < nobat; i++ {
		min := i
		for j := i + 1; j < nobat; j++ {
			if o[min].stok > o[j].stok {
				min = j
			}
		}
		o[i], o[min] = o[min], o[i]
	}
	fmt.Println("Berhasil Diurutkan Secara Ascending!")
}
func InsertionSortODescending(o*[MAX]obat, nobat int){
	for i := 1; i < nobat; i++ {
		x := o[i]
		j := i - 1
		for j >= 0 && strings.ToLower(o[j].kodeObat) < strings.ToLower(x.kodeObat) {
			o[j+1] = o[j]
			j--
		}
		o[j+1] = x
	}
	fmt.Println("Berhasil Diurutkan Secara Descending!")
}
func SelectionSortODescending(o*[MAX]obat, nobat int){
	for i := 0; i < nobat; i++ {
		min := i
		for j := i + 1; j < nobat; j++ {
			if o[min].stok < o[j].stok {
				min = j
			}
		}
		o[i], o[min] = o[min], o[i]
	}
	fmt.Println("Berhasil Diurutkan Secara Descending!")
}

func TambahDokter(d *[MAX]dokter, ndokter *int, tambah dokter) {
	if *ndokter < MAX {
		d[*ndokter] = tambah
		*ndokter++
	}
	fmt.Println("Data Dokter Berhasil Ditambahkan!")
}

func TampilkanDataDokter(d *[MAX]dokter, ndokter int) {
	for i := 0; i < ndokter; i++ {
		fmt.Printf("Data Dokter Ke-%d :\n", i+1)
		fmt.Printf("ID: %d\nNama: %s\nSpesialis: %s\n", d[i].idDokter, d[i].namaDokter, d[i].spesialis)
	}
}

func SelectionSortDokter(d *[MAX]dokter, ndokter int) {
	for i := 0; i < ndokter; i++ {
		min := i
		for j := i + 1; j < ndokter; j++ {
			if d[j].idDokter < d[min].idDokter {
				min = j
			}
		}
		d[i], d[min] = d[min], d[i]
	}
}
func SelectionSortDDescending(d *[MAX]dokter, ndokter int) {
	for i := 0; i < ndokter; i++ {
		min := i
		for j := i + 1; j < ndokter; j++ {
			if d[j].idDokter > d[min].idDokter {
				min = j
			}
		}
		d[i], d[min] = d[min], d[i]
	}
	fmt.Println("Data Berhasil Diurutkan Secara Descending!")
}

func CariDokterID(d *[MAX]dokter, ndokter int, id int) int {
	SelectionSortDokter(d, ndokter)
	kiri := 0
	kanan := ndokter - 1
	for kiri <= kanan {
		mid := (kiri + kanan) / 2
		if d[mid].idDokter == id {
			return mid
		} else if id < d[mid].idDokter {
			kanan = mid - 1
		} else {
			kiri = mid + 1
		}
	}
	return -1
}

func CariDokterNama(d [MAX]dokter, ndokter int, nama string) int {
	for i := 0; i < ndokter; i++ {
		if d[i].namaDokter == nama {
			return i
		}
	}
	return -1
}

func HapusDokter(d *[MAX]dokter, ndokter *int, id int) {
	for i := 0; i < *ndokter; i++ {
		if d[i].idDokter == id {
			for j := i; j < *ndokter-1; j++ {
				d[j] = d[j+1]
			}
			*ndokter--
			fmt.Println("Data Berhasil Dihapus!")
			return
		}
	}
	fmt.Println("Dokter tidak ditemukan")
}

func InsertionSortDokter(d *[MAX]dokter, ndokter int) {
	for i := 1; i < ndokter; i++ {
		x := d[i]
		j := i - 1
		for j >= 0 && strings.ToLower(d[j].namaDokter) > strings.ToLower(x.namaDokter) {
			d[j+1] = d[j]
			j--
		}
		d[j+1] = x
	}
	fmt.Println("Data Berhasil Diurutkan Secara Ascending!")
}
func InsertionSortDDokter(d *[MAX]dokter, ndokter int) {
	for i := 1; i < ndokter; i++ {
		x := d[i]
		j := i - 1
		for j >= 0 && strings.ToLower(d[j].namaDokter) < strings.ToLower(x.namaDokter) {
			d[j+1] = d[j]
			j--
		}
		d[j+1] = x
	}
	fmt.Println("Data Berhasil Diurutkan Secara Descending!")
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	var pilih, pilihMenu int
	var npasien int
	var nobat int
	var ndokter int
	selesai := false
	for !selesai {
		fmt.Println("===Sistem monitoring Klinik===")
		fmt.Println("1. Kelola Data Pasien")
		fmt.Println("2. Kelola Data Obat")
		fmt.Println("3. Kelola Data Dokter")
		fmt.Println("4. Keluar")
		fmt.Println()
		fmt.Scan(&pilih)
		if pilih == 1 {
			kembali := false
			for !kembali {
				fmt.Println()
				fmt.Println("===Menu Pasien===")
				fmt.Println("1. Tambah Pasien")
				fmt.Println("2. Tampilkan Data Pasien")
				fmt.Println("3. Cari Pasien (Berdasarkan ID)")
				fmt.Println("4. Cari Pasien (Berdasarkan Nama)")
				fmt.Println("5. Hapus Pasien (Berdasarkan ID)")
				fmt.Println("6. Urutkan Pasien Asceding (Berdasarkan ID)")
				fmt.Println("7. Urutkan Pasien Ascending (Berdasarkan Nama)")
				fmt.Println("8. Urutkan Pasien Descending (Berdasarkan ID)")
				fmt.Println("9. Urutkan Pasien Descending (Berdasarkan Nama)")
				fmt.Println("10. Kembali")
				fmt.Println()
				fmt.Scan(&pilihMenu)
				reader.ReadString('\n')
				var p pasien
					if pilihMenu == 1 {
						fmt.Print("Masukkan ID Pasien: ")
						idPasien, _ := reader.ReadString('\n')
						p.id, _ = strconv.Atoi(strings.TrimSpace(idPasien))

						fmt.Print("Masukkan Nama Pasien: ")
						p.nama, _ = reader.ReadString('\n')
						p.nama = strings.TrimSpace(p.nama)

						fmt.Print("Masukkan Umur Pasien: ")
						umurPasien, _ := reader.ReadString('\n')
						p.umur, _ = strconv.Atoi(strings.TrimSpace(umurPasien))

						fmt.Print("Masukkan Jenis Penyakit Pasien: ")
						p.penyakit, _ = reader.ReadString('\n')
						p.penyakit = strings.TrimSpace(p.penyakit)

						TambahPasien(&Pasien, &npasien, p)
					}else if pilihMenu == 2 {
							TampilkanDataPasien(&Pasien, npasien)
					}else if pilihMenu == 3 {
							fmt.Print("Masukkan id : ")
							idStr,_ := reader.ReadString('\n')
							cariid, _ := strconv.Atoi(strings.TrimSpace(idStr))
							x := CariPasienID(&Pasien, npasien, cariid)
							if  x != -1 {
								fmt.Println("Data Ditemukan!")
								fmt.Printf("ID: %d\nNama: %s\nUmur: %d\nPenyakit: %s\n", Pasien[x].id, Pasien[x].nama, Pasien[x].umur, Pasien[x].penyakit)
							}else {
								fmt.Println("Data Tidak Ditemukan!")
							}
					}else if pilihMenu == 4 {
						fmt.Print("Masukkan nama : ")
						var carinama string 
						carinama, _ = reader.ReadString('\n')
						carinama = strings.TrimSpace(carinama)
						x := CariPasienNama(Pasien, npasien, carinama)
						if x != -1 {
							fmt.Println("Data Ditemukan!")
							fmt.Printf("ID: %d\nNama: %s\nUmur: %d\nPenyakit: %s\n", Pasien[x].id, Pasien[x].nama, Pasien[x].umur, Pasien[x].penyakit)
						}else{
							fmt.Println("Data Tidak Ditemukan!")
						}
					}else if pilihMenu == 5 {
						fmt.Print("Masukkan ID : ")
						idStr, _ := reader.ReadString('\n')
						hapusid, _ := strconv.Atoi(strings.TrimSpace(idStr))
						hapusPasien(&Pasien,&npasien,hapusid)
					}else if pilihMenu == 6 {
						SelectionSortPasien(&Pasien, npasien)
						fmt.Println("Data Pasien Berhasil Diurutkan Secara Ascending!")
					}else if pilihMenu == 7 {
						InsertionSortPasien(&Pasien, npasien)
					}else if pilihMenu == 8 {
						SelectionsortPDescending(&Pasien, npasien)					
					}else if pilihMenu == 9 {
						InsertionSortPDescending(&Pasien, npasien)
					}else if pilihMenu == 10 {
						kembali = true
					}
				}
		}else if pilih == 2  {
			kembali := false
			for !kembali {
				fmt.Println()
				fmt.Println("===Menu Obat===")
				fmt.Println("1. Tambah Obat")
				fmt.Println("2. Tampilkan Rincian Obat")
				fmt.Println("3. Cari Obat (Berdasarkan Kode)")
				fmt.Println("4. Hapus Obat (Berdasarkan Kode)")
				fmt.Println("5. Urutkan Obat Ascending (Berdasarkan Kode)")
				fmt.Println("6. Urutkan Obat Ascending (Berdasarkan Nama)")
				fmt.Println("7. Urutkan Obat Berdasarkan Stok")
				fmt.Println("8. Urutkan Obat Descending (Berdasarkan Kode)")
				fmt.Println("9. Urutkan stok Descending (Berdasarkan Stok)")
				fmt.Println("10. Kembali")
				fmt.Println()
				fmt.Scan(&pilihMenu)
				reader.ReadString('\n')
					if pilihMenu == 1 {
						var o obat
						fmt.Print("Masukkan Kode Obat: ")
						o.kodeObat,_= reader.ReadString('\n')
						o.kodeObat = strings.TrimSpace(o.kodeObat)
						fmt.Print("Masukkan Nama Obat : ")
						o.namaObat,_ = reader.ReadString('\n')
						o.namaObat = strings.TrimSpace(o.namaObat)
						fmt.Print("Masukkan jumlah obat: ")
						stokStr, _ := reader.ReadString('\n')
						o.stok,_ = strconv.Atoi(strings.TrimSpace(stokStr))
						TambahObat(&Obat, &nobat, o)
					}else if pilihMenu == 2 {
						TampilkanRincianObat(Obat, nobat)
					}else if pilihMenu == 3 {
						var carikode string 
							fmt.Print("Masukkan Kode Yang Dicari: ")
							carikode,_ = reader.ReadString('\n')
							carikode = strings.TrimSpace(carikode)
							x := CariObat(Obat, nobat, carikode)
							if x != -1 {
								fmt.Println("Obat Ditemukan!")
								fmt.Printf("Kode Obat : %s\nNama obat : %s\nStok: %d\n", Obat[x].kodeObat, Obat[x].namaObat, Obat[x].stok)
							}else {
								fmt.Println("Obat Tidak Ditemukan!")
							}
					}else if pilihMenu == 4 { 
						fmt.Print("Masukkan Kode Obat: ")
						hapuskode,_ := reader.ReadString('\n')
						hapuskode = strings.TrimSpace(hapuskode)
						HapusObat(&Obat, &nobat, hapuskode)
					}else if pilihMenu == 5 {
						UrutObatKode(&Obat, nobat)
					}else if pilihMenu == 6 {
						UrutObatNama(&Obat, nobat)
					}else if pilihMenu == 7 {
						UrutObatStok(&Obat, nobat)
					}else if  pilihMenu == 8 {
						InsertionSortODescending(&Obat, nobat)
					}else if pilihMenu == 9 {
						SelectionSortODescending(&Obat, nobat)
					}else if pilihMenu == 10 {
						kembali = true
					}
				}
		}else if pilih == 3 {
			kembali := false 
			for !kembali {
				fmt.Println()
				fmt.Println("===Menu Dokter===")
				fmt.Println("1. Tambah Dokter")
				fmt.Println("2. Tampilkan Data Dokter")
				fmt.Println("3. Cari Dokter (Berdasarkan ID)")
				fmt.Println("4. Cari Dokter (Berdasarkan Nama)")
				fmt.Println("5. Hapus Dokter (Berdasarkan ID)")
				fmt.Println("6. Urutkan Dokter Ascending (Berdasarkan ID)")
				fmt.Println("7. Urutkan Dokter Ascending (Berdasarkan Nama)")
				fmt.Println("8. Urutkan Dokter Descending (Berdasarkan ID)")
				fmt.Println("9. Urutkan Dokter Descending (Berdasarkan Nama)")
				fmt.Println("10. Kembali")
				fmt.Println()
				fmt.Scan(&pilihMenu)
				reader.ReadString('\n')
					if pilihMenu == 1 {
						var d dokter
						fmt.Print("Masukkan ID dokter : ")
						idStr,_ := reader.ReadString('\n')
						d.idDokter,_= strconv.Atoi(strings.TrimSpace(idStr))
						fmt.Print("Masukkan Nama Dokter : ")
						d.namaDokter,_ = reader.ReadString('\n')
						d.namaDokter = strings.TrimSpace(d.namaDokter)
						fmt.Print("Masukan Spesialisasi Dokter : ")
						d.spesialis,_= reader.ReadString('\n')
						d.spesialis = strings.TrimSpace(d.spesialis)
						TambahDokter(&Dokter, &ndokter, d )
					}else if pilihMenu == 2 {
						TampilkanDataDokter(&Dokter, ndokter)
					}else if pilihMenu == 3 {
						fmt.Print("Masukkan ID: ")
						idStr,_ := reader.ReadString('\n')
						cariid,_ := strconv.Atoi(strings.TrimSpace(idStr))
						x := CariDokterID(&Dokter, ndokter,cariid)
						if x != -1 {
							fmt.Println("Data Ditemukan!")
							fmt.Printf("ID: %d\nNama: %s\nSpesialis: %s\n", Dokter[x].idDokter, Dokter[x].namaDokter, Dokter[x].spesialis)

						}else {
							fmt.Println("Data Tidak Ditemukan!")
						}
					}else if pilihMenu == 4 {
						var carinama string 
						fmt.Print("Masukkan Nama: ")
						carinama,_ = reader.ReadString('\n')
						carinama = strings.TrimSpace(carinama)
						x := CariDokterNama(Dokter,ndokter,carinama)
						if x != -1 {
							fmt.Println("Data Ditemukan!")
							fmt.Printf("ID: %d\nNama: %s\nSpesialis: %s\n", Dokter[x].idDokter, Dokter[x].namaDokter, Dokter[x].spesialis)

						}else {
							fmt.Println("Data Tidak Ditemukan!")
						}
					}else if pilihMenu == 5 { 
						fmt.Print("Masukkan ID: ")
						idStr,_ := reader.ReadString('\n')
						hapusdokter,_ := strconv.Atoi(strings.TrimSpace(idStr))
						HapusDokter(&Dokter, &ndokter, hapusdokter)
					}else if pilihMenu == 6 {
						SelectionSortDokter(&Dokter, ndokter)
						fmt.Println("Data Berhasil Diurutkan Secara Ascending!")
					}else if pilihMenu == 7 {
						InsertionSortDokter(&Dokter, ndokter)
					}else if pilihMenu == 8 {
						SelectionSortDDescending(&Dokter, ndokter)
					}else if pilihMenu == 9 {
						InsertionSortDDokter(&Dokter, ndokter)
					}else if pilihMenu == 10 {
						kembali = true
					}
				}
		} else if pilih == 4 {
			selesai = true
		}
	}
}
