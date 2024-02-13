package main

import (
	"fmt"
	"strconv"
)

type Akun struct {
	NoHP     string
	Password string
	Balance  float64
}

type Transaksi struct {
	Dari   string
	Ke     string
	Jumlah float64
}

var users map[string]Akun
var transaksi []Transaksi

func main() {
	users = make(map[string]Akun)

	var choice int
	for {
		fmt.Println("1. Buat akun")
		fmt.Println("2. Login")
		fmt.Println("3. Top-up Balance")
		fmt.Println("4. Transfer Balance")
		fmt.Println("5. Lihat Profil")
		fmt.Println("6. Lihat Histori Transaksi")
		fmt.Println("7. Exit")

		fmt.Print("Masukkan pilihan: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			createAccount()
		case 2:
			login()
		case 3:
			topUpBalance()
		case 4:
			transferBalance()
		case 5:
			viewProfile()
		case 6:
			viewTransactionHistory()
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func createAccount() {
	var phoneNumber, password string
	fmt.Print("Masukkan No. HP: ")
	fmt.Scanln(&phoneNumber)

	if _, ok := users[phoneNumber]; ok {
		fmt.Println("Akun sudah ada.")
		return
	}

	fmt.Print("Masukkan password: ")
	fmt.Scanln(&password)

	users[phoneNumber] = Akun{phoneNumber, password, 0.0}
	fmt.Println("Akun berhasil dibuat.")
}

func login() {
	var phoneNumber, password string
	fmt.Print("Masukkan No. HP: ")
	fmt.Scanln(&phoneNumber)

	fmt.Print("Masukkan password: ")
	fmt.Scanln(&password)

	if user, ok := users[phoneNumber]; ok {
		if user.Password == password {
			fmt.Println("Login berhasil.")
			return
		}
	}
	fmt.Println("No. HP dan password salah.")
}

func topUpBalance() {
	var phoneNumber string
	var amount float64
	fmt.Print("Masukkan No. HP: ")
	fmt.Scanln(&phoneNumber)

	if user, ok := users[phoneNumber]; ok {
		fmt.Print("Masukkan jumlah top up: ")
		fmt.Scanln(&amount)
		user.Balance += amount
		fmt.Println("Top-up berhasil.")
	} else {
		fmt.Println("Akun tidak ada.")
	}
}

func transferBalance() {
	var from, to string
	var amount float64
	fmt.Print("Masukkan No. HP: ")
	fmt.Scanln(&from)

	fmt.Print("Masukkan No. HP penerima: ")
	fmt.Scanln(&to)

	fmt.Print("Masukkan jumlah transfer: ")
	fmt.Scanln(&amount)

	if user, ok := users[from]; ok {
		if user.Balance >= amount {
			user.Balance -= amount
			user.Balance += amount
			transaksi = append(transaksi, Transaksi{from, to, amount})
			fmt.Println("Transfer successful.")
		} else {
			fmt.Println("Balance tidak cukup.")
		}
	} else {
		fmt.Println("Akun tidak ada.")
	}
}

func viewProfile() {
	var phoneNumber string
	fmt.Print("Masukan No. HP untuk melihat profil: ")
	fmt.Scanln(&phoneNumber)

	if user, ok := users[phoneNumber]; ok {
		fmt.Println("No. HP:", user.NoHP)
		fmt.Println("Balance:", user.Balance)
	} else {
		fmt.Println("Akun tidak ada.")
	}
}

func viewTransactionHistory() {
	fmt.Println("History Transaksi:")
	for i, transaction := range transaksi {
		fmt.Println("Transaksi", i+1)
		fmt.Println("Dari:", transaction.Dari)
		fmt.Println("Ke:", transaction.Ke)
		fmt.Println("Jumlah:", strconv.FormatFloat(transaction.Jumlah, 'f', 2, 64))
	}
}
