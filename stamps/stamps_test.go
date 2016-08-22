package stamps_test

import (
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
)

func TestStampDataString(t *testing.T) {
	t.Log(stampData.String())
}
func TestStampDataJSON(t *testing.T) {
	t.Log(stampData.JSON())
}

func TestStampInventoryString(t *testing.T) {
	t.Log(stampInventory.String())
}
func TestStampInventoryJSON(t *testing.T) {
	t.Log(stampInventory.JSON())
}
