// Code generated by go-swagger; DO NOT EDIT.

// Copyright Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package silence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSilencesHandlerFunc turns a function with the right signature into a get silences handler
type GetSilencesHandlerFunc func(GetSilencesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSilencesHandlerFunc) Handle(params GetSilencesParams) middleware.Responder {
	return fn(params)
}

// GetSilencesHandler interface for that can handle valid get silences params
type GetSilencesHandler interface {
	Handle(GetSilencesParams) middleware.Responder
}

// NewGetSilences creates a new http.Handler for the get silences operation
func NewGetSilences(ctx *middleware.Context, handler GetSilencesHandler) *GetSilences {
	return &GetSilences{Context: ctx, Handler: handler}
}

/*GetSilences swagger:route GET /api/v2/silences silence getSilences

Get a list of silences

*/
type GetSilences struct {
	Context *middleware.Context
	Handler GetSilencesHandler
}

func (o *GetSilences) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSilencesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
