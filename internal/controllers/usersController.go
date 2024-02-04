package controllers

import "github.com/gin-gonic/gin"

type UsersController struct {
}

func NewUsersController() *UsersController {
	return &UsersController{}
}

func (u *UsersController) CreateUser() (interface{}, error) {
	return nil, nil
}

func (u *UsersController) GetUser() (interface{}, error) {
	return nil, nil
}

func (u *UsersController) GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func (c *UsersController) UpdateUser() (interface{}, error) {
	return nil, nil
}

func (c *UsersController) DeleteUser() (interface{}, error) {
	return nil, nil
}
