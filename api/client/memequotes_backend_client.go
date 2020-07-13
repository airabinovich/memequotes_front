package client

import (
	"fmt"
	"github.com/airabinovich/memequotes_front/api/config"
	httpCommons "github.com/airabinovich/memequotes_front/api/http"
	"github.com/airabinovich/memequotes_front/api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MemequotesBackendClient interface {
	// GetAllCharacters retrieves all character
	GetAllCharacters(c *gin.Context) ([]model.CharacterResult, error)

	// GetAllPhrasesForCharacter retrieves all phrases for a character
	GetAllPhrasesForCharacter(ctx *gin.Context, characterID int64) ([]model.PhraseResult, error)
}

type MemequotesBackendHttpClient struct {
	host   string
	client *http.Client
}

func NewMemequotesBackendClient(host string, client *http.Client) MemequotesBackendHttpClient {
	return MemequotesBackendHttpClient{
		host:   host,
		client: client,
	}
}

func NewMemequotesBackendDefaultClient() MemequotesBackendHttpClient {
	host := fmt.Sprintf("%s:%d",
		config.Conf.GetString("client.memequotes_backend.host"),
		config.Conf.GetInt32("client.memequotes_backend.port"),
	)
	return NewMemequotesBackendClient(
		host,
		httpCommons.Client(),
	)
}
