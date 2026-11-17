package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func db() {
	// 1. Koneksi ke Database SQLite
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	// 2. AutoMigrate untuk memastikan tabel 'users' siap
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("Gagal membuat tabel: %v", err)
	}

	// 3. Buka file data.txt lama
	txtFile := "user_data.txt"
	file, err := os.Open(txtFile)
	if err != nil {
		log.Fatalf("Gagal membuka file %s: %v", txtFile, err)
	}
	defer file.Close()

	var usersToMigrate []User
	scanner := bufio.NewScanner(file)
	lineNum := 0

	// 4. Parsing isi file baris demi baris
	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // Abaikan baris kosong
		}

		parts := strings.Split(line, ",")
		if len(parts) < 3 {
			fmt.Printf(" Baris %d diabaikan (format tidak sesuai): %s\n", lineNum, line)
			continue
		}

		user := User{
			Username: strings.TrimSpace(parts[0]),
			Nama:     strings.TrimSpace(parts[1]),
			Password: strings.TrimSpace(parts[2]),
		}

		usersToMigrate = append(usersToMigrate, user)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error saat membaca file: %v", err)
	}

	if len(usersToMigrate) == 0 {
		fmt.Println("Tidak ada data baru yang ditemukan untuk dimigrasi.")
		return
	}

	// 5. Simpan ke database menggunakan GORM (Clause OnConflict untuk menghindari error duplikat)
	// 5. Simpan ke database. Gunakan OnConflict DoNothing supaya baris duplikat
	// (Username punya unique) tidak memicu error, melainkan dilewati.
	var inserted, skipped int
	for _, user := range usersToMigrate {
		result := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)
		if result.Error != nil {
			log.Printf("Gagal memigrasikan user %s: %v", user.Username, result.Error)
			skipped++
			continue
		}
		if result.RowsAffected > 0 {
			inserted++
		} else {
			skipped++
		}
	}

	if inserted == 0 && skipped == 0 {
		fmt.Println("Tidak ada data baru yang ditemukan untuk dimigrasi.")
		return
	}

	fmt.Printf(" Berhasil memindahkan %d data dari %s ke SQLite (app.db)!\n", inserted, txtFile)
}
