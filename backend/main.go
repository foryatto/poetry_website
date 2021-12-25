package main

import (
	"backend/app"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app.Run()
}
