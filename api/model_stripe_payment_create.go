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

// StripePaymentCreate struct for StripePaymentCreate
type StripePaymentCreate struct {
	Data StripePaymentCreateData `json:"data"`
}

// NewStripePaymentCreate instantiates a new StripePaymentCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripePaymentCreate(data StripePaymentCreateData) *StripePaymentCreate {
	this := StripePaymentCreate{}
	this.Data = data
	return &this
}

// NewStripePaymentCreateWithDefaults instantiates a new StripePaymentCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripePaymentCreateWithDefaults() *StripePaymentCreate {
	this := StripePaymentCreate{}
	return &this
}

// GetData returns the Data field value
func (o *StripePaymentCreate) GetData() StripePaymentCreateData {
	if o == nil {
		var ret StripePaymentCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *StripePaymentCreate) GetDataOk() (*StripePaymentCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *StripePaymentCreate) SetData(v StripePaymentCreateData) {
	o.Data = v
}

func (o StripePaymentCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableStripePaymentCreate struct {
	value *StripePaymentCreate
	isSet bool
}

func (v NullableStripePaymentCreate) Get() *StripePaymentCreate {
	return v.value
}

func (v *NullableStripePaymentCreate) Set(val *StripePaymentCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableStripePaymentCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableStripePaymentCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripePaymentCreate(val *StripePaymentCreate) *NullableStripePaymentCreate {
	return &NullableStripePaymentCreate{value: val, isSet: true}
}

func (v NullableStripePaymentCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripePaymentCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}