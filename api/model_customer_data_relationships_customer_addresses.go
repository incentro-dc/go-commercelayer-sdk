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

// CustomerDataRelationshipsCustomerAddresses struct for CustomerDataRelationshipsCustomerAddresses
type CustomerDataRelationshipsCustomerAddresses struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewCustomerDataRelationshipsCustomerAddresses instantiates a new CustomerDataRelationshipsCustomerAddresses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerDataRelationshipsCustomerAddresses(type_ string, id string) *CustomerDataRelationshipsCustomerAddresses {
	this := CustomerDataRelationshipsCustomerAddresses{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewCustomerDataRelationshipsCustomerAddressesWithDefaults instantiates a new CustomerDataRelationshipsCustomerAddresses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerDataRelationshipsCustomerAddressesWithDefaults() *CustomerDataRelationshipsCustomerAddresses {
	this := CustomerDataRelationshipsCustomerAddresses{}
	var type_ string = "customer_addresses"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *CustomerDataRelationshipsCustomerAddresses) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerDataRelationshipsCustomerAddresses) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerDataRelationshipsCustomerAddresses) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CustomerDataRelationshipsCustomerAddresses) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomerDataRelationshipsCustomerAddresses) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomerDataRelationshipsCustomerAddresses) SetId(v string) {
	o.Id = v
}

func (o CustomerDataRelationshipsCustomerAddresses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerDataRelationshipsCustomerAddresses struct {
	value *CustomerDataRelationshipsCustomerAddresses
	isSet bool
}

func (v NullableCustomerDataRelationshipsCustomerAddresses) Get() *CustomerDataRelationshipsCustomerAddresses {
	return v.value
}

func (v *NullableCustomerDataRelationshipsCustomerAddresses) Set(val *CustomerDataRelationshipsCustomerAddresses) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerDataRelationshipsCustomerAddresses) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerDataRelationshipsCustomerAddresses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerDataRelationshipsCustomerAddresses(val *CustomerDataRelationshipsCustomerAddresses) *NullableCustomerDataRelationshipsCustomerAddresses {
	return &NullableCustomerDataRelationshipsCustomerAddresses{value: val, isSet: true}
}

func (v NullableCustomerDataRelationshipsCustomerAddresses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerDataRelationshipsCustomerAddresses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
