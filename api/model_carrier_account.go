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

// CarrierAccount struct for CarrierAccount
type CarrierAccount struct {
	Data *CarrierAccountData `json:"data,omitempty"`
}

// NewCarrierAccount instantiates a new CarrierAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCarrierAccount() *CarrierAccount {
	this := CarrierAccount{}
	return &this
}

// NewCarrierAccountWithDefaults instantiates a new CarrierAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCarrierAccountWithDefaults() *CarrierAccount {
	this := CarrierAccount{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CarrierAccount) GetData() CarrierAccountData {
	if o == nil || o.Data == nil {
		var ret CarrierAccountData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CarrierAccount) GetDataOk() (*CarrierAccountData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CarrierAccount) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given CarrierAccountData and assigns it to the Data field.
func (o *CarrierAccount) SetData(v CarrierAccountData) {
	o.Data = &v
}

func (o CarrierAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCarrierAccount struct {
	value *CarrierAccount
	isSet bool
}

func (v NullableCarrierAccount) Get() *CarrierAccount {
	return v.value
}

func (v *NullableCarrierAccount) Set(val *CarrierAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableCarrierAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableCarrierAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarrierAccount(val *CarrierAccount) *NullableCarrierAccount {
	return &NullableCarrierAccount{value: val, isSet: true}
}

func (v NullableCarrierAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCarrierAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
