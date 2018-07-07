package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/andoshin11/euphro-ddd/src/usecase"
	"github.com/gin-gonic/gin"
)

type MuseumHandler interface {
	GetAll(c *gin.Context)
	GetNeighbors(c *gin.Context)
}

type museumHandler struct {
	MuseumUsecase usecase.MuseumUsecase
}

func NewMuseumHandler(u usecase.MuseumUsecase) MuseumHandler {
	return &museumHandler{
		MuseumUsecase: u,
	}
}

func (h *museumHandler) GetAll(c *gin.Context) {
	ctx := context.Background()

	museums, err := h.MuseumUsecase.GetAll(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{"museums": museums})
}

func (h *museumHandler) GetNeighbors(c *gin.Context) {
	ctx := context.Background()

	lat, err := strconv.ParseFloat(c.Query("lat"), 64)
	lng, err := strconv.ParseFloat(c.Query("lng"), 64)
	distance, err := strconv.Atoi(c.Query("distance"))

	if err != nil {
		log.Fatalln(err)
	}

	museums, err := h.MuseumUsecase.GetNeighbors(ctx, lat, lng, distance)

	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{"museums": museums})
}
