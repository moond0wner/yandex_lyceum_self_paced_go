// "Новый читатель"

package main

import "fmt"

type User struct {
	Name     string
	Age      int
	IsActive bool
}

func NewUser(name string, age int) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("name is empty for user")
	}
	if age == 0 {
		age = 18
	}
	return &User{Name: name, Age: age, IsActive: true}, nil

}
