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

func (u User) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *User) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewUser(firstName , lastName, birthdate string) (*User, error ){

	if FirstName == "" || LastName == "" || Birthdate == "" {
		return nil, errors.New("first name, last name or birth date missing")
	}
	return User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
