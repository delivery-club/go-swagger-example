// Code generated by go-swagger; DO NOT EDIT.

package hello

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/delivery-club/go-swagger-example/example3/pkg/example3/models"
)

// HelloWorldFullReader is a Reader for the HelloWorldFull structure.
type HelloWorldFullReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HelloWorldFullReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHelloWorldFullOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHelloWorldFullBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHelloWorldFullInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHelloWorldFullOK creates a HelloWorldFullOK with default headers values
func NewHelloWorldFullOK() *HelloWorldFullOK {
	return &HelloWorldFullOK{}
}

/*HelloWorldFullOK handles this case with default header values.

successful response
*/
type HelloWorldFullOK struct {
	Payload *models.HelloWorld
}

func (o *HelloWorldFullOK) Error() string {
	return fmt.Sprintf("[GET /hello-world][%d] helloWorldFullOK  %+v", 200, o.Payload)
}

func (o *HelloWorldFullOK) GetPayload() *models.HelloWorld {
	return o.Payload
}

func (o *HelloWorldFullOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HelloWorld)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHelloWorldFullBadRequest creates a HelloWorldFullBadRequest with default headers values
func NewHelloWorldFullBadRequest() *HelloWorldFullBadRequest {
	return &HelloWorldFullBadRequest{}
}

/*HelloWorldFullBadRequest handles this case with default header values.

Bad request
*/
type HelloWorldFullBadRequest struct {
	Payload *models.Error
}

func (o *HelloWorldFullBadRequest) Error() string {
	return fmt.Sprintf("[GET /hello-world][%d] helloWorldFullBadRequest  %+v", 400, o.Payload)
}

func (o *HelloWorldFullBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *HelloWorldFullBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHelloWorldFullInternalServerError creates a HelloWorldFullInternalServerError with default headers values
func NewHelloWorldFullInternalServerError() *HelloWorldFullInternalServerError {
	return &HelloWorldFullInternalServerError{}
}

/*HelloWorldFullInternalServerError handles this case with default header values.

Internal server error
*/
type HelloWorldFullInternalServerError struct {
	Payload *models.Error
}

func (o *HelloWorldFullInternalServerError) Error() string {
	return fmt.Sprintf("[GET /hello-world][%d] helloWorldFullInternalServerError  %+v", 500, o.Payload)
}

func (o *HelloWorldFullInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *HelloWorldFullInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
