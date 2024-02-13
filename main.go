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

	users["1234567890"] = Akun{"1234567890", "password123", 100.0}

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
	fmt.Print("Enter phone number: ")
	fmt.Scanln(&phoneNumber)

	if _, ok := users[phoneNumber]; ok {
		fmt.Println("Account already exists.")
		return
	}

	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	users[phoneNumber] = Akun{phoneNumber, password, 0.0}
	fmt.Println("Account created successfully.")
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
			users[to].Balance += amount
			transaksi = append(transaksi, Transaksi{from, to, amount})
			fmt.Println("Transfer successful.")
		} else {
			fmt.Println("Insufficient balance.")
		}
	} else {
		fmt.Println("Account not found.")
	}
}

func viewProfile() {
	var phoneNumber string
	fmt.Print("Enter phone number to view profile: ")
	fmt.Scanln(&phoneNumber)

	if user, ok := users[phoneNumber]; ok {
		fmt.Println("Phone Number:", user.NoHP)
		fmt.Println("Balance:", user.Balance)
	} else {
		fmt.Println("Account not found.")
	}
}

func viewTransactionHistory() {
	fmt.Println("Transaction History:")
	for i, transaction := range transaksi {
		fmt.Println("Transaction", i+1)
		fmt.Println("From:", transaction.Dari)
		fmt.Println("To:", transaction.Ke)
		fmt.Println("Amount:", strconv.FormatFloat(transaction.Jumlah, 'f', 2, 64))
	}
}
