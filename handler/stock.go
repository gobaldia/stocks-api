package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gobaldia/stocks-api/upstream"
)

func GetQuote(ctx *gin.Context) {
	symbol := ctx.Query("symbol")

	alphaVantageClient := upstream.NewAlphaVantageClient()
	encrypterClient := upstream.NewEncrypterClient()

	encryptedQuote, err := processQuoteRetrievalAndEncryption(symbol, alphaVantageClient, encrypterClient)

	if err != nil {
		ctx.String(500, err.Error())
		ctx.Abort()
		return
	}

	ctx.JSON(200, &encryptedQuote)
}

func processQuoteRetrievalAndEncryption(symbol string, alphaVantageApi upstream.AlphaVantageApi, encrypterApi upstream.EncrypterApi) (*upstream.EncryptedQuote, error) {

	quote, err := alphaVantageApi.GetQuote(symbol)

	if err != nil {
		return nil, err
	}

	encryptedQuote, err := encrypterApi.EncryptQuote(*quote)

	return encryptedQuote, err
}
