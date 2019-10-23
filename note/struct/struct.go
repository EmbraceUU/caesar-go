package _struct

import "fmt"

type BoolStruct struct {
	Bool1 bool `json:"bool_1"`
	Bool2 bool `json:"bool_2"`
}

type NetInData struct {
	Id         int32
	NetInValue float64
}

type PairHttpData struct {
	Status PairStatusData `json:"status"`
	Data   []PairData     `json:"data"`
}

type PairStatusData struct {
	TimeStamp    string `json:"timestamp"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage error  `json:"error_message"`
	Elapsed      int    `json:"elapsed"`
	CreditCount  int    `json:"credit_count"`
}

type PairData struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Symbol      string   `json:"symbol"`
	Slug        string   `json:"slug"`
	CirSupply   float64  `json:"circulating_supply"`
	TotalSupply float64  `json:"total_supply"`
	MaxSupply   float64  `json:"max_supply"`
	DataAdded   string   `json:"data_added"`
	NumMarket   int      `json:"num_market_pairs"`
	Tags        []string `json:"tags"`
	CmcRank     int      `json:"cmc_rank"`
	Quote       Quote    `json:"quote"`
}

type Quote struct {
	USD QuoteData `json:"USD"`
}

type QuoteData struct {
	Price        float64 `json:"price"`
	Volume24H    float64 `json:"volume_24h"`
	PerChange1H  float64 `json:"percent_change_1h"`
	PerChange24H float64 `json:"percent_change_24h"`
	PerChange7D  float64 `json:"percent_change_7d"`
	MarketCap    float64 `json:"market_cap"`
	LastUpdate   string  `json:"last_updated"`
}

type CurrencyNetInTrendData struct {
	NetInPeriod int32
	NetInData   []NetInData
}

func GetBoolStruct() CurrencyNetInTrendData {
	var Bs CurrencyNetInTrendData
	return Bs
}

func GetPairHttpData() {
	var PD PairHttpData
	if PD.Data == nil || len(PD.Data) == 0 {
		fmt.Println(true)
		return
	}
	fmt.Println(false)
}
