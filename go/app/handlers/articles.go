package handlers

import (
	"fmt"
	"strconv"

	"../models"
	"github.com/gofiber/fiber/v2"
)

func GetArticles(c *fiber.Ctx) error {
	company_rnc := c.Params("rnc")
	branch_office, _ := strconv.Atoi(c.Params("branch_office"))
	return models.SendSuccess(c, models.GetArticlesByBranchOffices(company_rnc, branch_office))
}

func GetArticle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return models.SendUnprocessableEntity(c)
	}
	article := models.GetArticleById(int64(id))

	if article.Id == 0 {
		return models.SendNotFound(c)
	}
	return models.SendSuccess(c, article)
}

func SaveArticle(c *fiber.Ctx) error {
	article := models.ParseData(c)
	fmt.Println(article)
	if article.Name == "" || article.Id_Branch_OfficesVSDepartments == 0 {
		return models.SendUnprocessableEntity(c)
	}
	article.Save()
	return models.SendSuccess(c, article)
}

func UpdateArticle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return models.SendUnprocessableEntity(c)
	}
	article := models.GetArticleById(int64(id))

	if article.Id == 0 {
		return models.SendNotFound(c)
	}

	articleRequest := models.ParseData(c)

	article.Name = articleRequest.Name
	article.Description = articleRequest.Description
	article.Feature = articleRequest.Feature
	article.Status = articleRequest.Status
	article.Id_Type_article = article.Id_Type_article
	article.Id_Branch_OfficesVSDepartments = articleRequest.Id_Branch_OfficesVSDepartments

	article.Save()
	return models.SendSuccess(c, article)
}

func DeleteArticle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return models.SendUnprocessableEntity(c)
	}

	rowsAffect := models.DeleteArticleById(int64(id))
	if rowsAffect == 0 {
		return models.SendNotFound(c)
	}

	return models.SendNoContent(c)
}

// func GetArticlesByCompany(c *fiber.Ctx) error {
// 	rnc := c.Params("rnc")
// 	companie := models.GetCompanyByRnc(rnc)

// 	if companie.Rnc == "" {
// 		return models.SendNotFound(c)
// 	}

// 	return models.SendSuccess(c, models.GetArticleByCompany(rnc))
// }

// func GetArticlesByBranchOffices(c *fiber.Ctx) error {
// 	id_ := c.Params("id")
// 	return nil
// }
