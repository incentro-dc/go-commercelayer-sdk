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

// CustomerUpdate struct for CustomerUpdate
type CustomerUpdate struct {
	Data CustomerUpdateData `json:"data"`
}

// NewCustomerUpdate instantiates a new CustomerUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerUpdate(data CustomerUpdateData) *CustomerUpdate {
	this := CustomerUpdate{}
	this.Data = data
	return &this
}

// NewCustomerUpdateWithDefaults instantiates a new CustomerUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerUpdateWithDefaults() *CustomerUpdate {
	this := CustomerUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *CustomerUpdate) GetData() CustomerUpdateData {
	if o == nil {
		var ret CustomerUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomerUpdate) GetDataOk() (*CustomerUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomerUpdate) SetData(v CustomerUpdateData) {
	o.Data = v
}

func (o CustomerUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerUpdate struct {
	value *CustomerUpdate
	isSet bool
}

func (v NullableCustomerUpdate) Get() *CustomerUpdate {
	return v.value
}

func (v *NullableCustomerUpdate) Set(val *CustomerUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerUpdate(val *CustomerUpdate) *NullableCustomerUpdate {
	return &NullableCustomerUpdate{value: val, isSet: true}
}

func (v NullableCustomerUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}