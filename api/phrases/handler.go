package phrases

import (
	"fmt"
	"github.com/airabinovich/memequotes_front/api/client"
	commonContext "github.com/airabinovich/memequotes_front/api/context"
	"github.com/airabinovich/memequotes_front/api/rest"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var memequotesBackClient client.MemequotesBackendClient

func Initialize() {
	memequotesBackClient = client.NewMemequotesBackendDefaultClient()
}

func GetAllPhrasesForCharacter(c *gin.Context) {
	rest.ErrorWrapper(getAllPhrasesForCharacter, c)
}

func getAllPhrasesForCharacter(c *gin.Context) *rest.APIError {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)

	characterId, err := strconv.ParseInt(c.Param("character-id"), 10, 64)
	if err != nil {
		logger.Error("getting character with non-numeric characterId", err)
		return rest.NewBadRequest(err.Error())
	}

	logger.Debug(fmt.Sprintf("Getting phrases for character id %d", characterId))
	phrases, err := memequotesBackClient.GetAllPhrasesForCharacter(c, characterId)
	if err != nil {
		logger.Error("get character by id", err)
		return rest.NewInternalServerError(err.Error())
	}
	c.JSON(http.StatusOK, phrases)
	return nil
}
