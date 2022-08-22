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

// GETLineItemOptions200ResponseDataInnerRelationshipsLineItem struct for GETLineItemOptions200ResponseDataInnerRelationshipsLineItem
type GETLineItemOptions200ResponseDataInnerRelationshipsLineItem struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewGETLineItemOptions200ResponseDataInnerRelationshipsLineItem instantiates a new GETLineItemOptions200ResponseDataInnerRelationshipsLineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETLineItemOptions200ResponseDataInnerRelationshipsLineItem(type_ string, id string) *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem {
	this := GETLineItemOptions200ResponseDataInnerRelationshipsLineItem{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGETLineItemOptions200ResponseDataInnerRelationshipsLineItemWithDefaults instantiates a new GETLineItemOptions200ResponseDataInnerRelationshipsLineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETLineItemOptions200ResponseDataInnerRelationshipsLineItemWithDefaults() *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem {
	this := GETLineItemOptions200ResponseDataInnerRelationshipsLineItem{}
	var type_ string = "line_items"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) SetId(v string) {
	o.Id = v
}

func (o GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETLineItemOptions200ResponseDataInnerRelationshipsLineItem struct {
	value *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem
	isSet bool
}

func (v NullableGETLineItemOptions200ResponseDataInnerRelationshipsLineItem) Get() *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem {
	return v.value
}

func (v *NullableGETLineItemOptions200ResponseDataInnerRelationshipsLineItem) Set(val *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGETLineItemOptions200ResponseDataInnerRelationshipsLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGETLineItemOptions200ResponseDataInnerRelationshipsLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETLineItemOptions200ResponseDataInnerRelationshipsLineItem(val *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) *NullableGETLineItemOptions200ResponseDataInnerRelationshipsLineItem {
	return &NullableGETLineItemOptions200ResponseDataInnerRelationshipsLineItem{value: val, isSet: true}
}

func (v NullableGETLineItemOptions200ResponseDataInnerRelationshipsLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETLineItemOptions200ResponseDataInnerRelationshipsLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
