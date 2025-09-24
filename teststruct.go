package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type mahasiswa struct {
	nama    string
	NIM     string
	Jurusan string
	IPK     float64
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan Data Mahasiswa")

	fmt.Print("Masukkan Nama: ")
	inputNama, _ := reader.ReadString('\n')
	nama := strings.TrimSpace(inputNama)

	fmt.Print("Masukkan NIM: ")
	inputNIM, _ := reader.ReadString('\n')
	NIM := strings.TrimSpace(inputNIM)
	fmt.Print("Masukkan Jurusan: ")
	inputJurusan, _ := reader.ReadString('\n')
	Jurusan := strings.TrimSpace(inputJurusan)

	fmt.Print("Masukkan IPK: ")
	inputIPK, _ := reader.ReadString('\n')
	IPK, err := strconv.ParseFloat(strings.TrimSpace(inputIPK), 64)
	if err != nil {
		fmt.Println("Input IPK tidak valid. Harap masukkan angka desimal.")
		return
	}
	maharu := mahasiswa{
		nama:    nama,
		NIM:     NIM,
		Jurusan: Jurusan,
		IPK:     IPK,
	}
	fmt.Print("Data Mahasiswa berhasil disimpan")
	fmt.Printf("Nama: %s\n", maharu.nama)
	fmt.Printf("NIM: %s\n", maharu.NIM)
	fmt.Printf("Jurusan: %s\n", maharu.Jurusan)
	fmt.Printf("IPK: %.2f\n", maharu.IPK)
}
