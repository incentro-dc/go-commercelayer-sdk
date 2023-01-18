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

// CustomerAddressCreate struct for CustomerAddressCreate
type CustomerAddressCreate struct {
	Data CustomerAddressCreateData `json:"data"`
}

// NewCustomerAddressCreate instantiates a new CustomerAddressCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAddressCreate(data CustomerAddressCreateData) *CustomerAddressCreate {
	this := CustomerAddressCreate{}
	this.Data = data
	return &this
}

// NewCustomerAddressCreateWithDefaults instantiates a new CustomerAddressCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAddressCreateWithDefaults() *CustomerAddressCreate {
	this := CustomerAddressCreate{}
	return &this
}

// GetData returns the Data field value
func (o *CustomerAddressCreate) GetData() CustomerAddressCreateData {
	if o == nil {
		var ret CustomerAddressCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreate) GetDataOk() (*CustomerAddressCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomerAddressCreate) SetData(v CustomerAddressCreateData) {
	o.Data = v
}

func (o CustomerAddressCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerAddressCreate struct {
	value *CustomerAddressCreate
	isSet bool
}

func (v NullableCustomerAddressCreate) Get() *CustomerAddressCreate {
	return v.value
}

func (v *NullableCustomerAddressCreate) Set(val *CustomerAddressCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAddressCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAddressCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAddressCreate(val *CustomerAddressCreate) *NullableCustomerAddressCreate {
	return &NullableCustomerAddressCreate{value: val, isSet: true}
}

func (v NullableCustomerAddressCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAddressCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
