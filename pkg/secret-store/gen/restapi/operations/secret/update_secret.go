///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package secret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateSecretHandlerFunc turns a function with the right signature into a update secret handler
type UpdateSecretHandlerFunc func(UpdateSecretParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateSecretHandlerFunc) Handle(params UpdateSecretParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateSecretHandler interface for that can handle valid update secret params
type UpdateSecretHandler interface {
	Handle(UpdateSecretParams, interface{}) middleware.Responder
}

// NewUpdateSecret creates a new http.Handler for the update secret operation
func NewUpdateSecret(ctx *middleware.Context, handler UpdateSecretHandler) *UpdateSecret {
	return &UpdateSecret{Context: ctx, Handler: handler}
}

/*UpdateSecret swagger:route PUT /{secretName} secret updateSecret

UpdateSecret update secret API

*/
type UpdateSecret struct {
	Context *middleware.Context
	Handler UpdateSecretHandler
}

func (o *UpdateSecret) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateSecretParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
