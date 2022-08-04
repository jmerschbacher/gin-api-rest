package main

import (
	"github.com/jmerschbacher/gin-api-rest/database"
	"github.com/jmerschbacher/gin-api-rest/routes"
)

func main() {
	database.ConexaoDB()
	routes.HandleRequests()
}
