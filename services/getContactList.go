package services

import (
	"contact-list/models"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func GetContactList() []models.User {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	err := godotenv.Load()
	if err != nil {
		panic("Can't access env")
	}
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("Can't connect db")
	} else {
		fmt.Println("Success connect db")
	}
	defer conn.Close(context.Background())

	rows, _ := conn.Query(context.Background(), `SELECT "id", "name", "user_contact"."phone" AS "phone" FROM "users"
	JOIN "user_contact" ON "user_contact"."id_user" = "users"."id";`)
	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		panic("Can't collect data from db")
	}
	fmt.Println(users)
	return users
}
