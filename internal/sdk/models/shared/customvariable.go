// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Type - Custom variable type
type Type string

const (
	TypeOrderTable  Type = "order_table"
	TypeCustom      Type = "custom"
	TypeJourneyLink Type = "journey_link"
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
	case "order_table":
		fallthrough
	case "custom":
		fallthrough
	case "journey_link":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type CustomVariable struct {
	// The tags of custom variable
	Tags   []string `json:"_tags,omitempty"`
	Config any      `json:"config,omitempty"`
	// Creation time
	CreatedAt *string `json:"created_at,omitempty"`
	// Created by
	CreatedBy *string `json:"created_by,omitempty"`
	// The helper function logic
	HelperLogic *string `json:"helper_logic,omitempty"`
	// The helper function parameter's names
	HelperParams []string `json:"helper_params,omitempty"`
	// ID
	ID *string `json:"id,omitempty"`
	// The key which is used for Handlebar variable syntax {{"{{"}}key}}
	Key string `json:"key"`
	// Custom variable name
	Name *string `json:"name,omitempty"`
	// Handlebar template that used to generate the variable content
	Template string `json:"template"`
	// Custom variable type
	Type *Type `json:"type,omitempty"`
	// Last update time
	UpdatedAt *string `json:"updated_at,omitempty"`
	// Updated by
	UpdatedBy *string `json:"updated_by,omitempty"`
}

func (o *CustomVariable) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CustomVariable) GetConfig() any {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CustomVariable) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CustomVariable) GetCreatedBy() *string {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *CustomVariable) GetHelperLogic() *string {
	if o == nil {
		return nil
	}
	return o.HelperLogic
}

func (o *CustomVariable) GetHelperParams() []string {
	if o == nil {
		return nil
	}
	return o.HelperParams
}

func (o *CustomVariable) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CustomVariable) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *CustomVariable) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CustomVariable) GetTemplate() string {
	if o == nil {
		return ""
	}
	return o.Template
}

func (o *CustomVariable) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *CustomVariable) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *CustomVariable) GetUpdatedBy() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedBy
}

type CustomVariableInput struct {
	// The tags of custom variable
	Tags   []string `json:"_tags,omitempty"`
	Config any      `json:"config,omitempty"`
	// The helper function logic
	HelperLogic *string `json:"helper_logic,omitempty"`
	// The helper function parameter's names
	HelperParams []string `json:"helper_params,omitempty"`
	// The key which is used for Handlebar variable syntax {{"{{"}}key}}
	Key string `json:"key"`
	// Custom variable name
	Name *string `json:"name,omitempty"`
	// Handlebar template that used to generate the variable content
	Template string `json:"template"`
	// Custom variable type
	Type *Type `json:"type,omitempty"`
}

func (o *CustomVariableInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CustomVariableInput) GetConfig() any {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CustomVariableInput) GetHelperLogic() *string {
	if o == nil {
		return nil
	}
	return o.HelperLogic
}

func (o *CustomVariableInput) GetHelperParams() []string {
	if o == nil {
		return nil
	}
	return o.HelperParams
}

func (o *CustomVariableInput) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *CustomVariableInput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CustomVariableInput) GetTemplate() string {
	if o == nil {
		return ""
	}
	return o.Template
}

func (o *CustomVariableInput) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}
