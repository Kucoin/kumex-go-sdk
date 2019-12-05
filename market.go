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
type TradeHistoriesModel []*TradeHistoryModel

// TradeHistories returns a list the latest trades for a symbol.
func (as *ApiService) TradeHistories(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/trade/history", map[string]string{"symbol": symbol})
	return as.Call(req)
}


// A TickerModel represents a market ticker for all trading pairs in the market (including 24h volume).
type TickerModel struct {
	Symbol      string `json:"symbol"`
	Buy         string `json:"buy"`
	Sell        string `json:"sell"`
	ChangeRate  string `json:"changeRate"`
	ChangePrice string `json:"changePrice"`
	High        string `json:"high"`
	Low         string `json:"low"`
	Vol         string `json:"vol"`
	VolValue    string `json:"volValue"`
	Last        string `json:"last"`
}

//// A TickersModel is the set of *MarketTickerModel.
//type TickersModel []*TickerModel
//
//// TickersResponseModel represents the response model of MarketTickers().
//type TickersResponseModel struct {
//	Time    int64        `json:"time"`
//	Tickers TickersModel `json:"ticker"`
//}
//
//// Tickers returns all tickers as TickersResponseModel for all trading pairs in the market (including 24h volume).
//func (as *ApiService) Tickers(symbol string) (*ApiResponse, error) {
//	req := NewRequest(http.MethodGet, "api/v1/ticker", map[string]string{"symbol": symbol})
//	return as.Call(req)
//}

// A Stats24hrModel represents 24 hr stats for the symbol.
// Volume is in base currency units.
// Open, high, low are in quote currency units.
type Stats24hrModel struct {
	Symbol      string `json:"symbol"`
	ChangeRate  string `json:"changeRate"`
	ChangePrice string `json:"changePrice"`
	Open        string `json:"open"`
	Close       string `json:"close"`
	High        string `json:"high"`
	Low         string `json:"low"`
	Vol         string `json:"vol"`
	VolValue    string `json:"volValue"`
}

// Stats24hr returns 24 hr stats for the symbol. volume is in base currency units. open, high, low are in quote currency units.
func (as *ApiService) Stats24hr(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/stats", map[string]string{"symbol": symbol})
	return as.Call(req)
}

// Markets returns the transaction currencies for the entire trading market.
func (as *ApiService) Markets() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/markets", nil)
	return as.Call(req)
}

// A PartOrderBookModel represents a list of open orders for a symbol, a part of Order Book within 100 depth for each side(ask or bid).
type PartOrderBookModel struct {
	Sequence string     `json:"sequence"`
	Bids     [][]string `json:"bids"`
	Asks     [][]string `json:"asks"`
}

// AggregatedPartOrderBook returns a list of open orders(aggregated) for a symbol.
func (as *ApiService) AggregatedPartOrderBook(symbol string, depth int64) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/orderbook/level2_"+IntToString(depth), map[string]string{"symbol": symbol})
	return as.Call(req)
}

// A FullOrderBookModel represents a list of open orders for a symbol, with full depth.
type FullOrderBookModel struct {
	Sequence string     `json:"sequence"`
	Bids     [][]string `json:"bids"`
	Asks     [][]string `json:"asks"`
}

// AggregatedFullOrderBook returns a list of open orders(aggregated) for a symbol.
func (as *ApiService) AggregatedFullOrderBook(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v2/market/orderbook/level2", map[string]string{"symbol": symbol})
	return as.Call(req)
}

// AtomicFullOrderBook returns a list of open orders for a symbol.
// Level-3 order book includes all bids and asks (non-aggregated, each item in Level-3 means a single order).
func (as *ApiService) AtomicFullOrderBook(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/orderbook/level3", map[string]string{"symbol": symbol})
	return as.Call(req)
}

// KLineModel represents the k lines for a symbol.
// Rates are returned in grouped buckets based on requested type.
type KLineModel []string

// A KLinesModel is the set of *KLineModel.
type KLinesModel []*KLineModel

// KLines returns the k lines for a symbol.
// Data are returned in grouped buckets based on requested type.
// Parameter #2 typo is the type of candlestick patterns.
func (as *ApiService) KLines(symbol, typo string, startAt, endAt int64) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/candles", map[string]string{
		"symbol":  symbol,
		"type":    typo,
		"startAt": IntToString(startAt),
		"endAt":   IntToString(endAt),
	})
	return as.Call(req)
}
