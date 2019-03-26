// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/generated/models"
)

// GetMicroBlockTransactionByHashAndIndexReader is a Reader for the GetMicroBlockTransactionByHashAndIndex structure.
type GetMicroBlockTransactionByHashAndIndexReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMicroBlockTransactionByHashAndIndexReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMicroBlockTransactionByHashAndIndexOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetMicroBlockTransactionByHashAndIndexBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetMicroBlockTransactionByHashAndIndexNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMicroBlockTransactionByHashAndIndexOK creates a GetMicroBlockTransactionByHashAndIndexOK with default headers values
func NewGetMicroBlockTransactionByHashAndIndexOK() *GetMicroBlockTransactionByHashAndIndexOK {
	return &GetMicroBlockTransactionByHashAndIndexOK{}
}

/*GetMicroBlockTransactionByHashAndIndexOK handles this case with default header values.

Successful operation
*/
type GetMicroBlockTransactionByHashAndIndexOK struct {
	Payload *models.GenericSignedTx
}

func (o *GetMicroBlockTransactionByHashAndIndexOK) Error() string {
	return fmt.Sprintf("[GET /micro-blocks/hash/{hash}/transactions/index/{index}][%d] getMicroBlockTransactionByHashAndIndexOK  %+v", 200, o.Payload)
}

func (o *GetMicroBlockTransactionByHashAndIndexOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericSignedTx)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMicroBlockTransactionByHashAndIndexBadRequest creates a GetMicroBlockTransactionByHashAndIndexBadRequest with default headers values
func NewGetMicroBlockTransactionByHashAndIndexBadRequest() *GetMicroBlockTransactionByHashAndIndexBadRequest {
	return &GetMicroBlockTransactionByHashAndIndexBadRequest{}
}

/*GetMicroBlockTransactionByHashAndIndexBadRequest handles this case with default header values.

Invalid hash or index
*/
type GetMicroBlockTransactionByHashAndIndexBadRequest struct {
	Payload *models.Error
}

func (o *GetMicroBlockTransactionByHashAndIndexBadRequest) Error() string {
	return fmt.Sprintf("[GET /micro-blocks/hash/{hash}/transactions/index/{index}][%d] getMicroBlockTransactionByHashAndIndexBadRequest  %+v", 400, o.Payload)
}

func (o *GetMicroBlockTransactionByHashAndIndexBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMicroBlockTransactionByHashAndIndexNotFound creates a GetMicroBlockTransactionByHashAndIndexNotFound with default headers values
func NewGetMicroBlockTransactionByHashAndIndexNotFound() *GetMicroBlockTransactionByHashAndIndexNotFound {
	return &GetMicroBlockTransactionByHashAndIndexNotFound{}
}

/*GetMicroBlockTransactionByHashAndIndexNotFound handles this case with default header values.

Block not found
*/
type GetMicroBlockTransactionByHashAndIndexNotFound struct {
	Payload *models.Error
}

func (o *GetMicroBlockTransactionByHashAndIndexNotFound) Error() string {
	return fmt.Sprintf("[GET /micro-blocks/hash/{hash}/transactions/index/{index}][%d] getMicroBlockTransactionByHashAndIndexNotFound  %+v", 404, o.Payload)
}

func (o *GetMicroBlockTransactionByHashAndIndexNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
