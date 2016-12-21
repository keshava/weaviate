package adapters


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaveAdaptersGetHandlerFunc turns a function with the right signature into a weave adapters get handler
type WeaveAdaptersGetHandlerFunc func(WeaveAdaptersGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaveAdaptersGetHandlerFunc) Handle(params WeaveAdaptersGetParams) middleware.Responder {
	return fn(params)
}

// WeaveAdaptersGetHandler interface for that can handle valid weave adapters get params
type WeaveAdaptersGetHandler interface {
	Handle(WeaveAdaptersGetParams) middleware.Responder
}

// NewWeaveAdaptersGet creates a new http.Handler for the weave adapters get operation
func NewWeaveAdaptersGet(ctx *middleware.Context, handler WeaveAdaptersGetHandler) *WeaveAdaptersGet {
	return &WeaveAdaptersGet{Context: ctx, Handler: handler}
}

/*WeaveAdaptersGet swagger:route GET /adapters/{adapterId} adapters weaveAdaptersGet

Get an adapter.

*/
type WeaveAdaptersGet struct {
	Context *middleware.Context
	Handler WeaveAdaptersGetHandler
}

func (o *WeaveAdaptersGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaveAdaptersGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}