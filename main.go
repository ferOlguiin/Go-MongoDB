package main

import (
	"dbconnection/database"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	database.ConnectDB()

}
