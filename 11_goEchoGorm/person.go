package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
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

func (a *App) ListPeople(c echo.Context) error {
	sortOrder := "name"
	if c.QueryParam("sortBy") == "email" {
		sortOrder = "email"
	}

	people := []Person{}
	a.DB.Connection.Order(sortOrder).Find(&people)

	return c.JSON(http.StatusOK, people)
}

func (a *App) CreatePerson(c echo.Context) error {
	person := Person{}
	if err := c.Bind(&person); err != nil {
		return err
	}
	a.DB.Connection.Save(&person)
	return c.JSON(http.StatusOK, person)
}

func (a *App) DeletePerson(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	person := a.getPerson(id)

	a.DB.Connection.Delete(&person)
	return c.String(http.StatusOK, "Deletion Successful" + string(id))
}

func (a *App) UpdatePerson(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	person := a.getPerson(id)

	changes := Person{}
	if err := c.Bind(&changes); err != nil {
		return err
	}

	// What fields can be changed?
	person.Visible = changes.Visible

	a.DB.Connection.Save(&person)
	return c.JSON(http.StatusOK, person)
}

func (a *App) SeedPeople(c echo.Context) error {
	a.DB.Connection.Delete(&Person{})
	a.DB.Connection.Create(&Person{
		Name: "Ryan",
		Age: 32,
		Balance: 1000,
		Email: "ryanmccombe@example.com",
		Address: "Belfast",
		Visible: true,
	})

	a.DB.Connection.Create(&Person{
		Name: "Hannah",
		Age: 28,
		Balance: 1000,
		Email: "hannah_morrison@example.com",
		Address: "Belfast",
		Visible: true,
	})

	a.DB.Connection.Create(&Person{
		Name: "Testy McTestFace",
		Age: 45,
		Balance: 1000,
		Email: "a-tester@example.com",
		Address: "Dublin",
		Visible: false,
	})

	return c.String(http.StatusOK, "Seeded!")
}

func (a *App) getPerson(id int) Person {
	foundRecord := Person{}
	a.DB.Connection.First(&foundRecord, id)
	return foundRecord
}