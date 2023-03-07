/*
Package model ...
*/

package model

type CardResponse struct {
	Cards []Card
}

type Card struct {
	ID     string   `json:"ID"`
	Name   string   `json:"Name"`
	Types  []string `json:"Types"`
	HP     string   `json:"HP"`
	Rarity string   `json:"Rarity"`
}

type ClassicCard struct {
	ID     string `json:"ID"`
	Name   string `json:"Name"`
	Type   string `json:"Type"`
	HP     string `json:"HP"`
	Rarity string `json:"Rarity"`
}

type ClassicOutput struct {
	Data []ClassicCard `json:"Cards"`
}
type ClassicResponse struct {
	Data []Card `json:"Data"`
}

type CardSearchOutput struct {
	Data []map[string]interface{} `json:"Cards"`
}
type CardSearchResponse struct {
	Data []map[string]interface{} `json:"Data"`
}
