package polymorphism


/* Prinsip polymorphism pada OOP adalah konsep di mana suatu objek berbeda-beda dapat diakses melalui satu interface. 
Sebuah objek polymorphic dapat beradaptasi dengan metode apapun yang diimplementasikan pada objek tersebut, 
dan setiap class memiliki interpretasinya tersendiri terhadap interfacenya.
 */

import (
	"fmt"
	"testing"
)

type user struct {
	name string
	Age  int
}

// interface class user
type userService interface {
	getUserAge() string
}

func (u user) getName() string {
	return u.name
}

func (u user) getUserAge() string {
	return fmt.Sprintf("%s is %d years old", u.getName(), u.Age)
}

//implementasikan interface userService
func NewUser(name string, age int) userService {
	return user{name: name, Age: age}
}

func TestPolymorphim(t *testing.T) {
	user := NewUser("andiahmad", 25)
	fmt.Println(user.getUserAge())
}
