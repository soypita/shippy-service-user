package main

import (
	"context"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	pb "github.com/soypita/shippy-service-user/proto/user"
)

type User struct {
	Id       string
	Name     string
	Company  string
	Email    string
	Password string
}

func MarshalUser(user *pb.User) *User {
	return &User{
		Id:       user.Id,
		Name:     user.Name,
		Company:  user.Company,
		Email:    user.Email,
		Password: user.Password,
	}
}

func UnmarshalUser(user *User) *pb.User {
	return &pb.User{
		Id:       user.Id,
		Name:     user.Name,
		Company:  user.Company,
		Email:    user.Email,
		Password: user.Password,
	}
}

func UnmarshalUsers(users []*User) []*pb.User {
	res := make([]*pb.User, 0)
	for _, val := range users {
		res = append(res, UnmarshalUser(val))
	}
	return res
}

type Repository interface {
	GetAll(ctx context.Context) ([]*User, error)
	Get(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *pb.User) error
	GetByEmailAndPassword(ctx context.Context, user *User) (*User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("Id", uuid.String())
}

func (ur *UserRepository) GetAll(ctx context.Context) ([]*User, error) {
	var users []*User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) Get(ctx context.Context, id string) (*User, error) {
	var user *User
	user.Id = id
	if err := ur.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) GetByEmailAndPassword(user *User) (*User, error) {
	if err := ur.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) Create(user *User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
