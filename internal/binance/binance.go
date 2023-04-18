package binance

import (
	"bytes"
	"discord-monitor/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func SendReqBuy() (string, error) {
	url := "https://p2p.binance.com/bapi/c2c/v2/friendly/c2c/adv/search"

	body := &models.Post{
		ProMerchantAds: false,
		Page:           1,
		Rows:           1,
		PayTypes:       []string{"TinkoffNew"},
		Countries:      []string{},
		PublisherType:  nil,
		Asset:          "USDT",
		Fiat:           "RUB",
		TradeType:      "BUY",
	}

	b, err := json.Marshal(body)
	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	//var proxy = "https://" + user + ":" + password + "@" + ip + ":" + port
	//var proxy1 = "https://87.123.33.20.8118"
	//Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}, Timeout: time.Second * 5

	// Creating the proxy url
	//proxyUrl, err := url2.Parse(proxy1)
	//if err != nil {
	//	log.Fatal(err)
	//}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)

	var buf models.Response
	json.Unmarshal(data, &buf)

	return buf.Data[0].Adv.Price, nil
}

func SendReqSell() (string, error) {
	url := "https://p2p.binance.com/bapi/c2c/v2/friendly/c2c/adv/search"

	body := &models.Post{
		ProMerchantAds: false,
		Page:           1,
		Rows:           1,
		PayTypes:       []string{"TinkoffNew"},
		Countries:      []string{},
		PublisherType:  nil,
		Asset:          "USDT",
		Fiat:           "RUB",
		TradeType:      "SELL",
	}

	b, err := json.Marshal(body)
	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)

	var buf models.Response
	json.Unmarshal(data, &buf)

	return buf.Data[0].Adv.Price, nil
}
