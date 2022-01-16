package main

import (
	"github.com/ZakhN/registration_service/models"
	"github.com/ZakhN/registration_service/rerpository"
)

type Service struct {
	Repository rerpository.Repository
}

func (s Service) registerUser() (models.User, error) {
	return s.Repository.GetUserByCredentials("", "")
}
