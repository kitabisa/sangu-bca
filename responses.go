package bca

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

type AccStatementsResp struct {
	StartDate    string                    `json:"StartDate"`
	EndDate      string                    `json:"EndDate"`
	Currency     string                    `json:"Currency"`
	StartBalance string                    `json:"StartBalance"`
	Data         []AccStatementsDetailResp `json:"Data"`
}

type AccStatementsDetailResp struct {
	TransactionDate   string `json:"TransactionDate"`
	BranchCode        string `json:"BranchCode"`
	TransactionType   string `json:"TransactionType"`
	TransactionAmount string `json:"TransactionAmount"`
	TransactionName   string `json:"TransactionName"`
	Trailer           string `json:"Trailer"`
}

type FailedResponse struct {
	ErrorCode    string `json:"ErrorCode"`
	ErrorMessage struct {
		Indonesian string `json:"Indonesian"`
		English    string `json:"English"`
	} `json:"ErrorMessage"`
}
