package main

import (
  "fmt"
)

const (
  host = "localhost"
  port = 5432
  user = "postgres"
  password = "your-password"
  dbname = "usegolang_dev"
)
func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
  "password=%s dbname=%s sslmode=disable",
  host, port, user, password, dbname)
}

