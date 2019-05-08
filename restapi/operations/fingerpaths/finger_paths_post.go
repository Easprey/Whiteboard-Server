// Code generated by go-swagger; DO NOT EDIT.

package fingerpaths

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// FingerPathsPostHandlerFunc turns a function with the right signature into a finger paths post handler
type FingerPathsPostHandlerFunc func(FingerPathsPostParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FingerPathsPostHandlerFunc) Handle(params FingerPathsPostParams) middleware.Responder {
	return fn(params)
}

// FingerPathsPostHandler interface for that can handle valid finger paths post params
type FingerPathsPostHandler interface {
	Handle(FingerPathsPostParams) middleware.Responder
}

// NewFingerPathsPost creates a new http.Handler for the finger paths post operation
func NewFingerPathsPost(ctx *middleware.Context, handler FingerPathsPostHandler) *FingerPathsPost {
	return &FingerPathsPost{Context: ctx, Handler: handler}
}

/*FingerPathsPost swagger:route POST /boards/{boardName}/fingerpaths fingerpaths fingerPathsPost

Adds fingerpaths to the associated board.

*/
type FingerPathsPost struct {
	Context *middleware.Context
	Handler FingerPathsPostHandler
}

func (o *FingerPathsPost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFingerPathsPostParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}