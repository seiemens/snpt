package models

type Snippet struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Cookie  string `json:"cookie"`
}
