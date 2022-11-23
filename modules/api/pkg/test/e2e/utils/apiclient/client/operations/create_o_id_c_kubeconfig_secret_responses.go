// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// CreateOIDCKubeconfigSecretReader is a Reader for the CreateOIDCKubeconfigSecret structure.
type CreateOIDCKubeconfigSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOIDCKubeconfigSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOIDCKubeconfigSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateOIDCKubeconfigSecretCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateOIDCKubeconfigSecretUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateOIDCKubeconfigSecretForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateOIDCKubeconfigSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateOIDCKubeconfigSecretOK creates a CreateOIDCKubeconfigSecretOK with default headers values
func NewCreateOIDCKubeconfigSecretOK() *CreateOIDCKubeconfigSecretOK {
	return &CreateOIDCKubeconfigSecretOK{}
}

/*
CreateOIDCKubeconfigSecretOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type CreateOIDCKubeconfigSecretOK struct {
}

// IsSuccess returns true when this create o Id c kubeconfig secret o k response has a 2xx status code
func (o *CreateOIDCKubeconfigSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create o Id c kubeconfig secret o k response has a 3xx status code
func (o *CreateOIDCKubeconfigSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create o Id c kubeconfig secret o k response has a 4xx status code
func (o *CreateOIDCKubeconfigSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create o Id c kubeconfig secret o k response has a 5xx status code
func (o *CreateOIDCKubeconfigSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create o Id c kubeconfig secret o k response a status code equal to that given
func (o *CreateOIDCKubeconfigSecretOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateOIDCKubeconfigSecretOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/kubeconfig/secret][%d] createOIdCKubeconfigSecretOK ", 200)
}

func (o *CreateOIDCKubeconfigSecretOK) String() string {
	return fmt.Sprintf("[GET /api/v2/kubeconfig/secret][%d] createOIdCKubeconfigSecretOK ", 200)
}

func (o *CreateOIDCKubeconfigSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOIDCKubeconfigSecretCreated creates a CreateOIDCKubeconfigSecretCreated with default headers values
func NewCreateOIDCKubeconfigSecretCreated() *CreateOIDCKubeconfigSecretCreated {
	return &CreateOIDCKubeconfigSecretCreated{}
}

/*
CreateOIDCKubeconfigSecretCreated describes a response with status code 201, with default header values.

EmptyResponse is a empty response
*/
type CreateOIDCKubeconfigSecretCreated struct {
}

