// Code generated by go-swagger; DO NOT EDIT.

package ruler

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/percona/pmm/api/grafana/gmodels"
)

// RouteGetRulesConfigReader is a Reader for the RouteGetRulesConfig structure.
type RouteGetRulesConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RouteGetRulesConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRouteGetRulesConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRouteGetRulesConfigOK creates a RouteGetRulesConfigOK with default headers values
func NewRouteGetRulesConfigOK() *RouteGetRulesConfigOK {
	return &RouteGetRulesConfigOK{}
}

/* RouteGetRulesConfigOK describes a response with status code 200, with default header values.

NamespaceConfigResponse
*/
type RouteGetRulesConfigOK struct {
	Payload gmodels.NamespaceConfigResponse
}

func (o *RouteGetRulesConfigOK) Error() string {
	return fmt.Sprintf("[GET /api/ruler/{Recipient}/api/v1/rules][%d] routeGetRulesConfigOK  %+v", 200, o.Payload)
}
func (o *RouteGetRulesConfigOK) GetPayload() gmodels.NamespaceConfigResponse {
	return o.Payload
}

func (o *RouteGetRulesConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
