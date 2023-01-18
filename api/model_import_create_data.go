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

// ImportCreateData struct for ImportCreateData
type ImportCreateData struct {
	// The resource's type
	Type          string                               `json:"type"`
	Attributes    POSTImports201ResponseDataAttributes `json:"attributes"`
	Relationships map[string]interface{}               `json:"relationships,omitempty"`
}

// NewImportCreateData instantiates a new ImportCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportCreateData(type_ string, attributes POSTImports201ResponseDataAttributes) *ImportCreateData {
	this := ImportCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewImportCreateDataWithDefaults instantiates a new ImportCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportCreateDataWithDefaults() *ImportCreateData {
	this := ImportCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *ImportCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ImportCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ImportCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ImportCreateData) GetAttributes() POSTImports201ResponseDataAttributes {
	if o == nil {
		var ret POSTImports201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ImportCreateData) GetAttributesOk() (*POSTImports201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ImportCreateData) SetAttributes(v POSTImports201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ImportCreateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportCreateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ImportCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *ImportCreateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o ImportCreateData) MarshalJSON() ([]byte, error) {
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

type NullableImportCreateData struct {
	value *ImportCreateData
	isSet bool
}

func (v NullableImportCreateData) Get() *ImportCreateData {
	return v.value
}

func (v *NullableImportCreateData) Set(val *ImportCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableImportCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableImportCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportCreateData(val *ImportCreateData) *NullableImportCreateData {
	return &NullableImportCreateData{value: val, isSet: true}
}

func (v NullableImportCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
