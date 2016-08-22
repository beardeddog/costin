package coins

import (
	"encoding/json"
	"fmt"
	"time"
)

// CoinInventory struct contains information about what is owned
type CoinInventory struct {
	DocumentType    string    `json:"documenttype"`    // coininventory
	IDNumber        string    `json:"idnumber"`        // km number
	Year            string    `json:"year"`            // 1994
	Mint            string    `json:"mint"`            // S, D, P, (P)
	Variant         string    `json:"variant"`         // 7 feathers
	Notes           string    `json:"notes"`           // from KDW collection
	PurchasePrice   float32   `json:"purchaseprice"`   // 8.50
	PurchaseDate    string    `json:"purchasedate"`    // 12-23-2014
	StorageLocation string    `json:"storagelocation"` // dansco
	Grade           string    `json:"grade"`           // AU-50, F-12, PRF-65
	Slabbed         string    `json:"slabbed"`         // yes or no
	QuantityOwned   int       `json:"quanityowned"`    // 1
	DateChanged     time.Time `json:"datechanged"`     // auto generated
}

/*
{
	"documenttype":"coininventory",
	"idnumber": "202",
	"year": "1964"
	"mint": "",
	"variant": "",
	"notes":"",
	"purchaseprice": 0.25,
	"purchasedate": "",
	"storagelocation":"Dansco",
	"grade": "XF-40",
	"slabbed": "no"
	"quantityowned": 1,
	"datechanged": "2016-08-13 14:34:23"
}
*/

// String ...
func (c *CoinInventory) String() string {
	return fmt.Sprintf("%s, %s, %s, %s, %s, %s, %.2f, %s, %s, %s, %s, %d, %v",
		c.DocumentType, c.IDNumber, c.Year, c.Mint, c.Variant, c.Notes,
		c.PurchasePrice, c.PurchaseDate, c.StorageLocation, c.Grade,
		c.Slabbed, c.QuantityOwned, c.DateChanged)
}

// JSON ...
func (c *CoinInventory) JSON() string {
	js, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		fmt.Println("error MarshalIndent the CoinInventory struct")
		return ""
	}
	return string(js)
}
