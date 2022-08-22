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

// CustomerGroup struct for CustomerGroup
type CustomerGroup struct {
	Data CustomerGroupData `json:"data"`
}

// NewCustomerGroup instantiates a new CustomerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerGroup(data CustomerGroupData) *CustomerGroup {
	this := CustomerGroup{}
	this.Data = data
	return &this
}

// NewCustomerGroupWithDefaults instantiates a new CustomerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerGroupWithDefaults() *CustomerGroup {
	this := CustomerGroup{}
	return &this
}

// GetData returns the Data field value
func (o *CustomerGroup) GetData() CustomerGroupData {
	if o == nil {
		var ret CustomerGroupData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomerGroup) GetDataOk() (*CustomerGroupData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomerGroup) SetData(v CustomerGroupData) {
	o.Data = v
}

func (o CustomerGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerGroup struct {
	value *CustomerGroup
	isSet bool
}

func (v NullableCustomerGroup) Get() *CustomerGroup {
	return v.value
}

func (v *NullableCustomerGroup) Set(val *CustomerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerGroup(val *CustomerGroup) *NullableCustomerGroup {
	return &NullableCustomerGroup{value: val, isSet: true}
}

func (v NullableCustomerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
