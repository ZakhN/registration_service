package rerpository

import "github.com/ZakhN/registration_service/models"

type PostgresRepo struct {
}

func (p PostgresRepo) GetUserByCredentials(password string, login string) (models.User, error) {
	return models.User{}, nil
}
