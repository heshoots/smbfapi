package handlers

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/heshoots/smbfapi/gen/restapi/operations"
	"github.com/heshoots/smbfapi/pkg/challonge"
)

func GetTournamentHandler(params operations.GetTournamentParams) middleware.Responder {

	tournament, err := challonge.GetTournament(params.ID)
	if err != nil {
		fmt.Println(err)
		return operations.NewGetTournamentBadRequest()
	}
	return operations.NewGetTournamentOK().WithPayload(tournament)
}

func SearchTournamentsHandler(params operations.SearchTournamentsParams) middleware.Responder {
	tournaments, err := challonge.GetTournaments()
	if err != nil {
		fmt.Println(err)
		operations.NewSearchTournamentsBadRequest()
	}
	return operations.NewSearchTournamentsOK().WithPayload(tournaments)
}
