package rerpository

import "github.com/ZakhN/registration_service/models"

type Repository interface {
	GetUserByCredentials(password string, login string) (models.User, error)
}
