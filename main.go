package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	var dir string
	flag.StringVar(&dir, "dir", "./static", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	r.HandleFunc("/score/2020/qualification/{problem}", handleBooks).Methods("POST")
	r.HandleFunc("/score/2019/qualification/{problem}", handleSlideshow).Methods("POST")
	r.HandleFunc("/score/2018/qualification/{problem}", handleRides).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8080",
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	fmt.Println("Scoreboard listening on localhost:8080...")
	fmt.Println("1. GET problem files from 'localhost:8080/static/2020/qualification/a_example.txt'")
	fmt.Println("2. POST your solution string to 'localhost:8080/score/2020/qualification/a_example' to receive a score")

	log.Fatal(srv.ListenAndServe())
}
