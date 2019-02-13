package models

import (
	"gopkg.in/mgo.v2/bson"
)

// User Struct
type User struct {
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
	ID     bson.ObjectId `json:"id" bson:"_id"`
}

// UserManagerInterface : user manager methods
type UserManagerInterface interface {
	GetUser(id string) (*User, error)
	GetAllUsers() ([]*User, error)
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(user *User) error
}

// UserRepositoryInterface : user repository methods
type UserRepositoryInterface interface {
	FindUserById(id string) (*User, error)
	FindAllUsers() ([]*User, error)
	Insert(user *User) error
	Update(user *User) error
	Delete(id string) error
	IsNotFoundErr(err error) bool
	IsAlreadyExistErr(err error) bool
}
