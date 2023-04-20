package main

import (
	"fmt"
	f "forum"

	_ "github.com/mattn/go-sqlite3"
)

var url = "localhost:8080"

func main() {
	fmt.Println(f.HashPassword("ADMIN"))
}
