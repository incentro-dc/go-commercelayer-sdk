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

// LineItemCreateData struct for LineItemCreateData
type LineItemCreateData struct {
	// The resource's type
	Type          string                           `json:"type"`
	Attributes    LineItemCreateDataAttributes     `json:"attributes"`
	Relationships *LineItemCreateDataRelationships `json:"relationships,omitempty"`
}

// NewLineItemCreateData instantiates a new LineItemCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemCreateData(type_ string, attributes LineItemCreateDataAttributes) *LineItemCreateData {
	this := LineItemCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewLineItemCreateDataWithDefaults instantiates a new LineItemCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemCreateDataWithDefaults() *LineItemCreateData {
	this := LineItemCreateData{}
	var type_ string = "line_items"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *LineItemCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LineItemCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LineItemCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *LineItemCreateData) GetAttributes() LineItemCreateDataAttributes {
	if o == nil {
		var ret LineItemCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *LineItemCreateData) GetAttributesOk() (*LineItemCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *LineItemCreateData) SetAttributes(v LineItemCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *LineItemCreateData) GetRelationships() LineItemCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret LineItemCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemCreateData) GetRelationshipsOk() (*LineItemCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *LineItemCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given LineItemCreateDataRelationships and assigns it to the Relationships field.
func (o *LineItemCreateData) SetRelationships(v LineItemCreateDataRelationships) {
	o.Relationships = &v
}

func (o LineItemCreateData) MarshalJSON() ([]byte, error) {
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

type NullableLineItemCreateData struct {
	value *LineItemCreateData
	isSet bool
}

func (v NullableLineItemCreateData) Get() *LineItemCreateData {
	return v.value
}

func (v *NullableLineItemCreateData) Set(val *LineItemCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemCreateData(val *LineItemCreateData) *NullableLineItemCreateData {
	return &NullableLineItemCreateData{value: val, isSet: true}
}

func (v NullableLineItemCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
