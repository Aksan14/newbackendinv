package service

import (
	"godesaapps/dto"
	"godesaapps/model"
)

type WebsiteContentService interface {
	GetContent() (*model.WebsiteContent, error)
	UpdateContent(req *dto.WebsiteContentRequest) error
}
