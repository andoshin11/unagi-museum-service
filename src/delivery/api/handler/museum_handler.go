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
	GetByID(c *gin.Context)
}

type museumHandler struct {
	museumUsecase usecase.MuseumUsecase
}

// NewMuseumHandler returns a handler
func NewMuseumHandler(u usecase.MuseumUsecase) MuseumHandler {
	return &museumHandler{
		museumUsecase: u,
	}
}

func (h *museumHandler) GetAll(c *gin.Context) {
	ctx := context.Background()

	museums, err := h.museumUsecase.GetAll(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.JSON(http.StatusOK, gin.H{"items": museums})
}

func (h *museumHandler) GetNeighbors(c *gin.Context) {
	ctx := context.Background()

	lat, err := strconv.ParseFloat(c.Query("lat"), 64)
	lng, err := strconv.ParseFloat(c.Query("lng"), 64)
	distance, err := strconv.Atoi(c.Query("distance"))

	if err != nil {
		log.Fatalln(err)
	}

	museums, err := h.museumUsecase.GetNeighbors(ctx, lat, lng, distance)

	if err != nil {
		log.Fatalln(err)
	}
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.JSON(http.StatusOK, gin.H{"items": museums})
}

func (h *museumHandler) GetByID(c *gin.Context) {
	ctx := context.Background()

	id := c.Param("id")

	museum, err := h.museumUsecase.GetByID(ctx, id)

	if err != nil {
		log.Fatalln(err)
	}
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.JSON(http.StatusOK, gin.H{"item": museum})
}
