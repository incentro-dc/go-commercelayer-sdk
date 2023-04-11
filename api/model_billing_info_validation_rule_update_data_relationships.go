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

// BillingInfoValidationRuleUpdateDataRelationships struct for BillingInfoValidationRuleUpdateDataRelationships
type BillingInfoValidationRuleUpdateDataRelationships struct {
	Market *BillingInfoValidationRuleCreateDataRelationshipsMarket `json:"market,omitempty"`
}

// NewBillingInfoValidationRuleUpdateDataRelationships instantiates a new BillingInfoValidationRuleUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInfoValidationRuleUpdateDataRelationships() *BillingInfoValidationRuleUpdateDataRelationships {
	this := BillingInfoValidationRuleUpdateDataRelationships{}
	return &this
}

// NewBillingInfoValidationRuleUpdateDataRelationshipsWithDefaults instantiates a new BillingInfoValidationRuleUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInfoValidationRuleUpdateDataRelationshipsWithDefaults() *BillingInfoValidationRuleUpdateDataRelationships {
	this := BillingInfoValidationRuleUpdateDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *BillingInfoValidationRuleUpdateDataRelationships) GetMarket() BillingInfoValidationRuleCreateDataRelationshipsMarket {
	if o == nil || o.Market == nil {
		var ret BillingInfoValidationRuleCreateDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInfoValidationRuleUpdateDataRelationships) GetMarketOk() (*BillingInfoValidationRuleCreateDataRelationshipsMarket, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *BillingInfoValidationRuleUpdateDataRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given BillingInfoValidationRuleCreateDataRelationshipsMarket and assigns it to the Market field.
func (o *BillingInfoValidationRuleUpdateDataRelationships) SetMarket(v BillingInfoValidationRuleCreateDataRelationshipsMarket) {
	o.Market = &v
}

func (o BillingInfoValidationRuleUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	return json.Marshal(toSerialize)
}

type NullableBillingInfoValidationRuleUpdateDataRelationships struct {
	value *BillingInfoValidationRuleUpdateDataRelationships
	isSet bool
}

func (v NullableBillingInfoValidationRuleUpdateDataRelationships) Get() *BillingInfoValidationRuleUpdateDataRelationships {
	return v.value
}

func (v *NullableBillingInfoValidationRuleUpdateDataRelationships) Set(val *BillingInfoValidationRuleUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInfoValidationRuleUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInfoValidationRuleUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInfoValidationRuleUpdateDataRelationships(val *BillingInfoValidationRuleUpdateDataRelationships) *NullableBillingInfoValidationRuleUpdateDataRelationships {
	return &NullableBillingInfoValidationRuleUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableBillingInfoValidationRuleUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInfoValidationRuleUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
