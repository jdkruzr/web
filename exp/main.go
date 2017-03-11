package main

import (
  "fmt"
  // "bufio"
  // "os"
  
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
	Email	string `gorm:"not null;unique_index"`
  Orders []Order
}

type Order struct {
	gorm.Model
	UserID	uint
	Amount	int
	Description string
}

func createOrder(db *gorm.DB, user User, amount int, desc string) {
	db.Create(&Order{
	  UserID:	user.ID,
	  Amount: 	amount,
	  Description:  desc,
	})
  if db.Error != nil {
    panic(db.Error)
  }
}
	
func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
	  panic(err)
	}

  defer db.Close()
	
	db.LogMode(true)
	
	db.AutoMigrate(&User{}, &Order{})
	
	// name, email := getInfo()
	
	var user User
	db.Preload("Orders").First(&user)
	if db.Error != nil {
	  panic(db.Error)
	}
	
  //createOrder(db, user, 1001, "Fake Description #1")
  //createOrder(db, user, 9999, "Fake Description #2")
  //createOrder(db, user, 8800, "Fake Description #3")

  fmt.Println("Email:", user.Email)
  fmt.Println("Number of orders:", len(user.Orders))
  fmt.Println("Orders:", user.Orders)

	fmt.Printf("%+v\n", user)
}
