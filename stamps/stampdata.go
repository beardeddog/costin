package stamps

import (
	"encoding/json"
	"fmt"
	"time"
)

// StampData stuct contains information about each coin
type StampData struct {
	DocumentType string    `json:"documenttype"` // stampdata
	IDNumber     string    `json:"idnumber"`     // scott number
	Country      string    `json:"country"`      // USA
	Name         string    `json:"name"`         // Lewis and Clark Expedition
	Description  string    `json:"description"`  // subject
	Series       string    `json:"series"`       // Famous Americans
	Keywords     string    `json:"keywords"`     // ...
	Denomination string    `json:"denomination"` // 80c, 1.00, ..
	DateChanged  time.Time `json:"datechanged"`  // auto generated
}

/*
{
  "documenttype": "stampdata",
	"idnumber": "C46",
	"country": "USA",
	"name": "Diamond Head, Honolulu, Hawaii",
	"description": "",
	"series": "",
	"keywords": "",
	"denomination": "80c",
	"datechanged": "2016-08-13 14:34:23"
}
*/

// String ...
func (c *StampData) String() string {
	return fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %v",
		c.DocumentType, c.IDNumber, c.Country, c.Description, c.Name, c.Series,
		c.Keywords, c.Denomination, c.DateChanged)
}

// JSON ...
func (c *StampData) JSON() string {
	js, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		fmt.Println("error MarshalIndent the StampData struct")
		return ""
	}
	return string(js)
}
