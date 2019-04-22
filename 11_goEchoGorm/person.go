package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
)

type Person struct {
	gorm.Model
	Name string
	Age int
	Balance int
	Email string
	Address string
	Visible bool
}

// Return all people; provide method to sort by name or email
func (a *App) ListPeople(c echo.Context) error {
	sortOrder := c.QueryParam("sortBy")
	return c.String(http.StatusOK, "List people!  sort by " + sortOrder)
}

func (a *App) CreatePerson(c echo.Context) error {
	return c.String(http.StatusOK, "Creation Endpoint!")
}

func (a *App) DeletePerson(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Deletion Endpoint! id: " + id)
}

func (a *App) UpdatePerson(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Update Endpoint! id: " + id)
}

func (a *App) SeedPeople(c echo.Context) error {
	a.DB.Connection.Create(&Person{
		Name: "Ryan",
		Age: 32,
		Balance: 1000,
		Email: "ryanmccombe@gmail.com",
		Address: "Belfast",
		Visible: true,
	})

	return c.String(http.StatusOK, "Seeded!")
}