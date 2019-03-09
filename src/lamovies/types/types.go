package types

// Movie ...
type Movie struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsRented bool   `json:"is_rented"`
}
