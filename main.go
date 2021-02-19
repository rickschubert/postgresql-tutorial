package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "turm"
	password = "admin"
	dbname   = "koffer"
)

func main() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	// insert
	// hardcoded
	// insertStmt := `insert into "accounts"("user_id", "username", "password", "email", "created_on", "last_login") values(1, 'John', 'pass', 'dasfafd@gmx.de', '2004-10-19 10:23:54', '2004-10-19 10:23:54')`
	// _, e := db.Exec(insertStmt)
	// CheckError(e)

	// dynamic
	// insertDynStmt := `insert into "accounts"("user_id", "username", "password", "email", "created_on", "last_login") values($1,$2,$3,$4,$5,$6)`
	// _, e := db.Exec(insertDynStmt, 2, "dddddd", "pass", "dasfafddddd@gmx.de", "2004-10-19 10:23:54", "2004-10-19 10:23:54")
	// CheckError(e)

	// update
	// updateStmt := `update "accounts" set "username"=$1, "password"=$2 where "user_id"=$3`
	// _, e := db.Exec(updateStmt, "Mary", "new password", 2)
	// CheckError(e)

	rows, err := db.Query(`SELECT "username", "password" FROM "accounts"`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var username string
		var password string

		err = rows.Scan(&username, &password)
		CheckError(err)

		fmt.Println(username, password)
	}

	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
