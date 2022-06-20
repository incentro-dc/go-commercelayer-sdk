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

// InventoryReturnLocationDataRelationshipsInventoryModel struct for InventoryReturnLocationDataRelationshipsInventoryModel
type InventoryReturnLocationDataRelationshipsInventoryModel struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewInventoryReturnLocationDataRelationshipsInventoryModel instantiates a new InventoryReturnLocationDataRelationshipsInventoryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryReturnLocationDataRelationshipsInventoryModel(type_ string, id string) *InventoryReturnLocationDataRelationshipsInventoryModel {
	this := InventoryReturnLocationDataRelationshipsInventoryModel{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewInventoryReturnLocationDataRelationshipsInventoryModelWithDefaults instantiates a new InventoryReturnLocationDataRelationshipsInventoryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryReturnLocationDataRelationshipsInventoryModelWithDefaults() *InventoryReturnLocationDataRelationshipsInventoryModel {
	this := InventoryReturnLocationDataRelationshipsInventoryModel{}
	var type_ string = "inventory_models"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InventoryReturnLocationDataRelationshipsInventoryModel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationDataRelationshipsInventoryModel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InventoryReturnLocationDataRelationshipsInventoryModel) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InventoryReturnLocationDataRelationshipsInventoryModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationDataRelationshipsInventoryModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InventoryReturnLocationDataRelationshipsInventoryModel) SetId(v string) {
	o.Id = v
}

func (o InventoryReturnLocationDataRelationshipsInventoryModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
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
