package handler

import (
	"errors"
	"github.com/gobaldia/stocks-api/upstream"
	"testing"
)

func TestProcessQuoteRetrievalAndEncryptionOk(t *testing.T) {
	upstream.GetQuoteMock = func(symbol string) (*upstream.Quote, error) {
		return &upstream.Quote{
			Symbol:               "SYM",
			CurrentValue:         100,
			PreviousClose:        100,
			Open:                 100,
			MarketCapitalization: 10000,
			Volume:               100,
		}, nil
	}

	upstream.EncryptQuoteMock = func(quote upstream.Quote) (*upstream.EncryptedQuote, error) {
		return &upstream.EncryptedQuote{
			Value: "3ncrypt3dQu0t3",
		}, nil
	}

	alphaVantageMock := upstream.AlphaVantageMock{}
	encrypterMock := upstream.EncrypterMock{}

	encryptedQuote, err := processQuoteRetrievalAndEncryption("SYM", alphaVantageMock, encrypterMock)

	if err != nil {
		t.Errorf("Error was expected to be nil, but it was %v", err)
	}

	if encryptedQuote.Value != "3ncrypt3dQu0t3" {
		t.Errorf("Value was expected to be \"3ncrypt3dQu0t3\", but it was %v", encryptedQuote.Value)
	}
}

func TestProcessQuoteRetrievalAndEncryptionAlphaVantageError(t *testing.T) {
	upstream.GetQuoteMock = func(symbol string) (*upstream.Quote, error) {
		return nil, errors.New("some error")
	}

	upstream.EncryptQuoteMock = func(quote upstream.Quote) (*upstream.EncryptedQuote, error) {
		return &upstream.EncryptedQuote{
			Value: "3ncrypt3dQu0t3",
		}, nil
	}

	alphaVantageMock := upstream.AlphaVantageMock{}
	encrypterMock := upstream.EncrypterMock{}

	encryptedQuote, err := processQuoteRetrievalAndEncryption("SYM", alphaVantageMock, encrypterMock)

	if err == nil {
		t.Errorf("Error was expected to be present, but it was not")
	}

	if err.Error() != "some error" {
		t.Errorf("Error message was expected to be \"some error\", but it was \"%v\"", err)
	}

	if encryptedQuote != nil {
		t.Errorf("Encrypted quote was expected to be nil, but it was %v", encryptedQuote)
	}
}

func TestProcessQuoteRetrievalAndEncryptionEncrypterError(t *testing.T) {
	upstream.GetQuoteMock = func(symbol string) (*upstream.Quote, error) {
		return &upstream.Quote{
			Symbol:               "SYM",
			CurrentValue:         100,
			PreviousClose:        100,
			Open:                 100,
			MarketCapitalization: 10000,
			Volume:               100,
		}, nil
	}

	upstream.EncryptQuoteMock = func(quote upstream.Quote) (*upstream.EncryptedQuote, error) {
		return nil, errors.New("some error")
	}

	alphaVantageMock := upstream.AlphaVantageMock{}
	encrypterMock := upstream.EncrypterMock{}

	encryptedQuote, err := processQuoteRetrievalAndEncryption("SYM", alphaVantageMock, encrypterMock)

	if err == nil {
		t.Errorf("Error was expected to be present, but it was not")
	}

	if err.Error() != "some error" {
		t.Errorf("Error message was expected to be \"some error\", but it was \"%v\"", err)
	}

	if encryptedQuote != nil {
		t.Errorf("Encrypted quote was expected to be nil, but it was %v", encryptedQuote)
	}
}

func TestProcessQuoteRetrievalAndEncryptionMultipleErrors(t *testing.T) {
	upstream.GetQuoteMock = func(symbol string) (*upstream.Quote, error) {
		return nil, errors.New("some error on Alpha Vantage")
	}

	upstream.EncryptQuoteMock = func(quote upstream.Quote) (*upstream.EncryptedQuote, error) {
		return nil, errors.New("some error on Encrypter")
	}

	alphaVantageMock := upstream.AlphaVantageMock{}
	encrypterMock := upstream.EncrypterMock{}

	encryptedQuote, err := processQuoteRetrievalAndEncryption("SYM", alphaVantageMock, encrypterMock)

	if err == nil {
		t.Errorf("Error was expected to be present, but it was not")
	}

	if err.Error() != "some error on Alpha Vantage" {
		t.Errorf("Error message was expected to be \"some error on Alpha Vantage\", but it was \"%v\"", err)
	}

	if encryptedQuote != nil {
		t.Errorf("Encrypted quote was expected to be nil, but it was %v", encryptedQuote)
	}
}