package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)


type UserTest struct{
	id int
	name string
}

func main() {

	//Устанавливаем соединение GO_Train_first
	connStr := "user=egorovvv password=12345 dbname=GO_Train_first sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Устанавливаем соединение GO_Train_target
	connStr1 := "user=egorovvv password=12345 dbname=GO_Train_target sslmode=disable"
	db_tg, err_tg := sql.Open("postgres", connStr1)
	if err_tg != nil {
		panic(err_tg)
	}
	defer db_tg.Close()



	rows, err := db.Query("select * from public.contact")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	contacts := []UserTest{}

	// Добавляем с slice набор из таблицы public.contact
	for rows.Next(){
		p := UserTest{}
		err := rows.Scan(&p.id, &p.name)
		if err != nil{
			fmt.Println(err)
			continue
		}
		contacts = append(contacts, p)
	}

	// Вывод содержимого таблицы public.contact из GO_Train_first
	for _, p := range contacts{
		fmt.Println(p.id, p.name)
	}


	for _, p := range contacts{

		result, err := db_tg.Exec("insert into public.contact_target(id, name) VALUES ($1, $2)",
			p.id, p.name)
		if err != nil{
			panic(err)
			continue
		}

		fmt.Println(result.RowsAffected())

	}


}
