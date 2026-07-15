package ui

import (
	"contact-list/models"
	"contact-list/services"
	"fmt"
)

type Print struct{}

func (u Print) HomeMenu() {
	fmt.Printf("================================\n")
	fmt.Println("          DASHBOARD")
	fmt.Printf("================================\n\n")
	fmt.Printf("(1) Contact List\n(2) Create Contact\n\n")
	fmt.Printf("================================\n")
	fmt.Printf("(x) Exit\n\n")
}

func (u Print) ContactList() {
	data := services.GetContactList()
	fmt.Printf("* ALL CONTACTS\n\n")
	for x, val := range data {
		fmt.Printf("%d. Name : %s\n", x+1, val.Name)
		fmt.Printf("   Phone: %s\n\n", val.Phone)
	}
	fmt.Printf("\n================================\n")
	fmt.Printf("(0) Back   (1) Delete   (3) Edit\n\n")
}

func (u Print) ContactListDel() []models.User {
	data := services.GetContactList()
	fmt.Printf("* DELETE CONTACT\n\n")
	for x, val := range data {
		fmt.Printf("(%d) Name : %s\n", x+1, val.Name)
		fmt.Printf("   Phone: %s\n\n", val.Phone)
	}
	fmt.Printf("\n================================\n")
	return data
}
