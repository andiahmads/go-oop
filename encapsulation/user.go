package encapsulation

/* Encapsulation juga sering digunakan karena terdapat fitur information-hiding mechanism.
Mekanisme ini menghilangkan akses publik ke atribut-atribut yang terdapat di dalam “kapsul” tersebut.
Metode ini dapat memudahkan kamu dalam mendefinisikan atribut apa saja yang dapat dibaca dan diperbarui.
*/

import "fmt"

type user struct {
	name string
	Age  int
}

func (u user) getName() string {
	return u.name
}

func (u user) getUserAndAge() string {
	return fmt.Sprintf("%s is %d years old", u.getName(), u.Age)
}

func NewUser(name string, age int) user {
	return user{name: name, Age: age}
}
