package service

import (
	"api/models"
	"api/pkg/repository"
)

type Empl_postService struct {
	empl repository.Employee
	repo repository.Empl_post
}

func NewEmpl_postService(repo repository.Empl_post, empl repository.Employee) *Empl_postService {
	return &Empl_postService{repo: repo, empl: empl}
}

func (s *Empl_postService) Create(Empl_post models.Employee_Post, idempl int, iddep int) (models.Employee_Post, error) {
	_, err := s.empl.GetById(idempl, iddep)
	if err != nil {
		return models.Employee_Post{}, err
	}

	return s.repo.Create(Empl_post, idempl)
}

func (s *Empl_postService) GetAll(idempl int, iddep int) ([]models.Employee_Post, error) {
	_, err := s.empl.GetById(idempl, iddep)
	if err != nil {
		return nil, err
	}

	return s.repo.GetAll(idempl)
}

func (s *Empl_postService) GetById(id int, idempl int, iddep int) (models.Employee_Post, error) {
	_, err := s.empl.GetById(idempl, iddep)
	if err != nil {
		return models.Employee_Post{}, err
	}

	return s.repo.GetById(id, idempl)
}

func (s *Empl_postService) Delete(idempl int, iddep int) error {
	_, err := s.empl.GetById(idempl, iddep)
	if err != nil {
		return err
	}

	return s.repo.Delete(idempl)
}
