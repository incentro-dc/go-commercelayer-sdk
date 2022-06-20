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

// AuthorizationDataRelationshipsCaptures struct for AuthorizationDataRelationshipsCaptures
type AuthorizationDataRelationshipsCaptures struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewAuthorizationDataRelationshipsCaptures instantiates a new AuthorizationDataRelationshipsCaptures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationDataRelationshipsCaptures(type_ string, id string) *AuthorizationDataRelationshipsCaptures {
	this := AuthorizationDataRelationshipsCaptures{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAuthorizationDataRelationshipsCapturesWithDefaults instantiates a new AuthorizationDataRelationshipsCaptures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationDataRelationshipsCapturesWithDefaults() *AuthorizationDataRelationshipsCaptures {
	this := AuthorizationDataRelationshipsCaptures{}
	var type_ string = "captures"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *AuthorizationDataRelationshipsCaptures) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuthorizationDataRelationshipsCaptures) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuthorizationDataRelationshipsCaptures) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AuthorizationDataRelationshipsCaptures) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AuthorizationDataRelationshipsCaptures) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AuthorizationDataRelationshipsCaptures) SetId(v string) {
	o.Id = v
}

func (o AuthorizationDataRelationshipsCaptures) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizationDataRelationshipsCaptures struct {
	value *AuthorizationDataRelationshipsCaptures
	isSet bool
}

func (v NullableAuthorizationDataRelationshipsCaptures) Get() *AuthorizationDataRelationshipsCaptures {
	return v.value
}

func (v *NullableAuthorizationDataRelationshipsCaptures) Set(val *AuthorizationDataRelationshipsCaptures) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationDataRelationshipsCaptures) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationDataRelationshipsCaptures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationDataRelationshipsCaptures(val *AuthorizationDataRelationshipsCaptures) *NullableAuthorizationDataRelationshipsCaptures {
	return &NullableAuthorizationDataRelationshipsCaptures{value: val, isSet: true}
}

func (v NullableAuthorizationDataRelationshipsCaptures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationDataRelationshipsCaptures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
