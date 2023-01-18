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

// ExportCreateData struct for ExportCreateData
type ExportCreateData struct {
	// The resource's type
	Type          string                               `json:"type"`
	Attributes    POSTExports201ResponseDataAttributes `json:"attributes"`
	Relationships map[string]interface{}               `json:"relationships,omitempty"`
}

// NewExportCreateData instantiates a new ExportCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportCreateData(type_ string, attributes POSTExports201ResponseDataAttributes) *ExportCreateData {
	this := ExportCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewExportCreateDataWithDefaults instantiates a new ExportCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportCreateDataWithDefaults() *ExportCreateData {
	this := ExportCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *ExportCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExportCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExportCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ExportCreateData) GetAttributes() POSTExports201ResponseDataAttributes {
	if o == nil {
		var ret POSTExports201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ExportCreateData) GetAttributesOk() (*POSTExports201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ExportCreateData) SetAttributes(v POSTExports201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ExportCreateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportCreateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ExportCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *ExportCreateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o ExportCreateData) MarshalJSON() ([]byte, error) {
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

type NullableExportCreateData struct {
	value *ExportCreateData
	isSet bool
}

func (v NullableExportCreateData) Get() *ExportCreateData {
	return v.value
}

func (v *NullableExportCreateData) Set(val *ExportCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableExportCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableExportCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportCreateData(val *ExportCreateData) *NullableExportCreateData {
	return &NullableExportCreateData{value: val, isSet: true}
}

func (v NullableExportCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
