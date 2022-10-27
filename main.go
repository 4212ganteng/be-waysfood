package main

import (
	"fmt"
	"go-batch2/database"
	"go-batch2/pkg/mysql"
	"go-batch2/routes"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// env

	env := godotenv.Load()
	if env != nil {
		panic("failed to load env file")
	}

	mysql.DatabaseInit()

	database.RunMigration()

	r := mux.NewRouter()

	routes.RoutesInit(r.PathPrefix("/api/v1").Subrouter())

	//path file
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	fmt.Println("server running on port 8000")
	http.ListenAndServe("127.0.0.1:8000", r)
}
