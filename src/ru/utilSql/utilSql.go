package utilSql

import (
	"database/sql"
	"fmt"
	"log"
)

type UserTest struct {
	Id   int
	Name string
}

var UserT UserTest

func SqlQueryId(id int, db *sql.DB) []UserTest {

	rows, err := db.Query(`SELECT id, name FROM public.contact where id=$1`, id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var contacts = make([]UserTest, 0)
	var p UserTest
	// Добавляем с slice набор из таблицы public.contact
	for rows.Next() {

		err := rows.Scan(&p.Id, &p.Name)
		if err != nil {
			log.Fatal(err)
		}
		contacts = append(contacts, p)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return contacts
}

func SqlQueryAll(db *sql.DB) []UserTest {

	rows, err := db.Query(`SELECT id, name FROM public.contact`)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var contacts = make([]UserTest, 0)
	var p UserTest
	// Добавляем с slice набор из таблицы public.contact
	for rows.Next() {

		err := rows.Scan(&p.Id, &p.Name)
		if err != nil {
			log.Fatal(err)
		}
		contacts = append(contacts, p)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return contacts
}

func SqlAddUserTest(p UserTest, db *sql.DB) {

	result, err := db.Exec("insert into public.contact(id, name) VALUES ($1, $2)",
		p.Id, p.Name)
	if err != nil {
		panic(err)

	}

	fmt.Println(result.RowsAffected())

}
