package app

import (
    "fmt"
    "github.com/joho/godotenv"
    "log"
    "net/http"
    "os"
    "simple-go-restapi/app/handler"
    "simple-go-restapi/app/sqldriver"

    _ "github.com/go-sql-driver/mysql"

    "github.com/gorilla/mux"
)

type App struct {
    Router *mux.Router
    Handler *handler.Handler
}

func (a *App) Initialize() {
    myEnv, err := godotenv.Read()

    DB_HOST := myEnv["DB_HOST"]
    DB_PORT := myEnv["DB_PORT"]
    DB_NAME := myEnv["DB_NAME"]
    DB_USERNAME := myEnv["DB_USERNAME"]
    DB_PASSWORD := myEnv["DB_PASSWORD"]

    connection, err := sqldriver.ConnectSQL(DB_HOST, DB_PORT, DB_USERNAME, DB_PASSWORD, DB_NAME)
    if err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }

    a.Router = mux.NewRouter()
    a.Handler = handler.NewHandler(connection)
    a.setRouters()
}

func (a *App) setRouters() {
    a.Get("/user/{name}", a.Handler.GetByID)
    a.Put("/user", a.Handler.Create)
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
    a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
    a.Router.HandleFunc(path, f).Methods("PUT")
}

func (a *App) Run(host string) {
    log.Fatal(http.ListenAndServe(host, a.Router))
}
