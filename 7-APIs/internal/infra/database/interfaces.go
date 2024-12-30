package database

import "github.com/AndreCDiniz/PosGoExpert/7-APIs/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
