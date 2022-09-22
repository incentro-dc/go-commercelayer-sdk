/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// InventoryReturnLocationDataRelationshipsInventoryModel struct for InventoryReturnLocationDataRelationshipsInventoryModel
type InventoryReturnLocationDataRelationshipsInventoryModel struct {
	Data InventoryReturnLocationDataRelationshipsInventoryModelData `json:"data"`
}

// NewInventoryReturnLocationDataRelationshipsInventoryModel instantiates a new InventoryReturnLocationDataRelationshipsInventoryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryReturnLocationDataRelationshipsInventoryModel(data InventoryReturnLocationDataRelationshipsInventoryModelData) *InventoryReturnLocationDataRelationshipsInventoryModel {
	this := InventoryReturnLocationDataRelationshipsInventoryModel{}
	this.Data = data
	return &this
}

// NewInventoryReturnLocationDataRelationshipsInventoryModelWithDefaults instantiates a new InventoryReturnLocationDataRelationshipsInventoryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryReturnLocationDataRelationshipsInventoryModelWithDefaults() *InventoryReturnLocationDataRelationshipsInventoryModel {
	this := InventoryReturnLocationDataRelationshipsInventoryModel{}
	return &this
}

// GetData returns the Data field value
func (o *InventoryReturnLocationDataRelationshipsInventoryModel) GetData() InventoryReturnLocationDataRelationshipsInventoryModelData {
	if o == nil {
		var ret InventoryReturnLocationDataRelationshipsInventoryModelData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationDataRelationshipsInventoryModel) GetDataOk() (*InventoryReturnLocationDataRelationshipsInventoryModelData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InventoryReturnLocationDataRelationshipsInventoryModel) SetData(v InventoryReturnLocationDataRelationshipsInventoryModelData) {
	o.Data = v
}

func (o InventoryReturnLocationDataRelationshipsInventoryModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryReturnLocationDataRelationshipsInventoryModel struct {
	value *InventoryReturnLocationDataRelationshipsInventoryModel
	isSet bool
}

func (v NullableInventoryReturnLocationDataRelationshipsInventoryModel) Get() *InventoryReturnLocationDataRelationshipsInventoryModel {
	return v.value
}

func (v *NullableInventoryReturnLocationDataRelationshipsInventoryModel) Set(val *InventoryReturnLocationDataRelationshipsInventoryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryReturnLocationDataRelationshipsInventoryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryReturnLocationDataRelationshipsInventoryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryReturnLocationDataRelationshipsInventoryModel(val *InventoryReturnLocationDataRelationshipsInventoryModel) *NullableInventoryReturnLocationDataRelationshipsInventoryModel {
	return &NullableInventoryReturnLocationDataRelationshipsInventoryModel{value: val, isSet: true}
}

func (v NullableInventoryReturnLocationDataRelationshipsInventoryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryReturnLocationDataRelationshipsInventoryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}