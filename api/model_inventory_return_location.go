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

// InventoryReturnLocation struct for InventoryReturnLocation
type InventoryReturnLocation struct {
	Data *InventoryReturnLocationData `json:"data,omitempty"`
}

// NewInventoryReturnLocation instantiates a new InventoryReturnLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryReturnLocation() *InventoryReturnLocation {
	this := InventoryReturnLocation{}
	return &this
}

// NewInventoryReturnLocationWithDefaults instantiates a new InventoryReturnLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryReturnLocationWithDefaults() *InventoryReturnLocation {
	this := InventoryReturnLocation{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InventoryReturnLocation) GetData() InventoryReturnLocationData {
	if o == nil || o.Data == nil {
		var ret InventoryReturnLocationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocation) GetDataOk() (*InventoryReturnLocationData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InventoryReturnLocation) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given InventoryReturnLocationData and assigns it to the Data field.
func (o *InventoryReturnLocation) SetData(v InventoryReturnLocationData) {
	o.Data = &v
}

func (o InventoryReturnLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryReturnLocation struct {
	value *InventoryReturnLocation
	isSet bool
}

func (v NullableInventoryReturnLocation) Get() *InventoryReturnLocation {
	return v.value
}

func (v *NullableInventoryReturnLocation) Set(val *InventoryReturnLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryReturnLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryReturnLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryReturnLocation(val *InventoryReturnLocation) *NullableInventoryReturnLocation {
	return &NullableInventoryReturnLocation{value: val, isSet: true}
}

func (v NullableInventoryReturnLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryReturnLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
