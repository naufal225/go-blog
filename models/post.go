package models

type Post struct {
	ID int `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Content string `json:"content"`

	UserID int `json:"user_id"`
}