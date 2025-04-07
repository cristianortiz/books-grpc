package model

// Book model used acrross different parts o tiers of the books app
// but only for bussines logic
type Book struct {
	Isbn      int    `json:"isbn"`
	Name      string `json:"name"`
	Publisher int    `json:"publisher"`
}
