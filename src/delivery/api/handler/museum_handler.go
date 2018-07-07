package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/andoshin11/euphro-ddd/src/usecase"
	"github.com/gin-gonic/gin"
)

type MuseumHandler interface {
	GetAll(ctx *gin.Context)
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
