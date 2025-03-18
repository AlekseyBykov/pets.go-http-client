package model

type AssetsResponse struct {
	Data      []AssetData `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

type AssetResponse struct {
	Data      AssetData `json:"data"`
	Timestamp int64     `json:"timestamp"`
}
