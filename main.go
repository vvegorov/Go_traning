package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"projectDBRest/src/ru/handle"
	_ "projectDBRest/src/ru/handle"
	"projectDBRest/src/ru/initDb"
	_ "projectDBRest/src/ru/initDb"
)

//var rows *sql.Rows

func main() {

	router := mux.NewRouter()
	buildBlockRoutes(router)

	initDb.Db = initDb.InitDB()
	defer initDb.Db.Close()

	log.Fatal(http.ListenAndServe(":8000", router))

}

func buildBlockRoutes(router *mux.Router) {
	prefix := "/books"
	router.HandleFunc(prefix, handle.GetInfo).Methods("GET")
	router.HandleFunc("/book"+"/{id}", handle.GetListById).Methods("GET")
	router.HandleFunc(prefix, handle.AddUserTest).Methods("POST")

}

