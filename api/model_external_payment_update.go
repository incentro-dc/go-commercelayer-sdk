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

// ExternalPaymentUpdate struct for ExternalPaymentUpdate
type ExternalPaymentUpdate struct {
	Data ExternalPaymentUpdateData `json:"data"`
}

// NewExternalPaymentUpdate instantiates a new ExternalPaymentUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPaymentUpdate(data ExternalPaymentUpdateData) *ExternalPaymentUpdate {
	this := ExternalPaymentUpdate{}
	this.Data = data
	return &this
}

// NewExternalPaymentUpdateWithDefaults instantiates a new ExternalPaymentUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPaymentUpdateWithDefaults() *ExternalPaymentUpdate {
	this := ExternalPaymentUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *ExternalPaymentUpdate) GetData() ExternalPaymentUpdateData {
	if o == nil {
		var ret ExternalPaymentUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExternalPaymentUpdate) GetDataOk() (*ExternalPaymentUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ExternalPaymentUpdate) SetData(v ExternalPaymentUpdateData) {
	o.Data = v
}

func (o ExternalPaymentUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableExternalPaymentUpdate struct {
	value *ExternalPaymentUpdate
	isSet bool
}

func (v NullableExternalPaymentUpdate) Get() *ExternalPaymentUpdate {
	return v.value
}

func (v *NullableExternalPaymentUpdate) Set(val *ExternalPaymentUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPaymentUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPaymentUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPaymentUpdate(val *ExternalPaymentUpdate) *NullableExternalPaymentUpdate {
	return &NullableExternalPaymentUpdate{value: val, isSet: true}
}

func (v NullableExternalPaymentUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPaymentUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
