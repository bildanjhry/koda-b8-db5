package actions

import (
	"contact-list/services"
	"contact-list/ui"
	"contact-list/utils"
	"fmt"
	"os"
	"strconv"
)

var dis = ui.Print{}

func HomeMenu() {
	for {
		dis.HomeMenu()
		res, _ := utils.Io("\nInsert Input: ")
		switch res {
		case "1":
			utils.ClearTerm(0, "")
			contactList()
		case "2":
			utils.ClearTerm(0, "")
			createContact()
		case "x":
			utils.ClearTerm(1, "See ya")
			os.Exit(0)
		}

	}
}

func contactList() {
	for {
		dis.ContactList()
		res, _ := utils.Io("\nInsert Input: ")
		switch res {
		case "0":
			utils.ClearTerm(0, "")
			return
		case "1":
			utils.ClearTerm(0, "")
			deleteContact()
		case "2":
			utils.ClearTerm(0, "")
			editContact()
		default:
			utils.ClearTerm(1, "Invalid Input")
		}

	}
}

func createContact() {
	name, _ := utils.Io("\nInsert name :")
	phone, _ := utils.Io("Insert Phone :")
	res := services.CreateContact(name, phone)
	fmt.Println(res)
	_, err := utils.Io("\nPress enter to go back ")
	if err == nil {
		utils.ClearTerm(0, "")
		return
	}
}

func editContact() {
	data := dis.ContactListEdit()
	res, _ := utils.Io("\nInsert number: ")
	var id int
	name := ""
	for x, val := range data {
		if res == strconv.Itoa(x+1) {
			name = val.Name
			id = val.Id
		}
	}
	if name == "" {
		utils.ClearTerm(1, "Invalid Input")
		editContact()
		return
	}
	fmt.Printf("You select %s\n", name)
	newPhone, _ := utils.Io("\nInsert new Phone Number : ")
	fmt.Printf("\nName : %s", name)
	fmt.Printf("\nPhone : %s", newPhone)
	response, _ := utils.Io("\nSave new Phone Number? (y/n): ")
	if response == "y" {
		services.UpdateContact(&id, &newPhone)
		utils.ClearTerm(0, "")
		return
	} else {
		utils.ClearTerm(0, "")
		editContact()
	}
}

func deleteContact() {
	data := dis.ContactListDel()
	res, _ := utils.Io("\nInsert number: ")
	var id int
	name := ""
	for x, val := range data {
		if res == strconv.Itoa(x+1) {
			name = val.Name
			id = val.Id
		}
	}
	if name == "" {
		utils.ClearTerm(1, "Invalid Input")
		deleteContact()
		return
	}
	fmt.Printf("You select %s\n", name)
	response, _ := utils.Io("\nAre you sure to delete? (y/n): ")
	if response == "y" {
		services.DeleteContact(&id)
		utils.ClearTerm(0, "")
		return
	} else {
		utils.ClearTerm(0, "")
		deleteContact()
	}
}
