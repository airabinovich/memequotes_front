package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airabinovich/memequotes_front/api/config"
	commonContext "github.com/airabinovich/memequotes_front/api/context"
	httpCommons "github.com/airabinovich/memequotes_front/api/http"
	"github.com/airabinovich/memequotes_front/api/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type wrappedPhrasesResponse struct {
	Results []model.PhraseResult `json:"results"`
}

func (client MemequotesBackendHttpClient) GetAllPhrasesForCharacter(ctx *gin.Context, characterID int64) ([]model.PhraseResult, error) {
	logger := commonContext.Logger(ctx)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/character/%d/phrases", client.host, characterID), nil)
	if err != nil {
		logger.Error("Request cannot be created", err)
		return []model.PhraseResult{}, err
	}

	var resp *http.Response
	err = httpCommons.Retry(
		func() error {
			var fnErr error
			resp, fnErr = client.client.Do(req)
			if fnErr != nil {
				logger.Error("Request cannot be performed", err)
				return fnErr
			}

			if resp.StatusCode >= http.StatusInternalServerError {
				defer resp.Body.Close()
				fnErr = errors.New("internal server error")
				bodyAsString, _ := httpCommons.ResponseBodyAsString(resp)
				logger.Error(fmt.Sprintf("Backend server error status: %d - url: %s - body: %s", resp.StatusCode, resp.Request.URL.String(), bodyAsString), fnErr)
				return fnErr
			}

			return nil
		}, 2, time.Duration(config.Conf.GetTimeDuration("client.memequotes_backend.sleep_retry")))
	if err != nil {
		return []model.PhraseResult{}, err
	}

	defer resp.Body.Close()

	if (resp.StatusCode / 100) != 2 {
		bodyAsString, _ := httpCommons.ResponseBodyAsString(resp)
		logger.Warn(fmt.Sprintf("Backend bad request - status: %d - url: %s - body: %s", resp.StatusCode, resp.Request.URL.String(), bodyAsString))
		return []model.PhraseResult{}, err
	}

	var result wrappedPhrasesResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		logger.Error("Backend response cannot be parsed", err)
		return []model.PhraseResult{}, err
	}
	return result.Results, nil
}
