// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DollarRelation struct {
	EntityID *string `json:"entity_id,omitempty"`
}

func (o *DollarRelation) GetEntityID() *string {
	if o == nil {
		return nil
	}
	return o.EntityID
}

type BaseRelation struct {
	DollarRelation []DollarRelation `json:"$relation,omitempty"`
}

func (o *BaseRelation) GetDollarRelation() []DollarRelation {
	if o == nil {
		return nil
	}
	return o.DollarRelation
}