package upstream

type AlphaVantageMock struct {}

var GetQuoteMock func(symbol string) (*Quote, error)

func (mock AlphaVantageMock) GetQuote(symbol string) (*Quote, error) {
	return GetQuoteMock(symbol)
}
