// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/epilot-dev/terraform-provider-epilot-variable/internal/sdk/internal/utils"
)

// ContextData - If context data is avaialble, this data will be used for variable replace.
type ContextData struct {
}

type VariableParameters struct {
	TemplateType TemplateType `json:"template_type"`
	// 2-letter language code (ISO 639-1)
	Language *string `default:"de" json:"language"`
	// The main entity ID. Use main entity in order to use the variable without schema slug prefix - or just pass directly to other object ID.
	MainEntityID *string `json:"main_entity_id,omitempty"`
	// Brand ID
	BrandID *float64 `json:"brand_id,omitempty"`
	// User ID
	UserID *string `json:"user_id,omitempty"`
	// Organization ID of the user
	UserOrgID *string `json:"user_org_id,omitempty"`
	// Custom variables with specified values form other services.
	CustomVariables []ExternalCustomVariable `json:"custom_variables,omitempty"`
	// If context data is avaialble, this data will be used for variable replace.
	ContextData *ContextData `json:"context_data,omitempty"`
	// The name of email template
	TemplateName *string `json:"template_name,omitempty"`
	// The tags of email template
	TemplateTags []string `json:"template_tags,omitempty"`
	// The version of the variables syntax supported. Default is 1.0
	VariablesVersion *string `json:"variables_version,omitempty"`
}

func (v VariableParameters) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *VariableParameters) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *VariableParameters) GetTemplateType() TemplateType {
	if o == nil {
		return TemplateType("")
	}
	return o.TemplateType
}

func (o *VariableParameters) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *VariableParameters) GetMainEntityID() *string {
	if o == nil {
		return nil
	}
	return o.MainEntityID
}

func (o *VariableParameters) GetBrandID() *float64 {
	if o == nil {
		return nil
	}
	return o.BrandID
}

func (o *VariableParameters) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *VariableParameters) GetUserOrgID() *string {
	if o == nil {
		return nil
	}
	return o.UserOrgID
}

func (o *VariableParameters) GetCustomVariables() []ExternalCustomVariable {
	if o == nil {
		return nil
	}
	return o.CustomVariables
}

func (o *VariableParameters) GetContextData() *ContextData {
	if o == nil {
		return nil
	}
	return o.ContextData
}

func (o *VariableParameters) GetTemplateName() *string {
	if o == nil {
		return nil
	}
	return o.TemplateName
}

func (o *VariableParameters) GetTemplateTags() []string {
	if o == nil {
		return nil
	}
	return o.TemplateTags
}

func (o *VariableParameters) GetVariablesVersion() *string {
	if o == nil {
		return nil
	}
	return o.VariablesVersion
}