package main

type ShipmentInput struct {
	Sender   string  `json:"sender" binding:"required"`   // ISO-3166-1 alpha-2
	Receiver string  `json:"receiver" binding:"required"` // ISO-3166-1 alpha-2
	Weight   float64 `json:"weight" binding:"required"`   // in Kg
}

type ShipmentResponse struct {
	Sender   string  `json:"sender"`   // ISO-3166-1 alpha-2
	Receiver string  `json:"receiver"` // ISO-3166-1 alpha-2
	Weight   float64 `json:"weight"`   // in Kg
	Price    float64 `json:"price"`
}
