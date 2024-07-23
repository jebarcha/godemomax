package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("failed to read balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	accountBalance, err := getBalanceFromFile() //10000.0

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------")
		//panic(err)
		return
	}

	for i := 0; i < 20; i++ {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("your deposit: ")
			var depositoAmount float64
			fmt.Scan(&depositoAmount)

			if depositoAmount <= 0 {
				fmt.Println("Invalid amount. Must be greather than 0")
				continue
			}

			accountBalance += depositoAmount
			fmt.Println("Your new balance is", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("your withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greather than 0")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("You cannot withdraw an amount greater than your balance")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Your new balance is", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}
}
