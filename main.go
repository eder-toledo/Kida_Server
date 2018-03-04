package main

import (
	"os"
	"net/http"
)
func home(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"index.html")
}
func main() {
	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fileServer))
	http.HandleFunc("/",home)

port := os.Getenv("PORT")

if port == "" {
	port = "8000"
}

http.ListenAndServe(":"+port,nil)
}
