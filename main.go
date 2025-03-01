package main

import (
	"context"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/islu/HW0720/internal/router"
	"github.com/islu/HW0720/internal/usecase"
)

func initAppConfig() *usecase.ApplicationParams {
	env := mustGetenv("NODE_ENV")

	// Database
	dbHost := mustGetenv("DB_HOST")
	dbPort := mustGetenv("DB_PORT")
	dbName := mustGetenv("DB_NAME")
	dbUser := mustGetenv("DB_USER")
	dbPwd := mustGetenv("DB_PASS")
	dbSchemaName := "public"

	// Alchemy
	alchemyAPIKey := mustGetenv("ALCHEMY_API_KEY")

	return &usecase.ApplicationParams{
		Environment:   env,
		DBHost:        dbHost,
		DBPort:        dbPort,
		DBName:        dbName,
		DBUser:        dbUser,
		DBPassword:    dbPwd,
		DBSchemaName:  dbSchemaName,
		AlchemyAPIKey: alchemyAPIKey,
	}
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Fatal Error in connect_connector.go: %s environment variable not set.\n", k)
	}
	return v
}

func main() {
	// Init configuration
	config := initAppConfig()

	app, err := usecase.NewApplication(context.Background(), config)
	if err != nil {
		log.Fatalln(err)
	}

	setupGinRoute(app)
}

// Setup Gin route and run Gin server
//
//	@title			HW0720
//	@version		1.0
//	@description	HW0720
//	@host			localhost:8080
//	@basePath		/api/v1
func setupGinRoute(app *usecase.Application) {
	r := router.RegisterHandlers(app)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
