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

// AdyenPayment struct for AdyenPayment
type AdyenPayment struct {
	Data *AdyenPaymentData `json:"data,omitempty"`
}

// NewAdyenPayment instantiates a new AdyenPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenPayment() *AdyenPayment {
	this := AdyenPayment{}
	return &this
}

// NewAdyenPaymentWithDefaults instantiates a new AdyenPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenPaymentWithDefaults() *AdyenPayment {
	this := AdyenPayment{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AdyenPayment) GetData() AdyenPaymentData {
	if o == nil || o.Data == nil {
		var ret AdyenPaymentData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPayment) GetDataOk() (*AdyenPaymentData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AdyenPayment) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given AdyenPaymentData and assigns it to the Data field.
func (o *AdyenPayment) SetData(v AdyenPaymentData) {
	o.Data = &v
}

func (o AdyenPayment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAdyenPayment struct {
	value *AdyenPayment
	isSet bool
}

func (v NullableAdyenPayment) Get() *AdyenPayment {
	return v.value
}

func (v *NullableAdyenPayment) Set(val *AdyenPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenPayment(val *AdyenPayment) *NullableAdyenPayment {
	return &NullableAdyenPayment{value: val, isSet: true}
}

func (v NullableAdyenPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
