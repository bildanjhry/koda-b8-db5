package main

import (
	"contact-list/services"
	"fmt"
)

func main() {
	//response := services.CreateContact("Dona", "07372213")
	id := 14
	services.DeleteContact(&id)
	res := services.GetContactList()
	fmt.Println(res)
}
