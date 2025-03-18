package model

type AssetData struct {
	ID           string `json:"id"`
	Rank         string `json:"rank"`
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	MaxSupply    string `json:"maxSupply"`
	MarketCapUSD string `json:"marketCapUSD"`
	VolumeUSD24h string `json:"volumeUSD24h"`
	PriceUSD     string `json:"priceUSD"`
}

func (a AssetData) Info() string {
	return "[ID] " + a.ID + " | [RANK] " + a.Rank + " | [SYMBOL] " + a.Symbol +
		" | [NAME] " + a.Name + " | [PRICE] " + a.PriceUSD
}
