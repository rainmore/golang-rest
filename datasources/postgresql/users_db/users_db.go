package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	DBClient *sql.DB
)

func Init() {
	// databaseSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
	// 	"go_lang_example",
	// 	"{change-it}",
	// 	"localhost",
	// 	"go_lang_example",
	// )

	databaseSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		"localhost",
		5432,
		"go_lang_example",
		"{change-it}",
		"go_lang_example",
		"example",
	)

	var err error
	DBClient, err = sql.Open("postgres", databaseSourceName)

	if err != nil {
		panic(err)
	}

	Ping()

	log.Println("database successfully configured")
}

func Ping() {
	if err := DBClient.Ping(); err != nil {
		panic(err)
	}
}
