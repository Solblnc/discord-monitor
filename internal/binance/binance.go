package binance

import (
	"bytes"
	"discord-monitor/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func PostReq() error {
	// HTTP endpoint
	url := "https://p2p.binance.com/bapi/c2c/v2/friendly/c2c/adv/search"

	body := &models.Post{
		ProMerchantAds: false,
		Page:           1,
		Rows:           2,
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

	//user := "87"
	//password := "123"
	//ip := "33"
	//port := "20"
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

	var prettyJson bytes.Buffer
	json.Indent(&prettyJson, data, "", "\t")

	var buf models.Response

	json.Unmarshal(data, &buf)

	fmt.Println(buf)

	return nil
}
