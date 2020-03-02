package main

import (
	"fmt"
	"net/http"

	"github.com/jeremy-miller/lenslocked.com/models"

	"github.com/jeremy-miller/lenslocked.com/controllers"

	"github.com/gorilla/mux"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	//us.DestructiveReset()
	us.AutoMigrate()

	static := controllers.NewStatic()
	users := controllers.NewUsers(us)

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.Handle("/", static.Home).Methods("GET")
	r.Handle("/contact", static.Contact).Methods("GET")
	r.Handle("/signup", users.NewView).Methods("GET")
	r.HandleFunc("/signup", users.Create).Methods("POST")
	r.Handle("/login", users.LoginView).Methods("GET")
	r.HandleFunc("/login", users.Login).Methods("POST")

	http.ListenAndServe(":3000", r)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "<h1>Sorry, but we couldn't find the page you were looking for</h1>")
}
