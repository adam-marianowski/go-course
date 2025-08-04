package user

import (
	"errors"
	"fmt"
	"time"
)

func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "admin",
			lastName:  "adminay",
			birthDate: "02/02/2012",
			createdAt: time.Now(),
		},
	}
}

// constructor
func New(fName, lName, bDate string) (*User, error) {
	if fName == "" || lName == "" || bDate == "" {
		return nil, errors.New("all values are required")
	}

	return &User{
		firstName: fName,
		lastName:  lName,
		birthDate: bDate,
		createdAt: time.Now(),
	}, nil
}
