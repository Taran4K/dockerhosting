package service

import (
	"api/models"
	"api/pkg/repository"
)

type StrategyService struct {
	org  repository.Organization
	repo repository.Strategy
}

func NewStrategyService(repo repository.Strategy, org repository.Organization) *StrategyService {
	return &StrategyService{repo: repo, org: org}
}

func (s *StrategyService) Create(strategy models.Strategy, idorg int) (models.Strategy, error) {
	_, err := s.org.GetById(idorg)
	if err != nil {
		return models.Strategy{}, err
	}

	return s.repo.Create(strategy, idorg)
}

func (s *StrategyService) GetAll(idorg int) ([]models.Strategy, error) {
	_, err := s.org.GetById(idorg)
	if err != nil {
		return nil, err
	}

	return s.repo.GetAll(idorg)
}

func (s *StrategyService) GetById(id int, idorg int) (models.Strategy, error) {
	_, err := s.org.GetById(idorg)
	if err != nil {
		return models.Strategy{}, err
	}

	return s.repo.GetById(id, idorg)
}

func (s *StrategyService) Delete(id int, idorg int) error {
	_, err := s.org.GetById(idorg)
	if err != nil {
		return err
	}

	return s.repo.Delete(id, idorg)
}

func (s *StrategyService) Update(id int, strategy models.Strategy, idorg int) (models.Strategy, error) {
	_, err := s.org.GetById(idorg)
	if err != nil {
		return models.Strategy{}, err
	}

	return s.repo.Update(id, strategy, idorg)
}
