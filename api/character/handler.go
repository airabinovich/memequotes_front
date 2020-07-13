package character

import (
	"github.com/airabinovich/memequotes_front/api/client"
	commonContext "github.com/airabinovich/memequotes_front/api/context"
	"github.com/airabinovich/memequotes_front/api/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

var memequotesBackClient client.MemequotesBackendClient

func Initialize() {
	memequotesBackClient = client.NewMemequotesBackendDefaultClient()
}

func GetAllCharacters(c *gin.Context) {
	rest.ErrorWrapper(getAllCharacters, c)
}

func getAllCharacters(c *gin.Context) *rest.APIError {
	ctx := commonContext.RequestContext(c)
	logger := commonContext.Logger(ctx)

	logger.Debug("Getting all characters")
	chs, err := memequotesBackClient.GetAllCharacters(c)
	if err != nil {
		logger.Error("get all character", err)
		return rest.NewInternalServerError(err.Error())
	}

	c.JSON(http.StatusOK, chs)
	return nil
}
