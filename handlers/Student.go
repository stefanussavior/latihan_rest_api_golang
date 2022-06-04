package handlers

import (
	"latihan_project_golang/REST_API_FIBER/database"
	"latihan_project_golang/REST_API_FIBER/model"

	"github.com/gofiber/fiber/v2"
)

func GetStudent(c *fiber.Ctx) error {
	var Student []model.Student

	database.Database.Find(&Student)
	return c.Status(200).JSON(Student)
}
