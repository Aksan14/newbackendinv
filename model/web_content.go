package model

type WebsiteContent struct {
	ID          int64  `json:"id"`
	Logo        string `json:"logo"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
}