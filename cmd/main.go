package main

import (
	"fmt"
	"github.com/AlekseyBykov/pets.go-http-client/internal/client"
	"log"
	"os"
)

func main() {
	c := client.NewCoinCapClient(os.Stdout)

	assets, err := c.GetAssets()
	if err != nil {
		log.Fatal(err)
	}

	for _, asset := range assets {
		fmt.Println(asset.Info())
	}

	asset, err := c.GetAsset("bitcoin")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(asset.Info())
}
