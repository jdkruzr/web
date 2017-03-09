package main

import (
  "fmt"
  "bufio"
  "os"
  
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

func getInfo() (name, email string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, _ = reader.ReadString('\n')
	fmt.Println("What is your email?")
	email, _ = reader.ReadString('\n')
	return name, email
}

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
	  panic(err)
	}
	
	db.LogMode(true)
	
	defer db.Close()
	
	db.AutoMigrate(&User{})
	
	name, email := getInfo()
	
	u := &User{
		Name: name,
		Email: email,
	}
	if err = db.Create(u).Error; err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", u)
}
