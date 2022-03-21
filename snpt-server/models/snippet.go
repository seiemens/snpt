/*
Created by Ramon
Date: 7.3.22
functions:
	Snippet struct
*/
package models

type Snippet struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Cookie  string `json:"cookie"`
}
