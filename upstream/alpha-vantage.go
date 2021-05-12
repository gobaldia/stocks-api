package upstream

import (
	"encoding/json"
	"github.com/gobaldia/stocks-api/config"

	//"fmt"
	//"net/http"
	"fmt"
	"net/http"
)

type Quote struct {
	Symbol string
	CurrentValue float64
	//Variation float64
	PreviousClose float64
	Open float64
	//Bid float64
	//Ask float64
	MarketCapitalization float64
	Volume int
	//AverageVolume float64
}

type AlphaVantageClient struct {
	BaseURL string
	APIKey string
}

func NewAlphaVantageClient() *AlphaVantageClient {
	return &AlphaVantageClient{
		BaseURL: "https://www.alphavantage.co",
		APIKey: config.GetConfig().AlphaVantageApiKey,
	}
}

type AlphaVantageGlobalQuote struct {
	Symbol           string  `json:"01. symbol"`
	Open             float64 `json:"02. open,string"`
	High             float64 `json:"03. high,string"`
	Low              float64 `json:"04. low,string"`
	Price            float64 `json:"05. price,string"`
	Volume           int     `json:"06. volume,string"`
	LatestTradingDay string  `json:"07. latest trading day"`
	PreviousClose    float64 `json:"08. previous close,string"`
	Change           float64 `json:"09. change,string"`
}

type AlphaVantageGlobalQuoteResponse struct {
	GlobalQuote AlphaVantageGlobalQuote `json:"Global Quote"`
}

type AlphaVantageApi interface {
	GetQuote(symbol string) (*Quote, error)
}

func (avc *AlphaVantageClient) GetQuote(symbol string) (*Quote, error) {
	quotesEndpoint := fmt.Sprintf("%s/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", avc.BaseURL, symbol, avc.APIKey)

	r, err := http.Get(quotesEndpoint)

	if err != nil {
		return nil, fmt.Errorf("failed retrieving quote: %s", err.Error())
	}

	var quoteResponse AlphaVantageGlobalQuoteResponse

	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&quoteResponse); err != nil {
		return nil, fmt.Errorf("failed decoding quote: %s", err.Error())
	}

	quote := Quote{
		Symbol:               quoteResponse.GlobalQuote.Symbol,
		CurrentValue:         quoteResponse.GlobalQuote.Price,
		//Variation:            quoteResponse.GlobalQuote.,
		PreviousClose:        quoteResponse.GlobalQuote.PreviousClose,
		Open:                 quoteResponse.GlobalQuote.Open,
		//Bid:                  quoteResponse.GlobalQuote.Bi,
		//Ask:                  quoteResponse.GlobalQuote.Ask,
		MarketCapitalization: float64(quoteResponse.GlobalQuote.Volume) * quoteResponse.GlobalQuote.Price,
		Volume:               quoteResponse.GlobalQuote.Volume,
		//AverageVolume:        quoteResponse.GlobalQuote.A,
	}

	return &quote, err
}