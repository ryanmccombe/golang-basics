package main

import (
	"./database"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	app := &App{
		DB: database.Connect(),
	}

	app.PrepareDB()

	e.GET("/people", app.ListPeople)
	e.POST("/people", app.CreatePerson)
	e.DELETE("/people/:id", app.DeletePerson)
	e.PATCH("/people/:id", app.UpdatePerson)

	e.GET("/people/seed", app.SeedPeople)

	e.Logger.Fatal(e.Start(":3001"))
}

func (a *App) PrepareDB() {
	a.DB.Connection.AutoMigrate(&Person{})
}