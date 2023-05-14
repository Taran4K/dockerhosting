package service

import (
	"api/models"
	"api/pkg/repository"
	gomail "gopkg.in/mail.v2"
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

func (s *UserService) UpdateProfile(id int, User models.UserAllData) (models.User, error) {
	return s.repo.UpdateProfile(id, User)
}

func (s *UserService) UpdateEmployee(empl models.Employee) (models.Employee, error) {
	return s.repo.UpdateEmployee(empl)
}

func (s *UserService) UpdatePassword(id int, oldpass string, newpass string, userall models.UserAllData) (string, error) {
	if oldpass != newpass {
		if generatePasswordHash(oldpass) == userall.User.Password {
			return s.repo.UpdatePassword(id, generatePasswordHash(newpass))
		} else {
			return "Старый пароль неверный", nil
		}
	} else {
		return "Старый и новый пароль совапдают", nil
	}
}

func (s *UserService) CreateTask(Task models.Task) (models.Task, error) {

	return s.repo.CreateTask(Task)
}

func (s *UserService) GetAllTask(idempl int) ([]models.Task, error) {

	return s.repo.GetAllTask(idempl)
}

func (s *UserService) GetTaskById(id int) (models.Task, error) {

	return s.repo.GetTaskById(id)
}

func (s *UserService) DeleteTask(id int) error {

	return s.repo.DeleteTask(id)
}

func (s *UserService) DeleteLogTask(id int) error {

	return s.repo.DeleteLogTask(id)
}

func (s *UserService) UpdateTask(id int, Task models.Task) (models.Task, error) {

	return s.repo.UpdateTask(id, Task)
}

func (s *UserService) CreateGoal(Goal models.Goal) (models.Goal, error) {

	return s.repo.CreateGoal(Goal)
}

func (s *UserService) GetAllGoal(idempl int) ([]models.Goal, error) {

	return s.repo.GetAllGoal(idempl)
}

func (s *UserService) GetGoalById(id int) (models.Goal, error) {

	return s.repo.GetGoalById(id)
}

func (s *UserService) DeleteGoal(id int) error {

	return s.repo.DeleteGoal(id)
}

func (s *UserService) DeleteLogGoal(id int) error {

	return s.repo.DeleteLogGoal(id)
}

func (s *UserService) UpdateGoal(id int, goal models.Goal) (models.Goal, error) {

	return s.repo.UpdateGoal(id, goal)
}

func (s *UserService) CreateStrat(strat models.Strategy) (models.Strategy, error) {

	return s.repo.CreateStrat(strat)
}

func (s *UserService) GetAllStrat(idempl int) ([]models.Strategy, error) {

	return s.repo.GetAllStrat(idempl)
}

func (s *UserService) GetStratById(id int) (models.Strategy, error) {

	return s.repo.GetStratById(id)
}

func (s *UserService) DeleteStrat(id int) error {

	return s.repo.DeleteStrat(id)
}

func (s *UserService) DeleteLogStrat(id int) error {

	return s.repo.DeleteLogStrat(id)
}

func (s *UserService) UpdateStrat(id int, strat models.Strategy) (models.Strategy, error) {

	return s.repo.UpdateStrat(id, strat)
}

func (s *UserService) SendMail(email, title, message string) (string, error) {
	mail := gomail.NewMessage()
	mail.SetHeader("From", mailsender)
	mail.SetHeader("To", email)
	mail.SetHeader("Subject", title)
	mail.SetBody("text/plain", message)

	send := gomail.NewDialer("smtp.gmail.com", 587, mailsender, mailpassword)

	if err := send.DialAndSend(mail); err != nil {
		return "Ошибка", err
	}

	return "Успешная отправка", nil
}
