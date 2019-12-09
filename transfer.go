package kumex

import "net/http"

// A
type TransferOutModel struct {
	ApplyId string `json:"applyId"`
}

// Transfer Funds to KuCoin-Main Account.
func (as *ApiService) TransferOut(bizNo, amount string) (*ApiResponse, error) {
	p := map[string]string{
		"bizNo":  bizNo,
		"amount": amount,
	}
	req := NewRequest(http.MethodPost, "/api/v1/transfer-out", p)
	return as.Call(req)
}

// A DepositModel represents a deposit record.
type TransferModel struct {
	ApplyId   string `json:"applyId"`
	Currency  string `json:"currency"`
	Status    string `json:"status"`
	Amount    string `json:"amount"`
	Reason    string `json:"reason"`
	Offset    int64  `json:"offset"`
	CreatedAt int64  `json:"createdAt"`
}

type TransfersModel []*TransferModel

// Deposits returns a list of deposit.
func (as *ApiService) TransferList(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/transfer-list", params)
	return as.Call(req)
}

// CancelWithdrawalResultModel represents the result of CancelWithdrawal().
type CancelTransferModel struct {
	ApplyId string `json:"applyId"`
}

// Cancel Transfer-Out Request.
func (as *ApiService) CancelTransfer(applyId string) (*ApiResponse, error) {
	p := map[string]string{
		"applyId": applyId,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/cancel/transfer-out", p)
	return as.Call(req)
}
