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

// NewUploadNewBoardCoverParams creates a new UploadNewBoardCoverParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadNewBoardCoverParams() *UploadNewBoardCoverParams {
	return &UploadNewBoardCoverParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadNewBoardCoverParamsWithTimeout creates a new UploadNewBoardCoverParams object
// with the ability to set a timeout on a request.
func NewUploadNewBoardCoverParamsWithTimeout(timeout time.Duration) *UploadNewBoardCoverParams {
	return &UploadNewBoardCoverParams{
		timeout: timeout,
	}
}

// NewUploadNewBoardCoverParamsWithContext creates a new UploadNewBoardCoverParams object
// with the ability to set a context for a request.
func NewUploadNewBoardCoverParamsWithContext(ctx context.Context) *UploadNewBoardCoverParams {
	return &UploadNewBoardCoverParams{
		Context: ctx,
	}
}

// NewUploadNewBoardCoverParamsWithHTTPClient creates a new UploadNewBoardCoverParams object
// with the ability to set a custom HTTPClient for a request.
func NewUploadNewBoardCoverParamsWithHTTPClient(client *http.Client) *UploadNewBoardCoverParams {
	return &UploadNewBoardCoverParams{
		HTTPClient: client,
	}
}

/* UploadNewBoardCoverParams contains all the parameters to send to the API endpoint
   for the upload new board cover operation.

   Typically these are written to a http.Request.
*/
type UploadNewBoardCoverParams struct {

	// Cover.
	Cover runtime.NamedReadCloser

	// TagID.
	TagID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upload new board cover params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadNewBoardCoverParams) WithDefaults() *UploadNewBoardCoverParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload new board cover params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadNewBoardCoverParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upload new board cover params
func (o *UploadNewBoardCoverParams) WithTimeout(timeout time.Duration) *UploadNewBoardCoverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload new board cover params
func (o *UploadNewBoardCoverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload new board cover params
func (o *UploadNewBoardCoverParams) WithContext(ctx context.Context) *UploadNewBoardCoverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload new board cover params
func (o *UploadNewBoardCoverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload new board cover params
func (o *UploadNewBoardCoverParams) WithHTTPClient(client *http.Client) *UploadNewBoardCoverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload new board cover params
func (o *UploadNewBoardCoverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCover adds the cover to the upload new board cover params
func (o *UploadNewBoardCoverParams) WithCover(cover runtime.NamedReadCloser) *UploadNewBoardCoverParams {
	o.SetCover(cover)
	return o
}

// SetCover adds the cover to the upload new board cover params
func (o *UploadNewBoardCoverParams) SetCover(cover runtime.NamedReadCloser) {
	o.Cover = cover
}

// WithTagID adds the tagID to the upload new board cover params
func (o *UploadNewBoardCoverParams) WithTagID(tagID string) *UploadNewBoardCoverParams {
	o.SetTagID(tagID)
	return o
}

// SetTagID adds the tagId to the upload new board cover params
func (o *UploadNewBoardCoverParams) SetTagID(tagID string) {
	o.TagID = tagID
}

// WriteToRequest writes these params to a swagger request
func (o *UploadNewBoardCoverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	// form file param cover
	if err := r.SetFileParam("cover", o.Cover); err != nil {
		return err
	}

	// path param tagId
	if err := r.SetPathParam("tagId", o.TagID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
