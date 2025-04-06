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

func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewUser(FirstName , LastName, Birthdate string) (*User, error ){

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
