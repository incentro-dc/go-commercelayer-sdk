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

// SkuListItemData struct for SkuListItemData
type SkuListItemData struct {
	// The resource's type
	Type          string                                        `json:"type"`
	Attributes    GETSkuListItems200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *SkuListItemDataRelationships                 `json:"relationships,omitempty"`
}

// NewSkuListItemData instantiates a new SkuListItemData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListItemData(type_ string, attributes GETSkuListItems200ResponseDataInnerAttributes) *SkuListItemData {
	this := SkuListItemData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewSkuListItemDataWithDefaults instantiates a new SkuListItemData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListItemDataWithDefaults() *SkuListItemData {
	this := SkuListItemData{}
	return &this
}

// GetType returns the Type field value
func (o *SkuListItemData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SkuListItemData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SkuListItemData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SkuListItemData) GetAttributes() GETSkuListItems200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETSkuListItems200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SkuListItemData) GetAttributesOk() (*GETSkuListItems200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SkuListItemData) SetAttributes(v GETSkuListItems200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SkuListItemData) GetRelationships() SkuListItemDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret SkuListItemDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListItemData) GetRelationshipsOk() (*SkuListItemDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SkuListItemData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SkuListItemDataRelationships and assigns it to the Relationships field.
func (o *SkuListItemData) SetRelationships(v SkuListItemDataRelationships) {
	o.Relationships = &v
}

func (o SkuListItemData) MarshalJSON() ([]byte, error) {
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

type NullableSkuListItemData struct {
	value *SkuListItemData
	isSet bool
}

func (v NullableSkuListItemData) Get() *SkuListItemData {
	return v.value
}

func (v *NullableSkuListItemData) Set(val *SkuListItemData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListItemData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListItemData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListItemData(val *SkuListItemData) *NullableSkuListItemData {
	return &NullableSkuListItemData{value: val, isSet: true}
}

func (v NullableSkuListItemData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListItemData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
