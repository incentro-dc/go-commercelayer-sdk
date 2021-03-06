/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LineItemDataRelationshipsLineItemOptions struct for LineItemDataRelationshipsLineItemOptions
type LineItemDataRelationshipsLineItemOptions struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewLineItemDataRelationshipsLineItemOptions instantiates a new LineItemDataRelationshipsLineItemOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemDataRelationshipsLineItemOptions(type_ string, id string) *LineItemDataRelationshipsLineItemOptions {
	this := LineItemDataRelationshipsLineItemOptions{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewLineItemDataRelationshipsLineItemOptionsWithDefaults instantiates a new LineItemDataRelationshipsLineItemOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemDataRelationshipsLineItemOptionsWithDefaults() *LineItemDataRelationshipsLineItemOptions {
	this := LineItemDataRelationshipsLineItemOptions{}
	var type_ string = "line_item_options"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *LineItemDataRelationshipsLineItemOptions) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LineItemDataRelationshipsLineItemOptions) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LineItemDataRelationshipsLineItemOptions) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *LineItemDataRelationshipsLineItemOptions) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LineItemDataRelationshipsLineItemOptions) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LineItemDataRelationshipsLineItemOptions) SetId(v string) {
	o.Id = v
}

func (o LineItemDataRelationshipsLineItemOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableLineItemDataRelationshipsLineItemOptions struct {
	value *LineItemDataRelationshipsLineItemOptions
	isSet bool
}

func (v NullableLineItemDataRelationshipsLineItemOptions) Get() *LineItemDataRelationshipsLineItemOptions {
	return v.value
}

func (v *NullableLineItemDataRelationshipsLineItemOptions) Set(val *LineItemDataRelationshipsLineItemOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemDataRelationshipsLineItemOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemDataRelationshipsLineItemOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemDataRelationshipsLineItemOptions(val *LineItemDataRelationshipsLineItemOptions) *NullableLineItemDataRelationshipsLineItemOptions {
	return &NullableLineItemDataRelationshipsLineItemOptions{value: val, isSet: true}
}

func (v NullableLineItemDataRelationshipsLineItemOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemDataRelationshipsLineItemOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
