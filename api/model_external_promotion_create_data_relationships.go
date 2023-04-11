/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ExternalPromotionCreateDataRelationships struct for ExternalPromotionCreateDataRelationships
type ExternalPromotionCreateDataRelationships struct {
	Market                   *BillingInfoValidationRuleCreateDataRelationshipsMarket           `json:"market,omitempty"`
	PromotionRules           *ExternalPromotionCreateDataRelationshipsPromotionRules           `json:"promotion_rules,omitempty"`
	OrderAmountPromotionRule *ExternalPromotionCreateDataRelationshipsOrderAmountPromotionRule `json:"order_amount_promotion_rule,omitempty"`
	SkuListPromotionRule     *ExternalPromotionCreateDataRelationshipsSkuListPromotionRule     `json:"sku_list_promotion_rule,omitempty"`
	CouponCodesPromotionRule *CouponCreateDataRelationshipsPromotionRule                       `json:"coupon_codes_promotion_rule,omitempty"`
}

// NewExternalPromotionCreateDataRelationships instantiates a new ExternalPromotionCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPromotionCreateDataRelationships() *ExternalPromotionCreateDataRelationships {
	this := ExternalPromotionCreateDataRelationships{}
	return &this
}

// NewExternalPromotionCreateDataRelationshipsWithDefaults instantiates a new ExternalPromotionCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPromotionCreateDataRelationshipsWithDefaults() *ExternalPromotionCreateDataRelationships {
	this := ExternalPromotionCreateDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *ExternalPromotionCreateDataRelationships) GetMarket() BillingInfoValidationRuleCreateDataRelationshipsMarket {
	if o == nil || o.Market == nil {
		var ret BillingInfoValidationRuleCreateDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionCreateDataRelationships) GetMarketOk() (*BillingInfoValidationRuleCreateDataRelationshipsMarket, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *ExternalPromotionCreateDataRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given BillingInfoValidationRuleCreateDataRelationshipsMarket and assigns it to the Market field.
func (o *ExternalPromotionCreateDataRelationships) SetMarket(v BillingInfoValidationRuleCreateDataRelationshipsMarket) {
	o.Market = &v
}

// GetPromotionRules returns the PromotionRules field value if set, zero value otherwise.
func (o *ExternalPromotionCreateDataRelationships) GetPromotionRules() ExternalPromotionCreateDataRelationshipsPromotionRules {
	if o == nil || o.PromotionRules == nil {
		var ret ExternalPromotionCreateDataRelationshipsPromotionRules
		return ret
	}
	return *o.PromotionRules
}

// GetPromotionRulesOk returns a tuple with the PromotionRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionCreateDataRelationships) GetPromotionRulesOk() (*ExternalPromotionCreateDataRelationshipsPromotionRules, bool) {
	if o == nil || o.PromotionRules == nil {
		return nil, false
	}
	return o.PromotionRules, true
}

// HasPromotionRules returns a boolean if a field has been set.
func (o *ExternalPromotionCreateDataRelationships) HasPromotionRules() bool {
	if o != nil && o.PromotionRules != nil {
		return true
	}

	return false
}

// SetPromotionRules gets a reference to the given ExternalPromotionCreateDataRelationshipsPromotionRules and assigns it to the PromotionRules field.
func (o *ExternalPromotionCreateDataRelationships) SetPromotionRules(v ExternalPromotionCreateDataRelationshipsPromotionRules) {
	o.PromotionRules = &v
}

// GetOrderAmountPromotionRule returns the OrderAmountPromotionRule field value if set, zero value otherwise.
func (o *ExternalPromotionCreateDataRelationships) GetOrderAmountPromotionRule() ExternalPromotionCreateDataRelationshipsOrderAmountPromotionRule {
	if o == nil || o.OrderAmountPromotionRule == nil {
		var ret ExternalPromotionCreateDataRelationshipsOrderAmountPromotionRule
		return ret
	}
	return *o.OrderAmountPromotionRule
}

// GetOrderAmountPromotionRuleOk returns a tuple with the OrderAmountPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionCreateDataRelationships) GetOrderAmountPromotionRuleOk() (*ExternalPromotionCreateDataRelationshipsOrderAmountPromotionRule, bool) {
	if o == nil || o.OrderAmountPromotionRule == nil {
		return nil, false
	}
	return o.OrderAmountPromotionRule, true
}

// HasOrderAmountPromotionRule returns a boolean if a field has been set.
func (o *ExternalPromotionCreateDataRelationships) HasOrderAmountPromotionRule() bool {
	if o != nil && o.OrderAmountPromotionRule != nil {
		return true
	}

	return false
}

