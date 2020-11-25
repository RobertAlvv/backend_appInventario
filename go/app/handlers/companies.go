package handlers

import (
	"../models"

	"github.com/gofiber/fiber/v2"
)

func GetCompany(c *fiber.Ctx) error {
	companie := models.GetCompanyByRnc(c.Params("rnc"))
	if companie.Rnc == "" {
		return models.SendNotFound(c)
	}
	return models.SendSuccess(c, companie)
}

func GetCompanies(c *fiber.Ctx) error {
	return models.SendSuccess(c, models.GetCompanies())
}
