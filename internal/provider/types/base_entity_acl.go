// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type BaseEntityACL struct {
	Delete []types.String `tfsdk:"delete"`
	Edit   []types.String `tfsdk:"edit"`
	View   []types.String `tfsdk:"view"`
}
