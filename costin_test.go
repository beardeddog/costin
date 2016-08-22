package main

import (
	"costin/coins"
	"costin/config"
	"costin/stamps"
	"testing"
	"time"
)

var (
	stampData = stamps.StampData{
		DocumentType: "stampdata",
		IDNumber:     "C46",
		Country:      "USA",
		Name:         "Diamond Head, Honolulu, Hawaii",
		Description:  "",
		Series:       "",
		Keywords:     "",
		Denomination: "80c",
		DateChanged:  time.Now(),
	}
	stampInventory = stamps.StampInventory{
		DocumentType:         "stampinventory",
		IDNumber:             "C46",
		Format:               "block",
		FormatNumberOfStamps: 4,
		PlateNumber:          "23045",
		PostalMarkings:       "",
		AdditionalMarkings:   "",
		SelvageMarks:         "",
		PurchasePrice:        17.95,
		PurchaseDate:         "08-13-2016",
		StorageLocation:      "blue binder",
		Used:                 "no",
		Condition:            "very fine",
		Gum:                  "original",
		Faults:               "",
		Slabbed:              "no",
		QuantityOwned:        1,
		DateChanged:          time.Now(),
	}
	coinData = coins.CoinData{
		DocumentType: "coindata",
		IDNumber:     "202",
		Country:      "USA",
		Coin:         "half dollar",
		Series:       "Kennedy",
		Denomination: "0.50",
		Metal:        "silver",
		PercentMetal: 0.90,
		WeightGrams:  12.5,
		DiameterMM:   30.6,
		DateChanged:  time.Now(),
	}
	coinInventory = coins.CoinInventory{
		DocumentType:    "coininventory",
		IDNumber:        "202",
		Year:            "1964",
		Mint:            "",
		Variant:         "",
		Notes:           "",
		PurchasePrice:   1.15,
		PurchaseDate:    "08-12-2016",
		StorageLocation: "Dansco",
		Grade:           "XF-40",
		Slabbed:         "no",
		QuantityOwned:   1,
		DateChanged:     time.Now(),
	}

	values = []string{"B", "C"}
	cfg    = config.Config{
		Type:       "config",
		Collection: "stamps",
		Field:      "Gum",
		Values:     values,
	}
)

func TestString(t *testing.T) {
	t.Log(stampData.String())
	t.Log(stampInventory.String())
	t.Log(coinData.String())
	t.Log(coinInventory.String())
	t.Log(cfg.String())
}
func TestJSON(t *testing.T) {
	t.Log(stampData.JSON())
	t.Log(stampInventory.JSON())
	t.Log(coinData.JSON())
	t.Log(coinInventory.JSON())
	t.Log(cfg.JSON())
}
