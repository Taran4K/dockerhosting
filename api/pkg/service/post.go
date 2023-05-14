package service

import (
	"api/models"
	"api/pkg/repository"
)

type PostService struct {
	dep  repository.Department
	repo repository.Post
}

func NewPostService(repo repository.Post, dep repository.Department) *PostService {
	return &PostService{repo: repo, dep: dep}
}

func (s *PostService) Create(post models.Post, iddep int, idorg int) (models.Post, error) {
	_, err := s.dep.GetById(iddep, idorg)
	if err != nil {
		return models.Post{}, err
	}

	return s.repo.Create(post, iddep)
}

func (s *PostService) GetAll(iddep int, idorg int) ([]models.Post, error) {
	_, err := s.dep.GetById(iddep, idorg)
	if err != nil {
		return nil, err
	}

	return s.repo.GetAll(iddep)
}

func (s *PostService) GetById(id int, iddep int, idorg int) (models.Post, error) {
	_, err := s.dep.GetById(iddep, idorg)
	if err != nil {
		return models.Post{}, err
	}

	return s.repo.GetById(id, iddep)
}

func (s *PostService) Delete(id int, iddep int, idorg int) error {
	_, err := s.dep.GetById(iddep, idorg)
	if err != nil {
		return err
	}

	return s.repo.Delete(id, iddep)
}

func (s *PostService) Update(id int, post models.Post, iddep int, idorg int) (models.Post, error) {
	_, err := s.dep.GetById(iddep, idorg)
	if err != nil {
		return models.Post{}, err
	}

	return s.repo.Update(id, post, iddep)
}
