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

// InventoryModelUpdateData struct for InventoryModelUpdateData
type InventoryModelUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                                        `json:"id"`
	Attributes    PATCHInventoryModelsInventoryModelId200ResponseDataAttributes `json:"attributes"`
	Relationships map[string]interface{}                                        `json:"relationships,omitempty"`
}

// NewInventoryModelUpdateData instantiates a new InventoryModelUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryModelUpdateData(type_ string, id string, attributes PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) *InventoryModelUpdateData {
	this := InventoryModelUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewInventoryModelUpdateDataWithDefaults instantiates a new InventoryModelUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryModelUpdateDataWithDefaults() *InventoryModelUpdateData {
	this := InventoryModelUpdateData{}
	var type_ string = "inventory_models"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InventoryModelUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InventoryModelUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InventoryModelUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InventoryModelUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InventoryModelUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InventoryModelUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *InventoryModelUpdateData) GetAttributes() PATCHInventoryModelsInventoryModelId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHInventoryModelsInventoryModelId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *InventoryModelUpdateData) GetAttributesOk() (*PATCHInventoryModelsInventoryModelId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *InventoryModelUpdateData) SetAttributes(v PATCHInventoryModelsInventoryModelId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InventoryModelUpdateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryModelUpdateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InventoryModelUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *InventoryModelUpdateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o InventoryModelUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryModelUpdateData struct {
	value *InventoryModelUpdateData
	isSet bool
}

func (v NullableInventoryModelUpdateData) Get() *InventoryModelUpdateData {
	return v.value
}

func (v *NullableInventoryModelUpdateData) Set(val *InventoryModelUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryModelUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryModelUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryModelUpdateData(val *InventoryModelUpdateData) *NullableInventoryModelUpdateData {
	return &NullableInventoryModelUpdateData{value: val, isSet: true}
}

func (v NullableInventoryModelUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryModelUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
