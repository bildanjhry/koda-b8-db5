package services

import (
	"contact-list/utils"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func UpdateContact(id *int, phone *string) {
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
	commentTag, err := conn.Exec(context.Background(), `UPDATE "user_contact" SET phone=$2 WHERE id_user=$1`,
		id, phone)
	if err != nil {
		panic(err)
	}
	if commentTag.RowsAffected() != 1 {
		panic("No Row Found")
	}
}
