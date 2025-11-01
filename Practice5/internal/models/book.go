package models

type Book struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Genre string  `json:"genre"`
}

type BookFilter struct {
	Genre  string
	Sort   string
	Limit  int
	Offset int
}
