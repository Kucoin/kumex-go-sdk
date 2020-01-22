package kumex

import (
	"net/http"
)

// An AccountModel represents an account.
type AccountModel struct {
	AccountEquity    float64 `json:"accountEquity"`
	UnrealisedPNL    float64 `json:"unrealisedPNL"`
	MarginBalance    float64 `json:"marginBalance"`
	PositionMargin   float64 `json:"positionMargin"`
	OrderMargin      float64 `json:"orderMargin"`
	FrozenFunds      float64 `json:"frozenFunds"`
	AvailableBalance float64 `json:"availableBalance"`
}

// An AccountsModel is the set of *AccountModel.
type AccountsModel []*AccountModel

// AccountOverview returns a list of accounts.
// See the Deposits section for documentation on how to deposit funds to begin trading.
func (as *ApiService) AccountOverview() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/account-overview", nil)
	return as.Call(req)
}

// A TransactionHistoryModel represents a sub-account user.
type TransactionHistoryModel struct {
	Time          string `json:"time"`
	Type          string `json:"type"`
	Amount        string `json:"amount"`
	Fee           string `json:"fee"`
	AccountEquity string `json:"accountEquity"`
	Status        string `json:"status"`
	Remarks       string `json:"remark"`
	Offset        string `json:"offset"`
}

// An TransactionHistoryListModel the set of *TransactionHistoryModel.
type TransactionHistoryListModel []*TransactionHistoryModel

// TransactionHistory returns a list of ledgers.
// Account activity either increases or decreases your account balance.
// Items are paginated and sorted latest first.
func (as *ApiService) TransactionHistory(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/transaction-history", params)
	return as.Call(req)
}
