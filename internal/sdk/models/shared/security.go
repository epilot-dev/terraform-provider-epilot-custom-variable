// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type Security struct {
	EpilotAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
	EpilotOrg  *string `security:"scheme,type=apiKey,subtype=header,name=x-epilot-org-id"`
}

func (o *Security) GetEpilotAuth() *string {
	if o == nil {
		return nil
	}
	return o.EpilotAuth
}

func (o *Security) GetEpilotOrg() *string {
	if o == nil {
		return nil
	}
	return o.EpilotOrg
}
