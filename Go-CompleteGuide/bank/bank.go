package main

import "fmt"
import "os"
import "strconv"

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText:=fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() float64  {
	data, _ :=os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _:=strconv.ParseFloat(balanceText, 64)
	return balance
}


func main() {

	var accountBalance = getBalanceFromFile()

	var choices = [4]string{"Check balance", "Deposit money", "Withdraw money", "Exit"}
	fmt.Println("Welcome to the Bank of Job!")
	fmt.Println("How can I help you today?")

	for {

		fmt.Println("\nPlease choose from the options below:")

		for index, choice := range choices {
			fmt.Println(fmt.Sprintf("%v", index+1) + ". " + choice)
		}

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
			fmt.Println("Withdrawal made. You're new balance is", accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing to work with us today!")
			return
		}

	}

}
