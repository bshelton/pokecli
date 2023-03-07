/*
Package api makes api requests to the pokemontcg api.
*/
package api

import (
	"encoding/json"

	"github.com/bshelton/pokecli/pkg/logger"
	"github.com/bshelton/pokecli/pkg/model"
)

/* SearchClassic uses the cards api and transforms the result into a Classic mode result
 */
func SearchClassic() (*model.ClassicOutput, error) {
	apiResponse, err := Search("/cards", "rarity:Rare (types:fire OR types:grass) hp:[90 TO *]", "id,name,types,hp,rarity", "id", 10)

	if err != nil {
		return nil, err
	}

	var cards model.ClassicResponse

	err = json.Unmarshal([]byte(*apiResponse), &cards)
	if err != nil {
		logger.Error("%s", err)
		return nil, err
	}

	var output model.ClassicOutput
	var cardList []model.ClassicCard
	for card := range cards.Data {
		cardType := ""
		if len(cards.Data[card].Types) > 0 {
			cardType = cards.Data[card].Types[0]
		}
		cardList = append(cardList, model.ClassicCard{
			ID:     cards.Data[card].ID,
			Name:   cards.Data[card].Name,
			Type:   cardType,
			HP:     cards.Data[card].HP,
			Rarity: cards.Data[card].Rarity,
		})
	}
	output.Data = cardList

	return &output, nil
}

/* SearchCards uses the cards api and returns a CardSearchResponse model
* @param query - The search query
* @param fields - The fields to include in the response, if empty then it will return all fields
* @param orderby - The field to order the return data by
* @param limit - The number of cards to search for
*
 */
func SearchCards(query string, fields string, orderby string, limit int) (*model.CardSearchResponse, error) {
	apiResponse, err := Search("/cards", query, fields, orderby, limit)

	if err != nil {
		return nil, err
	}

	var cards model.CardSearchResponse

	err = json.Unmarshal([]byte(*apiResponse), &cards)

	if err != nil {
		return nil, err
	}

	return &cards, nil
}
