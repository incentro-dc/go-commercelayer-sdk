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

// GETAuthorizations200ResponseDataInnerRelationshipsVoids struct for GETAuthorizations200ResponseDataInnerRelationshipsVoids
type GETAuthorizations200ResponseDataInnerRelationshipsVoids struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewGETAuthorizations200ResponseDataInnerRelationshipsVoids instantiates a new GETAuthorizations200ResponseDataInnerRelationshipsVoids object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAuthorizations200ResponseDataInnerRelationshipsVoids(type_ string, id string) *GETAuthorizations200ResponseDataInnerRelationshipsVoids {
	this := GETAuthorizations200ResponseDataInnerRelationshipsVoids{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGETAuthorizations200ResponseDataInnerRelationshipsVoidsWithDefaults instantiates a new GETAuthorizations200ResponseDataInnerRelationshipsVoids object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAuthorizations200ResponseDataInnerRelationshipsVoidsWithDefaults() *GETAuthorizations200ResponseDataInnerRelationshipsVoids {
	this := GETAuthorizations200ResponseDataInnerRelationshipsVoids{}
	var type_ string = "voids"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *GETAuthorizations200ResponseDataInnerRelationshipsVoids) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsVoids) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GETAuthorizations200ResponseDataInnerRelationshipsVoids) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GETAuthorizations200ResponseDataInnerRelationshipsVoids) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GETAuthorizations200ResponseDataInnerRelationshipsVoids) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GETAuthorizations200ResponseDataInnerRelationshipsVoids) SetId(v string) {
	o.Id = v
}

func (o GETAuthorizations200ResponseDataInnerRelationshipsVoids) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETAuthorizations200ResponseDataInnerRelationshipsVoids struct {
	value *GETAuthorizations200ResponseDataInnerRelationshipsVoids
	isSet bool
}

func (v NullableGETAuthorizations200ResponseDataInnerRelationshipsVoids) Get() *GETAuthorizations200ResponseDataInnerRelationshipsVoids {
	return v.value
}

func (v *NullableGETAuthorizations200ResponseDataInnerRelationshipsVoids) Set(val *GETAuthorizations200ResponseDataInnerRelationshipsVoids) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAuthorizations200ResponseDataInnerRelationshipsVoids) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAuthorizations200ResponseDataInnerRelationshipsVoids) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAuthorizations200ResponseDataInnerRelationshipsVoids(val *GETAuthorizations200ResponseDataInnerRelationshipsVoids) *NullableGETAuthorizations200ResponseDataInnerRelationshipsVoids {
	return &NullableGETAuthorizations200ResponseDataInnerRelationshipsVoids{value: val, isSet: true}
}

func (v NullableGETAuthorizations200ResponseDataInnerRelationshipsVoids) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAuthorizations200ResponseDataInnerRelationshipsVoids) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
