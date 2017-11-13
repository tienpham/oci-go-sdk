// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"net/http"
)

// DeleteCpeRequest wrapper for the DeleteCpe operation
type DeleteCpeRequest struct {

	// The OCID of the CPE.
	CpeID *string `mandatory:"true" contributesTo:"path" name:"cpeId"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

func (request DeleteCpeRequest) String() string {
	return common.PointerString(request)
}

// DeleteCpeResponse wrapper for the DeleteCpe operation
type DeleteCpeResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response DeleteCpeResponse) String() string {
	return common.PointerString(response)
}