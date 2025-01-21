package app

import (
	"fetch-project/model"
	"fetch-project/schema"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-errors/errors"
)

var data = make(map[string]model.Reciept)

func Process(opts *schema.Reciept) (map[string]string, error) {

	now := time.Now().UTC()
	id := strconv.FormatInt(now.UnixMilli(), 10)

	purchaseDate, err := time.Parse("2006-01-02", opts.PurchaseDate)
	if err != nil {
		return nil, errors.New("The receipt is invalid.")
	}

	purchaseTime, err := time.Parse("15:04", opts.PurchaseTime)
	if err != nil {
		return nil, errors.New("The receipt is invalid.")
	}

	total, err := strconv.ParseFloat(opts.Total, 64)
	if err != nil {
		return nil, errors.New("The receipt is invalid.")
	}

	reciept := model.Reciept{
		ID:           id,
		Retailer:     opts.Retailer,
		PurchaseDate: purchaseDate,
		PurchaseTime: purchaseTime,
		Total:        total,
	}

	var items []model.Item

	for _, item := range opts.Items {

		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return nil, errors.New("The receipt is invalid.")
		}
		newItem := model.Item{
			ShortDescription: strings.TrimSpace(item.ShortDescription),
			Price:            price,
		}
		items = append(items, newItem)
	}
	reciept.Items = items

	data[id] = reciept

	resp := map[string]string{
		"id": id,
	}

	return resp, nil
}

func GetPoint(id string) (map[string]int, error, int) {

	reciept, exists := data[id]
	if !exists {
		return nil, errors.New("No receipt found for that ID."), http.StatusNotFound
	}
	points := 0

	points += countAlphanumeric(reciept.Retailer)
	if math.Floor(reciept.Total) == float64(reciept.Total) {
		points += 50
	}

	if math.Mod(reciept.Total, 0.25) == 0 {
		points += 25
	}

	pairs := len(reciept.Items) / 2
	points += pairs * 5

	for _, item := range reciept.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			points += int(math.Ceil(item.Price * 0.2))
		}
	}

	if reciept.PurchaseDate.Day()%2 != 0 {
		points += 6
	}

	hour := reciept.PurchaseTime.Hour()
	minute := reciept.PurchaseTime.Minute()

	if (hour == 14 && minute > 0) || (hour > 14 && hour < 16) {
		points += 10
	}
	resp := map[string]int{
		"points": points,
	}
	return resp, nil, http.StatusOK
}
