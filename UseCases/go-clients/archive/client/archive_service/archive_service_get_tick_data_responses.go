// Code generated by go-swagger; DO NOT EDIT.

package archive_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/qubic/integration/go-clients/archive/models"
)

// ArchiveServiceGetTickDataReader is a Reader for the ArchiveServiceGetTickData structure.
type ArchiveServiceGetTickDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArchiveServiceGetTickDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewArchiveServiceGetTickDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewArchiveServiceGetTickDataDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewArchiveServiceGetTickDataOK creates a ArchiveServiceGetTickDataOK with default headers values
func NewArchiveServiceGetTickDataOK() *ArchiveServiceGetTickDataOK {
	return &ArchiveServiceGetTickDataOK{}
}

/*
ArchiveServiceGetTickDataOK describes a response with status code 200, with default header values.

A successful response.
*/
type ArchiveServiceGetTickDataOK struct {
	Payload *models.PbGetTickDataResponse
}

// IsSuccess returns true when this archive service get tick data o k response has a 2xx status code
func (o *ArchiveServiceGetTickDataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this archive service get tick data o k response has a 3xx status code
func (o *ArchiveServiceGetTickDataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive service get tick data o k response has a 4xx status code
func (o *ArchiveServiceGetTickDataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this archive service get tick data o k response has a 5xx status code
func (o *ArchiveServiceGetTickDataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this archive service get tick data o k response a status code equal to that given
func (o *ArchiveServiceGetTickDataOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the archive service get tick data o k response
func (o *ArchiveServiceGetTickDataOK) Code() int {
	return 200
}

func (o *ArchiveServiceGetTickDataOK) Error() string {
	return fmt.Sprintf("[GET /ticks/{tickNumber}/tick-data][%d] archiveServiceGetTickDataOK  %+v", 200, o.Payload)
}

func (o *ArchiveServiceGetTickDataOK) String() string {
	return fmt.Sprintf("[GET /ticks/{tickNumber}/tick-data][%d] archiveServiceGetTickDataOK  %+v", 200, o.Payload)
}

func (o *ArchiveServiceGetTickDataOK) GetPayload() *models.PbGetTickDataResponse {
	return o.Payload
}

func (o *ArchiveServiceGetTickDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PbGetTickDataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArchiveServiceGetTickDataDefault creates a ArchiveServiceGetTickDataDefault with default headers values
func NewArchiveServiceGetTickDataDefault(code int) *ArchiveServiceGetTickDataDefault {
	return &ArchiveServiceGetTickDataDefault{
		_statusCode: code,
	}
}

/*
ArchiveServiceGetTickDataDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ArchiveServiceGetTickDataDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this archive service get tick data default response has a 2xx status code
func (o *ArchiveServiceGetTickDataDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this archive service get tick data default response has a 3xx status code
func (o *ArchiveServiceGetTickDataDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this archive service get tick data default response has a 4xx status code
func (o *ArchiveServiceGetTickDataDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this archive service get tick data default response has a 5xx status code
func (o *ArchiveServiceGetTickDataDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this archive service get tick data default response a status code equal to that given
func (o *ArchiveServiceGetTickDataDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the archive service get tick data default response
func (o *ArchiveServiceGetTickDataDefault) Code() int {
	return o._statusCode
}

func (o *ArchiveServiceGetTickDataDefault) Error() string {
	return fmt.Sprintf("[GET /ticks/{tickNumber}/tick-data][%d] ArchiveService_GetTickData default  %+v", o._statusCode, o.Payload)
}

func (o *ArchiveServiceGetTickDataDefault) String() string {
	return fmt.Sprintf("[GET /ticks/{tickNumber}/tick-data][%d] ArchiveService_GetTickData default  %+v", o._statusCode, o.Payload)
}

func (o *ArchiveServiceGetTickDataDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ArchiveServiceGetTickDataDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
