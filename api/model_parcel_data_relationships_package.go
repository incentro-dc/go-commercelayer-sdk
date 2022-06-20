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

// ParcelDataRelationshipsPackage struct for ParcelDataRelationshipsPackage
type ParcelDataRelationshipsPackage struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewParcelDataRelationshipsPackage instantiates a new ParcelDataRelationshipsPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelDataRelationshipsPackage(type_ string, id string) *ParcelDataRelationshipsPackage {
	this := ParcelDataRelationshipsPackage{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewParcelDataRelationshipsPackageWithDefaults instantiates a new ParcelDataRelationshipsPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelDataRelationshipsPackageWithDefaults() *ParcelDataRelationshipsPackage {
	this := ParcelDataRelationshipsPackage{}
	var type_ string = "packages"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ParcelDataRelationshipsPackage) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ParcelDataRelationshipsPackage) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ParcelDataRelationshipsPackage) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ParcelDataRelationshipsPackage) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ParcelDataRelationshipsPackage) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ParcelDataRelationshipsPackage) SetId(v string) {
	o.Id = v
}

func (o ParcelDataRelationshipsPackage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableParcelDataRelationshipsPackage struct {
	value *ParcelDataRelationshipsPackage
	isSet bool
}

func (v NullableParcelDataRelationshipsPackage) Get() *ParcelDataRelationshipsPackage {
	return v.value
}

func (v *NullableParcelDataRelationshipsPackage) Set(val *ParcelDataRelationshipsPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelDataRelationshipsPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelDataRelationshipsPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelDataRelationshipsPackage(val *ParcelDataRelationshipsPackage) *NullableParcelDataRelationshipsPackage {
	return &NullableParcelDataRelationshipsPackage{value: val, isSet: true}
}

func (v NullableParcelDataRelationshipsPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelDataRelationshipsPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}