// IsSuccess returns true when this create o Id c kubeconfig secret created response has a 2xx status code
func (o *CreateOIDCKubeconfigSecretCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create o Id c kubeconfig secret created response has a 3xx status code
func (o *CreateOIDCKubeconfigSecretCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create o Id c kubeconfig secret created response has a 4xx status code
func (o *CreateOIDCKubeconfigSecretCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create o Id c kubeconfig secret created response has a 5xx status code
func (o *CreateOIDCKubeconfigSecretCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create o Id c kubeconfig secret created response a status code equal to that given
func (o *CreateOIDCKubeconfigSecretCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateOIDCKubeconfigSecretCreated) Error() string {
	return fmt.Sprintf("[GET /api/v2/kubeconfig/secret][%d] createOIdCKubeconfigSecretCreated ", 201)
}

func (o *CreateOIDCKubeconfigSecretCreated) String() string {
	return fmt.Sprintf("[GET /api/v2/kubeconfig/secret][%d] createOIdCKubeconfigSecretCreated ", 201)
}

func (o *CreateOIDCKubeconfigSecretCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOIDCKubeconfigSecretUnauthorized creates a CreateOIDCKubeconfigSecretUnauthorized with default headers values
func NewCreateOIDCKubeconfigSecretUnauthorized() *CreateOIDCKubeconfigSecretUnauthorized {
	return &CreateOIDCKubeconfigSecretUnauthorized{}
}

/*
CreateOIDCKubeconfigSecretUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateOIDCKubeconfigSecretUnauthorized struct {
}

// IsSuccess returns true when this create o Id c kubeconfig secret unauthorized response has a 2xx status code
func (o *CreateOIDCKubeconfigSecretUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create o Id c kubeconfig secret unauthorized response has a 3xx status code
func (o *CreateOIDCKubeconfigSecretUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create o Id c kubeconfig secret unauthorized response has a 4xx status code
func (o *CreateOIDCKubeconfigSecretUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create o Id c kubeconfig secret unauthorized response has a 5xx status code
func (o *CreateOIDCKubeconfigSecretUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create o Id c kubeconfig secret unauthorized response a status code equal to that given
func (o *CreateOIDCKubeconfigSecretUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateOIDCKubeconfigSecretUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/kubeconfig/secret][%d] createOIdCKubeconfigSecretUnauthorized ", 401)
}

func (o *CreateOIDCKubeconfigSecretUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/kubeconfig/secret][%d] createOIdCKubeconfigSecretUnauthorized ", 401)
}

func (o *CreateOIDCKubeconfigSecretUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOIDCKubeconfigSecretForbidden creates a CreateOIDCKubeconfigSecretForbidden with default headers values
func NewCreateOIDCKubeconfigSecretForbidden() *CreateOIDCKubeconfigSecretForbidden {
	return &CreateOIDCKubeconfigSecretForbidden{}
}

/*
CreateOIDCKubeconfigSecretForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateOIDCKubeconfigSecretForbidden struct {
}

// IsSuccess returns true when this create o Id c kubeconfig secret forbidden response has a 2xx status code
func (o *CreateOIDCKubeconfigSecretForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create o Id c kubeconfig secret forbidden response has a 3xx status code
func (o *CreateOIDCKubeconfigSecretForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create o Id c kubeconfig secret forbidden response has a 4xx status code
func (o *CreateOIDCKubeconfigSecretForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create o Id c kubeconfig secret forbidden response has a 5xx status code
func (o *CreateOIDCKubeconfigSecretForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create o Id c kubeconfig secret forbidden response a status code equal to that given
func (o *CreateOIDCKubeconfigSecretForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateOIDCKubeconfigSecretForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/kubeconfig/secret][%d] createOIdCKubeconfigSecretForbidden ", 403)
}

func (o *CreateOIDCKubeconfigSecretForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/kubeconfig/secret][%d] createOIdCKubeconfigSecretForbidden ", 403)
}

func (o *CreateOIDCKubeconfigSecretForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOIDCKubeconfigSecretDefault creates a CreateOIDCKubeconfigSecretDefault with default headers values
func NewCreateOIDCKubeconfigSecretDefault(code int) *CreateOIDCKubeconfigSecretDefault {
	return &CreateOIDCKubeconfigSecretDefault{
		_statusCode: code,
	}
}

/*
CreateOIDCKubeconfigSecretDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateOIDCKubeconfigSecretDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create o ID c kubeconfig secret default response
func (o *CreateOIDCKubeconfigSecretDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create o ID c kubeconfig secret default response has a 2xx status code
func (o *CreateOIDCKubeconfigSecretDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create o ID c kubeconfig secret default response has a 3xx status code
func (o *CreateOIDCKubeconfigSecretDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create o ID c kubeconfig secret default response has a 4xx status code
func (o *CreateOIDCKubeconfigSecretDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create o ID c kubeconfig secret default response has a 5xx status code
func (o *CreateOIDCKubeconfigSecretDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create o ID c kubeconfig secret default response a status code equal to that given
func (o *CreateOIDCKubeconfigSecretDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateOIDCKubeconfigSecretDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/kubeconfig/secret][%d] createOIDCKubeconfigSecret default  %+v", o._statusCode, o.Payload)
}

func (o *CreateOIDCKubeconfigSecretDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/kubeconfig/secret][%d] createOIDCKubeconfigSecret default  %+v", o._statusCode, o.Payload)
}

func (o *CreateOIDCKubeconfigSecretDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateOIDCKubeconfigSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}