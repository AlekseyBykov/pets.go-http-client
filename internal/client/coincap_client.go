package client

import (
	"encoding/json"
	"fmt"
	"github.com/AlekseyBykov/pets.go-http-client/internal/model"
	"io"
	"net/http"
	"time"
)

type CoinCapClient struct {
	httpClient *http.Client
	baseURL    string
}

func NewCoinCapClient(logger io.Writer) *CoinCapClient {
	return &CoinCapClient{
		baseURL: "https://api.coincap.io/v2",
		httpClient: &http.Client{
			Transport: &loggingRoundTripper{
				logger: logger,
				next:   http.DefaultTransport,
			},
			Timeout: 10 * time.Second,
		},
	}
}

func (c *CoinCapClient) GetAssets() ([]model.AssetData, error) {
	resp, err := c.httpClient.Get(c.baseURL + "/assets")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var assetsResponse model.AssetsResponse
	if err = json.NewDecoder(resp.Body).Decode(&assetsResponse); err != nil {
		return nil, err
	}

	return assetsResponse.Data, nil
}

func (c *CoinCapClient) GetAsset(id string) (*model.AssetData, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf("%s/assets/%s", c.baseURL, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var assetResponse model.AssetResponse
	if err = json.NewDecoder(resp.Body).Decode(&assetResponse); err != nil {
		return nil, err
	}

	return &assetResponse.Data, nil
}
