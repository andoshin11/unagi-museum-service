package main

import (
	"context"
	"log"
	"net/http"

	"github.com/andoshin11/euphro-ddd/src/delivery/api/handler"
	"github.com/andoshin11/euphro-ddd/src/gateway"
	"github.com/andoshin11/euphro-ddd/src/repository"
	"github.com/andoshin11/euphro-ddd/src/usecase"
	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"test": "hoge"})
}

func main() {
	router := gin.Default()

	ctx := context.Background()
	client, err := gateway.NewFirestoreClient(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	museumRepository := repository.NewMuseumRepository(client)
	museumUseCase := usecase.NewMuseumUsecase(museumRepository)

	// handlers
	museumHandler := handler.NewMuseumHandler(museumUseCase)

	// API namespace
	v1 := router.Group("/api/v1")
	{
		v1.GET("/test", test)
		v1.GET("/museums", museumHandler.GetAll)
		v1.GET("/museums/:id", museumHandler.GetByID)
		v1.GET("/neighbors", museumHandler.GetNeighbors)
	}

	router.Run(":8000")
}
