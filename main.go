package main

import (
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Akun struct {
	gorm.Model
	NoHP     string `gorm:"unique"`
	Password string
	Balance  float64
}

type Transaksi struct {
	gorm.Model
	Dari   string
	Ke     string
	Jumlah float64
}

var db *gorm.DB

func main() {
	var err error
	dsn := "root:@tcp(localhost:3306)/groupproject?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	err = db.AutoMigrate(&Akun{}, &Transaksi{})
	if err != nil {
		log.Fatalf("Failed to auto migrate models: %v", err)
	}

	var menu int
	for {
		fmt.Println("1. Buat akun")
		fmt.Println("2. Login")
		fmt.Println("3. Top-up Balance")
		fmt.Println("4. Transfer Balance")
		fmt.Println("5. Lihat Profil")
		fmt.Println("6. Lihat Histori Transaksi")
		fmt.Println("7. Exit")

		fmt.Print("Masukkan pilihan: ")
		fmt.Scanln(&menu)

		switch menu {
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
			fmt.Println("Pilihan salah.")
		}
	}
}

func createAccount() {
	var phoneNumber, password string
	fmt.Print("Masukkan No. HP: ")
	fmt.Scanln(&phoneNumber)

	var existingUser Akun
	result := db.Where("no_hp = ?", phoneNumber).First(&existingUser)
	if result.Error == nil {
		fmt.Println("Akun sudah ada.")
		return
	}

	fmt.Print("Masukkan password: ")
	fmt.Scanln(&password)

	newUser := Akun{NoHP: phoneNumber, Password: password}
	result = db.Create(&newUser)
	if result.Error != nil {
		log.Fatalf("Failed to create account: %v", result.Error)
	}
	fmt.Println("Akun berhasil dibuat.")
}

func login() {
	var phoneNumber, password string
	fmt.Print("Masukkan No. HP: ")
	fmt.Scanln(&phoneNumber)

	fmt.Print("Masukkan password: ")
	fmt.Scanln(&password)

	var user Akun
	result := db.Where("no_hp = ? AND password = ?", phoneNumber, password).First(&user)
	if result.Error == nil {
		fmt.Println("Login berhasil.")
		return
	}
	fmt.Println("No. HP dan password salah.")
}

func topUpBalance() {
	var phoneNumber string
	var amount float64
	fmt.Print("Masukkan No. HP: ")
	fmt.Scanln(&phoneNumber)

	var user Akun
	result := db.Where("no_hp = ?", phoneNumber).First(&user)
	if result.Error != nil {
		fmt.Println("Akun tidak ada.")
		return
	}

	fmt.Print("Masukkan jumlah top up: ")
	fmt.Scanln(&amount)

	user.Balance += amount
	result = db.Save(&user)
	if result.Error != nil {
		log.Fatalf("Failed to top up balance: %v", result.Error)
	}
	fmt.Println("Top-up berhasil.")
}

func transferBalance() {
	var from, to string
	var amount float64
	fmt.Print("Masukkan No. HP pengirim: ")
	fmt.Scanln(&from)

	fmt.Print("Masukkan No. HP penerima: ")
	fmt.Scanln(&to)

	fmt.Print("Masukkan jumlah transfer: ")
	fmt.Scanln(&amount)

	var userFrom, userTo Akun
	resultFrom := db.Where("no_hp = ?", from).First(&userFrom)
	resultTo := db.Where("no_hp = ?", to).First(&userTo)
	if resultFrom.Error != nil {
		fmt.Println("Akun pengirim tidak ada.")
		return
	}
	if resultTo.Error != nil {
		fmt.Println("Akun penerima tidak ada.")
		return
	}

	if userFrom.Balance < amount {
		fmt.Println("Balance tidak cukup.")
		return
	}

	userFrom.Balance -= amount
	userTo.Balance += amount

	tx := db.Begin()
	if err := tx.Error; err != nil {
		log.Fatalf("Error beginning transaction: %v", err)
	}

	var result = tx.Save(&userFrom)
	if result.Error != nil {
		tx.Rollback()
		log.Fatalf("Error saving sender account: %v", result.Error)
	}

	result = tx.Save(&userTo)
	if result.Error != nil {
		tx.Rollback()
		log.Fatalf("Error saving receiver account: %v", result.Error)
	}

	tx.Commit()
	fmt.Println("Transfer berhasil.")
}

func viewProfile() {
	var phoneNumber string
	fmt.Print("Masukan No. HP untuk melihat profil: ")
	fmt.Scanln(&phoneNumber)

	var user Akun
	result := db.Where("no_hp = ?", phoneNumber).First(&user)
	if result.Error != nil {
		fmt.Println("Akun tidak ada.")
		return
	}

	fmt.Println("No. HP:", user.NoHP)
	fmt.Println("Balance:", user.Balance)
}

func viewTransactionHistory() {
	var transactions []Transaksi
	result := db.Find(&transactions)
	if result.Error != nil {
		log.Fatalf("Failed to fetch transaction history: %v", result.Error)
	}
	fmt.Println("History Transaksi:")
	for i, transaction := range transactions {
		fmt.Println("Transaksi", i+1)
		fmt.Println("Dari:", transaction.Dari)
		fmt.Println("Ke:", transaction.Ke)
		fmt.Println("Jumlah:", strconv.FormatFloat(transaction.Jumlah, 'f', 2, 64))
	}
}
