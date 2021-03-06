// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetTournamentHandlerFunc turns a function with the right signature into a get tournament handler
type GetTournamentHandlerFunc func(GetTournamentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTournamentHandlerFunc) Handle(params GetTournamentParams) middleware.Responder {
	return fn(params)
}

// GetTournamentHandler interface for that can handle valid get tournament params
type GetTournamentHandler interface {
	Handle(GetTournamentParams) middleware.Responder
}

// NewGetTournament creates a new http.Handler for the get tournament operation
func NewGetTournament(ctx *middleware.Context, handler GetTournamentHandler) *GetTournament {
	return &GetTournament{Context: ctx, Handler: handler}
}

/*GetTournament swagger:route GET /tournaments/{id} getTournament

get specific tournament

Access values for a tournament by id


*/
type GetTournament struct {
	Context *middleware.Context
	Handler GetTournamentHandler
}

func (o *GetTournament) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTournamentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
