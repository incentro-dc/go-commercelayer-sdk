/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETReturns200ResponseDataInnerRelationshipsOriginAddressData struct for GETReturns200ResponseDataInnerRelationshipsOriginAddressData
type GETReturns200ResponseDataInnerRelationshipsOriginAddressData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETReturns200ResponseDataInnerRelationshipsOriginAddressData instantiates a new GETReturns200ResponseDataInnerRelationshipsOriginAddressData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETReturns200ResponseDataInnerRelationshipsOriginAddressData() *GETReturns200ResponseDataInnerRelationshipsOriginAddressData {
	this := GETReturns200ResponseDataInnerRelationshipsOriginAddressData{}
	return &this
}

// NewGETReturns200ResponseDataInnerRelationshipsOriginAddressDataWithDefaults instantiates a new GETReturns200ResponseDataInnerRelationshipsOriginAddressData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETReturns200ResponseDataInnerRelationshipsOriginAddressDataWithDefaults() *GETReturns200ResponseDataInnerRelationshipsOriginAddressData {
	this := GETReturns200ResponseDataInnerRelationshipsOriginAddressData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETReturns200ResponseDataInnerRelationshipsOriginAddressData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturns200ResponseDataInnerRelationshipsOriginAddressData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETReturns200ResponseDataInnerRelationshipsOriginAddressData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETReturns200ResponseDataInnerRelationshipsOriginAddressData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETReturns200ResponseDataInnerRelationshipsOriginAddressData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturns200ResponseDataInnerRelationshipsOriginAddressData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETReturns200ResponseDataInnerRelationshipsOriginAddressData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETReturns200ResponseDataInnerRelationshipsOriginAddressData) SetId(v string) {
	o.Id = &v
}

func (o GETReturns200ResponseDataInnerRelationshipsOriginAddressData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETReturns200ResponseDataInnerRelationshipsOriginAddressData struct {
	value *GETReturns200ResponseDataInnerRelationshipsOriginAddressData
	isSet bool
}

func (v NullableGETReturns200ResponseDataInnerRelationshipsOriginAddressData) Get() *GETReturns200ResponseDataInnerRelationshipsOriginAddressData {
	return v.value
}

func (v *NullableGETReturns200ResponseDataInnerRelationshipsOriginAddressData) Set(val *GETReturns200ResponseDataInnerRelationshipsOriginAddressData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETReturns200ResponseDataInnerRelationshipsOriginAddressData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETReturns200ResponseDataInnerRelationshipsOriginAddressData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETReturns200ResponseDataInnerRelationshipsOriginAddressData(val *GETReturns200ResponseDataInnerRelationshipsOriginAddressData) *NullableGETReturns200ResponseDataInnerRelationshipsOriginAddressData {
	return &NullableGETReturns200ResponseDataInnerRelationshipsOriginAddressData{value: val, isSet: true}
}

func (v NullableGETReturns200ResponseDataInnerRelationshipsOriginAddressData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETReturns200ResponseDataInnerRelationshipsOriginAddressData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}