package main
import "fmt"

func main() {
	var password string
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)
	var hasUpper, hasLower, hasDigit bool
	var length, i int
	var ch byte			