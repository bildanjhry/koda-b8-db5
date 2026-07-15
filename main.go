package main

import (
	"contact-list/actions"
	"contact-list/utils"
)

func main() {
	utils.ClearTerm(0, "")
	actions.HomeMenu()
}
