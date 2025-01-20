package schema

type Item struct {
	ShortDescription string `json:"shortDescription" validate:"required,validateRegex=^[\\w\\s\\-]+$"`
	Price            string `json:"price" validate:"required,validateRegex=^\\d+\\.\\d{2}$"`
}

type Reciept struct {
	Retailer     string `json:"retailer,omitempty" validate:"required,validateRegex=^[\\w\\s\\-&]+$"`
	PurchaseDate string `json:"purchaseDate" validate:"required,validateDate=2006-01-02"`
	PurchaseTime string `json:"purchaseTime" validate:"required,validateTime=15:04"`
	Total        string `json:"total" validate:"required,validateRegex=^\\d+\\.\\d{2}$"`
	Items        []Item `json:"items"  validate:"min=1,dive"`
}
