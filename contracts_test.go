package kumex

import "testing"

func TestApiService_ActiveContracts(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.ActiveContracts()
	if err != nil {
		t.Fatal(err)
	}
	o := &ContractsModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case o.BaseCurrency == "":
		t.Error("Empty key 'baseCurrency'")
	case o.Symbol == "":
		t.Error("Empty key 'symbol'")
	case o.FairMethod == "":
		t.Error("Empty key 'fairMethod'")
	case o.IndexSymbol == "":
		t.Error("Empty key 'indexSymbol'")
	}
}

func TestApiService_Contracts(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.Contracts("XBTUSDM")
	if err != nil {
		t.Fatal(err)
	}
	o := &ContractsModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case o.BaseCurrency == "":
		t.Error("Empty key 'baseCurrency'")
	case o.Symbol == "":
		t.Error("Empty key 'symbol'")
	case o.FairMethod == "":
		t.Error("Empty key 'fairMethod'")
	case o.IndexSymbol == "":
		t.Error("Empty key 'indexSymbol'")
	}
}
