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

// InventoryStockLocation struct for InventoryStockLocation
type InventoryStockLocation struct {
	Data InventoryStockLocationData `json:"data"`
}

// NewInventoryStockLocation instantiates a new InventoryStockLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryStockLocation(data InventoryStockLocationData) *InventoryStockLocation {
	this := InventoryStockLocation{}
	this.Data = data
	return &this
}

// NewInventoryStockLocationWithDefaults instantiates a new InventoryStockLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryStockLocationWithDefaults() *InventoryStockLocation {
	this := InventoryStockLocation{}
	return &this
}

// GetData returns the Data field value
func (o *InventoryStockLocation) GetData() InventoryStockLocationData {
	if o == nil {
		var ret InventoryStockLocationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InventoryStockLocation) GetDataOk() (*InventoryStockLocationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InventoryStockLocation) SetData(v InventoryStockLocationData) {
	o.Data = v
}

func (o InventoryStockLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryStockLocation struct {
	value *InventoryStockLocation
	isSet bool
}

func (v NullableInventoryStockLocation) Get() *InventoryStockLocation {
	return v.value
}

func (v *NullableInventoryStockLocation) Set(val *InventoryStockLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryStockLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryStockLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryStockLocation(val *InventoryStockLocation) *NullableInventoryStockLocation {
	return &NullableInventoryStockLocation{value: val, isSet: true}
}

func (v NullableInventoryStockLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryStockLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}