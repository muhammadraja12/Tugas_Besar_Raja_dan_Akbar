package main

import "fmt"

func main() {

		//Input password dari user
		var password string
		fmt.Print("Masukkan password: ")
		fmt.Scan(&password)
 
		// Inisialisasi variabel
		var hasUpper, hasLower, hasDigit bool 
		var length, i int 
		var ch byte

		// Set awal variabel
		hasUpper = false
		hasLower = false
		hasDigit = false

		// Karakter yang ada pada password
		for i = 0; i < len(password); i++ { 
			ch = password[i]
			length++

			//cek tipe karakter
			if ch >= 'A' && ch <= 'Z' {
				hasUpper = true
			} else if ch >= 'a' && ch <= 'z' {
				hasLower = true
			} else if ch >= '0' && ch <= '9' {
				hasDigit = true
			} 
		}

			
		// Requirment penilaian kekuatan password
		if length < 8 {
			fmt.Println("Password minimal 8 karakter.")
		} else if hasUpper && hasLower && hasDigit {
			fmt.Println("Password KUAT.")
		} else if (hasUpper && hasLower) || (hasUpper && hasDigit) || (hasLower && hasDigit) {
			fmt.Println("Password SEDANG.")
		} else {
			fmt.Println("Password LEMAH.")
		}
	}

