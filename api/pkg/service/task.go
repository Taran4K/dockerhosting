package service

import (
	"api/models"
	"api/pkg/repository"
)

type TaskService struct {
	empl repository.Employee
	repo repository.Task
}

func NewTaskService(repo repository.Task, empl repository.Employee) *TaskService {
	return &TaskService{repo: repo, empl: empl}
}

func (s *TaskService) Create(Task models.Task, idempl int, iddep int) (models.Task, error) {
	_, err := s.empl.GetById(idempl, iddep)
	if err != nil {
		return models.Task{}, err
	}

	return s.repo.Create(Task, idempl)
}

func (s *TaskService) GetAll(idempl int, iddep int) ([]models.Task, error) {
	_, err := s.empl.GetById(idempl, iddep)
	if err != nil {
		return nil, err
	}

	return s.repo.GetAll(idempl)
}

func (s *TaskService) GetById(id int, idempl int, iddep int) (models.Task, error) {
	_, err := s.empl.GetById(idempl, iddep)
	if err != nil {
		return models.Task{}, err
	}

	return s.repo.GetById(id, idempl)
}

func (s *TaskService) Delete(id int, idempl int, iddep int) error {
	_, err := s.empl.GetById(idempl, iddep)
	if err != nil {
		return err
	}

	return s.repo.Delete(id, idempl)
}

func (s *TaskService) Update(id int, Task models.Task, idempl int, iddep int) (models.Task, error) {
	_, err := s.empl.GetById(idempl, iddep)
	if err != nil {
		return models.Task{}, err
	}

	return s.repo.Update(id, Task, idempl)
}
