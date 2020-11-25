package models

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	writer      *fiber.Ctx
}

func CreateDefaultResponse(c *fiber.Ctx) Response {
	return Response{
		Status:      http.StatusOK,
		contentType: "json",
		writer:      c,
	}
}

func SendSuccess(c *fiber.Ctx, data interface{}) error {
	response := CreateDefaultResponse(c)
	response.Data = data
	response.Message = "OK"
	return response.Send()
}

func SendUnprocessableEntity(c *fiber.Ctx) error {
	response := CreateDefaultResponse(c)
	return response.UnprocessableEntity()
}

func (this *Response) UnprocessableEntity() error {
	this.Message = "Unprocessable Entity"
	this.Status = http.StatusUnprocessableEntity
	return this.Send()
}

func SendNotFound(c *fiber.Ctx) error {
	response := CreateDefaultResponse(c)
	return response.NotFound()
}

func (this *Response) NotFound() error {
	this.Message = "Resource not found"
	this.Status = http.StatusNotFound
	return this.Send()
}

func SendNoContent(c *fiber.Ctx) error {
	response := CreateDefaultResponse(c)
	return response.NoContent()
}

func (this *Response) NoContent() error {
	this.Message = "No Content."
	this.Status = http.StatusNoContent
	return this.Send()
}

func SendConflict(c *fiber.Ctx) error {
	response := CreateDefaultResponse(c)
	return response.Conflict()
}

func (this *Response) Conflict() error {
	this.Message = "This resource already exists"
	this.Status = http.StatusConflict
	return this.Send()
}

func (this *Response) Send() error {
	this.writer.Status(this.Status)
	this.writer.Type(this.contentType)
	return this.writer.Send(formatJSON(this))
}

func formatJSON(this *Response) []byte {

	replaceSlashJSON := strings.NewReplacer(
		"\\", "",
		"\"{", "{",
		"}\"", "}",
		"u0026", "&",
	)
	outputByte, _ := json.MarshalIndent(this, "", " ")

	return []byte(replaceSlashJSON.Replace(string(outputByte)))
}
