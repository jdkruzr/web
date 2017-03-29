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
  Create(user *User) error
  Update(user *User) error
  Delete(id uint) error
}

type UserGorm struct {
  *gorm.DB
}

func NewUserGorm(connectionInfo string) (*UserGorm, error) {
  db, err := gorm.Open("postgres", connectionInfo)
  if err != nil {
    return nil, err
  }
  return &UserGorm{db}, nil
}

func (ug *UserGorm) ByID(id uint) *User {
  ret := &User{}
  err := ug.DB.First(ret, id).Error
  switch err {
  case nil:
    return ret
  case gorm.ErrRecordNotFound:
    return nil
  default:
    panic (err)
  }
}

func (ug *UserGorm) ByEmail(email string) *User {
  return nil
}

func (ug *UserGorm) Create(user *User) error {
  return ug.DB.Create(user).Error
}

func (ug *UserGorm) Update(user *User) {
  return nil
}

func (ug *UserGorm) Delete(id uint) {
  return nil
}

// We'll get rid of this in production.
func (ug *UserGorm) DestructiveReset() {
  ug.DropTableIfExists(&User{})
  ug.AutoMigrate(&User{})
}

