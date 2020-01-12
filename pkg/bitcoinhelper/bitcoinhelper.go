package bitcoinhelper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
)

// BitcoinHelper provides helper functions
type BitcoinHelper struct {
	Ticker         *Ticker
	BitcoinAddress *BitcoinAddress
	verbose        bool
}

// Ticker stores different currency bitcoin data
type Ticker struct {
	USD struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"USD"`
	AUD struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"AUD"`
	BRL struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"BRL"`
	CAD struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"CAD"`
	CHF struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"CHF"`
	CLP struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"CLP"`
	CNY struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"CNY"`
	DKK struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"DKK"`
	EUR struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"EUR"`
	GBP struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"GBP"`
	HKD struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"HKD"`
	INR struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"INR"`
	ISK struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"ISK"`
	JPY struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"JPY"`
	KRW struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"KRW"`
	NZD struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"NZD"`
	PLN struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"PLN"`
	RUB struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"RUB"`
	SEK struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"SEK"`
	SGD struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"SGD"`
	THB struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"THB"`
	TWD struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"TWD"`
}

// Init will initialize the BitcoinHelper object
func (bh *BitcoinHelper) Init(verbose bool) {
	bh.verbose = verbose
}

func (bh *BitcoinHelper) getField(v *Ticker, currency string, field string) (float64, error) {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(currency)
	if f.IsValid() {
		field := f.FieldByName(field)
		if field.IsValid() {
			return field.Float(), nil
		}
	}

	return 0, fmt.Errorf("Field %s does not exist", field)
}

// GetMarketPrice ...
func (bh *BitcoinHelper) GetMarketPrice(currency string, field string) float64 {
	var price float64
	price = 0.0

	response, err := http.Get("https://blockchain.info/ticker")
	if err != nil {
		panic(fmt.Sprintf("The HTTP request failed with error %s\n", err))
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		// fmt.Println(string(data))

		var ticker Ticker
		json.Unmarshal(data, &ticker)

		price, err = bh.getField(&ticker, currency, field)
		if err != nil {
			panic(err)
		}

	}

	return price
}

// BitcoinAddress contains data from a bitcoin address such as balance
type BitcoinAddress struct {
	ErrNo int `json:"err_no"`
	Data  struct {
		Address             string `json:"address"`
		Received            int    `json:"received"`
		Sent                int    `json:"sent"`
		Balance             int    `json:"balance"`
		TxCount             int    `json:"tx_count"`
		UnconfirmedTxCount  int    `json:"unconfirmed_tx_count"`
		UnconfirmedReceived int    `json:"unconfirmed_received"`
		UnconfirmedSent     int    `json:"unconfirmed_sent"`
		UnspentTxCount      int    `json:"unspent_tx_count"`
		FirstTx             string `json:"first_tx"`
		LastTx              string `json:"last_tx"`
	} `json:"data"`
}

// ShowAddressBalance ...
func (bh *BitcoinHelper) ShowAddressBalance() {
	btcAddr := os.Getenv("BITCOIN_ADDR")
	url := fmt.Sprintf("https://chain.api.btc.com/v3/address/%s", btcAddr)
	response, err := http.Get(url)
	if err != nil {
		panic(fmt.Sprintf("The HTTP request failed with error %s\n", err))
	}
	data, _ := ioutil.ReadAll(response.Body)

	var btcAddress BitcoinAddress
	json.Unmarshal(data, &btcAddress)
	fmt.Println("mBtc", float64(btcAddress.Data.Balance)/100000.0)
	fmt.Println("Btc", float64(btcAddress.Data.Balance)/100000000.0)
}

// GetAddressBalance ...
func (bh *BitcoinHelper) GetAddressBalance() int64 {
	var balance int64
	balance = 0.0

	btcAddr := os.Getenv("BITCOIN_ADDR")
	url := fmt.Sprintf("https://chain.api.btc.com/v3/address/%s", btcAddr)
	response, err := http.Get(url)
	if err != nil {
		panic(fmt.Sprintf("The HTTP request failed with error %s\n", err))
	}
	data, _ := ioutil.ReadAll(response.Body)

	var btcAddress BitcoinAddress
	json.Unmarshal(data, &btcAddress)
	balance = int64(btcAddress.Data.Balance)

	return balance
}
