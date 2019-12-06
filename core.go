package bca

import (
	"fmt"
	"io"
	"net/url"
	"strings"
)

// CoreGateway struct
type CoreGateway struct {
	Client Client
}

// Call : base method to call API
func (gateway *CoreGateway) Call(method, path string, header map[string]string, body io.Reader, v interface{}, x interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	path = gateway.Client.BcaBaseURL + path

	return gateway.Client.Call(method, path, header, body, v, x)
}

func (gateway *CoreGateway) GetToken() (TokenResponse, FailedResponse, error) {
	if gateway.Client.BcaGetTokenURL == "" {
		gateway.Client.BcaGetTokenURL = "/api/oauth/token"
	}
	respSuccess := TokenResponse{}
	respFailed := FailedResponse{}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	err := gateway.Call("POST", gateway.Client.BcaGetTokenURL, headers, strings.NewReader(data.Encode()), &respSuccess, &respFailed)
	if err != nil {
		return respSuccess, respFailed, err
	}

	return respSuccess, respFailed, nil
}

func (gateway *CoreGateway) AccStatement(req AccStatementsReq) (respS AccStatementsResp, respF FailedResponse, err error) {
	// default path
	if gateway.Client.BcaAccStatementsURL == "" {
		gateway.Client.BcaAccStatementsURL =
			fmt.Sprint("/banking/v3/corporates/",
				gateway.Client.BcaCompanyID,
				"/accounts/",
				req.AccountNumber,
				"/statements")
	}
	// query params
	q := url.Values{}
	q.Set("EndDate", req.EndDate)
	q.Set("StartDate", req.StartDate)
	gateway.Client.BcaAccStatementsURL = fmt.Sprint(gateway.Client.BcaAccStatementsURL, "?", q.Encode())

	// bca signature
	signature := BcaSignature{
		APISecret:   gateway.Client.BcaApiSecret,
		AccessToken: req.Token,
		HTTPMethod:  "GET",
		RelativeURL: gateway.Client.BcaAccStatementsURL,
		RequestBody: "",
		Timestamp:   getBcaTimestamp(),
	}
	bcaSignature, err := generateBcaSignature(signature)
	if err != nil {
		return
	}

	headers := map[string]string{
		"Authorization":   fmt.Sprintf("Bearer %v", req.Token),
		"Content-Type":    "application/json",
		"Origin":          gateway.Client.Origin,
		"X-BCA-Key":       gateway.Client.BcaApiKey,
		"X-BCA-Timestamp": signature.Timestamp,
		"X-BCA-Signature": bcaSignature,
	}

	err = gateway.Call("GET", gateway.Client.BcaAccStatementsURL, headers, nil, &respS, &respF)
	if err != nil {
		return
	}

	return
}
