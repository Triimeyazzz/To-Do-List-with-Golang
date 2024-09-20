package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	tasks := []string{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== To-Do List ===")
		fmt.Println("1. Tambah Tugas")
		fmt.Println("2. Lihat Tugas")
		fmt.Println("3. Hapus Tugas")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih opsi (1-4): ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Print("Masukkan tugas baru: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			tasks = append(tasks, task)
			fmt.Println("Tugas berhasil ditambahkan!")
		case "2":
			if len(tasks) == 0 {
				fmt.Println("Tidak ada tugas!")
			} else {
				fmt.Println("Daftar Tugas:")
				for i, task := range tasks {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}
		case "3":
			if len(tasks) == 0 {
				fmt.Println("Tidak ada tugas untuk dihapus!")
			} else {
				fmt.Print("Masukkan nomor tugas yang akan dihapus: ")
				taskNumStr, _ := reader.ReadString('\n')
				taskNumStr = strings.TrimSpace(taskNumStr)
				taskNum, err := strconv.Atoi(taskNumStr)
				if err != nil || taskNum < 1 || taskNum > len(tasks) {
					fmt.Println("Nomor tugas tidak valid!")
				} else {
					tasks = append(tasks[:taskNum-1], tasks[taskNum:]...)
					fmt.Println("Tugas berhasil dihapus!")
				}
			}
		case "4":
			fmt.Println("Keluar dari aplikasi...")
			return
		default:
			fmt.Println("Opsi tidak valid! Pilih antara 1-4.")
		}
	}
}
