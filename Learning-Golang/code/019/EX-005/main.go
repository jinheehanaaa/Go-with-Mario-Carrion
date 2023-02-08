package main

import "fmt"

type User struct {
	Name      string
	BirthYear uint
}

func main() {
	// 1
	//user := NewUser("mario", 1990)

	// 2
	user := User{
		Name:      "mario",
		BirthYear: 1900,
	}

	fmt.Println(user)
}

// 1
// NewUser ... (Unnecessary function)
// You can define initializer like this,
// but you don't need this
/*
func NewUser(name string, birthYear uint) User {
	return User{
		Name:      name,
		BirthYear: birthYear,
	}
}
*/

// Unnecessary function
/*
// BirthYear ...
func (u User) GetBirthYear() uint {
	return u.BirthYear
}
*/

// Unnecessary function
/*
// SetBirthYear ...
func (u *User) SetBirthYear(birthYear uint) {
	u.BirthYear = birthYear
}
*/
