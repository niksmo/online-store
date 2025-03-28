package scheme

import "time"

type Order struct {
	UserID     int32     `json:"userID"`
	OrderID    int64     `json:"orderID"`
	Items      []Product `json:"items"`
	TotalPrice float32   `json:"totalPrice"`
	Date       time.Time `json:"date"`
}
