// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/epilot-dev/terraform-provider-epilot-custom-variable/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *CustomVariableDataSourceModel) RefreshFromSharedCustomVariable(resp *shared.CustomVariable) {
	if resp != nil {
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		if resp.Config == nil {
			r.Config = types.StringNull()
		} else {
			configResult, _ := json.Marshal(resp.Config)
			r.Config = types.StringValue(string(configResult))
		}
		r.CreatedAt = types.StringPointerValue(resp.CreatedAt)
		r.CreatedBy = types.StringPointerValue(resp.CreatedBy)
		r.HelperLogic = types.StringPointerValue(resp.HelperLogic)
		r.HelperParams = []types.String{}
		for _, v := range resp.HelperParams {
			r.HelperParams = append(r.HelperParams, types.StringValue(v))
		}
		r.ID = types.StringPointerValue(resp.ID)
		r.Key = types.StringPointerValue(resp.Key)
		r.Name = types.StringPointerValue(resp.Name)
		r.Template = types.StringPointerValue(resp.Template)
		if resp.Type != nil {
			r.Type = types.StringValue(string(*resp.Type))
		} else {
			r.Type = types.StringNull()
		}
		r.UpdatedAt = types.StringPointerValue(resp.UpdatedAt)
		r.UpdatedBy = types.StringPointerValue(resp.UpdatedBy)
	}
}
