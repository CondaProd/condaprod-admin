package service

import (
	"condaprod-admin/pkg/models"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) Run() {
	e := echo.New()
	e.POST("/auth", s.Auth)

	r := e.Group("/api", middleware.JWTWithConfig(middleware.JWTConfig{
		Claims: &models.Claims{},
		SuccessHandler: func(c echo.Context) {
			user := c.Get("user").(*models.Claims).Username
			c.Set("username", user)
		},
		SigningKey: []byte("GGWPSosatbElena20081342CTSpAwNHackersDownsAHAHHAHAhahahMidShadowFiend"),
	}))

	r.GET("/stats/ws", s.StatsWSHandler) // endpoint for get linux machine stats in real time

	e.Start("0.0.0.0:8080")
}
