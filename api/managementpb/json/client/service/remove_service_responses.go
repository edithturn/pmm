// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RemoveServiceReader is a Reader for the RemoveService structure.
type RemoveServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRemoveServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveServiceOK creates a RemoveServiceOK with default headers values
func NewRemoveServiceOK() *RemoveServiceOK {
	return &RemoveServiceOK{}
}

/* RemoveServiceOK describes a response with status code 200, with default header values.

A successful response.
*/
type RemoveServiceOK struct {
	Payload interface{}
}

func (o *RemoveServiceOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/Service/Remove][%d] removeServiceOk  %+v", 200, o.Payload)
}

func (o *RemoveServiceOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RemoveServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveServiceDefault creates a RemoveServiceDefault with default headers values
func NewRemoveServiceDefault(code int) *RemoveServiceDefault {
	return &RemoveServiceDefault{
		_statusCode: code,
	}
}

/* RemoveServiceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RemoveServiceDefault struct {
	_statusCode int

	Payload *RemoveServiceDefaultBody
}

// Code gets the status code for the remove service default response
func (o *RemoveServiceDefault) Code() int {
	return o._statusCode
}

func (o *RemoveServiceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/Service/Remove][%d] RemoveService default  %+v", o._statusCode, o.Payload)
}

func (o *RemoveServiceDefault) GetPayload() *RemoveServiceDefaultBody {
	return o.Payload
}

func (o *RemoveServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(RemoveServiceDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RemoveServiceBody remove service body
swagger:model RemoveServiceBody
*/
type RemoveServiceBody struct {
	// ServiceType describes supported Service types.
	// Enum: [SERVICE_TYPE_INVALID MYSQL_SERVICE MONGODB_SERVICE POSTGRESQL_SERVICE PROXYSQL_SERVICE HAPROXY_SERVICE EXTERNAL_SERVICE]
	ServiceType *string `json:"service_type,omitempty"`

	// Service ID or Service Name is required.
	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

	// agent id
	AgentID string `json:"agent_id,omitempty"`
}

// Validate validates this remove service body
func (o *RemoveServiceBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateServiceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var removeServiceBodyTypeServiceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SERVICE_TYPE_INVALID","MYSQL_SERVICE","MONGODB_SERVICE","POSTGRESQL_SERVICE","PROXYSQL_SERVICE","HAPROXY_SERVICE","EXTERNAL_SERVICE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		removeServiceBodyTypeServiceTypePropEnum = append(removeServiceBodyTypeServiceTypePropEnum, v)
	}
}

const (

	// RemoveServiceBodyServiceTypeSERVICETYPEINVALID captures enum value "SERVICE_TYPE_INVALID"
	RemoveServiceBodyServiceTypeSERVICETYPEINVALID string = "SERVICE_TYPE_INVALID"

	// RemoveServiceBodyServiceTypeMYSQLSERVICE captures enum value "MYSQL_SERVICE"
	RemoveServiceBodyServiceTypeMYSQLSERVICE string = "MYSQL_SERVICE"

	// RemoveServiceBodyServiceTypeMONGODBSERVICE captures enum value "MONGODB_SERVICE"
	RemoveServiceBodyServiceTypeMONGODBSERVICE string = "MONGODB_SERVICE"

	// RemoveServiceBodyServiceTypePOSTGRESQLSERVICE captures enum value "POSTGRESQL_SERVICE"
	RemoveServiceBodyServiceTypePOSTGRESQLSERVICE string = "POSTGRESQL_SERVICE"

	// RemoveServiceBodyServiceTypePROXYSQLSERVICE captures enum value "PROXYSQL_SERVICE"
	RemoveServiceBodyServiceTypePROXYSQLSERVICE string = "PROXYSQL_SERVICE"

	// RemoveServiceBodyServiceTypeHAPROXYSERVICE captures enum value "HAPROXY_SERVICE"
	RemoveServiceBodyServiceTypeHAPROXYSERVICE string = "HAPROXY_SERVICE"

	// RemoveServiceBodyServiceTypeEXTERNALSERVICE captures enum value "EXTERNAL_SERVICE"
	RemoveServiceBodyServiceTypeEXTERNALSERVICE string = "EXTERNAL_SERVICE"
)

// prop value enum
func (o *RemoveServiceBody) validateServiceTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, removeServiceBodyTypeServiceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *RemoveServiceBody) validateServiceType(formats strfmt.Registry) error {
	if swag.IsZero(o.ServiceType) { // not required
		return nil
	}

	// value enum
	if err := o.validateServiceTypeEnum("body"+"."+"service_type", "body", *o.ServiceType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this remove service body based on context it is used
func (o *RemoveServiceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveServiceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveServiceBody) UnmarshalBinary(b []byte) error {
	var res RemoveServiceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoveServiceDefaultBody remove service default body
swagger:model RemoveServiceDefaultBody
*/
type RemoveServiceDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*RemoveServiceDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this remove service default body
func (o *RemoveServiceDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RemoveServiceDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RemoveService default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RemoveService default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this remove service default body based on the context it is used
func (o *RemoveServiceDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RemoveServiceDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RemoveService default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RemoveService default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RemoveServiceDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveServiceDefaultBody) UnmarshalBinary(b []byte) error {
	var res RemoveServiceDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoveServiceDefaultBodyDetailsItems0 remove service default body details items0
swagger:model RemoveServiceDefaultBodyDetailsItems0
*/
type RemoveServiceDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this remove service default body details items0
func (o *RemoveServiceDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this remove service default body details items0 based on context it is used
func (o *RemoveServiceDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveServiceDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveServiceDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res RemoveServiceDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
