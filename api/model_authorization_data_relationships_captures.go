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

// AuthorizationDataRelationshipsCaptures struct for AuthorizationDataRelationshipsCaptures
type AuthorizationDataRelationshipsCaptures struct {
	Data *AuthorizationDataRelationshipsCapturesData `json:"data,omitempty"`
}

// NewAuthorizationDataRelationshipsCaptures instantiates a new AuthorizationDataRelationshipsCaptures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationDataRelationshipsCaptures() *AuthorizationDataRelationshipsCaptures {
	this := AuthorizationDataRelationshipsCaptures{}
	return &this
}

// NewAuthorizationDataRelationshipsCapturesWithDefaults instantiates a new AuthorizationDataRelationshipsCaptures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationDataRelationshipsCapturesWithDefaults() *AuthorizationDataRelationshipsCaptures {
	this := AuthorizationDataRelationshipsCaptures{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AuthorizationDataRelationshipsCaptures) GetData() AuthorizationDataRelationshipsCapturesData {
	if o == nil || o.Data == nil {
		var ret AuthorizationDataRelationshipsCapturesData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDataRelationshipsCaptures) GetDataOk() (*AuthorizationDataRelationshipsCapturesData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AuthorizationDataRelationshipsCaptures) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given AuthorizationDataRelationshipsCapturesData and assigns it to the Data field.
func (o *AuthorizationDataRelationshipsCaptures) SetData(v AuthorizationDataRelationshipsCapturesData) {
	o.Data = &v
}

func (o AuthorizationDataRelationshipsCaptures) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
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
