package oopdasar01

import (
	"fmt"
	"testing"
)

type user struct {
	name string
}

func (u user) getName() string {
	return u.name
}

func NewUser(name string) user {
	return user{name: name}
}

func TestUser(t *testing.T) {
	user := NewUser("andiahmads")
	showUser := user.getName()
	fmt.Println(showUser)
}
