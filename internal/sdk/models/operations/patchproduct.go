// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/models/shared"
	"net/http"
)

type PatchProductRequest struct {
	// Product to patch
	ProductPatch shared.ProductPatch `request:"mediaType=application/json"`
	// The product id
	ProductID string `pathParam:"style=simple,explode=false,name=productId"`
}

func (o *PatchProductRequest) GetProductPatch() shared.ProductPatch {
	if o == nil {
		return shared.ProductPatch{}
	}
	return o.ProductPatch
}

func (o *PatchProductRequest) GetProductID() string {
	if o == nil {
		return ""
	}
	return o.ProductID
}

type PatchProductResponse struct {
	// Any error based on client data errors
	ClientError *shared.ClientError
	// HTTP response content type for this operation
	ContentType string
	// Product entity response
	Product *shared.Product
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Any error based on the server-side
	ServerError *shared.ServerError
}

func (o *PatchProductResponse) GetClientError() *shared.ClientError {
	if o == nil {
		return nil
	}
	return o.ClientError
}

func (o *PatchProductResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchProductResponse) GetProduct() *shared.Product {
	if o == nil {
		return nil
	}
	return o.Product
}

func (o *PatchProductResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchProductResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchProductResponse) GetServerError() *shared.ServerError {
	if o == nil {
		return nil
	}
	return o.ServerError
}