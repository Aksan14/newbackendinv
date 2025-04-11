package dto

type WebsiteContentRequest struct {
	Logo        string `json:"logo"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
}
