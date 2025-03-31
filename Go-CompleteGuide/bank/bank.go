package main

import (
	"fmt"
)

func main() {

	var accountBalance = 1000.0

	var choices = [4]string{"Check balance", "Deposit money", "Withdraw money", "Exit"}
	fmt.Println("Welcome to the Bank of Job!")
	fmt.Println("How can I help you today?")

	for {
		for index, choice := range choices {
			fmt.Println(fmt.Sprintf("%v", index+1) + ". " + choice)
		}

		var choice int

		fmt.Print("Please choose from the above options:")

		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("You're balance is", accountBalance)
		} else if choice == 2 {
			var depositAmount float64
			fmt.Print("Please enter the amount you wish to deposit:")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount - deposit must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Deposit made. You're new balance is", accountBalance)
		} else if choice == 3 {
			var withdrawalAmount float64
			fmt.Print("Please enter the amount you wish to withdraw:")
			fmt.Scan(&withdrawalAmount)
			if withdrawalAmount <= 0 {
				fmt.Println("Invalid - withdrawal amount must be greater than 0.")
				continue
			}
			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficent funds, withdrawal can't exceed current account balance: ", accountBalance)
				continue
			}

			accountBalance += withdrawalAmount
			fmt.Println("Withdrawal made. You're new balance is", accountBalance)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}
	fmt.Println("Thanks for choosing to work with us today!")
}
