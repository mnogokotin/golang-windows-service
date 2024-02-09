package app

import (
	"context"
	"github.com/kardianos/service"
	"github.com/mnogokotin/golang-packages/database/postgres"
	"github.com/mnogokotin/golang-windows-service/internal/config"
	"github.com/mnogokotin/golang-windows-service/internal/service/csv"
	"github.com/mnogokotin/golang-windows-service/internal/service/database"
	"github.com/mnogokotin/golang-windows-service/internal/service/file"
	"log"
	"os"
	"time"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	pg, err := postgres.New(cfg.PG.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Close()

	ctx, cancel := context.WithCancel(context.Background())
	go task(ctx, cfg.Service.UpdateInterval, cfg.Service.OutputFilePath, cfg.Service.CsvSeparator, pg)

	time.Sleep(cfg.Service.CancelInterval)
	cancel()
}

func (p *program) Stop(s service.Service) error {
	<-time.After(1 * time.Second)
	return nil
}

func Run() {
	svcConfig := &service.Config{
		Name:        "GolangService",
		DisplayName: "Golang service",
	}

	p := &program{}
	s, err := service.New(p, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) > 1 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}

func task(ctx context.Context, updateInterval time.Duration, outputFilePath string, csvSeparator string, pg *postgres.Postgres) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			func() {
				f, err := file.OpenOrCreateFileOnRead(outputFilePath)
				defer f.Close()
				if err != nil {
					log.Fatal(err)
				}

				lastLineId, err := csv.GetLastLineId(f, csvSeparator)
				if err != nil {
					lastLineId = "0"
				}

				eventModels := database.GetEventModelsWithGreaterId(pg, lastLineId)
				if len(eventModels) > 0 {
					f, err := file.OpenFileOnWriteAtTheEnd(outputFilePath)
					defer f.Close()
					if err != nil {
						log.Fatal(err)
					}

					err = file.WriteModelsToFile(f, csvSeparator, eventModels)
					if err != nil {
						log.Fatal(err)
					}
				}
			}()
		}
		time.Sleep(updateInterval)
	}
}
