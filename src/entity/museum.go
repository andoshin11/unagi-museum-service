package entity

import "time"

// Museum type definition
type Museum struct {
	Identifier string    `firestore:"identifier" json:"identifier"`
	CreatedAt  time.Time `firestore:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time `firestore:"updatedAt" json:"updatedAt"`
	Name       string    `firestore:"name" json:"name"`
	Address    string    `firestore:"address" json:"address"`
	Img        string    `firestore:"img" json:"img"`
	Entry      string    `firestore:"entry" json:"entry"`
	SiteURL    string    `firestore:"siteUrl" json:"siteUrl"`
	Lat        float64   `firestore:"lat" json:"lat"`
	Lng        float64   `firestore:"lng" json:"lng"`
}
