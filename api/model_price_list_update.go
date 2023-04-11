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

// PriceListUpdate struct for PriceListUpdate
type PriceListUpdate struct {
	Data PriceListUpdateData `json:"data"`
}

// NewPriceListUpdate instantiates a new PriceListUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceListUpdate(data PriceListUpdateData) *PriceListUpdate {
	this := PriceListUpdate{}
	this.Data = data
	return &this
}

// NewPriceListUpdateWithDefaults instantiates a new PriceListUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceListUpdateWithDefaults() *PriceListUpdate {
	this := PriceListUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *PriceListUpdate) GetData() PriceListUpdateData {
	if o == nil {
		var ret PriceListUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PriceListUpdate) GetDataOk() (*PriceListUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PriceListUpdate) SetData(v PriceListUpdateData) {
	o.Data = v
}

func (o PriceListUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePriceListUpdate struct {
	value *PriceListUpdate
	isSet bool
}

func (v NullablePriceListUpdate) Get() *PriceListUpdate {
	return v.value
}

func (v *NullablePriceListUpdate) Set(val *PriceListUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceListUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceListUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceListUpdate(val *PriceListUpdate) *NullablePriceListUpdate {
	return &NullablePriceListUpdate{value: val, isSet: true}
}

func (v NullablePriceListUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceListUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
