package main

import (
	"finworker/internal/app"
)

// @title			E2E
// @version		0.0
// @description	Описание взаимодействия и работы внутренних моделей E2E CRM
//
// @license.name	MIT License
// @license.url	https://github.com/GeekchanskiY/E2E/blob/main/LICENSE
//
// @host			localhost:8080
// @BasePath		/
//
// @contact.name	Repo
// @contact.url	https://github.com/GeekchanskiY/E2E
//
// @accept			json
// @produce		json
// @schemes		http
func main() {
	app.NewApp().Run()
}
