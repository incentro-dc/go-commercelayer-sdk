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

// SkuListCreateData struct for SkuListCreateData
type SkuListCreateData struct {
	// The resource's type
	Type          interface{}                             `json:"type"`
	Attributes    POSTSkuLists201ResponseDataAttributes   `json:"attributes"`
	Relationships *CouponRecipientCreateDataRelationships `json:"relationships,omitempty"`
}

// NewSkuListCreateData instantiates a new SkuListCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListCreateData(type_ interface{}, attributes POSTSkuLists201ResponseDataAttributes) *SkuListCreateData {
	this := SkuListCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewSkuListCreateDataWithDefaults instantiates a new SkuListCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListCreateDataWithDefaults() *SkuListCreateData {
	this := SkuListCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SkuListCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SkuListCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SkuListCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SkuListCreateData) GetAttributes() POSTSkuLists201ResponseDataAttributes {
	if o == nil {
		var ret POSTSkuLists201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SkuListCreateData) GetAttributesOk() (*POSTSkuLists201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SkuListCreateData) SetAttributes(v POSTSkuLists201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SkuListCreateData) GetRelationships() CouponRecipientCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CouponRecipientCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListCreateData) GetRelationshipsOk() (*CouponRecipientCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SkuListCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CouponRecipientCreateDataRelationships and assigns it to the Relationships field.
func (o *SkuListCreateData) SetRelationships(v CouponRecipientCreateDataRelationships) {
	o.Relationships = &v
}

func (o SkuListCreateData) MarshalJSON() ([]byte, error) {
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

type NullableSkuListCreateData struct {
	value *SkuListCreateData
	isSet bool
}

func (v NullableSkuListCreateData) Get() *SkuListCreateData {
	return v.value
}

func (v *NullableSkuListCreateData) Set(val *SkuListCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListCreateData(val *SkuListCreateData) *NullableSkuListCreateData {
	return &NullableSkuListCreateData{value: val, isSet: true}
}

func (v NullableSkuListCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
