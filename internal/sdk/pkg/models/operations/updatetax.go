// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateTaxRequest struct {
	// Tax to update (complete update)
	TaxCreate shared.TaxCreate `request:"mediaType=application/json"`
	// The tax id
	TaxID string `pathParam:"style=simple,explode=false,name=taxId"`
}

func (o *UpdateTaxRequest) GetTaxCreate() shared.TaxCreate {
	if o == nil {
		return shared.TaxCreate{}
	}
	return o.TaxCreate
}

func (o *UpdateTaxRequest) GetTaxID() string {
	if o == nil {
		return ""
	}
	return o.TaxID
}

type UpdateTaxResponse struct {
	// Any error based on client data errors
	ClientError *shared.ClientError
	// HTTP response content type for this operation
	ContentType string
	// Any error based on the server-side
	ServerError *shared.ServerError
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Tax entity response
	Tax *shared.Tax
}

func (o *UpdateTaxResponse) GetClientError() *shared.ClientError {
	if o == nil {
		return nil
	}
	return o.ClientError
}

func (o *UpdateTaxResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateTaxResponse) GetServerError() *shared.ServerError {
	if o == nil {
		return nil
	}
	return o.ServerError
}

func (o *UpdateTaxResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateTaxResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateTaxResponse) GetTax() *shared.Tax {
	if o == nil {
		return nil
	}
	return o.Tax
}
