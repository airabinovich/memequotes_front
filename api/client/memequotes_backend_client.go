package client

import (
	"github.com/airabinovich/memequotes_front/api/model"
	"github.com/gin-gonic/gin"
)

type MemequotesBackendClient interface {
	// GetAll retrieves all character
	GetAll(c *gin.Context) ([]model.CharacterResult, error)
}
