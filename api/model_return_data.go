/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.7.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ReturnData struct for ReturnData
type ReturnData struct {
	// The resource's type
	Type          string                                       `json:"type"`
	Attributes    GETReturns200ResponseDataInnerAttributes     `json:"attributes"`
	Relationships *GETReturns200ResponseDataInnerRelationships `json:"relationships,omitempty"`
}

// NewReturnData instantiates a new ReturnData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnData(type_ string, attributes GETReturns200ResponseDataInnerAttributes) *ReturnData {
	this := ReturnData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewReturnDataWithDefaults instantiates a new ReturnData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnDataWithDefaults() *ReturnData {
	this := ReturnData{}
	var type_ string = "returns"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ReturnData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReturnData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReturnData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ReturnData) GetAttributes() GETReturns200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETReturns200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ReturnData) GetAttributesOk() (*GETReturns200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ReturnData) SetAttributes(v GETReturns200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ReturnData) GetRelationships() GETReturns200ResponseDataInnerRelationships {
	if o == nil || o.Relationships == nil {
		var ret GETReturns200ResponseDataInnerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnData) GetRelationshipsOk() (*GETReturns200ResponseDataInnerRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ReturnData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GETReturns200ResponseDataInnerRelationships and assigns it to the Relationships field.
func (o *ReturnData) SetRelationships(v GETReturns200ResponseDataInnerRelationships) {
	o.Relationships = &v
}

func (o ReturnData) MarshalJSON() ([]byte, error) {
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

type NullableReturnData struct {
	value *ReturnData
	isSet bool
}

func (v NullableReturnData) Get() *ReturnData {
	return v.value
}

func (v *NullableReturnData) Set(val *ReturnData) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnData) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnData(val *ReturnData) *NullableReturnData {
	return &NullableReturnData{value: val, isSet: true}
}

func (v NullableReturnData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
