// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
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

// ChangeAzureDatabaseExporterReader is a Reader for the ChangeAzureDatabaseExporter structure.
type ChangeAzureDatabaseExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeAzureDatabaseExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeAzureDatabaseExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeAzureDatabaseExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeAzureDatabaseExporterOK creates a ChangeAzureDatabaseExporterOK with default headers values
func NewChangeAzureDatabaseExporterOK() *ChangeAzureDatabaseExporterOK {
	return &ChangeAzureDatabaseExporterOK{}
}

/*ChangeAzureDatabaseExporterOK handles this case with default header values.

A successful response.
*/
type ChangeAzureDatabaseExporterOK struct {
	Payload *ChangeAzureDatabaseExporterOKBody
}

func (o *ChangeAzureDatabaseExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeAzureDatabaseExporter][%d] changeAzureDatabaseExporterOk  %+v", 200, o.Payload)
}

func (o *ChangeAzureDatabaseExporterOK) GetPayload() *ChangeAzureDatabaseExporterOKBody {
	return o.Payload
}

func (o *ChangeAzureDatabaseExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeAzureDatabaseExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeAzureDatabaseExporterDefault creates a ChangeAzureDatabaseExporterDefault with default headers values
func NewChangeAzureDatabaseExporterDefault(code int) *ChangeAzureDatabaseExporterDefault {
	return &ChangeAzureDatabaseExporterDefault{
		_statusCode: code,
	}
}

/*ChangeAzureDatabaseExporterDefault handles this case with default header values.

An unexpected error response.
*/
type ChangeAzureDatabaseExporterDefault struct {
	_statusCode int

	Payload *ChangeAzureDatabaseExporterDefaultBody
}

// Code gets the status code for the change azure database exporter default response
func (o *ChangeAzureDatabaseExporterDefault) Code() int {
	return o._statusCode
}

func (o *ChangeAzureDatabaseExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeAzureDatabaseExporter][%d] ChangeAzureDatabaseExporter default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeAzureDatabaseExporterDefault) GetPayload() *ChangeAzureDatabaseExporterDefaultBody {
	return o.Payload
}

func (o *ChangeAzureDatabaseExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeAzureDatabaseExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeAzureDatabaseExporterBody change azure database exporter body
swagger:model ChangeAzureDatabaseExporterBody
*/
type ChangeAzureDatabaseExporterBody struct {

	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// common
	Common *ChangeAzureDatabaseExporterParamsBodyCommon `json:"common,omitempty"`
}

// Validate validates this change azure database exporter body
func (o *ChangeAzureDatabaseExporterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeAzureDatabaseExporterBody) validateCommon(formats strfmt.Registry) error {

	if swag.IsZero(o.Common) { // not required
		return nil
	}

	if o.Common != nil {
		if err := o.Common.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "common")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeAzureDatabaseExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeAzureDatabaseExporterBody) UnmarshalBinary(b []byte) error {
	var res ChangeAzureDatabaseExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeAzureDatabaseExporterDefaultBody change azure database exporter default body
swagger:model ChangeAzureDatabaseExporterDefaultBody
*/
type ChangeAzureDatabaseExporterDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this change azure database exporter default body
func (o *ChangeAzureDatabaseExporterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeAzureDatabaseExporterDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("ChangeAzureDatabaseExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeAzureDatabaseExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeAzureDatabaseExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeAzureDatabaseExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeAzureDatabaseExporterOKBody change azure database exporter OK body
swagger:model ChangeAzureDatabaseExporterOKBody
*/
type ChangeAzureDatabaseExporterOKBody struct {

	// azure database exporter
	AzureDatabaseExporter *ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporter `json:"azure_database_exporter,omitempty"`
}

// Validate validates this change azure database exporter OK body
func (o *ChangeAzureDatabaseExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAzureDatabaseExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeAzureDatabaseExporterOKBody) validateAzureDatabaseExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.AzureDatabaseExporter) { // not required
		return nil
	}

	if o.AzureDatabaseExporter != nil {
		if err := o.AzureDatabaseExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeAzureDatabaseExporterOk" + "." + "azure_database_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeAzureDatabaseExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeAzureDatabaseExporterOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeAzureDatabaseExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporter AzureDatabaseExporter runs on Generic or Container Node and exposes RemoteAzure Node metrics.
swagger:model ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporter
*/
type ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Node identifier.
	NodeID string `json:"node_id,omitempty"`

	// Azure database subscription ID.
	AzureDatabaseSubscriptionID string `json:"azure_database_subscription_id,omitempty"`

	// Azure database resource type (mysql, maria, postgres)
	AzureDatabaseResourceType string `json:"azure_database_resource_type,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// AgentStatus represents actual Agent status.
	//
	//  - STARTING: Agent is starting.
	//  - RUNNING: Agent is running.
	//  - WAITING: Agent encountered error and will be restarted automatically soon.
	//  - STOPPING: Agent is stopping.
	//  - DONE: Agent finished.
	//  - UNKNOWN: Agent is not connected, we don't know anything about it's state.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE UNKNOWN]
	Status *string `json:"status,omitempty"`

	// Listen port for scraping metrics (the same for several configurations).
	ListenPort int64 `json:"listen_port,omitempty"`

	// True if the exporter operates in push metrics mode.
	PushMetricsEnabled bool `json:"push_metrics_enabled,omitempty"`
}

// Validate validates this change azure database exporter OK body azure database exporter
func (o *ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var changeAzureDatabaseExporterOkBodyAzureDatabaseExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeAzureDatabaseExporterOkBodyAzureDatabaseExporterTypeStatusPropEnum = append(changeAzureDatabaseExporterOkBodyAzureDatabaseExporterTypeStatusPropEnum, v)
	}
}

const (

	// ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusSTARTING captures enum value "STARTING"
	ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusSTARTING string = "STARTING"

	// ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusRUNNING captures enum value "RUNNING"
	ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusRUNNING string = "RUNNING"

	// ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusWAITING captures enum value "WAITING"
	ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusWAITING string = "WAITING"

	// ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusSTOPPING captures enum value "STOPPING"
	ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusSTOPPING string = "STOPPING"

	// ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusDONE captures enum value "DONE"
	ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusDONE string = "DONE"

	// ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusUNKNOWN captures enum value "UNKNOWN"
	ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporterStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (o *ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, changeAzureDatabaseExporterOkBodyAzureDatabaseExporterTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("changeAzureDatabaseExporterOk"+"."+"azure_database_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporter) UnmarshalBinary(b []byte) error {
	var res ChangeAzureDatabaseExporterOKBodyAzureDatabaseExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeAzureDatabaseExporterParamsBodyCommon ChangeCommonAgentParams contains parameters that can be changed for all Agents.
swagger:model ChangeAzureDatabaseExporterParamsBodyCommon
*/
type ChangeAzureDatabaseExporterParamsBodyCommon struct {

	// Enable this Agent. Can't be used with disabled.
	Enable bool `json:"enable,omitempty"`

	// Disable this Agent. Can't be used with enabled.
	Disable bool `json:"disable,omitempty"`

	// Replace all custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Remove all custom user-assigned labels.
	RemoveCustomLabels bool `json:"remove_custom_labels,omitempty"`

	// Enables push metrics with vmagent, can't be used with disable_push_metrics.
	// Can't be used with agent version lower then 2.12 and unsupported agents.
	EnablePushMetrics bool `json:"enable_push_metrics,omitempty"`

	// Disables push metrics, pmm-server starts to pull it, can't be used with enable_push_metrics.
	DisablePushMetrics bool `json:"disable_push_metrics,omitempty"`
}

// Validate validates this change azure database exporter params body common
func (o *ChangeAzureDatabaseExporterParamsBodyCommon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeAzureDatabaseExporterParamsBodyCommon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeAzureDatabaseExporterParamsBodyCommon) UnmarshalBinary(b []byte) error {
	var res ChangeAzureDatabaseExporterParamsBodyCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
