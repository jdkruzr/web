package main

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
)

const (
  host = "localhost"
  port = 5432
  user = "web"
  password = "supersecurepassword"
  dbname = "web_dev"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
	  panic(err)
	}

	err = db.Ping()
	if err != nil {
	  panic(err)
	}

	fmt.Println("Successfully connected!")
	// db.Close()

	_, err = db.Exec(`
		INSERT INTO users(name, email)
		VALUES($1, $2)`,
		"Butts", "butts@gmail.com")
	
	if err != nil {
		panic(err)
	}
	
	var id int
	row := db.QueryRow(`
		INSERT INTO users(name, email)
		VALUES($1, $2)
		RETURNING id`,
		"Bonk", "boop@asdf.gov")
	err = row.Scan(&id)
	if err != nil {
		panic(err)
	}
	
	var name, email string
	row = db.QueryRow(`
		SELECT id, name, email
		FROM users
		WHERE id=$1`, 1)
	err = row.Scan(&id, &name, &email)
	if err != nil {
		panic(err)
	}
	fmt.Println("ID:", id, "Name:", name, "Email:", email)
	

}