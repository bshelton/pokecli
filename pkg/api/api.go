/*
Package api makes api requests to the pokemontcg api.
*/
package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/bshelton/pokecli/pkg/config"
	"github.com/bshelton/pokecli/pkg/logger"
)

/* Search function builds an escaped query based the following parameters
* @param path - The API endpoint
* @param fields - The fields to include in the response, if empty then it will return all fields
* @param query - The search query
* @param orderby - The field to order the return data by
* @param limit - The number of cards to search for
*
 */
func Search(path string, query string, fields string, orderby string, limit int) (*string, error) {
	// The API expects that all order by field names to be lowercased.
	orderby = strings.ToLower(orderby)

	params := "?q=" + url.QueryEscape(query) + "&pageSize=" + strconv.Itoa(limit) + "&orderBy=" + url.QueryEscape(orderby)

	if fields != "" {
		params = params + "&select=" + url.QueryEscape(fields)
	}

	return APIRequest(path + params)
}

func APIRequest(path string) (*string, error) {
	client := &http.Client{}
	finalURL := config.BaseURL + path

	logger.Debug(fmt.Sprintf("Sending request to:  %s", finalURL))
	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	// Add API Key if env var is set.
	apiKey, ok := os.LookupEnv(config.APIKeyEnvVarKey)
	if ok {
		req.Header.Add(config.AuthHeaderKey, apiKey)
	} else {
		logger.Warn(fmt.Sprintf("Api Key not set so therefor rate limiting is reduced. Set %s for a better experience", config.APIKeyEnvVarKey))
	}

	resp, err := client.Do(req)

	if err != nil {
		logger.Error("Could not create request, %v", err.Error())
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	bodyString := string(bodyBytes)

	return &bodyString, nil
}
