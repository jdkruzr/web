package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  )

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

type UserGorm struct {
  *gorm.DB
}

func NewUserGorm(connectionInfo string) (*UserGorm, error) {
  db, err := gorm.Open("postgres", connectionInfo)
  if err != nil {
    return nil, err
  }
  return &UserGorm(db), nil
}

func (ug *UserGorm) ByID(id uint) *User {
  return nil
}

func (ug *UserGorm) ByEmail(email string) *User {
}

func (ug *UserGorm) Create(user *User) {
}

func (ug *UserGorm) Update(user *User) {
}

func (ug *UserGorm) Delete(id uint) {
}


