/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SkuCreateData struct for SkuCreateData
type SkuCreateData struct {
	// The resource's type
	Type          string                      `json:"type"`
	Attributes    SkuCreateDataAttributes     `json:"attributes"`
	Relationships *SkuCreateDataRelationships `json:"relationships,omitempty"`
}

// NewSkuCreateData instantiates a new SkuCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuCreateData(type_ string, attributes SkuCreateDataAttributes) *SkuCreateData {
	this := SkuCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewSkuCreateDataWithDefaults instantiates a new SkuCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuCreateDataWithDefaults() *SkuCreateData {
	this := SkuCreateData{}
	var type_ string = "skus"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *SkuCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SkuCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SkuCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SkuCreateData) GetAttributes() SkuCreateDataAttributes {
	if o == nil {
		var ret SkuCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SkuCreateData) GetAttributesOk() (*SkuCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SkuCreateData) SetAttributes(v SkuCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SkuCreateData) GetRelationships() SkuCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret SkuCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuCreateData) GetRelationshipsOk() (*SkuCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SkuCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SkuCreateDataRelationships and assigns it to the Relationships field.
func (o *SkuCreateData) SetRelationships(v SkuCreateDataRelationships) {
	o.Relationships = &v
}

func (o SkuCreateData) MarshalJSON() ([]byte, error) {
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

type NullableSkuCreateData struct {
	value *SkuCreateData
	isSet bool
}

func (v NullableSkuCreateData) Get() *SkuCreateData {
	return v.value
}

func (v *NullableSkuCreateData) Set(val *SkuCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuCreateData(val *SkuCreateData) *NullableSkuCreateData {
	return &NullableSkuCreateData{value: val, isSet: true}
}

func (v NullableSkuCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
