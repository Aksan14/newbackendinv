package service

import (
	"godesaapps/dto"
	"godesaapps/model"
	"godesaapps/repository"
)

type websiteContentService struct {
	repo repository.WebsiteContentRepository
}

func NewWebsiteContentService(repo repository.WebsiteContentRepository) WebsiteContentService {
	return &websiteContentService{repo}
}

func (s *websiteContentService) GetContent() (*model.WebsiteContent, error) {
	return s.repo.GetContent()
}

func (s *websiteContentService) UpdateContent(req *dto.WebsiteContentRequest) error {
	content, err := s.repo.GetContent()
	if err != nil {
		return err
	}

	content.Logo = req.Logo
	content.Title = req.Title
	content.Description = req.Description
	content.Address = req.Address
	content.Email = req.Email
	content.Phone = req.Phone

	return s.repo.UpdateContent(content)
}
