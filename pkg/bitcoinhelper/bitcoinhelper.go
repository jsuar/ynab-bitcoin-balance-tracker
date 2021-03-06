package bitcoinhelper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"

	"github.com/jsuar/ynab-bitcoin-balance-tracker/pkg/envhelper"
	"go.uber.org/zap"
)

// BitcoinHelper provides helper functions
type BitcoinHelper struct {
	Ticker         *Ticker
	BitcoinAddress *BitcoinAddress
	verbose        bool
	logger         *zap.SugaredLogger
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
func (bh *BitcoinHelper) Init(verbose bool, logger *zap.SugaredLogger) {
	bh.verbose = verbose
	bh.logger = logger
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
func (bh *BitcoinHelper) GetMarketPrice(currency string, field string) (float64, error) {
	var price float64

	url := "https://blockchain.info/ticker"
	response, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("The HTTP request failed with error %s", err)
	}
	if response.StatusCode != 200 {
		return 0, fmt.Errorf("Request to %s failed with status code: %d", url, response.StatusCode)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	var ticker Ticker
	json.Unmarshal(data, &ticker)

	price, err = bh.getField(&ticker, currency, field)
	if err != nil {
		return 0, err
	}

	return price, nil
}

// BitcoinAddress contains data from a bitcoin address such as balance
type BitcoinAddress struct {
	ErrNo int `json:"err_no"`
	Data  struct {
		Address  string `json:"address"`
		Received int    `json:"received"`
		Sent     int    `json:"sent"`
		Balance  int    `json:"balance"`
		// Balance             int    `json:"banlance"`
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
func (bh *BitcoinHelper) ShowAddressBalance() error {
	envVar := "BITCOIN_ADDR"
	btcAddr, err := envhelper.GetRequiredEnv(envVar)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	url := fmt.Sprintf("https://chain.api.btc.com/v3/address/%s", btcAddr)
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("The HTTP request failed with error %s", err)
	}
	if response.StatusCode != 200 {
		return fmt.Errorf("Request to %s failed with status code: %d", url, response.StatusCode)
	}

	data, _ := ioutil.ReadAll(response.Body)
	var btcAddress BitcoinAddress
	err = json.Unmarshal(data, &btcAddress)
	if err != nil {
		return err
	}

	fmt.Println("Btc", float64(btcAddress.Data.Balance)/100000000.0)
	fmt.Println("mBtc", float64(btcAddress.Data.Balance)/100000.0)

	return nil
}

// GetAddressBalance ...
func (bh *BitcoinHelper) GetAddressBalance() (int64, error) {
	var balance int64

	btcAddr, err := envhelper.GetRequiredEnv("BITCOIN_ADDR")
	bh.logger.Infof("BITCOIN_ADDR=%s", btcAddr)
	if err != nil {
		return 0, err
	}

	url := fmt.Sprintf("https://chain.api.btc.com/v3/address/%s", btcAddr)
	bh.logger.Infof("url: %s", url)
	response, err := http.Get(url)
	bh.logger.Infof("response: %s\n", response)
	bh.logger.Infof("status: %s\n", response.Status)
	if err != nil {
		return 0, fmt.Errorf("The HTTP request failed with error %s", err)
	}
	if response.StatusCode != 200 {
		return 0, fmt.Errorf("Request to %s failed with status code: %d", url, response.StatusCode)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, fmt.Errorf("ReadAll failed with: %s", err)
	}
	bh.logger.Infof("body data: %s\n", data)
	var btcAddress BitcoinAddress
	err = json.Unmarshal(data, &btcAddress)
	if err != nil {
		return 0, err
	}

	bh.logger.Infof("unmarshalled balance: %d\n", btcAddress.Data.Balance)
	balance = int64(btcAddress.Data.Balance)

	return balance, nil
}
