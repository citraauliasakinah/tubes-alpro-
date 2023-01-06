package main

import (
	"fmt"
	"os"
	"os/exec"
)

const N int = 8
const M int = 22

const T int = 5

type penumpang struct {
	np                string
	nama              string
	usia              int
	kelas_penerbangan string
}

type tiket struct {
	kode_tiket int
	data       [T]penumpang
	jumlah     int
}

type bayi struct {
	np          string
	tempatDuduk string
}

type pesawat struct {
	bisnis         [N][4]penumpang
	ekonomi        [M][6]penumpang
	jumlah_bisnis  int
	jumlah_ekonomi int
}

type arrBayi struct {
	data   [100]bayi
	jumlah int
}

// ---------

var plane pesawat
var karcis [100]tiket
var dataBayi arrBayi

// ---------

func menu(pilih *int) {
	fmt.Println("SELAMAT DATANG DI tiket.ai\nSILAHKAN PILIH OPTION BERIKUT\n1. lakukan reservasi\n2. lihat kursi pesawat\n3. ...\n0. exit (selesai)")

	fmt.Scan(&*pilih)
}

func printPesawat() {
	for i := 0; i < N; i++ {
		for j := 0; j < 4; j++ {
			if plane.bisnis[i][j].np == "" {
				fmt.Print("{}", " ")
			} else {
				fmt.Print(plane.bisnis[i][j], " ")
			}
		}
		fmt.Println()
	}
	for i := 0; i < M; i++ {
		for j := 0; j < 6; j++ {
			if plane.ekonomi[i][j].np == "" {
				fmt.Print("{}", " ")
			} else {
				fmt.Print(plane.ekonomi[i][j], " ")
			}
		}
		fmt.Println()
	}
}

func reservasi() {
	var inputed bool = false

	var data penumpang

	fmt.Print("Masukkan kelas penerbangan :")
	fmt.Scan(&data.kelas_penerbangan)

	for i := 0; i < T && !inputed; i++ {
		fmt.Println("Data Penumpang", i+1)

		fmt.Print("Masukkan nomor passport : ")
		fmt.Scan(&data.np)

		fmt.Print("Masukkan nama anda : ")
		fmt.Scan(&data.nama)

		fmt.Print("Masukkan usia anda : ")
		fmt.Scan(&data.usia)

		if data.kelas_penerbangan == "ekonomi" {

			var stop bool = false

			for i := 0; i < M && !stop; i++ {
				for j := 0; j < 6 && !stop; j++ {
					if plane.ekonomi[i][j].np == "" {
						plane.ekonomi[i][j] = data

						stop = true
					}
				}
			}
		} else {
			var stop bool = false

			for i := 0; i < N && !stop; i++ {
				for j := 0; j < 4 && !stop; j++ {
					if plane.bisnis[i][j].np == "" {
						plane.bisnis[i][j] = data

						stop = true
					}
				}
			}
		}

		printPesawat()

		var confirm string
		fmt.Print("Masih ada lagi? (Y/N)")

		fmt.Scan(&confirm)
		if confirm != "Y" && confirm != "y" {
			inputed = true
		}
	}
}

func main() {
	var running bool = true

	for running {

		var pilih int
		menu(&pilih)

		switch pilih {
		case 1:
			reservasi()
		case 2:
			printPesawat()
		case 3:
			fmt.Print("case 3")
		default:
			fmt.Print("exit")
			running = false
		}

		fmt.Scanln()
		fmt.Scanln()

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
