package main

import (
	"fmt"
	"example.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err.Error())
		fmt.Println("----")
		panic("Cannot continue with Bank App!")
	}

	fmt.Println("Welcome to the Bank of Job!")
	fmt.Println("How can I help you today?")
	
	
	for {
		
		fmt.Println("\nPlease choose from the options below:")
		
		showOptions()	

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("You're balance is", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Please enter the amount you wish to deposit:")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount - deposit must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("\nDeposit made. You're new balance is", accountBalance)
		case 3:
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

			accountBalance -= withdrawalAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Withdrawal made. You're new balance is", accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing to work with us today!")
			return
		}

	}

}

