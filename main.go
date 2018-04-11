package main

import (
	"encoding/json"
	"log"
	"net/http"

	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

/*
var db *sql.DB
*/
type User struct {
	id        string    `json:"id"`
	FirstName string    `json:"firstname,omitempty"`
	LastName  string    `json:"lastname,omitempty"`
	UserName  string    `json:"username,omitempty"`
	CreateAt  time.Time `json:"createat"`
}

// ListUser almacena a todos los usuarios
var Listusers = make(map[string]User)
var id int

func main() {
	/*GetConnection()*/
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/User", GetUsersHandler).Methods("GET")
	r.HandleFunc("/api/User", PostUserHandler).Methods("POST")
	r.HandleFunc("/api/User{id}", PutUserHandler).Methods("PUT")
	r.HandleFunc("/api/User{id}", DeleteUserHandler).Methods("DELETE")
	/*CloseConnection()*/
	log.Fatal(http.ListenAndServe(":8080", r))
}

//GetUsersHandler lista todos los usuarios
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []User
	for _, v := range Listusers {
		users = append(users, v)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//PostUserHandler crea un usuario en la base de datos
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	user.CreateAt = time.Now()
	id++
	k := strconv.Itoa(id)
	Listusers[k] = user
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

//PutUserHandler Actualiza un usuario en base al id
func PutUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	k := params["id"]
	var userupdate User
	err := json.NewDecoder(r.Body).Decode(&userupdate)
	if err != nil {
		panic(err)
	}
	if user, ok := Listusers[k]; ok {
		userupdate.CreateAt = user.CreateAt
		delete(Listusers, k)
		Listusers[k] = userupdate
	} else {
		log.Printf("No encontramos el id %s", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

//DeleteUserHandler elimina un usuario en base al id
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	k := params["id"]

	if _, ok := Listusers[k]; ok {

		delete(Listusers, k)

	} else {
		log.Printf("No encontramos el id %s", k)
	}
	w.WriteHeader(http.StatusNoContent)

}

/*
//GetConnection genera la conexion ala BD
func GetConnection() {
	if connection, err := sql.Open("mysql", "root:@tcp(localhost:3303)/kida"); err != nil {
		panic(err.Error())
	} else {
		db = connection
	}
	fmt.Println("Conexion exitosa")
}

//CloseConnection cierra la conexion ala DB
func CloseConnection() {
	db.Close()
}
*/
