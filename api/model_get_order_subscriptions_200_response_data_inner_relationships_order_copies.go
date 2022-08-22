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

// GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies struct for GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies
type GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewGETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies instantiates a new GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies(type_ string, id string) *GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies {
	this := GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopiesWithDefaults instantiates a new GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopiesWithDefaults() *GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies {
	this := GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies{}
	var type_ string = "order_copies"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies) SetId(v string) {
	o.Id = v
}

func (o GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies struct {
	value *GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies
	isSet bool
}

func (v NullableGETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies) Get() *GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies {
	return v.value
}

func (v *NullableGETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies) Set(val *GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies(val *GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies) *NullableGETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies {
	return &NullableGETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies{value: val, isSet: true}
}

func (v NullableGETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}