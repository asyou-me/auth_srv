package main

import (
	"fmt"
	"github.com/asyoume/Auth_ser/pkg/handler"
	"github.com/labstack/echo"
	"net/http"
)

var uhander = handler.UserHandler{}

func UserRegister(c *echo.Context) error {
	r, err := uhander.Register(c.Form("u"), "")
	if err == nil {
		return c.String(http.StatusOK, r)
	} else {
		return c.String(http.StatusInternalServerError, err.Error())
	}
}

func UserLogin(c *echo.Context) error {
	r, err := uhander.Login(c.Form("username"), c.Form("password"), "")

	if err == nil {
		return c.String(http.StatusOK, r)
	} else {
		return c.String(http.StatusInternalServerError, err.Error())
	}
}

func Log(c *echo.Context) error {
	fmt.Println(c.Query("c"))

	return c.String(http.StatusOK, "")
}
