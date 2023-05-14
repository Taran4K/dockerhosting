package service

import (
	"api/models"
	"api/pkg/repository"
)

type OrganizationService struct {
	repo repository.Organization
	empl repository.Employee
}

func NewOrganizationService(repo repository.Organization, empl repository.Employee) *OrganizationService {
	return &OrganizationService{repo: repo, empl: empl}
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

func (s *OrganizationService) GetDirector(id int) (models.Employee, error) {
	empls, _ := s.empl.GetOrganizationAll(id)

	return s.repo.GetDirector(empls)
}

func (s *OrganizationService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *OrganizationService) Update(id int, organization models.Organization) (models.Organization, error) {
	return s.repo.Update(id, organization)
}

func (s *OrganizationService) UpdateDirector(director models.Employee, olddir, newdir int) (string, error) {
	if olddir != 0 && director.ID_Employee != 0 {
		if director.ID_Employee == olddir {
			if olddir != newdir {
				oldorg, err := s.empl.GetOrg(olddir)
				if err != nil {
					return "Неверный ключ сотрудника", err
				}

				dirorg, err := s.empl.GetOrg(director.ID_Employee)
				if err != nil {
					return "Неверный ключ сотрудника", err
				}

				neworg, err2 := s.empl.GetOrg(newdir)
				if err2 != nil {
					return "Неверный ключ сотрудника", err2
				}

				if oldorg.Id_Organization == neworg.Id_Organization && oldorg.Id_Organization == dirorg.Id_Organization {
					return s.repo.UpdateDirector(olddir, newdir)
				} else {
					return "Неверный ключ сотрудника", err2
				}
			} else {
				return "Ключи нового и старого директора совпадают", nil
			}
		} else {
			return "Неверный ключ директора", nil
		}
	} else {
		return s.repo.CreateDirector(newdir)
	}
}
