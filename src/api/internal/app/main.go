package app

import (
	"context"
	"github.com/kardianos/service"
	"github.com/mnogokotin/golang-windows-service/internal/service/task"
	"os"
	"time"
)

type Config struct {
	updateInteval int
}

func newConfig() *Config {
	return &Config{updateInteval: 1}
}

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	c := newConfig()
	ctx, cancel := context.WithCancel(context.Background())
	go task1(ctx, time.Duration(c.updateInteval))
	time.Sleep(10 * time.Minute)
	cancel()
}

func (p *program) Stop(s service.Service) error {
	<-time.After(1 * time.Second)
	return nil
}

func Main() {
	svcConfig := &service.Config{
		Name:        "GolangService",
		DisplayName: "Golang Service",
		Description: "Golang Service (Windows)",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		panic(err)
	}
	if len(os.Args) > 1 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			panic(err)
		}
		return
	}

	logger, err = s.Logger(nil)
	if err != nil {
		panic(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}

func task1(ctx context.Context, updateInterval time.Duration) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			task.ReadFromDbAndWriteToFile()
		}
		time.Sleep(updateInterval * time.Second)
	}
}
