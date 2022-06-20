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

// GiftCardCreateData struct for GiftCardCreateData
type GiftCardCreateData struct {
	// The resource's type
	Type          string                           `json:"type"`
	Attributes    GiftCardCreateDataAttributes     `json:"attributes"`
	Relationships *GiftCardCreateDataRelationships `json:"relationships,omitempty"`
}

// NewGiftCardCreateData instantiates a new GiftCardCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardCreateData(type_ string, attributes GiftCardCreateDataAttributes) *GiftCardCreateData {
	this := GiftCardCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewGiftCardCreateDataWithDefaults instantiates a new GiftCardCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardCreateDataWithDefaults() *GiftCardCreateData {
	this := GiftCardCreateData{}
	var type_ string = "gift_cards"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *GiftCardCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GiftCardCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GiftCardCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *GiftCardCreateData) GetAttributes() GiftCardCreateDataAttributes {
	if o == nil {
		var ret GiftCardCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GiftCardCreateData) GetAttributesOk() (*GiftCardCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GiftCardCreateData) SetAttributes(v GiftCardCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GiftCardCreateData) GetRelationships() GiftCardCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret GiftCardCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardCreateData) GetRelationshipsOk() (*GiftCardCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GiftCardCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GiftCardCreateDataRelationships and assigns it to the Relationships field.
func (o *GiftCardCreateData) SetRelationships(v GiftCardCreateDataRelationships) {
	o.Relationships = &v
}

func (o GiftCardCreateData) MarshalJSON() ([]byte, error) {
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

type NullableGiftCardCreateData struct {
	value *GiftCardCreateData
	isSet bool
}

func (v NullableGiftCardCreateData) Get() *GiftCardCreateData {
	return v.value
}

func (v *NullableGiftCardCreateData) Set(val *GiftCardCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardCreateData(val *GiftCardCreateData) *NullableGiftCardCreateData {
	return &NullableGiftCardCreateData{value: val, isSet: true}
}

func (v NullableGiftCardCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}