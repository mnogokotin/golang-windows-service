package app

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/mnogokotin/golang-windows-service/internal/service"
	"os"
	"strconv"
	"time"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	updateInterval, _ := strconv.Atoi(os.Getenv("UPDATE_INTEVAL"))
	go task(ctx, time.Duration(updateInterval))
	time.Sleep(10 * time.Minute)
	cancel()
}

func task(ctx context.Context, updateInterval time.Duration) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			service.ReadFromDbAndWriteToFile()
		}
		time.Sleep(updateInterval * time.Second)
	}
}
