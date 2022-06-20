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

// ExternalPaymentCreate struct for ExternalPaymentCreate
type ExternalPaymentCreate struct {
	Data ExternalPaymentCreateData `json:"data"`
}

// NewExternalPaymentCreate instantiates a new ExternalPaymentCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPaymentCreate(data ExternalPaymentCreateData) *ExternalPaymentCreate {
	this := ExternalPaymentCreate{}
	this.Data = data
	return &this
}

// NewExternalPaymentCreateWithDefaults instantiates a new ExternalPaymentCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPaymentCreateWithDefaults() *ExternalPaymentCreate {
	this := ExternalPaymentCreate{}
	return &this
}

// GetData returns the Data field value
func (o *ExternalPaymentCreate) GetData() ExternalPaymentCreateData {
	if o == nil {
		var ret ExternalPaymentCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExternalPaymentCreate) GetDataOk() (*ExternalPaymentCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ExternalPaymentCreate) SetData(v ExternalPaymentCreateData) {
	o.Data = v
}

func (o ExternalPaymentCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableExternalPaymentCreate struct {
	value *ExternalPaymentCreate
	isSet bool
}

func (v NullableExternalPaymentCreate) Get() *ExternalPaymentCreate {
	return v.value
}

func (v *NullableExternalPaymentCreate) Set(val *ExternalPaymentCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPaymentCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPaymentCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPaymentCreate(val *ExternalPaymentCreate) *NullableExternalPaymentCreate {
	return &NullableExternalPaymentCreate{value: val, isSet: true}
}

func (v NullableExternalPaymentCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPaymentCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
