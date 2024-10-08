// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-custom-variable/internal/sdk/models/shared"
	"net/http"
)

type UpdateCustomVariableRequest struct {
	// Custom variable ID
	ID             string                 `pathParam:"style=simple,explode=false,name=id"`
	CustomVariable *shared.CustomVariable `request:"mediaType=application/json"`
}

func (o *UpdateCustomVariableRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateCustomVariableRequest) GetCustomVariable() *shared.CustomVariable {
	if o == nil {
		return nil
	}
	return o.CustomVariable
}

type UpdateCustomVariableResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	CustomVariable *shared.CustomVariable
}

func (o *UpdateCustomVariableResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateCustomVariableResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateCustomVariableResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateCustomVariableResponse) GetCustomVariable() *shared.CustomVariable {
	if o == nil {
		return nil
	}
	return o.CustomVariable
}
