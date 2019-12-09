package kumex

import (
	"net/http"
)

type PositionModel struct {
	Id                string `json:"id"`
	Symbol            string `json:"symbol"`
	AutoDeposit       bool   `json:"autoDeposit"`
	MaintMarginReq    string `json:"maintMarginReq"`
	RiskLimit         string `json:"riskLimit"`
	RealLeverage      string `json:"realLeverage"`
	CrossMode         bool   `json:"crossMode"`
	DelevPercentage   string `json:"delevPercentage"`
	OpeningTimestamp  string `json:"openingTimestamp"`
	CurrentTimestamp  string `json:"currentTimestamp"`
	CurrentQty        string `json:"currentQty"`
	CurrentCost       string `json:"currentCost"`
	CurrentComm       string `json:"currentComm"`
	UnrealisedCost    string `json:"unrealisedCost"`
	RealisedGrossCost string `json:"realisedGrossCost"`
	RealisedCost      string `json:"realisedCost"`
	IsOpen            bool   `json:"isOpen"`
	MarkPrice         string `json:"markPrice"`
	MarkValue         string `json:"markValue"`
	PosCost           string `json:"posCost"`
	PosCross          string `json:"posCross"`
	PosInit           string `json:"posInit"`
	PosComm           string `json:"posComm"`
	PosLoss           string `json:"posLoss"`
	PosMargin         string `json:"posMargin"`
	PosMaint          string `json:"posMaint"`
	MaintMargin       string `json:"maintMargin"`
	RealisedGrossPnl  string `json:"realisedGrossPnl"`
	RealisedPnl       string `json:"realisedPnl"`
	UnrealisedPnl     string `json:"unrealisedPnl"`
	UnrealisedPnlPcnt string `json:"unrealisedPnlPcnt"`
	UnrealisedRoePcnt string `json:"unrealisedRoePcnt"`
	AvgEntryPrice     string `json:"avgEntryPrice"`
	LiquidationPrice  string `json:"liquidationPrice"`
	BankruptPrice     string `json:"bankruptPrice"`
}

// Get Position Details.
func (as *ApiService) position(symbol string) (*ApiResponse, error) {
	p := map[string]string{}
	if symbol != "" {
		p["symbol"] = symbol
	}
	req := NewRequest(http.MethodGet, "/api/v1/position", p)
	return as.Call(req)
}

// Get Position List.
func (as *ApiService) positions() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/positions", nil)
	return as.Call(req)
}

// Enable/Disable of Auto-Deposit Margin.
func (as *ApiService) autoDepositStatus(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/position/margin/auto-deposit-status", params)
	return as.Call(req)
}

// Add Margin Manually.
func (as *ApiService) depositMargin(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/position/margin/deposit-margin", params)
	return as.Call(req)
}
