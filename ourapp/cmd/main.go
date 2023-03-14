package main

import (
	"flag"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":8080", "webserver port")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.ServeFile(w, r, "./ui/html/login.html")
	})

	router.POST("/login", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		email := r.FormValue("email")
		password := r.FormValue("password")

		// TODO: Verify email and password

		fmt.Fprintf(w, "Logged in as %s, %s", email, password)
	})

	infoLog.Printf("Server is starting on %s ", *addr)
	err := http.ListenAndServe(*addr, router)
	errorLog.Fatal(err)

}
