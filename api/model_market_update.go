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

// MarketUpdate struct for MarketUpdate
type MarketUpdate struct {
	Data MarketUpdateData `json:"data"`
}

// NewMarketUpdate instantiates a new MarketUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketUpdate(data MarketUpdateData) *MarketUpdate {
	this := MarketUpdate{}
	this.Data = data
	return &this
}

// NewMarketUpdateWithDefaults instantiates a new MarketUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketUpdateWithDefaults() *MarketUpdate {
	this := MarketUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *MarketUpdate) GetData() MarketUpdateData {
	if o == nil {
		var ret MarketUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MarketUpdate) GetDataOk() (*MarketUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MarketUpdate) SetData(v MarketUpdateData) {
	o.Data = v
}

func (o MarketUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableMarketUpdate struct {
	value *MarketUpdate
	isSet bool
}

func (v NullableMarketUpdate) Get() *MarketUpdate {
	return v.value
}

func (v *NullableMarketUpdate) Set(val *MarketUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketUpdate(val *MarketUpdate) *NullableMarketUpdate {
	return &NullableMarketUpdate{value: val, isSet: true}
}

func (v NullableMarketUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
