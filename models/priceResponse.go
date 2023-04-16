package models

type Response struct {
	Data []struct {
		Adv struct {
			Price string `json:"price"`
		} `json:"adv"`
	} `json:"data"`
}
