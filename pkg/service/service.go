package service

import "github.com/labstack/echo"

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) Run() {
	e := echo.New()
	e.POST("/auth", s.Auth)
	e.Start("0.0.0.0:8080")
}
