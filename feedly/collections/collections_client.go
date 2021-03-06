// Code generated by go-swagger; DO NOT EDIT.

package collections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new collections API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for collections API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateOrUpdateCollection(params *CreateOrUpdateCollectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOrUpdateCollectionOK, error)

	GetFeedCollections(params *GetFeedCollectionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFeedCollectionsOK, error)

	GetPersonalCollectionDetail(params *GetPersonalCollectionDetailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPersonalCollectionDetailOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateOrUpdateCollection Create or update a personal collection
*/
func (a *Client) CreateOrUpdateCollection(params *CreateOrUpdateCollectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOrUpdateCollectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOrUpdateCollectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateOrUpdateCollection",
		Method:             "POST",
		PathPattern:        "/collections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateOrUpdateCollectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateOrUpdateCollectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateOrUpdateCollection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFeedCollections Get the list of personal collections
*/
func (a *Client) GetFeedCollections(params *GetFeedCollectionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFeedCollectionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFeedCollectionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetFeedCollections",
		Method:             "GET",
		PathPattern:        "/collections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFeedCollectionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFeedCollectionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetFeedCollections: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPersonalCollectionDetail Get details about a personal collection
*/
func (a *Client) GetPersonalCollectionDetail(params *GetPersonalCollectionDetailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPersonalCollectionDetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPersonalCollectionDetailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPersonalCollectionDetail",
		Method:             "GET",
		PathPattern:        "/collections/{collectionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPersonalCollectionDetailReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPersonalCollectionDetailOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPersonalCollectionDetail: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
