package upstream

type EncrypterMock struct{}

var EncryptQuoteMock func(quote Quote) (*EncryptedQuote, error)

func (mock EncrypterMock) EncryptQuote(quote Quote) (*EncryptedQuote, error) {
	return EncryptQuoteMock(quote)
}