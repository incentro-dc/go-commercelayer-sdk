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

// InventoryModelData struct for InventoryModelData
type InventoryModelData struct {
	// The resource's type
	Type          string                                               `json:"type"`
	Attributes    GETInventoryModels200ResponseDataInnerAttributes     `json:"attributes"`
	Relationships *GETInventoryModels200ResponseDataInnerRelationships `json:"relationships,omitempty"`
}

// NewInventoryModelData instantiates a new InventoryModelData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryModelData(type_ string, attributes GETInventoryModels200ResponseDataInnerAttributes) *InventoryModelData {
	this := InventoryModelData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewInventoryModelDataWithDefaults instantiates a new InventoryModelData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryModelDataWithDefaults() *InventoryModelData {
	this := InventoryModelData{}
	var type_ string = "inventory_models"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InventoryModelData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InventoryModelData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InventoryModelData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *InventoryModelData) GetAttributes() GETInventoryModels200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETInventoryModels200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *InventoryModelData) GetAttributesOk() (*GETInventoryModels200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *InventoryModelData) SetAttributes(v GETInventoryModels200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InventoryModelData) GetRelationships() GETInventoryModels200ResponseDataInnerRelationships {
	if o == nil || o.Relationships == nil {
		var ret GETInventoryModels200ResponseDataInnerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryModelData) GetRelationshipsOk() (*GETInventoryModels200ResponseDataInnerRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InventoryModelData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GETInventoryModels200ResponseDataInnerRelationships and assigns it to the Relationships field.
func (o *InventoryModelData) SetRelationships(v GETInventoryModels200ResponseDataInnerRelationships) {
	o.Relationships = &v
}

func (o InventoryModelData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryModelData struct {
	value *InventoryModelData
	isSet bool
}

func (v NullableInventoryModelData) Get() *InventoryModelData {
	return v.value
}

func (v *NullableInventoryModelData) Set(val *InventoryModelData) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryModelData) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryModelData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryModelData(val *InventoryModelData) *NullableInventoryModelData {
	return &NullableInventoryModelData{value: val, isSet: true}
}

func (v NullableInventoryModelData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryModelData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
