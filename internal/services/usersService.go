package services

import (
	"github.com/miguelToscano/newsletter/internal/repositories/usersRepository"
)

type UsersService struct {
	usersRepository usersRepository.UsersRepository
}

func NewUsersService(usersRepository usersRepository.UsersRepository) *UsersService {
	return &UsersService{usersRepository: usersRepository}
}

func (us *UsersService) CreateUser() (interface{}, error) {
	usersRepository.CreateUser()
	return nil, nil
}

func (us *UsersService) GetUser() (interface{}, error) {
	usersRepository.GetUser()
	return nil, nil
}

func (us *UsersService) GetUsers() (interface{}, error) {
	usersRepository.GetUsers()
	return nil, nil
}

func (us *UsersService) UpdateUser() (interface{}, error) {
	usersRepository.UpdateUser()
	return nil, nil
}

func (us *UsersService) DeleteUser() (interface{}, error) {
	usersRepository.DeleteUser()
	return nil, nil
}
