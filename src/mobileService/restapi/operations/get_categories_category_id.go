// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCategoriesCategoryIDHandlerFunc turns a function with the right signature into a get categories category ID handler
type GetCategoriesCategoryIDHandlerFunc func(GetCategoriesCategoryIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCategoriesCategoryIDHandlerFunc) Handle(params GetCategoriesCategoryIDParams) middleware.Responder {
	return fn(params)
}

// GetCategoriesCategoryIDHandler interface for that can handle valid get categories category ID params
type GetCategoriesCategoryIDHandler interface {
	Handle(GetCategoriesCategoryIDParams) middleware.Responder
}

// NewGetCategoriesCategoryID creates a new http.Handler for the get categories category ID operation
func NewGetCategoriesCategoryID(ctx *middleware.Context, handler GetCategoriesCategoryIDHandler) *GetCategoriesCategoryID {
	return &GetCategoriesCategoryID{Context: ctx, Handler: handler}
}

/*GetCategoriesCategoryID swagger:route GET /categories/{categoryId} getCategoriesCategoryId

Gets category articles

Gets all the articles for an category

*/
type GetCategoriesCategoryID struct {
	Context *middleware.Context
	Handler GetCategoriesCategoryIDHandler
}

func (o *GetCategoriesCategoryID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetCategoriesCategoryIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
