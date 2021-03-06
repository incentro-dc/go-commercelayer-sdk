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

// RefundData struct for RefundData
type RefundData struct {
	// The resource's type
	Type          string                   `json:"type"`
	Attributes    RefundDataAttributes     `json:"attributes"`
	Relationships *RefundDataRelationships `json:"relationships,omitempty"`
}

// NewRefundData instantiates a new RefundData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefundData(type_ string, attributes RefundDataAttributes) *RefundData {
	this := RefundData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewRefundDataWithDefaults instantiates a new RefundData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundDataWithDefaults() *RefundData {
	this := RefundData{}
	var type_ string = "refunds"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *RefundData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RefundData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RefundData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *RefundData) GetAttributes() RefundDataAttributes {
	if o == nil {
		var ret RefundDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *RefundData) GetAttributesOk() (*RefundDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *RefundData) SetAttributes(v RefundDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *RefundData) GetRelationships() RefundDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret RefundDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundData) GetRelationshipsOk() (*RefundDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *RefundData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RefundDataRelationships and assigns it to the Relationships field.
func (o *RefundData) SetRelationships(v RefundDataRelationships) {
	o.Relationships = &v
}

func (o RefundData) MarshalJSON() ([]byte, error) {
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

type NullableRefundData struct {
	value *RefundData
	isSet bool
}

func (v NullableRefundData) Get() *RefundData {
	return v.value
}

func (v *NullableRefundData) Set(val *RefundData) {
	v.value = val
	v.isSet = true
}

func (v NullableRefundData) IsSet() bool {
	return v.isSet
}

func (v *NullableRefundData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefundData(val *RefundData) *NullableRefundData {
	return &NullableRefundData{value: val, isSet: true}
}

func (v NullableRefundData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefundData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
