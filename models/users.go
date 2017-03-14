package models

import "github.com/jinzhu/gorm"

type User struct {
  gorm.Model
  Name string
  Email string `gorm:"not null;unique_index"`
}

type UserService interface {
  ByID(id uint) *User
  ByEmail(email string) *User
  Create(user *User)
  Update(user *User)
  Delete(id uint)
}


