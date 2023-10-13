// Code generated by go-swagger; DO NOT EDIT.

// Copyright Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package alertinfo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/prometheus/alertmanager/api/v2/models"
)

// GetAlertInfosReader is a Reader for the GetAlertInfos structure.
type GetAlertInfosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertInfosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertInfosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAlertInfosBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAlertInfosInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAlertInfosOK creates a GetAlertInfosOK with default headers values
func NewGetAlertInfosOK() *GetAlertInfosOK {
	return &GetAlertInfosOK{}
}

/*
GetAlertInfosOK describes a response with status code 200, with default header values.

Get alerts response
*/
type GetAlertInfosOK struct {
	Payload *models.GettableAlertInfos
}

// IsSuccess returns true when this get alert infos o k response has a 2xx status code
func (o *GetAlertInfosOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alert infos o k response has a 3xx status code
func (o *GetAlertInfosOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert infos o k response has a 4xx status code
func (o *GetAlertInfosOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alert infos o k response has a 5xx status code
func (o *GetAlertInfosOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert infos o k response a status code equal to that given
func (o *GetAlertInfosOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAlertInfosOK) Error() string {
	return fmt.Sprintf("[GET /alertinfos][%d] getAlertInfosOK  %+v", 200, o.Payload)
}

func (o *GetAlertInfosOK) String() string {
	return fmt.Sprintf("[GET /alertinfos][%d] getAlertInfosOK  %+v", 200, o.Payload)
}

func (o *GetAlertInfosOK) GetPayload() *models.GettableAlertInfos {
	return o.Payload
}

func (o *GetAlertInfosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GettableAlertInfos)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertInfosBadRequest creates a GetAlertInfosBadRequest with default headers values
func NewGetAlertInfosBadRequest() *GetAlertInfosBadRequest {
	return &GetAlertInfosBadRequest{}
}

/*
GetAlertInfosBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetAlertInfosBadRequest struct {
	Payload string
}

// IsSuccess returns true when this get alert infos bad request response has a 2xx status code
func (o *GetAlertInfosBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alert infos bad request response has a 3xx status code
func (o *GetAlertInfosBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert infos bad request response has a 4xx status code
func (o *GetAlertInfosBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alert infos bad request response has a 5xx status code
func (o *GetAlertInfosBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert infos bad request response a status code equal to that given
func (o *GetAlertInfosBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAlertInfosBadRequest) Error() string {
	return fmt.Sprintf("[GET /alertinfos][%d] getAlertInfosBadRequest  %+v", 400, o.Payload)
}

func (o *GetAlertInfosBadRequest) String() string {
	return fmt.Sprintf("[GET /alertinfos][%d] getAlertInfosBadRequest  %+v", 400, o.Payload)
}

func (o *GetAlertInfosBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetAlertInfosBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertInfosInternalServerError creates a GetAlertInfosInternalServerError with default headers values
func NewGetAlertInfosInternalServerError() *GetAlertInfosInternalServerError {
	return &GetAlertInfosInternalServerError{}
}

/*
GetAlertInfosInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetAlertInfosInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this get alert infos internal server error response has a 2xx status code
func (o *GetAlertInfosInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alert infos internal server error response has a 3xx status code
func (o *GetAlertInfosInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert infos internal server error response has a 4xx status code
func (o *GetAlertInfosInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alert infos internal server error response has a 5xx status code
func (o *GetAlertInfosInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get alert infos internal server error response a status code equal to that given
func (o *GetAlertInfosInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAlertInfosInternalServerError) Error() string {
	return fmt.Sprintf("[GET /alertinfos][%d] getAlertInfosInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlertInfosInternalServerError) String() string {
	return fmt.Sprintf("[GET /alertinfos][%d] getAlertInfosInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlertInfosInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetAlertInfosInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}