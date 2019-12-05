package kumex

import (
	"net/http"
)

// A TickerLevel1Model represents ticker include only the inside (i.e. best) bid and ask data, last price and last trade size.
type TickerLevel1Model struct {
	Sequence     string `json:"sequence"`
	Symbol       string `json:"symbol"`
	Side         string `json:"side"`
	Size         int    `json:"size"`
	Price        string `json:"price"`
	BestBidSize  string `json:"bestBidSize"`
	BestBidPrice string `json:"bestBidPrice"`
	BestAskSize  string `json:"bestAskSize"`
	BestAskPrice string `json:"bestAskPrice"`
	TradeId      string `json:"tradeId"`
	Ts           int64  `json:"ts"`
}

// TickerLevel1 returns the ticker include only the inside (i.e. best) bid and ask data, last price and last trade size.
func (as *ApiService) Ticker(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/ticker", map[string]string{"symbol": symbol})
	return as.Call(req)
}

type Level2SnapshotModel struct {
	Symbol   string     `json:"symbol"`
	Sequence string     `json:"sequence"`
	Asks     [][]string `json:"asks"`
	Bids     [][]string `json:"bids"`
}

// TickerLevel1 returns the ticker include only the inside (i.e. best) bid and ask data, last price and last trade size.
func (as *ApiService) Level2Snapshot(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/level2/snapshot", map[string]string{"symbol": symbol})
	return as.Call(req)
}

type Level2MessageQueryModel struct {
	Symbol   string `json:"symbol"`
	Sequence string `json:"sequence"`
	Change   string `json:"change"`
}

type Level2MessageQueryListModel []*Level2MessageQueryModel

func (as *ApiService) Level2MessageQuery(symbol string, start, end int64) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/level2/message/query", map[string]string{
		"symbol": symbol,
		"start":  IntToString(start),
		"end":    IntToString(end),
	})
	return as.Call(req)
}

type Level3SnapshotModel struct {
	Symbol   string     `json:"symbol"`
	Sequence string     `json:"sequence"`
	Asks     [][]string `json:"asks"`
	Bids     [][]string `json:"bids"`
}

func (as *ApiService) Level3Snapshot(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/level3/snapshot", map[string]string{"symbol": symbol})
	return as.Call(req)
}

type Level3MessageQueryModel struct {
	Symbol    string `json:"symbol"`
	Sequence  string `json:"sequence"`
	Side      string `json:"side"`
	OrderTime int64  `json:"orderTime"`
	Size      int    `json:"size"`
	OrderId   string `json:"orderId"`
	Price     string `json:"price"`
	Type      string `json:"type"`
	ClientOid string `json:"clientOid"`
	Ts        int64  `json:"ts"`
}

type Level3MessageQueryListModel []*Level3MessageQueryModel

func (as *ApiService) Level3MessageQuery(symbol string, start, end int64) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/level3/message/query", map[string]string{
		"symbol": symbol,
		"start":  IntToString(start),
		"end":    IntToString(end),
	})
	return as.Call(req)
}

// A TradeHistoryModel represents a the latest trades for a symbol.
type TradeHistoryModel struct {
	Sequence     string `json:"sequence"`
	TradeId      string `json:"tradeId"`
	TakerOrderId string `json:"takerOrderId"`
	MakerOrderId string `json:"makerOrderId"`
	Price        string `json:"price"`
	Size         string `json:"size"`
	Side         string `json:"side"`
	Time         int64  `json:"t"`
}

// A TradeHistoriesModel is the set of *TradeHistoryModel.
type TradesHistoryModel []*TradeHistoryModel

// TradeHistories returns a list the latest trades for a symbol.
func (as *ApiService) TradeHistory(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/trade/history", map[string]string{"symbol": symbol})
	return as.Call(req)
}

type InterestModel struct {
	Symbol      string `json:"symbol"`
	Granularity string `json:"granularity"`
	TimePoint   string `json:"timePoint"`
	Value       string `json:"value"`
}

type InterestsModel []*InterestModel

// Deposits returns a list of deposit.
func (as *ApiService) InterestQuery(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/interest/query", params)
	return as.Call(req)
}

type IndexModel struct {
	Symbol          string     `json:"symbol"`
	Granularity     string     `json:"granularity"`
	TimePoint       string     `json:"timePoint"`
	Value           string     `json:"value"`
	DecomposionList [][]string `json:"decomposionList"`
}

type IndexQueryModel []*IndexModel

// Deposits returns a list of deposit.
func (as *ApiService) IndexQuery(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/interest/query", params)
	return as.Call(req)
}

type MarkPriceModel struct {
	Symbol      string `json:"symbol"`
	Granularity string `json:"granularity"`
	TimePoint   string `json:"timePoint"`
	Value       string `json:"value"`
	IndexPrice  string `json:"indexPrice"`
}

func (as *ApiService) MarkPrice(Symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/mark-price/"+Symbol+"/current", nil)
	return as.Call(req)
}

type PremiumModel struct {
	Symbol          string     `json:"symbol"`
	Granularity     string     `json:"granularity"`
	TimePoint       string     `json:"timePoint"`
	Value           string     `json:"value"`
}

type PremiumsModel []*PremiumModel

// Deposits returns a list of deposit.
func (as *ApiService) PremiumQuery(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/premium/query", params)
	return as.Call(req)
}


type FundingRateModel struct {
	Symbol      string `json:"symbol"`
	Granularity string `json:"granularity"`
	TimePoint   string `json:"timePoint"`
	Value       string `json:"value"`
	PredictedValue  string `json:"predictedValue"`
}

func (as *ApiService) FundingRate(Symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/funding-rate/"+Symbol+"/current", nil)
	return as.Call(req)
}
