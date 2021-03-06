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

// StripePayment struct for StripePayment
type StripePayment struct {
	Data StripePaymentData `json:"data"`
}

// NewStripePayment instantiates a new StripePayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripePayment(data StripePaymentData) *StripePayment {
	this := StripePayment{}
	this.Data = data
	return &this
}

// NewStripePaymentWithDefaults instantiates a new StripePayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripePaymentWithDefaults() *StripePayment {
	this := StripePayment{}
	return &this
}

// GetData returns the Data field value
func (o *StripePayment) GetData() StripePaymentData {
	if o == nil {
		var ret StripePaymentData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *StripePayment) GetDataOk() (*StripePaymentData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *StripePayment) SetData(v StripePaymentData) {
	o.Data = v
}

func (o StripePayment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableStripePayment struct {
	value *StripePayment
	isSet bool
}

func (v NullableStripePayment) Get() *StripePayment {
	return v.value
}

func (v *NullableStripePayment) Set(val *StripePayment) {
	v.value = val
	v.isSet = true
}

func (v NullableStripePayment) IsSet() bool {
	return v.isSet
}

func (v *NullableStripePayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripePayment(val *StripePayment) *NullableStripePayment {
	return &NullableStripePayment{value: val, isSet: true}
}

func (v NullableStripePayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripePayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
