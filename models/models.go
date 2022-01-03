package models

type HouseDetails struct {
	Name     string  `json:"name"`
	Address  string  `json:"address"`
	Locality string  `json:"locality"`
	Country  string  `json:"country"`
	PinCode  string  `json:"pinCode"`
	Amount   float32 `db:"amount"`
}
