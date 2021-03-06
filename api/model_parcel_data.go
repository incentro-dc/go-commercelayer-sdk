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

// ParcelData struct for ParcelData
type ParcelData struct {
	// The resource's type
	Type          string                   `json:"type"`
	Attributes    ParcelDataAttributes     `json:"attributes"`
	Relationships *ParcelDataRelationships `json:"relationships,omitempty"`
}

// NewParcelData instantiates a new ParcelData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelData(type_ string, attributes ParcelDataAttributes) *ParcelData {
	this := ParcelData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewParcelDataWithDefaults instantiates a new ParcelData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelDataWithDefaults() *ParcelData {
	this := ParcelData{}
	var type_ string = "parcels"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ParcelData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ParcelData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ParcelData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ParcelData) GetAttributes() ParcelDataAttributes {
	if o == nil {
		var ret ParcelDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ParcelData) GetAttributesOk() (*ParcelDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ParcelData) SetAttributes(v ParcelDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ParcelData) GetRelationships() ParcelDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ParcelDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelData) GetRelationshipsOk() (*ParcelDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ParcelData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ParcelDataRelationships and assigns it to the Relationships field.
func (o *ParcelData) SetRelationships(v ParcelDataRelationships) {
	o.Relationships = &v
}

func (o ParcelData) MarshalJSON() ([]byte, error) {
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

type NullableParcelData struct {
	value *ParcelData
	isSet bool
}

func (v NullableParcelData) Get() *ParcelData {
	return v.value
}

func (v *NullableParcelData) Set(val *ParcelData) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelData) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelData(val *ParcelData) *NullableParcelData {
	return &NullableParcelData{value: val, isSet: true}
}

func (v NullableParcelData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
