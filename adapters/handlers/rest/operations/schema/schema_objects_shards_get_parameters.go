//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewSchemaObjectsShardsGetParams creates a new SchemaObjectsShardsGetParams object
//
// There are no default values defined in the spec.
func NewSchemaObjectsShardsGetParams() SchemaObjectsShardsGetParams {

	return SchemaObjectsShardsGetParams{}
}

// SchemaObjectsShardsGetParams contains all the bound params for the schema objects shards get operation
// typically these are obtained from a http.Request
//
// swagger:parameters schema.objects.shards.get
type SchemaObjectsShardsGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ClassName string
	/*
	  In: query
	*/
	Tenant *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSchemaObjectsShardsGetParams() beforehand.
func (o *SchemaObjectsShardsGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rClassName, rhkClassName, _ := route.Params.GetOK("className")
	if err := o.bindClassName(rClassName, rhkClassName, route.Formats); err != nil {
		res = append(res, err)
	}

	qTenant, qhkTenant, _ := qs.GetOK("tenant")
	if err := o.bindTenant(qTenant, qhkTenant, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindClassName binds and validates parameter ClassName from path.
func (o *SchemaObjectsShardsGetParams) bindClassName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ClassName = raw

	return nil
}

// bindTenant binds and validates parameter Tenant from query.
func (o *SchemaObjectsShardsGetParams) bindTenant(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Tenant = &raw

	return nil
}
