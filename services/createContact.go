package services

import (
	"contact-list/utils"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreateContact(name string, number string) string {
	res := "Success create contact"
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
			res = "Failed Create contact"
		}
	}()

	utils.LoadEnv()

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("Can't connect db")
	}
	defer conn.Close(context.Background())

	rows, err := conn.Exec(context.Background(), `INSERT INTO "users" ("name") VALUES ($1)`, name)
	if err != nil || rows.RowsAffected() != 1 {
		panic("Failed create contact")
	} else {
		user := GetContactByName(name)

		contact, err := conn.Exec(context.Background(), `INSERT INTO "user_contact" ("id_user", "phone") VALUES ($1, $2)
		RETURNING "id_user" "phone"`, user.Id, number)
		if err != nil || contact.RowsAffected() != 1 {
			panic("Can't collect data from db")
		}
	}
	return res
}
