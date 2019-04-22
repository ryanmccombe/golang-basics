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

	e.GET("/users", app.ListPeople)
	e.POST("/users", app.CreatePerson)
	e.DELETE("/users/:id", app.DeletePerson)
	e.PATCH("/users/:id", app.UpdatePerson)

	e.GET("/users/seed", app.SeedPeople)

	e.Logger.Fatal(e.Start(":3001"))
}

func (a *App) PrepareDB() {
	a.DB.Connection.AutoMigrate(&Person{})
}