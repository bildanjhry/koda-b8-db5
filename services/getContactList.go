package services

import (
	"contact-list/models"
	"contact-list/utils"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func GetContactList() []models.User {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	utils.LoadEnv()

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close(context.Background())

	rows, _ := conn.Query(context.Background(), `SELECT "id", "name", "user_contact"."phone" AS "phone" FROM "users"
	JOIN "user_contact" ON "user_contact"."id_user" = "users"."id";`)

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		panic("Can't collect data from db")
	}
	return users
}
