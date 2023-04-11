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

// AdjustmentCreate struct for AdjustmentCreate
type AdjustmentCreate struct {
	Data AdjustmentCreateData `json:"data"`
}

// NewAdjustmentCreate instantiates a new AdjustmentCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdjustmentCreate(data AdjustmentCreateData) *AdjustmentCreate {
	this := AdjustmentCreate{}
	this.Data = data
	return &this
}

// NewAdjustmentCreateWithDefaults instantiates a new AdjustmentCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdjustmentCreateWithDefaults() *AdjustmentCreate {
	this := AdjustmentCreate{}
	return &this
}

// GetData returns the Data field value
func (o *AdjustmentCreate) GetData() AdjustmentCreateData {
	if o == nil {
		var ret AdjustmentCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AdjustmentCreate) GetDataOk() (*AdjustmentCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AdjustmentCreate) SetData(v AdjustmentCreateData) {
	o.Data = v
}

func (o AdjustmentCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAdjustmentCreate struct {
	value *AdjustmentCreate
	isSet bool
}

func (v NullableAdjustmentCreate) Get() *AdjustmentCreate {
	return v.value
}

func (v *NullableAdjustmentCreate) Set(val *AdjustmentCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAdjustmentCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAdjustmentCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdjustmentCreate(val *AdjustmentCreate) *NullableAdjustmentCreate {
	return &NullableAdjustmentCreate{value: val, isSet: true}
}

func (v NullableAdjustmentCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdjustmentCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
