package app

import (
	"condaprod-admin/pkg/service"
)

func Run() {
	s := service.New()
	s.Run()
}
