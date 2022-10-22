package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type City struct {
	AID    int
	UserID string
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/batsa")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")

	results, err := db.Query("SELECT AID,User_ID FROM admins")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {

		var city City
		err := results.Scan(&city.AID, &city.UserID)

		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("%v\n", city)
	}
}
