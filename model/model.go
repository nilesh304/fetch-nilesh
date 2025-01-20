package model

import (
	"time"
)

type Item struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price"`
}

type Reciept struct {
	ID           string    `json:"id"`
	Retailer     string    `json:"retailer,omitempty"`
	PurchaseDate time.Time `json:"purchaseDate"`
	PurchaseTime time.Time `json:"purchaseTime"`
	Total        float64   `json:"total"`
	Items        []Item    `json:"items"`
}
