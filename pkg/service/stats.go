package service

import (
	"net/http"
	"time"

	"github.com/c9s/goprocinfo/linux"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *Service) StatsWSHandler(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	go func() {
		stat, err := linux.ReadCPUInfo("/proc/cpuinfo")
		if err != nil {
			ws.WriteJSON(map[string]interface{}{
				"error": err,
			})
			return
		}
		ws.WriteJSON(stat)
		for {
			memory, err := linux.ReadMemInfo("/proc/meminfo")
			if err != nil {
				return
			}
			ws.WriteJSON(memory)
			time.Sleep(time.Second * 1)
		}
	}()

	return nil
}
