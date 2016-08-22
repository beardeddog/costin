package coins_test

import (
	"costin/coins"
	"testing"
	"time"
)

var (
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
)

func TestCoinDataString(t *testing.T) {
	t.Log(coinData.String())
}
func TestCoinDataJSON(t *testing.T) {
	t.Log(coinData.JSON())
}

func TestCoinInventoryString(t *testing.T) {
	t.Log(coinInventory.String())
}
func TestCoinInventoryJSON(t *testing.T) {
	t.Log(coinInventory.JSON())
}
