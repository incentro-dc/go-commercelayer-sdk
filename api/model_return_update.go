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

// ReturnUpdate struct for ReturnUpdate
type ReturnUpdate struct {
	Data ReturnUpdateData `json:"data"`
}

// NewReturnUpdate instantiates a new ReturnUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnUpdate(data ReturnUpdateData) *ReturnUpdate {
	this := ReturnUpdate{}
	this.Data = data
	return &this
}

// NewReturnUpdateWithDefaults instantiates a new ReturnUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnUpdateWithDefaults() *ReturnUpdate {
	this := ReturnUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *ReturnUpdate) GetData() ReturnUpdateData {
	if o == nil {
		var ret ReturnUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ReturnUpdate) GetDataOk() (*ReturnUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ReturnUpdate) SetData(v ReturnUpdateData) {
	o.Data = v
}

func (o ReturnUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableReturnUpdate struct {
	value *ReturnUpdate
	isSet bool
}

func (v NullableReturnUpdate) Get() *ReturnUpdate {
	return v.value
}

func (v *NullableReturnUpdate) Set(val *ReturnUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnUpdate(val *ReturnUpdate) *NullableReturnUpdate {
	return &NullableReturnUpdate{value: val, isSet: true}
}

func (v NullableReturnUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
