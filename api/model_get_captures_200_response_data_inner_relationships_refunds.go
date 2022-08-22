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

// GETCaptures200ResponseDataInnerRelationshipsRefunds struct for GETCaptures200ResponseDataInnerRelationshipsRefunds
type GETCaptures200ResponseDataInnerRelationshipsRefunds struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewGETCaptures200ResponseDataInnerRelationshipsRefunds instantiates a new GETCaptures200ResponseDataInnerRelationshipsRefunds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCaptures200ResponseDataInnerRelationshipsRefunds(type_ string, id string) *GETCaptures200ResponseDataInnerRelationshipsRefunds {
	this := GETCaptures200ResponseDataInnerRelationshipsRefunds{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGETCaptures200ResponseDataInnerRelationshipsRefundsWithDefaults instantiates a new GETCaptures200ResponseDataInnerRelationshipsRefunds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCaptures200ResponseDataInnerRelationshipsRefundsWithDefaults() *GETCaptures200ResponseDataInnerRelationshipsRefunds {
	this := GETCaptures200ResponseDataInnerRelationshipsRefunds{}
	var type_ string = "refunds"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *GETCaptures200ResponseDataInnerRelationshipsRefunds) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GETCaptures200ResponseDataInnerRelationshipsRefunds) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GETCaptures200ResponseDataInnerRelationshipsRefunds) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GETCaptures200ResponseDataInnerRelationshipsRefunds) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GETCaptures200ResponseDataInnerRelationshipsRefunds) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GETCaptures200ResponseDataInnerRelationshipsRefunds) SetId(v string) {
	o.Id = v
}

func (o GETCaptures200ResponseDataInnerRelationshipsRefunds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETCaptures200ResponseDataInnerRelationshipsRefunds struct {
	value *GETCaptures200ResponseDataInnerRelationshipsRefunds
	isSet bool
}

func (v NullableGETCaptures200ResponseDataInnerRelationshipsRefunds) Get() *GETCaptures200ResponseDataInnerRelationshipsRefunds {
	return v.value
}

func (v *NullableGETCaptures200ResponseDataInnerRelationshipsRefunds) Set(val *GETCaptures200ResponseDataInnerRelationshipsRefunds) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCaptures200ResponseDataInnerRelationshipsRefunds) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCaptures200ResponseDataInnerRelationshipsRefunds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCaptures200ResponseDataInnerRelationshipsRefunds(val *GETCaptures200ResponseDataInnerRelationshipsRefunds) *NullableGETCaptures200ResponseDataInnerRelationshipsRefunds {
	return &NullableGETCaptures200ResponseDataInnerRelationshipsRefunds{value: val, isSet: true}
}

func (v NullableGETCaptures200ResponseDataInnerRelationshipsRefunds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCaptures200ResponseDataInnerRelationshipsRefunds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
