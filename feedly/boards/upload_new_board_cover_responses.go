// Code generated by go-swagger; DO NOT EDIT.

package boards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UploadNewBoardCoverReader is a Reader for the UploadNewBoardCover structure.
type UploadNewBoardCoverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadNewBoardCoverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadNewBoardCoverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUploadNewBoardCoverOK creates a UploadNewBoardCoverOK with default headers values
func NewUploadNewBoardCoverOK() *UploadNewBoardCoverOK {
	return &UploadNewBoardCoverOK{}
}

/* UploadNewBoardCoverOK describes a response with status code 200, with default header values.

Success
*/
type UploadNewBoardCoverOK struct {
}

func (o *UploadNewBoardCoverOK) Error() string {
	return fmt.Sprintf("[POST /boards/{tagId}][%d] uploadNewBoardCoverOK ", 200)
}

func (o *UploadNewBoardCoverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
