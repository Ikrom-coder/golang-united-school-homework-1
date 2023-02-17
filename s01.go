package structs

import (
	"reflect"
)

type User struct {
	lastName  string
	firstName string
}

func New() User {
	return User{}
}

type UserInterface interface {
	SetFirstName(string)
	SetLastName(string)
	FullName() string
}

func (u *User) SetFirstName(s string) {
	u.firstName = s
}

func (u *User) SetLastName(s string) {
	u.lastName = s
}

func (u *User) FullName() string {
	return u.lastName + " " + u.firstName
}

func ResetUser(user *User) {
	user.SetLastName("")
	user.SetFirstName("")
}

func IsUser(user interface{}) bool {
	if reflect.TypeOf(user) == reflect.TypeOf(User{}) {
		return true
	}
	return false
}

func ProcessUser(user UserInterface) string {
	fName := "qq"
	lName := "ww"
	user.SetFirstName(fName)
	user.SetLastName(lName)
	return user.FullName()
}
