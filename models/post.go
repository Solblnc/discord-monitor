package models

type Post struct {
	ProMerchantAds bool        `json:"proMerchantAds"`
	Page           int         `json:"page"`
	Rows           int         `json:"rows"`
	PayTypes       []string    `json:"payTypes"`
	Countries      []string    `json:"countries"`
	PublisherType  interface{} `json:"publisherType"`
	Asset          string      `json:"asset"`
	Fiat           string      `json:"fiat"`
	TradeType      string      `json:"tradeType"`
}
