package main

import (
    "simple-go-restapi/app"
)

func main() {
    app := &app.App{}
    app.Initialize()
    app.Run(":3000")
}

