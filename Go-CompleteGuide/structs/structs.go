package main

import (
	"fmt"	
	"example.com/structs/user"
)



func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser,err:= user.New(firstName,lastName,birthdate)

	if err !=nil {
		fmt.Println(err.Error())
		return
	} 

	admin := user.NewAdmin("test@example.com", "password")

	admin.OutputUserDetails()

	admin.ClearUserName()

	admin.OutputUserDetails()

	

	// ... do something awesome with that gathered data!

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
