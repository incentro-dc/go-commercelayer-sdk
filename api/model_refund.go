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

// checks if the Refund type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Refund{}

// Refund struct for Refund
type Refund struct {
	Data *RefundData `json:"data,omitempty"`
}

// NewRefund instantiates a new Refund object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefund() *Refund {
	this := Refund{}
	return &this
}

// NewRefundWithDefaults instantiates a new Refund object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundWithDefaults() *Refund {
	this := Refund{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Refund) GetData() RefundData {
	if o == nil || IsNil(o.Data) {
		var ret RefundData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetDataOk() (*RefundData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Refund) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given RefundData and assigns it to the Data field.
func (o *Refund) SetData(v RefundData) {
	o.Data = &v
}

func (o Refund) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Refund) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableRefund struct {
	value *Refund
	isSet bool
}

func (v NullableRefund) Get() *Refund {
	return v.value
}

func (v *NullableRefund) Set(val *Refund) {
	v.value = val
	v.isSet = true
}

func (v NullableRefund) IsSet() bool {
	return v.isSet
}

func (v *NullableRefund) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefund(val *Refund) *NullableRefund {
	return &NullableRefund{value: val, isSet: true}
}

func (v NullableRefund) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefund) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
