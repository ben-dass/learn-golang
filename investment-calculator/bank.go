package main

import "fmt"

func main() {
    //var accountBalance float64 = 1000

    printOptions()

    var choice int
    fmt.Print("\nSelected Option: ")
    fmt.Scan(&choice)

    if choice == 1 {

    }
}

func printOptions() {
    fmt.Println("\nWelcome to Go Bank!")
    fmt.Println("Please select an option:")
    fmt.Println("  1. Check Balance")
    fmt.Println("  2. Deposit")
    fmt.Println("  3. Withdraw")
    fmt.Println("  4. Exit")
}
