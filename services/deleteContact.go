package services

import (
	"contact-list/utils"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func DeleteContact(id *int) {
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
	commentTag, err := conn.Exec(context.Background(), `DELETE FROM "user_contact" WHERE id_user=$1`,
		id)
	if err != nil {
		panic(err)
	}
	commentTagUser, err := conn.Exec(context.Background(), `DELETE FROM "users" WHERE id=$1`,
		id)
	if err != nil {
		panic(err)
	}
	if commentTag.RowsAffected() != 1 || commentTagUser.RowsAffected() != 1 {
		panic("No Row Found")
	}
}
