/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CustomerCreateDataRelationshipsCustomerGroup struct for CustomerCreateDataRelationshipsCustomerGroup
type CustomerCreateDataRelationshipsCustomerGroup struct {
	Data CustomerDataRelationshipsCustomerGroupData `json:"data"`
}

// NewCustomerCreateDataRelationshipsCustomerGroup instantiates a new CustomerCreateDataRelationshipsCustomerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerCreateDataRelationshipsCustomerGroup(data CustomerDataRelationshipsCustomerGroupData) *CustomerCreateDataRelationshipsCustomerGroup {
	this := CustomerCreateDataRelationshipsCustomerGroup{}
	this.Data = data
	return &this
}

// NewCustomerCreateDataRelationshipsCustomerGroupWithDefaults instantiates a new CustomerCreateDataRelationshipsCustomerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerCreateDataRelationshipsCustomerGroupWithDefaults() *CustomerCreateDataRelationshipsCustomerGroup {
	this := CustomerCreateDataRelationshipsCustomerGroup{}
	return &this
}

// GetData returns the Data field value
func (o *CustomerCreateDataRelationshipsCustomerGroup) GetData() CustomerDataRelationshipsCustomerGroupData {
	if o == nil {
		var ret CustomerDataRelationshipsCustomerGroupData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomerCreateDataRelationshipsCustomerGroup) GetDataOk() (*CustomerDataRelationshipsCustomerGroupData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomerCreateDataRelationshipsCustomerGroup) SetData(v CustomerDataRelationshipsCustomerGroupData) {
	o.Data = v
}

func (o CustomerCreateDataRelationshipsCustomerGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerCreateDataRelationshipsCustomerGroup struct {
	value *CustomerCreateDataRelationshipsCustomerGroup
	isSet bool
}

func (v NullableCustomerCreateDataRelationshipsCustomerGroup) Get() *CustomerCreateDataRelationshipsCustomerGroup {
	return v.value
}

func (v *NullableCustomerCreateDataRelationshipsCustomerGroup) Set(val *CustomerCreateDataRelationshipsCustomerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerCreateDataRelationshipsCustomerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerCreateDataRelationshipsCustomerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerCreateDataRelationshipsCustomerGroup(val *CustomerCreateDataRelationshipsCustomerGroup) *NullableCustomerCreateDataRelationshipsCustomerGroup {
	return &NullableCustomerCreateDataRelationshipsCustomerGroup{value: val, isSet: true}
}

func (v NullableCustomerCreateDataRelationshipsCustomerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerCreateDataRelationshipsCustomerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
