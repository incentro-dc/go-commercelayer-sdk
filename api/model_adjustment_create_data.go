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

// AdjustmentCreateData struct for AdjustmentCreateData
type AdjustmentCreateData struct {
	// The resource's type
	Type          string                         `json:"type"`
	Attributes    AdjustmentCreateDataAttributes `json:"attributes"`
	Relationships map[string]interface{}         `json:"relationships,omitempty"`
}

// NewAdjustmentCreateData instantiates a new AdjustmentCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdjustmentCreateData(type_ string, attributes AdjustmentCreateDataAttributes) *AdjustmentCreateData {
	this := AdjustmentCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAdjustmentCreateDataWithDefaults instantiates a new AdjustmentCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdjustmentCreateDataWithDefaults() *AdjustmentCreateData {
	this := AdjustmentCreateData{}
	var type_ string = "adjustments"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *AdjustmentCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AdjustmentCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AdjustmentCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AdjustmentCreateData) GetAttributes() AdjustmentCreateDataAttributes {
	if o == nil {
		var ret AdjustmentCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AdjustmentCreateData) GetAttributesOk() (*AdjustmentCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AdjustmentCreateData) SetAttributes(v AdjustmentCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AdjustmentCreateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdjustmentCreateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AdjustmentCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *AdjustmentCreateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o AdjustmentCreateData) MarshalJSON() ([]byte, error) {
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

type NullableAdjustmentCreateData struct {
	value *AdjustmentCreateData
	isSet bool
}

func (v NullableAdjustmentCreateData) Get() *AdjustmentCreateData {
	return v.value
}

func (v *NullableAdjustmentCreateData) Set(val *AdjustmentCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAdjustmentCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAdjustmentCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdjustmentCreateData(val *AdjustmentCreateData) *NullableAdjustmentCreateData {
	return &NullableAdjustmentCreateData{value: val, isSet: true}
}

func (v NullableAdjustmentCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdjustmentCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
