// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// SearchTournamentsHandlerFunc turns a function with the right signature into a search tournaments handler
type SearchTournamentsHandlerFunc func(SearchTournamentsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SearchTournamentsHandlerFunc) Handle(params SearchTournamentsParams) middleware.Responder {
	return fn(params)
}

// SearchTournamentsHandler interface for that can handle valid search tournaments params
type SearchTournamentsHandler interface {
	Handle(SearchTournamentsParams) middleware.Responder
}

// NewSearchTournaments creates a new http.Handler for the search tournaments operation
func NewSearchTournaments(ctx *middleware.Context, handler SearchTournamentsHandler) *SearchTournaments {
	return &SearchTournaments{Context: ctx, Handler: handler}
}

/*SearchTournaments swagger:route GET /tournaments searchTournaments

get multiple tournament

get tournaments from api


*/
type SearchTournaments struct {
	Context *middleware.Context
	Handler SearchTournamentsHandler
}

func (o *SearchTournaments) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSearchTournamentsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
