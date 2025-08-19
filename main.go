package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type assetsResponse struct {
	Data []assetData `json:"data"`
}

type assetData struct {
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

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	return l.next.RoundTrip(r)
}

func main() {
	client := &http.Client{
		Transport: &loggingRoundTripper{
			logger: os.Stdout,
			next:   http.DefaultTransport,
		},
		Timeout: time.Second * 15,
	}

	req, err := http.NewRequest("GET", "https://rest.coincap.io/v3/assets?limit=5", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer ключ")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Raw response:", string(body)) // посмотреть что вернул API

	var r assetsResponse
	if err = json.Unmarshal(body, &r); err != nil {
		log.Fatal(err)
	}

	for _, asset := range r.Data {
		fmt.Printf("%s (%s): $%s\n", asset.Name, asset.Symbol, asset.PriceUSD)
	}
}

