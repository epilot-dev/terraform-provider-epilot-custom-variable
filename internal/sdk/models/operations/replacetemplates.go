// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-variable/internal/sdk/models/shared"
	"net/http"
)

type ReplaceTemplatesRequestBody struct {
	Inputs     []string                   `json:"inputs,omitempty"`
	Parameters *shared.VariableParameters `json:"parameters,omitempty"`
}

func (o *ReplaceTemplatesRequestBody) GetInputs() []string {
	if o == nil {
		return nil
	}
	return o.Inputs
}

func (o *ReplaceTemplatesRequestBody) GetParameters() *shared.VariableParameters {
	if o == nil {
		return nil
	}
	return o.Parameters
}

// ReplaceTemplatesResponseBody - ok
type ReplaceTemplatesResponseBody struct {
	Outputs []string `json:"outputs,omitempty"`
}

func (o *ReplaceTemplatesResponseBody) GetOutputs() []string {
	if o == nil {
		return nil
	}
	return o.Outputs
}

type ReplaceTemplatesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// ok
	Object *ReplaceTemplatesResponseBody
}

func (o *ReplaceTemplatesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReplaceTemplatesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReplaceTemplatesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ReplaceTemplatesResponse) GetObject() *ReplaceTemplatesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}