package main

import (
	"bufio"
	"fmt"
	"os"
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

	fmt.print("Masukkan NIM: ")
	inputNIM, _ := reader.ReadString()
	NIM := strings.TrimSpace(inputNIM)
}
