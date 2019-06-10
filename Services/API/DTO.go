package apiService

import "time"

type Badge struct {
	ID        string    `json:"ID"`
	LastName  string    `json:"nom"`
	FirstName string    `json:"prenom"`
	CreatedAt time.Time `json:"ajout"`
}
