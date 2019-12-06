package bca

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BcaTestSuite struct {
	suite.Suite
	client Client
}

func TestBcaTestSuite(t *testing.T) {
	suite.Run(t, new(BcaTestSuite))
}

func (b *BcaTestSuite) SetupSuite() {
	b.client = Client{
		BcaBaseURL:      "https://bca-hostname",
		BcaApiKey:       "your-bca-api-key",
		BcaApiSecret:    "your-bca-api-secret",
		BcaClientID:     "your-bca-client-id",
		BcaClientSecret: "your-bca-client-secret",
		BcaCompanyID:    "your-bca-company-id",
		Origin:          "your-origin-host",
	}
}

func (b *BcaTestSuite) TestGetTokenSuccess() {
	core := CoreGateway{
		Client: b.client,
	}

	tokenResp, failResp, err := core.GetToken()

	assert.NotEqual(b.T(), TokenResponse{}, tokenResp)
	assert.Equal(b.T(), FailedResponse{}, failResp)
	assert.Equal(b.T(), nil, err)
}

func (b *BcaTestSuite) TestGetAccStatementSuccess() {
	core := CoreGateway{
		Client: b.client,
	}

	tokenResp, failResp, err := core.GetToken()
	assert.Equal(b.T(), nil, err)
	assert.Equal(b.T(), FailedResponse{}, failResp)

	req := AccStatementsReq{
		AccountNumber: "0201245680",
		StartDate:     "2016-08-29",
		EndDate:       "2016-09-01",
		Token:         tokenResp.AccessToken,
	}

	_, accRespErr, err := core.AccStatement(req)
	assert.Equal(b.T(), nil, err)
	assert.Equal(b.T(), FailedResponse{}, failResp)
}
