package service

import (
	"api/models"
	"api/pkg/repository"
)

type DepartmentService struct {
	org  repository.Organization
	repo repository.Department
}

func NewDepartmentService(repo repository.Department, org repository.Organization) *DepartmentService {
	return &DepartmentService{repo: repo, org: org}
}

func (s *DepartmentService) Create(department models.Department, idorg int) (models.Department, error) {
	_, err := s.org.GetById(idorg)
	if err != nil {
		return models.Department{}, err
	}

	return s.repo.Create(department, idorg)
}

func (s *DepartmentService) GetAll(idorg int) ([]models.Department, error) {
	_, err := s.org.GetById(idorg)
	if err != nil {
		return nil, err
	}

	return s.repo.GetAll(idorg)
}

func (s *DepartmentService) GetById(id int, idorg int) (models.Department, error) {
	_, err := s.org.GetById(idorg)
	if err != nil {
		return models.Department{}, err
	}

	return s.repo.GetById(id, idorg)
}

func (s *DepartmentService) Delete(id int, idorg int) error {
	_, err := s.org.GetById(idorg)
	if err != nil {
		return err
	}

	return s.repo.Delete(id, idorg)
}

func (s *DepartmentService) Update(id int, department models.Department, idorg int) (models.Department, error) {
	_, err := s.org.GetById(idorg)
	if err != nil {
		return models.Department{}, err
	}

	return s.repo.Update(id, department, idorg)
}
