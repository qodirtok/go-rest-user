package cmd

import (
	"go-rest-user/config"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func goDotEnvVariable() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	// return os.Getenv()
}

func Cmd() {

	goDotEnvVariable()

	config.Connection()

	rest := gin.Default()
	config.Http(rest)

	rest.Run(":4748")
}
