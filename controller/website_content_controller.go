package controller

import (
	"godesaapps/dto"
	"godesaapps/model"
)

type ContentGo interface {
	GetContent() (*model.WebsiteContent, error)
	UpdateContent(req *dto.WebsiteContentRequest) error
}