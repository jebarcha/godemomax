package main

import (
	"bank/fileops"
	"fmt"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------")
		//panic(err)
		//return
	}

	for i := 0; i < 20; i++ {
		presentOptions()

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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}
}
