package upstream

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gobaldia/stocks-api/config"
	"net/http"
)

type EncrypterClient struct {
	BaseURL string
}

func NewEncrypterClient() *EncrypterClient {
	return &EncrypterClient{
		BaseURL: config.GetConfig().EncrypterApi,
	}
}

type EncryptedQuote struct {
	Value string
}

type EncrypterApi interface {
	EncryptQuote(quote Quote) (*EncryptedQuote, error)
}

func (ec *EncrypterClient) EncryptQuote(quote Quote) (*EncryptedQuote, error) {
	encryptionEndpoint := fmt.Sprintf("%s/encrypt", ec.BaseURL)

	jsonQuote, err := json.Marshal(quote)

	r, err := http.Post(encryptionEndpoint, "application/json", bytes.NewBuffer(jsonQuote))
	if err != nil {
		return nil, fmt.Errorf("failed encrypting quote: %s", err.Error())
	}

	var encryptedQuoteResponse EncryptedQuote

	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&encryptedQuoteResponse); err != nil {
		return nil, fmt.Errorf("failed decoding encrypted quote: %s", err.Error())
	}

	return &encryptedQuoteResponse, err
}
