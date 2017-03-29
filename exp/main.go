package main

import (
  "fmt"
  "github.com/jdkruzr/web/models"
)

const (
  host = "localhost"
  port = 5432
  user = "web"
  password = "supersecurepassword"
  dbname = "web_dev"
)

func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
  "password=%s dbname=%s sslmode=disable",
  host, port, user, password, dbname)

  ug, err := models.NewUserGorm(psqlInfo)
  if err != nil {
    panic(err)
  }

  ug.DestructiveReset()

  user := &models.User{
    Name: "Jon Calhoun",
    Email:  "jon@calhoun.io",
  }

  fmt.Println("Let's make a guy!")
  err = ug.Create(user)
  if err != nil {
    panic(err)
  }

  fmt.Println("Created this guy: ", user)

  fmt.Println("Retrieving user ID: ", user.ID)
  userByID := ug.ByID(user.ID)
  if userByID == nil {
    panic("No user found by ID " + user.ID + "!")
  }
  fmt.Println("Found the user: ", userByID)

  fmt.Println("Updating user: ", user)
  user.Email = "jarett@reticulum.us"
  if err := ug.Update(user); err != nil {
    panic(err)
  }
  fmt.Println("Updated user: ", user)
  
  fmt.Println("Retrieving the user w/ email:", user.Email)
  userByEmail := ug.ByEmail(user.Email)
  if userByEmail == nil {
    panic("No user found by email!")
  }
  fmt.Println("Found the user:", userByEmail)

  fmt.Println("Deleting the user...")
  if err := ug.Delete(user.ID); err != nil {
    panic(err)
  }

  //Verify that the user was deleted
  
}


