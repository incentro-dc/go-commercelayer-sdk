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

// AuthorizationUpdate struct for AuthorizationUpdate
type AuthorizationUpdate struct {
	Data AuthorizationUpdateData `json:"data"`
}

// NewAuthorizationUpdate instantiates a new AuthorizationUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationUpdate(data AuthorizationUpdateData) *AuthorizationUpdate {
	this := AuthorizationUpdate{}
	this.Data = data
	return &this
}

// NewAuthorizationUpdateWithDefaults instantiates a new AuthorizationUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationUpdateWithDefaults() *AuthorizationUpdate {
	this := AuthorizationUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *AuthorizationUpdate) GetData() AuthorizationUpdateData {
	if o == nil {
		var ret AuthorizationUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AuthorizationUpdate) GetDataOk() (*AuthorizationUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AuthorizationUpdate) SetData(v AuthorizationUpdateData) {
	o.Data = v
}

func (o AuthorizationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizationUpdate struct {
	value *AuthorizationUpdate
	isSet bool
}

func (v NullableAuthorizationUpdate) Get() *AuthorizationUpdate {
	return v.value
}

func (v *NullableAuthorizationUpdate) Set(val *AuthorizationUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationUpdate(val *AuthorizationUpdate) *NullableAuthorizationUpdate {
	return &NullableAuthorizationUpdate{value: val, isSet: true}
}

func (v NullableAuthorizationUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
