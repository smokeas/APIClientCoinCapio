package main

import (
	"fmt"
	"log"
	"time"

	"coincap/apiclientcoincapio"
)

func main() {
	coincapClient, err := apiclientcoincapio.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}

	assetData, err := coincapClient.GetAssets()
	if err != nil {
		log.Fatal(err)
	}

	for _, asset := range assetData {
		fmt.Println(asset.Info())
	}
}
