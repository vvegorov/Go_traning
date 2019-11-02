package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var db *sql.DB
var target []UserTest

type UserTest struct {
	Id   int
	Name string
}

//var rows *sql.Rows

func main() {

	router := mux.NewRouter()
	buildBlockRoutes(router)

	log.Fatal(http.ListenAndServe(":8000", router))

}

func buildBlockRoutes(router *mux.Router) {
	prefix := "/books"
	router.HandleFunc(prefix, GetInfo).Methods("GET")
	router.HandleFunc(prefix+"/{id}", GetListById).Methods("GET")
	router.HandleFunc(prefix, AddUserTest).Methods("POST")

}

func GetListById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Converts the id parameter from a string to an int
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err == nil {
		log.Println("Get info about id #", id)

		db := InitDB()
		defer db.Close()

		target := SqlQueryId(id, db)
		k := target[0]
		fmt.Println(k)

		json.NewEncoder(w).Encode(target)

	} else {
		log.Fatal(err.Error())
	}

}

func GetInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := InitDB()
	defer db.Close()

	target := SqlQueryAll(db)

	json.NewEncoder(w).Encode(target)
}

func AddUserTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userT UserTest
	var x int
	_ = json.NewDecoder(r.Body).Decode(&userT)
	// x += 1
	x = rand.Intn(1000000)
	userT.Id = x

	db := InitDB()
	defer db.Close()

	SqlAddUserTest(userT, db)

	json.NewEncoder(w).Encode(userT)
}

func InitDB() *sql.DB {
	connStr := "user=egorovvv password=12345 dbname=GO_Train_first sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return db

}

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
