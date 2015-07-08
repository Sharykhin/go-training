package controllers

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Get(c *echo.Context) error {
	var data Response = Response{Status: 200, Message: "Hello World"}
	return c.JSON(http.StatusOK, data)
}
