package scheme

import "time"

type Order struct {
	UserID     int       `json:"userID"`
	OrderID    int       `json:"orderID"`
	Items      []Product `json:"items"`
	TotalPrice float64   `json:"totalPrice"`
	Date       time.Time `json:"date"`
}
