package repository

import (
	"database/sql"
	"godesaapps/model"
)

type websiteContentRepository struct {
	db *sql.DB
}

func NewWebsiteContentRepository(db *sql.DB) WebsiteContentRepository {
	return &websiteContentRepository{db}
}

func (r *websiteContentRepository) GetContent() (*model.WebsiteContent, error) {
	row := r.db.QueryRow(`SELECT id, logo, title, description, address, email, phone FROM website_content LIMIT 1`)
	var content model.WebsiteContent
	err := row.Scan(&content.ID, &content.Logo, &content.Title, &content.Description, &content.Address, &content.Email, &content.Phone)
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func (r *websiteContentRepository) UpdateContent(data *model.WebsiteContent) error {
	_, err := r.db.Exec(`
		UPDATE website_content SET 
			logo = ?, 
			title = ?, 
			description = ?, 
			address = ?, 
			email = ?, 
			phone = ?
		WHERE id = ?`,
		data.Logo,
		data.Title,
		data.Description,
		data.Address,
		data.Email,
		data.Phone,
		data.ID,
	)
	return err
}
