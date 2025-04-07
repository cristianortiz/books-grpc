package model

// defines a contract that other structs can implement. However, it is not directly related to Book or DBBook structs
type Tabler interface {
	TableName() string
}

// for interaction with DB, same fields that model/books.go, the idea is to use this struct
// only with DB ops like and ORM, and separate bussines logic (book model) from DB specifics details
type DBBook struct {
	Isbn      int    `json:"isbn"`
	Name      string `json:"name"`
	Publisher int    `json:"publisher"`
}

// return the DB table name associated with this specific Book model
func (DBBook) TableName() string {
	return "books"
}
