package handlers

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/heshoots/smbfapi/gen/restapi/operations"
	"github.com/heshoots/smbfapi/pkg/challonge"
	"github.com/heshoots/smbfapi/pkg/facebook"
	"time"
)

func GetEventHandler(params operations.GetEventParams) middleware.Responder {
	response, err := facebook.GetEvent(params.ID, params.APIKey)
	if err != nil {
		fmt.Println(err)
		return operations.NewGetEventBadRequest()
	}

	start, err := time.Parse("2006-01-02T15:04:05-0700", response.StartTime)
	if err == nil {
		tournaments, err := challonge.GetTournamentsForDate(start)
		if err == nil {
			response.Tournaments = tournaments
		}
	}
	return operations.NewGetEventOK().WithPayload(response)
}

func SearchEventsHandler(params operations.SearchEventsParams) middleware.Responder {
	response, err := facebook.GetEvents(params.APIKey)
	if err != nil {
		fmt.Println(err)
		return operations.NewSearchEventsBadRequest()
	}
	return operations.NewSearchEventsOK().WithPayload(response)
}
