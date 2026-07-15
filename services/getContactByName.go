package services

import (
	"contact-list/models"
	"contact-list/utils"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func GetContactByName(name string) models.UserName {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	utils.LoadEnv()

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("Can't connect db")
	}
	defer conn.Close(context.Background())

	rows, _ := conn.Query(context.Background(), `SELECT "id", "name" FROM "users" WHERE name=$1`,
		name)
	users, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.UserName])
	if err != nil {
		panic("Can't collect data from db users")
	}
	return users
}
