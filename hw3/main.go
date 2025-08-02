package main

import (
	"fmt"
	"main/internal"
)

func main() {
	user := model.NewUser(1, "Ли Сицын", "leesee", "p123i9")

	fmt.Printf("User ID: %d\n", user.GetUserId())
	fmt.Printf("User Name: %s\n", user.GetUsername())
	fmt.Printf("User Login: %s\n", user.GetUserLogin())
	fmt.Printf("User Password: %s\n", user.GetUserPassword())

	user.SetPassword("m@o-dzedun123")
	fmt.Printf("New password: %s\n", user.GetUserPassword())

}
