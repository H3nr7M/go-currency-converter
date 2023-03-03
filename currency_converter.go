package main
// strconv is a package in the Go standard library that implements conversions to and from string representations of basic data types.
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

//const is a keyword in Go that declares a constant value. The value of a constant must be known at compile time.
const (
    exchangeRate = 0.84 // exchange rate from USD to EUR
)

func main() {
	// for read input from the user
    scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the currency converter!")
    fmt.Println("Enter an amount in USD to convert to EUR:")
	// Scan() reads the next token from the input, which is space-delimited by default.
    scanner.Scan()
	//usdStr is a string
    usdStr := scanner.Text()

    usd, err := strconv.ParseFloat(usdStr, 64) //64 is the bit size
	//if err is not nil, that is a word for null in Go
    if err != nil {
		//Println is a function in the fmt package that prints a line to standard output.
        fmt.Println("Invalid input:", err)
        return
    }

    eur := usd * exchangeRate

    fmt.Printf("%.2f USD is equal to %.2f EUR\n", usd, eur)
}
