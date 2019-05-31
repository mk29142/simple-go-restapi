package main_test

import (
    "bytes"
    "encoding/json"
    "github.com/joho/godotenv"
    "net/http"
    "net/http/httptest"
    "os"
    "simple-go-restapi/app"
    "simple-go-restapi/app/models"
    "simple-go-restapi/app/sqldriver"
    "testing"

    _ "github.com/go-sql-driver/mysql"
)

var a *app.App
var db *sqldriver.DB

func TestMain(m *testing.M) {
	a = &app.App{}
	a.Initialize()
	addTestDataToDB()

	code := m.Run()

	removeTestDataFromDB()
	os.Exit(code)
}

func TestGetUser(t *testing.T) {
    user := getUser(t)
    assertBody(&user, t)
}

func getUser(t *testing.T) models.User {
    req, _ := http.NewRequest("GET", "/user/test_name", nil)
    response := executeRequest(req)
    if response.Code != http.StatusOK {
        t.Fail()
    }
    user := models.User{}
    _ = json.NewDecoder(response.Body).Decode(&user)
    return user
}

func TestCreate(t *testing.T) {
    removeTestDataFromDB()
    jsonStr := []byte(`{"name":"test_name", "age":23, "school":"test_school"}`)

    req, _ := http.NewRequest("PUT", "/user", bytes.NewBuffer(jsonStr))
    response := executeRequest(req)

    if response.Code != http.StatusCreated {
        t.Fail()
    }

    user := getUser(t)
    removeTestDataFromDB()
    assertBody(&user, t)
}

func assertBody(user *models.User, t *testing.T) {
    if user.Name != "test_name" {
        t.Fail()
    }
    if user.Age != 23 {
        t.Fail()
    }
    if user.School != "test_school" {
        t.Fail()
    }
}


func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    a.Router.ServeHTTP(rr, req)
    return rr
}

func removeTestDataFromDB() {
    query := "DELETE FROM Users WHERE Name='test_name'"
    stmt, err := db.SQL.Prepare(query)
    if err != nil {
        panic("couldn't remove from db")
    }
    _, err = stmt.Exec()
    defer stmt.Close()
    if err != nil {
        panic("couldn't remove from db")
    }
}

func addTestDataToDB() {
    myEnv, _ := godotenv.Read(".env")

    DB_HOST := myEnv["DB_HOST"]
    DB_PORT := myEnv["DB_PORT"]
    DB_NAME := myEnv["DB_NAME"]
    DB_USERNAME := myEnv["DB_USERNAME"]
    DB_PASSWORD := myEnv["DB_PASSWORD"]

    connection, err := sqldriver.ConnectSQL(DB_HOST, DB_PORT, DB_USERNAME, DB_PASSWORD, DB_NAME)
    db = connection

    query := "INSERT INTO Users (Name, Age, School) VALUES (?, ?, ?)"

    stmt, err := connection.SQL.Prepare(query)
    if err != nil {
        panic("couldn't add into db")
    }

    _, err = stmt.Exec("test_name", 23, "test_school")
    defer stmt.Close()
    if err != nil {
        panic("couldn't add into db")
    }
}
