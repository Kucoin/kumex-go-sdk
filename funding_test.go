package kumex

import "testing"

func TestApiService_FundingHistory(t *testing.T) {
	s := NewApiServiceFromEnv()
	p := &PaginationParam{CurrentPage: 1, PageSize: 10}
	rsp, err := s.FundingHistory(map[string]string{"symbol":"XBTUSDM"}, p)
	if err != nil {
		t.Fatal(err)
	}

	os := FundingListModel{}
	if _, err := rsp.ReadPaginationData(&os); err != nil {
		t.Fatal(err)
	}
	for _, o := range os {
		t.Log(ToJsonString(o))
		switch {
		case o.Id == "":
			t.Error("Empty key 'id'")
		case o.Symbol == "":
			t.Error("Empty key 'symbol'")
		case o.MarkPrice == "":
			t.Error("Empty key 'markPrice'")
		case o.FundingRate == "":
			t.Error("Empty key 'fundingRate'")
		case o.Funding == "":
			t.Error("Empty key 'funding'")
		}
	}
}
