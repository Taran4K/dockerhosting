package service

import (
	"api/models"
	"api/pkg/repository"
)

type OperationService struct {
	org  repository.Organization
	repo repository.Finances_Operations
}

func NewOperationService(repo repository.Finances_Operations, org repository.Organization) *OperationService {
	return &OperationService{repo: repo, org: org}
}

func (s *OperationService) Create(operation models.Finances_Operations, idorg int) (models.Finances_Operations, error) {
	_, err := s.org.GetById(idorg)
	if err != nil {
		return models.Finances_Operations{}, err
	}

	return s.repo.Create(operation, idorg)
}

func (s *OperationService) GetAll(idorg int) ([]models.Finances_Operations, error) {
	_, err := s.org.GetById(idorg)
	if err != nil {
		return nil, err
	}

	return s.repo.GetAll(idorg)
}

func (s *OperationService) GetById(id int, idorg int) (models.Finances_Operations, error) {
	_, err := s.org.GetById(idorg)
	if err != nil {
		return models.Finances_Operations{}, err
	}

	return s.repo.GetById(id, idorg)
}

func (s *OperationService) Delete(id int, idorg int) error {
	_, err := s.org.GetById(idorg)
	if err != nil {
		return err
	}

	return s.repo.Delete(id, idorg)
}

func (s *OperationService) Update(id int, operation models.Finances_Operations, idorg int) (models.Finances_Operations, error) {
	_, err := s.org.GetById(idorg)
	if err != nil {
		return models.Finances_Operations{}, err
	}

	return s.repo.Update(id, operation, idorg)
}
