package main

import (
	"github.com/gin-gonic/gin"
    "github.com/leofideliss/mocks/routes"
)

func main() {
    router := gin.Default()
    routes.SetupRoutes(router)
    router.Run()
}

