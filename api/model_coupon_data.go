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

// CouponData struct for CouponData
type CouponData struct {
	// The resource's type
	Type          string                   `json:"type"`
	Attributes    CouponDataAttributes     `json:"attributes"`
	Relationships *CouponDataRelationships `json:"relationships,omitempty"`
}

// NewCouponData instantiates a new CouponData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponData(type_ string, attributes CouponDataAttributes) *CouponData {
	this := CouponData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCouponDataWithDefaults instantiates a new CouponData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponDataWithDefaults() *CouponData {
	this := CouponData{}
	var type_ string = "coupons"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *CouponData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CouponData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CouponData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CouponData) GetAttributes() CouponDataAttributes {
	if o == nil {
		var ret CouponDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CouponData) GetAttributesOk() (*CouponDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CouponData) SetAttributes(v CouponDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CouponData) GetRelationships() CouponDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CouponDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponData) GetRelationshipsOk() (*CouponDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CouponData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CouponDataRelationships and assigns it to the Relationships field.
func (o *CouponData) SetRelationships(v CouponDataRelationships) {
	o.Relationships = &v
}

func (o CouponData) MarshalJSON() ([]byte, error) {
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

type NullableCouponData struct {
	value *CouponData
	isSet bool
}

func (v NullableCouponData) Get() *CouponData {
	return v.value
}

func (v *NullableCouponData) Set(val *CouponData) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponData) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponData(val *CouponData) *NullableCouponData {
	return &NullableCouponData{value: val, isSet: true}
}

func (v NullableCouponData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}