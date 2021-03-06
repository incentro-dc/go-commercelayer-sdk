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

// LineItemOptionUpdateData struct for LineItemOptionUpdateData
type LineItemOptionUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                 `json:"id"`
	Attributes    LineItemOptionUpdateDataAttributes     `json:"attributes"`
	Relationships *LineItemOptionUpdateDataRelationships `json:"relationships,omitempty"`
}

// NewLineItemOptionUpdateData instantiates a new LineItemOptionUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemOptionUpdateData(type_ string, id string, attributes LineItemOptionUpdateDataAttributes) *LineItemOptionUpdateData {
	this := LineItemOptionUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewLineItemOptionUpdateDataWithDefaults instantiates a new LineItemOptionUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemOptionUpdateDataWithDefaults() *LineItemOptionUpdateData {
	this := LineItemOptionUpdateData{}
	var type_ string = "line_item_options"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *LineItemOptionUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LineItemOptionUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LineItemOptionUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *LineItemOptionUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LineItemOptionUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LineItemOptionUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *LineItemOptionUpdateData) GetAttributes() LineItemOptionUpdateDataAttributes {
	if o == nil {
		var ret LineItemOptionUpdateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *LineItemOptionUpdateData) GetAttributesOk() (*LineItemOptionUpdateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *LineItemOptionUpdateData) SetAttributes(v LineItemOptionUpdateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *LineItemOptionUpdateData) GetRelationships() LineItemOptionUpdateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret LineItemOptionUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionUpdateData) GetRelationshipsOk() (*LineItemOptionUpdateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *LineItemOptionUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given LineItemOptionUpdateDataRelationships and assigns it to the Relationships field.
func (o *LineItemOptionUpdateData) SetRelationships(v LineItemOptionUpdateDataRelationships) {
	o.Relationships = &v
}

func (o LineItemOptionUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableLineItemOptionUpdateData struct {
	value *LineItemOptionUpdateData
	isSet bool
}

func (v NullableLineItemOptionUpdateData) Get() *LineItemOptionUpdateData {
	return v.value
}

func (v *NullableLineItemOptionUpdateData) Set(val *LineItemOptionUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemOptionUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemOptionUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemOptionUpdateData(val *LineItemOptionUpdateData) *NullableLineItemOptionUpdateData {
	return &NullableLineItemOptionUpdateData{value: val, isSet: true}
}

func (v NullableLineItemOptionUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemOptionUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
