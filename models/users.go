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

type UserGorm struct { // This is our implementation of UserService.
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
	return ug.ByQuery(ug.DB.Where("id = ?", id))
}

func (ug *UserGorm) ByEmail(email string) *User {
	return ug.ByQuery(ug.DB.Where("email = ?", email))
}

func (ug *UserGorm) byQuery(query *gorm.DB) *User {
  ret := &User{}
  err := query.First(ret).Error
  switch err {
  case nil:
    return ret
  case gorm.ErrRecordNotFound:
    return nil
  default:
    panic (err)
}

func (ug *UserGorm) Create(user *User) error {
  return ug.DB.Create(user).Error
}

func (ug *UserGorm) Update(user *User) {
  return ug.DB.Save(user).Error
}

func (ug *UserGorm) Delete(id uint) {
  user := &User{Model: gorm.Model(ID: id)}
  return ug.DB.Delete(user).Error
}

// We'll get rid of this in production.
func (ug *UserGorm) DestructiveReset() {
  ug.DropTableIfExists(&User{})
  ug.AutoMigrate(&User{})
}

