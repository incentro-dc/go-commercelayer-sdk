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

// checks if the InventoryReturnLocationCreateDataRelationshipsInventoryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryReturnLocationCreateDataRelationshipsInventoryModel{}

// InventoryReturnLocationCreateDataRelationshipsInventoryModel struct for InventoryReturnLocationCreateDataRelationshipsInventoryModel
type InventoryReturnLocationCreateDataRelationshipsInventoryModel struct {
	Data InventoryReturnLocationDataRelationshipsInventoryModelData `json:"data"`
}

// NewInventoryReturnLocationCreateDataRelationshipsInventoryModel instantiates a new InventoryReturnLocationCreateDataRelationshipsInventoryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryReturnLocationCreateDataRelationshipsInventoryModel(data InventoryReturnLocationDataRelationshipsInventoryModelData) *InventoryReturnLocationCreateDataRelationshipsInventoryModel {
	this := InventoryReturnLocationCreateDataRelationshipsInventoryModel{}
	this.Data = data
	return &this
}

// NewInventoryReturnLocationCreateDataRelationshipsInventoryModelWithDefaults instantiates a new InventoryReturnLocationCreateDataRelationshipsInventoryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryReturnLocationCreateDataRelationshipsInventoryModelWithDefaults() *InventoryReturnLocationCreateDataRelationshipsInventoryModel {
	this := InventoryReturnLocationCreateDataRelationshipsInventoryModel{}
	return &this
}

// GetData returns the Data field value
func (o *InventoryReturnLocationCreateDataRelationshipsInventoryModel) GetData() InventoryReturnLocationDataRelationshipsInventoryModelData {
	if o == nil {
		var ret InventoryReturnLocationDataRelationshipsInventoryModelData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationCreateDataRelationshipsInventoryModel) GetDataOk() (*InventoryReturnLocationDataRelationshipsInventoryModelData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InventoryReturnLocationCreateDataRelationshipsInventoryModel) SetData(v InventoryReturnLocationDataRelationshipsInventoryModelData) {
	o.Data = v
}

func (o InventoryReturnLocationCreateDataRelationshipsInventoryModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryReturnLocationCreateDataRelationshipsInventoryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableInventoryReturnLocationCreateDataRelationshipsInventoryModel struct {
	value *InventoryReturnLocationCreateDataRelationshipsInventoryModel
	isSet bool
}

func (v NullableInventoryReturnLocationCreateDataRelationshipsInventoryModel) Get() *InventoryReturnLocationCreateDataRelationshipsInventoryModel {
	return v.value
}

func (v *NullableInventoryReturnLocationCreateDataRelationshipsInventoryModel) Set(val *InventoryReturnLocationCreateDataRelationshipsInventoryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryReturnLocationCreateDataRelationshipsInventoryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryReturnLocationCreateDataRelationshipsInventoryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryReturnLocationCreateDataRelationshipsInventoryModel(val *InventoryReturnLocationCreateDataRelationshipsInventoryModel) *NullableInventoryReturnLocationCreateDataRelationshipsInventoryModel {
	return &NullableInventoryReturnLocationCreateDataRelationshipsInventoryModel{value: val, isSet: true}
}

func (v NullableInventoryReturnLocationCreateDataRelationshipsInventoryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryReturnLocationCreateDataRelationshipsInventoryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}