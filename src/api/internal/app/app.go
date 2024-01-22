package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	//"os"
	//"strings"
)

func init() {
	if err := godotenv.Load("/app/.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Run() {
	//trustedProxies := strings.Split(os.Getenv("TRUSTED_PROXIES"), " ")
	fmt.Print("run")
}
