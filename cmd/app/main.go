package main

import "github.com/lastdoctor/emma-app-go/internal/app"

func main() {
	const configsDir = "configs"
	app.Run(configsDir)
}
