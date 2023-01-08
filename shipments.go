package main

type ShipmentInput struct {
	Sender   string  // ISO-3166-1 alpha-2
	Receiver string  // ISO-3166-1 alpha-2
	Weight   float32 // in Kg
}

type ShipmentResponse struct {
	Sender   string  `json:"sender"`   // ISO-3166-1 alpha-2
	Receiver string  `json:"receiver"` // ISO-3166-1 alpha-2
	Weight   float32 `json:"weight"`   // in Kg
	Price    float32 `json:"price"`
}
