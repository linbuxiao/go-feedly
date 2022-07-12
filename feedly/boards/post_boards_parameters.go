// Code generated by go-swagger; DO NOT EDIT.

package boards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPostBoardsParams creates a new PostBoardsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostBoardsParams() *PostBoardsParams {
	return &PostBoardsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostBoardsParamsWithTimeout creates a new PostBoardsParams object
// with the ability to set a timeout on a request.
func NewPostBoardsParamsWithTimeout(timeout time.Duration) *PostBoardsParams {
	return &PostBoardsParams{
		timeout: timeout,
	}
}

// NewPostBoardsParamsWithContext creates a new PostBoardsParams object
// with the ability to set a context for a request.
func NewPostBoardsParamsWithContext(ctx context.Context) *PostBoardsParams {
	return &PostBoardsParams{
		Context: ctx,
	}
}

// NewPostBoardsParamsWithHTTPClient creates a new PostBoardsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostBoardsParamsWithHTTPClient(client *http.Client) *PostBoardsParams {
	return &PostBoardsParams{
		HTTPClient: client,
	}
}

/* PostBoardsParams contains all the parameters to send to the API endpoint
   for the post boards operation.

   Typically these are written to a http.Request.
*/
type PostBoardsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post boards params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostBoardsParams) WithDefaults() *PostBoardsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post boards params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostBoardsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post boards params
func (o *PostBoardsParams) WithTimeout(timeout time.Duration) *PostBoardsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post boards params
func (o *PostBoardsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post boards params
func (o *PostBoardsParams) WithContext(ctx context.Context) *PostBoardsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post boards params
func (o *PostBoardsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post boards params
func (o *PostBoardsParams) WithHTTPClient(client *http.Client) *PostBoardsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post boards params
func (o *PostBoardsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostBoardsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
