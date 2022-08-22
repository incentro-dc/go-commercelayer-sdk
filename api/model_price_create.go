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

// PriceCreate struct for PriceCreate
type PriceCreate struct {
	Data PriceCreateData `json:"data"`
}

// NewPriceCreate instantiates a new PriceCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceCreate(data PriceCreateData) *PriceCreate {
	this := PriceCreate{}
	this.Data = data
	return &this
}

// NewPriceCreateWithDefaults instantiates a new PriceCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceCreateWithDefaults() *PriceCreate {
	this := PriceCreate{}
	return &this
}

// GetData returns the Data field value
func (o *PriceCreate) GetData() PriceCreateData {
	if o == nil {
		var ret PriceCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PriceCreate) GetDataOk() (*PriceCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PriceCreate) SetData(v PriceCreateData) {
	o.Data = v
}

func (o PriceCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePriceCreate struct {
	value *PriceCreate
	isSet bool
}

func (v NullablePriceCreate) Get() *PriceCreate {
	return v.value
}

func (v *NullablePriceCreate) Set(val *PriceCreate) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceCreate) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceCreate(val *PriceCreate) *NullablePriceCreate {
	return &NullablePriceCreate{value: val, isSet: true}
}

func (v NullablePriceCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
