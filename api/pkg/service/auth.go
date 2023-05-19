package service

import (
	"api/models"
	"api/pkg/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	gomail "gopkg.in/mail.v2"
	"time"
)

const (
	salt         = "dfghjkweghkjhfgaHJ098BH"
	signinKey    = "kjew7hH3jhkf8jdkl32KDwWDJdi32d"
	tokenTTL     = 12 * time.Hour
	mailsender   = "organizationcontrollapp@gmail.com"
	mailpassword = "dqcatyyhwzknpvjf"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) Authorize(login, password string, idempl int) (string, error) {
	user, err := s.repo.GetUser(login, generatePasswordHash(password))

	if err != nil {
		return "", err
	}
	println(user.Employee_ID)
	println(idempl)
	if user.Employee_ID == idempl {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(tokenTTL).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			user.Id_user,
		})
		return token.SignedString([]byte(signinKey))
	} else {
		return "Почта не подходит к данному пользователю", err
	}
}

func (s *AuthService) GetAll(id int) (models.UserAllData, error) {
	userall, err := s.repo.GetAll(id)

	if err != nil {
		return models.UserAllData{}, err
	}

	return userall, err
}

func (s *AuthService) GetEmployee(email string) (int, error) {
	user, err := s.repo.GetEmployee(email)

	if err != nil {
		return 0, err
	}

	return user, nil
}

func (s *AuthService) EmailCheck(email, code string) (string, error) {
	mail := gomail.NewMessage()
	mail.SetHeader("From", mailsender)
	mail.SetHeader("To", email)
	mail.SetHeader("Subject", "Код для подтверждения почты")
	mail.SetBody("text/plain", code)

	send := gomail.NewDialer("smtp.gmail.com", 587, mailsender, mailpassword)

	if err := send.DialAndSend(mail); err != nil {
		return "Ошибка", err
	}

	return "Успешная отправка", nil
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signinKey), nil
	})

	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
