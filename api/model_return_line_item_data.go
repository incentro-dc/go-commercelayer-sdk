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

// ReturnLineItemData struct for ReturnLineItemData
type ReturnLineItemData struct {
	// The resource's type
	Type          string                           `json:"type"`
	Attributes    ReturnLineItemDataAttributes     `json:"attributes"`
	Relationships *ReturnLineItemDataRelationships `json:"relationships,omitempty"`
}

// NewReturnLineItemData instantiates a new ReturnLineItemData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnLineItemData(type_ string, attributes ReturnLineItemDataAttributes) *ReturnLineItemData {
	this := ReturnLineItemData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewReturnLineItemDataWithDefaults instantiates a new ReturnLineItemData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnLineItemDataWithDefaults() *ReturnLineItemData {
	this := ReturnLineItemData{}
	var type_ string = "return_line_items"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ReturnLineItemData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReturnLineItemData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReturnLineItemData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ReturnLineItemData) GetAttributes() ReturnLineItemDataAttributes {
	if o == nil {
		var ret ReturnLineItemDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ReturnLineItemData) GetAttributesOk() (*ReturnLineItemDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ReturnLineItemData) SetAttributes(v ReturnLineItemDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ReturnLineItemData) GetRelationships() ReturnLineItemDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ReturnLineItemDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLineItemData) GetRelationshipsOk() (*ReturnLineItemDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ReturnLineItemData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ReturnLineItemDataRelationships and assigns it to the Relationships field.
func (o *ReturnLineItemData) SetRelationships(v ReturnLineItemDataRelationships) {
	o.Relationships = &v
}

func (o ReturnLineItemData) MarshalJSON() ([]byte, error) {
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

type NullableReturnLineItemData struct {
	value *ReturnLineItemData
	isSet bool
}

func (v NullableReturnLineItemData) Get() *ReturnLineItemData {
	return v.value
}

func (v *NullableReturnLineItemData) Set(val *ReturnLineItemData) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnLineItemData) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnLineItemData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnLineItemData(val *ReturnLineItemData) *NullableReturnLineItemData {
	return &NullableReturnLineItemData{value: val, isSet: true}
}

func (v NullableReturnLineItemData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnLineItemData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}