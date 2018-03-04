package main

import (
	"os"
	"net/http"
)
func home(w http.ResponseWriter, r *http.Request){
	http.ServerFile(w,r,"index.html")
}
func main() {
	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fileServer))

port := os.Getenv("PORT")

if port == "" {
	port = "8000"
}
http.HandleFunc("/",home)
http.ListenAndServe(":"+port,nil)
}
