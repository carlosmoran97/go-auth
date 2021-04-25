package controllers

import (
	"io"
	"os"

	"github.com/carlosmoran97/go-auth/database"
	"github.com/carlosmoran97/go-auth/models"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		io.WriteString(os.Stderr, err.Error())
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&user)

	c.JSON(user)

}

func Login(c *fiber.Ctx) {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		io.WriteString(os.Stderr, err.Error())
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		c.JSON(fiber.Map{
			"message": "user not found",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(fiber.Map{
			"message": "incorrect password",
		})
		return
	}

	c.JSON(user)
}
