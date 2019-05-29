package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"simple-go-restapi/app/handler"
	"simple-go-restapi/app/sqldriver"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	SQL *sql.DB
	Handler *handler.Handler
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize() {
	//dbName := os.Getenv("DB_NAME")
	//dbPass := os.Getenv("DB_PASS")
	//dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")

	connection, err := sqldriver.ConnectSQL("localhost", "3306", "root", "", "goDB")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	a.SQL = connection.SQL
	a.Router = mux.NewRouter()
	a.Handler = handler.NewHandler(connection)
	a.setRouters(a.Handler)
}

// setRouters sets the all required routers
func (a *App) setRouters(userHandler *handler.Handler) {
	// Routing for handling the projects
	a.Get("/users/{name}", a.handleRequest(userHandler.GetByID))

}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}
