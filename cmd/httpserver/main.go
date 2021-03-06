package httpserver

import (
	"github.com/Haleluak/hexagonal/internal/core/services"
	"github.com/Haleluak/hexagonal/internal/handlers"
	"github.com/Haleluak/hexagonal/internal/repositories"
	"github.com/Haleluak/hexagonal/pkg/uidgen"
	"github.com/gin-gonic/gin"
)

func main() {
	gamesRepository := repositories.NewMemKVS()
	gamesService := services.New(gamesRepository, uidgen.New())
	gamesHandler := handlers.NewHTTPHandler(gamesService)

	router := gin.New()
	router.GET("/games/:id", gamesHandler.Get)
	router.POST("/games", gamesHandler.Create)
	router.PUT("/games/:id", gamesHandler.RevealCell)

	router.Run(":8080")
}