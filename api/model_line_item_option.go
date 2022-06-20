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

// LineItemOption struct for LineItemOption
type LineItemOption struct {
	Data LineItemOptionData `json:"data"`
}

// NewLineItemOption instantiates a new LineItemOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemOption(data LineItemOptionData) *LineItemOption {
	this := LineItemOption{}
	this.Data = data
	return &this
}

// NewLineItemOptionWithDefaults instantiates a new LineItemOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemOptionWithDefaults() *LineItemOption {
	this := LineItemOption{}
	return &this
}

// GetData returns the Data field value
func (o *LineItemOption) GetData() LineItemOptionData {
	if o == nil {
		var ret LineItemOptionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *LineItemOption) GetDataOk() (*LineItemOptionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *LineItemOption) SetData(v LineItemOptionData) {
	o.Data = v
}

func (o LineItemOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableLineItemOption struct {
	value *LineItemOption
	isSet bool
}

func (v NullableLineItemOption) Get() *LineItemOption {
	return v.value
}

func (v *NullableLineItemOption) Set(val *LineItemOption) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemOption) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemOption(val *LineItemOption) *NullableLineItemOption {
	return &NullableLineItemOption{value: val, isSet: true}
}

func (v NullableLineItemOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
