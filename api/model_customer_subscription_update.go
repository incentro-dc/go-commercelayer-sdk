/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CustomerSubscriptionUpdate struct for CustomerSubscriptionUpdate
type CustomerSubscriptionUpdate struct {
	Data CustomerSubscriptionUpdateData `json:"data"`
}

// NewCustomerSubscriptionUpdate instantiates a new CustomerSubscriptionUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerSubscriptionUpdate(data CustomerSubscriptionUpdateData) *CustomerSubscriptionUpdate {
	this := CustomerSubscriptionUpdate{}
	this.Data = data
	return &this
}

// NewCustomerSubscriptionUpdateWithDefaults instantiates a new CustomerSubscriptionUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerSubscriptionUpdateWithDefaults() *CustomerSubscriptionUpdate {
	this := CustomerSubscriptionUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *CustomerSubscriptionUpdate) GetData() CustomerSubscriptionUpdateData {
	if o == nil {
		var ret CustomerSubscriptionUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomerSubscriptionUpdate) GetDataOk() (*CustomerSubscriptionUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomerSubscriptionUpdate) SetData(v CustomerSubscriptionUpdateData) {
	o.Data = v
}

func (o CustomerSubscriptionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerSubscriptionUpdate struct {
	value *CustomerSubscriptionUpdate
	isSet bool
}

func (v NullableCustomerSubscriptionUpdate) Get() *CustomerSubscriptionUpdate {
	return v.value
}

func (v *NullableCustomerSubscriptionUpdate) Set(val *CustomerSubscriptionUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerSubscriptionUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerSubscriptionUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerSubscriptionUpdate(val *CustomerSubscriptionUpdate) *NullableCustomerSubscriptionUpdate {
	return &NullableCustomerSubscriptionUpdate{value: val, isSet: true}
}

func (v NullableCustomerSubscriptionUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerSubscriptionUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
