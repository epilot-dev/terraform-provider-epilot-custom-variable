// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Type string

const (
	TypeSimple  Type = "simple"
	TypePartial Type = "partial"
)

func (e Type) ToPointer() *Type {
	return &e
}
func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "simple":
		fallthrough
	case "partial":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type VariableResult struct {
	Type *Type `json:"type,omitempty"`
	// Payload for the QR data
	Qrdata *string `json:"qrdata,omitempty"`
	// Variable group
	Group *string `json:"group,omitempty"`
	// The value which is used to insert to template
	Insert *string `json:"insert,omitempty"`
	// Variable description
	Description *string `json:"description,omitempty"`
}

func (o *VariableResult) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *VariableResult) GetQrdata() *string {
	if o == nil {
		return nil
	}
	return o.Qrdata
}

func (o *VariableResult) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *VariableResult) GetInsert() *string {
	if o == nil {
		return nil
	}
	return o.Insert
}

func (o *VariableResult) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}
