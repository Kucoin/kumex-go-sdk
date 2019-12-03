package kumex

import "net/http"

type FundingModel struct {
	Id           string `json:"id"`
	Symbol       string `json:"symbol"`
	TimePoint    string `json:"timePoint"`
	FundingRate  string `json:"fundingRate"`
	MarkPrice    string `json:"markPrice"`
	PositionQty  string `json:"positionQty"`
	PositionCost string `json:"positionCost"`
	Funding      string `json:"funding"`
}

// A OrdersModel is the set of *OrderModel.
type FundingListModel []*FundingModel

// Orders returns a list your current orders.
func (as *ApiService) FundingHistory(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/funding-history", params)
	return as.Call(req)
}
