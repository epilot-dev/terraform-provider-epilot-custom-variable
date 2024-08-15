// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-product/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"math/big"
	"time"
)

func (r *PriceDataSourceModel) RefreshFromSharedPrice(resp *shared.Price) {
	if resp != nil {
		r.ACL.Delete = []types.String{}
		for _, v := range resp.ACL.Delete {
			r.ACL.Delete = append(r.ACL.Delete, types.StringValue(v))
		}
		r.ACL.Edit = []types.String{}
		for _, v := range resp.ACL.Edit {
			r.ACL.Edit = append(r.ACL.Edit, types.StringValue(v))
		}
		r.ACL.View = []types.String{}
		for _, v := range resp.ACL.View {
			r.ACL.View = append(r.ACL.View, types.StringValue(v))
		}
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		r.ID = types.StringValue(resp.ID)
		r.Org = types.StringValue(resp.Org)
		r.Owners = []tfTypes.BaseEntityOwner{}
		if len(r.Owners) > len(resp.Owners) {
			r.Owners = r.Owners[:len(resp.Owners)]
		}
		for ownersCount, ownersItem := range resp.Owners {
			var owners1 tfTypes.BaseEntityOwner
			owners1.OrgID = types.StringValue(ownersItem.OrgID)
			owners1.UserID = types.StringPointerValue(ownersItem.UserID)
			if ownersCount+1 > len(r.Owners) {
				r.Owners = append(r.Owners, owners1)
			} else {
				r.Owners[ownersCount].OrgID = owners1.OrgID
				r.Owners[ownersCount].UserID = owners1.UserID
			}
		}
		r.Schema = types.StringValue(resp.Schema)
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.Title = types.StringValue(resp.Title)
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
		r.Active = types.BoolValue(resp.Active)
		if resp.BillingDurationAmount != nil {
			r.BillingDurationAmount = types.NumberValue(big.NewFloat(float64(*resp.BillingDurationAmount)))
		} else {
			r.BillingDurationAmount = types.NumberNull()
		}
		if resp.BillingDurationUnit != nil {
			r.BillingDurationUnit = types.StringValue(string(*resp.BillingDurationUnit))
		} else {
			r.BillingDurationUnit = types.StringNull()
		}
		r.Description = types.StringValue(resp.Description)
		r.IsCompositePrice = types.BoolPointerValue(resp.IsCompositePrice)
		r.IsTaxInclusive = types.BoolPointerValue(resp.IsTaxInclusive)
		r.LongDescription = types.StringPointerValue(resp.LongDescription)
		if resp.NoticeTimeAmount != nil {
			r.NoticeTimeAmount = types.NumberValue(big.NewFloat(float64(*resp.NoticeTimeAmount)))
		} else {
			r.NoticeTimeAmount = types.NumberNull()
		}
		if resp.NoticeTimeUnit != nil {
			r.NoticeTimeUnit = types.StringValue(string(*resp.NoticeTimeUnit))
		} else {
			r.NoticeTimeUnit = types.StringNull()
		}
		if resp.PriceComponents == nil {
			r.PriceComponents = nil
		} else {
			r.PriceComponents = &tfTypes.PriceCreatePriceComponents{}
			r.PriceComponents.DollarRelation = []tfTypes.PriceComponentRelation{}
			if len(r.PriceComponents.DollarRelation) > len(resp.PriceComponents.DollarRelation) {
				r.PriceComponents.DollarRelation = r.PriceComponents.DollarRelation[:len(resp.PriceComponents.DollarRelation)]
			}
			for dollarRelationCount, dollarRelationItem := range resp.PriceComponents.DollarRelation {
				var dollarRelation1 tfTypes.PriceComponentRelation
				dollarRelation1.Tags = []types.String{}
				for _, v := range dollarRelationItem.Tags {
					dollarRelation1.Tags = append(dollarRelation1.Tags, types.StringValue(v))
				}
				dollarRelation1.EntityID = types.StringPointerValue(dollarRelationItem.EntityID)
				if dollarRelationCount+1 > len(r.PriceComponents.DollarRelation) {
					r.PriceComponents.DollarRelation = append(r.PriceComponents.DollarRelation, dollarRelation1)
				} else {
					r.PriceComponents.DollarRelation[dollarRelationCount].Tags = dollarRelation1.Tags
					r.PriceComponents.DollarRelation[dollarRelationCount].EntityID = dollarRelation1.EntityID
				}
			}
		}
		if resp.PriceDisplayInJourneys != nil {
			r.PriceDisplayInJourneys = types.StringValue(string(*resp.PriceDisplayInJourneys))
		} else {
			r.PriceDisplayInJourneys = types.StringNull()
		}
		if resp.PricingModel != nil {
			r.PricingModel = types.StringValue(string(*resp.PricingModel))
		} else {
			r.PricingModel = types.StringNull()
		}
		if resp.RenewalDurationAmount != nil {
			r.RenewalDurationAmount = types.NumberValue(big.NewFloat(float64(*resp.RenewalDurationAmount)))
		} else {
			r.RenewalDurationAmount = types.NumberNull()
		}
		if resp.RenewalDurationUnit != nil {
			r.RenewalDurationUnit = types.StringValue(string(*resp.RenewalDurationUnit))
		} else {
			r.RenewalDurationUnit = types.StringNull()
		}
		if resp.Tax == nil {
			r.Tax = types.StringNull()
		} else {
			taxResult, _ := json.Marshal(resp.Tax)
			r.Tax = types.StringValue(string(taxResult))
		}
		if resp.TerminationTimeAmount != nil {
			r.TerminationTimeAmount = types.NumberValue(big.NewFloat(float64(*resp.TerminationTimeAmount)))
		} else {
			r.TerminationTimeAmount = types.NumberNull()
		}
		if resp.TerminationTimeUnit != nil {
			r.TerminationTimeUnit = types.StringValue(string(*resp.TerminationTimeUnit))
		} else {
			r.TerminationTimeUnit = types.StringNull()
		}
		r.Tiers = []tfTypes.PriceTier{}
		if len(r.Tiers) > len(resp.Tiers) {
			r.Tiers = r.Tiers[:len(resp.Tiers)]
		}
		for tiersCount, tiersItem := range resp.Tiers {
			var tiers1 tfTypes.PriceTier
			if tiersItem.DisplayMode != nil {
				tiers1.DisplayMode = types.StringValue(string(*tiersItem.DisplayMode))
			} else {
				tiers1.DisplayMode = types.StringNull()
			}
			if tiersItem.FlatFeeAmount != nil {
				tiers1.FlatFeeAmount = types.NumberValue(big.NewFloat(float64(*tiersItem.FlatFeeAmount)))
			} else {
				tiers1.FlatFeeAmount = types.NumberNull()
			}
			tiers1.FlatFeeAmountDecimal = types.StringPointerValue(tiersItem.FlatFeeAmountDecimal)
			if tiersItem.UnitAmount != nil {
				tiers1.UnitAmount = types.NumberValue(big.NewFloat(float64(*tiersItem.UnitAmount)))
			} else {
				tiers1.UnitAmount = types.NumberNull()
			}
			tiers1.UnitAmountDecimal = types.StringPointerValue(tiersItem.UnitAmountDecimal)
			if tiersItem.UpTo != nil {
				tiers1.UpTo = types.NumberValue(big.NewFloat(float64(*tiersItem.UpTo)))
			} else {
				tiers1.UpTo = types.NumberNull()
			}
			if tiersCount+1 > len(r.Tiers) {
				r.Tiers = append(r.Tiers, tiers1)
			} else {
				r.Tiers[tiersCount].DisplayMode = tiers1.DisplayMode
				r.Tiers[tiersCount].FlatFeeAmount = tiers1.FlatFeeAmount
				r.Tiers[tiersCount].FlatFeeAmountDecimal = tiers1.FlatFeeAmountDecimal
				r.Tiers[tiersCount].UnitAmount = tiers1.UnitAmount
				r.Tiers[tiersCount].UnitAmountDecimal = tiers1.UnitAmountDecimal
				r.Tiers[tiersCount].UpTo = tiers1.UpTo
			}
		}
		if resp.Type != nil {
			r.Type = types.StringValue(string(*resp.Type))
		} else {
			r.Type = types.StringNull()
		}
		r.Unit = types.StringPointerValue(resp.Unit)
		if resp.UnitAmount != nil {
			r.UnitAmount = types.NumberValue(big.NewFloat(float64(*resp.UnitAmount)))
		} else {
			r.UnitAmount = types.NumberNull()
		}
		r.UnitAmountCurrency = types.StringPointerValue(resp.UnitAmountCurrency)
		r.UnitAmountDecimal = types.StringPointerValue(resp.UnitAmountDecimal)
		r.VariablePrice = types.BoolPointerValue(resp.VariablePrice)
	}
}
