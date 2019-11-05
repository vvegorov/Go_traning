package initDb

import "database/sql"

var Db *sql.DB

func InitDB() *sql.DB {
	connStr := "user=egorovvv password=12345 dbname=GO_Train_first sslmode=disable"
	Db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return Db
}
