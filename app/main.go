package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/mgo.v2"
)

func main() {

	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")

	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	log.Println("database url:", dsn)

	session, err := mgo.Dial(dsn)
	if err != nil {
		log.Fatalln("erro na conexão com o mongodb:", err.Error())
	}
	defer session.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s %s%s %s", r.Method, r.Host, r.URL.Path, r.Proto)
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if err := session.Ping(); err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		} else {
			w.Write([]byte("OK"))
		}
	})

	log.Println("App rodando...")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
