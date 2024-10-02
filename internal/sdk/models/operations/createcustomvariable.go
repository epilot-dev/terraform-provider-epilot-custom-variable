// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-custom-variable/internal/sdk/models/shared"
	"net/http"
)

type CreateCustomVariableResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	CustomVariable *shared.CustomVariable
}

func (o *CreateCustomVariableResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateCustomVariableResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateCustomVariableResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateCustomVariableResponse) GetCustomVariable() *shared.CustomVariable {
	if o == nil {
		return nil
	}
	return o.CustomVariable
}
