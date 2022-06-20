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

// AddressCreateData struct for AddressCreateData
type AddressCreateData struct {
	// The resource's type
	Type          string                      `json:"type"`
	Attributes    AddressCreateDataAttributes `json:"attributes"`
	Relationships *AddressDataRelationships   `json:"relationships,omitempty"`
}

// NewAddressCreateData instantiates a new AddressCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressCreateData(type_ string, attributes AddressCreateDataAttributes) *AddressCreateData {
	this := AddressCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAddressCreateDataWithDefaults instantiates a new AddressCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressCreateDataWithDefaults() *AddressCreateData {
	this := AddressCreateData{}
	var type_ string = "addresses"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *AddressCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddressCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddressCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AddressCreateData) GetAttributes() AddressCreateDataAttributes {
	if o == nil {
		var ret AddressCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AddressCreateData) GetAttributesOk() (*AddressCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AddressCreateData) SetAttributes(v AddressCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AddressCreateData) GetRelationships() AddressDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AddressDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressCreateData) GetRelationshipsOk() (*AddressDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AddressCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AddressDataRelationships and assigns it to the Relationships field.
func (o *AddressCreateData) SetRelationships(v AddressDataRelationships) {
	o.Relationships = &v
}

func (o AddressCreateData) MarshalJSON() ([]byte, error) {
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

type NullableAddressCreateData struct {
	value *AddressCreateData
	isSet bool
}

func (v NullableAddressCreateData) Get() *AddressCreateData {
	return v.value
}

func (v *NullableAddressCreateData) Set(val *AddressCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressCreateData(val *AddressCreateData) *NullableAddressCreateData {
	return &NullableAddressCreateData{value: val, isSet: true}
}

func (v NullableAddressCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
