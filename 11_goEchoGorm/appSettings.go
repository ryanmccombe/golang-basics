package main

import (
	"./database"
)

type App struct {
	DB *database.Settings
	Test string
}
