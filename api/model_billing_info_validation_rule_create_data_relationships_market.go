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

// BillingInfoValidationRuleCreateDataRelationshipsMarket struct for BillingInfoValidationRuleCreateDataRelationshipsMarket
type BillingInfoValidationRuleCreateDataRelationshipsMarket struct {
	Data AvalaraAccountDataRelationshipsMarketsData `json:"data"`
}

// NewBillingInfoValidationRuleCreateDataRelationshipsMarket instantiates a new BillingInfoValidationRuleCreateDataRelationshipsMarket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInfoValidationRuleCreateDataRelationshipsMarket(data AvalaraAccountDataRelationshipsMarketsData) *BillingInfoValidationRuleCreateDataRelationshipsMarket {
	this := BillingInfoValidationRuleCreateDataRelationshipsMarket{}
	this.Data = data
	return &this
}

// NewBillingInfoValidationRuleCreateDataRelationshipsMarketWithDefaults instantiates a new BillingInfoValidationRuleCreateDataRelationshipsMarket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInfoValidationRuleCreateDataRelationshipsMarketWithDefaults() *BillingInfoValidationRuleCreateDataRelationshipsMarket {
	this := BillingInfoValidationRuleCreateDataRelationshipsMarket{}
	return &this
}

// GetData returns the Data field value
func (o *BillingInfoValidationRuleCreateDataRelationshipsMarket) GetData() AvalaraAccountDataRelationshipsMarketsData {
	if o == nil {
		var ret AvalaraAccountDataRelationshipsMarketsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BillingInfoValidationRuleCreateDataRelationshipsMarket) GetDataOk() (*AvalaraAccountDataRelationshipsMarketsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BillingInfoValidationRuleCreateDataRelationshipsMarket) SetData(v AvalaraAccountDataRelationshipsMarketsData) {
	o.Data = v
}

func (o BillingInfoValidationRuleCreateDataRelationshipsMarket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBillingInfoValidationRuleCreateDataRelationshipsMarket struct {
	value *BillingInfoValidationRuleCreateDataRelationshipsMarket
	isSet bool
}

func (v NullableBillingInfoValidationRuleCreateDataRelationshipsMarket) Get() *BillingInfoValidationRuleCreateDataRelationshipsMarket {
	return v.value
}

func (v *NullableBillingInfoValidationRuleCreateDataRelationshipsMarket) Set(val *BillingInfoValidationRuleCreateDataRelationshipsMarket) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInfoValidationRuleCreateDataRelationshipsMarket) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInfoValidationRuleCreateDataRelationshipsMarket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInfoValidationRuleCreateDataRelationshipsMarket(val *BillingInfoValidationRuleCreateDataRelationshipsMarket) *NullableBillingInfoValidationRuleCreateDataRelationshipsMarket {
	return &NullableBillingInfoValidationRuleCreateDataRelationshipsMarket{value: val, isSet: true}
}

func (v NullableBillingInfoValidationRuleCreateDataRelationshipsMarket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInfoValidationRuleCreateDataRelationshipsMarket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
