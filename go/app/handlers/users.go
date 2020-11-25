package handlers

import (
	"strconv"

	"../middleware"
	"../models"
	"github.com/gofiber/fiber/v2"
)

func LoginUser(c *fiber.Ctx) error {
	user := new(models.User)
	err := c.BodyParser(user)
	if err != nil {
		return models.SendUnprocessableEntity(c)
	}
	user.GetUserID()
	if user.Id == 0 {
		return models.SendNotFound(c)
	}
	user.TokenJWT = middleware.SignToken("tokenKey", string(user.Id))
	return models.SendSuccess(c, user)
}

func SaveUser(c *fiber.Ctx) error {
	user := new(models.User)
	err := c.BodyParser(user)
	if err != nil {
		return models.SendUnprocessableEntity(c)
	}
	exist, _ := user.ExistUsername()

	if exist == true {
		return models.SendConflict(c)
	}

	user.Save()
	return models.SendSuccess(c, user)
}

func GetUsers(c *fiber.Ctx) error {
	return models.SendSuccess(c, models.GetUsers())
}

func GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return models.SendUnprocessableEntity(c)
	}
	user := models.GetUserById(int64(id))

	if user.Id == 0 {
		return models.SendNotFound(c)
	}
	return models.SendSuccess(c, user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return models.SendUnprocessableEntity(c)
	}
	rowsAffect := models.DeleteUserById(int64(id))

	if rowsAffect == 0 {
		return models.SendNotFound(c)
	}

	return models.SendNoContent(c)
}

func UpdateUser(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	user := models.GetUserById(int64(id))
	if user.Id == 0 {
		return models.SendNotFound(c)
	}

	userRequest := new(models.User)
	err = c.BodyParser(userRequest)
	if err != nil {
		return models.SendUnprocessableEntity(c)
	}

	user.Username = userRequest.Username
	user.Password = userRequest.Password
	user.Save()
	return models.SendSuccess(c, user)
}
