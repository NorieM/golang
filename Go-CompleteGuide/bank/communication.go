package main

import "fmt"

func showOptions() {
	var choices = [4]string{"Check balance", "Deposit money", "Withdraw money", "Exit"}
	for index, choice := range choices {
		fmt.Println(fmt.Sprintf("%v", index+1) + ". " + choice)
	}
}