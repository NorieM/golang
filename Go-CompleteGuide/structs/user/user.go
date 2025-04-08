package user

import (
	"fmt"
	"time"
	"errors"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	Email string
	Password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		Email: email,
		Password: password,
		User: User {
			firstName: "",
			lastName: "",
			birthdate: "----",
			createdAt: time.Now(),
		},
	 }
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(FirstName , LastName, Birthdate string) (*User, error ){

	if FirstName == "" || LastName == "" || Birthdate == "" {
		return nil, errors.New("first name, last name or birth date missing")
	}
	return &User{
		firstName: FirstName,
		lastName:  LastName,
		birthdate: Birthdate,
		createdAt: time.Now(),
	}, nil
}
