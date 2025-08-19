package apiclientcoincapio

import "fmt"

type AssetsResponse struct {
	Data      []AssetData `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

type AssetResponse struct {
	Data      AssetData `json:"data"`
	Timestamp int64     `json:"timestamp"`
}

type AssetData struct {
	ID            string `json:"id"`
	Rank          string `json:"rank"`
	Symbol        string `json:"symbol"`
	Name          string `json:"name"`
	Supply        string `json:"supply"`
	MaxSupply     string `json:"maxSupply"`
	MarketCapUSD  string `json:"marketCapUsd"`
	VolumeUSD24Hr string `json:"volumeUsd24Hr"`
	PriceUSD      string `json:"priceUsd"`
}

func (d AssetData) Info() string {
	return fmt.Sprintf("[id] %s | [RANK] %s | [SYMBOL] %s | [NAME] %s | [PRICE] %s",
		d.ID, d.Rank, d.Symbol, d.Name, d.PriceUSD)
}
