/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GiftCardData struct for GiftCardData
type GiftCardData struct {
	// The resource's type
	Type          interface{}                                `json:"type"`
	Attributes    GETGiftCards200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *GiftCardDataRelationships                 `json:"relationships,omitempty"`
}

// NewGiftCardData instantiates a new GiftCardData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardData(type_ interface{}, attributes GETGiftCards200ResponseDataInnerAttributes) *GiftCardData {
	this := GiftCardData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewGiftCardDataWithDefaults instantiates a new GiftCardData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardDataWithDefaults() *GiftCardData {
	this := GiftCardData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GiftCardData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GiftCardData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *GiftCardData) GetAttributes() GETGiftCards200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETGiftCards200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GiftCardData) GetAttributesOk() (*GETGiftCards200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GiftCardData) SetAttributes(v GETGiftCards200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GiftCardData) GetRelationships() GiftCardDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret GiftCardDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardData) GetRelationshipsOk() (*GiftCardDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GiftCardData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GiftCardDataRelationships and assigns it to the Relationships field.
func (o *GiftCardData) SetRelationships(v GiftCardDataRelationships) {
	o.Relationships = &v
}

func (o GiftCardData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
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

type NullableGiftCardData struct {
	value *GiftCardData
	isSet bool
}

func (v NullableGiftCardData) Get() *GiftCardData {
	return v.value
}

func (v *NullableGiftCardData) Set(val *GiftCardData) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardData) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardData(val *GiftCardData) *NullableGiftCardData {
	return &NullableGiftCardData{value: val, isSet: true}
}

func (v NullableGiftCardData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
