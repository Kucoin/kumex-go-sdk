package main

import (
	"github.com/Kucoin/kumex-go-sdk"
	"log"
)

func main() {
	//s := kumex.NewApiServiceFromEnv()
	s := kumex.NewApiService(
		kumex.ApiKeyOption(""),
		kumex.ApiSecretOption(""),
		kumex.ApiPassPhraseOption(""),
	)
	serverTime(s)
	accounts(s)
	orders(s)
	websocket(s)

}

func serverTime(s *kumex.ApiService) {
	rsp, err := s.ServerTime()
	if err != nil {
		log.Printf("Error: %s", err.Error())
		// Handle error
		return
	}

	var ts int64
	if err := rsp.ReadData(&ts); err != nil {
		// Handle error
		return
	}
	log.Printf("The server time: %d", ts)
}

func accounts(s *kumex.ApiService) {
	rsp, err := s.AccountOverview()
	if err != nil {
		// Handle error
		return
	}

	as := kumex.AccountsModel{}
	if err := rsp.ReadData(&as); err != nil {
		// Handle error
		return
	}

	for _, a := range as {
		log.Printf("Available balance: %s %s => %s", a.AccountEquity, a.OrderMargin, a.AvailableBalance)
	}
}

func orders(s *kumex.ApiService) {
	rsp, err := s.Orders(map[string]string{}, &kumex.PaginationParam{CurrentPage: 1, PageSize: 10})
	if err != nil {
		// Handle error
		return
	}

	os := kumex.OrdersModel{}
	pa, err := rsp.ReadPaginationData(&os)
	if err != nil {
		// Handle error
		return
	}
	log.Printf("Total num: %d, total page: %d", pa.TotalNum, pa.TotalPage)
	for _, o := range os {
		log.Printf("Order: %s, %s, %s", o.Id, o.Type, o.Price)
	}
}
func websocket(s *kumex.ApiService) {
	rsp, err := s.WebSocketPublicToken()
	if err != nil {
		// Handle error
		return
	}

	tk := &kumex.WebSocketTokenModel{}
	if err := rsp.ReadData(tk); err != nil {
		// Handle error
		return
	}

	c := s.NewWebSocketClient(tk)

	mc, ec, err := c.Connect()
	if err != nil {
		// Handle error
		return
	}

	ch1 := kumex.NewSubscribeMessage("/contractMarket/ticker:XBTUSDM", false)
	ch2 := kumex.NewSubscribeMessage("/contractMarket/ticker:XBTUSDM", false)
	uch := kumex.NewUnsubscribeMessage("/contractMarket/ticker:XBTUSDM", false)

	if err := c.Subscribe(ch1, ch2); err != nil {
		// Handle error
		return
	}

	var i = 0
	for {
		select {
		case err := <-ec:
			c.Stop() // Stop subscribing the WebSocket feed
			log.Printf("Error: %s", err.Error())
			// Handle error
			return
		case msg := <-mc:
			// log.Printf("Received: %s", kumex.ToJsonString(m))
			t := &kumex.TickerLevel1Model{}
			if err := msg.ReadData(t); err != nil {
				log.Printf("Failure to read: %s", err.Error())
				return
			}
			log.Printf("Ticker: %s, %s, %s, %s", msg.Topic, t.Sequence, t.Price, t.Size)
			i++
			if i == 5 {
				log.Println("Unsubscribe XBTUSDM")
				if err = c.Unsubscribe(uch); err != nil {
					log.Printf("Error: %s", err.Error())
					// Handle error
					return
				}
			}
			if i == 10 {
				log.Println("Subscribe XBTUSDM")
				if err = c.Subscribe(ch2); err != nil {
					log.Printf("Error: %s", err.Error())
					// Handle error
					return
				}
			}
			if i == 15 {
				log.Println("Exit subscription")
				c.Stop() // Stop subscribing the WebSocket feed
				return
			}
		}
	}
}
