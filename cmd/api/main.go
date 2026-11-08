package main

import (
	"github.com/berchmanchua/golang-viral-api/internal/aws"
	"github.com/berchmanchua/golang-viral-api/internal/client"
	"github.com/berchmanchua/golang-viral-api/internal/config"
	"github.com/berchmanchua/golang-viral-api/internal/db"
	"github.com/berchmanchua/golang-viral-api/internal/server"
	"github.com/berchmanchua/golang-viral-api/internal/service"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.LoadConfig,
			db.NewPostgresDB,
			db.NewRedisClient,
			client.NewRetryableHTTPClient,
			aws.NewDynamoDBClient,
			aws.NewKinesisClient,
			service.NewViralService,
			server.NewViralHandler,
			server.NewRouter,
		),
		fx.Invoke(server.StartHTTPServer),
	)

	app.Run()
}
