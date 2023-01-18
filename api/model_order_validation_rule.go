/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// OrderValidationRule struct for OrderValidationRule
type OrderValidationRule struct {
	Data *OrderValidationRuleData `json:"data,omitempty"`
}

// NewOrderValidationRule instantiates a new OrderValidationRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderValidationRule() *OrderValidationRule {
	this := OrderValidationRule{}
	return &this
}

// NewOrderValidationRuleWithDefaults instantiates a new OrderValidationRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderValidationRuleWithDefaults() *OrderValidationRule {
	this := OrderValidationRule{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrderValidationRule) GetData() OrderValidationRuleData {
	if o == nil || o.Data == nil {
		var ret OrderValidationRuleData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderValidationRule) GetDataOk() (*OrderValidationRuleData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrderValidationRule) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given OrderValidationRuleData and assigns it to the Data field.
func (o *OrderValidationRule) SetData(v OrderValidationRuleData) {
	o.Data = &v
}

func (o OrderValidationRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableOrderValidationRule struct {
	value *OrderValidationRule
	isSet bool
}

func (v NullableOrderValidationRule) Get() *OrderValidationRule {
	return v.value
}

func (v *NullableOrderValidationRule) Set(val *OrderValidationRule) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderValidationRule) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderValidationRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderValidationRule(val *OrderValidationRule) *NullableOrderValidationRule {
	return &NullableOrderValidationRule{value: val, isSet: true}
}

func (v NullableOrderValidationRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderValidationRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
