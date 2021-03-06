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

// AddressUpdateData struct for AddressUpdateData
type AddressUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                      `json:"id"`
	Attributes    AddressUpdateDataAttributes `json:"attributes"`
	Relationships *AddressDataRelationships   `json:"relationships,omitempty"`
}

// NewAddressUpdateData instantiates a new AddressUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressUpdateData(type_ string, id string, attributes AddressUpdateDataAttributes) *AddressUpdateData {
	this := AddressUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewAddressUpdateDataWithDefaults instantiates a new AddressUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressUpdateDataWithDefaults() *AddressUpdateData {
	this := AddressUpdateData{}
	var type_ string = "addresses"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *AddressUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddressUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddressUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AddressUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AddressUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AddressUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *AddressUpdateData) GetAttributes() AddressUpdateDataAttributes {
	if o == nil {
		var ret AddressUpdateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AddressUpdateData) GetAttributesOk() (*AddressUpdateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AddressUpdateData) SetAttributes(v AddressUpdateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AddressUpdateData) GetRelationships() AddressDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AddressDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressUpdateData) GetRelationshipsOk() (*AddressDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AddressUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AddressDataRelationships and assigns it to the Relationships field.
func (o *AddressUpdateData) SetRelationships(v AddressDataRelationships) {
	o.Relationships = &v
}

func (o AddressUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableAddressUpdateData struct {
	value *AddressUpdateData
	isSet bool
}

func (v NullableAddressUpdateData) Get() *AddressUpdateData {
	return v.value
}

func (v *NullableAddressUpdateData) Set(val *AddressUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressUpdateData(val *AddressUpdateData) *NullableAddressUpdateData {
	return &NullableAddressUpdateData{value: val, isSet: true}
}

func (v NullableAddressUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
