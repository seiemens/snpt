package models

type Snippet struct {
	ID       float64 `json:"id"`
	Title    string  `json:"title"`
	Content  string  `json:"content"`
	AuthorID float64 `json:"authorID"`
}
