package service

import (
	"api/models"
	"api/pkg/repository"
)

type DepartmentService struct {
	org  repository.Organization
	repo repository.Department
	empl repository.Employee
}

func NewDepartmentService(repo repository.Department, org repository.Organization, empl repository.Employee) *DepartmentService {
	return &DepartmentService{repo: repo, org: org, empl: empl}
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

func (s *DepartmentService) GetRucovoditel(id int) (models.Employee, error) {
	return s.repo.GetRucovoditel(id)
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

func (s *DepartmentService) UpdateRucovoditel(rucovoditel models.Employee, oldruc, newruc int) (string, error) {
	print(oldruc)
	print(rucovoditel.ID_Employee)
	print(newruc)
	if oldruc != 0 && rucovoditel.ID_Employee != 0 {
		if rucovoditel.ID_Employee == oldruc {
			if oldruc != newruc {
				oldorg, err := s.empl.GetDep(oldruc)
				if err != nil {
					return "Неверный ключ сотрудника", err
				}

				dirorg, err := s.empl.GetDep(rucovoditel.ID_Employee)
				if err != nil {
					return "Неверный ключ сотрудника", err
				}

				neworg, err2 := s.empl.GetDep(newruc)
				if err2 != nil {
					return "Неверный ключ сотрудника", err2
				}

				if oldorg.ID_Department == neworg.ID_Department && oldorg.ID_Department == dirorg.ID_Department {
					return s.repo.UpdateRucovoditel(oldruc, newruc)
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
		return s.repo.CreateRucovoditel(newruc)
	}
}
