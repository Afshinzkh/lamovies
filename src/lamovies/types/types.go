package types

import "time"

// Movie ...
type Movie struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	DateAdded time.Time `json:"date_added"`
}
