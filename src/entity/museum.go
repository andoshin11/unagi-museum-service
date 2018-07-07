package entity

// Museum type definition
type Museum struct {
	ID        string  `json:"id"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name"`
	Img       string  `json:"img"`
}
