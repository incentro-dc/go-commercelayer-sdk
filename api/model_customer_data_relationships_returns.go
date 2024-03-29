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

// checks if the CustomerDataRelationshipsReturns type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerDataRelationshipsReturns{}

// CustomerDataRelationshipsReturns struct for CustomerDataRelationshipsReturns
type CustomerDataRelationshipsReturns struct {
	Data *CustomerDataRelationshipsReturnsData `json:"data,omitempty"`
}

// NewCustomerDataRelationshipsReturns instantiates a new CustomerDataRelationshipsReturns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerDataRelationshipsReturns() *CustomerDataRelationshipsReturns {
	this := CustomerDataRelationshipsReturns{}
	return &this
}

// NewCustomerDataRelationshipsReturnsWithDefaults instantiates a new CustomerDataRelationshipsReturns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerDataRelationshipsReturnsWithDefaults() *CustomerDataRelationshipsReturns {
	this := CustomerDataRelationshipsReturns{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CustomerDataRelationshipsReturns) GetData() CustomerDataRelationshipsReturnsData {
	if o == nil || IsNil(o.Data) {
		var ret CustomerDataRelationshipsReturnsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerDataRelationshipsReturns) GetDataOk() (*CustomerDataRelationshipsReturnsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CustomerDataRelationshipsReturns) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CustomerDataRelationshipsReturnsData and assigns it to the Data field.
func (o *CustomerDataRelationshipsReturns) SetData(v CustomerDataRelationshipsReturnsData) {
	o.Data = &v
}

func (o CustomerDataRelationshipsReturns) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerDataRelationshipsReturns) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
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
