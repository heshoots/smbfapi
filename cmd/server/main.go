package main

import (
	"log"

	"github.com/alexflint/go-arg"
	"github.com/go-openapi/loads"
	"github.com/heshoots/smbfapi/gen/restapi"
	"github.com/heshoots/smbfapi/gen/restapi/operations"
	"github.com/heshoots/smbfapi/pkg/handlers"
	"github.com/spf13/viper"
)

type cliArgs struct {
	Port int `arg:"-p,help:port to listen to"`
}

var (
	args = &cliArgs{
		Port: 8080,
	}
)

func main() {
	arg.MustParse(args)

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewSmbfAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = args.Port

	api.GetEventHandler = operations.GetEventHandlerFunc(handlers.GetEventHandler)
	api.SearchEventsHandler = operations.SearchEventsHandlerFunc(handlers.SearchEventsHandler)

	api.GetTournamentHandler = operations.GetTournamentHandlerFunc(handlers.GetTournamentHandler)
	api.SearchTournamentsHandler = operations.SearchTournamentsHandlerFunc(handlers.SearchTournamentsHandler)

	viper.AutomaticEnv()
	viper.SetEnvPrefix("smbfapi")
	viper.BindEnv("challongekey")

	// Start listening using having the handlers and port
	// already set up.
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
