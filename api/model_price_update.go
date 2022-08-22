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

// PriceUpdate struct for PriceUpdate
type PriceUpdate struct {
	Data PriceUpdateData `json:"data"`
}

// NewPriceUpdate instantiates a new PriceUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceUpdate(data PriceUpdateData) *PriceUpdate {
	this := PriceUpdate{}
	this.Data = data
	return &this
}

// NewPriceUpdateWithDefaults instantiates a new PriceUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceUpdateWithDefaults() *PriceUpdate {
	this := PriceUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *PriceUpdate) GetData() PriceUpdateData {
	if o == nil {
		var ret PriceUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PriceUpdate) GetDataOk() (*PriceUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PriceUpdate) SetData(v PriceUpdateData) {
	o.Data = v
}

func (o PriceUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePriceUpdate struct {
	value *PriceUpdate
	isSet bool
}

func (v NullablePriceUpdate) Get() *PriceUpdate {
	return v.value
}

func (v *NullablePriceUpdate) Set(val *PriceUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceUpdate(val *PriceUpdate) *NullablePriceUpdate {
	return &NullablePriceUpdate{value: val, isSet: true}
}

func (v NullablePriceUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
