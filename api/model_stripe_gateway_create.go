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

// StripeGatewayCreate struct for StripeGatewayCreate
type StripeGatewayCreate struct {
	Data StripeGatewayCreateData `json:"data"`
}

// NewStripeGatewayCreate instantiates a new StripeGatewayCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripeGatewayCreate(data StripeGatewayCreateData) *StripeGatewayCreate {
	this := StripeGatewayCreate{}
	this.Data = data
	return &this
}

// NewStripeGatewayCreateWithDefaults instantiates a new StripeGatewayCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripeGatewayCreateWithDefaults() *StripeGatewayCreate {
	this := StripeGatewayCreate{}
	return &this
}

// GetData returns the Data field value
func (o *StripeGatewayCreate) GetData() StripeGatewayCreateData {
	if o == nil {
		var ret StripeGatewayCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *StripeGatewayCreate) GetDataOk() (*StripeGatewayCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *StripeGatewayCreate) SetData(v StripeGatewayCreateData) {
	o.Data = v
}

func (o StripeGatewayCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableStripeGatewayCreate struct {
	value *StripeGatewayCreate
	isSet bool
}

func (v NullableStripeGatewayCreate) Get() *StripeGatewayCreate {
	return v.value
}

func (v *NullableStripeGatewayCreate) Set(val *StripeGatewayCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeGatewayCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeGatewayCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeGatewayCreate(val *StripeGatewayCreate) *NullableStripeGatewayCreate {
	return &NullableStripeGatewayCreate{value: val, isSet: true}
}

func (v NullableStripeGatewayCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeGatewayCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
