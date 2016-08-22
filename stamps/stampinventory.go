package stamps

import (
	"encoding/json"
	"fmt"
	"time"
)

// StampInventory struct contains information about what is owned
type StampInventory struct {
	DocumentType         string    `json:"documenttype"`         // stampinventory
	IDNumber             string    `json:"idnumber"`             // scott number
	Format               string    `json:"format"`               // single, block, sheet, strip, ...
	FormatNumberOfStamps int       `json:"formatnumberofstamps"` // 1
	PlateNumber          string    `json:"platenumber"`          // 12345
	PostalMarkings       string    `json:"postalmarkings"`       // none
	AdditionalMarkings   string    `json:"additionalmarkings"`   // none
	SelvageMarks         string    `json:"selvagemarks"`         // none
	PurchasePrice        float32   `json:"purchaseprice"`        // 4.20
	PurchaseDate         string    `json:"purchasedate"`         // 8-12-2016
	StorageLocation      string    `json:"storagelocation"`      // blue binder
	Used                 string    `json:"used"`                 // yes or no
	Condition            string    `json:"condition"`            // never hinged
	Gum                  string    `json:"gum"`                  // original
	Faults               string    `json:"faults"`               // none
	Slabbed              string    `json:"slabbed"`              // no
	QuantityOwned        int       `json:"quanityowned"`         // 1
	DateChanged          time.Time `json:"datechanged"`          // auto generated
}

/*
{
  "documenttype": "stampinventory",
	"idnumber": "C46",
	"format": "block",
	"formatnumberofstamps": 4,
	"platenumber": "23045",
	"postalmarkings": "",
	"additionalmarkings": "",
	"selvagemarks": "",
	"purchaseprice": "",
	"purchasedate": "",
	"storagelocation": "blue binder",
	"used": "no",
	"condition": "very fine",
	"gum": "original",
	"faults": "",
	"slabbed": "no",
	"quantityowned": 1,
  "datechanged": "2016-08-13 14:34:23"
}
*/

// String ...
func (c *StampInventory) String() string {
	return fmt.Sprintf("%s, %s, %s, %d, %s, %s, %s, %s, %.2f, %s, %s, %s, %s, %s, %s, %s, %d, %v",
		c.DocumentType, c.IDNumber, c.Format, c.FormatNumberOfStamps, c.PlateNumber,
		c.PostalMarkings, c.AdditionalMarkings, c.SelvageMarks, c.PurchasePrice,
		c.PurchaseDate, c.StorageLocation, c.Used, c.Condition, c.Gum, c.Faults,
		c.Slabbed, c.QuantityOwned, c.DateChanged)
}

// JSON ...
func (c *StampInventory) JSON() string {
	js, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		fmt.Println("error MarshalIndent the StampInventory struct")
		return ""
	}
	return string(js)
}
