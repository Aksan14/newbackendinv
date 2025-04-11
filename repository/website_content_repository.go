package repository

import (
	"godesaapps/model"
)

type WebsiteContentRepository interface {
	GetContent() (*model.WebsiteContent, error)
	UpdateContent(data *model.WebsiteContent) error
}
