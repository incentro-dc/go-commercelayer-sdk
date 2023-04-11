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

// AdjustmentUpdate struct for AdjustmentUpdate
type AdjustmentUpdate struct {
	Data AdjustmentUpdateData `json:"data"`
}

// NewAdjustmentUpdate instantiates a new AdjustmentUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdjustmentUpdate(data AdjustmentUpdateData) *AdjustmentUpdate {
	this := AdjustmentUpdate{}
	this.Data = data
	return &this
}

// NewAdjustmentUpdateWithDefaults instantiates a new AdjustmentUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdjustmentUpdateWithDefaults() *AdjustmentUpdate {
	this := AdjustmentUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *AdjustmentUpdate) GetData() AdjustmentUpdateData {
	if o == nil {
		var ret AdjustmentUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AdjustmentUpdate) GetDataOk() (*AdjustmentUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AdjustmentUpdate) SetData(v AdjustmentUpdateData) {
	o.Data = v
}

func (o AdjustmentUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAdjustmentUpdate struct {
	value *AdjustmentUpdate
	isSet bool
}

func (v NullableAdjustmentUpdate) Get() *AdjustmentUpdate {
	return v.value
}

func (v *NullableAdjustmentUpdate) Set(val *AdjustmentUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAdjustmentUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAdjustmentUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdjustmentUpdate(val *AdjustmentUpdate) *NullableAdjustmentUpdate {
	return &NullableAdjustmentUpdate{value: val, isSet: true}
}

func (v NullableAdjustmentUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdjustmentUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
