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

// CustomerDataRelationshipsCustomerGroup struct for CustomerDataRelationshipsCustomerGroup
type CustomerDataRelationshipsCustomerGroup struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewCustomerDataRelationshipsCustomerGroup instantiates a new CustomerDataRelationshipsCustomerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerDataRelationshipsCustomerGroup(type_ string, id string) *CustomerDataRelationshipsCustomerGroup {
	this := CustomerDataRelationshipsCustomerGroup{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewCustomerDataRelationshipsCustomerGroupWithDefaults instantiates a new CustomerDataRelationshipsCustomerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerDataRelationshipsCustomerGroupWithDefaults() *CustomerDataRelationshipsCustomerGroup {
	this := CustomerDataRelationshipsCustomerGroup{}
	var type_ string = "customer_groups"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *CustomerDataRelationshipsCustomerGroup) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerDataRelationshipsCustomerGroup) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerDataRelationshipsCustomerGroup) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CustomerDataRelationshipsCustomerGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomerDataRelationshipsCustomerGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomerDataRelationshipsCustomerGroup) SetId(v string) {
	o.Id = v
}

func (o CustomerDataRelationshipsCustomerGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerDataRelationshipsCustomerGroup struct {
	value *CustomerDataRelationshipsCustomerGroup
	isSet bool
}

func (v NullableCustomerDataRelationshipsCustomerGroup) Get() *CustomerDataRelationshipsCustomerGroup {
	return v.value
}

func (v *NullableCustomerDataRelationshipsCustomerGroup) Set(val *CustomerDataRelationshipsCustomerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerDataRelationshipsCustomerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerDataRelationshipsCustomerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerDataRelationshipsCustomerGroup(val *CustomerDataRelationshipsCustomerGroup) *NullableCustomerDataRelationshipsCustomerGroup {
	return &NullableCustomerDataRelationshipsCustomerGroup{value: val, isSet: true}
}

func (v NullableCustomerDataRelationshipsCustomerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerDataRelationshipsCustomerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
