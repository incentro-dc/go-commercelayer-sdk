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

// GETMarkets200ResponseDataInnerRelationshipsMerchant struct for GETMarkets200ResponseDataInnerRelationshipsMerchant
type GETMarkets200ResponseDataInnerRelationshipsMerchant struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewGETMarkets200ResponseDataInnerRelationshipsMerchant instantiates a new GETMarkets200ResponseDataInnerRelationshipsMerchant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETMarkets200ResponseDataInnerRelationshipsMerchant(type_ string, id string) *GETMarkets200ResponseDataInnerRelationshipsMerchant {
	this := GETMarkets200ResponseDataInnerRelationshipsMerchant{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGETMarkets200ResponseDataInnerRelationshipsMerchantWithDefaults instantiates a new GETMarkets200ResponseDataInnerRelationshipsMerchant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETMarkets200ResponseDataInnerRelationshipsMerchantWithDefaults() *GETMarkets200ResponseDataInnerRelationshipsMerchant {
	this := GETMarkets200ResponseDataInnerRelationshipsMerchant{}
	var type_ string = "merchants"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchant) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchant) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchant) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchant) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchant) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GETMarkets200ResponseDataInnerRelationshipsMerchant) SetId(v string) {
	o.Id = v
}

func (o GETMarkets200ResponseDataInnerRelationshipsMerchant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETMarkets200ResponseDataInnerRelationshipsMerchant struct {
	value *GETMarkets200ResponseDataInnerRelationshipsMerchant
	isSet bool
}

func (v NullableGETMarkets200ResponseDataInnerRelationshipsMerchant) Get() *GETMarkets200ResponseDataInnerRelationshipsMerchant {
	return v.value
}

func (v *NullableGETMarkets200ResponseDataInnerRelationshipsMerchant) Set(val *GETMarkets200ResponseDataInnerRelationshipsMerchant) {
	v.value = val
	v.isSet = true
}

func (v NullableGETMarkets200ResponseDataInnerRelationshipsMerchant) IsSet() bool {
	return v.isSet
}

func (v *NullableGETMarkets200ResponseDataInnerRelationshipsMerchant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETMarkets200ResponseDataInnerRelationshipsMerchant(val *GETMarkets200ResponseDataInnerRelationshipsMerchant) *NullableGETMarkets200ResponseDataInnerRelationshipsMerchant {
	return &NullableGETMarkets200ResponseDataInnerRelationshipsMerchant{value: val, isSet: true}
}

func (v NullableGETMarkets200ResponseDataInnerRelationshipsMerchant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETMarkets200ResponseDataInnerRelationshipsMerchant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}