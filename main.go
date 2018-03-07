package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func api(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " Ruta para /api ")
}
func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)
	mux.HandleFunc("/api", api)
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Corriendo en el puerto 8080")
	log.Fatal(server.ListenAndServe())

	// Conectar a base de datos y comprobar error

	db, err := sql.Open("mysql", "be7eddfee4e767: bdead22e @ /KIDADB") //cambiar por tus datos
	if err != nil {
		panic(err)
	}
	// db cerrar base de datos
	defer db.Close()

	// test para coneccion de base de datos
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

}
