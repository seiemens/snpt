package main

type user struct {
	ID       float64 `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Email    float64 `json:"email"`
}

type snippet struct {
	ID       float64 `json:"id"`
	Title    string  `json:"title"`
	Content  string  `json:"content"`
	AuthorID float64 `json:"authorID"`
}
