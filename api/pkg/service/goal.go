package service

import (
	"api/models"
	"api/pkg/repository"
)

type GoalService struct {
	dep  repository.Department
	repo repository.Goal
}

func NewGoalService(repo repository.Goal, dep repository.Department) *GoalService {
	return &GoalService{repo: repo, dep: dep}
}

func (s *GoalService) Create(Goal models.Goal, iddep int, idorg int) (models.Goal, error) {
	_, err := s.dep.GetById(iddep, idorg)
	if err != nil {
		return models.Goal{}, err
	}

	return s.repo.Create(Goal, iddep)
}

func (s *GoalService) GetAll(iddep int, idorg int) ([]models.Goal, error) {
	_, err := s.dep.GetById(iddep, idorg)
	if err != nil {
		return nil, err
	}

	return s.repo.GetAll(iddep)
}

func (s *GoalService) GetById(id int, iddep int, idorg int) (models.Goal, error) {
	_, err := s.dep.GetById(iddep, idorg)
	if err != nil {
		return models.Goal{}, err
	}

	return s.repo.GetById(id, iddep)
}

func (s *GoalService) Delete(id int, iddep int, idorg int) error {
	_, err := s.dep.GetById(iddep, idorg)
	if err != nil {
		return err
	}

	return s.repo.Delete(id, iddep)
}

func (s *GoalService) Update(id int, Goal models.Goal, iddep int, idorg int) (models.Goal, error) {
	_, err := s.dep.GetById(iddep, idorg)
	if err != nil {
		return models.Goal{}, err
	}

	return s.repo.Update(id, Goal, iddep)
}
