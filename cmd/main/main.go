package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/bshelton/pokecli/pkg/api"
	"github.com/bshelton/pokecli/pkg/logger"
	"github.com/bshelton/pokecli/pkg/model"
)

func logLevelHelpMessage() string {
	return "Debug - 0\nInfo - 1\nWarning - 2\nError -3\nFatal -4"
}

func main() {
	var cardLimit int
	var sortBy string
	var fields string
	var classicMode bool
	var logLevel int
	flag.IntVar(&cardLimit, "limit", 250, "The number of cards to return")
	flag.StringVar(&sortBy, "sortBy", "id", "The field to sort by")
	flag.StringVar(&fields, "fields", "", "A comma seperated list of fields to include in the response. If empty it will include all fields")
	flag.BoolVar(&classicMode, "classic", false, "In classic mode the cli returns exactly what was asked for.")
	flag.IntVar(&logLevel, "loglevel", 4, logLevelHelpMessage())
	flag.Parse()

	logger.Setup("PokeCLI: ", logLevel)

	if classicMode {
		runClassicMode()
	} else {
		runFlexMode(cardLimit, sortBy, fields)
	}
}

/*
 * runFlexMode function takes in parameters to construct an api call.
 * @param cardLimit - The number of cards to return
 * @param sortBy - The field to sort by
 * @param fields - The fields to select in the query
 */
func runFlexMode(cardLimit int, sortBy string, fields string) {
	if cardLimit < 1 {
		log.Fatal("You entered a limit less than 1, please search for at least 1 card.")
	}

	logger.Info("Running in Flex mode.")
	logger.Info(fmt.Sprintf("Getting %d cards", cardLimit))

	cardResponse, err := api.SearchCards("rarity:Rare (types:fire OR types:grass) hp:[90 TO *]", fields, sortBy, cardLimit)

	if err != nil {
		logger.Fatal(err.Error())
	}

	fmt.Println(prettyOutput(model.CardSearchOutput(*cardResponse)))
}

/*
 * runClassicMode returns the exact data requests in the assessment.
 */
func runClassicMode() {
	logger.Info("Running in Classic mode.")
	apiResponse, err := api.SearchClassic()

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(prettyOutput(apiResponse))
}

/* prettyOutput displays the cards in JSON format
 * @param input - Any interface to output
 *
 */
func prettyOutput(input any) string {
	out, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		logger.Error(err)
		return "An error occurred formatting response"
	}
	// Display '&' instead of unicode
	return string(bytes.Replace(out, []byte("\\u0026"), []byte("&"), -1))
}
