// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/models/shared"
	"net/http"
)

type PatchTaxRequest struct {
	// Tax to patch
	TaxPatch shared.TaxPatch `request:"mediaType=application/json"`
	// The tax id
	TaxID string `pathParam:"style=simple,explode=false,name=taxId"`
}

func (o *PatchTaxRequest) GetTaxPatch() shared.TaxPatch {
	if o == nil {
		return shared.TaxPatch{}
	}
	return o.TaxPatch
}

func (o *PatchTaxRequest) GetTaxID() string {
	if o == nil {
		return ""
	}
	return o.TaxID
}

type PatchTaxResponse struct {
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

func (o *PatchTaxResponse) GetClientError() *shared.ClientError {
	if o == nil {
		return nil
	}
	return o.ClientError
}

func (o *PatchTaxResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchTaxResponse) GetServerError() *shared.ServerError {
	if o == nil {
		return nil
	}
	return o.ServerError
}

func (o *PatchTaxResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchTaxResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchTaxResponse) GetTax() *shared.Tax {
	if o == nil {
		return nil
	}
	return o.Tax
}