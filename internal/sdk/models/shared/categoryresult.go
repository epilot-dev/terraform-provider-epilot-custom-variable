// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CategoryResult struct {
	Category    *string `json:"category,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (o *CategoryResult) GetCategory() *string {
	if o == nil {
		return nil
	}
	return o.Category
}

func (o *CategoryResult) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}