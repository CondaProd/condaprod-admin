package service

import (
	"condaprod-admin/pkg/models"
	"os/exec"

	"github.com/labstack/echo"
)

func (s *Service) Auth(c echo.Context) error {
	user := &models.User{}
	c.Bind(user)
	cmd := exec.Command("su", "-", user.Login)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	if _, err := stdin.Write([]byte(user.Password)); err != nil {
		return err
	}

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
}
