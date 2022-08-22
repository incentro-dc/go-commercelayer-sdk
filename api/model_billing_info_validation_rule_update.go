/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.7.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// BillingInfoValidationRuleUpdate struct for BillingInfoValidationRuleUpdate
type BillingInfoValidationRuleUpdate struct {
	Data BillingInfoValidationRuleUpdateData `json:"data"`
}

// NewBillingInfoValidationRuleUpdate instantiates a new BillingInfoValidationRuleUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInfoValidationRuleUpdate(data BillingInfoValidationRuleUpdateData) *BillingInfoValidationRuleUpdate {
	this := BillingInfoValidationRuleUpdate{}
	this.Data = data
	return &this
}

// NewBillingInfoValidationRuleUpdateWithDefaults instantiates a new BillingInfoValidationRuleUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInfoValidationRuleUpdateWithDefaults() *BillingInfoValidationRuleUpdate {
	this := BillingInfoValidationRuleUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *BillingInfoValidationRuleUpdate) GetData() BillingInfoValidationRuleUpdateData {
	if o == nil {
		var ret BillingInfoValidationRuleUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BillingInfoValidationRuleUpdate) GetDataOk() (*BillingInfoValidationRuleUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BillingInfoValidationRuleUpdate) SetData(v BillingInfoValidationRuleUpdateData) {
	o.Data = v
}

func (o BillingInfoValidationRuleUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBillingInfoValidationRuleUpdate struct {
	value *BillingInfoValidationRuleUpdate
	isSet bool
}

func (v NullableBillingInfoValidationRuleUpdate) Get() *BillingInfoValidationRuleUpdate {
	return v.value
}

func (v *NullableBillingInfoValidationRuleUpdate) Set(val *BillingInfoValidationRuleUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInfoValidationRuleUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInfoValidationRuleUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInfoValidationRuleUpdate(val *BillingInfoValidationRuleUpdate) *NullableBillingInfoValidationRuleUpdate {
	return &NullableBillingInfoValidationRuleUpdate{value: val, isSet: true}
}

func (v NullableBillingInfoValidationRuleUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInfoValidationRuleUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
