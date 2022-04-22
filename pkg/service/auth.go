package service

import (
	"condaprod-admin/pkg/models"
	"os/exec"

	"github.com/labstack/echo"
)

func (s *Service) Auth(c echo.Context) error {
	user := &models.User{}
	c.Bind(user)
	cmd := exec.Command("su", "-", user.Login, "-c", "whoami")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	if _, err := stdin.Write([]byte(user.Password + "\n")); err != nil {
		return err
	}

	stdout, err := cmd.Output()
	if err != nil {
		return err
	}

	if string(stdout[:len(stdout)-1]) != user.Login {
		return c.String(400, string(stdout))
	}

	return c.JSON(200, echo.Map{}) // todo
}
