package handle

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"projectDBRest/src/ru/initDb"
	"projectDBRest/src/ru/utilSql"
	"strconv"
)

func GetListById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Converts the id parameter from a string to an int
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err == nil {
		log.Println("Get info about id #", id)

		target := utilSql.SqlQueryId(id, initDb.Db)
		k := target[0]
		fmt.Println(k)

		json.NewEncoder(w).Encode(target)

	} else {
		log.Fatal(err.Error())
	}

}

func GetInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	target := utilSql.SqlQueryAll(initDb.Db)

	json.NewEncoder(w).Encode(target)
}

func AddUserTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var x int
	//var userT UserTest
	_ = json.NewDecoder(r.Body).Decode(&utilSql.UserT)
	// x += 1
	x = rand.Intn(1000000)
	utilSql.UserT.Id = x

	utilSql.SqlAddUserTest(utilSql.UserT, initDb.Db)

	json.NewEncoder(w).Encode(utilSql.UserT)
}
