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

// SkuListUpdateData struct for SkuListUpdateData
type SkuListUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                      `json:"id"`
	Attributes    SkuListUpdateDataAttributes `json:"attributes"`
	Relationships map[string]interface{}      `json:"relationships,omitempty"`
}

// NewSkuListUpdateData instantiates a new SkuListUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListUpdateData(type_ string, id string, attributes SkuListUpdateDataAttributes) *SkuListUpdateData {
	this := SkuListUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewSkuListUpdateDataWithDefaults instantiates a new SkuListUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListUpdateDataWithDefaults() *SkuListUpdateData {
	this := SkuListUpdateData{}
	var type_ string = "sku_lists"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *SkuListUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SkuListUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SkuListUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SkuListUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SkuListUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SkuListUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *SkuListUpdateData) GetAttributes() SkuListUpdateDataAttributes {
	if o == nil {
		var ret SkuListUpdateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SkuListUpdateData) GetAttributesOk() (*SkuListUpdateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SkuListUpdateData) SetAttributes(v SkuListUpdateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SkuListUpdateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListUpdateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SkuListUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *SkuListUpdateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o SkuListUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableSkuListUpdateData struct {
	value *SkuListUpdateData
	isSet bool
}

func (v NullableSkuListUpdateData) Get() *SkuListUpdateData {
	return v.value
}

func (v *NullableSkuListUpdateData) Set(val *SkuListUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListUpdateData(val *SkuListUpdateData) *NullableSkuListUpdateData {
	return &NullableSkuListUpdateData{value: val, isSet: true}
}

func (v NullableSkuListUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}