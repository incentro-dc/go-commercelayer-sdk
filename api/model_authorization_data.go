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

// AuthorizationData struct for AuthorizationData
type AuthorizationData struct {
	// The resource's type
	Type          string                          `json:"type"`
	Attributes    AuthorizationDataAttributes     `json:"attributes"`
	Relationships *AuthorizationDataRelationships `json:"relationships,omitempty"`
}

// NewAuthorizationData instantiates a new AuthorizationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationData(type_ string, attributes AuthorizationDataAttributes) *AuthorizationData {
	this := AuthorizationData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAuthorizationDataWithDefaults instantiates a new AuthorizationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationDataWithDefaults() *AuthorizationData {
	this := AuthorizationData{}
	var type_ string = "authorizations"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *AuthorizationData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuthorizationData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuthorizationData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AuthorizationData) GetAttributes() AuthorizationDataAttributes {
	if o == nil {
		var ret AuthorizationDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AuthorizationData) GetAttributesOk() (*AuthorizationDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AuthorizationData) SetAttributes(v AuthorizationDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AuthorizationData) GetRelationships() AuthorizationDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AuthorizationDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationData) GetRelationshipsOk() (*AuthorizationDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AuthorizationData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AuthorizationDataRelationships and assigns it to the Relationships field.
func (o *AuthorizationData) SetRelationships(v AuthorizationDataRelationships) {
	o.Relationships = &v
}

func (o AuthorizationData) MarshalJSON() ([]byte, error) {
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

type NullableAuthorizationData struct {
	value *AuthorizationData
	isSet bool
}

func (v NullableAuthorizationData) Get() *AuthorizationData {
	return v.value
}

func (v *NullableAuthorizationData) Set(val *AuthorizationData) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationData) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationData(val *AuthorizationData) *NullableAuthorizationData {
	return &NullableAuthorizationData{value: val, isSet: true}
}

func (v NullableAuthorizationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
