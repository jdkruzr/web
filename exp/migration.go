package main

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
)

const (
  host = "localhost"
  port = 5432
  user = "web"
  password = "supersecurepassword"
  dbname = "web_dev"
)

type User struct {
	gorm.Model
	Name	string
	Email	string `gorm:"unique_index"`
}

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
	  panic(err)
	}
	
	db.LogMode(true)
	
	defer db.Close()
	
	db.AutoMigrate(&User{}) // So this makes a table based on what it sees in the User struct.
}
