package main

import (
	"VacancyService/internal/app"
	_ "github.com/lib/pq"
)

func main() {
	bootstrap := app.NewApp()
	bootstrap.Run()
}