// SetOrderAmountPromotionRule gets a reference to the given ExternalPromotionCreateDataRelationshipsOrderAmountPromotionRule and assigns it to the OrderAmountPromotionRule field.
func (o *ExternalPromotionCreateDataRelationships) SetOrderAmountPromotionRule(v ExternalPromotionCreateDataRelationshipsOrderAmountPromotionRule) {
	o.OrderAmountPromotionRule = &v
}

// GetSkuListPromotionRule returns the SkuListPromotionRule field value if set, zero value otherwise.
func (o *ExternalPromotionCreateDataRelationships) GetSkuListPromotionRule() ExternalPromotionCreateDataRelationshipsSkuListPromotionRule {
	if o == nil || o.SkuListPromotionRule == nil {
		var ret ExternalPromotionCreateDataRelationshipsSkuListPromotionRule
		return ret
	}
	return *o.SkuListPromotionRule
}

// GetSkuListPromotionRuleOk returns a tuple with the SkuListPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionCreateDataRelationships) GetSkuListPromotionRuleOk() (*ExternalPromotionCreateDataRelationshipsSkuListPromotionRule, bool) {
	if o == nil || o.SkuListPromotionRule == nil {
		return nil, false
	}
	return o.SkuListPromotionRule, true
}

// HasSkuListPromotionRule returns a boolean if a field has been set.
func (o *ExternalPromotionCreateDataRelationships) HasSkuListPromotionRule() bool {
	if o != nil && o.SkuListPromotionRule != nil {
		return true
	}

	return false
}

// SetSkuListPromotionRule gets a reference to the given ExternalPromotionCreateDataRelationshipsSkuListPromotionRule and assigns it to the SkuListPromotionRule field.
func (o *ExternalPromotionCreateDataRelationships) SetSkuListPromotionRule(v ExternalPromotionCreateDataRelationshipsSkuListPromotionRule) {
	o.SkuListPromotionRule = &v
}

// GetCouponCodesPromotionRule returns the CouponCodesPromotionRule field value if set, zero value otherwise.
func (o *ExternalPromotionCreateDataRelationships) GetCouponCodesPromotionRule() CouponCreateDataRelationshipsPromotionRule {
	if o == nil || o.CouponCodesPromotionRule == nil {
		var ret CouponCreateDataRelationshipsPromotionRule
		return ret
	}
	return *o.CouponCodesPromotionRule
}

// GetCouponCodesPromotionRuleOk returns a tuple with the CouponCodesPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionCreateDataRelationships) GetCouponCodesPromotionRuleOk() (*CouponCreateDataRelationshipsPromotionRule, bool) {
	if o == nil || o.CouponCodesPromotionRule == nil {
		return nil, false
	}
	return o.CouponCodesPromotionRule, true
}

// HasCouponCodesPromotionRule returns a boolean if a field has been set.
func (o *ExternalPromotionCreateDataRelationships) HasCouponCodesPromotionRule() bool {
	if o != nil && o.CouponCodesPromotionRule != nil {
		return true
	}

	return false
}

// SetCouponCodesPromotionRule gets a reference to the given CouponCreateDataRelationshipsPromotionRule and assigns it to the CouponCodesPromotionRule field.
func (o *ExternalPromotionCreateDataRelationships) SetCouponCodesPromotionRule(v CouponCreateDataRelationshipsPromotionRule) {
	o.CouponCodesPromotionRule = &v
}

func (o ExternalPromotionCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.PromotionRules != nil {
		toSerialize["promotion_rules"] = o.PromotionRules
	}
	if o.OrderAmountPromotionRule != nil {
		toSerialize["order_amount_promotion_rule"] = o.OrderAmountPromotionRule
	}
	if o.SkuListPromotionRule != nil {
		toSerialize["sku_list_promotion_rule"] = o.SkuListPromotionRule
	}
	if o.CouponCodesPromotionRule != nil {
		toSerialize["coupon_codes_promotion_rule"] = o.CouponCodesPromotionRule
	}
	return json.Marshal(toSerialize)
}

type NullableExternalPromotionCreateDataRelationships struct {
	value *ExternalPromotionCreateDataRelationships
	isSet bool
}

func (v NullableExternalPromotionCreateDataRelationships) Get() *ExternalPromotionCreateDataRelationships {
	return v.value
}

func (v *NullableExternalPromotionCreateDataRelationships) Set(val *ExternalPromotionCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPromotionCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPromotionCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPromotionCreateDataRelationships(val *ExternalPromotionCreateDataRelationships) *NullableExternalPromotionCreateDataRelationships {
	return &NullableExternalPromotionCreateDataRelationships{value: val, isSet: true}
}

func (v NullableExternalPromotionCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPromotionCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
