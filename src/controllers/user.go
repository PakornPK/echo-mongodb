package controllers

import (
	"net/http"

	"github.com/echo_CRUD/src/database"
	"github.com/echo_CRUD/src/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllUser(c echo.Context) error {
	client := database.Client
	users, err := models.ReadAllUser(client)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	client := database.Client
	data := bson.D{
		{Key: "name", Value: "pakorn"},
		{Key: "last_name", Value: "R."},
	}
	models.InsertUser(client, data)
	return c.String(http.StatusOK, "OK")
}
