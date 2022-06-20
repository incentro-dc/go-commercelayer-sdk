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

// CustomerDataRelationshipsReturns struct for CustomerDataRelationshipsReturns
type CustomerDataRelationshipsReturns struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewCustomerDataRelationshipsReturns instantiates a new CustomerDataRelationshipsReturns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerDataRelationshipsReturns(type_ string, id string) *CustomerDataRelationshipsReturns {
	this := CustomerDataRelationshipsReturns{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewCustomerDataRelationshipsReturnsWithDefaults instantiates a new CustomerDataRelationshipsReturns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerDataRelationshipsReturnsWithDefaults() *CustomerDataRelationshipsReturns {
	this := CustomerDataRelationshipsReturns{}
	var type_ string = "returns"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *CustomerDataRelationshipsReturns) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerDataRelationshipsReturns) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerDataRelationshipsReturns) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CustomerDataRelationshipsReturns) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomerDataRelationshipsReturns) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomerDataRelationshipsReturns) SetId(v string) {
	o.Id = v
}

func (o CustomerDataRelationshipsReturns) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerDataRelationshipsReturns struct {
	value *CustomerDataRelationshipsReturns
	isSet bool
}

func (v NullableCustomerDataRelationshipsReturns) Get() *CustomerDataRelationshipsReturns {
	return v.value
}

func (v *NullableCustomerDataRelationshipsReturns) Set(val *CustomerDataRelationshipsReturns) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerDataRelationshipsReturns) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerDataRelationshipsReturns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerDataRelationshipsReturns(val *CustomerDataRelationshipsReturns) *NullableCustomerDataRelationshipsReturns {
	return &NullableCustomerDataRelationshipsReturns{value: val, isSet: true}
}

func (v NullableCustomerDataRelationshipsReturns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerDataRelationshipsReturns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}