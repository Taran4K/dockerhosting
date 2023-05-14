package service

import (
	"api/models"
	"api/pkg/repository"
)

type UserService struct {
	empl repository.Employee
	repo repository.User
}

func NewUserService(repo repository.User, empl repository.Employee) *UserService {
	return &UserService{repo: repo, empl: empl}
}

func (s *UserService) Create(User models.User, idempl int, iddep int) (models.User, error) {
	_, err := s.empl.GetById(idempl, iddep)
	if err != nil {
		return models.User{}, err
	}

	return s.repo.Create(User, idempl)
}

func (s *UserService) GetAll(idempl int, iddep int) ([]models.User, error) {
	_, err := s.empl.GetById(idempl, iddep)
	if err != nil {
		return nil, err
	}

	return s.repo.GetAll(idempl)
}

func (s *UserService) GetById(id int, idempl int, iddep int) (models.User, error) {
	_, err := s.empl.GetById(idempl, iddep)
	if err != nil {
		return models.User{}, err
	}

	return s.repo.GetById(id, idempl)
}

func (s *UserService) Delete(id int, idempl int, iddep int) error {
	_, err := s.empl.GetById(idempl, iddep)
	if err != nil {
		return err
	}

	return s.repo.Delete(id, idempl)
}

func (s *UserService) Update(id int, User models.User, idempl int, iddep int) (models.User, error) {
	_, err := s.empl.GetById(idempl, iddep)
	if err != nil {
		return models.User{}, err
	}

	return s.repo.Update(id, User, idempl)
}
