package coins

import (
	"encoding/json"
	"fmt"
	"time"
)

// CoinData stuct contains information about each coin
type CoinData struct {
	DocumentType string    `json:"documenttype"` // coindata
	IDNumber     string    `json:"idnumber"`     // km number
	Country      string    `json:"country"`      // km number
	Coin         string    `json:"coin"`         // dollar, half dollar, ...
	Series       string    `json:"series"`       // washington head, wheat ears
	Denomination string    `json:"denomination"` // 1, 0.50,
	Metal        string    `json:"metal"`        // silver
	PercentMetal float64   `json:"percentmetal"` // 0.90
	WeightGrams  float64   `json:"weightgrams"`  // 26.9600
	DiameterMM   float64   `json:"diametermm"`   // 39.0
	DateChanged  time.Time `json:"datechanged"`  // auto generated
}

/*
{
	"documenttype":"coindata",
	"idnumber": "202",
	"country": "USA",
	"coin": "half dollar",
	"series": "Kennedy",
	"denomination": "0.50",
	"metal": "silver",
	"percentmetal": 0.90,
	"weightgrams": 12.5,
	"diametermm": 30.6,
	"datechanged": "2016-08-13 14:34:23"
}
*/

// String ...
func (c *CoinData) String() string {
	return fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %f, %f, %f, %v",
		c.DocumentType, c.IDNumber, c.Country, c.Coin, c.Series, c.Denomination,
		c.Metal, c.PercentMetal, c.WeightGrams, c.DiameterMM, c.DateChanged)
}

// JSON ...
func (c *CoinData) JSON() string {
	js, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		fmt.Println("error MarshalIndent the CoinData struct")
		return ""
	}
	return string(js)
}
