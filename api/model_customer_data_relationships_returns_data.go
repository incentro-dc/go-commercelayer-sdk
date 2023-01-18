/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CustomerDataRelationshipsReturnsData struct for CustomerDataRelationshipsReturnsData
type CustomerDataRelationshipsReturnsData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource's id
	Id *string `json:"id,omitempty"`
}

// NewCustomerDataRelationshipsReturnsData instantiates a new CustomerDataRelationshipsReturnsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerDataRelationshipsReturnsData() *CustomerDataRelationshipsReturnsData {
	this := CustomerDataRelationshipsReturnsData{}
	return &this
}

// NewCustomerDataRelationshipsReturnsDataWithDefaults instantiates a new CustomerDataRelationshipsReturnsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerDataRelationshipsReturnsDataWithDefaults() *CustomerDataRelationshipsReturnsData {
	this := CustomerDataRelationshipsReturnsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CustomerDataRelationshipsReturnsData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerDataRelationshipsReturnsData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomerDataRelationshipsReturnsData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CustomerDataRelationshipsReturnsData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomerDataRelationshipsReturnsData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerDataRelationshipsReturnsData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerDataRelationshipsReturnsData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomerDataRelationshipsReturnsData) SetId(v string) {
	o.Id = &v
}

func (o CustomerDataRelationshipsReturnsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerDataRelationshipsReturnsData struct {
	value *CustomerDataRelationshipsReturnsData
	isSet bool
}

func (v NullableCustomerDataRelationshipsReturnsData) Get() *CustomerDataRelationshipsReturnsData {
	return v.value
}

func (v *NullableCustomerDataRelationshipsReturnsData) Set(val *CustomerDataRelationshipsReturnsData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerDataRelationshipsReturnsData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerDataRelationshipsReturnsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerDataRelationshipsReturnsData(val *CustomerDataRelationshipsReturnsData) *NullableCustomerDataRelationshipsReturnsData {
	return &NullableCustomerDataRelationshipsReturnsData{value: val, isSet: true}
}

func (v NullableCustomerDataRelationshipsReturnsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerDataRelationshipsReturnsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
