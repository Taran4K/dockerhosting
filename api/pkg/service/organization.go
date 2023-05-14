package service

import (
	"api/models"
	"api/pkg/repository"
)

type OrganizationService struct {
	repo repository.Organization
}

func NewOrganizationService(repo repository.Organization) *OrganizationService {
	return &OrganizationService{repo: repo}
}

func (s *OrganizationService) Create(organization models.Organization) (models.Organization, error) {
	return s.repo.Create(organization)
}

func (s *OrganizationService) GetAll() ([]models.Organization, error) {
	return s.repo.GetAll()
}

func (s *OrganizationService) GetById(id int) (models.Organization, error) {
	return s.repo.GetById(id)
}

func (s *OrganizationService) GetByKey(key string) (models.Organization, error) {
	return s.repo.GetByKey(key)
}

func (s *OrganizationService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *OrganizationService) Update(id int, organization models.Organization) (models.Organization, error) {
	return s.repo.Update(id, organization)
}
