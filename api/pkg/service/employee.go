package service

import (
	"api/models"
	"api/pkg/repository"
)

type EmployeeService struct {
	dep  repository.Department
	repo repository.Employee
}

func NewEmployeeService(repo repository.Employee, dep repository.Department) *EmployeeService {
	return &EmployeeService{repo: repo, dep: dep}
}

func (s *EmployeeService) Create(Employee models.Employee, iddep int, idorg int) (models.Employee, error) {
	_, err := s.dep.GetById(iddep, idorg)
	if err != nil {
		return models.Employee{}, err
	}

	return s.repo.Create(Employee, iddep)
}

func (s *EmployeeService) GetAll(iddep int, idorg int) ([]models.Employee, error) {
	_, err := s.dep.GetById(iddep, idorg)
	if err != nil {
		return nil, err
	}

	return s.repo.GetAll(iddep)
}

func (s *EmployeeService) GetById(id int, iddep int, idorg int) (models.Employee, error) {
	_, err := s.dep.GetById(iddep, idorg)
	if err != nil {
		return models.Employee{}, err
	}

	return s.repo.GetById(id, iddep)
}

func (s *EmployeeService) Delete(id int, iddep int, idorg int) error {
	_, err := s.dep.GetById(iddep, idorg)
	if err != nil {
		return err
	}

	return s.repo.Delete(id, iddep)
}

func (s *EmployeeService) Update(id int, Employee models.Employee, iddep int, idorg int) (models.Employee, error) {
	_, err := s.dep.GetById(iddep, idorg)
	if err != nil {
		return models.Employee{}, err
	}

	return s.repo.Update(id, Employee, iddep)
}
