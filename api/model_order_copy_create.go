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

// OrderCopyCreate struct for OrderCopyCreate
type OrderCopyCreate struct {
	Data OrderCopyCreateData `json:"data"`
}

// NewOrderCopyCreate instantiates a new OrderCopyCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCopyCreate(data OrderCopyCreateData) *OrderCopyCreate {
	this := OrderCopyCreate{}
	this.Data = data
	return &this
}

// NewOrderCopyCreateWithDefaults instantiates a new OrderCopyCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCopyCreateWithDefaults() *OrderCopyCreate {
	this := OrderCopyCreate{}
	return &this
}

// GetData returns the Data field value
func (o *OrderCopyCreate) GetData() OrderCopyCreateData {
	if o == nil {
		var ret OrderCopyCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OrderCopyCreate) GetDataOk() (*OrderCopyCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OrderCopyCreate) SetData(v OrderCopyCreateData) {
	o.Data = v
}

func (o OrderCopyCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableOrderCopyCreate struct {
	value *OrderCopyCreate
	isSet bool
}

func (v NullableOrderCopyCreate) Get() *OrderCopyCreate {
	return v.value
}

func (v *NullableOrderCopyCreate) Set(val *OrderCopyCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCopyCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCopyCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCopyCreate(val *OrderCopyCreate) *NullableOrderCopyCreate {
	return &NullableOrderCopyCreate{value: val, isSet: true}
}

func (v NullableOrderCopyCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCopyCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
