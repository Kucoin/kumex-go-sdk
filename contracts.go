package kumex

import "net/http"

// A ContractsModel is the struct.
type ContractsModel struct {
	BaseCurrency       string `json:"baseCurrency"`
	FairMethod         string `json:"fairMethod"`
	FundingBaseSymbol  string `json:"fundingBaseSymbol"`
	FundingQuoteSymbol string `json:"fundingQuoteSymbol"`
	FundingRateSymbol  string `json:"fundingRateSymbol"`
	IndexSymbol        string `json:"indexSymbol"`
	IsDeleverage       bool   `json:"isDeleverage"`
	InitialMargin      string `json:"baseCurrency"`
	IsInverse          bool   `json:"isInverse"`
	IsQuanto           bool   `json:"isQuanto"`
	LotSize            string `json:"lotSize"`
	MaintainMargin     string `json:"maintainMargin"`
	MakerFeeRate       string `json:"makerFeeRate"`
	MakerFixFee        string `json:"makerFixFee"`
	MarkMethod         string `json:"markMethod"`
	MaxOrderQty        string `json:"maxOrderQty"`
	MaxPrice           string `json:"maxPrice"`
	MaxRiskLimit       string `json:"maxRiskLimit"`
	MinRiskLimit       string `json:"minRiskLimit"`
	Multiplier         string `json:"multiplier"`
	QuoteCurrency      string `json:"quoteCurrency"`
	RiskStep           string `json:"riskStep"`
	RootSymbol         string `json:"rootSymbol"`
	Status             string `json:"status"`
	Symbol             string `json:"symbol"`
	TakerFeeRate       string `json:"takerFeeRate"`
	TakerFixFee        string `json:"takerFixFee"`
	TickSize           string `json:"tickSize"`
}

// ActiveContracts Get Open Contract List.
func (as *ApiService) ActiveContracts() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/contracts/active", nil)
	return as.Call(req)
}

// Contracts Get Order Info. of the Contract.
func (as *ApiService) Contracts(symbol string) (*ApiResponse, error) {
	p := map[string]string{}
	if symbol != "" {
		p["symbol"] = symbol
	}
	req := NewRequest(http.MethodGet, "/api/v1/position", p)
	return as.Call(req)
}
