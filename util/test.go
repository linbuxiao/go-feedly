package util

import "github.com/jarcoal/httpmock"

func NewJsonResponderFromString(status int, text string) httpmock.Responder {
	r := httpmock.NewStringResponse(status, text)
	r.Header.Set("Content-Type", "application/json")
	return httpmock.ResponderFromResponse(r)
}
