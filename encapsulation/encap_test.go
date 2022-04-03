package encapsulation

import (
	"fmt"
	"testing"
)

func TestEncaptulation(t *testing.T) {
	user := NewUser("andiahmad", 25)
	fmt.Println(user.getName())
}
